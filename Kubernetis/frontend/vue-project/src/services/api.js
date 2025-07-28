// API service for backend communication
// Replace with your actual server IP and port
// Examples:
// const API_BASE_URL = 'http://192.168.1.100:30800'  // Local network IP
// const API_BASE_URL = 'http://your-cluster-ip:30800'  // Kubernetes cluster IP
// const API_BASE_URL = 'http://your-domain.com:30800'  // Domain name
const API_BASE_URL = 'http://88.99.80.96:30800'  // Change this to your server's IP:PORT

class ApiService {
  constructor() {
    this.baseURL = API_BASE_URL
  }

  // Helper method for making HTTP requests
  async request(endpoint, options = {}) {
    const url = `${this.baseURL}${endpoint}`
    const config = {
      headers: {
        'Content-Type': 'application/json',
        ...options.headers,
      },
      ...options,
    }

    try {
      const response = await fetch(url, config)
      
      if (!response.ok) {
        const errorText = await response.text()
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
  async createDatabase(username, dbname) {
    return await this.request('/createDB', {
      method: 'POST',
      body: JSON.stringify({
        username,
        dbname
      })
    })
  }

  // Delete a database
  async deleteDatabase(username, dbname) {
    return await this.request('/deleteDB', {
      method: 'DELETE',
      body: JSON.stringify({
        username,
        dbname
      })
    })
  }

  // Get database access information
  async getDatabaseAccess(username, dbname) {
    return await this.request('/getDBAccess', {
      method: 'POST',
      body: JSON.stringify({
        username,
        dbname
      })
    })
  }

  // Get user databases
  async getUserDatabases(username) {
    return await this.request('/getUserDatabases', {
      method: 'POST',
      body: JSON.stringify({
        username
      })
    })
  }
}

export default new ApiService()
