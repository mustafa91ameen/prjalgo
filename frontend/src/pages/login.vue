<template>
  <div class="login-page">
    <div class="login-container">
      <!-- Login Card with Header -->
      <div class="login-box">
        <!-- Header -->
        <div class="login-header">
          <div class="header-gradient-line"></div>
          <div class="header-content">
            <v-icon size="40" color="white">mdi-shield-lock</v-icon>
            <div class="header-text">
              <h1 class="main-title">نظام إدارة المشاريع</h1>
              <p class="subtitle">تسجيل الدخول</p>
            </div>
          </div>
        </div>

        <!-- Login Form -->
        <div class="login-form-section">
          <form @submit.prevent="handleLogin">
            <!-- Error Message -->
            <v-alert
              v-if="errorMessage"
              type="error"
              variant="tonal"
              class="mb-4"
              density="compact"
            >
              {{ errorMessage }}
            </v-alert>

            <!-- Username Field -->
            <v-text-field
              v-model="formData.username"
              label="اسم المستخدم"
              prepend-inner-icon="mdi-account"
              variant="outlined"
              density="comfortable"
              class="mb-4"
              required
              :disabled="loading"
              hide-details="auto"
            ></v-text-field>

            <!-- Password Field -->
            <v-text-field
              v-model="formData.password"
              :type="showPassword ? 'text' : 'password'"
              label="كلمة المرور"
              prepend-inner-icon="mdi-lock"
              :append-inner-icon="showPassword ? 'mdi-eye-off' : 'mdi-eye'"
              @click:append-inner="showPassword = !showPassword"
              variant="outlined"
              density="comfortable"
              class="mb-6"
              required
              :disabled="loading"
              hide-details="auto"
            ></v-text-field>

            <!-- Login Button -->
            <v-btn
              type="submit"
              size="large"
              block
              :loading="loading"
              :disabled="loading"
              class="login-btn"
            >
              <v-icon start>mdi-login</v-icon>
              دخول
            </v-btn>
          </form>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { handleLogin as loginUser } from '@/services/authService'
import { useAuthStore } from '@/stores/auth'

const router = useRouter()
const authStore = useAuthStore()

const formData = ref({
  username: '',
  password: ''
})

const loading = ref(false)
const errorMessage = ref('')
const showPassword = ref(false)

const handleLogin = async () => {
  // Validate fields
  if (!formData.value.username || !formData.value.password) {
    errorMessage.value = 'يرجى إدخال اسم المستخدم وكلمة المرور'
    return
  }

  loading.value = true
  errorMessage.value = ''

  try {
    const session = await loginUser({
      username: formData.value.username.trim(),
      password: formData.value.password,
    })

    // Set user in store
    authStore.setUser({ id: session.userId, username: session.username })

    // Fetch user pages and permissions
    await authStore.fetchPages()

    await router.push('/')
  } catch (error) {
    const status = error?.status
    const message = error?.message || ''
    if (status === 0) {
      errorMessage.value = 'فشل الاتصال بالخادم. يرجى التأكد من تشغيل الخادم والمحاولة مجدداً.'
    } else if (status === 401) {
      if (message.includes('inactive')) {
        errorMessage.value = 'حسابك غير مفعل. يرجى التواصل مع المسؤول.'
      } else {
        errorMessage.value = 'اسم المستخدم أو كلمة المرور غير صحيحة.'
      }
    } else {
      errorMessage.value = error?.message || 'حدث خطأ غير متوقع. يرجى المحاولة مرة أخرى.'
    }
    console.error('Login error:', error)
  } finally {
    loading.value = false
  }
}
</script>

