<template>
  <div class="roles-container">
    <!-- Page Header Component -->
    <PageHeader
      title="إدارة الأدوار والصلاحيات"
      subtitle="إدارة وتنظيم جميع الأدوار وصلاحيات النظام"
      badge="الأدوار"
      badgeType="primary"
      class="roles-header"
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
      <!-- Total Roles Card -->
      <v-card class="stat-card" elevation="0">
        <div class="stat-card-content">
          <div class="stat-icon role">
            <i class="mdi mdi-shield-account"></i>
          </div>
          <div class="stat-info">
            <div class="stat-label">إجمالي الأدوار</div>
            <div class="stat-value">{{ totalRoles }}</div>
          </div>
        </div>
      </v-card>

      <!-- Total Pages Card -->
      <v-card class="stat-card" elevation="0">
        <div class="stat-card-content">
          <div class="stat-icon page">
            <i class="mdi mdi-file-document-multiple"></i>
          </div>
          <div class="stat-info">
            <div class="stat-label">إجمالي الصفحات</div>
            <div class="stat-value">{{ totalPagesCount }}</div>
          </div>
        </div>
      </v-card>
    </div>

    <!-- Roles List Header -->
    <div class="roles-list-header">
      <div class="list-header-content">
        <div class="list-header-info">
          <h2 class="list-header-title">
            <i class="mdi mdi-format-list-bulleted"></i>
            قائمة الأدوار
          </h2>
          <p class="list-header-subtitle">عرض جميع الأدوار والصلاحيات</p>
        </div>
        <div class="list-header-actions">
          <button v-if="canCreate" class="list-action-btn primary" @click="openAddDialog">
            <i class="mdi mdi-plus"></i>
            إضافة دور جديد
          </button>
        </div>
      </div>
    </div>

    <!-- Roles Table -->
    <v-card class="mb-6">
      <v-data-table
        :headers="headers"
        :items="roles"
        :loading="loading"
        class="elevation-1"
      >
        <template v-slot:item.name="{ item }">
          <div class="d-flex align-center">
            <v-avatar size="36" color="primary" class="me-3">
              <i class="mdi mdi-shield" style="color: white; font-size: 18px;"></i>
            </v-avatar>
            <div>
              <div class="font-weight-medium">{{ item.name }}</div>
              <div class="text-caption text-grey">{{ item.description || 'لا يوجد وصف' }}</div>
            </div>
          </div>
        </template>
        <template v-slot:item.pagesCount="{ item }">
          <v-chip color="info" size="small" variant="flat">
            {{ item.pagesCount }} صفحة
          </v-chip>
        </template>
        <template v-slot:item.permissions="{ item }">
          <v-chip
            v-for="(permission, index) in item.permissions.slice(0, 3)"
            :key="index"
            size="small"
            class="me-1"
            color="primary"
          >
            {{ permission }}
          </v-chip>
          <v-chip
            v-if="item.permissions.length > 3"
            size="small"
            color="grey"
          >
            +{{ item.permissions.length - 3 }}
          </v-chip>
        </template>
        <template v-slot:item.actions="{ item }">
          <v-btn
            size="small"
            color="info"
            variant="tonal"
            icon
            class="me-1"
            @click="viewRole(item)"
            title="عرض التفاصيل"
          >
            <i class="mdi mdi-eye"></i>
          </v-btn>
          <v-btn
            v-if="canUpdate"
            size="small"
            color="primary"
            variant="tonal"
            icon
            class="me-1"
            @click="editRole(item)"
            title="تعديل"
          >
            <i class="mdi mdi-pencil"></i>
          </v-btn>
          <v-btn
            v-if="canUpdate"
            size="small"
            color="warning"
            variant="tonal"
            icon
            class="me-1"
            @click="openPermissionsDialog(item)"
            title="إدارة الصلاحيات"
          >
            <i class="mdi mdi-key"></i>
          </v-btn>
          <v-btn
            v-if="canDelete"
            size="small"
            color="error"
            variant="tonal"
            icon
            @click="confirmDeleteRole(item)"
            title="حذف"
          >
            <i class="mdi mdi-delete"></i>
          </v-btn>
        </template>
      </v-data-table>
    </v-card>

    <!-- Add/Edit Role Dialog -->
    <v-dialog v-model="showAddDialog" max-width="700" persistent>
      <v-card class="role-dialog">
        <v-card-title class="dialog-header">
          <i :class="editingRole ? 'mdi mdi-pencil' : 'mdi mdi-shield-plus'"></i>
          {{ editingRole ? 'تعديل الدور' : 'إضافة دور جديد' }}
        </v-card-title>

        <v-card-text class="dialog-content">
          <v-form ref="roleForm" v-model="formValid">
            <v-row>
              <!-- اسم الدور -->
              <v-col cols="12">
                <v-text-field
                  v-model="newRole.name"
                  label="اسم الدور"
                  variant="outlined"
                  density="comfortable"
                  :rules="[v => !!v || 'اسم الدور مطلوب']"
                  required
                  autofocus
                ></v-text-field>
              </v-col>

              <!-- وصف الدور -->
              <v-col cols="12">
                <v-textarea
                  v-model="newRole.description"
                  label="وصف الدور"
                  variant="outlined"
                  density="comfortable"
                  rows="3"
                ></v-textarea>
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
          <button class="dialog-btn save" @click="saveRole" :disabled="!formValid || saving">
            <i class="mdi mdi-content-save"></i>
            {{ saving ? 'جاري الحفظ...' : 'حفظ' }}
          </button>
        </v-card-actions>
      </v-card>
    </v-dialog>

    <!-- View Role Dialog -->
    <v-dialog v-model="showViewDialog" max-width="700">
      <v-card class="role-dialog">
        <v-card-title class="dialog-header">
          <i class="mdi mdi-shield-account"></i>
          تفاصيل الدور
        </v-card-title>

        <v-card-text v-if="selectedRole" class="dialog-content">
          <div class="text-center mb-4">
            <v-avatar size="80" color="primary">
              <i class="mdi mdi-shield" style="color: white; font-size: 40px;"></i>
            </v-avatar>
            <h3 class="mt-3" style="color: rgba(255,255,255,0.95);">{{ selectedRole.name }}</h3>
            <p class="text-caption" style="color: rgba(255,255,255,0.7);">{{ selectedRole.description || 'لا يوجد وصف' }}</p>
          </div>

          <h4 class="mb-3" style="color: rgba(255,255,255,0.9);">الصفحات والصلاحيات:</h4>
          <v-list density="compact" bg-color="transparent">
            <v-list-item
              v-for="rp in viewRolePages"
              :key="rp.id"
              class="mb-2 role-page-item"
            >
              <template v-slot:prepend>
                <i :class="rp.page?.icon || 'mdi mdi-file'" style="color: #8b5cf6; font-size: 24px; margin-left: 12px;"></i>
              </template>
              <v-list-item-title style="color: rgba(255,255,255,0.95);">{{ rp.page?.name || 'صفحة غير معروفة' }}</v-list-item-title>
              <v-list-item-subtitle>
                <v-chip
                  v-for="perm in rp.permissions"
                  :key="perm"
                  size="x-small"
                  :color="getPermissionColor(perm)"
                  class="me-1"
                >
                  {{ getPermissionText(perm) }}
                </v-chip>
              </v-list-item-subtitle>
            </v-list-item>
            <v-list-item v-if="viewRolePages.length === 0">
              <v-list-item-title style="color: rgba(255,255,255,0.5);">لا توجد صفحات مخصصة لهذا الدور</v-list-item-title>
            </v-list-item>
          </v-list>
        </v-card-text>

        <v-card-actions class="dialog-actions">
          <v-spacer></v-spacer>
          <button class="dialog-btn save" @click="showViewDialog = false">
            <i class="mdi mdi-close"></i>
            إغلاق
          </button>
        </v-card-actions>
      </v-card>
    </v-dialog>

    <!-- Permissions Dialog -->
    <v-dialog v-model="showPermissionsDialog" max-width="900" persistent>
      <v-card class="role-dialog">
        <v-card-title class="dialog-header">
          <i class="mdi mdi-key"></i>
          إدارة صلاحيات الدور - {{ selectedRole?.name }}
        </v-card-title>

        <v-card-text class="dialog-content" style="max-height: 70vh; overflow-y: auto;">
          <v-alert type="info" variant="tonal" class="mb-4">
            حدد الصفحات التي يمكن لهذا الدور الوصول إليها، واختر الصلاحيات المناسبة لكل صفحة
          </v-alert>

          <v-expansion-panels variant="accordion" class="permissions-panels">
            <v-expansion-panel
              v-for="page in allPages"
              :key="page.id"
              class="mb-2"
            >
              <v-expansion-panel-title class="page-panel-title">
                <div class="d-flex align-center justify-space-between w-100">
                  <div class="d-flex align-center">
                    <v-checkbox
                      v-model="pageEnabled[page.id]"
                      hide-details
                      density="compact"
                      color="primary"
                      @click.stop
                      @change="togglePagePermissions(page.id)"
                    />
                    <i :class="page.icon || 'mdi mdi-file'" style="color: #8b5cf6; font-size: 24px; margin: 0 8px;"></i>
                    <div>
                      <span class="font-weight-bold" style="color: rgba(255,255,255,0.95);">{{ page.name }}</span>
                      <span class="text-caption ms-2" style="color: rgba(255,255,255,0.5);">({{ page.route }})</span>
                    </div>
                  </div>
                  <v-chip
                    v-if="getSelectedPermissionsCount(page.id) > 0"
                    color="primary"
                    size="small"
                    class="me-4"
                  >
                    {{ getSelectedPermissionsCount(page.id) }} صلاحية
                  </v-chip>
                </div>
              </v-expansion-panel-title>
              <v-expansion-panel-text class="page-panel-content">
                <div class="permissions-grid">
                  <div
                    v-for="perm in getPermissionsForPage(page)"
                    :key="perm.value"
                    class="permission-checkbox"
                  >
                    <v-checkbox-btn
                      v-model="pagePermissions[page.id]"
                      :value="perm.value"
                      :color="perm.color"
                    />
                    <span class="permission-label">{{ perm.label }}</span>
                  </div>
                </div>
              </v-expansion-panel-text>
            </v-expansion-panel>
          </v-expansion-panels>
        </v-card-text>

        <v-card-actions class="dialog-actions">
          <v-spacer></v-spacer>
          <button class="dialog-btn cancel" @click="closePermissionsDialog">
            <i class="mdi mdi-close"></i>
            إلغاء
          </button>
          <button class="dialog-btn save" @click="savePermissions" :disabled="savingPermissions">
            <i class="mdi mdi-content-save"></i>
            {{ savingPermissions ? 'جاري الحفظ...' : 'حفظ الصلاحيات' }}
          </button>
        </v-card-actions>
      </v-card>
    </v-dialog>

    <!-- Delete Confirmation Dialog -->
    <v-dialog v-model="showDeleteDialog" max-width="500">
      <v-card class="role-dialog">
        <v-card-title class="dialog-header" style="background: linear-gradient(135deg, #dc2626 0%, #b91c1c 100%) !important;">
          <i class="mdi mdi-delete-alert"></i>
          تأكيد الحذف
        </v-card-title>

        <v-card-text v-if="selectedRole" class="dialog-content text-center">
          <v-avatar size="60" color="error" class="mb-3">
            <i class="mdi mdi-shield" style="color: white; font-size: 30px;"></i>
          </v-avatar>
          <h4 style="color: rgba(255,255,255,0.95);">{{ selectedRole.name }}</h4>

          <v-alert type="error" variant="tonal" class="my-4">
            تحذير: هذا الإجراء لا يمكن التراجع عنه!
          </v-alert>

          <p style="color: rgba(255,255,255,0.7);">
            هل أنت متأكد من حذف هذا الدور نهائياً؟
          </p>
        </v-card-text>

        <v-card-actions class="dialog-actions">
          <v-spacer></v-spacer>
          <button class="dialog-btn cancel" @click="showDeleteDialog = false">
            إلغاء
          </button>
          <button class="dialog-btn delete" @click="deleteRoleItem" :disabled="deleting">
            <i class="mdi mdi-delete"></i>
            {{ deleting ? 'جاري الحذف...' : 'حذف نهائي' }}
          </button>
        </v-card-actions>
      </v-card>
    </v-dialog>

  </div>
