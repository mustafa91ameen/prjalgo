<template>
  <div class="users-container">
    <!-- Page Header Component -->
    <PageHeader
      title="إدارة المستخدمين"
      subtitle="إدارة وتتبع جميع المستخدمين والصلاحيات"
      badge="المستخدمون"
      badgeType="info"
      class="users-header"
    >
      <template #actions>
        <button class="page-action-btn secondary">
          <i class="mdi mdi-export"></i>
          تصدير
        </button>
        <button class="page-icon-btn">
          <i class="mdi mdi-dots-vertical"></i>
        </button>
      </template>
    </PageHeader>

    <!-- Statistics Cards -->
    <div class="stats-grid">
      <!-- Total Users Card -->
      <v-card class="stat-card" elevation="0">
        <div class="stat-card-content">
          <div class="stat-icon primary">
            <i class="mdi mdi-account-group"></i>
          </div>
          <div class="stat-info">
            <div class="stat-label">إجمالي المستخدمين</div>
            <div class="stat-value">{{ pagination.total }}</div>
          </div>
        </div>
      </v-card>

      <!-- Active Users Card -->
      <v-card class="stat-card" elevation="0">
        <div class="stat-card-content">
          <div class="stat-icon success">
            <i class="mdi mdi-account-check"></i>
          </div>
          <div class="stat-info">
            <div class="stat-label">مستخدمين نشطين</div>
            <div class="stat-value">{{ activeUsers }}</div>
          </div>
        </div>
      </v-card>

      <!-- Inactive Users Card -->
      <v-card class="stat-card" elevation="0">
        <div class="stat-card-content">
          <div class="stat-icon warning">
            <i class="mdi mdi-account-alert"></i>
          </div>
          <div class="stat-info">
            <div class="stat-label">مستخدمين غير نشطين</div>
            <div class="stat-value">{{ inactiveUsers }}</div>
          </div>
        </div>
      </v-card>
    </div>

    <!-- Users List Header -->
    <div class="users-list-header">
      <div class="list-header-content">
        <div class="list-header-info">
          <h2 class="list-header-title">
            <i class="mdi mdi-format-list-bulleted"></i>
            قائمة المستخدمين
          </h2>
          <p class="list-header-subtitle">عرض جميع المستخدمين والصلاحيات</p>
        </div>
        <div class="list-header-actions">
          <button v-if="canCreate" class="list-action-btn primary" @click="openAddDialog">
            <i class="mdi mdi-plus"></i>
            إضافة مستخدم جديد
          </button>
        </div>
      </div>
    </div>

    <!-- Users Table -->
    <v-card class="mb-6">
      <v-data-table
        :headers="headers"
        :items="users"
        :loading="loading"
        class="elevation-1"
      >
        <template v-slot:item.status="{ item }">
          <v-chip
            :color="getStatusColor(item.status)"
            size="small"
          >
            {{ getStatusText(item.status) }}
          </v-chip>
        </template>
        <template v-slot:item.actions="{ item }">
          <v-btn
            v-if="canUpdate && canEditUser(item)"
            size="small"
            color="primary"
            class="me-2"
            @click="editUser(item)"
          >
            <i class="mdi mdi-pencil"></i>
          </v-btn>
          <v-btn
            v-if="canUpdatePassword && canEditUser(item)"
            size="small"
            color="purple"
            class="me-2"
            @click="openPasswordDialog(item)"
          >
            <i class="mdi mdi-lock-reset"></i>
          </v-btn>
          <v-btn
            v-if="canUpdateStatus && canEditUser(item)"
            size="small"
            :color="item.status === 'active' ? 'warning' : 'success'"
            class="me-2"
            @click="toggleUserStatus(item)"
          >
            <i :class="item.status === 'active' ? 'mdi mdi-account-off' : 'mdi mdi-account-check'"></i>
          </v-btn>
          <v-btn
            v-if="canDelete && canEditUser(item)"
            size="small"
            color="error"
            @click="confirmDeleteUser(item)"
          >
            <i class="mdi mdi-delete"></i>
          </v-btn>
        </template>
      </v-data-table>

      <!-- Pagination -->
      <div class="d-flex justify-center pa-4" v-if="pagination.totalPages > 1">
        <v-pagination
          v-model="currentPage"
          :length="pagination.totalPages"
          :total-visible="5"
          @update:model-value="changePage"
        ></v-pagination>
      </div>
    </v-card>

    <!-- Add/Edit User Dialog -->
    <v-dialog v-model="showAddDialog" max-width="900" persistent>
      <v-card class="user-dialog">
        <v-card-title class="dialog-header">
          <i class="mdi mdi-account-plus"></i>
          {{ editingUser ? 'تعديل المستخدم' : 'إضافة مستخدم جديد' }}
        </v-card-title>

        <v-card-text class="dialog-content">
          <v-form ref="userForm" v-model="formValid">
            <v-row>
              <!-- اسم المستخدم -->
              <v-col cols="12" md="6">
                <v-text-field
                  v-model="newUser.username"
                  label="اسم المستخدم"
                  variant="outlined"
                  density="comfortable"
                  :rules="[v => !!v || 'اسم المستخدم مطلوب']"
                  :disabled="!!editingUser"
                  required
                ></v-text-field>
              </v-col>

              <!-- الاسم الكامل -->
              <v-col cols="12" md="6">
                <v-text-field
                  v-model="newUser.fullName"
                  label="الاسم الكامل"
                  variant="outlined"
                  density="comfortable"
                  :rules="[v => !!v || 'الاسم الكامل مطلوب']"
                  required
                ></v-text-field>
              </v-col>

              <!-- البريد الإلكتروني -->
              <v-col cols="12" md="6">
                <v-text-field
                  v-model="newUser.email"
                  label="البريد الإلكتروني"
                  variant="outlined"
                  density="comfortable"
                  type="email"
                  :rules="[
                    v => !!v || 'البريد الإلكتروني مطلوب',
                    v => /.+@.+\..+/.test(v) || 'البريد الإلكتروني غير صحيح'
                  ]"
                  required
                ></v-text-field>
              </v-col>

              <!-- رقم الهاتف -->
              <v-col cols="12" md="6">
                <v-text-field
                  v-model="newUser.phone"
                  label="رقم الهاتف"
                  variant="outlined"
                  density="comfortable"
                  :rules="[v => !!v || 'رقم الهاتف مطلوب']"
                  required
                ></v-text-field>
              </v-col>

              <!-- المسمى الوظيفي -->
              <v-col cols="12" md="6">
                <v-text-field
                  v-model="newUser.jobTitle"
                  label="المسمى الوظيفي"
                  variant="outlined"
                  density="comfortable"
                  :rules="[v => !!v || 'المسمى الوظيفي مطلوب']"
                  required
                ></v-text-field>
              </v-col>

              <!-- كلمة المرور (فقط للإضافة) -->
              <v-col cols="12" md="6" v-if="!editingUser">
                <v-text-field
                  v-model="newUser.password"
                  label="كلمة المرور"
                  variant="outlined"
                  density="comfortable"
                  type="password"
                  :rules="[
                    v => !!v || 'كلمة المرور مطلوبة',
                    v => v.length >= 8 || 'كلمة المرور يجب أن تكون 8 أحرف على الأقل'
                  ]"
                  required
                ></v-text-field>
              </v-col>
            </v-row>
          </v-form>
        </v-card-text>

        <v-card-actions class="dialog-actions">
          <v-spacer></v-spacer>
          <button class="dialog-btn cancel" @click="closeDialog">
            <i class="mdi mdi-close"></i>
            إلغاء
          </button>
          <button class="dialog-btn save" @click="saveUser" :disabled="!formValid || saving">
            <i class="mdi mdi-content-save"></i>
            {{ saving ? 'جاري الحفظ...' : 'حفظ' }}
          </button>
        </v-card-actions>
      </v-card>
    </v-dialog>

    <!-- Delete Confirmation Dialog -->
    <v-dialog v-model="showDeleteDialog" max-width="400">
      <v-card class="user-dialog">
        <v-card-title class="dialog-header">
          <i class="mdi mdi-alert"></i>
          تأكيد الحذف
        </v-card-title>
        <v-card-text class="dialog-content">
          هل أنت متأكد من حذف المستخدم "{{ userToDelete?.fullName }}"؟
        </v-card-text>
        <v-card-actions class="dialog-actions">
          <v-spacer></v-spacer>
          <button class="dialog-btn cancel" @click="showDeleteDialog = false">إلغاء</button>
          <button class="dialog-btn save" style="background: linear-gradient(135deg, #ef4444 0%, #dc2626 100%) !important;" @click="deleteUser">حذف</button>
        </v-card-actions>
      </v-card>
    </v-dialog>

    <!-- Password Change Dialog -->
    <v-dialog v-model="showPasswordDialog" max-width="500" persistent>
      <v-card class="user-dialog">
        <v-card-title class="dialog-header">
          <i class="mdi mdi-lock-reset"></i>
          تغيير كلمة المرور
        </v-card-title>
        <v-card-text class="dialog-content">
          <p class="mb-4" style="color: rgba(255,255,255,0.7);">
            تغيير كلمة المرور للمستخدم: <strong style="color: #06b6d4;">{{ userToChangePassword?.fullName }}</strong>
          </p>
          <v-form ref="passwordForm" v-model="passwordFormValid">
            <v-text-field
              v-model="newPassword"
              label="كلمة المرور الجديدة"
              variant="outlined"
              density="comfortable"
              type="password"
              :rules="[
                v => !!v || 'كلمة المرور مطلوبة',
                v => v.length >= 8 || 'كلمة المرور يجب أن تكون 8 أحرف على الأقل'
              ]"
              required
            ></v-text-field>
            <v-text-field
              v-model="confirmPassword"
              label="تأكيد كلمة المرور"
              variant="outlined"
              density="comfortable"
              type="password"
              :rules="[
                v => !!v || 'تأكيد كلمة المرور مطلوب',
                v => v === newPassword || 'كلمتا المرور غير متطابقتين'
              ]"
              required
            ></v-text-field>
          </v-form>
        </v-card-text>
        <v-card-actions class="dialog-actions">
          <v-spacer></v-spacer>
          <button class="dialog-btn cancel" @click="closePasswordDialog">
            <i class="mdi mdi-close"></i>
            إلغاء
          </button>
          <button class="dialog-btn save" style="background: linear-gradient(135deg, #9333ea 0%, #7c3aed 100%) !important;" @click="changePassword" :disabled="!passwordFormValid || savingPassword">
            <i class="mdi mdi-lock-check"></i>
            {{ savingPassword ? 'جاري الحفظ...' : 'تغيير' }}
          </button>
        </v-card-actions>
      </v-card>
    </v-dialog>

  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import PageHeader from '../components/PageHeader.vue'
