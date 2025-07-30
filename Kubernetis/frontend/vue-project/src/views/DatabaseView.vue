<template>
  <div class="database-management">
    <!-- Header Section -->
    <div class="page-header">
      <div class="header-content">
        <h1 class="page-title">Database Management</h1>
        <p class="page-subtitle">Create, manage, and access PostgreSQL databases in your Kubernetes cluster</p>
      </div>
    </div>

    <!-- Main Content -->
    <div class="main-content">
      <!-- Action Cards -->
      <div class="action-section">
        <div class="action-cards">
          <!-- Create Database Card -->
          <div class="action-card">
            <div class="card-header">
              <div class="card-icon create">
                <svg width="24" height="24" fill="currentColor" viewBox="0 0 16 16">
                  <path d="M8 4a.5.5 0 0 1 .5.5v3h3a.5.5 0 0 1 0 1h-3v3a.5.5 0 0 1-1 0v-3h-3a.5.5 0 0 1 0-1h3v-3A.5.5 0 0 1 8 4z"/>
                </svg>
              </div>
              <h3>Create Database</h3>
            </div>
            <p>Create a new PostgreSQL database cluster with external access</p>
            <form @submit.prevent="createDatabase" class="db-form">
              <div class="form-group">
                <label for="create-dbname">Database Name</label>
                <input 
                  id="create-dbname"
                  v-model="createForm.dbname" 
                  type="text" 
                  placeholder="Enter database name"
                  required
                />
              </div>
              <button type="submit" class="btn btn-primary" :disabled="isLoading">
                <span v-if="isLoading" class="spinner"></span>
                {{ isLoading ? 'Creating...' : 'Create Database' }}
              </button>
            </form>
          </div>

          <!-- Get Access Info Card -->
          <div class="action-card">
            <div class="card-header">
              <div class="card-icon info">
                <svg width="24" height="24" fill="currentColor" viewBox="0 0 16 16">
                  <path d="M8 16A8 8 0 1 0 8 0a8 8 0 0 0 0 16zm.93-9.412-1 4.705c-.07.34.029.533.304.533.194 0 .487-.07.686-.246l-.088.416c-.287.346-.92.598-1.465.598-.703 0-1.002-.422-.808-1.319l.738-3.468c.064-.293.006-.399-.287-.47l-.451-.081.082-.381 2.29-.287zM8 5.5a1 1 0 1 1 0-2 1 1 0 0 1 0 2z"/>
                </svg>
              </div>
              <h3>Get Access Info</h3>
            </div>
            <p>Retrieve connection information and credentials for an existing database</p>
            <form @submit.prevent="getAccessInfo" class="db-form">
              <div class="form-group">
                <label for="access-dbname">Database Name</label>
                <input 
                  id="access-dbname"
                  v-model="accessForm.dbname" 
                  type="text" 
                  placeholder="Enter database name"
                  required
                />
              </div>
              <button type="submit" class="btn btn-info" :disabled="isLoading">
                <span v-if="isLoading" class="spinner"></span>
                {{ isLoading ? 'Getting Info...' : 'Get Access Info' }}
              </button>
            </form>
          </div>

          <!-- Delete Database Card -->
          <div class="action-card">
            <div class="card-header">
              <div class="card-icon delete">
                <svg width="24" height="24" fill="currentColor" viewBox="0 0 16 16">
                  <path d="M5.5 5.5A.5.5 0 0 1 6 6v6a.5.5 0 0 1-1 0V6a.5.5 0 0 1 .5-.5zm2.5 0a.5.5 0 0 1 .5.5v6a.5.5 0 0 1-1 0V6a.5.5 0 0 1 .5-.5zm3 .5a.5.5 0 0 0-1 0v6a.5.5 0 0 0 1 0V6z"/>
                  <path fill-rule="evenodd" d="M14.5 3a1 1 0 0 1-1 1H13v9a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V4h-.5a1 1 0 0 1-1-1V2a1 1 0 0 1 1-1H6a1 1 0 0 1 1-1h2a1 1 0 0 1 1 1h3.5a1 1 0 0 1 1 1v1zM4.118 4 4 4.059V13a1 1 0 0 0 1 1h6a1 1 0 0 0 1-1V4.059L11.882 4H4.118zM2.5 3V2h11v1h-11z"/>
                </svg>
              </div>
              <h3>Delete Database</h3>
            </div>
            <p>Permanently delete a PostgreSQL database cluster</p>
            <form @submit.prevent="deleteDatabase" class="db-form">
              <div class="form-group">
                <label for="delete-dbname">Database Name</label>
                <input 
                  id="delete-dbname"
                  v-model="deleteForm.dbname" 
                  type="text" 
                  placeholder="Enter database name"
                  required
                />
              </div>
              <button type="submit" class="btn btn-danger" :disabled="isLoading">
                <span v-if="isLoading" class="spinner"></span>
                {{ isLoading ? 'Deleting...' : 'Delete Database' }}
              </button>
            </form>
          </div>
        </div>
      </div>

      <!-- User Databases Section -->
      <div class="user-databases-section">
        <div class="section-card">
          <div class="card-header">
            <div class="card-icon list">
              <svg width="24" height="24" fill="currentColor" viewBox="0 0 16 16">
                <path d="M8 3c-1.552 0-2.94.707-3.905 1.028.965.32 2.353 1.028 3.905 1.028s2.94-.707 3.905-1.028C11.94 3.707 10.552 3 8 3z"/>
                <path d="M4 9a1 1 0 0 1-1-1V6a1 1 0 0 1 1-1v4zM13 6a1 1 0 0 0-1 1v2a1 1 0 0 0 1 1V6zM4 11a1 1 0 0 1-1-1v-1a1 1 0 0 1 1-1v3zM13 10a1 1 0 0 0-1 1v1a1 1 0 0 0 1 1v-3z"/>
                <path d="M5 2c0-1.105.895-2 2-2h2c1.105 0 2 .895 2 2v2c0 1.105-.895 2-2 2H7c-1.105 0-2-.895-2-2V2z"/>
              </svg>
            </div>
            <h3>My Databases</h3>
          </div>
          <p>View and manage all your existing databases</p>
          
          <form @submit.prevent="getUserDatabases" class="db-form">
            <button type="submit" class="btn btn-secondary" :disabled="isLoading">
              <span v-if="isLoading" class="spinner"></span>
              {{ isLoading ? 'Loading...' : 'Get My Databases' }}
            </button>
          </form>

          <!-- Database List -->
          <div v-if="userDatabases" class="databases-list">
            <div class="databases-header">
              <h4>Found {{ (userDatabases.databases && userDatabases.databases.length) || 0 }} database(s)</h4>
              <div class="status-info">
                <span class="status-badge success">{{ userDatabases.status }}</span>
                <span class="cluster-name">Cluster: {{ userDatabases.cluster_name }}</span>
              </div>
            </div>
            
            <div v-if="userDatabases.databases && userDatabases.databases.length > 0" class="database-buttons">
              <button 
                v-for="dbName in userDatabases.databases" 
                :key="dbName"
                @click="accessDatabase(dbName)"
                class="btn btn-database"
                :disabled="isLoading"
              >
                <svg width="16" height="16" fill="currentColor" viewBox="0 0 16 16" class="db-icon">
                  <path d="M4.318 2.687C5.234 2.271 6.536 2 8 2s2.766.27 3.682.687C12.644 3.125 13 3.627 13 4c0 .374-.356.875-1.318 1.313C10.766 5.729 9.464 6 8 6s-2.766-.27-3.682-.687C3.356 4.875 3 4.373 3 4c0-.374.356-.875 1.318-1.313ZM13 5.698V7c0 .374-.356.875-1.318 1.313C10.766 8.729 9.464 9 8 9s-2.766-.27-3.682-.687C3.356 7.875 3 7.373 3 7V5.698c.271.202.58.378.904.525C4.978 6.711 6.427 7 8 7s3.022-.289 4.096-.777A4.92 4.92 0 0 0 13 5.698ZM14 4c0-1.313-.967-2.5-2.682-3.113C10.363.264 9.223 0 8 0s-2.363.264-3.318.887C2.967 1.5 2 2.687 2 4v9c0 1.313.967 2.5 2.682 3.113C5.637 16.736 6.777 17 8 17s2.363-.264 3.318-.887C13.033 15.5 14 14.313 14 13V4ZM3 9.698V11c0 .374.356.875 1.318 1.313C5.234 12.729 6.536 13 8 13s2.766-.27 3.682-.687C12.644 11.875 13 11.373 13 11V9.698c-.271.202-.58.378-.904.525C11.022 10.711 9.573 11 8 11s-3.022-.289-4.096-.777A4.92 4.92 0 0 1 3 9.698Z"/>
                </svg>
                {{ dbName }}
              </button>
            </div>
            
            <div v-else class="no-databases">
              <p>No databases found for this user.</p>
            </div>
          </div>
        </div>
      </div>

      <!-- Results Section -->
      <div v-if="result" class="result-section">
        <div class="result-card" :class="result.type">
          <div class="result-header">
            <div class="result-icon">
              <svg v-if="result.type === 'success'" width="24" height="24" fill="currentColor" viewBox="0 0 16 16">
                <path d="M16 8A8 0 1 1 0 8a8 8 0 0 1 16 0zm-3.97-3.03a.75.75 0 0 0-1.08.022L7.477 9.417 5.384 7.323a.75.75 0 0 0-1.06 1.061L6.97 11.03a.75.75 0 0 0 1.079-.02l3.992-4.99a.75.75 0 0 0-.01-1.05z"/>
              </svg>
              <svg v-else width="24" height="24" fill="currentColor" viewBox="0 0 16 16">
                <path d="M16 8A8 0 1 1 0 8a8 8 0 0 1 16 0zM8 4a.905.905 0 0 0-.9.995l.35 3.507a.552.552 0 0 0 1.1 0l.35-3.507A.905.905 0 0 0 8 4zm.002 6a1 1 0 1 0 0 2 1 1 0 0 0 0-2z"/>
              </svg>
            </div>
            <h3>{{ result.title }}</h3>
            <button @click="clearResult" class="close-btn">
              <svg width="16" height="16" fill="currentColor" viewBox="0 0 16 16">
                <path d="M2.146 2.854a.5.5 0 1 1 .708-.708L8 7.293l5.146-5.147a.5.5 0 0 1 .708.708L8.707 8l5.147 5.146a.5.5 0 0 1-.708.708L8 8.707l-5.146 5.147a.5.5 0 0 1-.708-.708L7.293 8 2.146 2.854Z"/>
              </svg>
            </button>
          </div>
          <div class="result-content">
            <p>{{ result.message }}</p>
            
            <!-- Database Access Information (only for getDBAccess) -->
            <div v-if="result.data && result.data.status === 'success'" class="access-info">
              <div class="info-grid">
                <div class="info-item">
                  <label>Database Name:</label>
                  <span>{{ result.data.database_name }}</span>
                </div>
                <div class="info-item">
                  <label>Username:</label>
                  <span>{{ result.data.username }}</span>
                </div>
                <div class="info-item">
                  <label>Cluster Name:</label>
                  <span>{{ result.data.cluster_name }}</span>
                </div>
                <div class="info-item">
                  <label>Internal Host:</label>
                  <span>{{ result.data.internal_host }}</span>
                </div>
                <div class="info-item">
                  <label>Internal Port:</label>
                  <span>{{ result.data.internal_port }}</span>
                </div>
                <div class="info-item">
                  <label>External Host:</label>
                  <span>{{ result.data.external_host }}</span>
                </div>
                <div class="info-item">
                  <label>External Port:</label>
                  <span>{{ result.data.external_port }}</span>
                </div>
              </div>

              <!-- Connection Information -->
              <div class="connection-section">
                <h4>Connection Information</h4>
                
                <div class="connection-item">
                  <label>Password:</label>
                  <div class="password-field">
                    <input 
                      :type="showPassword ? 'text' : 'password'" 
                      :value="result.data.actual_password" 
                      readonly
                    />
                    <button @click="showPassword = !showPassword" class="toggle-password">
                      <svg v-if="showPassword" width="16" height="16" fill="currentColor" viewBox="0 0 16 16">
                        <path d="M13.359 11.238C15.06 9.72 16 8 16 8s-3-5.5-8-5.5a7.028 7.028 0 0 0-2.79.588l.77.771A5.944 5.944 0 0 1 8 3.5c2.12 0 3.879 1.168 5.168 2.457A13.134 13.134 0 0 1 14.828 8c-.058.087-.122.183-.195.288-.335.48-.83 1.12-1.465 1.755-.165.165-.337.328-.517.486l.708.709z"/>
                        <path d="M11.297 9.176a3.5 3.5 0 0 0-4.474-4.474l.823.823a2.5 2.5 0 0 1 2.829 2.829l.822.822zm-2.943 1.299.822.822a3.5 3.5 0 0 1-4.474-4.474l.823.823a2.5 2.5 0 0 0 2.829 2.829z"/>
                        <path d="M3.35 5.47c-.18.16-.353.322-.518.487A13.134 13.134 0 0 0 1.172 8l.195.288c.335.48.83 1.12 1.465 1.755C4.121 11.332 5.881 12.5 8 12.5c.716 0 1.39-.133 2.02-.36l.77.772A7.029 7.029 0 0 1 8 13.5C3 13.5 0 8 0 8s.939-1.721 2.641-3.238l.708.708zm10.296 8.884-12-12 .708-.708 12 12-.708.708z"/>
                      </svg>
                      <svg v-else width="16" height="16" fill="currentColor" viewBox="0 0 16 16">
                        <path d="M16 8s-3-5.5-8-5.5S0 8 0 8s3 5.5 8 5.5S16 8 16 8zM1.173 8a13.133 13.133 0 0 1 1.66-2.043C4.12 4.668 5.88 3.5 8 3.5c2.12 0 3.879 1.168 5.168 2.457A13.133 13.133 0 0 1 14.828 8c-.058.087-.122.183-.195.288-.335.48-.83 1.12-1.465 1.755C11.879 11.332 10.119 12.5 8 12.5c-2.12 0-3.879-1.168-5.168-2.457A13.134 13.134 0 0 1 1.172 8z"/>
                        <path d="M8 5.5a2.5 2.5 0 1 0 0 5 2.5 2.5 0 0 0 0-5zM4.5 8a3.5 3.5 0 1 1 7 0 3.5 3.5 0 0 1-7 0z"/>
                      </svg>
                    </button>
                  </div>
                </div>

                <div class="connection-item">
                  <label>Internal Connection String:</label>
                  <div class="connection-string">
                    <input :value="result.data.connection_string" readonly />
                    <button @click="copyToClipboard(result.data.connection_string)" class="copy-btn">
                      <svg width="16" height="16" fill="currentColor" viewBox="0 0 16 16">
                        <path d="M4 1.5H3a2 2 0 0 0-2 2V14a2 2 0 0 0 2 2h10a2 2 0 0 0 2-2V3.5a2 2 0 0 0-2-2h-1v1h1a1 1 0 0 1 1 1V14a1 1 0 0 1-1 1H3a1 1 0 0 1-1-1V3.5a1 1 0 0 1 1-1h1v-1z"/>
                        <path d="M9.5 1a.5.5 0 0 1 .5.5v1a.5.5 0 0 1-.5.5h-3a.5.5 0 0 1-.5-.5v-1a.5.5 0 0 1 .5-.5h3zm-3-1A1.5 1.5 0 0 0 5 1.5v1A1.5 1.5 0 0 0 6.5 4h3A1.5 1.5 0 0 0 11 2.5v-1A1.5 1.5 0 0 0 9.5 0h-3z"/>
                      </svg>
                    </button>
                  </div>
                </div>

                <div class="connection-item">
                  <label>External Connection String:</label>
                  <div class="connection-string">
                    <input :value="result.data.external_connection_string" readonly />
                    <button @click="copyToClipboard(result.data.external_connection_string)" class="copy-btn">
                      <svg width="16" height="16" fill="currentColor" viewBox="0 0 16 16">
                        <path d="M4 1.5H3a2 2 0 0 0-2 2V14a2 2 0 0 0 2 2h10a2 2 0 0 0 2-2V3.5a2 2 0 0 0-2-2h-1v1h1a1 1 0 0 1 1 1V14a1 1 0 0 1-1 1H3a1 1 0 0 1-1-1V3.5a1 1 0 0 1 1-1h1v-1z"/>
                        <path d="M9.5 1a.5.5 0 0 1 .5.5v1a.5.5 0 0 1-.5.5h-3a.5.5 0 0 1-.5-.5v-1a.5.5 0 0 1 .5-.5h3zm-3-1A1.5 1.5 0 0 0 5 1.5v1A1.5 1.5 0 0 0 6.5 4h3A1.5 1.5 0 0 0 11 2.5v-1A1.5 1.5 0 0 0 9.5 0h-3z"/>
                      </svg>
                    </button>
                  </div>
                </div>

                <div class="connection-item">
                  <label>PSQL Command:</label>
                  <div class="connection-string">
                    <input :value="result.data.psql_command" readonly />
                    <button @click="copyToClipboard(result.data.psql_command)" class="copy-btn">
                      <svg width="16" height="16" fill="currentColor" viewBox="0 0 16 16">
                        <path d="M4 1.5H3a2 2 0 0 0-2 2V14a2 2 0 0 0 2 2h10a2 2 0 0 0 2-2V3.5a2 2 0 0 0-2-2h-1v1h1a1 1 0 0 1 1 1V14a1 1 0 0 1-1 1H3a1 1 0 0 1-1-1V3.5a1 1 0 0 1 1-1h1v-1z"/>
                        <path d="M9.5 1a.5.5 0 0 1 .5.5v1a.5.5 0 0 1-.5.5h-3a.5.5 0 0 1-.5-.5v-1a.5.5 0 0 1 .5-.5h3zm-3-1A1.5 1.5 0 0 0 5 1.5v1A1.5 1.5 0 0 0 6.5 4h3A1.5 1.5 0 0 0 11 2.5v-1A1.5 1.5 0 0 0 9.5 0h-3z"/>
                      </svg>
                    </button>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- Error Section -->
      <div v-if="error" class="error-section">
        <div class="error-card">
          <div class="error-header">
            <div class="error-icon">
              <svg width="24" height="24" fill="currentColor" viewBox="0 0 16 16">
                <path d="M16 8A8 0 1 1 0 8a8 8 0 0 1 16 0zM8 4a.905.905 0 0 0-.9.995l.35 3.507a.552.552 0 0 0 1.1 0l.35-3.507A.905.905 0 0 0 8 4zm.002 6a1 1 0 1 0 0 2 1 1 0 0 0 0-2z"/>
              </svg>
            </div>
            <h3>Error</h3>
            <button @click="clearError" class="close-btn">
              <svg width="16" height="16" fill="currentColor" viewBox="0 0 16 16">
                <path d="M2.146 2.854a.5.5 0 1 1 .708-.708L8 7.293l5.146-5.147a.5.5 0 0 1 .708.708L8.707 8l5.147 5.146a.5.5 0 0 1-.708.708L8 8.707l-5.146 5.147a.5.5 0 0 1-.708-.708L7.293 8 2.146 2.854Z"/>
              </svg>
            </button>
          </div>
          <div class="error-content">
            <p>{{ error }}</p>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, reactive } from 'vue'