</template>

<script setup>
import { ref, reactive, computed, onMounted } from 'vue'
import PageHeader from '../components/PageHeader.vue'
import { listRoles, createRole, updateRole, deleteRole as apiDeleteRole } from '@/api/roles'
import { listPages } from '@/api/pages'
import { getRolePagesByRoleId, bulkUpdateRolePages, listRolePages } from '@/api/rolePages'
import { usePermissions } from '@/composables/usePermissions'
import { useToast } from '@/composables/useToast'
import { DEFAULT_PAGE, DEFAULT_LIMIT } from '@/constants/pagination'

const { success, error: showError } = useToast()

const { canCreate, canUpdate, canDelete } = usePermissions('/roles')

const loading = ref(false)
const saving = ref(false)
const deleting = ref(false)
const savingPermissions = ref(false)

// Dialog states
const showAddDialog = ref(false)
const showViewDialog = ref(false)
const showPermissionsDialog = ref(false)
const showDeleteDialog = ref(false)

const formValid = ref(false)
const roleForm = ref(null)
const editingRole = ref(null)
const selectedRole = ref(null)

// Pagination
const page = ref(DEFAULT_PAGE)
const limit = ref(DEFAULT_LIMIT)
const total = ref(0)

// Data
const roles = ref([])
const allPages = ref([])
const allRolePages = ref([])
const viewRolePages = ref([])

