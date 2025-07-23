<template>
  <div class="profile-page">
    <div class="profile-container">
      <div class="profile-header">
        <h1>User Profile</h1>
        <p>Manage your CloudOps account and preferences</p>
      </div>

      <!-- User Profile Card -->
      <div class="profile-card">
        <div class="profile-avatar-section">
          <div class="profile-avatar">
            <img v-if="userProfile.picture" :src="userProfile.picture" :alt="userProfile.name" />
            <div v-else class="avatar-placeholder">
              <svg width="48" height="48" fill="currentColor" viewBox="0 0 16 16">
                <path d="M8 8a3 3 0 1 0 0-6 3 3 0 0 0 0 6zm2-3a2 2 0 1 1-4 0 2 2 0 0 1 4 0zm4 8c0 1-1 1-1 1H3s-1 0-1-1 1-4 6-4 6 3 6 4zm-1-.004c-.001-.246-.154-.986-.832-1.664C11.516 10.68 10.289 10 8 10c-2.29 0-3.516.68-4.168 1.332-.678.678-.83 1.418-.832 1.664h10z"/>
              </svg>
            </div>
          </div>
          <div class="profile-info">
            <h2>{{ userProfile.name }}</h2>
            <p class="user-email">{{ userProfile.email }}</p>
            <span class="user-status">Active</span>
          </div>
        </div>

        <div class="profile-details">
          <div class="detail-section">
            <h3>Account Information</h3>
            <div class="detail-grid">
              <div class="detail-item">
                <span class="label">User ID:</span>
                <span class="value">{{ userProfile.sub?.substring(0, 20) }}...</span>
              </div>
              <div class="detail-item">
                <span class="label">Email:</span>
                <span class="value">{{ userProfile.email }}</span>
              </div>
              <div class="detail-item">
                <span class="label">Name:</span>
                <span class="value">{{ userProfile.name }}</span>
              </div>
              <div class="detail-item">
                <span class="label">Provider:</span>
                <span class="value">Zitadel</span>
              </div>
            </div>
          </div>

          <div class="detail-section">
            <h3>Cloud Resources</h3>
            <div class="resource-stats">
              <div class="resource-item">
                <div class="resource-icon">
                  <svg width="24" height="24" viewBox="0 0 24 24" fill="currentColor">
                    <path d="M4 6H20V18H4V6ZM20 4H4C2.89 4 2 4.89 2 6V18C2 19.11 2.89 20 4 20H20C21.11 20 22 19.11 22 18V6C22 4.89 21.11 4 20 4Z"/>
                  </svg>
                </div>
                <div class="resource-info">
                  <span class="resource-count">5</span>
                  <span class="resource-label">Nodes</span>
                </div>
              </div>
              <div class="resource-item">
                <div class="resource-icon">
                  <svg width="24" height="24" viewBox="0 0 24 24" fill="currentColor">
                    <path d="M12 2C13.1 2 14 2.9 14 4C14 5.1 13.1 6 12 6C10.9 6 10 5.1 10 4C10 2.9 10.9 2 12 2ZM21 9V7L15 1H9V3H14.5L20 8.5V9H21ZM10 8V6H4V8H10ZM9 19V10H15V19C15 20.1 14.1 21 13 21H11C9.9 21 9 20.1 9 19Z"/>
                  </svg>
                </div>
                <div class="resource-info">
                  <span class="resource-count">42</span>
                  <span class="resource-label">Pods</span>
                </div>
              </div>
              <div class="resource-item">
                <div class="resource-icon">
                  <svg width="24" height="24" viewBox="0 0 24 24" fill="currentColor">
                    <path d="M19 3H5C3.9 3 3 3.9 3 5V19C3 20.1 3.9 21 5 21H19C20.1 21 21 20.1 21 19V5C21 3.9 20.1 3 19 3ZM19 19H5V5H19V19ZM17 12H15V10H17V12ZM7 8H9V10H7V8ZM11 16H13V18H11V16Z"/>
                  </svg>
                </div>
                <div class="resource-info">
                  <span class="resource-count">8</span>
                  <span class="resource-label">Namespaces</span>
                </div>
              </div>
            </div>
          </div>

          <div class="detail-section">
            <h3>Recent Activity</h3>
            <div class="activity-list">
              <div class="activity-item">
                <div class="activity-icon success">
                  <svg width="16" height="16" fill="currentColor" viewBox="0 0 16 16">
                    <path d="M10.067.87a2.89 2.89 0 0 0-4.134 0l-.622.638-.89-.011a2.89 2.89 0 0 0-2.924 2.924l.01.89-.636.622a2.89 2.89 0 0 0 0 4.134l.637.622-.011.89a2.89 2.89 0 0 0 2.924 2.924l.89-.01.622.636a2.89 2.89 0 0 0 4.134 0l.622-.637.89.011a2.89 2.89 0 0 0 2.924-2.924l-.01-.89.636-.622a2.89 2.89 0 0 0 0-4.134l-.637-.622.011-.89a2.89 2.89 0 0 0-2.924-2.924l-.89.01-.622-.636zm.287 5.984-3 3a.5.5 0 0 1-.708 0l-1.5-1.5a.5.5 0 1 1 .708-.708L7 8.793l2.646-2.647a.5.5 0 0 1 .708.708z"/>
                  </svg>
                </div>
                <div class="activity-content">
                  <span class="activity-text">Successfully deployed nginx-app</span>
                  <span class="activity-time">2 hours ago</span>
                </div>
              </div>
              <div class="activity-item">
                <div class="activity-icon info">
                  <svg width="16" height="16" fill="currentColor" viewBox="0 0 16 16">
                    <path d="M8 16A8 8 0 1 0 8 0a8 8 0 0 0 0 16zm.93-9.412-1 4.705c-.07.34.029.533.304.533.194 0 .487-.07.686-.246l-.088.416c-.287.346-.92.598-1.465.598-.703 0-1.002-.422-.808-1.319l.738-3.468c.064-.293.006-.399-.287-.47l-.451-.081.082-.381 2.29-.287zM8 5.5a1 1 0 1 1 0-2 1 1 0 0 1 0 2z"/>
                  </svg>
                </div>
                <div class="activity-content">
                  <span class="activity-text">Scaled backend service to 3 replicas</span>
                  <span class="activity-time">6 hours ago</span>
                </div>
              </div>
              <div class="activity-item">
                <div class="activity-icon info">
                  <svg width="16" height="16" fill="currentColor" viewBox="0 0 16 16">
                    <path d="M8 16A8 8 0 1 0 8 0a8 8 0 0 0 0 16zm.93-9.412-1 4.705c-.07.34.029.533.304.533.194 0 .487-.07.686-.246l-.088.416c-.287.346-.92.598-1.465.598-.703 0-1.002-.422-.808-1.319l.738-3.468c.064-.293.006-.399-.287-.47l-.451-.081.082-.381 2.29-.287zM8 5.5a1 1 0 1 1 0-2 1 1 0 0 1 0 2z"/>
                  </svg>
                </div>
                <div class="activity-content">
                  <span class="activity-text">Updated cluster configuration</span>
                  <span class="activity-time">1 day ago</span>
                </div>
              </div>
            </div>
          </div>
        </div>

        <div class="profile-actions">
          <button class="btn btn-primary" @click="goToDashboard">
            <svg width="16" height="16" fill="currentColor" viewBox="0 0 16 16">
              <path d="m8 3.293 6 6V13.5a1.5 1.5 0 0 1-1.5 1.5h-9A1.5 1.5 0 0 1 2 13.5V9.293l6-6zm5-.793V6l-2-2V2.5a.5.5 0 0 1 .5-.5h1a.5.5 0 0 1 .5.5z"/>
              <path d="M7.293 1.5a1 1 0 0 1 1.414 0l6.647 6.646a.5.5 0 0 1-.708.708L8 2.207 1.354 8.854a.5.5 0 1 1-.708-.708L7.293 1.5z"/>
            </svg>
            Go to Dashboard
          </button>
          <button class="btn btn-secondary" @click="editProfile">
            <svg width="16" height="16" fill="currentColor" viewBox="0 0 16 16">
              <path d="M12.146.146a.5.5 0 0 1 .708 0l3 3a.5.5 0 0 1 0 .708L7.5 12.207 4 13l.793-3.5L12.146.146zM11.5 1.207 13.793 3.5 12.5 4.793 10.207 2.5 11.5 1.207zM10.5 2.5 12.793 4.793 5.293 12.293 3 13l.707-3.293L10.5 2.5z"/>
            </svg>
            Edit Profile
          </button>
          <button class="btn btn-warning" @click="debugZitadelData">
            <svg width="16" height="16" fill="currentColor" viewBox="0 0 16 16">
              <path d="M8 15A7 7 0 1 1 8 1a7 7 0 0 1 0 14zm0 1A8 8 0 1 0 8 0a8 8 0 0 0 0 16z"/>
              <path d="M7.002 11a1 1 0 1 1 2 0 1 1 0 0 1-2 0zM7.1 4.995a.905.905 0 1 1 1.8 0l-.35 3.507a.552.552 0 0 1-1.1 0L7.1 4.995z"/>
            </svg>
            Debug Data
          </button>
          <button class="btn btn-danger" @click="logout">
            <svg width="16" height="16" fill="currentColor" viewBox="0 0 16 16">
              <path fill-rule="evenodd" d="M10 12.5a.5.5 0 0 1-.5.5h-8a.5.5 0 0 1-.5-.5v-9a.5.5 0 0 1 .5-.5h8a.5.5 0 0 1 .5.5v2a.5.5 0 0 0 1 0v-2A1.5 1.5 0 0 0 9.5 2h-8A1.5 1.5 0 0 0 0 3.5v9A1.5 1.5 0 0 0 1.5 14h8a1.5 1.5 0 0 0 1.5-1.5v-2a.5.5 0 0 0-1 0v2z"/>
              <path fill-rule="evenodd" d="M15.854 8.354a.5.5 0 0 0 0-.708l-3-3a.5.5 0 0 0-.708.708L14.293 7.5H5.5a.5.5 0 0 0 0 1h8.793l-2.147 2.146a.5.5 0 0 0 .708.708l3-3z"/>
            </svg>
            Logout
          </button>
        </div>
      </div>

      <!-- Debug Data Display -->
      <div v-if="showDebugData" class="debug-card">
        <div class="debug-header">
          <h2>Zitadel Debug Information</h2>
          <button class="debug-close" @click="showDebugData = false">
            <svg width="16" height="16" fill="currentColor" viewBox="0 0 16 16">
              <path d="M4.646 4.646a.5.5 0 0 1 .708 0L8 7.293l2.646-2.647a.5.5 0 0 1 .708.708L8.707 8l2.647 2.646a.5.5 0 0 1-.708.708L8 8.707l-2.646 2.647a.5.5 0 0 1-.708-.708L7.293 8 4.646 5.354a.5.5 0 0 1 0-.708z"/>
            </svg>
          </button>
        </div>
        
        <div class="debug-content">
          <div class="debug-section">
            <h3>Stored Tokens</h3>
            <div class="debug-data">
              <div class="debug-item">
                <span class="debug-label">Access Token:</span>
                <pre class="debug-value">{{ debugData.accessToken || 'Not found' }}</pre>
              </div>
              <div class="debug-item">
                <span class="debug-label">ID Token:</span>
                <pre class="debug-value">{{ debugData.idToken || 'Not found' }}</pre>
              </div>
              <div class="debug-item">
                <span class="debug-label">Refresh Token:</span>
                <pre class="debug-value">{{ debugData.refreshToken || 'Not found' }}</pre>
              </div>
            </div>
          </div>

          <div class="debug-section">
            <h3>Decoded ID Token Payload</h3>
            <pre class="debug-json">{{ JSON.stringify(debugData.decodedIdToken, null, 2) }}</pre>
          </div>

          <div class="debug-section">
            <h3>OAuth State & PKCE</h3>
            <div class="debug-data">
              <div class="debug-item">
                <span class="debug-label">OAuth State:</span>
                <pre class="debug-value">{{ debugData.oauthState || 'Not found' }}</pre>
              </div>
              <div class="debug-item">
                <span class="debug-label">Code Verifier:</span>
                <pre class="debug-value">{{ debugData.codeVerifier || 'Not found' }}</pre>
              </div>
            </div>
          </div>

          <div class="debug-section">
            <h3>User Profile (Parsed)</h3>
            <pre class="debug-json">{{ JSON.stringify(userProfile, null, 2) }}</pre>
          </div>

          <div class="debug-section" v-if="debugData.userInfo">
            <h3>UserInfo Endpoint Response</h3>
            <pre class="debug-json">{{ JSON.stringify(debugData.userInfo, null, 2) }}</pre>
          </div>

          <div class="debug-section">
            <h3>Zitadel Configuration</h3>
            <pre class="debug-json">{{ JSON.stringify(debugData.config, null, 2) }}</pre>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'

