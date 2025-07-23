<template>
  <div class="login">
    <div class="login-container">
      <div class="login-form">
        <div class="zitadel-logo">
          <svg width="48" height="48" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
            <path d="M12 2L2 7L12 12L22 7L12 2Z" stroke="currentColor" stroke-width="2" stroke-linejoin="round"/>
            <path d="M2 17L12 22L22 17" stroke="currentColor" stroke-width="2" stroke-linejoin="round"/>
            <path d="M2 12L12 17L22 12" stroke="currentColor" stroke-width="2" stroke-linejoin="round"/>
          </svg>
        </div>
        
        <h1>Login to CloudOps</h1>
        
        <p class="login-subtitle">
          Secure authentication powered by Zitadel Identity Platform
        </p>
        
        <div class="login-content">
          <!-- Login Button -->
          <button 
            class="btn-zitadel" 
            @click="redirectToZitadel" 
            :disabled="isLoading"
          >
            <svg class="zitadel-icon" width="20" height="20" viewBox="0 0 24 24" fill="none">
              <path d="M12 2L2 7L12 12L22 7L12 2Z" stroke="currentColor" stroke-width="2" stroke-linejoin="round"/>
              <path d="M2 17L12 22L22 17" stroke="currentColor" stroke-width="2" stroke-linejoin="round"/>
              <path d="M2 12L12 17L22 12" stroke="currentColor" stroke-width="2" stroke-linejoin="round"/>
            </svg>
            {{ isLoading ? 'Redirecting...' : 'Continue with Zitadel' }}
          </button>
          
          <!-- Security info -->
          <div class="security-info">
            <div class="security-item">
              <svg width="16" height="16" fill="currentColor" viewBox="0 0 16 16">
                <path d="M8 1a2 2 0 0 1 2 2v4H6V3a2 2 0 0 1 2-2zm3 6V3a3 3 0 0 0-6 0v4a2 2 0 0 0-2 2v5a2 2 0 0 0 2-2V9a2 2 0 0 0-2-2z"/>
              </svg>
              <span>Secure OAuth2/OIDC Authentication</span>
            </div>
            <div class="security-item">
              <svg width="16" height="16" fill="currentColor" viewBox="0 0 16 16">
                <path d="M10.067.87a2.89 2.89 0 0 0-4.134 0l-.622.638-.89-.011a2.89 2.89 0 0 0-2.924 2.924l.01.89-.636.622a2.89 2.89 0 0 0 0 4.134l.637.622-.011.89a2.89 2.89 0 0 0 2.924 2.924l.89-.01.622.636a2.89 2.89 0 0 0 4.134 0l.622-.637.89.011a2.89 2.89 0 0 0 2.924-2.924l-.01-.89.636-.622a2.89 2.89 0 0 0 0-4.134l-.637-.622.011-.89a2.89 2.89 0 0 0-2.924-2.924l-.89.01-.622-.636zm.287 5.984-3 3a.5.5 0 0 1-.708 0l-1.5-1.5a.5.5 0 1 1 .708-.708L7 8.793l2.646-2.647a.5.5 0 0 1 .708.708z"/>
              </svg>
              <span>Enterprise-grade Security</span>
            </div>
            <div class="security-item">
              <svg width="16" height="16" fill="currentColor" viewBox="0 0 16 16">
                <path d="M8 16A8 8 0 1 0 8 0a8 8 0 0 0 0 16zm.93-9.412-1 4.705c-.07.34.029.533.304.533.194 0 .487-.07.686-.246l-.088.416c-.287.346-.92.598-1.465.598-.703 0-1.002-.422-.808-1.319l.738-3.468c.064-.293.006-.399-.287-.47l-.451-.081.082-.381 2.29-.287zM8 5.5a1 1 0 1 1 0-2 1 1 0 0 1 0 2z"/>
              </svg>
              <span>Single Sign-On (SSO) Ready</span>
            </div>
          </div>
        </div>
        
        <div class="login-footer">
          <p>Powered by <a href="https://zitadel.com" target="_blank" rel="noopener">Zitadel</a></p>
          <p><router-link to="/">← Back to Home</router-link></p>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted, computed } from 'vue'
import { useRouter } from 'vue-router'

const router = useRouter()
const isLoading = ref(false)
const userProfile = ref(null)
const isLoggedIn = computed(() => !!userProfile.value)

// Zitadel configuration - Update these values with your Zitadel instance
const zitadelConfig = {
  authority: 'https://openstack-demo-fgevuc.us1.zitadel.cloud',
  authorizationEndpoint: 'https://openstack-demo-fgevuc.us1.zitadel.cloud/oauth/v2/authorize',
  clientId: '330117514837175600',
  redirectUri: `http://localhost:3000/auth/callback`,
  scope: 'openid profile email',
  responseType: 'code',
}

