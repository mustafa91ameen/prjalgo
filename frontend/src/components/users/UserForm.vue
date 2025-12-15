<template>
  <v-dialog :model-value="modelValue" @update:model-value="$emit('update:modelValue', $event)" max-width="800px" persistent>
    <v-card class="add-user-dialog">
      <v-card-title class="dialog-header">
        <div class="dialog-title">
          <v-icon size="32" :color="isEdit ? 'success' : 'primary'" class="me-3">
            {{ isEdit ? 'mdi-account-edit' : 'mdi-account-plus' }}
          </v-icon>
          <h2>{{ isEdit ? 'تعديل المستخدم' : 'إضافة مستخدم جديد' }}</h2>
        </div>
        <v-btn
          icon="mdi-close"
          variant="text"
          @click="closeDialog"
          class="close-btn"
        />
      </v-card-title>

      <v-divider />

      <v-card-text class="dialog-content">
        <v-form ref="formRef" v-model="formValid" lazy-validation>
          <v-row>
            <!-- Avatar -->
            <v-col cols="12" class="text-center mb-4">
              <v-avatar size="100" class="user-avatar-upload">
                <v-img
                  :src="localUser.avatar || defaultAvatar"
                  alt="صورة المستخدم"
                />
              </v-avatar>
              <div class="mt-2">
                <v-btn
                  size="small"
                  color="primary"
                  variant="outlined"
                  prepend-icon="mdi-camera"
                >
                  تحديد صورة
                </v-btn>
              </div>
            </v-col>

            <!-- First Name (Add mode) -->
            <v-col v-if="!isEdit" cols="12" md="6">
              <v-text-field
                v-model="localUser.firstName"
                label="الاسم الأول *"
                :rules="nameRules"
                required
                variant="outlined"
                prepend-inner-icon="mdi-account"
              />
            </v-col>

            <!-- Last Name (Add mode) -->
            <v-col v-if="!isEdit" cols="12" md="6">
              <v-text-field
                v-model="localUser.lastName"
                label="الاسم الأخير *"
                :rules="nameRules"
                required
                variant="outlined"
                prepend-inner-icon="mdi-account"
              />
            </v-col>

            <!-- Full Name (Edit mode) -->
            <v-col v-if="isEdit" cols="12" md="6">
              <v-text-field
                v-model="localUser.name"
                label="الاسم *"
                :rules="nameRules"
                required
                variant="outlined"
                prepend-inner-icon="mdi-account"
              />
            </v-col>

            <!-- Email -->
            <v-col cols="12" md="6">
              <v-text-field
                v-model="localUser.email"
                label="البريد الإلكتروني *"
                :rules="emailRules"
                required
                variant="outlined"
                prepend-inner-icon="mdi-email"
                type="email"
              />
            </v-col>

            <!-- Phone -->
            <v-col cols="12" md="6">
              <v-text-field
                v-model="localUser.phone"
                label="رقم الهاتف"
                variant="outlined"
                prepend-inner-icon="mdi-phone"
                type="tel"
              />
            </v-col>

            <!-- Role -->
            <v-col cols="12" md="6">
              <v-select
                v-model="localUser.role"
                :items="roleOptions"
                label="الدور *"
                :rules="requiredRules"
                required
                variant="outlined"
                prepend-inner-icon="mdi-account-tie"
                class="black-dropdown-select"
              />
            </v-col>

            <!-- Department -->
            <v-col cols="12" md="6">
              <v-select
                v-model="localUser.department"
                :items="departmentOptions"
                label="القسم *"
                :rules="requiredRules"
                required
                variant="outlined"
                prepend-inner-icon="mdi-office-building"
                class="black-dropdown-select"
              />
            </v-col>

            <!-- Status -->
            <v-col cols="12" md="6">
              <v-select
                v-model="localUser.status"
                :items="statusOptions"
                label="الحالة *"
                :rules="requiredRules"
                required
                variant="outlined"
                prepend-inner-icon="mdi-account-check"
                class="black-dropdown-select"
              />
            </v-col>

            <!-- Password (Add mode only) -->
            <v-col v-if="!isEdit" cols="12" md="6">
              <v-text-field
                v-model="localUser.password"
                label="كلمة المرور *"
                :rules="passwordRules"
                required
                variant="outlined"
                prepend-inner-icon="mdi-lock"
                :type="showPassword ? 'text' : 'password'"
                :append-inner-icon="showPassword ? 'mdi-eye' : 'mdi-eye-off'"
                @click:append-inner="showPassword = !showPassword"
              />
            </v-col>

            <!-- Notes -->
            <v-col cols="12">
              <v-textarea
                v-model="localUser.notes"
                label="ملاحظات"
                variant="outlined"
                prepend-inner-icon="mdi-note-text"
                rows="3"
                auto-grow
              />
            </v-col>
          </v-row>
        </v-form>
      </v-card-text>

      <v-divider />

      <v-card-actions class="dialog-actions">
        <v-spacer />
        <v-btn
          color="grey"
          variant="outlined"
          @click="closeDialog"
          class="me-2"
        >
          إلغاء
        </v-btn>
        <v-btn
          :color="isEdit ? 'success' : 'primary'"
          variant="elevated"
          @click="saveUser"
          :loading="loading"
          :disabled="!formValid"
        >
          {{ isEdit ? 'حفظ التعديلات' : 'حفظ المستخدم' }}
        </v-btn>
      </v-card-actions>
    </v-card>
  </v-dialog>
