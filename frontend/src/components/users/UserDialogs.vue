<template>
  <!-- Reset Password Dialog -->
  <v-dialog :model-value="resetPasswordDialog" @update:model-value="$emit('update:resetPasswordDialog', $event)" max-width="500px">
    <v-card class="reset-password-dialog">
      <v-card-title class="dialog-header">
        <div class="dialog-title">
          <v-icon size="32" color="warning" class="me-3">mdi-key-change</v-icon>
          <h2>إعادة تعيين كلمة المرور</h2>
        </div>
        <v-btn
          icon="mdi-close"
          variant="text"
          @click="$emit('update:resetPasswordDialog', false)"
          class="close-btn"
        />
      </v-card-title>

      <v-divider />

      <v-card-text v-if="user" class="pa-6">
        <div class="text-center mb-4">
          <v-avatar size="60">
            <v-img :src="user.avatar || defaultAvatar" />
          </v-avatar>
          <h4 class="mt-2">{{ user.name }}</h4>
          <p class="text-caption">{{ user.email }}</p>
        </div>

        <v-alert
          type="warning"
          variant="tonal"
          class="mb-4"
        >
          سيتم إرسال كلمة مرور جديدة إلى البريد الإلكتروني للمستخدم
        </v-alert>

        <p class="text-body-2 text-center">
          هل أنت متأكد من إعادة تعيين كلمة المرور لهذا المستخدم؟
        </p>
      </v-card-text>

      <v-divider />

      <v-card-actions class="dialog-actions">
        <v-spacer />
        <v-btn
          color="grey"
          variant="outlined"
          @click="$emit('update:resetPasswordDialog', false)"
          class="me-2"
        >
          إلغاء
        </v-btn>
        <v-btn
          color="warning"
          variant="elevated"
          @click="$emit('confirm-reset')"
          :loading="resetLoading"
        >
          إعادة تعيين
        </v-btn>
      </v-card-actions>
    </v-card>
  </v-dialog>

  <!-- Delete Confirmation Dialog -->
  <v-dialog :model-value="deleteDialog" @update:model-value="$emit('update:deleteDialog', $event)" max-width="500px">
    <v-card class="delete-confirm-dialog">
      <v-card-title class="dialog-header">
        <div class="dialog-title">
          <v-icon size="32" color="error" class="me-3">mdi-delete-alert</v-icon>
          <h2>تأكيد الحذف</h2>
        </div>
        <v-btn
          icon="mdi-close"
          variant="text"
          @click="$emit('update:deleteDialog', false)"
          class="close-btn"
        />
      </v-card-title>

      <v-divider />

      <v-card-text v-if="user" class="pa-6">
        <div class="text-center mb-4">
          <v-avatar size="60">
            <v-img :src="user.avatar || defaultAvatar" />
          </v-avatar>
          <h4 class="mt-2">{{ user.name }}</h4>
          <p class="text-caption">{{ user.email }}</p>
        </div>

        <v-alert
          type="error"
          variant="tonal"
          class="mb-4"
        >
          تحذير: هذا الإجراء لا يمكن التراجع عنه!
        </v-alert>

        <p class="text-body-2 text-center">
          هل أنت متأكد من حذف هذا المستخدم نهائياً؟
        </p>
      </v-card-text>

      <v-divider />

      <v-card-actions class="dialog-actions">
        <v-spacer />
        <v-btn
          color="grey"
          variant="outlined"
          @click="$emit('update:deleteDialog', false)"
          class="me-2"
        >
          إلغاء
        </v-btn>
        <v-btn
          color="error"
          variant="elevated"
          @click="$emit('confirm-delete')"
          :loading="deleteLoading"
        >
          حذف نهائي
        </v-btn>
      </v-card-actions>
    </v-card>
  </v-dialog>
</template>

<script setup>
defineProps({
  resetPasswordDialog: {
    type: Boolean,
    default: false
  },
  deleteDialog: {
    type: Boolean,
    default: false
  },
  user: {
    type: Object,
    default: null
  },
  resetLoading: {
    type: Boolean,
    default: false
  },
  deleteLoading: {
    type: Boolean,
    default: false
  }
})

defineEmits([
  'update:resetPasswordDialog',
  'update:deleteDialog',
  'confirm-reset',
  'confirm-delete'
])

const defaultAvatar = 'https://ui-avatars.com/api/?name=User&background=667eea&color=fff&size=128'
</script>

<style scoped>
@import './styles/users.css';
</style>
