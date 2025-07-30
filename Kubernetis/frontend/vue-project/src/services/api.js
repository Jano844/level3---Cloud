// API service for backend communication
// Replace with your actual server IP and port
// Examples:
// const API_BASE_URL = 'http://192.168.1.100:30800'  // Local network IP
// const API_BASE_URL = 'http://your-cluster-ip:30800'  // Kubernetes cluster IP
// const API_BASE_URL = 'http://your-domain.com:30800'  // Domain name
const API_BASE_IP = '88.99.80.96'
const API_BASE_URL = `http://${API_BASE_IP}:30800`  // Change this to your server's IP:PORT

class ApiService {
  constructor() {
    this.baseURL = API_BASE_URL
  }

  // Helper method for making HTTP requests
  async request(endpoint, options = {}) {
    const url = `${this.baseURL}${endpoint}`
    
    // Get JWT access token from localStorage for authentication
    const accessToken = localStorage.getItem('access_token')
    
    const config = {
      headers: {
        'Content-Type': 'application/json',
        // Include JWT token in Authorization header if available
        ...(accessToken && { 'Authorization': `Bearer ${accessToken}` }),
        ...options.headers,
      },
      ...options,
    }

    try {
      const response = await fetch(url, config)
      
      if (!response.ok) {
        const errorText = await response.text()
        
        // Handle authentication errors - redirect to login
        if (response.status === 401) {
          // Clear invalid tokens
          localStorage.removeItem('access_token')
          localStorage.removeItem('refresh_token')
          localStorage.removeItem('id_token')
          sessionStorage.removeItem('oauth_state')
          sessionStorage.removeItem('code_verifier')
          
          // // Redirect to login page
          // if (typeof window !== 'undefined' && window.location.pathname !== '/login') {
          //   window.location.href = '/login'
          // }
          
          throw new Error('Authentication required. Please log in again.')
        }
        
        throw new Error(`HTTP ${response.status}: ${errorText}`)
      }

      // Check if response has content
      const contentType = response.headers.get('content-type')
      if (contentType && contentType.includes('application/json')) {
        return await response.json()
      } else {
        return await response.text()
      }
    } catch (error) {
      console.error(`API request failed for ${endpoint}:`, error)
      throw error
    }
  }

  // Health check
  async healthCheck() {
    return await this.request('/')
  }

  // Get all pods
  async getPods() {
    return await this.request('/pods')
  }

  // Create a new database
  async createDatabase(dbname) {
    return await this.request('/createDB', {
      method: 'POST',
      body: JSON.stringify({
        dbname
      })
    })
  }

  // Delete a database
  async deleteDatabase(dbname) {
    return await this.request('/deleteDB', {
      method: 'DELETE',
      body: JSON.stringify({
        dbname
      })
    })
  }

  // Get database access information
  async getDatabaseAccess(dbname) {
    return await this.request('/getDBAccess', {
      method: 'POST',
      body: JSON.stringify({
        dbname
      })
    })
  }

  // Get user databases
  async getUserDatabases() {
    return await this.request('/getUserDatabases', {
      method: 'POST',
      body: JSON.stringify({})
    })
  }

  // Check if user is authenticated
  isAuthenticated() {
    const accessToken = localStorage.getItem('access_token')
    const idToken = localStorage.getItem('id_token')
    
    if (!accessToken || !idToken) {
      return false
    }
    
    try {
      // Check if ID token is expired
      const payload = this.parseJWT(idToken)
      if (payload.exp && payload.exp * 1000 < Date.now()) {
        return false
      }
      return true
    } catch (error) {
      return false
    }
  }

  // Get current user information from stored tokens
  getCurrentUser() {
    const idToken = localStorage.getItem('id_token')
    if (!idToken) {
      return null
    }
    
    try {
      const payload = this.parseJWT(idToken)
      return {
        username: payload.preferred_username || payload.sub,
        email: payload.email,
        name: payload.name || payload.given_name,
        sub: payload.sub,
        exp: payload.exp
      }
    } catch (error) {
      return null
    }
  }

  // Helper method to parse JWT tokens
  parseJWT(token) {
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

  // Clear authentication data
  clearAuth() {
    localStorage.removeItem('access_token')
    localStorage.removeItem('refresh_token')
    localStorage.removeItem('id_token')
    sessionStorage.removeItem('oauth_state')
    sessionStorage.removeItem('code_verifier')
  }
}

export default new ApiService()
