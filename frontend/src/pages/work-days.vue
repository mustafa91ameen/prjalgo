<template>
  <div class="work-days-container">
    <!-- Page Header Component -->
    <PageHeader
      title="أيام العمل"
      subtitle="عرض وإدارة أيام العمل للمشروع"
      badge="أيام العمل"
      badgeType="info"
      class="work-days-header"
    >
      <template #actions>
        <button class="page-action-btn secondary" @click="$router.back()">
          <i class="mdi mdi-arrow-right"></i>
          رجوع
        </button>
      </template>
    </PageHeader>

    <!-- Project Info Card -->
    <v-card class="project-info-card mb-6">
      <v-card-text>
        <h3 class="project-name">اسم المشروع</h3>
        <p class="project-details">تفاصيل المشروع هنا</p>
      </v-card-text>
    </v-card>

    <!-- Work Days List Header -->
    <div class="work-days-list-header">
      <div class="list-header-content">
        <div class="list-header-info">
          <h2 class="list-header-title">
            <i class="mdi mdi-calendar-clock"></i>
            قائمة أيام العمل
          </h2>
          <p class="list-header-subtitle">عرض وإدارة جميع أيام العمل في المشروع</p>
        </div>
        <div class="list-header-actions">
          <button v-if="canCreate" class="list-action-btn primary" @click="openDialog">
            <i class="mdi mdi-plus"></i>
            إضافة يوم عمل جديد
          </button>
        </div>
      </div>
    </div>

    <!-- Work Days Table -->
    <v-card>
      <v-data-table
        :headers="headers"
        :items="workDays"
        :loading="loading"
        class="elevation-1"
      >
        <template v-slot:item.workDate="{ item }">
          {{ formatDate(item.workDate) }}
        </template>
        <template v-slot:item.totalCost="{ item }">
          {{ formatCurrency(item.totalCost) }}
        </template>
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
            size="small"
            color="#06b6d4"
            class="me-1"
            variant="tonal"
            icon
            @click="viewWorkDay(item)"
          >
            <i class="mdi mdi-eye"></i>
          </v-btn>
          <v-btn
            v-if="canUpdate"
            size="small"
            color="#f59e0b"
            class="me-1"
            variant="tonal"
            icon
            @click="editWorkDay(item)"
          >
            <i class="mdi mdi-pencil"></i>
          </v-btn>
          <v-btn
            v-if="canDelete"
            size="small"
            color="#ef4444"
            class="me-1"
            variant="tonal"
            icon
            @click="deleteWorkDayHandler(item)"
          >
            <i class="mdi mdi-delete"></i>
          </v-btn>
          <v-btn
            size="small"
            color="#10b981"
            variant="tonal"
            icon
          >
            <i class="mdi mdi-file-excel"></i>
          </v-btn>
          <v-btn
            v-if="canUpdate"
            size="small"
            :color="item.status === 'completed' ? '#f59e0b' : '#10b981'"
            variant="tonal"
            icon
            @click="toggleWorkDayStatus(item)"
            :title="item.status === 'completed' ? 'إلغاء الإكمال' : 'تحديد كمكتمل'"
          >
            <i :class="item.status === 'completed' ? 'mdi mdi-undo' : 'mdi mdi-check-circle'"></i>
          </v-btn>
        </template>
      </v-data-table>
    </v-card>

    <!-- Add Work Day Dialog -->
    <v-dialog v-model="showAddDialog" max-width="700" persistent>
      <v-card class="add-work-day-dialog" dir="rtl">
        <v-card-title class="dialog-header">
          <i :class="editingWorkDay ? 'mdi mdi-pencil' : 'mdi mdi-calendar-plus'"></i>
          <span>{{ editingWorkDay ? 'تعديل يوم العمل' : 'إضافة يوم عمل جديد' }}</span>
          <v-spacer></v-spacer>
          <v-btn icon size="small" @click="closeDialog">
            <i class="mdi mdi-close"></i>
          </v-btn>
        </v-card-title>

        <v-card-text class="dialog-body">
          <v-form ref="form" v-model="valid">
            <div class="form-grid">
              <v-text-field v-model="workDayForm.workDate" label="تاريخ العمل" prepend-inner-icon="mdi-calendar-today" variant="outlined" type="date" :rules="[v => !!v || 'تاريخ العمل مطلوب']" required></v-text-field>
              <v-select v-model="workDayForm.workSubCategoryId" label="فئة العمل" prepend-inner-icon="mdi-tag" variant="outlined" :items="workSubCategories" item-title="name" item-value="id" clearable></v-select>
              <v-textarea v-model="workDayForm.description" label="الوصف" prepend-inner-icon="mdi-text" variant="outlined" rows="2" class="full-width"></v-textarea>
              <v-textarea v-model="workDayForm.notes" label="ملاحظات" prepend-inner-icon="mdi-note-text" variant="outlined" rows="2" class="full-width"></v-textarea>
            </div>
          </v-form>
        </v-card-text>

        <v-card-actions class="dialog-actions">
          <v-btn color="grey" variant="outlined" @click="closeDialog">إلغاء</v-btn>
          <v-spacer></v-spacer>
          <v-btn color="primary" variant="elevated" @click="saveWorkDay" :disabled="!valid">حفظ</v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>

    <!-- Delete Confirmation Dialog -->
    <v-dialog v-model="showDeleteDialog" max-width="500">
      <v-card class="delete-dialog" dir="rtl">
        <v-card-title class="dialog-header delete-header">
          <i class="mdi mdi-delete-alert"></i>
          تأكيد حذف يوم العمل
        </v-card-title>

        <v-card-text v-if="deletingWorkDay" class="dialog-body text-center">
          <v-avatar size="60" color="error" class="mb-3">
            <i class="mdi mdi-calendar-remove" style="color: white; font-size: 30px;"></i>
          </v-avatar>
          <h4 style="color: rgba(255,255,255,0.95);">{{ formatDate(deletingWorkDay.workDate) }}</h4>
          <p v-if="deletingWorkDay.description" style="color: rgba(255,255,255,0.7); margin-top: 8px;">
            {{ deletingWorkDay.description }}
          </p>

          <v-alert type="error" variant="tonal" class="my-4" style="text-align: right; direction: rtl;">
            تحذير: هذا الإجراء لا يمكن التراجع عنه!
          </v-alert>

          <p style="color: rgba(255,255,255,0.7);">
            هل أنت متأكد من حذف يوم العمل هذا نهائياً؟
          </p>
        </v-card-text>

        <v-card-actions class="dialog-actions">
          <v-spacer></v-spacer>
          <button class="dialog-btn cancel" @click="showDeleteDialog = false">
            إلغاء
          </button>
          <button class="dialog-btn delete" @click="confirmDelete" :disabled="deleteLoading">
            <i class="mdi mdi-delete"></i>
            {{ deleteLoading ? 'جاري الحذف...' : 'حذف نهائي' }}
          </button>
        </v-card-actions>
      </v-card>
    </v-dialog>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import PageHeader from '../components/PageHeader.vue'
