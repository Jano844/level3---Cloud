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
        
        <h1 v-if="!isLoggedIn">Login with Zitadel</h1>
        <h1 v-else>Welcome Back!</h1>
        
        <p class="login-subtitle" v-if="!isLoggedIn">
          Secure authentication powered by Zitadel Identity Platform
        </p>
        <p class="login-subtitle" v-else>
          You are successfully logged in to your account
        </p>
        
        <div class="login-content">
          <!-- Login Button - Show only when not logged in -->
          <button 
            v-if="!isLoggedIn"
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

          <!-- Profile Section - Show only when logged in -->
          <div v-if="isLoggedIn" class="profile-section">
            <div class="profile-card">
              <div class="profile-avatar">
                <img v-if="userProfile.picture" :src="userProfile.picture" :alt="userProfile.name" />
                <div v-else class="avatar-placeholder">
                  <svg width="32" height="32" fill="currentColor" viewBox="0 0 16 16">
                    <path d="M8 8a3 3 0 1 0 0-6 3 3 0 0 0 0 6zm2-3a2 2 0 1 1-4 0 2 2 0 0 1 4 0zm4 8c0 1-1 1-1 1H3s-1 0-1-1 1-4 6-4 6 3 6 4zm-1-.004c-.001-.246-.154-.986-.832-1.664C11.516 10.68 10.289 10 8 10c-2.29 0-3.516.68-4.168 1.332-.678.678-.83 1.418-.832 1.664h10z"/>
                  </svg>
                </div>
              </div>
              <div class="profile-info">
                <h3>{{ userProfile.name }}</h3>
                <p class="email">{{ userProfile.email }}</p>
                <div class="profile-details">
                  <div class="detail-item" v-if="userProfile.username">
                    <span class="label">Username:</span>
                    <span class="value">{{ userProfile.username }}</span>
                  </div>
                  <div class="detail-item" v-if="userProfile.givenName">
                    <span class="label">First Name:</span>
                    <span class="value">{{ userProfile.givenName }}</span>
                  </div>
                  <div class="detail-item" v-if="userProfile.familyName">
                    <span class="label">Last Name:</span>
                    <span class="value">{{ userProfile.familyName }}</span>
                  </div>
                  <div class="detail-item">
                    <span class="label">User ID:</span>
                    <span class="value">{{ userProfile.sub?.substring(0, 20) }}...</span>
                  </div>
                  <div class="detail-item" v-if="userProfile.iss">
                    <span class="label">Issuer:</span>
                    <span class="value">{{ userProfile.iss.replace('https://', '').substring(0, 30) }}...</span>
                  </div>
                </div>
              </div>
            </div>
            
            <div class="profile-actions">
              <button class="btn-secondary" @click="goToHome">
                <svg width="16" height="16" fill="currentColor" viewBox="0 0 16 16">
                  <path d="m8 3.293 6 6V13.5a1.5 1.5 0 0 1-1.5 1.5h-9A1.5 1.5 0 0 1 2 13.5V9.293l6-6zm5-.793V6l-2-2V2.5a.5.5 0 0 1 .5-.5h1a.5.5 0 0 1 .5.5z"/>
                  <path d="M7.293 1.5a1 1 0 0 1 1.414 0l6.647 6.646a.5.5 0 0 1-.708.708L8 2.207 1.354 8.854a.5.5 0 1 1-.708-.708L7.293 1.5z"/>
                </svg>
                Go to Home
              </button>
              <button class="btn-debug" @click="printDebugInfo">
                <svg width="16" height="16" fill="currentColor" viewBox="0 0 16 16">
                  <path d="M6 10.5a.5.5 0 0 1 .5-.5h3a.5.5 0 0 1 0 1h-3a.5.5 0 0 1-.5-.5zm-2-3a.5.5 0 0 1 .5-.5h7a.5.5 0 0 1 0 1h-7a.5.5 0 0 1-.5-.5zm-2-3a.5.5 0 0 1 .5-.5h11a.5.5 0 0 1 0 1h-11a.5.5 0 0 1-.5-.5z"/>
                </svg>
                Show Debug Info
              </button>
              <button class="btn-logout" @click="logout">
                <svg width="16" height="16" fill="currentColor" viewBox="0 0 16 16">
                  <path fill-rule="evenodd" d="M10 12.5a.5.5 0 0 1-.5.5h-8a.5.5 0 0 1-.5-.5v-9a.5.5 0 0 1 .5-.5h8a.5.5 0 0 1 .5.5v2a.5.5 0 0 0 1 0v-2A1.5 1.5 0 0 0 9.5 2h-8A1.5 1.5 0 0 0 0 3.5v9A1.5 1.5 0 0 0 1.5 14h8a1.5 1.5 0 0 0 1.5-1.5v-2a.5.5 0 0 0-1 0v2z"/>
                  <path fill-rule="evenodd" d="M15.854 8.354a.5.5 0 0 0 0-.708l-3-3a.5.5 0 0 0-.708.708L14.293 7.5H5.5a.5.5 0 0 0 0 1h8.793l-2.147 2.146a.5.5 0 0 0 .708.708l3-3z"/>
                </svg>
                Logout
              </button>
            </div>

            <!-- Debug section - Remove this after debugging -->
            <details class="debug-section">
              <summary>Debug Information (Click to expand)</summary>
              <pre class="debug-content">{{ JSON.stringify(userProfile.fullPayload, null, 2) }}</pre>
            </details>
          </div>
          
          <!-- Security info - Show only when not logged in -->
          <div v-if="!isLoggedIn" class="security-info">
            <div class="security-item">
              <svg width="16" height="16" fill="currentColor" viewBox="0 0 16 16">
                <path d="M8 1a2 2 0 0 1 2 2v4H6V3a2 2 0 0 1 2-2zm3 6V3a3 3 0 0 0-6 0v4a2 2 0 0 0-2 2v5a2 2 0 0 0 2 2h6a2 2 0 0 0 2-2V9a2 2 0 0 0-2-2z"/>
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

