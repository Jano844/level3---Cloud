// main.go
package main

import (
	"context"
	"crypto/rsa"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"os"
	"strings"
	"time"

	"myapp/myDatabase"

	"github.com/golang-jwt/jwt/v5"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

// CORS middleware
func enableCORS(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set CORS headers
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With")
		w.Header().Set("Access-Control-Max-Age", "86400")

		// Handle preflight requests
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		// Call the next handler
		next.ServeHTTP(w, r)
	})
}

func main() {
	// Initialize Zitadel JWT validation
	if err := initZitadelConfig(); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to initialize Zitadel config: %v\n", err)
		fmt.Fprintf(os.Stderr, "Make sure ZITADEL_ISSUER environment variable is set\n")
		os.Exit(1)
	}

	config, err := rest.InClusterConfig()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to get in-cluster config: %v\n", err)
		os.Exit(1)
	}

	var clientset *kubernetes.Clientset
	clientset, err = kubernetes.NewForConfig(config)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create clientset: %v\n", err)
		os.Exit(1)
	}

	// Health check endpoint (no JWT required)
	http.HandleFunc("/", enableCORS(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "jano844/level3-test:v1.2.10\n")
	}))

	// Protected endpoints (JWT required)
	http.HandleFunc("/pods", enableCORS(jwtAuthMiddleware(func(w http.ResponseWriter, r *http.Request) {
		handlePods(w, r, clientset)
	})))

	http.HandleFunc("/createDB", enableCORS(jwtAuthMiddleware(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			myDatabase.CreateNewDatabase(w, r, clientset)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})))

	http.HandleFunc("/deleteDB", enableCORS(jwtAuthMiddleware(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodDelete {
			myDatabase.DeleteDatabase(w, r, clientset)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})))

	http.HandleFunc("/getDBAccess", enableCORS(jwtAuthMiddleware(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			myDatabase.GetDatabaseAccess(w, r, clientset)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})))

	http.HandleFunc("/getUserDatabases", enableCORS(jwtAuthMiddleware(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			myDatabase.GetUserDatabases(w, r, clientset)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})))

	fmt.Println("Server starting with JWT authentication enabled...")
	http.ListenAndServe(":8080", nil)
}