// New role data
const newRole = ref({
  name: '',
  description: ''
})

// Permissions matrix - pageId -> array of permissions
const pagePermissions = reactive({})

// Track which pages are enabled (have at least one permission)
const pageEnabled = reactive({})

// Base permissions for all pages
const basePermissions = [
  { value: 'read', label: 'قراءة', color: 'success' },
  { value: 'create', label: 'إنشاء', color: 'primary' },
  { value: 'update', label: 'تعديل', color: 'warning' },
  { value: 'delete', label: 'حذف', color: 'error' }
]

// Extra permissions for specific pages
const extraPermissions = {
  '/users': [
    { value: 'updatePassword', label: 'تغيير كلمة المرور', color: 'purple' }
  ]
}

// Get available permissions for a specific page
const getPermissionsForPage = (page) => {
  const extra = extraPermissions[page.route] || []
  return [...basePermissions, ...extra]
}

// Table headers
const headers = [
  { title: 'الدور', key: 'name', align: 'start' },
  { title: 'عدد الصفحات', key: 'pagesCount', align: 'center' },
  { title: 'الصلاحيات', key: 'permissions', align: 'center' },
  { title: 'الإجراءات', key: 'actions', align: 'center', sortable: false }
]

// Computed properties
const totalRoles = computed(() => total.value || roles.value.length)
const totalPagesCount = computed(() => allPages.value.length)

