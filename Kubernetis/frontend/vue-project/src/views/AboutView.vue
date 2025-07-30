<template>
  <div class="about-page">
    <!-- Hero Section -->
    <div class="hero-section">
      <div class="hero-content">
        <div class="hero-text">
          <h1 class="hero-title">CloudOps Platform</h1>
          <p class="hero-subtitle">Enterprise-Grade Database-as-a-Service</p>
          <p class="hero-description">
            Deploy, manage, and scale PostgreSQL databases in the cloud with enterprise-level security,
            high availability, and seamless integration. Built on modern cloud-native technologies.
          </p>
        </div>
        <div class="hero-visual">
          <div class="tech-stack">
            <div class="tech-item">
              <svg width="40" height="40" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                <path d="M12 2L2 7L12 12L22 7L12 2Z"/>
                <path d="M2 17L12 22L22 17"/>
                <path d="M2 12L12 17L22 12"/>
              </svg>
              <span>Kubernetes</span>
            </div>
            <div class="tech-item">
              <svg width="40" height="40" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                <ellipse cx="12" cy="5" rx="9" ry="3"/>
                <path d="M21 12c0 1.66-4 3-9 3s-9-1.34-9-3"/>
                <path d="M3 5v14c0 1.66 4 3 9 3s9-1.34 9-3V5"/>
              </svg>
              <span>PostgreSQL</span>
            </div>
            <div class="tech-item">
              <svg width="40" height="40" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                <rect x="2" y="3" width="20" height="14" rx="2" ry="2"/>
                <line x1="8" y1="21" x2="16" y2="21"/>
                <line x1="12" y1="17" x2="12" y2="21"/>
              </svg>
              <span>OpenStack</span>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- Platform Overview -->
    <div class="section">
      <div class="container">
        <h2 class="section-title">Platform Overview</h2>
        <div class="overview-grid">
          <div class="overview-card">
            <div class="card-icon">
              <svg width="32" height="32" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                <path d="M12 22s8-4 8-10V5l-8-3-8 3v7c0 6 8 10 8 10z"/>
              </svg>
            </div>
            <h3>Secure Authentication</h3>
            <p>Enterprise-grade authentication powered by Zitadel Identity Platform with OAuth2/OIDC compliance and multi-factor authentication support.</p>
          </div>
          
          <div class="overview-card">
            <div class="card-icon">
              <svg width="32" height="32" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                <ellipse cx="12" cy="5" rx="9" ry="3"/>
                <path d="M21 12c0 1.66-4 3-9 3s-9-1.34-9-3"/>
                <path d="M3 5v14c0 1.66 4 3 9 3s9-1.34 9-3V5"/>
              </svg>
            </div>
            <h3>Database Management</h3>
            <p>Create, manage, and scale PostgreSQL databases on-demand with automated backups, monitoring, and performance optimization.</p>
          </div>
          
          <div class="overview-card">
            <div class="card-icon">
              <svg width="32" height="32" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                <path d="M12 2L2 7L12 12L22 7L12 2Z"/>
                <path d="M2 17L12 22L22 17"/>
                <path d="M2 12L12 17L22 12"/>
              </svg>
            </div>
            <h3>Cloud-Native Infrastructure</h3>
            <p>Built on Kubernetes and OpenStack for maximum scalability, reliability, and resource efficiency in a cloud-native environment.</p>
          </div>
        </div>
      </div>
    </div>

    <!-- Architecture Section -->
    <div class="section architecture-section">
      <div class="container">
        <h2 class="section-title">System Architecture</h2>
        <div class="architecture-content">
          <div class="architecture-text">
            <h3>Cloud Infrastructure</h3>
            <p>Our platform is built on a robust foundation:</p>
            <ul class="feature-list">
              <li><strong>OpenStack Cloud:</strong> Private cloud infrastructure providing virtualized compute, storage, and networking resources</li>
              <li><strong>Kubernetes Cluster:</strong> Container orchestration across multiple VM instances ensuring high availability and scalability</li>
              <li><strong>Go Backend:</strong> High-performance RESTful API service handling database operations and user management</li>
              <li><strong>PostgreSQL Operator:</strong> Automated database provisioning, scaling, and maintenance using Kubernetes operators</li>
              <li><strong>Zitadel Identity:</strong> Modern identity and access management with enterprise-grade security</li>
            </ul>
          </div>
          <div class="architecture-diagram">
            <div class="layer layer-top">
              <div class="layer-title">User Interface</div>
              <div class="components">
                <div class="component">Vue.js Frontend</div>
                <div class="component">Zitadel Auth</div>
              </div>
            </div>
            <div class="layer layer-middle">
              <div class="layer-title">Application Layer</div>
              <div class="components">
                <div class="component">Go API Backend</div>
                <div class="component">Database Operator</div>
              </div>
            </div>
            <div class="layer layer-bottom">
              <div class="layer-title">Infrastructure</div>
              <div class="components">
                <div class="component">Kubernetes Cluster</div>
                <div class="component">OpenStack VMs</div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- Features Section -->
    <div class="section features-section">
      <div class="container">
        <h2 class="section-title">Key Features</h2>
        <div class="features-grid">
          <div class="feature-item">
            <div class="feature-icon">
              <svg width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                <path d="M9 12l2 2 4-4"/>
                <path d="M21 12c0 1.66-4 3-9 3s-9-1.34-9-3"/>
                <path d="M3 5v14c0 1.66 4 3 9 3s9-1.34 9-3V5"/>
                <ellipse cx="12" cy="5" rx="9" ry="3"/>
              </svg>
            </div>
            <h4>Instant Database Provisioning</h4>
            <p>Create production-ready PostgreSQL databases in seconds with automated setup and configuration.</p>
          </div>

          <div class="feature-item">
            <div class="feature-icon">
              <svg width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                <path d="M12 22s8-4 8-10V5l-8-3-8 3v7c0 6 8 10 8 10z"/>
              </svg>
            </div>
            <h4>Enterprise Security</h4>
            <p>End-to-end encryption, network isolation, and role-based access control for maximum data protection.</p>
          </div>

          <div class="feature-item">
            <div class="feature-icon">
              <svg width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                <path d="M13 2L3 14h9l-1 8 10-12h-9l1-8z"/>
              </svg>
            </div>
            <h4>High Availability</h4>
            <p>99.9% uptime guarantee with automated failover, load balancing, and disaster recovery capabilities.</p>
          </div>

          <div class="feature-item">
            <div class="feature-icon">
              <svg width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                <polyline points="23 6 13.5 15.5 8.5 10.5 1 18"/>
                <polyline points="17 6 23 6 23 12"/>
              </svg>
            </div>
            <h4>Auto-Scaling</h4>
            <p>Dynamic resource allocation based on workload demands with seamless horizontal and vertical scaling.</p>
          </div>

          <div class="feature-item">
            <div class="feature-icon">
              <svg width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                <path d="M14 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V8z"/>
                <polyline points="14 2 14 8 20 8"/>
                <line x1="16" y1="13" x2="8" y2="13"/>
                <line x1="16" y1="17" x2="8" y2="17"/>
                <polyline points="10 9 9 9 8 9"/>
              </svg>
            </div>
            <h4>Connection Management</h4>
            <p>Get secure connection strings, credentials, and access information for all your database instances.</p>
          </div>

          <div class="feature-item">
            <div class="feature-icon">
              <svg width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                <rect x="3" y="3" width="18" height="18" rx="2" ry="2"/>
                <rect x="7" y="7" width="3" height="9"/>
                <rect x="14" y="7" width="3" height="5"/>
              </svg>
            </div>
            <h4>Real-time Monitoring</h4>
            <p>Comprehensive dashboards with performance metrics, resource usage, and health monitoring.</p>
          </div>
        </div>
      </div>
    </div>

    <!-- How It Works Section -->
    <div class="section how-it-works">
      <div class="container">
        <h2 class="section-title">How It Works</h2>
        <div class="steps-container">
          <div class="step">
            <div class="step-number">1</div>
            <div class="step-content">
              <h4>Create Account</h4>
              <p>Sign up using Zitadel's secure authentication system with your email or social login.</p>
            </div>
          </div>
          
          <div class="step">
            <div class="step-number">2</div>
            <div class="step-content">
              <h4>Request Database</h4>
              <p>Create a new PostgreSQL database instance with your preferred configuration and resources.</p>
            </div>
          </div>
          
          <div class="step">
            <div class="step-number">3</div>
            <div class="step-content">
              <h4>Get Access</h4>
              <p>Receive secure connection information including host, port, credentials, and SSL certificates.</p>
            </div>
          </div>
          
          <div class="step">
            <div class="step-number">4</div>
            <div class="step-content">
              <h4>Start Building</h4>
              <p>Connect your applications and start building with a fully managed, production-ready database.</p>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- Contact Section -->
    <div class="section contact-section">
      <div class="container">
        <h2 class="section-title">Get in Touch</h2>
        <div class="contact-content">
          <div class="contact-info">
            <h3>Need Help or Have Questions?</h3>
            <p>Our team is here to help you get the most out of the CloudOps platform. Whether you need technical support, have feature requests, or want to discuss enterprise solutions, we're ready to assist.</p>
            
            <div class="contact-methods">
              <div class="contact-method">
                <div class="contact-icon">
                  <svg width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                    <path d="M4 4h16c1.1 0 2 .9 2 2v12c0 1.1-.9 2-2 2H4c-1.1 0-2-.9-2-2V6c0-1.1.9-2 2-2z"/>
                    <polyline points="22,6 12,13 2,6"/>
                  </svg>
                </div>
                <div>
                  <h4>Email Support</h4>
                  <p>contact@cloudops-platform.tech</p>
                  <small>Response within 24 hours</small>
                </div>
              </div>
              
              <div class="contact-method">
                <div class="contact-icon">
                  <svg width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                    <path d="M9 11a3 3 0 1 0 6 0a3 3 0 0 0-6 0"/>
                    <path d="M17.657 16.657L13.414 20.9a1.998 1.998 0 0 1-2.827 0l-4.244-4.243a8 8 0 1 1 11.314 0z"/>
                  </svg>
                </div>
                <div>
                  <h4>Business Inquiries</h4>
                  <p>business@cloudops-platform.tech</p>
                  <small>For partnerships and enterprise solutions</small>
                </div>
              </div>
            </div>
          </div>
          
          <div class="contact-stats">
            <div class="stat-item">
              <div class="stat-number">99.9%</div>
              <div class="stat-label">Uptime SLA</div>
            </div>
            <div class="stat-item">
              <div class="stat-number">< 2min</div>
              <div class="stat-label">Database Provisioning</div>
            </div>
            <div class="stat-item">
              <div class="stat-number">24/7</div>
              <div class="stat-label">System Monitoring</div>
            </div>
            <div class="stat-item">
              <div class="stat-number">Enterprise</div>
              <div class="stat-label">Security Standards</div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'AboutView'
}
</script>