const router = useRouter()
const userProfile = ref({
  name: 'John Doe',
  email: 'john.doe@example.com',
  picture: null,
  sub: '1234567890abcdef1234567890abcdef12345678'
})

const showDebugData = ref(false)
const debugData = ref({
  accessToken: null,
  idToken: null,
  refreshToken: null,
  decodedIdToken: null,
  oauthState: null,
  codeVerifier: null,
  userInfo: null,
  config: null
})

onMounted(() => {
  // Check if user is logged in
  const accessToken = localStorage.getItem('access_token')
  const idToken = localStorage.getItem('id_token')
  
  if (!accessToken || !idToken) {
    // Redirect to login if not authenticated
    router.push('/login')
    return
  }

  // Try to get user profile from stored token
  try {
    const payload = parseJWT(idToken)
    userProfile.value = {
      name: payload.name || payload.preferred_username || 'User',
      email: payload.email || 'No email provided',
      picture: payload.picture || null,
      sub: payload.sub,
      iss: payload.iss,
      exp: payload.exp
    }
    
    // Check if token is expired
    if (payload.exp && payload.exp * 1000 < Date.now()) {
      console.log('Token expired, redirecting to login')
      logout()
    }
  } catch (error) {
    console.error('Error parsing user token:', error)
    logout()
  }
})

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

