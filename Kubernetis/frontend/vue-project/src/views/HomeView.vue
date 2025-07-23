<template>
  <div class="dashboard">
    <!-- Header Section -->
    <header class="dashboard-header">
      <div class="header-content">
        <div class="welcome-section">
          <h1 class="dashboard-title">
            {{ isLoggedIn ? `Welcome back, ${userName}!` : 'Welcome to Your Dashboard' }}
          </h1>
          <p class="dashboard-subtitle">
            {{ isLoggedIn ? 'Here\'s what\'s happening with your system today.' : 'Please log in to access your personalized dashboard.' }}
          </p>
        </div>
        <div class="header-actions">
          <button v-if="!isLoggedIn" @click="goToLogin" class="btn-primary">
            <svg width="16" height="16" fill="currentColor" viewBox="0 0 16 16">
              <path d="M8 1a2 2 0 0 1 2 2v4H6V3a2 2 0 0 1 2-2zm3 6V3a3 3 0 0 0-6 0v4a2 2 0 0 0-2 2v5a2 2 0 0 0 2 2h6a2 2 0 0 0 2-2V9a2 2 0 0 0-2-2z"/>
            </svg>
            Sign In
          </button>
          <div v-else class="user-menu">
            <div class="user-avatar">
              <img v-if="userProfile?.picture" :src="userProfile.picture" :alt="userName" />
              <div v-else class="avatar-placeholder">
                {{ userName.charAt(0).toUpperCase() }}
              </div>
            </div>
            <span class="user-name">{{ userName }}</span>
          </div>
        </div>
      </div>
    </header>

    <!-- Main Dashboard Content -->
    <main class="dashboard-main">
      <!-- Guest View - Not Logged In -->
      <div v-if="!isLoggedIn" class="guest-dashboard">
        <div class="guest-cards">
          <div class="guest-card">
            <div class="card-icon">
              <svg width="24" height="24" fill="currentColor" viewBox="0 0 16 16">
                <path d="M8 1a2 2 0 0 1 2 2v4H6V3a2 2 0 0 1 2-2zm3 6V3a3 3 0 0 0-6 0v4a2 2 0 0 0-2 2v5a2 2 0 0 0 2 2h6a2 2 0 0 0 2-2V9a2 2 0 0 0-2-2z"/>
              </svg>
            </div>
            <h3>Secure Authentication</h3>
            <p>Sign in with Zitadel for secure access to your personalized dashboard.</p>
            <button @click="goToLogin" class="card-action">Get Started</button>
          </div>
          
          <div class="guest-card">
            <div class="card-icon">
              <svg width="24" height="24" fill="currentColor" viewBox="0 0 16 16">
                <path d="M8 16A8 8 0 1 0 8 0a8 8 0 0 0 0 16zm.93-9.412-1 4.705c-.07.34.029.533.304.533.194 0 .487-.07.686-.246l-.088.416c-.287.346-.92.598-1.465.598-.703 0-1.002-.422-.808-1.319l.738-3.468c.064-.293.006-.399-.287-.47l-.451-.081.082-.381 2.29-.287zM8 5.5a1 1 0 1 1 0-2 1 1 0 0 1 0 2z"/>
              </svg>
            </div>
            <h3>Dashboard Features</h3>
            <p>Access analytics, system monitoring, and management tools after authentication.</p>
            <button @click="goToAbout" class="card-action">Learn More</button>
          </div>
          
          <div class="guest-card">
            <div class="card-icon">
              <svg width="24" height="24" fill="currentColor" viewBox="0 0 16 16">
                <path d="M9.405 1.05c-.413-1.4-2.397-1.4-2.81 0l-.1.34a1.464 1.464 0 0 1-2.105.872l-.31-.17c-1.283-.698-2.686.705-1.987 1.987l.169.311c.446.82.023 1.841-.872 2.105l-.34.1c-1.4.413-1.4 2.397 0 2.81l.34.1a1.464 1.464 0 0 1 .872 2.105l-.17.31c-.698 1.283.705 2.686 1.987 1.987l.311-.169a1.464 1.464 0 0 1 2.105.872l.1.34c.413 1.4 2.397 1.4 2.81 0l.1-.34a1.464 1.464 0 0 1 2.105-.872l.31.17c1.283.698 2.686-.705 1.987-1.987l-.169-.311a1.464 1.464 0 0 1 .872-2.105l.34-.1c1.4-.413 1.4-2.397 0-2.81l-.34-.1a1.464 1.464 0 0 1-.872-2.105l.17-.31c.698-1.283-.705-2.686-1.987-1.987l-.311.169a1.464 1.464 0 0 1-2.105-.872l-.1-.34zM8 10.93a2.929 2.929 0 1 1 0-5.86 2.929 2.929 0 0 1 0 5.858z"/>
              </svg>
            </div>
            <h3>System Integration</h3>
            <p>Seamlessly integrate with Kubernetes, MongoDB, and other cloud services.</p>
            <button class="card-action" disabled>Coming Soon</button>
          </div>
        </div>
      </div>

      <!-- Authenticated User Dashboard -->
      <div v-else class="user-dashboard">
        <!-- Statistics Cards -->
        <div class="stats-grid">
          <div class="stat-card">
            <div class="stat-icon success">
              <svg width="20" height="20" fill="currentColor" viewBox="0 0 16 16">
                <path d="M16 8A8 0 1 1 0 8a8 8 0 0 1 16 0zm-3.97-3.03a.75.75 0 0 0-1.08.022L7.477 9.417 5.384 7.323a.75.75 0 0 0-1.06 1.061L6.97 11.03a.75.75 0 0 0 1.079-.02l3.992-4.99a.75.75 0 0 0-.01-1.05z"/>
              </svg>
            </div>
            <div class="stat-content">
              <h3>System Status</h3>
              <p class="stat-value">Online</p>
              <span class="stat-label">All services running</span>
            </div>
          </div>

          <div class="stat-card">
            <div class="stat-icon info">
              <svg width="20" height="20" fill="currentColor" viewBox="0 0 16 16">
                <path d="M8 15A7 7 0 1 1 8 1a7 7 0 0 1 0 14zm0 1A8 8 0 1 0 8 0a8 8 0 0 0 0 16z"/>
                <path d="m8.93 6.588-2.29.287-.082.38.45.083c.294.07.352.176.288.469l-.738 3.468c-.194.897.105 1.319.808 1.319.545 0 1.178-.252 1.465-.598l.088-.416c-.2.176-.492.246-.686.246-.275 0-.375-.193-.304-.533L8.93 6.588zM9 4.5a1 1 0 1 1-2 0 1 1 0 0 1 2 0z"/>
              </svg>
            </div>
            <div class="stat-content">
              <h3>Session Time</h3>
              <p class="stat-value">{{ sessionDuration }}</p>
              <span class="stat-label">Active session</span>
            </div>
          </div>

          <div class="stat-card">
            <div class="stat-icon warning">
              <svg width="20" height="20" fill="currentColor" viewBox="0 0 16 16">
                <path d="M6 .278a.768.768 0 0 1 .08.858 7.208 7.208 0 0 0-.878 3.46c0 4.021 3.278 7.277 7.318 7.277.527 0 1.04-.055 1.533-.16a.787.787 0 0 1 .81.316.733.733 0 0 1-.031.893A8.349 8.349 0 0 1 8.344 16C3.734 16 0 12.286 0 7.71 0 4.266 2.114 1.312 5.124.06A.752.752 0 0 1 6 .278z"/>
              </svg>
            </div>
            <div class="stat-content">
              <h3>Theme</h3>
              <p class="stat-value">Light</p>
              <span class="stat-label">Current mode</span>
            </div>
          </div>

          <div class="stat-card">
            <div class="stat-icon primary">
              <svg width="20" height="20" fill="currentColor" viewBox="0 0 16 16">
                <path d="M1 11a1 1 0 0 1 1-1h2a1 1 0 0 1 1 1v3a1 1 0 0 1-1 1H2a1 1 0 0 1-1-1v-3zM7.5 4a1 1 0 0 1 1-1h1a1 1 0 0 1 1 1v10a1 1 0 0 1-1 1h-1a1 1 0 0 1-1-1V4zM14 7a1 1 0 0 1 1 1v6a1 1 0 0 1-1 1h-1a1 1 0 0 1-1-1V8a1 1 0 0 1 1-1h1z"/>
              </svg>
            </div>
            <div class="stat-content">
              <h3>Last Login</h3>
              <p class="stat-value">{{ lastLoginTime }}</p>
              <span class="stat-label">Recent activity</span>
            </div>
          </div>
        </div>

        <!-- Quick Actions -->
        <div class="quick-actions">
          <h2>Quick Actions</h2>
          <div class="actions-grid">
            <button class="action-card" @click="goToLogin">
              <div class="action-icon">
                <svg width="24" height="24" fill="currentColor" viewBox="0 0 16 16">
                  <path d="M8 8a3 3 0 1 0 0-6 3 3 0 0 0 0 6zm2-3a2 2 0 1 1-4 0 2 2 0 0 1 4 0zm4 8c0 1-1 1-1 1H3s-1 0-1-1 1-4 6-4 6 3 6 4zm-1-.004c-.001-.246-.154-.986-.832-1.664C11.516 10.68 10.289 10 8 10c-2.29 0-3.516.68-4.168 1.332-.678.678-.83 1.418-.832 1.664h10z"/>
                </svg>
              </div>
              <span>View Profile</span>
            </button>

            <button class="action-card" @click="goToAbout">
              <div class="action-icon">
                <svg width="24" height="24" fill="currentColor" viewBox="0 0 16 16">
                  <path d="M8 16A8 8 0 1 0 8 0a8 8 0 0 0 0 16zm.93-9.412-1 4.705c-.07.34.029.533.304.533.194 0 .487-.07.686-.246l-.088.416c-.287.346-.92.598-1.465.598-.703 0-1.002-.422-.808-1.319l.738-3.468c.064-.293.006-.399-.287-.47l-.451-.081.082-.381 2.29-.287zM8 5.5a1 1 0 1 1 0-2 1 1 0 0 1 0 2z"/>
                </svg>
              </div>
              <span>System Info</span>
            </button>

            <button class="action-card" @click="refreshData">
              <div class="action-icon">
                <svg width="24" height="24" fill="currentColor" viewBox="0 0 16 16">
                  <path fill-rule="evenodd" d="M8 3a5 5 0 1 0 4.546 2.914.5.5 0 0 1 .908-.417A6 6 0 1 1 8 2v1z"/>
                  <path d="M8 4.466V.534a.25.25 0 0 1 .41-.192l2.36 1.966c.12.1.12.284 0 .384L8.41 4.658A.25.25 0 0 1 8 4.466z"/>
                </svg>
              </div>
              <span>Refresh</span>
            </button>

            <button class="action-card" @click="logout">
              <div class="action-icon">
                <svg width="24" height="24" fill="currentColor" viewBox="0 0 16 16">
                  <path fill-rule="evenodd" d="M10 12.5a.5.5 0 0 1-.5.5h-8a.5.5 0 0 1-.5-.5v-9a.5.5 0 0 1 .5-.5h8a.5.5 0 0 1 .5.5v2a.5.5 0 0 0 1 0v-2A1.5 1.5 0 0 0 9.5 2h-8A1.5 1.5 0 0 0 0 3.5v9A1.5 1.5 0 0 0 1.5 14h8a1.5 1.5 0 0 0 1.5-1.5v-2a.5.5 0 0 0-1 0v2z"/>
                  <path fill-rule="evenodd" d="M15.854 8.354a.5.5 0 0 0 0-.708l-3-3a.5.5 0 0 0-.708.708L14.293 7.5H5.5a.5.5 0 0 0 0 1h8.793l-2.147 2.146a.5.5 0 0 0 .708.708l3-3z"/>
                </svg>
              </div>
              <span>Sign Out</span>
            </button>
          </div>
        </div>

        <!-- Recent Activity -->
        <div class="recent-activity">
          <h2>Recent Activity</h2>
          <div class="activity-list">
            <div class="activity-item">
              <div class="activity-icon">
                <svg width="16" height="16" fill="currentColor" viewBox="0 0 16 16">
                  <path d="M8 1a2 2 0 0 1 2 2v4H6V3a2 2 0 0 1 2-2zm3 6V3a3 3 0 0 0-6 0v4a2 2 0 0 0-2 2v5a2 2 0 0 0 2 2h6a2 2 0 0 0 2-2V9a2 2 0 0 0-2-2z"/>
                </svg>
              </div>
              <div class="activity-content">
                <p><strong>Successful login</strong> via Zitadel authentication</p>
                <span class="activity-time">{{ currentTime }}</span>
              </div>
            </div>
            
            <div class="activity-item">
              <div class="activity-icon">
                <svg width="16" height="16" fill="currentColor" viewBox="0 0 16 16">
                  <path fill-rule="evenodd" d="M8 3a5 5 0 1 0 4.546 2.914.5.5 0 0 1 .908-.417A6 6 0 1 1 8 2v1z"/>
                </svg>
              </div>
              <div class="activity-content">
                <p><strong>Dashboard loaded</strong> successfully</p>
                <span class="activity-time">{{ currentTime }}</span>
              </div>
            </div>
          </div>
        </div>
      </div>
    </main>
  </div>