<style scoped>
.about-page {
  min-height: 100vh;
  background: linear-gradient(135deg, #f5f7fa 0%, #c3cfe2 100%);
}

/* Hero Section */
.hero-section {
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  color: white;
  padding: 4rem 0;
  margin-bottom: 0;
}

.hero-content {
  max-width: 1200px;
  margin: 0 auto;
  padding: 0 2rem;
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 4rem;
  align-items: center;
}

.hero-title {
  font-size: 3.5rem;
  font-weight: 700;
  margin-bottom: 1rem;
  background: linear-gradient(45deg, #ffffff, #e0e7ff);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background-clip: text;
}

.hero-subtitle {
  font-size: 1.5rem;
  font-weight: 500;
  margin-bottom: 1.5rem;
  color: #e0e7ff;
}

.hero-description {
  font-size: 1.1rem;
  line-height: 1.7;
  color: #f1f5f9;
  margin-bottom: 2rem;
}

.tech-stack {
  display: flex;
  flex-direction: column;
  gap: 2rem;
}

.tech-item {
  display: flex;
  align-items: center;
  gap: 1rem;
  background: rgba(255, 255, 255, 0.1);
  padding: 1.5rem;
  border-radius: 12px;
  backdrop-filter: blur(10px);
  border: 1px solid rgba(255, 255, 255, 0.2);
}

.tech-item svg {
  color: #ffffff;
}

.tech-item span {
  font-size: 1.1rem;
  font-weight: 500;
}

/* Section Styles */
.section {
  padding: 5rem 0;
}

.container {
  max-width: 1200px;
  margin: 0 auto;
  padding: 0 2rem;
}

.section-title {
  font-size: 2.5rem;
  font-weight: 700;
  color: #2c3e50;
  text-align: center;
  margin-bottom: 3rem;
}

/* Overview Grid */
.overview-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(300px, 1fr));
  gap: 2rem;
}

