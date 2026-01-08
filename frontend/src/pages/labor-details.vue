<template>
  <div class="labor-details-container">
    <!-- Page Header -->
    <PageHeader
      title="تفاصيل الأيدي العاملة"
      subtitle="إدارة العمال والأجور اليومية"
      badge="العمال"
      badgeType="success"
      class="labor-details-header"
    >
      <template #actions>
        <button class="page-action-btn secondary" @click="$router.back()">
          <i class="mdi mdi-arrow-right"></i>
          رجوع
        </button>
        <button class="page-action-btn primary" @click="showAddLaborModal = true">
          <i class="mdi mdi-plus"></i>
          إضافة عامل
        </button>
      </template>
    </PageHeader>

    <!-- Summary Cards -->
    <div class="summary-section">
      <div class="summary-card workers-summary">
        <div class="summary-icon">
          <i class="mdi mdi-account-group"></i>
        </div>
        <div class="summary-info">
          <span class="summary-label">إجمالي العمال</span>
          <span class="summary-value">{{ labor.length }}</span>
        </div>
      </div>

      <div class="summary-card hours-summary">
        <div class="summary-icon">
          <i class="mdi mdi-clock-outline"></i>
        </div>
        <div class="summary-info">
          <span class="summary-label">إجمالي ساعات العمل</span>
          <span class="summary-value">{{ totalHours }} ساعة</span>
        </div>
      </div>

      <div class="summary-card salary-summary">
        <div class="summary-icon">
          <i class="mdi mdi-cash"></i>
        </div>
        <div class="summary-info">
          <span class="summary-label">إجمالي الأجور اليومية</span>
          <span class="summary-value">{{ formatCurrency(totalSalaries) }}</span>
        </div>
      </div>
    </div>

    <!-- Labor Section Header -->
    <div class="section-header">
      <div class="section-header-content">
        <i class="mdi mdi-account-hard-hat"></i>
        <span>جدول الأيدي العاملة</span>
      </div>
      <div class="section-header-stats">
        <div class="header-stat">
          <span class="header-stat-value">{{ labor.length }}</span>
          <span class="header-stat-label">عامل</span>
        </div>
        <div class="header-stat">
          <span class="header-stat-value">{{ formatCurrency(totalSalaries) }}</span>
          <span class="header-stat-label">إجمالي الأجور</span>
        </div>
      </div>
    </div>

    <!-- Labor Table -->
    <v-card class="data-card">
      <v-data-table
        :headers="laborHeaders"
        :items="labor"
        class="elevation-0"
      >
        <template v-slot:item.index="{ index }">
          {{ index + 1 }}
        </template>
        <template v-slot:item.salary="{ item }">
          <span class="salary-value">{{ formatCurrency(item.salary) }}</span>
        </template>
        <template v-slot:item.notes="{ item }">
          <v-tooltip v-if="item.notes" location="top" max-width="300">
            <template v-slot:activator="{ props }">
              <span v-bind="props" class="notes-cell text-truncate">{{ item.notes }}</span>
            </template>
            <span>{{ item.notes }}</span>
          </v-tooltip>
          <span v-else class="notes-empty">-</span>
        </template>
        <template v-slot:item.actions="{ item }">
          <v-btn v-if="canUpdate" size="small" color="primary" class="me-2" @click="editLabor(item)">
            <i class="mdi mdi-pencil"></i>
          </v-btn>
          <v-btn v-if="canDelete" size="small" color="error" @click="deleteLabor(item)">
            <i class="mdi mdi-delete"></i>
          </v-btn>
        </template>
      </v-data-table>
    </v-card>

    <!-- Add Labor Modal -->
    <div v-if="showAddLaborModal" class="modal-overlay" @click.self="showAddLaborModal = false">
      <div class="modal-container">
        <div class="modal-header">
          <h3>
            <i class="mdi mdi-account-plus"></i>
            إضافة عامل جديد
          </h3>
          <button class="modal-close" @click="showAddLaborModal = false">
            <i class="mdi mdi-close"></i>
          </button>
        </div>
        
        <div class="modal-body">
          <div class="form-group">
            <label>اسم العامل</label>
            <input type="text" v-model="newLabor.name" placeholder="أدخل اسم العامل" />
          </div>
          <div class="form-group">
            <label>الدور / المهنة</label>
            <input type="text" v-model="newLabor.role" placeholder="أدخل المهنة" />
          </div>
          <div class="form-row">
            <div class="form-group">
              <label>رقم الهاتف</label>
              <input type="text" v-model="newLabor.phone" placeholder="رقم الهاتف" />
            </div>
            <div class="form-group">
              <label>العنوان</label>
              <input type="text" v-model="newLabor.address" placeholder="العنوان" />
            </div>
          </div>
          <div class="form-row">
            <div class="form-group">
              <label>الكمية</label>
              <input type="number" v-model="newLabor.hours" placeholder="الكمية" />
            </div>
            <div class="form-group">
              <label>الأجر اليومي (د.ع)</label>
              <input type="number" v-model="newLabor.salary" placeholder="الأجر" />
            </div>
          </div>
          <div class="form-group">
            <label>ملاحظات</label>
            <textarea v-model="newLabor.notes" placeholder="ملاحظات إضافية" rows="2"></textarea>
          </div>
        </div>

        <div class="modal-footer">
          <button class="btn-cancel" @click="showAddLaborModal = false">إلغاء</button>
          <button class="btn-save" @click="addLabor">حفظ العامل</button>
        </div>
      </div>
    </div>

    <!-- Edit Labor Modal -->
    <div v-if="showEditLaborModal" class="modal-overlay" @click.self="closeEditLaborModal">
      <div class="modal-container">
        <div class="modal-header">
          <h3>
            <i class="mdi mdi-pencil"></i>
            تعديل بيانات العامل
          </h3>
          <button class="modal-close" @click="closeEditLaborModal">
            <i class="mdi mdi-close"></i>
          </button>
        </div>

        <div class="modal-body">
          <div class="form-group">
            <label>اسم العامل</label>
            <input type="text" v-model="editLaborForm.name" placeholder="أدخل اسم العامل" />
          </div>
          <div class="form-group">
            <label>الدور / المهنة</label>
            <input type="text" v-model="editLaborForm.role" placeholder="أدخل المهنة" />
          </div>
          <div class="form-row">
            <div class="form-group">
              <label>رقم الهاتف</label>
              <input type="text" v-model="editLaborForm.phone" placeholder="رقم الهاتف" />
            </div>
            <div class="form-group">
              <label>العنوان</label>
              <input type="text" v-model="editLaborForm.address" placeholder="العنوان" />
            </div>
          </div>
          <div class="form-row">
            <div class="form-group">
              <label>الكمية</label>
              <input type="number" v-model="editLaborForm.hours" placeholder="الكمية" />
            </div>
            <div class="form-group">
              <label>الأجر اليومي (د.ع)</label>
              <input type="number" v-model="editLaborForm.salary" placeholder="الأجر" />
            </div>
          </div>
          <div class="form-group">
            <label>ملاحظات</label>
            <textarea v-model="editLaborForm.notes" placeholder="ملاحظات إضافية" rows="2"></textarea>
          </div>
        </div>

        <div class="modal-footer">
          <button class="btn-cancel" @click="closeEditLaborModal">إلغاء</button>
          <button class="btn-save" @click="saveEditedLabor">حفظ التعديلات</button>
        </div>
      </div>
    </div>

  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { useRoute } from 'vue-router'
