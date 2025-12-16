<template>
  <v-dialog v-model="isOpen" max-width="600" persistent>
    <v-card class="modern-dialog">
      <!-- Header -->
      <v-card-title class="dialog-header">
        <div class="header-content">
          <v-icon class="header-icon" color="white">{{ icon }}</v-icon>
          <span class="header-title">{{ title }}</span>
        </div>
        <v-btn
          icon="mdi-close"
          variant="text"
          color="white"
          @click="closeDialog"
          class="close-btn"
        >
          <v-icon>mdi-close</v-icon>
        </v-btn>
      </v-card-title>

      <!-- Content -->
      <v-card-text class="dialog-content">
        <slot />
      </v-card-text>

      <!-- Footer -->
      <v-card-actions class="dialog-actions">
        <v-spacer></v-spacer>
        <v-btn
          color="primary"
          variant="elevated"
          @click="saveAction"
          class="save-btn"
        >
          <v-icon class="me-2">mdi-content-save</v-icon>
          {{ saveText }}
        </v-btn>
        <v-btn
          color="grey"
          variant="outlined"
          @click="closeDialog"
          class="cancel-btn"
        >
          إلغاء
        </v-btn>
      </v-card-actions>
    </v-card>
  </v-dialog>
</template>

<script setup>
import { computed } from 'vue'

const props = defineProps({
  open: {
    type: Boolean,
    default: false
  },
  title: {
    type: String,
    default: 'إضافة عنصر'
  },
  saveText: {
    type: String,
    default: 'حفظ'
  },
  icon: {
    type: String,
    default: 'mdi-plus'
  }
})

const emit = defineEmits(['close', 'save'])

const isOpen = computed({
  get: () => props.open,
  set: (value) => {
    if (!value) {
      emit('close')
    }
  }
})

const closeDialog = () => {
  emit('close')
}

const saveAction = () => {
  emit('save')
}
</script>

<style scoped>
.modern-dialog {
  border-radius: var(--radius-xl);
  overflow: hidden;
  box-shadow: var(--shadow-xl);
}

.dialog-header {
  background: var(--gradient-info);
  color: var(--text-white);
  padding: 0.75rem 1rem !important;
  display: flex;
  align-items: center;
  justify-content: space-between;
}

.header-content {
  display: flex;
  align-items: center;
  gap: 0.5rem;
}

.header-icon {
  font-size: var(--font-size-xl) !important;
}

.header-title {
  font-size: var(--font-size-base) !important;
  font-weight: 700 !important;
  line-height: 1.4 !important;
}

.close-btn {
  opacity: 0.9;
  transition: opacity 0.3s ease;
  min-width: var(--space-8) !important;
  width: var(--space-8) !important;
  height: var(--space-8) !important;
  padding: 0.25rem !important;
}

.close-btn :deep(.v-icon) {
  font-size: var(--font-size-xl) !important;
}

.close-btn:hover {
  opacity: 1;
  background: rgba(255, 255, 255, 0.1);
}

.dialog-content {
  padding: 2rem;
  background: white;
}

.dialog-actions {
  padding: 1.25rem 1.5rem !important;
  background: var(--color-slate-50) !important;
  border-top: var(--space-px) solid var(--border-light) !important;
  display: flex;
  justify-content: flex-end;
  gap: var(--space-4);
}

.save-btn,
.save-btn.v-btn {
  font-weight: 700 !important;
  text-transform: none !important;
  border-radius: var(--radius-xl) !important;
  padding: var(--space-3-5) 32px !important;
  background: var(--gradient-info) !important;
  color: var(--text-white) !important;
  font-size: var(--font-size-base) !important;
  font-family: 'Cairo', 'Tajawal', 'Arial', sans-serif !important;
  min-width: var(--space-36) !important;
  box-shadow: 
    0 6px 20px rgba(59, 130, 246, 0.4),
    0 3px 10px rgba(37, 99, 235, 0.3),
    inset 0 1px 0 rgba(255, 255, 255, 0.25) !important;
  border: 2px solid rgba(255, 255, 255, 0.4) !important;
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1) !important;
  letter-spacing: 0.5px !important;
  line-height: 1.5 !important;
  position: relative;
  overflow: hidden;
  display: inline-flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
  gap: var(--space-2-5);
}

.save-btn :deep(.v-btn__content) {
  color: var(--text-white) !important;
  font-weight: 700 !important;
  font-size: var(--font-size-base) !important;
  letter-spacing: 0.5px !important;
  display: flex !important;
  align-items: center !important;
  justify-content: center !important;
  gap: var(--space-2-5) !important;
}

.save-btn :deep(.v-icon) {
  color: var(--text-white) !important;
  font-size: var(--font-size-xl) !important;
  flex-shrink: 0;
}

.save-btn::before {
  content: '';
  position: absolute;
  top: 0;
  left: -100%;
  width: 100%;
  height: 100%;
  background: linear-gradient(90deg, transparent, rgba(255, 255, 255, 0.3), transparent);
  transition: left 0.6s ease;
  z-index: 1;
}

.save-btn:hover::before {
  left: 100%;
}

.save-btn:hover,
.save-btn.v-btn:hover {
  background: var(--gradient-info-deep) !important;
  box-shadow: 
    0 8px 24px rgba(59, 130, 246, 0.5),
    0 4px 12px rgba(37, 99, 235, 0.4),
    inset 0 1px 0 rgba(255, 255, 255, 0.4) !important;
  transform: translateY(-2px) scale(1.02) !important;
  border-color: rgba(255, 255, 255, 0.5) !important;
}

.cancel-btn,
.cancel-btn.v-btn {
  font-weight: 700 !important;
  text-transform: none !important;
  border-radius: var(--radius-xl) !important;
  padding: var(--space-3-5) 32px !important;
  border: 2px solid var(--color-slate-300) !important;
  color: var(--color-slate-700) !important;
  background: var(--text-white) !important;
  backdrop-filter: blur(10px) !important;
  font-size: var(--font-size-base) !important;
  font-family: 'Cairo', 'Tajawal', 'Arial', sans-serif !important;
  min-width: var(--space-36) !important;
  box-shadow: 
    0 4px 16px rgba(0, 0, 0, 0.1),
    0 2px 8px rgba(156, 163, 175, 0.15),
    inset 0 1px 0 rgba(255, 255, 255, 1) !important;
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1) !important;
  letter-spacing: 0.5px !important;
  line-height: 1.5 !important;
  position: relative;
  overflow: hidden;
  display: inline-flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
}

.cancel-btn :deep(.v-btn__content) {
  color: var(--color-slate-700) !important;
  font-weight: 700 !important;
  font-size: var(--font-size-base) !important;
  letter-spacing: 0.5px !important;
}

.cancel-btn::before {
  content: '';
  position: absolute;
  top: 0;
  left: -100%;
  width: 100%;
  height: 100%;
  background: linear-gradient(90deg, transparent, rgba(255, 255, 255, 0.2), transparent);
  transition: left 0.5s;
}

.cancel-btn:hover::before {
  left: 100%;
}

.cancel-btn:hover,
.cancel-btn.v-btn:hover {
  background: var(--text-white) !important;
  border-color: var(--color-slate-400) !important;
  color: var(--color-slate-800) !important;
  transform: translateY(-2px) scale(1.02) !important;
  box-shadow: 
    0 6px 20px rgba(0, 0, 0, 0.12),
    0 3px 10px rgba(156, 163, 175, 0.15),
    inset 0 1px 0 rgba(255, 255, 255, 1) !important;
}
</style>
