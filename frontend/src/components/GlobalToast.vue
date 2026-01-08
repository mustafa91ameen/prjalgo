<template>
  <Teleport to="body">
    <Transition name="toast">
      <div v-if="toast.show" :class="['custom-toast', toast.color]">
        <div class="toast-icon">
          <i :class="getToastIcon(toast.color)"></i>
        </div>
        <div class="toast-content">
          <span class="toast-message">{{ toast.message }}</span>
        </div>
        <button class="toast-close" @click="hideToast">
          <i class="mdi mdi-close"></i>
        </button>
      </div>
    </Transition>
  </Teleport>
</template>

<script setup>
import { useToast, getToastIcon } from '@/composables/useToast'

const { toast, hideToast } = useToast()
</script>

<style scoped>
/* Custom Toast Notification Styles */
.custom-toast {
  position: fixed;
  top: 24px;
  left: 50%;
  transform: translateX(-50%);
  display: flex;
  align-items: center;
  gap: 12px;
  padding: 14px 20px;
  border-radius: 12px;
  min-width: 320px;
  max-width: 500px;
  z-index: 99999;
  box-shadow: 0 8px 32px rgba(0, 0, 0, 0.3),
              0 0 0 1px rgba(255, 255, 255, 0.1);
  backdrop-filter: blur(16px);
  direction: rtl;
  text-align: right;
}

.custom-toast.success {
  background: linear-gradient(135deg, rgba(16, 185, 129, 0.95) 0%, rgba(5, 150, 105, 0.95) 100%);
  border: 1px solid rgba(52, 211, 153, 0.3);
}

.custom-toast.error {
  background: linear-gradient(135deg, rgba(220, 38, 38, 0.95) 0%, rgba(185, 28, 28, 0.95) 100%);
  border: 1px solid rgba(248, 113, 113, 0.3);
}

.custom-toast.warning {
  background: linear-gradient(135deg, rgba(245, 158, 11, 0.95) 0%, rgba(217, 119, 6, 0.95) 100%);
  border: 1px solid rgba(251, 191, 36, 0.3);
}

.custom-toast.info {
  background: linear-gradient(135deg, rgba(20, 184, 166, 0.95) 0%, rgba(13, 148, 136, 0.95) 100%);
  border: 1px solid rgba(45, 212, 191, 0.3);
}

.toast-icon {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 32px;
  height: 32px;
  border-radius: 50%;
  background: rgba(255, 255, 255, 0.2);
  flex-shrink: 0;
}

.toast-icon i {
  font-size: 18px;
  color: white;
}

.toast-content {
  flex: 1;
}

.toast-message {
  color: white;
  font-size: 14px;
  font-weight: 500;
  line-height: 1.5;
}

.toast-close {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 28px;
  height: 28px;
  border: none;
  background: rgba(255, 255, 255, 0.15);
  border-radius: 8px;
  cursor: pointer;
  transition: all 0.2s ease;
  flex-shrink: 0;
}

.toast-close:hover {
  background: rgba(255, 255, 255, 0.25);
}

.toast-close i {
  font-size: 16px;
  color: white;
}

/* Toast Animation */
.toast-enter-active {
  animation: toast-in 0.4s cubic-bezier(0.4, 0, 0.2, 1);
}

.toast-leave-active {
  animation: toast-out 0.3s cubic-bezier(0.4, 0, 0.2, 1);
}

@keyframes toast-in {
  0% {
    opacity: 0;
    transform: translateX(-50%) translateY(-20px) scale(0.95);
  }
  100% {
    opacity: 1;
    transform: translateX(-50%) translateY(0) scale(1);
  }
}

@keyframes toast-out {
  0% {
    opacity: 1;
    transform: translateX(-50%) translateY(0) scale(1);
  }
  100% {
    opacity: 0;
    transform: translateX(-50%) translateY(-10px) scale(0.95);
  }
}

/* Responsive Styles */
@media (max-width: 768px) {
  .custom-toast {
    min-width: 280px;
    max-width: 90vw;
    padding: 12px 16px;
  }
}

@media (max-width: 480px) {
  .custom-toast {
    min-width: unset;
    width: calc(100vw - 24px);
    max-width: unset;
    top: 12px;
    padding: 10px 14px;
  }

  .toast-message {
    font-size: 13px;
  }
}
</style>
