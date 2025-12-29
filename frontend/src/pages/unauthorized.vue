<template>
  <div class="unauthorized-page">
    <div class="unauthorized-container">
      <div class="unauthorized-box">
        <!-- Header -->
        <div class="unauthorized-header">
          <div class="header-gradient-line"></div>
          <div class="header-content">
            <v-icon size="48" color="white">mdi-shield-off</v-icon>
            <div class="header-text">
              <h1 class="main-title">غير مصرح</h1>
              <p class="subtitle">ليس لديك صلاحية للوصول لهذه الصفحة</p>
            </div>
          </div>
        </div>

        <!-- Content Section -->
        <div class="unauthorized-content">
          <div class="message-box">
            <v-icon size="64" color="#dc2626" class="mb-4">mdi-lock-alert</v-icon>
            <h2 class="error-code">403</h2>
            <p class="error-message">عذراً، ليس لديك الصلاحيات اللازمة للوصول إلى هذه الصفحة.</p>
            <p class="error-hint">إذا كنت تعتقد أن هذا خطأ، يرجى التواصل مع مسؤول النظام.</p>
          </div>

          <div class="actions">
            <v-btn
              size="large"
              class="back-btn"
              @click="goBack"
            >
              <v-icon start>mdi-arrow-right</v-icon>
              العودة للخلف
            </v-btn>
            <v-btn
              size="large"
              variant="outlined"
              class="home-btn"
              @click="goHome"
            >
              <v-icon start>mdi-home</v-icon>
              الصفحة الرئيسية
            </v-btn>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { useRouter } from 'vue-router'

const router = useRouter()

const goBack = () => {
  if (window.history.length > 2) {
    router.back()
  } else {
    router.push('/')
  }
}

const goHome = () => {
  router.push('/')
}
</script>

<style scoped>
.unauthorized-page {
  display: flex;
  justify-content: center;
  align-items: center;
  min-height: 100vh;
  background: linear-gradient(135deg, #fef2f2 0%, #fee2e2 50%, #fecaca 100%);
  padding: 20px;
}

.unauthorized-container {
  width: 100%;
  max-width: 480px;
}

.unauthorized-box {
  border-radius: 16px;
  overflow: hidden;
  box-shadow: 0 10px 25px -5px rgba(220, 38, 38, 0.2), 0 8px 10px -6px rgba(220, 38, 38, 0.1);
}

/* Header - Red gradient for error state */
.unauthorized-header {
  background: linear-gradient(135deg, #dc2626 0%, #ef4444 100%);
  position: relative;
  overflow: hidden;
}

.header-gradient-line {
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  height: 4px;
  background: linear-gradient(90deg, #fca5a5 0%, #fecaca 50%, #fee2e2 100%);
}

.header-content {
  padding: 24px;
  display: flex;
  align-items: center;
  gap: 16px;
}

.header-text {
  display: flex;
  flex-direction: column;
  gap: 4px;
}

.main-title {
  color: white;
  font-size: 1.5rem;
  font-weight: 600;
  margin: 0;
  letter-spacing: 0.3px;
  text-shadow: 0 1px 2px rgba(0, 0, 0, 0.2);
}

.subtitle {
  color: rgba(255, 255, 255, 0.9);
  font-size: 0.875rem;
  margin: 0;
}

/* Content Section */
.unauthorized-content {
  background: rgba(255, 255, 255, 0.95);
  backdrop-filter: blur(10px);
  padding: 32px 24px;
}

.message-box {
  text-align: center;
  margin-bottom: 24px;
}

.error-code {
  font-size: 3rem;
  font-weight: 700;
  color: #dc2626;
  margin: 0 0 8px 0;
}

.error-message {
  font-size: 1rem;
  color: #1e293b;
  margin: 0 0 8px 0;
  line-height: 1.6;
}

.error-hint {
  font-size: 0.875rem;
  color: #64748b;
  margin: 0;
}

/* Actions */
.actions {
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.back-btn {
  background: linear-gradient(135deg, #4338ca 0%, #6366f1 100%) !important;
  color: white !important;
  font-weight: 600 !important;
  letter-spacing: 0.5px;
  box-shadow: 0 4px 6px -1px rgba(67, 56, 202, 0.3) !important;
}

.back-btn:hover {
  background: linear-gradient(135deg, #3730a3 0%, #4338ca 100%) !important;
  box-shadow: 0 6px 12px rgba(67, 56, 202, 0.4) !important;
}

.back-btn :deep(.v-icon) {
  color: white !important;
}

.home-btn {
  border-color: #4338ca !important;
  color: #4338ca !important;
  font-weight: 600 !important;
}

.home-btn:hover {
  background: rgba(67, 56, 202, 0.1) !important;
}

/* Responsive */
@media (max-width: 768px) {
  .unauthorized-page {
    padding: 16px;
  }

  .unauthorized-container {
    max-width: 100%;
  }

  .header-content {
    padding: 20px;
    gap: 12px;
  }

  .main-title {
    font-size: 1.25rem;
  }

  .error-code {
    font-size: 2.5rem;
  }

  .unauthorized-content {
    padding: 24px 20px;
  }
}

@media (max-width: 480px) {
  .unauthorized-page {
    padding: 12px;
  }

  .header-content {
    padding: 16px;
    flex-direction: column;
    text-align: center;
  }

  .main-title {
    font-size: 1.15rem;
  }

  .subtitle {
    font-size: 0.8rem;
  }

  .error-code {
    font-size: 2rem;
  }

  .error-message {
    font-size: 0.9rem;
  }

  .unauthorized-content {
    padding: 20px 16px;
  }
}
</style>