// Permission helpers
const getPermissionColor = (perm) => {
  const colors = {
    read: 'success',
    create: 'primary',
    update: 'warning',
    delete: 'error',
    updatePassword: 'purple'
  }
  return colors[perm] || 'grey'
}

const getPermissionText = (perm) => {
  const texts = {
    read: 'قراءة',
    create: 'إنشاء',
    update: 'تعديل',
    delete: 'حذف',
    updatePassword: 'تغيير كلمة المرور'
  }
  return texts[perm] || perm
}

// Toggle all permissions for a page
const togglePagePermissions = (pageId) => {
  const page = allPages.value.find(p => p.id === pageId)
  if (pageEnabled[pageId]) {
    // Enable - select ALL permissions for this page
    const allPerms = getPermissionsForPage(page || { route: '' })
    pagePermissions[pageId] = allPerms.map(p => p.value)
  } else {
    // Disable - clear all permissions
    pagePermissions[pageId] = []
  }
}

// Get count of selected permissions for a page
const getSelectedPermissionsCount = (pageId) => {
  return pagePermissions[pageId]?.length || 0
}

// Get pages count for a role
const getRolePagesCount = (roleId) => {
  return allRolePages.value.filter(rp => rp.roleId === roleId).length
}

// Fetch all data
const fetchData = async () => {
  loading.value = true
  try {
    const [rolesResponse, pagesResponse, rolePagesResponse] = await Promise.all([
      listRoles({ page: page.value, limit: limit.value }),
      listPages({ limit: 100 }),
      listRolePages({ limit: 500 })
    ])

    if (rolesResponse.success) {
      const rolesData = rolesResponse.data.data || []
      roles.value = rolesData.map(r => ({
        id: r.id,
        name: r.name,
        description: r.description || '',
        pagesCount: 0,
        permissions: []
      }))
      total.value = rolesResponse.data.total || 0
    }

    if (pagesResponse.success) {
      // listPages returns items, not data
      allPages.value = pagesResponse.data.items || pagesResponse.data.data || []
    }

    if (rolePagesResponse.success) {
      // Handle paginated response - data.data contains the array
      const rpData = rolePagesResponse.data
      allRolePages.value = Array.isArray(rpData) ? rpData : (rpData?.data || [])

      // Update roles with pages count and permissions
      roles.value.forEach(role => {
        const rolePagesList = allRolePages.value.filter(rp => rp.roleId === role.id)
        role.pagesCount = rolePagesList.length

        // Get page names as permissions display
        role.permissions = rolePagesList.map(rp => {
          const page = allPages.value.find(p => p.id === rp.pageId)
          return page ? page.name : ''
        }).filter(name => name)
      })
    }
  } catch (error) {
    console.error('Error fetching data:', error)
    showError('حدث خطأ في جلب البيانات')
  } finally {
    loading.value = false
  }
}