</template>

<script setup>
import { ref, watch, computed } from 'vue'

const props = defineProps({
  modelValue: {
    type: Boolean,
    default: false
  },
  user: {
    type: Object,
    default: () => ({})
  },
  loading: {
    type: Boolean,
    default: false
  }
})

const emit = defineEmits(['update:modelValue', 'save'])

const isEdit = computed(() => !!props.user?.id)

const formRef = ref(null)
const formValid = ref(false)
const showPassword = ref(false)

const defaultAvatar = 'https://via.placeholder.com/100x100?text=صورة'

const localUser = ref({
  firstName: '',
  lastName: '',
  name: '',
  email: '',
  phone: '',
  role: '',
  department: '',
  status: 'active',
  password: '',
  notes: '',
  avatar: ''
})

// Options
const roleOptions = [
  { title: 'مدير عام', value: 'admin' },
  { title: 'مدير مشروع', value: 'project_manager' },
  { title: 'مهندس', value: 'engineer' },
  { title: 'محاسب', value: 'accountant' },
  { title: 'موظف', value: 'employee' },
  { title: 'مراجع', value: 'reviewer' },
  { title: 'مستخدم', value: 'user' }
]

const departmentOptions = [
  { title: 'الإدارة', value: 'الإدارة' },
  { title: 'المحاسبة', value: 'المحاسبة' },
  { title: 'الهندسة', value: 'الهندسة' },
  { title: 'الموارد البشرية', value: 'الموارد البشرية' },
  { title: 'الصيانة', value: 'الصيانة' },
  { title: 'التسويق', value: 'التسويق' },
  { title: 'المبيعات', value: 'المبيعات' },
  { title: 'تكنولوجيا المعلومات', value: 'تكنولوجيا المعلومات' }
]

const statusOptions = [
  { title: 'نشط', value: 'active' },
  { title: 'غير نشط', value: 'inactive' },
  { title: 'معلق', value: 'pending' },
  { title: 'محظور', value: 'banned' }
]

// Validation Rules
const nameRules = [
  v => !!v || 'الاسم مطلوب',
  v => (v && v.length >= 2) || 'الاسم يجب أن يكون على الأقل حرفين',
  v => (v && v.length <= 50) || 'الاسم يجب أن يكون أقل من 50 حرف'
]

const emailRules = [
  v => !!v || 'البريد الإلكتروني مطلوب',
  v => /.+@.+\..+/.test(v) || 'البريد الإلكتروني غير صحيح'
]

const passwordRules = [
  v => !!v || 'كلمة المرور مطلوبة',
  v => (v && v.length >= 6) || 'كلمة المرور يجب أن تكون على الأقل 6 أحرف',
  v => (v && v.length <= 20) || 'كلمة المرور يجب أن تكون أقل من 20 حرف'
]

const requiredRules = [
  v => !!v || 'هذا الحقل مطلوب'
]

const resetForm = () => {
  localUser.value = {
    firstName: '',
    lastName: '',
    name: '',
    email: '',
    phone: '',
    role: '',
    department: '',
    status: 'active',
    password: '',
    notes: '',
    avatar: ''
  }
  if (formRef.value) {
    formRef.value.resetValidation()
  }
  formValid.value = false
  showPassword.value = false
}

// Watch for user prop changes
watch(() => props.user, (newUser) => {
  if (newUser && Object.keys(newUser).length > 0) {
    localUser.value = { ...newUser }
  } else {
    resetForm()
  }
}, { immediate: true, deep: true })

const closeDialog = () => {
  emit('update:modelValue', false)
  resetForm()
}

const saveUser = async () => {
  if (!formRef.value) return

  const { valid } = await formRef.value.validate()
  if (!valid) return

  // Build user data
  const userData = { ...localUser.value }

  // For new users, combine first and last name
  if (!isEdit.value) {
    userData.name = `${userData.firstName} ${userData.lastName}`
  }

  emit('save', userData)
}
</script>

<style scoped>
@import './styles/users.css';
</style>