import { listUsers, createUser, updateUser, updateUserStatus, updateUserPassword, deleteUser as apiDeleteUser } from '@/api/users'
import { usePermissions } from '@/composables/usePermissions'
import { useToast } from '@/composables/useToast'
import { DEFAULT_PAGE, DEFAULT_LIMIT } from '@/constants/pagination'

const { success, error: showError } = useToast()
const { canCreate, canUpdate, canDelete, canUpdateStatus, canUpdatePassword } = usePermissions('/users')
import { useAuthStore } from '@/stores/auth'

const authStore = useAuthStore()
const currentUserId = computed(() => authStore.user?.id)

// Check if user can be edited (not self, not super admin)
const canEditUser = (user) => {
  // Can't edit yourself
  if (user.id === currentUserId.value) return false
  return true
}

const loading = ref(false)
const saving = ref(false)
const savingPassword = ref(false)
const showAddDialog = ref(false)
const showDeleteDialog = ref(false)
const showPasswordDialog = ref(false)
const formValid = ref(false)
const passwordFormValid = ref(false)
const userForm = ref(null)
const passwordForm = ref(null)
const editingUser = ref(null)
const userToDelete = ref(null)
const userToChangePassword = ref(null)
const newPassword = ref('')
const confirmPassword = ref('')
const currentPage = ref(DEFAULT_PAGE)