<script>
import { ref, onMounted, computed } from 'vue'
import { useRouter } from 'vue-router'

export default {
  name: 'LoginView',
  setup() {
    const router = useRouter()
    const isLoading = ref(false)
    const userProfile = ref(null)
    const isLoggedIn = computed(() => !!userProfile.value)

    // Zitadel configuration - Update these values with your Zitadel instance
    const zitadelConfig = {
      authority: 'https://openstack-demo-fgevuc.us1.zitadel.cloud', // Base Zitadel URL
      authorizationEndpoint: 'https://openstack-demo-fgevuc.us1.zitadel.cloud/oauth/v2/authorize', // Authorization endpoint
      clientId: '330117514837175600', // Your Zitadel client ID
      redirectUri: `${window.location.origin}/auth/callback`, // Dynamic based on current origin
      scope: 'openid profile email', // Requested scopes
      responseType: 'code', // Authorization code flow
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
      
      // Optionally redirect to Zitadel logout endpoint
      // window.location.href = `${zitadelConfig.authority}/oidc/v1/end_session`
    }

    // Navigate to home
    const goToHome = () => {
      router.push('/')
    }

    // Print debug information to console
    const printDebugInfo = () => {
      const accessToken = localStorage.getItem('access_token')
      const idToken = localStorage.getItem('id_token')
      const refreshToken = localStorage.getItem('refresh_token')
      
      console.log('=== COMPLETE DEBUG INFORMATION ===')
      console.log('Current timestamp:', new Date().toISOString())
      console.log('Current URL:', window.location.href)
      console.log('')
      
      console.log('--- STORED TOKENS ---')
      console.log('Access Token (first 50 chars):', accessToken ? accessToken.substring(0, 50) + '...' : 'None')
      console.log('ID Token (first 50 chars):', idToken ? idToken.substring(0, 50) + '...' : 'None')
      console.log('Refresh Token (first 50 chars):', refreshToken ? refreshToken.substring(0, 50) + '...' : 'None')
      console.log('')
      
      console.log('--- ZITADEL CONFIGURATION ---')
      console.log('Authority:', zitadelConfig.authority)
      console.log('Client ID:', zitadelConfig.clientId)
      console.log('Redirect URI:', zitadelConfig.redirectUri)
      console.log('Scopes:', zitadelConfig.scope)
      console.log('')
      
      console.log('--- USER PROFILE DATA ---')
      console.log('Is Logged In:', isLoggedIn.value)
      console.log('User Profile Object:', userProfile.value)
      console.log('')
      
      if (userProfile.value && userProfile.value.fullPayload) {
        console.log('--- ID TOKEN PAYLOAD (DECODED) ---')
        console.log('Full Payload:', userProfile.value.fullPayload)
        console.log('Subject (sub):', userProfile.value.fullPayload.sub)
        console.log('Issuer (iss):', userProfile.value.fullPayload.iss)
        console.log('Audience (aud):', userProfile.value.fullPayload.aud)
        console.log('Expiration (exp):', userProfile.value.fullPayload.exp, '(', new Date(userProfile.value.fullPayload.exp * 1000).toISOString(), ')')
        console.log('Issued At (iat):', userProfile.value.fullPayload.iat, '(', new Date(userProfile.value.fullPayload.iat * 1000).toISOString(), ')')
        console.log('')
      }
      
      console.log('--- SESSION STORAGE ---')
      console.log('OAuth State:', sessionStorage.getItem('oauth_state'))
      console.log('Code Verifier:', sessionStorage.getItem('code_verifier') ? 'Present' : 'None')
      console.log('')
      
      console.log('--- ENVIRONMENT INFO ---')
      console.log('User Agent:', navigator.userAgent)
      console.log('Current Origin:', window.location.origin)
      console.log('Current Protocol:', window.location.protocol)
      console.log('Current Host:', window.location.host)
      console.log('')
      
      console.log('================================')
      
      // Also show an alert to let the user know the info was printed
      alert('Debug information has been printed to the browser console. Open Developer Tools (F12) to view it.')
    }

    const redirectToZitadel = async () => {
      isLoading.value = true
      
      try {
        // Generate state parameter for security
        const state = generateRandomString(32)
        sessionStorage.setItem('oauth_state', state)
        
        // Generate code verifier and challenge for PKCE
        const codeVerifier = generateRandomString(128)
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
        
        // Debug: Log all the parameters
        console.log('=== ZITADEL DEBUG INFO ===')
        console.log('Current origin:', window.location.origin)
        console.log('Redirect URI being sent:', zitadelConfig.redirectUri)
        console.log('Client ID:', zitadelConfig.clientId)
        console.log('Authorization endpoint:', zitadelConfig.authorizationEndpoint)
        console.log('Full auth URL:', authUrl.toString())
        console.log('========================')
        
        // Redirect to Zitadel
        window.location.href = authUrl.toString()
        
      } catch (error) {
        console.error('Error redirecting to Zitadel:', error)
        alert('Failed to redirect to Zitadel. Please check the configuration.')
        isLoading.value = false
      }
    }

    // Helper function to generate random string
    const generateRandomString = (length) => {
      const charset = 'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-._~'
      let result = ''
      for (let i = 0; i < length; i++) {
        result += charset.charAt(Math.floor(Math.random() * charset.length))
      }
      return result
    }

    // Generate code challenge for PKCE
    const generateCodeChallenge = async (verifier) => {
      const encoder = new TextEncoder()
      const data = encoder.encode(verifier)
      const digest = await window.crypto.subtle.digest('SHA-256', data)
      return btoa(String.fromCharCode(...new Uint8Array(digest)))
        .replace(/\+/g, '-')
        .replace(/\//g, '_')
        .replace(/=/g, '')
    }

    // Check auth status on component mount
    onMounted(() => {
      checkAuthStatus()
    })

    return {
      isLoading,
      isLoggedIn,
      userProfile,
      redirectToZitadel,
      logout,
      goToHome,
      printDebugInfo
    }
  }
}
</script>

<style scoped>
.login {
  min-height: 100vh;
  display: flex;
  align-items: center;
  justify-content: center;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  padding: 2rem;
}

.login-container {
  background: white;
  border-radius: 16px;
  box-shadow: 0 25px 50px rgba(0, 0, 0, 0.15);
  overflow: hidden;
  max-width: 450px;
  width: 100%;
}

.login-form {
  padding: 3rem 2.5rem;
  text-align: center;
}

.zitadel-logo {
  margin-bottom: 1.5rem;
  color: #667eea;
}

.login-form h1 {
  margin-bottom: 0.5rem;
  color: #2c3e50;
  font-size: 2.2rem;
  font-weight: 600;
}

.login-subtitle {
  color: #7f8c8d;
  margin-bottom: 2.5rem;
  font-size: 1rem;
  line-height: 1.5;
}

.login-content {
  margin-bottom: 2rem;
}

.btn-zitadel {
  width: 100%;
  padding: 1rem 1.5rem;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  color: white;
  border: none;
  border-radius: 10px;
  font-size: 1.1rem;
  font-weight: 600;
  cursor: pointer;
  transition: all 0.3s ease;
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 0.75rem;
  margin-bottom: 2rem;
  box-shadow: 0 4px 15px rgba(102, 126, 234, 0.4);
}

.btn-zitadel:hover:not(:disabled) {
  transform: translateY(-2px);
  box-shadow: 0 8px 25px rgba(102, 126, 234, 0.5);
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
  background: #f8f9fa;
  border-radius: 10px;
  padding: 1.5rem;
  margin-top: 1rem;
}

.profile-section {
  margin-top: 1rem;
}

.profile-card {
  background: #f8f9fa;
  border-radius: 12px;
  padding: 2rem;
  margin-bottom: 1.5rem;
  text-align: center;
}

.profile-avatar {
  margin-bottom: 1.5rem;
}

.profile-avatar img {
  width: 80px;
  height: 80px;
  border-radius: 50%;
  object-fit: cover;
  border: 3px solid #667eea;
}

.avatar-placeholder {
  width: 80px;
  height: 80px;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  color: white;
  margin: 0 auto;
}

.profile-info h3 {
  margin-bottom: 0.5rem;
  color: #2c3e50;
  font-size: 1.5rem;
}

.profile-info .email {
  color: #3498db;
  margin-bottom: 1rem;
  font-size: 1rem;
  font-weight: 500;
}

.debug-section {
  margin-top: 1.5rem;
  background: #f8f9fa;
  border-radius: 8px;
  padding: 1rem;
}

.debug-section summary {
  cursor: pointer;
  font-weight: 500;
  color: #6c757d;
  margin-bottom: 0.5rem;
}

.debug-content {
  background: #2c3e50;
  color: #ecf0f1;
  padding: 1rem;
  border-radius: 4px;
  font-size: 0.8rem;
  overflow-x: auto;
  margin-top: 0.5rem;
}

.profile-details {
  background: white;
  border-radius: 8px;
  padding: 1rem;
  margin-top: 1rem;
}

.detail-item {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 0.5rem;
}

.detail-item:last-child {
  margin-bottom: 0;
}

.detail-item .label {
  font-weight: 500;
  color: #5a6c7d;
  font-size: 0.9rem;
}

.detail-item .value {
  color: #2c3e50;
  font-size: 0.9rem;
  font-family: monospace;
}

.profile-actions {
  display: flex;
  gap: 1rem;
  justify-content: center;
  flex-wrap: wrap;
}

.btn-secondary, .btn-debug, .btn-logout {
  padding: 0.75rem 1.5rem;
  border: none;
  border-radius: 8px;
  font-size: 1rem;
  font-weight: 500;
  cursor: pointer;
  transition: all 0.3s ease;
  display: flex;
  align-items: center;
  gap: 0.5rem;
  text-decoration: none;
}

.btn-secondary {
  background: #6c757d;
  color: white;
}

.btn-secondary:hover {
  background: #5a6268;
  transform: translateY(-1px);
}

.btn-debug {
  background: #17a2b8;
  color: white;
}

.btn-debug:hover {
  background: #138496;
  transform: translateY(-1px);
}

.btn-logout {
  background: #dc3545;
  color: white;
}

.btn-logout:hover {
  background: #c82333;
  transform: translateY(-1px);
}

.security-item {
  display: flex;
  align-items: center;
  gap: 0.75rem;
  margin-bottom: 0.75rem;
  color: #5a6c7d;
  font-size: 0.9rem;
}

.security-item:last-child {
  margin-bottom: 0;
}

.security-item svg {
  color: #28a745;
  flex-shrink: 0;
}

.login-footer {
  text-align: center;
  padding-top: 1.5rem;
  border-top: 1px solid #e9ecef;
}

.login-footer p {
  color: #7f8c8d;
  margin: 0.5rem 0;
  font-size: 0.9rem;
}

.login-footer a {
  color: #3498db;
  text-decoration: none;
  font-weight: 500;
  transition: color 0.3s ease;
}

.login-footer a:hover {
  color: #2980b9;
  text-decoration: underline;
}

@media (max-width: 480px) {
  .login {
    padding: 1rem;
  }
  
  .login-form {
    padding: 2rem 1.5rem;
  }
  
  .login-form h1 {
    font-size: 1.8rem;
  }
  
  .btn-zitadel {
    font-size: 1rem;
    padding: 0.875rem 1.25rem;
  }
  
  .security-info {
    padding: 1.25rem;
  }
}
</style>