import { listWorkDays, createWorkDay, updateWorkDay, deleteWorkDay as apiDeleteWorkday, listWorkSubCategories, completeWorkDay, uncompleteWorkDay } from '@/api/workdays'
import { usePermissions } from '@/composables/usePermissions'
import { useToast } from '@/composables/useToast'
import { DEFAULT_PAGE, DEFAULT_LIMIT } from '@/constants/pagination'

const route = useRoute()
const router = useRouter()
const { canCreate, canUpdate, canDelete } = usePermissions('/workdays')
const { success, error: showError } = useToast()

const loading = ref(false)

// Pagination
const page = ref(DEFAULT_PAGE)
const limit = ref(DEFAULT_LIMIT)
const total = ref(0)

// Project ID from route
const projectId = ref(route.query.project_id || null)

const workDays = ref([])

// Fetch workdays from API
const fetchWorkDays = async () => {
  loading.value = true
  try {
    const params = { page: page.value, limit: limit.value }
    if (projectId.value) {
      params.projectId = projectId.value
    }
    const response = await listWorkDays(params)
    if (response.success) {
      workDays.value = (response.data.items || []).map(w => ({
        id: w.id,
        projectId: w.projectId,
        workSubCategoryId: w.workSubCategoryId,
        workDate: w.workDate,
        description: w.description || '',
        notes: w.notes || '',
        status: w.status || 'pending',
        totalCost: w.totalCost || 0
      }))
      total.value = response.data.total || 0
    }
  } catch (error) {
    console.error('Error fetching workdays:', error)
    showError('حدث خطأ في جلب أيام العمل')
  } finally {
    loading.value = false
  }
}