const logout = () => {
  localStorage.removeItem('access_token')
  localStorage.removeItem('refresh_token')
  localStorage.removeItem('id_token')
  sessionStorage.removeItem('oauth_state')
  sessionStorage.removeItem('code_verifier')
  router.push('/login')
}

const goToDashboard = () => {
  router.push('/')
}

const editProfile = () => {
  alert('Edit profile functionality would be implemented here!')
}

const debugZitadelData = async () => {
  console.log('=== ZITADEL DEBUG DATA ===')
  
  // Gather all stored tokens and data
  const accessToken = localStorage.getItem('access_token')
  const idToken = localStorage.getItem('id_token')
  const refreshToken = localStorage.getItem('refresh_token')
  const oauthState = sessionStorage.getItem('oauth_state')
  const codeVerifier = sessionStorage.getItem('code_verifier')
  
  // Decode ID token if available
  let decodedIdToken = null
  if (idToken) {
    try {
      decodedIdToken = parseJWT(idToken)
    } catch (error) {
      console.error('Error decoding ID token:', error)
    }
  }
  
  // Zitadel configuration (from your login setup)
  const config = {
    issuer: 'https://zitadel-vvxe2s.zitadel.cloud',
    clientId: '284946386936062465@demo',
    redirectUri: `${window.location.origin}/auth/callback`,
    scopes: ['openid', 'profile', 'email'],
    responseType: 'code',
    codeChallenge: 'S256'
  }
  
  // Try to fetch user info from userinfo endpoint
  let userInfo = null
  if (accessToken) {
    try {
      const response = await fetch(`${config.issuer}/oidc/v1/userinfo`, {
        headers: {
          'Authorization': `Bearer ${accessToken}`
        }
      })
      if (response.ok) {
        userInfo = await response.json()
      }
    } catch (error) {
      console.error('Error fetching user info:', error)
    }
  }
  
  // Populate debug data
  debugData.value = {
    accessToken,
    idToken,
    refreshToken,
    decodedIdToken,
    oauthState,
    codeVerifier,
    userInfo,
    config
  }
  
  // Log everything to console
  console.log('Access Token:', accessToken)
  console.log('ID Token:', idToken)
  console.log('Refresh Token:', refreshToken)
  console.log('Decoded ID Token:', decodedIdToken)
  console.log('OAuth State:', oauthState)
  console.log('Code Verifier:', codeVerifier)
  console.log('UserInfo Response:', userInfo)
  console.log('Zitadel Config:', config)
  console.log('Current User Profile:', userProfile.value)
  console.log('=== END DEBUG DATA ===')
  
  // Show debug data on page
  showDebugData.value = true
}
</script>