.overview-card {
  background: white;
  padding: 2.5rem;
  border-radius: 16px;
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.07);
  text-align: center;
  transition: transform 0.3s ease, box-shadow 0.3s ease;
}

.overview-card:hover {
  transform: translateY(-8px);
  box-shadow: 0 20px 40px rgba(0, 0, 0, 0.1);
}

.overview-card .card-icon {
  width: 64px;
  height: 64px;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  color: white;
  margin: 0 auto 1.5rem;
}

.overview-card h3 {
  font-size: 1.4rem;
  font-weight: 600;
  color: #2c3e50;
  margin-bottom: 1rem;
}

.overview-card p {
  color: #6c757d;
  line-height: 1.6;
}

/* Architecture Section */
.architecture-section {
  background: white;
}

.architecture-content {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 4rem;
  align-items: start;
}

.architecture-text h3 {
  font-size: 1.8rem;
  font-weight: 600;
  color: #2c3e50;
  margin-bottom: 1rem;
}

.architecture-text p {
  color: #6c757d;
  line-height: 1.6;
  margin-bottom: 1.5rem;
}

.feature-list {
  list-style: none;
  padding: 0;
}

.feature-list li {
  color: #6c757d;
  line-height: 1.7;
  margin-bottom: 1rem;
  padding-left: 1.5rem;
  position: relative;
}