</template>

<script>
import { ref, onMounted, computed } from 'vue'
import { useRouter } from 'vue-router'

export default {
  name: 'HomeView',
  setup() {
    const router = useRouter()
    const currentTime = ref('')
    const sessionStartTime = ref(null)
    const userProfile = ref(null)
    
    // Computed properties
    const isLoggedIn = computed(() => !!userProfile.value)
    const userName = computed(() => {
      if (userProfile.value) {
        return userProfile.value.name || userProfile.value.email || 'User'
      }
      return ''
    })
    
    const sessionDuration = computed(() => {
      if (!sessionStartTime.value) return '0m'
      const now = new Date()
      const diff = Math.floor((now - sessionStartTime.value) / 1000 / 60)
      if (diff < 60) return `${diff}m`
      const hours = Math.floor(diff / 60)
      const minutes = diff % 60
      return `${hours}h ${minutes}m`
    })
    
    const lastLoginTime = computed(() => {
      return sessionStartTime.value ? sessionStartTime.value.toLocaleTimeString() : '--:--'
    })

    // Check authentication status
    const checkAuthStatus = () => {
      const accessToken = localStorage.getItem('access_token')
      const idToken = localStorage.getItem('id_token')
      
      if (accessToken && idToken) {
        try {
          // Parse the ID token to get user information
          const payload = parseJWT(idToken)
          userProfile.value = {
            name: payload.name || payload.preferred_username || payload.given_name || 'User',
            email: payload.email || 'No email provided',
            picture: payload.picture || null,
            sub: payload.sub,
            iss: payload.iss
          }
          
          // Set session start time if not already set
          if (!sessionStartTime.value) {
            sessionStartTime.value = new Date()
          }
        } catch (error) {
          console.error('Error parsing user token:', error)
          userProfile.value = null
        }
      }
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

    // Update current time
    const updateTime = () => {
      currentTime.value = new Date().toLocaleTimeString()
    }

    // Navigation methods
    const goToLogin = () => {
      router.push('/login')
    }

    const goToAbout = () => {
      router.push('/about')
    }

    // Action methods
    const refreshData = () => {
      checkAuthStatus()
      updateTime()
      // Add a subtle visual feedback
      alert('Dashboard data refreshed!')
    }

    const logout = () => {
      localStorage.removeItem('access_token')
      localStorage.removeItem('refresh_token')
      localStorage.removeItem('id_token')
      sessionStorage.removeItem('oauth_state')
      sessionStorage.removeItem('code_verifier')
      userProfile.value = null
      sessionStartTime.value = null
      alert('You have been logged out successfully.')
    }

    // Initialize
    onMounted(() => {
      checkAuthStatus()
      updateTime()
      // Update time every minute
      setInterval(updateTime, 60000)
    })

    return {
      isLoggedIn,
      userName,
      userProfile,
      sessionDuration,
      lastLoginTime,
      currentTime,
      goToLogin,
      goToAbout,
      refreshData,
      logout
    }
  }
}
</script>

<style scoped>
.dashboard {
  min-height: 100vh;
  background: linear-gradient(135deg, #f5f7fa 0%, #c3cfe2 100%);
}

.dashboard-header {
  background: white;
  border-bottom: 1px solid #e9ecef;
  padding: 2rem 0;
  box-shadow: 0 2px 4px rgba(0,0,0,0.1);
}

.header-content {
  max-width: 1200px;
  margin: 0 auto;
  padding: 0 2rem;
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.welcome-section h1 {
  font-size: 2.5rem;
  font-weight: 700;
  color: #2c3e50;
  margin-bottom: 0.5rem;
}

.dashboard-subtitle {
  color: #6c757d;
  font-size: 1.1rem;
  margin: 0;
}

.header-actions .btn-primary {
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  color: white;
  border: none;
  padding: 0.75rem 1.5rem;
  border-radius: 8px;
  font-weight: 500;
  display: flex;
  align-items: center;
  gap: 0.5rem;
  cursor: pointer;
  transition: all 0.3s ease;
}

.header-actions .btn-primary:hover {
  transform: translateY(-2px);
  box-shadow: 0 4px 15px rgba(102, 126, 234, 0.4);
}

.user-menu {
  display: flex;
  align-items: center;
  gap: 1rem;
}

.user-avatar {
  width: 40px;
  height: 40px;
  border-radius: 50%;
  overflow: hidden;
  border: 2px solid #667eea;
}

.user-avatar img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.avatar-placeholder {
  width: 100%;
  height: 100%;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  display: flex;
  align-items: center;
  justify-content: center;
  color: white;
  font-weight: 600;
  font-size: 1.2rem;
}

.user-name {
  font-weight: 500;
  color: #2c3e50;
}

.dashboard-main {
  max-width: 1200px;
  margin: 0 auto;
  padding: 2rem;
}

/* Guest Dashboard */
.guest-cards {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(300px, 1fr));
  gap: 2rem;
  margin-top: 2rem;
}

.guest-card {
  background: white;
  border-radius: 12px;
  padding: 2rem;
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.07);
  text-align: center;
  transition: transform 0.3s ease, box-shadow 0.3s ease;
}

.guest-card:hover {
  transform: translateY(-4px);
  box-shadow: 0 8px 25px rgba(0, 0, 0, 0.15);
}

.guest-card .card-icon {
  width: 60px;
  height: 60px;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  color: white;
  margin: 0 auto 1.5rem;
}

.guest-card h3 {
  color: #2c3e50;
  margin-bottom: 1rem;
  font-size: 1.3rem;
}

.guest-card p {
  color: #6c757d;
  margin-bottom: 1.5rem;
  line-height: 1.6;
}

.card-action {
  background: #667eea;
  color: white;
  border: none;
  padding: 0.75rem 1.5rem;
  border-radius: 6px;
  font-weight: 500;
  cursor: pointer;
  transition: background 0.3s ease;
}

.card-action:hover:not(:disabled) {
  background: #5a67d8;
}

.card-action:disabled {
  background: #a0aec0;
  cursor: not-allowed;
}

/* User Dashboard */
.stats-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
  gap: 1.5rem;
  margin-bottom: 3rem;
}