<style scoped>
.profile-page {
  min-height: 100vh;
  padding: 2rem;
  width: 100%;
}

.profile-container {
  max-width: 1000px;
  margin: 0 auto;
}

.profile-header {
  text-align: center;
  margin-bottom: 3rem;
}

.profile-header h1 {
  font-size: 2.5rem;
  font-weight: 700;
  color: white;
  margin-bottom: 0.5rem;
  text-shadow: 2px 2px 4px rgba(0, 0, 0, 0.3);
}

.profile-header p {
  font-size: 1.1rem;
  color: rgba(255, 255, 255, 0.9);
}

.profile-card {
  background: rgba(255, 255, 255, 0.95);
  border-radius: 20px;
  padding: 3rem;
  box-shadow: 0 25px 50px rgba(0, 0, 0, 0.2);
  backdrop-filter: blur(20px);
}

.profile-avatar-section {
  display: flex;
  align-items: center;
  gap: 2rem;
  margin-bottom: 3rem;
  padding-bottom: 2rem;
  border-bottom: 1px solid #e5e7eb;
}

.profile-avatar img {
  width: 100px;
  height: 100px;
  border-radius: 50%;
  object-fit: cover;
  border: 4px solid #667eea;
}

.avatar-placeholder {
  width: 100px;
  height: 100px;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  color: white;
}