// Check if user is already logged in
const checkAuthStatus = async () => {
  const accessToken = localStorage.getItem('access_token')
  const idToken = localStorage.getItem('id_token')
  
  console.log('=== AUTH STATUS DEBUG ===')
  console.log('Access token exists:', !!accessToken)
  console.log('ID token exists:', !!idToken)
  
  if (accessToken && idToken) {
    try {
      // Decode the ID token to get user information
      const payload = parseJWT(idToken)
      console.log('Parsed ID token payload:', payload)
      
      // Try to fetch additional user info from userinfo endpoint
      let additionalUserInfo = null
      try {
        additionalUserInfo = await fetchUserInfo(accessToken)
        console.log('Additional user info from /userinfo:', additionalUserInfo)
      } catch (error) {
        console.warn('Could not fetch additional user info:', error)
      }
      
      // Merge information from ID token and userinfo endpoint
      userProfile.value = {
        name: additionalUserInfo?.name || payload.name || payload.preferred_username || payload.given_name || 'User',
        email: additionalUserInfo?.email || payload.email || 'No email provided',
        picture: additionalUserInfo?.picture || payload.picture || null,
        username: additionalUserInfo?.preferred_username || payload.preferred_username || payload.sub,
        givenName: additionalUserInfo?.given_name || payload.given_name || '',
        familyName: additionalUserInfo?.family_name || payload.family_name || '',
        sub: payload.sub,
        iss: payload.iss,
        exp: payload.exp,
        fullPayload: payload // For debugging
      }
      
      console.log('Final user profile:', userProfile.value)
      
      // Check if token is expired
      if (payload.exp && payload.exp * 1000 < Date.now()) {
        console.log('Token expired, logging out')
        logout()
      } else {
        // If logged in, redirect to profile page
        router.push('/profile')
      }
    } catch (error) {
      console.error('Error parsing user token:', error)
      logout()
    }
  }
  console.log('========================')
}

// Fetch additional user information from Zitadel userinfo endpoint
const fetchUserInfo = async (accessToken) => {
  const response = await fetch(`${zitadelConfig.authority}/oidc/v1/userinfo`, {
    headers: {
      'Authorization': `Bearer ${accessToken}`,
      'Content-Type': 'application/json'
    }
  })
  
  if (!response.ok) {
    throw new Error(`Failed to fetch user info: ${response.status}`)
  }
  
  return await response.json()
}

// Parse JWT token
const parseJWT = (token) => {
  try {
    const base64Url = token.split('.')[1]
    const base64 = base64Url.replace(/-/g, '+').replace(/_/g, '/')
    const jsonPayload = decodeURIComponent(
      atob(base64)
        .split('')
        .map(c => '%' + ('00' + c.charCodeAt(0).toString(16)).slice(-2))
        .join('')
    )
    return JSON.parse(jsonPayload)
  } catch (error) {
    throw new Error('Invalid JWT token')
  }
}

// Logout function
const logout = () => {
  localStorage.removeItem('access_token')
  localStorage.removeItem('refresh_token')
  localStorage.removeItem('id_token')
  sessionStorage.removeItem('oauth_state')
  sessionStorage.removeItem('code_verifier')
  userProfile.value = null
  router.push('/login')
}

// Redirect to Zitadel for authentication
const redirectToZitadel = async () => {
  isLoading.value = true
  
  try {
    // Generate state parameter for security
    const state = generateSecureRandomString(32)
    sessionStorage.setItem('oauth_state', state)
    
    // Generate code verifier and challenge for PKCE
    const codeVerifier = generateCodeVerifier()
    sessionStorage.setItem('code_verifier', codeVerifier)
    
    // Generate code challenge
    const codeChallenge = await generateCodeChallenge(codeVerifier)
    
    // Build authorization URL
    const authUrl = new URL(zitadelConfig.authorizationEndpoint)
    authUrl.searchParams.append('client_id', zitadelConfig.clientId)
    authUrl.searchParams.append('redirect_uri', zitadelConfig.redirectUri)
    authUrl.searchParams.append('scope', zitadelConfig.scope)
    authUrl.searchParams.append('response_type', zitadelConfig.responseType)
    authUrl.searchParams.append('state', state)
    authUrl.searchParams.append('code_challenge', codeChallenge)
    authUrl.searchParams.append('code_challenge_method', 'S256')
    
    console.log('=== REDIRECT DEBUG ===')
    console.log('State:', state)
    console.log('Code verifier:', codeVerifier)
    console.log('Code challenge:', codeChallenge)
    console.log('Auth URL:', authUrl.toString())
    console.log('======================')
    
    // Redirect to Zitadel
    window.location.href = authUrl.toString()
    
  } catch (error) {
    console.error('Error redirecting to Zitadel:', error)
    alert('Failed to redirect to Zitadel. Please check the configuration.')
    isLoading.value = false
  }
}

