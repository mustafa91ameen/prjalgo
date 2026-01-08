<template>
  <div class="unauthorized-page">
    <div class="unauthorized-container">
      <div class="unauthorized-content">
        <!-- Icon -->
        <div class="icon-wrapper">
          <v-icon size="80" color="error">mdi-shield-lock-outline</v-icon>
        </div>

        <!-- Title -->
        <h1 class="unauthorized-title">غير مصرح</h1>

        <!-- Description -->
        <p class="unauthorized-description">
          عذراً، ليس لديك صلاحية الوصول إلى هذه الصفحة.
        </p>
        <p class="unauthorized-subdescription">
          يرجى التواصل مع مدير النظام إذا كنت تعتقد أن هذا خطأ.
        </p>

        <!-- Actions -->
        <div class="unauthorized-actions">
          <v-btn
            color="primary"
            size="large"
            @click="goHome"
            class="action-btn"
          >
            <v-icon start>mdi-home</v-icon>
            العودة للرئيسية
          </v-btn>

          <v-btn
            variant="outlined"
            color="secondary"
            size="large"
            @click="goBack"
            class="action-btn"
          >
            <v-icon start>mdi-arrow-right</v-icon>
            العودة للخلف
          </v-btn>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { useRouter } from 'vue-router'

const router = useRouter()

const goHome = () => {
  router.push('/')
}

const goBack = () => {
  router.back()
}
</script>

<style scoped>
.unauthorized-page {
  display: flex;
  justify-content: center;
  align-items: center;
  min-height: 100vh;
  background: linear-gradient(135deg, #0f172a 0%, #1e293b 50%, #0f172a 100%);
  padding: 20px;
}

.unauthorized-container {
  width: 100%;
  max-width: 500px;
}

.unauthorized-content {
  background: rgba(30, 41, 59, 0.95);
  backdrop-filter: blur(10px);
  border-radius: 16px;
  padding: 48px 32px;
  text-align: center;
  box-shadow: 0 10px 25px -5px rgba(0, 0, 0, 0.3);
  border: 1px solid rgba(255, 255, 255, 0.1);
}

.icon-wrapper {
  margin-bottom: 24px;
}

.unauthorized-title {
  color: #f87171;
  font-size: 2rem;
  font-weight: 700;
  margin: 0 0 16px 0;
}

.unauthorized-description {
  color: #e2e8f0;
  font-size: 1.1rem;
  margin: 0 0 8px 0;
  line-height: 1.6;
}

.unauthorized-subdescription {
  color: #94a3b8;
  font-size: 0.95rem;
  margin: 0 0 32px 0;
  line-height: 1.5;
}

.unauthorized-actions {
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.action-btn {
  text-transform: none !important;
  font-weight: 600 !important;
}

/* Responsive */
@media (max-width: 480px) {
  .unauthorized-content {
    padding: 32px 24px;
  }

  .unauthorized-title {
    font-size: 1.5rem;
  }

  .unauthorized-description {
    font-size: 1rem;
  }

  .unauthorized-subdescription {
    font-size: 0.9rem;
  }
}
</style>