.feature-list li::before {
  content: "→";
  color: #667eea;
  font-weight: bold;
  position: absolute;
  left: 0;
}

.architecture-diagram {
  background: #f8f9fa;
  border-radius: 16px;
  padding: 2rem;
  border: 2px solid #e9ecef;
}

.layer {
  margin-bottom: 1.5rem;
  padding: 1rem;
  border-radius: 12px;
  border: 2px solid transparent;
}

.layer-top {
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  color: white;
  border-color: #667eea;
}

.layer-middle {
  background: linear-gradient(135deg, #28a745 0%, #20c997 100%);
  color: white;
  border-color: #28a745;
}

.layer-bottom {
  background: linear-gradient(135deg, #ffc107 0%, #fd7e14 100%);
  color: white;
  border-color: #ffc107;
}

.layer-title {
  font-weight: 600;
  margin-bottom: 0.5rem;
  text-align: center;
}

.components {
  display: flex;
  justify-content: center;
  gap: 1rem;
  flex-wrap: wrap;
}

.component {
  background: rgba(255, 255, 255, 0.2);
  padding: 0.5rem 1rem;
  border-radius: 8px;
  font-size: 0.9rem;
  backdrop-filter: blur(10px);
}

/* Features Grid */
.features-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(350px, 1fr));
  gap: 2rem;
}