import ApiService from '../services/api.js'

// Get the API base IP for external host display
const API_BASE_IP = '88.99.80.96'

// Reactive form data
const createForm = reactive({
  dbname: ''
})

const accessForm = reactive({
  dbname: ''
})

const deleteForm = reactive({
  dbname: ''
})

// State management
const isLoading = ref(false)
const result = ref(null)
const error = ref(null)
const showPassword = ref(false)
const userDatabases = ref(null)

// Clear functions
const clearResult = () => {
  result.value = null
}

const clearError = () => {
  error.value = null
}

// Copy to clipboard function
const copyToClipboard = async (text) => {
  try {
    await navigator.clipboard.writeText(text)
    // Could add a toast notification here
    console.log('Copied to clipboard:', text)
  } catch (err) {
    console.error('Failed to copy to clipboard:', err)
  }
}

// API action functions
const createDatabase = async () => {
  if (!createForm.dbname) {
    error.value = 'Database name is required'
    return
  }

  isLoading.value = true
  error.value = null
  result.value = null

  try {
    const response = await ApiService.createDatabase(createForm.dbname)
    
    // Replace placeholders in any string fields that might contain IP placeholders
    if (response && typeof response === 'object') {
      Object.keys(response).forEach(key => {
        if (typeof response[key] === 'string') {
          response[key] = response[key]
            .replace(/<IP>/g, API_BASE_IP)
            .replace(/YOUR_CLUSTER_IP/g, API_BASE_IP)
        }
      })
    }
    
    result.value = {
      type: 'success',
      title: 'Database Created',
      message: `Successfully created database "${createForm.dbname}"`,
      data: response
    }

    // Clear form
    createForm.dbname = ''
  } catch (err) {
    error.value = `Failed to create database: ${err.message}`
  } finally {
    isLoading.value = false
  }
}

