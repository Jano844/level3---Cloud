<template>
  <div class="auth-callback">
    <div class="callback-container">
      <div class="callback-content">
        <div class="spinner" v-if="isProcessing">
          <svg class="animate-spin" width="48" height="48" viewBox="0 0 24 24">
            <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4" fill="none"></circle>
            <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 0 1 8-8v8H4z"></path>
          </svg>
        </div>
        
        <h2 v-if="isProcessing">Processing Authentication...</h2>
        <h2 v-else-if="error" class="error">Authentication Failed</h2>
        <h2 v-else class="success">Authentication Successful</h2>
        
        <p v-if="isProcessing">Please wait while we complete your login with Zitadel.</p>
        <p v-else-if="error" class="error-text">{{ error }}</p>
        <p v-else>You have been successfully authenticated. Redirecting...</p>
        
        <div v-if="error" class="error-actions">
          <button @click="retryLogin" class="btn-retry">Try Again</button>
          <router-link to="/" class="btn-home">Go Home</router-link>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { ref, onMounted } from 'vue'
import { useRouter, useRoute } from 'vue-router'

export default {
  name: 'AuthCallback',
  setup() {
    const router = useRouter()
    const route = useRoute()
    const isProcessing = ref(true)
    const error = ref(null)

    // Zitadel configuration - should match LoginView
    const zitadelConfig = {
      authority: 'https://openstack-demo-fgevuc.us1.zitadel.cloud',
      authorizationEndpoint: 'https://openstack-demo-fgevuc.us1.zitadel.cloud/oauth/v2/authorize',
      clientId: '330117514837175600',
      tokenEndpoint: 'https://openstack-demo-fgevuc.us1.zitadel.cloud/oauth/v2/token'
    }

    const handleCallback = async () => {
      try {
        const urlParams = new URLSearchParams(window.location.search)
        const code = urlParams.get('code')
        const state = urlParams.get('state')
        const storedState = sessionStorage.getItem('oauth_state')
        const codeVerifier = sessionStorage.getItem('code_verifier')

        // Validate state parameter
        if (!state || state !== storedState) {
          throw new Error('Invalid state parameter. Possible CSRF attack.')
        }

        if (!code) {
          throw new Error('Authorization code not found in callback URL.')
        }

        if (!codeVerifier) {
          throw new Error('Code verifier not found. Please start the login process again.')
        }

        // Exchange authorization code for tokens
        const tokenResponse = await exchangeCodeForTokens(code, codeVerifier)
        
        if (tokenResponse.access_token) {
          // Store tokens securely
          localStorage.setItem('access_token', tokenResponse.access_token)
          if (tokenResponse.refresh_token) {
            localStorage.setItem('refresh_token', tokenResponse.refresh_token)
          }
          if (tokenResponse.id_token) {
            localStorage.setItem('id_token', tokenResponse.id_token)
          }

          // Clean up session storage
          sessionStorage.removeItem('oauth_state')
          sessionStorage.removeItem('code_verifier')

          // Successful authentication
          isProcessing.value = false
          
          // Redirect after a short delay
          setTimeout(() => {
            router.push('/')
          }, 1500)
        } else {
          throw new Error('No access token received from Zitadel.')
        }

      } catch (err) {
        console.error('Authentication callback error:', err)
        error.value = err.message || 'An unexpected error occurred during authentication.'
        isProcessing.value = false
      }
    }

    const exchangeCodeForTokens = async (code, codeVerifier) => {
      const tokenData = {
        grant_type: 'authorization_code',
        client_id: zitadelConfig.clientId,
        code: code,
        redirect_uri: `${window.location.origin}/auth/callback`,
        code_verifier: codeVerifier
      }

      const response = await fetch(zitadelConfig.tokenEndpoint, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/x-www-form-urlencoded',
        },
        body: new URLSearchParams(tokenData)
      })

      if (!response.ok) {
        const errorData = await response.text()
        throw new Error(`Token exchange failed: ${response.status} ${errorData}`)
      }

      return await response.json()
    }

    const retryLogin = () => {
      // Clean up any stored state
      sessionStorage.removeItem('oauth_state')
      sessionStorage.removeItem('code_verifier')
      localStorage.removeItem('access_token')
      localStorage.removeItem('refresh_token')
      localStorage.removeItem('id_token')
      
      // Redirect to login
      router.push('/login')
    }

    onMounted(() => {
      handleCallback()
    })

    return {
      isProcessing,
      error,
      retryLogin
    }
  }
}
</script>

<style scoped>
.auth-callback {
  min-height: 100vh;
  display: flex;
  align-items: center;
  justify-content: center;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  padding: 2rem;
}

.callback-container {
  background: white;
  border-radius: 16px;
  box-shadow: 0 25px 50px rgba(0, 0, 0, 0.15);
  max-width: 500px;
  width: 100%;
}

.callback-content {
  padding: 3rem 2.5rem;
  text-align: center;
}

.spinner {
  margin: 0 auto 2rem;
  color: #667eea;
}

.animate-spin {
  animation: spin 1s linear infinite;
}

@keyframes spin {
  from {
    transform: rotate(0deg);
  }
  to {
    transform: rotate(360deg);
  }
}

h2 {
  margin-bottom: 1rem;
  font-size: 1.8rem;
}

h2.error {
  color: #e74c3c;
}

h2.success {
  color: #27ae60;
}

p {
  color: #7f8c8d;
  margin-bottom: 2rem;
  line-height: 1.6;
}

p.error-text {
  color: #e74c3c;
  background: #fdf2f2;
  padding: 1rem;
  border-radius: 8px;
  border: 1px solid #fecaca;
}

.error-actions {
  display: flex;
  gap: 1rem;
  justify-content: center;
  flex-wrap: wrap;
}

.btn-retry, .btn-home {
  padding: 0.75rem 1.5rem;
  border-radius: 8px;
  text-decoration: none;
  font-weight: 500;
  transition: all 0.3s ease;
  display: inline-block;
}

.btn-retry {
  background: #e74c3c;
  color: white;
  border: none;
  cursor: pointer;
}

.btn-retry:hover {
  background: #c0392b;
}

.btn-home {
  background: #6c757d;
  color: white;
}

.btn-home:hover {
  background: #5a6268;
}

@media (max-width: 480px) {
  .callback-content {
    padding: 2rem 1.5rem;
  }
  
  h2 {
    font-size: 1.5rem;
  }
  
  .error-actions {
    flex-direction: column;
  }
}
</style>