import PageHeader from '../components/PageHeader.vue'
import { listLaborByWorkDay, createLabor, updateLabor, deleteLabor as apiDeleteLabor } from '@/api/materials'
import { usePermissions } from '@/composables/usePermissions'
import { useToast } from '@/composables/useToast'

const route = useRoute()
const { success, error: showError } = useToast()
const { canCreate, canUpdate, canDelete } = usePermissions('/workdayLabor')

const loading = ref(false)
const showAddLaborModal = ref(false)
const showEditLaborModal = ref(false)
const editingLabor = ref(null)

// Workday ID from route
const workdayId = ref(route.query.workday_id || null)

// Labor data
const labor = ref([])

// Fetch labor from API
const fetchLabor = async () => {
  if (!workdayId.value) return

  loading.value = true
  try {
    const data = await listLaborByWorkDay(workdayId.value)
    labor.value = (data || []).map(l => ({
      id: l.id,
      name: l.workerName || '',
      role: l.jobTitle || '',
      phone: l.phone || '',
      address: l.address || '',
      hours: l.quantity || 0,
      salary: l.cost || 0,
      notes: l.notes || ''
    }))
  } catch (error) {
    console.error('Error fetching labor:', error)
    showError('حدث خطأ في جلب بيانات العمال')
  } finally {
    loading.value = false
  }
}

const laborHeaders = [
  { title: '#', key: 'index', align: 'center' },
  { title: 'الاسم', key: 'name', align: 'start' },
  { title: 'الدور', key: 'role', align: 'center' },
  { title: 'الكمية', key: 'hours', align: 'center' },
  { title: 'الأجر اليومي', key: 'salary', align: 'center' },
  { title: 'ملاحظات', key: 'notes', align: 'center' },
  { title: 'الإجراءات', key: 'actions', align: 'center', sortable: false }
]

// New labor form
const newLabor = ref({
  name: '',
  role: '',
  phone: '',
  address: '',
  hours: 8,
  salary: '',
  notes: ''
})

// Edit labor form
const editLaborForm = ref({
  name: '',
  role: '',
  phone: '',
  address: '',
  hours: 8,
  salary: '',
  notes: ''
})