.stat-card {
  background: white;
  border-radius: 12px;
  padding: 1.5rem;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.05);
  display: flex;
  align-items: center;
  gap: 1rem;
  transition: transform 0.2s ease;
}

.stat-card:hover {
  transform: translateY(-2px);
}

.stat-icon {
  width: 48px;
  height: 48px;
  border-radius: 10px;
  display: flex;
  align-items: center;
  justify-content: center;
  color: white;
}

.stat-icon.success {
  background: linear-gradient(135deg, #28a745 0%, #20c997 100%);
}

.stat-icon.info {
  background: linear-gradient(135deg, #17a2b8 0%, #007bff 100%);
}

.stat-icon.warning {
  background: linear-gradient(135deg, #ffc107 0%, #fd7e14 100%);
}

.stat-icon.primary {
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
}

.stat-content h3 {
  margin: 0 0 0.25rem;
  color: #6c757d;
  font-size: 0.9rem;
  font-weight: 500;
}

.stat-value {
  margin: 0 0 0.25rem;
  color: #2c3e50;
  font-size: 1.5rem;
  font-weight: 600;
}

.stat-label {
  color: #adb5bd;
  font-size: 0.8rem;
}

/* Quick Actions */
.quick-actions {
  margin-bottom: 3rem;
}

.quick-actions h2 {
  color: #2c3e50;
  margin-bottom: 1.5rem;
  font-size: 1.5rem;
}

.actions-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
  gap: 1rem;
}

.action-card {
  background: white;
  border: 1px solid #e9ecef;
  border-radius: 10px;
  padding: 1.5rem;
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 0.75rem;
  cursor: pointer;
  transition: all 0.3s ease;
  text-decoration: none;
  color: inherit;
}

.action-card:hover {
  border-color: #667eea;
  transform: translateY(-2px);
  box-shadow: 0 4px 12px rgba(102, 126, 234, 0.15);
}

.action-icon {
  width: 48px;
  height: 48px;
  background: #f8f9fa;
  border-radius: 10px;
  display: flex;
  align-items: center;
  justify-content: center;
  color: #667eea;
  transition: all 0.3s ease;
}

.action-card:hover .action-icon {
  background: #667eea;
  color: white;
}

.action-card span {
  font-weight: 500;
  color: #2c3e50;
}

/* Recent Activity */
.recent-activity h2 {
  color: #2c3e50;
  margin-bottom: 1.5rem;
  font-size: 1.5rem;
}

.activity-list {
  background: white;
  border-radius: 12px;
  padding: 1.5rem;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.05);
}

.activity-item {
  display: flex;
  align-items: flex-start;
  gap: 1rem;
  padding: 1rem 0;
  border-bottom: 1px solid #f1f3f4;
}

.activity-item:last-child {
  border-bottom: none;
  padding-bottom: 0;
}

.activity-icon {
  width: 32px;
  height: 32px;
  background: #f8f9fa;
  border-radius: 8px;
  display: flex;
  align-items: center;
  justify-content: center;
  color: #667eea;
  flex-shrink: 0;
}

.activity-content p {
  margin: 0 0 0.25rem;
  color: #2c3e50;
  font-size: 0.9rem;
}

.activity-time {
  color: #adb5bd;
  font-size: 0.8rem;
}

/* Responsive Design */
@media (max-width: 768px) {
  .header-content {
    flex-direction: column;
    gap: 1.5rem;
    text-align: center;
  }

  .welcome-section h1 {
    font-size: 2rem;
  }

  .dashboard-main {
    padding: 1rem;
  }

  .stats-grid {
    grid-template-columns: 1fr;
  }

  .actions-grid {
    grid-template-columns: repeat(2, 1fr);
  }

  .guest-cards {
    grid-template-columns: 1fr;
  }
}

@media (max-width: 480px) {
  .actions-grid {
    grid-template-columns: 1fr;
  }

  .stat-card {
    flex-direction: column;
    text-align: center;
  }
}
</style>