// Fetch work sub categories for dropdown
const fetchWorkSubCategories = async () => {
  try {
    const data = await listWorkSubCategories()
    workSubCategories.value = data || []
  } catch (error) {
    console.error('Error fetching work sub categories:', error)
  }
}

// Delete dialog state
const showDeleteDialog = ref(false)
const deletingWorkDay = ref(null)
const deleteLoading = ref(false)

const headers = [
  { title: 'التاريخ', key: 'workDate', align: 'start' },
  { title: 'الوصف', key: 'description', align: 'center' },
  { title: 'التكلفة الإجمالية', key: 'totalCost', align: 'center' },
  { title: 'الحالة', key: 'status', align: 'center' },
  { title: 'الإجراءات', key: 'actions', align: 'center', sortable: false }
]

const formatDate = (date) => {
  if (!date) return '-'
  return new Date(date).toLocaleDateString('en-US')
}

const formatCurrency = (amount) => {
  return new Intl.NumberFormat('en-US').format(amount || 0) + ' د.ع'
}

const getStatusColor = (status) => {
  const colors = {
    completed: 'success',
    inprogress: 'warning',
    pending: 'info'
  }
  return colors[status] || 'grey'
}

const getStatusText = (status) => {
  const texts = {
    completed: 'مكتمل',
    inprogress: 'قيد التنفيذ',
    pending: 'معلق'
  }
  return texts[status] || status
}

// Dialog state
const showAddDialog = ref(false)
const valid = ref(false)
const form = ref(null)
const editingWorkDay = ref(null) // Track which work day is being edited

// Work day form data
const workDayForm = ref({
  workDate: '',
  workSubCategoryId: null,
  description: '',
  notes: ''
})

// Work sub categories for dropdown
const workSubCategories = ref([])

// Open dialog for adding
const openDialog = () => {
  editingWorkDay.value = null
  resetForm()
  showAddDialog.value = true
}

// Open dialog for editing
const editWorkDay = (item) => {
  editingWorkDay.value = item
  workDayForm.value = {
    workDate: item.workDate ? item.workDate.split('T')[0] : '',
    workSubCategoryId: item.workSubCategoryId || null,
    description: item.description || '',
    notes: item.notes || ''
  }
  showAddDialog.value = true
}

// Close dialog
const closeDialog = () => {
  showAddDialog.value = false
  editingWorkDay.value = null
  resetForm()
}

// Reset form
const resetForm = () => {
  workDayForm.value = {
    workDate: '',
    workSubCategoryId: null,
    description: '',
    notes: ''
  }
  if (form.value) {
    form.value.reset()
  }
}