// Computed totals
const totalSalaries = computed(() => {
  return labor.value.reduce((sum, item) => sum + item.salary, 0)
})

const totalHours = computed(() => {
  return labor.value.reduce((sum, item) => sum + item.hours, 0)
})

// Methods
const formatCurrency = (amount) => {
  return new Intl.NumberFormat('en-US').format(amount) + ' د.ع'
}

const addLabor = async () => {
  if (!newLabor.value.name || !newLabor.value.salary || !workdayId.value) return

  loading.value = true
  try {
    const laborData = {
      workDayId: Number(workdayId.value),
      workerName: newLabor.value.name,
      jobTitle: newLabor.value.role || null,
      phone: newLabor.value.phone || null,
      address: newLabor.value.address || null,
      quantity: Number(newLabor.value.hours) || 1,
      cost: Number(newLabor.value.salary),
      notes: newLabor.value.notes || null
    }

    await createLabor(laborData)
    success('تم إضافة العامل بنجاح')
    newLabor.value = { name: '', role: '', phone: '', address: '', hours: 8, salary: '', notes: '' }
    showAddLaborModal.value = false
    fetchLabor()
  } catch (error) {
    console.error('Error adding labor:', error)
    showError('حدث خطأ في إضافة العامل')
  } finally {
    loading.value = false
  }
}

const deleteLabor = async (item) => {
  if (!confirm('هل أنت متأكد من حذف هذا العامل؟')) return

  loading.value = true
  try {
    await apiDeleteLabor(item.id)
    success('تم حذف العامل بنجاح')
    fetchLabor()
  } catch (error) {
    console.error('Error deleting labor:', error)
    showError('حدث خطأ في حذف العامل')
  } finally {
    loading.value = false
  }
}

// Edit labor
const editLabor = (item) => {
  editingLabor.value = item
  editLaborForm.value = {
    name: item.name || '',
    role: item.role || '',
    phone: item.phone || '',
    address: item.address || '',
    hours: item.hours || 8,
    salary: item.salary || 0,
    notes: item.notes || ''
  }
  showEditLaborModal.value = true
}

const saveEditedLabor = async () => {
  if (!editLaborForm.value.name || !editLaborForm.value.salary || !editingLabor.value) return

  loading.value = true
  try {
    const laborData = {
      workDayId: Number(workdayId.value),
      workerName: editLaborForm.value.name,
      jobTitle: editLaborForm.value.role || null,
      phone: editLaborForm.value.phone || null,
      address: editLaborForm.value.address || null,
      quantity: Number(editLaborForm.value.hours) || 1,
      cost: Number(editLaborForm.value.salary),
      notes: editLaborForm.value.notes || null
    }

    await updateLabor(editingLabor.value.id, laborData)
    success('تم تحديث بيانات العامل بنجاح')
    closeEditLaborModal()
    fetchLabor()
  } catch (error) {
    console.error('Error updating labor:', error)
    showError('حدث خطأ في تحديث بيانات العامل')
  } finally {
    loading.value = false
  }
}

const closeEditLaborModal = () => {
  showEditLaborModal.value = false
  editingLabor.value = null
  editLaborForm.value = { name: '', role: '', phone: '', address: '', hours: 8, salary: '', notes: '' }
}

// Initialize
onMounted(() => {
  fetchLabor()
})
</script>

<style scoped>
.labor-details-container {
  padding: 32px;
  max-width: 1600px;
  margin: 0 auto;
}