const getAccessInfo = async () => {
  if (!accessForm.dbname) {
    error.value = 'Database name is required'
    return
  }

  isLoading.value = true
  error.value = null
  result.value = null

  try {
    const response = await ApiService.getDatabaseAccess(accessForm.dbname)
    
    // Replace external_host with our API_BASE_IP
    if (response.external_host) {
      response.external_host = API_BASE_IP
    }
    
    // Replace IP in external connection string if it exists
    if (response.external_connection_string) {
      response.external_connection_string = response.external_connection_string.replace(
        /host=[^;]+/,
        `host=${API_BASE_IP}`
      )
    }
    
    // Replace placeholders in psql_command and any other fields
    if (response.psql_command) {
      response.psql_command = response.psql_command
        .replace(/<IP>/g, API_BASE_IP)
        .replace(/YOUR_CLUSTER_IP/g, API_BASE_IP)
    }
    
    // Replace placeholders in any other string fields that might contain IP placeholders
    Object.keys(response).forEach(key => {
      if (typeof response[key] === 'string') {
        response[key] = response[key]
          .replace(/<IP>/g, API_BASE_IP)
          .replace(/YOUR_CLUSTER_IP/g, API_BASE_IP)
      }
    })
    
    result.value = {
      type: 'success',
      title: 'Database Access Information',
      message: `Retrieved access information for database "${accessForm.dbname}"`,
      data: response
    }

    // Clear form
    accessForm.dbname = ''
  } catch (err) {
    error.value = `Failed to get database access info: ${err.message}`
  } finally {
    isLoading.value = false
  }
}