.profile-info h2 {
  font-size: 2rem;
  font-weight: 600;
  color: #374151;
  margin-bottom: 0.5rem;
}

.user-email {
  font-size: 1.1rem;
  color: #6b7280;
  margin-bottom: 1rem;
}

.user-status {
  background: #dcfce7;
  color: #166534;
  padding: 0.25rem 1rem;
  border-radius: 20px;
  font-size: 0.875rem;
  font-weight: 500;
}

.profile-details {
  display: grid;
  gap: 3rem;
  margin-bottom: 3rem;
}

.detail-section h3 {
  font-size: 1.5rem;
  font-weight: 600;
  color: #374151;
  margin-bottom: 1.5rem;
}

.detail-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
  gap: 1.5rem;
}

.detail-item {
  background: #f9fafb;
  padding: 1.5rem;
  border-radius: 12px;
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.detail-item .label {
  font-weight: 500;
  color: #6b7280;
}

.detail-item .value {
  color: #374151;
  font-weight: 600;
  font-family: monospace;
}

.resource-stats {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
  gap: 1.5rem;
}

.resource-item {
  background: #f9fafb;
  padding: 2rem;
  border-radius: 12px;
  display: flex;
  align-items: center;
  gap: 1rem;
  transition: transform 0.3s ease;
}

.resource-item:hover {
  transform: translateY(-2px);
}

.resource-icon {
  width: 50px;
  height: 50px;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  border-radius: 12px;
  display: flex;
  align-items: center;
  justify-content: center;
  color: white;
}

.resource-count {
  font-size: 1.5rem;
  font-weight: 700;
  color: #374151;
  display: block;
}

.resource-label {
  font-size: 0.875rem;
  color: #6b7280;
  font-weight: 500;
}

.activity-list {
  display: flex;
  flex-direction: column;
  gap: 1rem;
}

.activity-item {
  display: flex;
  align-items: center;
  gap: 1rem;
  padding: 1rem;
  background: #f9fafb;
  border-radius: 8px;
}

.activity-icon {
  width: 32px;
  height: 32px;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  color: white;
}

.activity-icon.success {
  background: #10b981;
}

.activity-icon.info {
  background: #3b82f6;
}

.activity-content {
  flex: 1;
  display: flex;
  flex-direction: column;
}

.activity-text {
  font-weight: 500;
  color: #374151;
}

.activity-time {
  font-size: 0.875rem;
  color: #6b7280;
}

.profile-actions {
  display: flex;
  gap: 1rem;
  justify-content: center;
  flex-wrap: wrap;
}

.btn {
  padding: 0.75rem 2rem;
  border: none;
  border-radius: 12px;
  font-weight: 600;
  cursor: pointer;
  transition: all 0.3s ease;
  display: flex;
  align-items: center;
  gap: 0.5rem;
  text-decoration: none;
}

.btn-primary {
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  color: white;
  box-shadow: 0 4px 15px rgba(102, 126, 234, 0.4);
}

.btn-primary:hover {
  transform: translateY(-2px);
  box-shadow: 0 8px 25px rgba(102, 126, 234, 0.6);
}

.btn-secondary {
  background: #6b7280;
  color: white;
}

.btn-secondary:hover {
  background: #5a6268;
  transform: translateY(-2px);
}

.btn-danger {
  background: #dc3545;
  color: white;
}

.btn-danger:hover {
  background: #c82333;
  transform: translateY(-2px);
}

.btn-warning {
  background: #ffc107;
  color: #212529;
}

.btn-warning:hover {
  background: #e0a800;
  transform: translateY(-2px);
}

/* Debug Styles */
.debug-card {
  background: rgba(255, 255, 255, 0.95);
  border-radius: 20px;
  padding: 2rem;
  margin-top: 2rem;
  box-shadow: 0 25px 50px rgba(0, 0, 0, 0.2);
  backdrop-filter: blur(20px);
  border: 2px solid #ffc107;
}

.debug-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 2rem;
  padding-bottom: 1rem;
  border-bottom: 1px solid #e5e7eb;
}

