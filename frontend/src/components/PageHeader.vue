<template>
  <div class="ds-page-header">
    <div class="ds-page-header-content">
      <div class="ds-page-header-text">
        <h1 class="ds-page-header-title">
          <v-icon v-if="icon" class="ds-page-header-icon" :size="iconSize" color="white">
            {{ icon }}
          </v-icon>
          {{ title }}
        </h1>
        <p v-if="subtitle" class="ds-page-header-subtitle">{{ subtitle }}</p>
      </div>
      <v-btn
        v-if="showBackButton"
        icon="mdi-arrow-right"
        @click="handleBack"
        class="ds-page-header-back-btn"
        size="small"
      >
        <v-icon size="20" color="white">mdi-arrow-right</v-icon>
      </v-btn>
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { useRouter } from 'vue-router'

const props = defineProps({
  title: {
    type: String,
    required: true
  },
  subtitle: {
    type: String,
    default: ''
  },
  icon: {
    type: String,
    default: ''
  },
  iconSize: {
    type: [String, Number],
    default: 18
  },
  showBackButton: {
    type: Boolean,
    default: false
  },
  backRoute: {
    type: String,
    default: null
  }
})

const router = useRouter()

const handleBack = () => {
  if (props.backRoute) {
    router.push(props.backRoute)
  } else {
    router.back()
  }
}
</script>

<style scoped>
.ds-page-header {
  background: var(--ds-gradient-header);
  background-size: 400% 400%;
  animation: gradientShift 8s ease infinite;
  color: var(--ds-text-white);
  padding: var(--ds-page-header-padding);
  border-radius: 0;
  margin: 0;
  margin-bottom: 1rem;
  margin-top: 0;
  position: relative;
  overflow: hidden;
  box-shadow: var(--ds-shadow-md);
  border: none;
  border-bottom: 1px solid rgba(255, 255, 255, 0.3);
  backdrop-filter: blur(15px);
  z-index: var(--ds-z-base);
  width: 100%;
  max-width: 100%;
  box-sizing: border-box;
}

.ds-page-header-content {
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 1rem;
  font-family: var(--ds-font-family);
  max-width: 100%;
}

.ds-page-header-text {
  flex: 1;
}

.ds-page-header-title {
  font-size: var(--ds-page-header-title-size);
  font-weight: var(--ds-font-weight-semibold);
  margin: 0;
  text-shadow: 0 2px 6px rgba(0, 0, 0, 0.3), 0 0 0 1px rgba(0, 0, 0, 0.1);
  color: var(--ds-text-white);
  letter-spacing: var(--ds-letter-spacing-normal);
  font-family: var(--ds-font-family);
  line-height: var(--ds-line-height-tight);
  padding: 0.2rem 0.6rem;
  border-radius: var(--ds-radius-md);
  display: inline-flex;
  align-items: center;
  gap: 0.5rem;
  backdrop-filter: blur(10px);
  box-shadow: 0 2px 8px rgba(26, 35, 126, 0.3);
}

.ds-page-header-icon {
  color: var(--ds-text-white);
}

.ds-page-header-subtitle {
  font-size: var(--ds-page-header-subtitle-size);
  opacity: 0.9;
  margin: 0.25rem 0 0 0;
  color: var(--ds-text-white);
  text-shadow: 0 2px 8px rgba(0, 0, 0, 0.5);
  font-weight: var(--ds-font-weight-medium);
  letter-spacing: var(--ds-letter-spacing-normal);
  font-family: var(--ds-font-family);
  line-height: var(--ds-line-height-tight);
}

.ds-page-header-back-btn {
  background: rgba(255, 255, 255, 0.25);
  backdrop-filter: blur(15px);
  border: 1px solid rgba(255, 255, 255, 0.3);
  box-shadow: var(--ds-shadow-sm);
  transition: var(--ds-transition-normal);
  min-width: 36px;
  height: 36px;
}

.ds-page-header-back-btn:hover {
  background: rgba(255, 255, 255, 0.35);
  transform: translateX(2px);
  box-shadow: var(--ds-shadow-md);
}

@keyframes gradientShift {
  0% { background-position: 0% 50%; }
  50% { background-position: 100% 50%; }
  100% { background-position: 0% 50%; }
}
</style>