const deleteDatabase = async () => {
  if (!deleteForm.dbname) {
    error.value = 'Database name is required'
    return
  }

  // Confirmation dialog
  if (!confirm(`Are you sure you want to delete database "${deleteForm.dbname}"? This action cannot be undone.`)) {
    return
  }

  isLoading.value = true
  error.value = null
  result.value = null

  try {
    const response = await ApiService.deleteDatabase(deleteForm.dbname)
    
    result.value = {
      type: 'success',
      title: 'Database Deleted',
      message: `Successfully deleted database "${deleteForm.dbname}"`,
      data: response
    }

    // Clear form
    deleteForm.dbname = ''
  } catch (err) {
    error.value = `Failed to delete database: ${err.message}`
  } finally {
    isLoading.value = false
  }
}

// Get user databases
const getUserDatabases = async () => {
  isLoading.value = true
  error.value = null
  result.value = null
  userDatabases.value = null

  try {
    const response = await ApiService.getUserDatabases()
    userDatabases.value = response
  } catch (err) {
    error.value = `Failed to get user databases: ${err.message}`
  } finally {
    isLoading.value = false
  }
}

// Access a specific database (calls getDBAccess endpoint)
const accessDatabase = async (dbname) => {
  isLoading.value = true
  error.value = null
  result.value = null

  try {
    const response = await ApiService.getDatabaseAccess(dbname)
    
    // Replace external_host with our API_BASE_IP
    if (response.external_host) {
      response.external_host = API_BASE_IP
    }
    
    // Replace IP in external connection string if it exists
    if (response.external_connection_string) {
      response.external_connection_string = response.external_connection_string.replace(
        /host=[^;]+/,
        `host=${API_BASE_IP}`
      )
    }
    
    // Replace placeholders in psql_command and any other fields
    if (response.psql_command) {
      response.psql_command = response.psql_command
        .replace(/<IP>/g, API_BASE_IP)
        .replace(/YOUR_CLUSTER_IP/g, API_BASE_IP)
    }
    
    // Replace placeholders in any other string fields that might contain IP placeholders
    Object.keys(response).forEach(key => {
      if (typeof response[key] === 'string') {
        response[key] = response[key]
          .replace(/<IP>/g, API_BASE_IP)
          .replace(/YOUR_CLUSTER_IP/g, API_BASE_IP)
      }
    })
    
    result.value = {
      type: 'success',
      title: 'Database Access Information',
      message: `Retrieved access information for database "${dbname}"`,
      data: response
    }
  } catch (err) {
    error.value = `Failed to get database access info: ${err.message}`
  } finally {
    isLoading.value = false
  }
}
</script>