.feature-item {
  background: white;
  padding: 2rem;
  border-radius: 12px;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.05);
  border-left: 4px solid #667eea;
}

.feature-icon {
  width: 48px;
  height: 48px;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  border-radius: 12px;
  display: flex;
  align-items: center;
  justify-content: center;
  color: white;
  margin-bottom: 1rem;
}

.feature-item h4 {
  font-size: 1.2rem;
  font-weight: 600;
  color: #2c3e50;
  margin-bottom: 0.5rem;
}

.feature-item p {
  color: #6c757d;
  line-height: 1.6;
}

/* How It Works */
.how-it-works {
  background: white;
}

.steps-container {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
  gap: 2rem;
}

.step {
  text-align: center;
  position: relative;
}

.step-number {
  width: 60px;
  height: 60px;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  color: white;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 1.5rem;
  font-weight: 700;
  margin: 0 auto 1.5rem;
}

.step-content h4 {
  font-size: 1.3rem;
  font-weight: 600;
  color: #2c3e50;
  margin-bottom: 0.5rem;
}

.step-content p {
  color: #6c757d;
  line-height: 1.6;
}

/* Contact Section */
.contact-section {
  background: linear-gradient(135deg, #2c3e50 0%, #34495e 100%);
  color: white;
}

.contact-section .section-title {
  color: white;
}

.contact-content {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 4rem;
  align-items: start;
}

.contact-info h3 {
  font-size: 1.8rem;
  font-weight: 600;
  margin-bottom: 1rem;
}

.contact-info p {
  line-height: 1.7;
  margin-bottom: 2rem;
  color: #e0e7ff;
}

.contact-methods {
  display: flex;
  flex-direction: column;
  gap: 2rem;
}

.contact-method {
  display: flex;
  gap: 1rem;
  align-items: flex-start;
}

.contact-icon {
  width: 48px;
  height: 48px;
  background: rgba(255, 255, 255, 0.1);
  border-radius: 12px;
  display: flex;
  align-items: center;
  justify-content: center;
  color: #667eea;
  flex-shrink: 0;
}

.contact-method h4 {
  font-size: 1.2rem;
  font-weight: 600;
  margin-bottom: 0.25rem;
}

.contact-method p {
  color: #667eea;
  font-weight: 500;
  margin-bottom: 0.25rem;
}

.contact-method small {
  color: #cbd5e0;
  font-size: 0.9rem;
}

.contact-stats {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 2rem;
}

.stat-item {
  background: rgba(255, 255, 255, 0.1);
  padding: 2rem;
  border-radius: 16px;
  text-align: center;
  backdrop-filter: blur(10px);
  border: 1px solid rgba(255, 255, 255, 0.2);
}

.stat-number {
  font-size: 2rem;
  font-weight: 700;
  color: #667eea;
  margin-bottom: 0.5rem;
}

.stat-label {
  color: #e0e7ff;
  font-size: 0.9rem;
  font-weight: 500;
}

/* Responsive Design */
@media (max-width: 768px) {
  .hero-content {
    grid-template-columns: 1fr;
    gap: 2rem;
    text-align: center;
  }
  
  .hero-title {
    font-size: 2.5rem;
  }
  
  .architecture-content,
  .contact-content {
    grid-template-columns: 1fr;
    gap: 2rem;
  }
  
  .steps-container {
    grid-template-columns: 1fr;
  }
  
  .contact-stats {
    grid-template-columns: 1fr;
  }
  
  .section {
    padding: 3rem 0;
  }
  
  .container {
    padding: 0 1rem;
  }
}

@media (max-width: 480px) {
  .hero-title {
    font-size: 2rem;
  }
  
  .section-title {
    font-size: 2rem;
  }
  
  .overview-grid {
    grid-template-columns: 1fr;
  }
  
  .features-grid {
    grid-template-columns: 1fr;
  }
}
</style>