// Pagination
const pagination = ref({
  total: 0,
  page: DEFAULT_PAGE,
  limit: DEFAULT_LIMIT,
  totalPages: 0
})

// New user data
const newUser = ref({
  username: '',
  fullName: '',
  email: '',
  phone: '',
  jobTitle: '',
  password: ''
})

const users = ref([])

const headers = [
  { title: 'الاسم', key: 'fullName', align: 'start' },
  { title: 'اسم المستخدم', key: 'username', align: 'center' },
  { title: 'البريد الإلكتروني', key: 'email', align: 'center' },
  { title: 'الهاتف', key: 'phone', align: 'center' },
  { title: 'المسمى الوظيفي', key: 'jobTitle', align: 'center' },
  { title: 'الحالة', key: 'status', align: 'center' },
  { title: 'الإجراءات', key: 'actions', align: 'center', sortable: false }
]

// Computed properties
const activeUsers = computed(() => {
  return users.value.filter(user => user.status === 'active').length
})

const inactiveUsers = computed(() => {
  return users.value.filter(user => user.status === 'inactive').length
})

// Methods
const fetchUsers = async () => {
  loading.value = true
  try {
    const response = await listUsers({ page: currentPage.value, limit: pagination.value.limit })
    users.value = response.data || []
    pagination.value = {
      total: response.total || 0,
      page: response.page || currentPage.value,
      limit: response.limit || DEFAULT_LIMIT,
      totalPages: response.totalPages || 0
    }
  } catch (error) {
    console.error('Error fetching users:', error)
    showError('حدث خطأ أثناء جلب المستخدمين')
  } finally {
    loading.value = false
  }
}