<style scoped>
.database-management {
  min-height: 100vh;
  background: linear-gradient(135deg, #f5f7fa 0%, #c3cfe2 100%);
  padding: 2rem 0;
}

.page-header {
  background: white;
  border-bottom: 1px solid #e9ecef;
  padding: 2rem 0;
  margin-bottom: 2rem;
  box-shadow: 0 2px 4px rgba(0,0,0,0.1);
}

.header-content {
  max-width: 1200px;
  margin: 0 auto;
  padding: 0 2rem;
  text-align: center;
}

.page-title {
  font-size: 2.5rem;
  font-weight: 700;
  color: #2c3e50;
  margin-bottom: 0.5rem;
}

.page-subtitle {
  font-size: 1.1rem;
  color: #6c757d;
  margin: 0;
}

.main-content {
  max-width: 1200px;
  margin: 0 auto;
  padding: 0 2rem;
}

.action-section {
  margin-bottom: 3rem;
}

.action-cards {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(350px, 1fr));
  gap: 2rem;
}

.action-card {
  background: white;
  border-radius: 12px;
  padding: 2rem;
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
  transition: transform 0.3s ease, box-shadow 0.3s ease;
}

.action-card:hover {
  transform: translateY(-4px);
  box-shadow: 0 8px 25px rgba(0, 0, 0, 0.15);
}