// Generate crypto-secure random string
const generateSecureRandomString = (length) => {
  const array = new Uint8Array(length)
  crypto.getRandomValues(array)
  
  // Convert to base64url
  const base64String = btoa(String.fromCharCode(...array))
  return base64String
    .replace(/\+/g, '-')
    .replace(/\//g, '_')
    .replace(/=/g, '')
    .substring(0, length)
}

// Generate PKCE code verifier using crypto-secure random
const generateCodeVerifier = () => {
  // Use crypto.getRandomValues for better randomness
  const array = new Uint8Array(96) // 96 bytes = 128 base64url chars
  crypto.getRandomValues(array)
  
  // Convert to base64url
  const base64String = btoa(String.fromCharCode(...array))
  return base64String
    .replace(/\+/g, '-')
    .replace(/\//g, '_')
    .replace(/=/g, '')
}

// Generate PKCE code challenge using SHA256
const generateCodeChallenge = async (codeVerifier) => {
  // Convert string to Uint8Array
  const encoder = new TextEncoder()
  const data = encoder.encode(codeVerifier)
  
  // Hash with SHA-256
  const hashBuffer = await crypto.subtle.digest('SHA-256', data)
  
  // Convert to base64url
  const hashArray = new Uint8Array(hashBuffer)
  const base64String = btoa(String.fromCharCode(...hashArray))
  
  // Convert to base64url format
  return base64String
    .replace(/\+/g, '-')
    .replace(/\//g, '_')
    .replace(/=/g, '')
}

// Check auth status on component mount
onMounted(() => {
  checkAuthStatus()
})
</script>

<style scoped>
.login {
  min-height: 100vh;
  display: flex;
  align-items: center;
  justify-content: center;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  padding: 20px;
}

.login-container {
  width: 100%;
  max-width: 450px;
  margin: 0 auto;
}

.login-form {
  background: white;
  padding: 3rem 2.5rem;
  border-radius: 16px;
  box-shadow: 0 20px 60px rgba(0, 0, 0, 0.1);
  text-align: center;
}

.zitadel-logo {
  margin-bottom: 1.5rem;
  color: #667eea;
}

h1 {
  font-size: 2rem;
  font-weight: 700;
  color: #1a202c;
  margin-bottom: 0.5rem;
}

.login-subtitle {
  color: #718096;
  font-size: 1rem;
  margin-bottom: 2rem;
  line-height: 1.5;
}

.login-content {
  margin-bottom: 2rem;
}

.btn-zitadel {
  width: 100%;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  color: white;
  border: none;
  padding: 1rem 1.5rem;
  font-size: 1rem;
  font-weight: 600;
  border-radius: 12px;
  cursor: pointer;
  transition: all 0.3s ease;
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 0.5rem;
  margin-bottom: 2rem;
}

.btn-zitadel:hover:not(:disabled) {
  transform: translateY(-2px);
  box-shadow: 0 8px 25px rgba(102, 126, 234, 0.4);
}

.btn-zitadel:disabled {
  opacity: 0.7;
  cursor: not-allowed;
  transform: none;
}

.zitadel-icon {
  flex-shrink: 0;
}

.security-info {
  display: flex;
  flex-direction: column;
  gap: 1rem;
  text-align: left;
}

.security-item {
  display: flex;
  align-items: center;
  gap: 0.75rem;
  color: #4a5568;
  font-size: 0.9rem;
}

.security-item svg {
  color: #48bb78;
  flex-shrink: 0;
}

.login-footer {
  border-top: 1px solid #e2e8f0;
  padding-top: 1.5rem;
  margin-top: 1.5rem;
  color: #718096;
  font-size: 0.9rem;
}

.login-footer p {
  margin: 0.5rem 0;
}

.login-footer a {
  color: #667eea;
  text-decoration: none;
  font-weight: 500;
}

.login-footer a:hover {
  text-decoration: underline;
}

/* Responsive design */
@media (max-width: 480px) {
  .login {
    padding: 10px;
  }
  
  .login-form {
    padding: 2rem 1.5rem;
  }
  
  h1 {
    font-size: 1.75rem;
  }
}
</style>