// Dialog functions
const openAddDialog = () => {
  editingRole.value = null
  newRole.value = { name: '', description: '' }
  showAddDialog.value = true
}

const editRole = (role) => {
  editingRole.value = role
  newRole.value = {
    name: role.name,
    description: role.description
  }
  showAddDialog.value = true
}

const closeDialog = () => {
  showAddDialog.value = false
  editingRole.value = null
  newRole.value = { name: '', description: '' }
}

const viewRole = async (role) => {
  selectedRole.value = role

  // Fetch role pages for this role
  try {
    viewRolePages.value = await getRolePagesByRoleId(role.id)
    // Add page info to each role page
    viewRolePages.value = viewRolePages.value.map(rp => ({
      ...rp,
      page: allPages.value.find(p => p.id === rp.pageId)
    }))
  } catch (error) {
    console.error('Error fetching role pages:', error)
    viewRolePages.value = []
  }

  showViewDialog.value = true
}

const confirmDeleteRole = (role) => {
  selectedRole.value = role
  showDeleteDialog.value = true
}

const openPermissionsDialog = async (role) => {
  selectedRole.value = role

  // Reset permissions and pageEnabled
  Object.keys(pagePermissions).forEach(key => {
    delete pagePermissions[key]
  })
  Object.keys(pageEnabled).forEach(key => {
    delete pageEnabled[key]
  })

  // Initialize all pages with empty arrays
  allPages.value.forEach(page => {
    pagePermissions[page.id] = []
    pageEnabled[page.id] = false
  })

  // Load existing permissions for this role
  try {
    const existingRolePages = await getRolePagesByRoleId(role.id)
    existingRolePages.forEach(rp => {
      if (rp.pageId && rp.permissions) {
        pagePermissions[rp.pageId] = [...rp.permissions]
        pageEnabled[rp.pageId] = rp.permissions.length > 0
      }
    })
  } catch (error) {
    console.error('Failed to load role permissions:', error)
  }

  showPermissionsDialog.value = true
}

const closePermissionsDialog = () => {
  showPermissionsDialog.value = false
  selectedRole.value = null
}

// Save functions
const saveRole = async () => {
  if (!formValid.value) return

  saving.value = true
  try {
    const roleData = {
      name: newRole.value.name,
      description: newRole.value.description
    }

    let response
    if (editingRole.value) {
      response = await updateRole(editingRole.value.id, roleData)
    } else {
      response = await createRole(roleData)
    }

    if (response.success) {
      success(editingRole.value ? 'تم تحديث الدور بنجاح' : 'تم إضافة الدور بنجاح')
      closeDialog()
      await fetchData()
    } else {
      showError(response.message || 'حدث خطأ')
    }
  } catch (error) {
    console.error('Error saving role:', error)
    showError('حدث خطأ في حفظ الدور')
  } finally {
    saving.value = false
  }
}

const deleteRoleItem = async () => {
  if (!selectedRole.value) return

  deleting.value = true
  try {
    const response = await apiDeleteRole(selectedRole.value.id)
    if (response.success) {
      success('تم حذف الدور بنجاح')
      showDeleteDialog.value = false
      selectedRole.value = null
      await fetchData()
    } else {
      showError(response.message || 'حدث خطأ في الحذف')
    }
  } catch (error) {
    console.error('Error deleting role:', error)
    showError('حدث خطأ في حذف الدور')
  } finally {
    deleting.value = false
  }
}