// /pods Handler
func handlePods(w http.ResponseWriter, r *http.Request, clientset *kubernetes.Clientset) {
	// List pods from all namespaces using empty string ""
	pods, err := clientset.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		http.Error(w, "Failed to list pods: "+err.Error(), http.StatusInternalServerError)
		return
	}

	type PodInfo struct {
		Name      string `json:"name"`
		Namespace string `json:"namespace"`
		Status    string `json:"status"`
	}

	var result []PodInfo
	for _, pod := range pods.Items {
		result = append(result, PodInfo{
			Name:      pod.Name,
			Namespace: pod.Namespace,
			Status:    string(pod.Status.Phase),
		})
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

// JWT validation structures
type ZitadelConfig struct {
	Issuer  string `json:"issuer"`
	JWKSURI string `json:"jwks_uri"`
}

type JWKS struct {
	Keys []JWK `json:"keys"`
}

type JWK struct {
	Kty string `json:"kty"`
	Kid string `json:"kid"`
	Use string `json:"use"`
	N   string `json:"n"`
	E   string `json:"e"`
}

type Claims struct {
	jwt.RegisteredClaims
	Sub   string `json:"sub"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

// Global variables for caching
var (
	zitadelConfig *ZitadelConfig
	jwksKeys      map[string]*rsa.PublicKey
	lastJWKSFetch time.Time
)

// JWT validation middleware
func jwtAuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Skip JWT validation for health check endpoint
		if r.URL.Path == "/" {
			next.ServeHTTP(w, r)
			return
		}

		// Get Authorization header
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Missing Authorization header", http.StatusUnauthorized)
			return
		}

		// Check if it starts with "Bearer "
		if !strings.HasPrefix(authHeader, "Bearer ") {
			http.Error(w, "Invalid Authorization header format", http.StatusUnauthorized)
			return
		}

		// Extract token
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		// Validate token
		claims, err := validateJWTToken(tokenString)
		if err != nil {
			fmt.Printf("JWT validation failed: %v\n", err)
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		// Add user info to request context
		ctx := context.WithValue(r.Context(), "user_id", claims.Sub)
		ctx = context.WithValue(ctx, "user_email", claims.Email)
		ctx = context.WithValue(ctx, "user_name", claims.Name)

		// Call next handler with updated context
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// Initialize Zitadel configuration
func initZitadelConfig() error {
	issuer := os.Getenv("ZITADEL_ISSUER")
	if issuer == "" {
		return fmt.Errorf("ZITADEL_ISSUER environment variable not set")
	}

	zitadelConfig = &ZitadelConfig{
		Issuer:  issuer,
		JWKSURI: issuer + "/oauth/v2/keys",
	}

	jwksKeys = make(map[string]*rsa.PublicKey)

	// Fetch JWKS keys initially
	return fetchJWKSKeys()
}

// Fetch JWKS keys from Zitadel
func fetchJWKSKeys() error {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := client.Get(zitadelConfig.JWKSURI)
	if err != nil {
		return fmt.Errorf("failed to fetch JWKS: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("JWKS endpoint returned status %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read JWKS response: %v", err)
	}

	var jwks JWKS
	if err := json.Unmarshal(body, &jwks); err != nil {
		return fmt.Errorf("failed to parse JWKS: %v", err)
	}

	// Convert JWK to RSA public keys
	keysLoaded := 0
	for _, key := range jwks.Keys {
		if key.Kty == "RSA" && key.Use == "sig" {
			pubKey, err := jwkToRSAPublicKey(key)
			if err != nil {
				fmt.Printf("Failed to convert JWK to RSA key: %v\n", err)
				continue
			}
			jwksKeys[key.Kid] = pubKey
			keysLoaded++
		}
	}

	lastJWKSFetch = time.Now()

	if keysLoaded == 0 {
		return fmt.Errorf("no RSA signing keys found in JWKS")
	}

	return nil
}

// Convert JWK to RSA Public Key
func jwkToRSAPublicKey(key JWK) (*rsa.PublicKey, error) {
	// Decode the modulus (n) from base64url
	nBytes, err := base64.RawURLEncoding.DecodeString(key.N)
	if err != nil {
		return nil, fmt.Errorf("failed to decode modulus: %v", err)
	}

	// Decode the exponent (e) from base64url
	eBytes, err := base64.RawURLEncoding.DecodeString(key.E)
	if err != nil {
		return nil, fmt.Errorf("failed to decode exponent: %v", err)
	}

	// Convert bytes to big integers
	n := new(big.Int).SetBytes(nBytes)

	// Convert exponent bytes to int
	var e int
	for _, b := range eBytes {
		e = e*256 + int(b)
	}

	// Create RSA public key
	publicKey := &rsa.PublicKey{
		N: n,
		E: e,
	}

	return publicKey, nil
}

// Validate JWT token
func validateJWTToken(tokenString string) (*Claims, error) {
	// Refresh JWKS keys if needed (every hour)
	if time.Since(lastJWKSFetch) > time.Hour {
		if err := fetchJWKSKeys(); err != nil {
			fmt.Printf("Failed to refresh JWKS keys: %v\n", err)
		}
	}

	// Parse token
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		// Check signing method
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// Get key ID from header
		kid, ok := token.Header["kid"].(string)
		if !ok {
			return nil, fmt.Errorf("missing kid in token header")
		}

		// Get public key
		pubKey, exists := jwksKeys[kid]
		if !exists {
			return nil, fmt.Errorf("unknown key ID: %s", kid)
		}

		return pubKey, nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("invalid token claims")
	}

	// Validate issuer
	if claims.Issuer != zitadelConfig.Issuer {
		return nil, fmt.Errorf("invalid issuer: %s", claims.Issuer)
	}

	// Check expiration
	if claims.ExpiresAt != nil && claims.ExpiresAt.Time.Before(time.Now()) {
		return nil, fmt.Errorf("token expired")
	}

	return claims, nil
}