const changePage = (page) => {
  currentPage.value = page
  fetchUsers()
}

const openAddDialog = () => {
  editingUser.value = null
  resetForm()
  showAddDialog.value = true
}

const editUser = (user) => {
  editingUser.value = user
  newUser.value = {
    username: user.username,
    fullName: user.fullName,
    email: user.email,
    phone: user.phone,
    jobTitle: user.jobTitle,
    password: ''
  }
  showAddDialog.value = true
}

const closeDialog = () => {
  showAddDialog.value = false
  editingUser.value = null
  resetForm()
}

const resetForm = () => {
  newUser.value = {
    username: '',
    fullName: '',
    email: '',
    phone: '',
    jobTitle: '',
    password: ''
  }
  if (userForm.value) {
    userForm.value.reset()
  }
}

const saveUser = async () => {
  if (!formValid.value) return

  saving.value = true
  try {
    if (editingUser.value) {
      // Update existing user
      await updateUser(editingUser.value.id, {
        fullName: newUser.value.fullName,
        email: newUser.value.email,
        phone: newUser.value.phone,
        jobTitle: newUser.value.jobTitle
      })
      success('تم تحديث المستخدم بنجاح')
    } else {
      // Create new user
      await createUser(newUser.value)
      success('تم إضافة المستخدم بنجاح')
    }
    closeDialog()
    fetchUsers()
  } catch (error) {
    console.error('Error saving user:', error)
    showError(error.message || 'حدث خطأ أثناء حفظ المستخدم')
  } finally {
    saving.value = false
  }
}

const toggleUserStatus = async (user) => {
  try {
    const newStatus = user.status === 'active' ? 'inactive' : 'active'
    await updateUserStatus(user.id, newStatus)
    success('تم تحديث حالة المستخدم بنجاح')
    fetchUsers()
  } catch (error) {
    console.error('Error updating user status:', error)
    showError(error.message || 'حدث خطأ أثناء تحديث الحالة')
  }
}

const confirmDeleteUser = (user) => {
  userToDelete.value = user
  showDeleteDialog.value = true
}

const deleteUser = async () => {
  if (!userToDelete.value) return

  try {
    await apiDeleteUser(userToDelete.value.id)
    success('تم حذف المستخدم بنجاح')
    showDeleteDialog.value = false
    userToDelete.value = null
    fetchUsers()
  } catch (error) {
    console.error('Error deleting user:', error)
    showError(error.message || 'حدث خطأ أثناء حذف المستخدم')
  }
}

// Password change functions
const openPasswordDialog = (user) => {
  userToChangePassword.value = user
  newPassword.value = ''
  confirmPassword.value = ''
  showPasswordDialog.value = true
}

const closePasswordDialog = () => {
  showPasswordDialog.value = false
  userToChangePassword.value = null
  newPassword.value = ''
  confirmPassword.value = ''
  if (passwordForm.value) {
    passwordForm.value.reset()
  }
}

const changePassword = async () => {
  if (!passwordFormValid.value || !userToChangePassword.value) return

  savingPassword.value = true
  try {
    await updateUserPassword(userToChangePassword.value.id, newPassword.value)
    success('تم تغيير كلمة المرور بنجاح')
    closePasswordDialog()
  } catch (error) {
    console.error('Error changing password:', error)
    showError(error.message || 'حدث خطأ أثناء تغيير كلمة المرور')
  } finally {
    savingPassword.value = false
  }
}