const savePermissions = async () => {
  if (!selectedRole.value) return

  savingPermissions.value = true
  try {
    // Build page permissions array
    const permissionsToSave = []
    Object.entries(pagePermissions).forEach(([pageId, perms]) => {
      if (perms && perms.length > 0) {
        permissionsToSave.push({
          pageId: Number(pageId),
          permissions: perms
        })
      }
    })

    const response = await bulkUpdateRolePages(selectedRole.value.id, permissionsToSave)
    if (response.success) {
      success('تم حفظ الصلاحيات بنجاح')
      closePermissionsDialog()
      await fetchData()
    } else {
      showError(response.message || 'حدث خطأ في حفظ الصلاحيات')
    }
  } catch (error) {
    console.error('Failed to save permissions:', error)
    showError('حدث خطأ في حفظ الصلاحيات')
  } finally {
    savingPermissions.value = false
  }
}

// Initialize
onMounted(() => {
  fetchData()
})
</script>

<style scoped>
.roles-container {
  padding: 32px;
  max-width: 1600px;
  margin: 0 auto;
}

/* Roles Header Custom Color */
.roles-header {
  background: linear-gradient(135deg, #018790 0%, #005461 100%) !important;
}

.roles-header::before {
  background: linear-gradient(135deg, #06b6d4 0%, #10b981 50%, #06b6d4 100%) !important;
}

/* Statistics Grid - Only 2 cards */
.stats-grid {
  display: grid;
  grid-template-columns: repeat(2, 1fr);
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

.stat-icon.role {
  background: linear-gradient(135deg, #8b5cf6 0%, #7c3aed 100%) !important;
  box-shadow: 0 6px 16px rgba(139, 92, 246, 0.4);
}

.stat-icon.page {
  background: linear-gradient(135deg, #06b6d4 0%, #0891b2 100%) !important;
  box-shadow: 0 6px 16px rgba(6, 182, 212, 0.4);
}

/* Roles List Header */
.roles-list-header {
  background: linear-gradient(135deg, #0a3d42 0%, #052428 100%);
  border-radius: 16px;
  margin-bottom: 24px;
  padding: 16px 24px;
  border: 2px solid transparent;
  position: relative;
  overflow: hidden;
}

.roles-list-header::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  border-radius: 16px;
  padding: 2px;
  background: linear-gradient(135deg, #8b5cf6 0%, #7c3aed 50%, #8b5cf6 100%);
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
  color: #8b5cf6;
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
  background: linear-gradient(135deg, #8b5cf6 0%, #7c3aed 100%);
  color: white;
  box-shadow: 0 4px 12px rgba(139, 92, 246, 0.3);
}

.list-action-btn.primary:hover {
  transform: translateY(-2px);
  box-shadow: 0 6px 16px rgba(139, 92, 246, 0.4);
}

.list-action-btn i {
  font-size: 16px;
}

/* Role Dialog */
.role-dialog {
  border-radius: 16px !important;
  background: linear-gradient(135deg, #0a3d42 0%, #052428 100%) !important;
  border: 2px solid transparent !important;
  position: relative;
  overflow: hidden;
  direction: rtl;
  text-align: right;
}

.role-dialog::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  border-radius: 16px;
  padding: 2px;
  background: linear-gradient(135deg, #8b5cf6 0%, #7c3aed 50%, #8b5cf6 100%);
  -webkit-mask: linear-gradient(#fff 0 0) content-box, linear-gradient(#fff 0 0);
  mask: linear-gradient(#fff 0 0) content-box, linear-gradient(#fff 0 0);
  -webkit-mask-composite: xor;
  mask-composite: exclude;
  pointer-events: none;
}

.dialog-header {
  background: rgba(139, 92, 246, 0.1) !important;
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
  color: #8b5cf6;
  font-size: 24px;
}

.dialog-content {
  padding: 24px !important;
  position: relative;
  z-index: 1;
  direction: rtl;
  text-align: right;
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
  color: #8b5cf6 !important;
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
  background: linear-gradient(135deg, #8b5cf6 0%, #7c3aed 100%);
  color: white;
  box-shadow: 0 4px 12px rgba(139, 92, 246, 0.3);
}

.dialog-btn.save:hover:not(:disabled) {
  transform: translateY(-2px);
  box-shadow: 0 6px 16px rgba(139, 92, 246, 0.4);
}

.dialog-btn.save:disabled {
  opacity: 0.5;
  cursor: not-allowed;
}

.dialog-btn.delete {
  background: linear-gradient(135deg, #dc2626 0%, #b91c1c 100%);
  color: white;
  box-shadow: 0 4px 12px rgba(220, 38, 38, 0.3);
}

.dialog-btn.delete:hover:not(:disabled) {
  transform: translateY(-2px);
  box-shadow: 0 6px 16px rgba(220, 38, 38, 0.4);
}

.dialog-btn.delete:disabled {
  opacity: 0.5;
  cursor: not-allowed;
}

.dialog-btn i {
  font-size: 18px;
}

/* Role Page Item */
.role-page-item {
  background: rgba(255, 255, 255, 0.05) !important;
  border-radius: 8px;
  margin-bottom: 8px;
}

/* Permissions Panels Styles */
.permissions-panels {
  border-radius: 12px;
}

.permissions-panels :deep(.v-expansion-panel) {
  background: rgba(255, 255, 255, 0.05) !important;
  border: 1px solid rgba(255, 255, 255, 0.1);
  border-radius: 8px !important;
  margin-bottom: 8px;
}

.permissions-panels :deep(.v-expansion-panel--active) {
  border-color: #8b5cf6;
}

.permissions-panels :deep(.v-expansion-panel-title) {
  background: transparent !important;
  color: rgba(255, 255, 255, 0.95) !important;
}

.permissions-panels :deep(.v-expansion-panel-title__overlay) {
  background: transparent !important;
}

.page-panel-title {
  min-height: 56px;
}

.page-panel-content {
  background: rgba(255, 255, 255, 0.02) !important;
  padding: 16px !important;
}

.permissions-panels :deep(.v-expansion-panel-text__wrapper) {
  background: transparent !important;
  padding: 16px !important;
}

.permissions-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(150px, 1fr));
  gap: 12px;
}

.permission-checkbox {
  background: rgba(255, 255, 255, 0.05) !important;
  border-radius: 8px;
  padding: 8px 16px;
  border: 1px solid rgba(255, 255, 255, 0.1);
  transition: all 0.2s;
  margin: 0 !important;
  display: flex;
  flex-direction: row-reverse;
  align-items: center;
  justify-content: space-between;
  gap: 8px;
  cursor: pointer;
}

.permission-checkbox:hover {
  border-color: #8b5cf6;
  background: rgba(139, 92, 246, 0.1) !important;
}

.permission-checkbox .permission-label {
  color: rgba(255, 255, 255, 0.9) !important;
  font-weight: 500;
  font-size: 14px;
  flex: 1;
}

.permission-checkbox :deep(.v-selection-control) {
  min-height: unset !important;
  flex: none;
}

.v-card {
  border-radius: 12px;
}

/* Responsive Styles */
@media (max-width: 1024px) {
  .roles-container {
    padding: 24px;
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

  .permissions-grid {
    grid-template-columns: repeat(2, 1fr);
  }
}

@media (max-width: 768px) {
  .roles-container {
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

  .roles-list-header {
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
  .role-dialog .dialog-header {
    padding: 16px !important;
    font-size: 18px !important;
  }

  .role-dialog .dialog-content {
    padding: 16px !important;
  }

  .role-dialog .dialog-actions {
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

  .permissions-grid {
    grid-template-columns: 1fr 1fr;
    gap: 8px;
  }

  .permission-checkbox {
    padding: 6px 12px;
  }
}

@media (max-width: 480px) {
  .roles-container {
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

  .roles-list-header {
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

  .role-dialog {
    border-radius: 12px !important;
  }

  .role-dialog .dialog-header {
    padding: 12px 16px !important;
    font-size: 16px !important;
    gap: 8px;
  }

  .role-dialog .dialog-header i {
    font-size: 20px;
  }

  .role-dialog .dialog-content {
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

  .permissions-grid {
    grid-template-columns: 1fr;
  }

  .permission-checkbox {
    padding: 8px 12px;
  }

  .permission-label {
    font-size: 13px;
  }

  /* Expansion panels mobile */
  .permissions-panels :deep(.v-expansion-panel-title) {
    padding: 12px !important;
    min-height: 48px !important;
  }

  .page-panel-title {
    min-height: 48px;
  }

  .page-panel-title .d-flex.align-center > div span.font-weight-bold {
    font-size: 14px;
  }

  .page-panel-title .d-flex.align-center > div span.text-caption {
    display: none;
  }
}
</style>