<style scoped>
.login-page {
  display: flex;
  justify-content: center;
  align-items: center;
  min-height: 100vh;
  background: linear-gradient(135deg, #1a2332 0%, #0d1721 100%);
  padding: 20px;
  direction: rtl;
}

.login-container {
  width: 100%;
  max-width: 420px;
}

.login-box {
  border-radius: 20px;
  overflow: hidden;
  box-shadow: 0 25px 50px -12px rgba(0, 0, 0, 0.5), 0 0 0 1px rgba(6, 182, 212, 0.2);
}

/* Header - Matching app theme */
.login-header {
  background: linear-gradient(135deg, #018790 0%, #005461 100%);
  position: relative;
  overflow: hidden;
}

.header-gradient-line {
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  height: 4px;
  background: linear-gradient(90deg, #10b981 0%, #06b6d4 50%, #14b8a6 100%);
}

.header-content {
  padding: 28px 24px;
  display: flex;
  align-items: center;
  gap: 16px;
}

.header-text {
  display: flex;
  flex-direction: column;
  gap: 6px;
}

.main-title {
  color: white;
  font-size: 1.35rem;
  font-weight: 700;
  margin: 0;
  letter-spacing: 0.3px;
  text-shadow: 0 2px 4px rgba(0, 0, 0, 0.3);
}

.subtitle {
  color: rgba(255, 255, 255, 0.85);
  font-size: 0.9rem;
  margin: 0;
  font-weight: 500;
}

/* Form Section - Matching app dark theme */
.login-form-section {
  background: linear-gradient(135deg, #0a3d42 0%, #052428 100%);
  padding: 28px 24px;
}

/* Text field styling */
.login-form-section :deep(.v-field) {
  background: rgba(255, 255, 255, 0.08) !important;
  border-radius: 12px !important;
}

.login-form-section :deep(.v-field__input) {
  color: rgba(255, 255, 255, 0.95) !important;
  direction: rtl !important;
  text-align: right !important;
}

.login-form-section :deep(.v-label) {
  color: rgba(255, 255, 255, 0.7) !important;
  right: 12px !important;
  left: auto !important;
}

.login-form-section :deep(.v-field--focused .v-label) {
  color: #10b981 !important;
}

.login-form-section :deep(.v-field__outline) {
  color: rgba(6, 182, 212, 0.3) !important;
}

.login-form-section :deep(.v-field--focused .v-field__outline) {
  color: #10b981 !important;
}

.login-form-section :deep(.v-field:hover .v-field__outline) {
  color: rgba(6, 182, 212, 0.5) !important;
}

.login-form-section :deep(.v-icon) {
  color: rgba(255, 255, 255, 0.6) !important;
}

.login-form-section :deep(.v-field--focused .v-icon) {
  color: #10b981 !important;
}

/* RTL field adjustments */
.login-form-section :deep(.v-field__prepend-inner) {
  padding-right: 0 !important;
  padding-left: 8px !important;
}

.login-form-section :deep(.v-field__append-inner) {
  padding-left: 0 !important;
  padding-right: 8px !important;
}

/* Login Button - Matching app green gradient */
.login-btn {
  background: linear-gradient(135deg, #10b981 0%, #059669 100%) !important;
  color: white !important;
  font-weight: 600 !important;
  letter-spacing: 0.5px;
  border-radius: 12px !important;
  box-shadow: 0 4px 15px rgba(16, 185, 129, 0.3) !important;
  transition: all 0.3s ease !important;
}

.login-btn:hover {
  background: linear-gradient(135deg, #059669 0%, #047857 100%) !important;
  box-shadow: 0 6px 20px rgba(16, 185, 129, 0.4) !important;
  transform: translateY(-2px);
}

.login-btn :deep(.v-icon) {
  color: white !important;
}

/* Error alert styling */
.login-form-section :deep(.v-alert) {
  background: rgba(239, 68, 68, 0.15) !important;
  border: 1px solid rgba(239, 68, 68, 0.3) !important;
  border-radius: 12px !important;
  direction: rtl !important;
}

/* Responsive */
@media (max-width: 768px) {
  .login-page {
    padding: 16px;
  }

  .login-container {
    max-width: 100%;
  }

  .header-content {
    padding: 24px 20px;
    gap: 12px;
  }

  .main-title {
    font-size: 1.2rem;
  }

  .subtitle {
    font-size: 0.85rem;
  }

  .login-form-section {
    padding: 24px 20px;
  }
}

@media (max-width: 480px) {
  .login-page {
    padding: 12px;
  }

  .login-container {
    max-width: 100%;
  }

  .login-box {
    border-radius: 16px;
  }

  .header-content {
    padding: 20px 16px;
    gap: 10px;
    flex-direction: column;
    text-align: center;
  }

  .header-content .v-icon {
    font-size: 36px !important;
  }

  .main-title {
    font-size: 1.15rem;
  }

  .subtitle {
    font-size: 0.8rem;
  }

  .login-form-section {
    padding: 20px 16px;
  }

  .login-btn {
    font-size: 0.95rem !important;
    border-radius: 10px !important;
  }
}

@media (max-width: 360px) {
  .login-page {
    padding: 8px;
  }

  .header-content {
    padding: 16px 12px;
  }

  .main-title {
    font-size: 1.05rem;
  }

  .login-form-section {
    padding: 16px 12px;
  }
}
</style>