// Save work day (create or update)
const saveWorkDay = async () => {
  if (!valid.value) return

  loading.value = true
  try {
    const workdayData = {
      projectId: parseInt(projectId.value),
      workDate: workDayForm.value.workDate ? new Date(workDayForm.value.workDate).toISOString() : null,
      workSubCategoryId: workDayForm.value.workSubCategoryId || null,
      description: workDayForm.value.description || null,
      notes: workDayForm.value.notes || null
    }

    let response
    if (editingWorkDay.value) {
      // Update existing work day
      response = await updateWorkDay(editingWorkDay.value.id, workdayData)
      if (response.success) {
        success('تم تعديل يوم العمل بنجاح')
        closeDialog()
        fetchWorkDays()
      } else {
        showError(response.message || 'حدث خطأ')
      }
    } else {
      // Create new work day
      response = await createWorkDay(workdayData)
      if (response.success) {
        success('تم إضافة يوم العمل بنجاح')
        closeDialog()
        fetchWorkDays()
      } else {
        showError(response.message || 'حدث خطأ')
      }
    }
  } catch (error) {
    console.error('Error saving workday:', error)
    showError('حدث خطأ في حفظ يوم العمل')
  } finally {
    loading.value = false
  }
}

// Delete work day - open confirmation dialog
const deleteWorkDayHandler = (item) => {
  deletingWorkDay.value = item
  showDeleteDialog.value = true
}

// Confirm delete
const confirmDelete = async () => {
  if (!deletingWorkDay.value) return

  deleteLoading.value = true
  try {
    const response = await apiDeleteWorkday(deletingWorkDay.value.id)
    if (response.success) {
      success('تم حذف يوم العمل بنجاح')
      showDeleteDialog.value = false
      deletingWorkDay.value = null
      fetchWorkDays()
    } else {
      showError(response.message || 'حدث خطأ في الحذف')
    }
  } catch (error) {
    console.error('Error deleting workday:', error)
    showError('حدث خطأ في حذف يوم العمل')
  } finally {
    deleteLoading.value = false
  }
}

// View work day details
const viewWorkDay = (item) => {
  router.push({ path: '/work-day-details', query: { id: item.id } })
}

// Toggle work day status (complete/uncomplete)
const toggleWorkDayStatus = async (item) => {
  loading.value = true
  try {
    let response
    if (item.status === 'completed') {
      // Uncomplete the work day
      response = await uncompleteWorkDay(item.id)
      if (response.success) {
        success('تم إلغاء إكمال يوم العمل')
        fetchWorkDays()
      } else {
        showError(response.message || 'حدث خطأ')
      }
    } else {
      // Complete the work day
      response = await completeWorkDay(item.id)
      if (response.success) {
        success('تم تحديد يوم العمل كمكتمل')
        fetchWorkDays()
      } else {
        showError(response.message || 'حدث خطأ')
      }
    }
  } catch (error) {
    console.error('Error toggling workday status:', error)
    showError('حدث خطأ في تغيير حالة يوم العمل')
  } finally {
    loading.value = false
  }
}

// Initialize
onMounted(() => {
  fetchWorkDays()
  fetchWorkSubCategories()
})
</script>

<style scoped>
.work-days-container {
  padding: 32px;
  max-width: 1600px;
  margin: 0 auto;
}