/* Header */
.labor-details-header {
  background: linear-gradient(135deg, #018790 0%, #005461 100%) !important;
}

.labor-details-header::before {
  background: linear-gradient(135deg, #10b981 0%, #059669 50%, #10b981 100%) !important;
}

/* Section Header */
.section-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-top: 32px;
  margin-bottom: 16px;
  padding: 20px 24px;
  background: linear-gradient(135deg, #0a3d42 0%, #052428 100%);
  border-radius: 12px;
  border: 2px solid rgba(16, 185, 129, 0.3);
}

.section-header-content {
  display: flex;
  align-items: center;
  gap: 12px;
  color: rgba(255, 255, 255, 0.95);
  font-size: 18px;
  font-weight: 700;
}

.section-header-content i {
  font-size: 24px;
  color: #10b981;
}

.section-header-stats {
  display: flex;
  gap: 32px;
}

.header-stat {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 4px;
}

.header-stat-value {
  font-size: 20px;
  font-weight: 800;
  background: linear-gradient(135deg, #10b981 0%, #06b6d4 100%);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background-clip: text;
}

.header-stat-label {
  font-size: 12px;
  color: rgba(255, 255, 255, 0.6);
}

/* Data Card */
.data-card {
  border-radius: 16px !important;
  background: linear-gradient(135deg, #0a3d42 0%, #052428 100%) !important;
  border: 2px solid rgba(16, 185, 129, 0.3) !important;
  overflow: hidden;
}

.salary-value {
  font-weight: 700;
  color: #10b981;
}

/* Notes cell with truncation */
.notes-cell {
  display: inline-block;
  max-width: 200px;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
  cursor: pointer;
  color: rgba(255, 255, 255, 0.8);
}

.notes-cell:hover {
  color: #10b981;
}

.notes-empty {
  color: rgba(255, 255, 255, 0.4);
}

/* Summary Section */
.summary-section {
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  gap: 24px;
  margin-top: 24px;
}

.summary-card {
  display: flex;
  align-items: center;
  gap: 16px;
  padding: 24px;
  background: linear-gradient(135deg, #0a3d42 0%, #052428 100%);
  border-radius: 16px;
  border: 2px solid rgba(16, 185, 129, 0.3);
}

.summary-icon {
  width: 60px;
  height: 60px;
  border-radius: 12px;
  display: flex;
  align-items: center;
  justify-content: center;
}

.summary-icon i {
  font-size: 28px;
  color: white;
}

.workers-summary .summary-icon {
  background: linear-gradient(135deg, #10b981 0%, #059669 100%);
}

.hours-summary .summary-icon {
  background: linear-gradient(135deg, #06b6d4 0%, #0891b2 100%);
}

.salary-summary .summary-icon {
  background: linear-gradient(135deg, #14b8a6 0%, #0d9488 100%);
}

.summary-info {
  display: flex;
  flex-direction: column;
  gap: 4px;
}

.summary-label {
  font-size: 14px;
  color: rgba(255, 255, 255, 0.7);
}

.summary-value {
  font-size: 24px;
  font-weight: 800;
  background: linear-gradient(135deg, #10b981 0%, #06b6d4 100%);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background-clip: text;
}

/* Modal Styles */
.modal-overlay {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: rgba(0, 0, 0, 0.7);
  backdrop-filter: blur(5px);
  display: flex;
  justify-content: center;
  align-items: center;
  z-index: 1000;
}

.modal-container {
  background: linear-gradient(135deg, #0a3d42 0%, #052428 100%);
  border-radius: 20px;
  border: 2px solid rgba(16, 185, 129, 0.4);
  width: 90%;
  max-width: 550px;
  max-height: 90vh;
  overflow-y: auto;
}

.modal-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 20px 24px;
  border-bottom: 1px solid rgba(16, 185, 129, 0.2);
}

.modal-header h3 {
  margin: 0;
  color: #fff;
  font-size: 20px;
  display: flex;
  align-items: center;
  gap: 10px;
}

.modal-header h3 i {
  color: #10b981;
}

.modal-close {
  background: rgba(255, 255, 255, 0.1);
  border: none;
  width: 36px;
  height: 36px;
  border-radius: 10px;
  color: rgba(255, 255, 255, 0.7);
  cursor: pointer;
}

.modal-body {
  padding: 24px;
  display: flex;
  flex-direction: column;
  gap: 16px;
}

.form-row {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 16px;
}

.form-group {
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.form-group label {
  color: rgba(255, 255, 255, 0.9);
  font-size: 14px;
  font-weight: 500;
}

.form-group input,
.form-group select,
.form-group textarea {
  background: rgba(255, 255, 255, 0.08);
  border: 1px solid rgba(16, 185, 129, 0.3);
  border-radius: 10px;
  padding: 12px 16px;
  color: #fff;
  font-size: 14px;
  outline: none;
}

.form-group input:focus,
.form-group select:focus,
.form-group textarea:focus {
  border-color: #10b981;
}

.form-group select option {
  background: #0a3d42;
  color: #fff;
}

.modal-footer {
  display: flex;
  justify-content: flex-end;
  gap: 12px;
  padding: 20px 24px;
  border-top: 1px solid rgba(16, 185, 129, 0.2);
}

.btn-cancel,
.btn-save {
  padding: 12px 24px;
  border-radius: 10px;
  font-size: 14px;
  font-weight: 600;
  cursor: pointer;
  border: none;
}

.btn-cancel {
  background: rgba(255, 255, 255, 0.1);
  color: rgba(255, 255, 255, 0.8);
}

.btn-save {
  background: linear-gradient(135deg, #10b981 0%, #059669 100%);
  color: white;
  box-shadow: 0 4px 15px rgba(16, 185, 129, 0.3);
}

/* Responsive */
@media (max-width: 1024px) {
  .summary-section {
    grid-template-columns: 1fr;
  }
  
  .section-header {
    flex-direction: column;
    gap: 16px;
  }
  
  .section-header-stats {
    width: 100%;
    justify-content: space-around;
  }
}

@media (max-width: 768px) {
  .form-row {
    grid-template-columns: 1fr;
  }
}
</style>