.debug-header h2 {
  font-size: 1.8rem;
  font-weight: 600;
  color: #374151;
  margin: 0;
}

.debug-close {
  background: #dc3545;
  color: white;
  border: none;
  border-radius: 50%;
  width: 32px;
  height: 32px;
  display: flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
  transition: background 0.3s ease;
}

.debug-close:hover {
  background: #c82333;
}

.debug-content {
  display: flex;
  flex-direction: column;
  gap: 2rem;
}

.debug-section {
  background: #f8f9fa;
  padding: 1.5rem;
  border-radius: 12px;
  border-left: 4px solid #ffc107;
}

.debug-section h3 {
  font-size: 1.2rem;
  font-weight: 600;
  color: #374151;
  margin-bottom: 1rem;
}

.debug-data {
  display: flex;
  flex-direction: column;
  gap: 1rem;
}

.debug-item {
  background: white;
  padding: 1rem;
  border-radius: 8px;
  border: 1px solid #e5e7eb;
}

.debug-label {
  font-weight: 600;
  color: #6b7280;
  display: block;
  margin-bottom: 0.5rem;
  font-size: 0.875rem;
  text-transform: uppercase;
  letter-spacing: 0.5px;
}

.debug-value {
  background: #f3f4f6;
  padding: 0.75rem;
  border-radius: 6px;
  font-family: 'Monaco', 'Menlo', 'Ubuntu Mono', monospace;
  font-size: 0.875rem;
  color: #374151;
  white-space: pre-wrap;
  word-break: break-all;
  margin: 0;
  border: 1px solid #d1d5db;
}

.debug-json {
  background: #1f2937;
  color: #f9fafb;
  padding: 1rem;
  border-radius: 8px;
  font-family: 'Monaco', 'Menlo', 'Ubuntu Mono', monospace;
  font-size: 0.875rem;
  white-space: pre-wrap;
  word-break: break-word;
  margin: 0;
  max-height: 400px;
  overflow-y: auto;
  border: 1px solid #374151;
}

/* Responsive Design */
@media (max-width: 768px) {
  .profile-page {
    padding: 1rem;
  }
  
  .profile-card {
    padding: 2rem;
  }
  
  .profile-avatar-section {
    flex-direction: column;
    text-align: center;
    gap: 1rem;
  }
  
  .detail-grid {
    grid-template-columns: 1fr;
  }
  
  .resource-stats {
    grid-template-columns: 1fr;
  }
  
  .profile-actions {
    flex-direction: column;
  }
  
  .btn {
    justify-content: center;
  }
}
</style>