/* Work Days Header Custom Color */
.work-days-header {
  background: linear-gradient(135deg, #018790 0%, #005461 100%) !important;
}

.work-days-header::before {
  background: linear-gradient(135deg, #06b6d4 0%, #10b981 50%, #06b6d4 100%) !important;
}

/* Project Info Card */
.project-info-card {
  border-radius: 16px !important;
  background: linear-gradient(135deg, #0a3d42 0%, #052428 100%) !important;
  border: 2px solid transparent !important;
  position: relative;
}

.project-info-card::before {
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

.project-name {
  font-size: 24px;
  font-weight: 700;
  color: rgba(255, 255, 255, 0.95);
  margin-bottom: 8px;
}

.project-details {
  font-size: 14px;
  color: rgba(255, 255, 255, 0.7);
  margin: 0;
}

/* Work Days List Header */
.work-days-list-header {
  background: linear-gradient(135deg, #0a3d42 0%, #052428 100%);
  border-radius: 16px;
  margin-bottom: 24px;
  padding: 16px 24px;
  border: 2px solid transparent;
  position: relative;
  overflow: hidden;
}

.work-days-list-header::before {
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
  color: #06b6d4;
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
  background: linear-gradient(135deg, #06b6d4 0%, #0891b2 100%);
  color: white;
  box-shadow: 0 4px 12px rgba(6, 182, 212, 0.3);
}

.list-action-btn.primary:hover {
  transform: translateY(-2px);
  box-shadow: 0 6px 16px rgba(6, 182, 212, 0.4);
}

.list-action-btn i {
  font-size: 16px;
}

/* Add Work Day Dialog */
.add-work-day-dialog {
  background: linear-gradient(135deg, #0a3d42 0%, #052428 100%) !important;
  border-radius: 20px !important;
  border: 2px solid transparent !important;
  position: relative;
  direction: rtl;
  text-align: right;
}

.add-work-day-dialog::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  border-radius: 20px;
  padding: 2px;
  background: linear-gradient(135deg, #06b6d4 0%, #10b981 50%, #14b8a6 100%);
  -webkit-mask: linear-gradient(#fff 0 0) content-box, linear-gradient(#fff 0 0);
  mask: linear-gradient(#fff 0 0) content-box, linear-gradient(#fff 0 0);
  -webkit-mask-composite: xor;
  mask-composite: exclude;
  pointer-events: none;
}

.dialog-header {
  background: rgba(6, 182, 212, 0.15) !important;
  color: rgba(255, 255, 255, 0.95) !important;
  padding: 20px 24px !important;
  font-size: 20px !important;
  font-weight: 700 !important;
  display: flex !important;
  align-items: center !important;
  gap: 12px !important;
  position: relative;
  z-index: 1;
}

.dialog-header i {
  color: #06b6d4;
  font-size: 28px;
}

.dialog-body {
  padding: 24px !important;
  position: relative;
  z-index: 1;
}

.form-grid {
  display: grid;
  grid-template-columns: repeat(2, 1fr);
  gap: 20px;
}

.form-grid .full-width {
  grid-column: 1 / -1;
}

.dialog-actions {
  padding: 16px 24px !important;
  border-top: 1px solid rgba(255, 255, 255, 0.1) !important;
  position: relative;
  z-index: 1;
}

:deep(.v-field) {
  background: rgba(255, 255, 255, 0.05) !important;
  border-radius: 12px !important;
  direction: rtl;
  text-align: right;
}

:deep(.v-field--focused) {
  background: rgba(255, 255, 255, 0.08) !important;
}

:deep(.v-label) {
  color: rgba(255, 255, 255, 0.7) !important;
  right: 12px !important;
  left: auto !important;
}

:deep(.v-field__input) {
  color: rgba(255, 255, 255, 0.95) !important;
  text-align: right;
  direction: rtl;
}

:deep(.v-icon) {
  color: rgba(255, 255, 255, 0.6) !important;
}

:deep(.v-field__prepend-inner) {
  margin-left: 8px;
  margin-right: 0;
}

:deep(.v-select__selection-text) {
  text-align: right;
}


.v-card {
  border-radius: 12px;
}

/* Delete Dialog */
.delete-dialog {
  background: linear-gradient(135deg, #0a3d42 0%, #052428 100%) !important;
  border-radius: 16px !important;
  border: 2px solid transparent !important;
  position: relative;
}

.delete-dialog::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  border-radius: 16px;
  padding: 2px;
  background: linear-gradient(135deg, #dc2626 0%, #b91c1c 50%, #dc2626 100%);
  -webkit-mask: linear-gradient(#fff 0 0) content-box, linear-gradient(#fff 0 0);
  mask: linear-gradient(#fff 0 0) content-box, linear-gradient(#fff 0 0);
  -webkit-mask-composite: xor;
  mask-composite: exclude;
  pointer-events: none;
}

.delete-header {
  background: linear-gradient(135deg, #dc2626 0%, #b91c1c 100%) !important;
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
</style>