.card-header {
  display: flex;
  align-items: center;
  gap: 1rem;
  margin-bottom: 1rem;
}

.card-icon {
  width: 48px;
  height: 48px;
  border-radius: 12px;
  display: flex;
  align-items: center;
  justify-content: center;
  color: white;
}

.card-icon.create {
  background: linear-gradient(135deg, #28a745 0%, #20c997 100%);
}

.card-icon.info {
  background: linear-gradient(135deg, #007bff 0%, #6f42c1 100%);
}

.card-icon.delete {
  background: linear-gradient(135deg, #dc3545 0%, #fd7e14 100%);
}

.card-header h3 {
  margin: 0;
  font-size: 1.25rem;
  font-weight: 600;
  color: #2c3e50;
}

.action-card p {
  color: #6c757d;
  margin-bottom: 1.5rem;
  line-height: 1.5;
}

.db-form {
  display: flex;
  flex-direction: column;
  gap: 1rem;
}

.form-group {
  display: flex;
  flex-direction: column;
  gap: 0.5rem;
}

.form-group label {
  font-weight: 500;
  color: #495057;
  font-size: 0.9rem;
}

.form-group input {
  padding: 0.75rem;
  border: 2px solid #e9ecef;
  border-radius: 8px;
  font-size: 1rem;
  transition: border-color 0.3s ease;
}

.form-group input:focus {
  outline: none;
  border-color: #667eea;
  box-shadow: 0 0 0 3px rgba(102, 126, 234, 0.1);
}

.btn {
  padding: 0.75rem 1.5rem;
  border: none;
  border-radius: 8px;
  font-size: 1rem;
  font-weight: 500;
  cursor: pointer;
  transition: all 0.3s ease;
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 0.5rem;
}

.btn:disabled {
  opacity: 0.6;
  cursor: not-allowed;
}

.btn-primary {
  background: linear-gradient(135deg, #28a745 0%, #20c997 100%);
  color: white;
}

.btn-primary:hover:not(:disabled) {
  transform: translateY(-2px);
  box-shadow: 0 4px 15px rgba(40, 167, 69, 0.4);
}

.btn-info {
  background: linear-gradient(135deg, #007bff 0%, #6f42c1 100%);
  color: white;
}

.btn-info:hover:not(:disabled) {
  transform: translateY(-2px);
  box-shadow: 0 4px 15px rgba(0, 123, 255, 0.4);
}

.btn-danger {
  background: linear-gradient(135deg, #dc3545 0%, #fd7e14 100%);
  color: white;
}

.btn-danger:hover:not(:disabled) {
  transform: translateY(-2px);
  box-shadow: 0 4px 15px rgba(220, 53, 69, 0.4);
}

.btn-secondary {
  background: linear-gradient(135deg, #6c757d 0%, #495057 100%);
  color: white;
}

.btn-secondary:hover:not(:disabled) {
  transform: translateY(-2px);
  box-shadow: 0 4px 15px rgba(108, 117, 125, 0.4);
}

.spinner {
  width: 16px;
  height: 16px;
  border: 2px solid transparent;
  border-top: 2px solid currentColor;
  border-radius: 50%;
  animation: spin 1s linear infinite;
}

@keyframes spin {
  0% { transform: rotate(0deg); }
  100% { transform: rotate(360deg); }
}

.result-section, .error-section {
  margin-bottom: 2rem;
}

.result-card, .error-card {
  background: white;
  border-radius: 12px;
  overflow: hidden;
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
}

.result-card.success {
  border-left: 4px solid #28a745;
}

.error-card {
  border-left: 4px solid #dc3545;
}

.result-header, .error-header {
  background: #f8f9fa;
  padding: 1rem 1.5rem;
  display: flex;
  align-items: center;
  gap: 1rem;
  justify-content: space-between;
}

.result-icon, .error-icon {
  width: 24px;
  height: 24px;
  color: #28a745;
}

.error-icon {
  color: #dc3545;
}

.result-header h3, .error-header h3 {
  margin: 0;
  font-size: 1.125rem;
  font-weight: 600;
  color: #2c3e50;
  flex: 1;
}

.close-btn {
  background: none;
  border: none;
  color: #6c757d;
  cursor: pointer;
  padding: 0;
  width: 24px;
  height: 24px;
  display: flex;
  align-items: center;
  justify-content: center;
  border-radius: 4px;
  transition: background-color 0.3s ease;
}

.close-btn:hover {
  background: #e9ecef;
}

.result-content, .error-content {
  padding: 1.5rem;
}

.result-content p, .error-content p {
  margin: 0 0 1rem;
  color: #495057;
}

.access-info {
  margin-top: 1.5rem;
}

.info-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
  gap: 1rem;
  margin-bottom: 2rem;
}

.info-item {
  display: flex;
  flex-direction: column;
  gap: 0.25rem;
}

.info-item label {
  font-weight: 600;
  color: #495057;
  font-size: 0.9rem;
}

.info-item span {
  color: #2c3e50;
  font-family: 'Monaco', 'Menlo', 'Ubuntu Mono', monospace;
  background: #f8f9fa;
  padding: 0.5rem;
  border-radius: 4px;
  border: 1px solid #e9ecef;
}

.connection-section {
  border-top: 1px solid #e9ecef;
  padding-top: 1.5rem;
  margin-top: 1.5rem;
}

.connection-section h4 {
  margin: 0 0 1rem;
  color: #2c3e50;
  font-size: 1.125rem;
}

.connection-item {
  margin-bottom: 1rem;
}

.connection-item label {
  display: block;
  font-weight: 600;
  color: #495057;
  font-size: 0.9rem;
  margin-bottom: 0.5rem;
}

.password-field, .connection-string {
  display: flex;
  gap: 0.5rem;
}

.password-field input, .connection-string input {
  flex: 1;
  padding: 0.5rem;
  border: 1px solid #e9ecef;
  border-radius: 4px;
  font-family: 'Monaco', 'Menlo', 'Ubuntu Mono', monospace;
  background: #f8f9fa;
}

.toggle-password, .copy-btn {
  background: #007bff;
  color: white;
  border: none;
  border-radius: 4px;
  padding: 0.5rem;
  cursor: pointer;
  transition: background-color 0.3s ease;
  display: flex;
  align-items: center;
  justify-content: center;
}

.toggle-password:hover, .copy-btn:hover {
  background: #0056b3;
}

/* User Databases Section Styles */
.user-databases-section {
  margin-bottom: 3rem;
}

.section-card {
  background: white;
  border-radius: 12px;
  padding: 2rem;
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
  transition: transform 0.3s ease, box-shadow 0.3s ease;
}

.section-card:hover {
  transform: translateY(-2px);
  box-shadow: 0 6px 20px rgba(0, 0, 0, 0.15);
}

.card-icon.list {
  background: linear-gradient(135deg, #6f42c1 0%, #e83e8c 100%);
}

.btn-secondary {
  background: linear-gradient(135deg, #6c757d 0%, #495057 100%);
  color: white;
}

.btn-secondary:hover:not(:disabled) {
  transform: translateY(-2px);
  box-shadow: 0 4px 15px rgba(108, 117, 125, 0.4);
}

.databases-list {
  margin-top: 1.5rem;
  padding-top: 1.5rem;
  border-top: 1px solid #e9ecef;
}

.databases-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 1rem;
  flex-wrap: wrap;
  gap: 1rem;
}

.databases-header h4 {
  margin: 0;
  color: #2c3e50;
  font-size: 1.125rem;
}

.status-info {
  display: flex;
  align-items: center;
  gap: 1rem;
  flex-wrap: wrap;
}

.status-badge {
  padding: 0.25rem 0.75rem;
  border-radius: 20px;
  font-size: 0.875rem;
  font-weight: 500;
  text-transform: uppercase;
}

.status-badge.success {
  background: #d4edda;
  color: #155724;
  border: 1px solid #c3e6cb;
}

.cluster-name {
  color: #6c757d;
  font-size: 0.9rem;
  font-family: 'Monaco', 'Menlo', 'Ubuntu Mono', monospace;
  background: #f8f9fa;
  padding: 0.25rem 0.5rem;
  border-radius: 4px;
}

.database-buttons {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(200px, 1fr));
  gap: 1rem;
}

.btn-database {
  background: linear-gradient(135deg, #17a2b8 0%, #138496 100%);
  color: white;
  padding: 1rem;
  border: none;
  border-radius: 8px;
  font-size: 1rem;
  font-weight: 500;
  cursor: pointer;
  transition: all 0.3s ease;
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 0.5rem;
  text-align: center;
}

.btn-database:hover:not(:disabled) {
  transform: translateY(-2px);
  box-shadow: 0 4px 15px rgba(23, 162, 184, 0.4);
  background: linear-gradient(135deg, #138496 0%, #117a8b 100%);
}

.btn-database:disabled {
  opacity: 0.6;
  cursor: not-allowed;
}

.db-icon {
  flex-shrink: 0;
}

.no-databases {
  text-align: center;
  padding: 2rem;
  color: #6c757d;
}

.no-databases p {
  margin: 0;
  font-style: italic;
}

/* Responsive Design */
@media (max-width: 768px) {
  .database-management {
    padding: 1rem 0;
  }

  .main-content {
    padding: 0 1rem;
  }

  .action-cards {
    grid-template-columns: 1fr;
    gap: 1.5rem;
  }

  .action-card {
    padding: 1.5rem;
  }

  .page-title {
    font-size: 2rem;
  }

  .info-grid {
    grid-template-columns: 1fr;
  }

  .password-field, .connection-string {
    flex-direction: column;
  }

  .databases-header {
    flex-direction: column;
    align-items: flex-start;
  }

  .database-buttons {
    grid-template-columns: 1fr;
  }

  .status-info {
    flex-direction: column;
    align-items: flex-start;
    gap: 0.5rem;
  }
}
</style>