const getStatusColor = (status) => {
  const colors = {
    active: 'success',
    inactive: 'error'
  }
  return colors[status] || 'grey'
}

const getStatusText = (status) => {
  const texts = {
    active: 'نشط',
    inactive: 'غير نشط'
  }
  return texts[status] || status
}

onMounted(() => {
  fetchUsers()
})
</script>

<style scoped>
.users-container {
  padding: 32px;
  max-width: 1600px;
  margin: 0 auto;
}

/* Users Header Custom Color */
.users-header {
  background: linear-gradient(135deg, #018790 0%, #005461 100%) !important;
}

.users-header::before {
  background: linear-gradient(135deg, #06b6d4 0%, #10b981 50%, #06b6d4 100%) !important;
}

/* Statistics Grid */
.stats-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
  gap: 16px;
  margin-bottom: 24px;
}

.stat-card {
  border-radius: 16px !important;
  overflow: hidden;
  transition: all 0.4s cubic-bezier(0.4, 0, 0.2, 1);
  background: linear-gradient(135deg, #0a3d42 0%, #052428 100%) !important;
  border: 2px solid transparent !important;
  position: relative;
}

.stat-card::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  border-radius: 16px;
  padding: 2px;
  background: linear-gradient(135deg, #06b6d4 0%, #10b981 50%, #14b8a6 100%);
  -webkit-mask: linear-gradient(#fff 0 0) content-box, linear-gradient(#fff 0 0);
  mask: linear-gradient(#fff 0 0) content-box, linear-gradient(#fff 0 0);
  -webkit-mask-composite: xor;
  mask-composite: exclude;
  pointer-events: none;
}

.stat-card:hover {
  transform: translateY(-4px) scale(1.01);
  box-shadow: 0 12px 24px rgba(6, 182, 212, 0.3),
              0 0 40px rgba(16, 185, 129, 0.2) !important;
}

.stat-card-content {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  padding: 20px 16px;
  text-align: center;
  position: relative;
  z-index: 1;
}

.stat-info {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 4px;
}

.stat-label {
  font-size: 13px;
  color: rgba(255, 255, 255, 0.7);
  margin-bottom: 2px;
  font-weight: 500;
  letter-spacing: 0.3px;
}

.stat-value {
  font-size: 28px;
  font-weight: 800;
  background: linear-gradient(135deg, #ffffff 0%, #e0e7ff 100%);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background-clip: text;
  margin-bottom: 4px;
}

.stat-icon {
  width: 48px;
  height: 48px;
  border-radius: 12px;
  display: flex;
  align-items: center;
  justify-content: center;
  color: white;
  font-size: 24px;
  margin-bottom: 12px;
  background: linear-gradient(135deg, #06b6d4 0%, #0891b2 100%) !important;
  box-shadow: 0 6px 16px rgba(6, 182, 212, 0.4);
}

.stat-icon.primary {
  background: linear-gradient(135deg, #3b82f6 0%, #2563eb 100%) !important;
  box-shadow: 0 6px 16px rgba(59, 130, 246, 0.4);
}

.stat-icon.success {
  background: linear-gradient(135deg, #10b981 0%, #059669 100%) !important;
  box-shadow: 0 6px 16px rgba(16, 185, 129, 0.4);
}

.stat-icon.warning {
  background: linear-gradient(135deg, #f59e0b 0%, #d97706 100%) !important;
  box-shadow: 0 6px 16px rgba(245, 158, 11, 0.4);
}

/* Users List Header */
.users-list-header {
  background: linear-gradient(135deg, #0a3d42 0%, #052428 100%);
  border-radius: 16px;
  margin-bottom: 24px;
  padding: 16px 24px;
  border: 2px solid transparent;
  position: relative;
  overflow: hidden;
}

.users-list-header::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  border-radius: 16px;
  padding: 2px;
  background: linear-gradient(135deg, #3b82f6 0%, #2563eb 50%, #3b82f6 100%);
  -webkit-mask: linear-gradient(#fff 0 0) content-box, linear-gradient(#fff 0 0);
  mask: linear-gradient(#fff 0 0) content-box, linear-gradient(#fff 0 0);
  -webkit-mask-composite: xor;
  mask-composite: exclude;
  pointer-events: none;
}

.list-header-content {
  display: flex;
  justify-content: space-between;
  align-items: center;
  position: relative;
  z-index: 1;
}

.list-header-info {
  flex: 1;
}

.list-header-title {
  font-size: 18px;
  font-weight: 700;
  color: rgba(255, 255, 255, 0.95);
  margin: 0 0 2px 0;
  display: flex;
  align-items: center;
  gap: 8px;
}

.list-header-title i {
  color: #3b82f6;
  font-size: 20px;
}

.list-header-subtitle {
  font-size: 12px;
  color: rgba(255, 255, 255, 0.6);
  margin: 0;
}

.list-header-actions {
  display: flex;
  gap: 12px;
}

.list-action-btn {
  padding: 8px 18px;
  border-radius: 12px;
  border: none;
  font-size: 13px;
  font-weight: 700;
  cursor: pointer;
  display: inline-flex;
  align-items: center;
  gap: 8px;
  transition: all 0.3s ease;
}

.list-action-btn.primary {
  background: linear-gradient(135deg, #3b82f6 0%, #2563eb 100%);
  color: white;
  box-shadow: 0 4px 12px rgba(59, 130, 246, 0.3);
}

.list-action-btn.primary:hover {
  transform: translateY(-2px);
  box-shadow: 0 6px 16px rgba(59, 130, 246, 0.4);
}

.list-action-btn i {
  font-size: 16px;
}

/* User Dialog */
.user-dialog {
  border-radius: 16px !important;
  background: linear-gradient(135deg, #0a3d42 0%, #052428 100%) !important;
  border: 2px solid transparent !important;
  position: relative;
  overflow: hidden;
  direction: rtl;
  text-align: right;
}

.user-dialog::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  border-radius: 16px;
  padding: 2px;
  background: linear-gradient(135deg, #3b82f6 0%, #2563eb 50%, #3b82f6 100%);
  -webkit-mask: linear-gradient(#fff 0 0) content-box, linear-gradient(#fff 0 0);
  mask: linear-gradient(#fff 0 0) content-box, linear-gradient(#fff 0 0);
  -webkit-mask-composite: xor;
  mask-composite: exclude;
  pointer-events: none;
}

.dialog-header {
  background: rgba(59, 130, 246, 0.1) !important;
  color: rgba(255, 255, 255, 0.95) !important;
  font-size: 20px !important;
  font-weight: 700 !important;
  padding: 20px 24px !important;
  display: flex;
  align-items: center;
  gap: 10px;
  position: relative;
  z-index: 1;
  direction: rtl;
  text-align: right;
}

.dialog-header i {
  color: #3b82f6;
  font-size: 24px;
}

.dialog-content {
  padding: 24px !important;
  position: relative;
  z-index: 1;
  direction: rtl;
  text-align: right;
  max-height: 60vh;
  overflow-y: auto;
}

.dialog-content::-webkit-scrollbar {
  display: none;
}

.dialog-content {
  -ms-overflow-style: none;
  scrollbar-width: none;
}

.dialog-content :deep(.v-field) {
  background: rgba(255, 255, 255, 0.08) !important;
  border-radius: 12px;
  direction: rtl;
  text-align: right;
}

.dialog-content :deep(.v-field__outline) {
  color: rgba(255, 255, 255, 0.2) !important;
}

.dialog-content :deep(.v-field--focused .v-field__outline) {
  color: #3b82f6 !important;
}

.dialog-content :deep(.v-label) {
  color: rgba(255, 255, 255, 0.7) !important;
  right: 12px !important;
  left: auto !important;
}

.dialog-content :deep(.v-field__input) {
  color: rgba(255, 255, 255, 0.95) !important;
  text-align: right;
  direction: rtl;
}

.dialog-content :deep(textarea) {
  color: rgba(255, 255, 255, 0.95) !important;
  text-align: right;
  direction: rtl;
}

.dialog-content :deep(.v-input__details) {
  direction: rtl;
  text-align: right;
}

.dialog-actions {
  padding: 16px 24px !important;
  border-top: 1px solid rgba(255, 255, 255, 0.1);
  position: relative;
  z-index: 1;
  direction: rtl;
  text-align: right;
}

.dialog-btn {
  padding: 10px 24px;
  border-radius: 12px;
  border: none;
  font-size: 14px;
  font-weight: 700;
  cursor: pointer;
  display: inline-flex;
  align-items: center;
  gap: 8px;
  transition: all 0.3s ease;
}

.dialog-btn.cancel {
  background: rgba(255, 255, 255, 0.1);
  color: rgba(255, 255, 255, 0.8);
  border: 1px solid rgba(255, 255, 255, 0.2);
}

.dialog-btn.cancel:hover {
  background: rgba(255, 255, 255, 0.15);
  transform: translateY(-2px);
}

.dialog-btn.save {
  background: linear-gradient(135deg, #3b82f6 0%, #2563eb 100%);
  color: white;
  box-shadow: 0 4px 12px rgba(59, 130, 246, 0.3);
}

.dialog-btn.save:hover:not(:disabled) {
  transform: translateY(-2px);
  box-shadow: 0 6px 16px rgba(59, 130, 246, 0.4);
}

.dialog-btn.save:disabled {
  opacity: 0.5;
  cursor: not-allowed;
}

.dialog-btn i {
  font-size: 18px;
}

.v-card {
  border-radius: 12px;
}

/* Responsive Styles */
@media (max-width: 1024px) {
  .users-container {
    padding: 24px;
  }

  .stats-grid {
    grid-template-columns: repeat(2, 1fr);
  }

  .list-header-content {
    flex-direction: column;
    align-items: flex-start;
    gap: 16px;
  }

  .list-header-actions {
    width: 100%;
  }

  .list-action-btn.primary {
    width: 100%;
    justify-content: center;
  }
}

@media (max-width: 768px) {
  .users-container {
    padding: 16px;
  }

  .stats-grid {
    grid-template-columns: 1fr;
    gap: 12px;
  }

  .stat-card-content {
    flex-direction: row;
    justify-content: flex-start;
    gap: 16px;
    padding: 16px;
    text-align: right;
  }

  .stat-icon {
    margin-bottom: 0;
  }

  .stat-info {
    align-items: flex-start;
  }

  .users-list-header {
    padding: 12px 16px;
  }

  .list-header-title {
    font-size: 16px;
  }

  .list-action-btn {
    padding: 10px 16px;
    font-size: 12px;
  }

  /* Dialog responsive */
  .user-dialog .dialog-header {
    padding: 16px !important;
    font-size: 18px !important;
  }

  .user-dialog .dialog-content {
    padding: 16px !important;
  }

  .user-dialog .dialog-actions {
    padding: 12px 16px !important;
    flex-wrap: wrap;
    gap: 8px;
  }

  .dialog-btn {
    padding: 8px 16px;
    font-size: 13px;
  }

  /* Table responsive */
  .v-data-table {
    font-size: 13px;
  }

  .v-data-table :deep(th),
  .v-data-table :deep(td) {
    padding: 8px 12px !important;
  }
}

@media (max-width: 480px) {
  .users-container {
    padding: 12px;
  }

  .stat-value {
    font-size: 24px;
  }

  .stat-label {
    font-size: 12px;
  }

  .stat-icon {
    width: 40px;
    height: 40px;
    font-size: 20px;
  }

  .users-list-header {
    padding: 12px;
    border-radius: 12px;
  }

  .list-header-title {
    font-size: 14px;
  }

  .list-header-subtitle {
    font-size: 11px;
  }

  .list-action-btn {
    padding: 8px 12px;
    font-size: 11px;
    border-radius: 10px;
  }

  /* Dialog responsive - full width on mobile */
  :deep(.v-dialog) {
    margin: 8px !important;
  }

  .user-dialog {
    border-radius: 12px !important;
  }

  .user-dialog .dialog-header {
    padding: 12px 16px !important;
    font-size: 16px !important;
    gap: 8px;
  }

  .user-dialog .dialog-header i {
    font-size: 20px;
  }

  .user-dialog .dialog-content {
    padding: 12px !important;
    max-height: 50vh;
  }

  .dialog-btn {
    padding: 8px 12px;
    font-size: 12px;
    flex: 1;
  }

  /* Table overflow handling */
  .v-card.mb-6 {
    overflow-x: auto;
  }

  .v-data-table :deep(.v-table__wrapper) {
    overflow-x: auto;
  }

  .v-data-table :deep(th),
  .v-data-table :deep(td) {
    padding: 6px 8px !important;
    font-size: 12px;
    white-space: nowrap;
  }

  .v-btn.me-2 {
    margin-inline-end: 4px !important;
  }

  /* Pagination responsive */
  .v-pagination :deep(.v-btn) {
    min-width: 32px !important;
    height: 32px !important;
  }
}
</style>
