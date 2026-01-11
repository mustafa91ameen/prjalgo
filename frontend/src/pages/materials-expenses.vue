<template>
  <div class="materials-expenses-container">
    <!-- Page Header -->
    <PageHeader
      title="المواد والمصاريف اليومية"
      subtitle="إدارة المواد والمصروفات اليومية للمشروع"
      badge="المواد"
      badgeType="info"
      class="materials-expenses-header"
    >
      <template #actions>
        <button class="page-action-btn secondary" @click="$router.back()">
          <i class="mdi mdi-arrow-right"></i>
          رجوع
        </button>
        <button v-if="canCreate" class="page-action-btn primary" @click="showAddMaterialModal = true">
          <i class="mdi mdi-plus"></i>
          إضافة مادة
        </button>
      </template>
    </PageHeader>

    <!-- Summary Cards -->
    <div class="summary-section">
      <div class="summary-card materials-summary">
        <div class="summary-icon">
          <i class="mdi mdi-package-variant"></i>
        </div>
        <div class="summary-info">
          <span class="summary-label">إجمالي تكلفة المواد</span>
          <span class="summary-value">{{ formatCurrency(totalMaterialsCost) }}</span>
        </div>
      </div>

      <div class="summary-card expenses-summary">
        <div class="summary-icon">
          <i class="mdi mdi-cash-minus"></i>
        </div>
        <div class="summary-info">
          <span class="summary-label">إجمالي المصاريف اليومية</span>
          <span class="summary-value">{{ formatCurrency(totalDailyExpenses) }}</span>
        </div>
      </div>

      <div class="summary-card total-summary">
        <div class="summary-icon">
          <i class="mdi mdi-calculator"></i>
        </div>
        <div class="summary-info">
          <span class="summary-label">الإجمالي الكلي</span>
          <span class="summary-value">{{ formatCurrency(grandTotal) }}</span>
        </div>
      </div>
    </div>

    <!-- Materials Section -->
    <div class="section-header">
      <div class="section-header-content">
        <i class="mdi mdi-package-variant"></i>
        <span>جدول المواد</span>
      </div>
      <button v-if="canCreate" class="add-btn" @click="showAddMaterialModal = true">
        <i class="mdi mdi-plus"></i>
        إضافة مادة
      </button>
    </div>

    <v-card class="data-card">
      <v-data-table
        :headers="materialsHeaders"
        :items="materials"
        class="elevation-0"
      >
        <template v-slot:item.index="{ index }">
          {{ index + 1 }}
        </template>
        <template v-slot:item.cost="{ item }">
          <span class="cost-value">{{ formatCurrency(item.cost) }}</span>
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
          <v-btn v-if="canUpdate" size="small" color="primary" class="me-2">
            <i class="mdi mdi-pencil"></i>
          </v-btn>
          <v-btn v-if="canDelete" size="small" color="error" @click="deleteMaterial(item)">
            <i class="mdi mdi-delete"></i>
          </v-btn>
        </template>
      </v-data-table>
    </v-card>

    <!-- Daily Expenses Section -->
    <div class="section-header mt-8">
      <div class="section-header-content">
        <i class="mdi mdi-cash-minus"></i>
        <span>جدول المصاريف اليومية</span>
      </div>
      <button v-if="canCreate" class="add-btn expense" @click="showAddExpenseModal = true">
        <i class="mdi mdi-plus"></i>
        إضافة مصروف
      </button>
    </div>

    <v-card class="data-card">
      <v-data-table
        :headers="expensesHeaders"
        :items="dailyExpenses"
        class="elevation-0"
      >
        <template v-slot:item.index="{ index }">
          {{ index + 1 }}
        </template>
        <template v-slot:item.amount="{ item }">
          <span class="cost-value expense">{{ formatCurrency(item.amount) }}</span>
        </template>
        <template v-slot:item.status="{ item }">
          <v-chip
            :color="getStatusColor(item.status)"
            size="small"
          >
            {{ getStatusText(item.status) }}
          </v-chip>
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
          <v-btn v-if="canUpdate" size="small" color="primary" class="me-2" @click="editExpense(item)">
            <i class="mdi mdi-pencil"></i>
          </v-btn>
          <v-btn v-if="canDelete" size="small" color="error" @click="deleteExpense(item)">
            <i class="mdi mdi-delete"></i>
          </v-btn>
        </template>
      </v-data-table>
    </v-card>

    <!-- Add Material Modal -->
    <div v-if="showAddMaterialModal" class="modal-overlay" @click.self="showAddMaterialModal = false">
      <div class="modal-container">
        <div class="modal-header">
          <h3>
            <i class="mdi mdi-package-variant"></i>
            إضافة مادة جديدة
          </h3>
          <button class="modal-close" @click="showAddMaterialModal = false">
            <i class="mdi mdi-close"></i>
          </button>
        </div>
        
        <div class="modal-body">
          <div class="form-group">
            <label>اسم المادة</label>
            <input type="text" v-model="newMaterial.name" placeholder="أدخل اسم المادة" />
          </div>
          <div class="form-group">
            <label>الكمية</label>
            <input type="text" v-model="newMaterial.quantity" placeholder="مثال: 50 كيس" />
          </div>
          <div class="form-group">
            <label>التكلفة (د.ع)</label>
            <input type="number" v-model="newMaterial.cost" placeholder="أدخل التكلفة" />
          </div>
          <div class="form-group">
            <label>ملاحظات</label>
            <textarea v-model="newMaterial.notes" placeholder="ملاحظات إضافية" rows="2"></textarea>
          </div>
        </div>

        <div class="modal-footer">
          <button class="btn-cancel" @click="showAddMaterialModal = false">إلغاء</button>
          <button class="btn-save" @click="addMaterial">حفظ المادة</button>
        </div>
      </div>
    </div>

    <!-- Add Expense Modal -->
    <div v-if="showAddExpenseModal" class="modal-overlay" @click.self="showAddExpenseModal = false">
      <div class="modal-container expense-modal">
        <div class="modal-header">
          <h3>
            <i class="mdi mdi-cash-minus"></i>
            إضافة مصروف يومي
          </h3>
          <button class="modal-close" @click="showAddExpenseModal = false">
            <i class="mdi mdi-close"></i>
          </button>
        </div>

        <div class="modal-body">
          <div class="form-group">
            <label>وصف المصروف</label>
            <input type="text" v-model="newExpense.description" placeholder="أدخل وصف المصروف" />
          </div>
          <div class="form-group">
            <label>نوع المصروف</label>
            <input type="text" v-model="newExpense.category" placeholder="أدخل نوع المصروف" />
          </div>
          <div class="form-group">
            <label>المبلغ (د.ع)</label>
            <input type="number" v-model="newExpense.amount" placeholder="أدخل المبلغ" />
          </div>
          <div class="form-group">
            <label>الحالة</label>
            <select v-model="newExpense.status">
              <option v-for="option in statusOptions" :key="option.value" :value="option.value">
                {{ option.text }}
              </option>
            </select>
          </div>
          <div class="form-group">
            <label>ملاحظات</label>
            <textarea v-model="newExpense.notes" placeholder="ملاحظات إضافية" rows="2"></textarea>
          </div>
        </div>

        <div class="modal-footer">
          <button class="btn-cancel" @click="showAddExpenseModal = false">إلغاء</button>
          <button class="btn-save expense" @click="addExpense">حفظ المصروف</button>
        </div>
      </div>
    </div>

    <!-- Edit Expense Modal -->
    <div v-if="showEditExpenseModal" class="modal-overlay" @click.self="closeEditExpenseModal">
      <div class="modal-container expense-modal">
        <div class="modal-header">
          <h3>
            <i class="mdi mdi-pencil"></i>
            تعديل المصروف
          </h3>
          <button class="modal-close" @click="closeEditExpenseModal">
            <i class="mdi mdi-close"></i>
          </button>
        </div>

        <div class="modal-body">
          <div class="form-group">
            <label>وصف المصروف</label>
            <input type="text" v-model="editExpenseForm.description" placeholder="أدخل وصف المصروف" />
          </div>
          <div class="form-group">
            <label>نوع المصروف</label>
            <input type="text" v-model="editExpenseForm.category" placeholder="أدخل نوع المصروف" />
          </div>
          <div class="form-group">
            <label>المبلغ (د.ع)</label>
            <input type="number" v-model="editExpenseForm.amount" placeholder="أدخل المبلغ" />
          </div>
          <div class="form-group">
            <label>الحالة</label>
            <select v-model="editExpenseForm.status">
              <option v-for="option in statusOptions" :key="option.value" :value="option.value">
                {{ option.text }}
              </option>
            </select>
          </div>
          <div class="form-group">
            <label>ملاحظات</label>
            <textarea v-model="editExpenseForm.notes" placeholder="ملاحظات إضافية" rows="2"></textarea>
          </div>
        </div>

        <div class="modal-footer">
          <button class="btn-cancel" @click="closeEditExpenseModal">إلغاء</button>
          <button class="btn-save expense" @click="saveEditedExpense">حفظ التعديلات</button>
        </div>
      </div>
    </div>

  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { useRoute } from 'vue-router'
import PageHeader from '../components/PageHeader.vue'
import { listMaterialsByWorkDay, createMaterial, deleteMaterial as apiDeleteMaterial } from '@/api/materials'
import { listExpensesByProject, createExpense, updateExpense, deleteExpense as apiDeleteExpense } from '@/api/expenses'
import { getWorkDay } from '@/api/workdays'
import { usePermissions } from '@/composables/usePermissions'
import { useToast } from '@/composables/useToast'

const route = useRoute()
const { success, error: showError } = useToast()
const { canCreate, canUpdate, canDelete } = usePermissions('/workdayMaterials')

const loading = ref(false)
const showAddMaterialModal = ref(false)
const showAddExpenseModal = ref(false)
const showEditExpenseModal = ref(false)
const editingExpense = ref(null)

// Workday ID from route
const workdayId = ref(route.query.workday_id || null)
const projectId = ref(null)

// Materials data
const materials = ref([])

const materialsHeaders = [
  { title: '#', key: 'index', align: 'center' },
  { title: 'اسم المادة', key: 'name', align: 'start' },
  { title: 'الكمية', key: 'quantity', align: 'center' },
  { title: 'التكلفة', key: 'cost', align: 'center' },
  { title: 'ملاحظات', key: 'notes', align: 'center' },
  { title: 'الإجراءات', key: 'actions', align: 'center', sortable: false }
]

// Daily Expenses data
const dailyExpenses = ref([])

const expensesHeaders = [
  { title: '#', key: 'index', align: 'center' },
  { title: 'الوصف', key: 'description', align: 'start' },
  { title: 'الفئة', key: 'category', align: 'center' },
  { title: 'المبلغ', key: 'amount', align: 'center' },
  { title: 'الحالة', key: 'status', align: 'center' },
  { title: 'ملاحظات', key: 'notes', align: 'center' },
  { title: 'الإجراءات', key: 'actions', align: 'center', sortable: false }
]

// New material form
const newMaterial = ref({
  name: '',
  quantity: '',
  cost: '',
  notes: ''
})

// New expense form
const newExpense = ref({
  description: '',
  category: '',
  amount: '',
  status: 'pending',
  notes: ''
})

// Status options for expense
const statusOptions = [
  { value: 'pending', text: 'معلق' },
  { value: 'approved', text: 'موافق عليه' },
  { value: 'rejected', text: 'مرفوض' }
]

// Edit expense form
const editExpenseForm = ref({
  description: '',
  category: '',
  amount: '',
  status: 'pending',
  notes: ''
})

// Computed totals
const totalMaterialsCost = computed(() => {
  return materials.value.reduce((sum, item) => sum + (item.cost * item.quantity), 0)
})

// Only count approved expenses in the total
const totalDailyExpenses = computed(() => {
  return dailyExpenses.value
    .filter(item => item.status === 'approved')
    .reduce((sum, item) => sum + item.amount, 0)
})

const grandTotal = computed(() => {
  return totalMaterialsCost.value + totalDailyExpenses.value
})

// Fetch workday to get projectId
const fetchWorkday = async () => {
  if (!workdayId.value) return
  try {
    const response = await getWorkDay(workdayId.value)
    if (response.success && response.data) {
      projectId.value = response.data.projectId || null
    }
  } catch (error) {
    console.error('Error fetching workday:', error)
  }
}

// Fetch materials from API
const fetchMaterials = async () => {
  if (!workdayId.value) return

  loading.value = true
  try {
    const data = await listMaterialsByWorkDay(workdayId.value)
    materials.value = (data || []).map(m => ({
      id: m.id,
      name: m.materialName,
      quantity: m.quantity || '',
      cost: m.cost || 0,
      notes: m.notes || ''
    }))
  } catch (error) {
    console.error('Error fetching materials:', error)
    showError('حدث خطأ في جلب بيانات المواد')
  } finally {
    loading.value = false
  }
}

// Methods
const formatCurrency = (amount) => {
  return new Intl.NumberFormat('en-US').format(amount) + ' د.ع'
}

// Status helpers
const getStatusColor = (status) => {
  const colors = {
    pending: 'warning',
    approved: 'success',
    rejected: 'error'
  }
  return colors[status] || 'grey'
}

const getStatusText = (status) => {
  const texts = {
    pending: 'معلق',
    approved: 'موافق عليه',
    rejected: 'مرفوض'
  }
  return texts[status] || status
}

const addMaterial = async () => {
  if (!newMaterial.value.name || !newMaterial.value.cost || !workdayId.value) return
  loading.value = true
  try {
    const materialData = {
      workDayId: Number(workdayId.value),
      materialName: newMaterial.value.name,
      quantity: Number(newMaterial.value.quantity) || 1,
      cost: Number(newMaterial.value.cost),
      notes: newMaterial.value.notes || null
    }

    await createMaterial(materialData)
    success('تم إضافة المادة بنجاح')
    newMaterial.value = { name: '', quantity: '', cost: '', notes: '' }
    showAddMaterialModal.value = false
    await fetchMaterials()
  } catch (error) {
    console.error('Error adding material:', error)
    showError('حدث خطأ في إضافة المادة')
  } finally {
    loading.value = false
  }
}

const deleteMaterial = async (item) => {
  if (!confirm('هل أنت متأكد من حذف هذه المادة؟')) return
  loading.value = true
  try {
    await apiDeleteMaterial(item.id)
    success('تم حذف المادة بنجاح')
    await fetchMaterials()
  } catch (error) {
    console.error('Error deleting material:', error)
    showError('حدث خطأ في حذف المادة')
  } finally {
    loading.value = false
  }
}

// Fetch daily expenses from API (filtered by project)
const fetchDailyExpenses = async () => {
  if (!projectId.value) return
  loading.value = true
  try {
    const response = await listExpensesByProject(projectId.value)
    const items = response.data || []
    dailyExpenses.value = items.map(e => ({
      id: e.id,
      description: e.name || '',
      category: e.type || '',
      amount: e.amount || 0,
      status: e.status || 'pending',
      notes: e.notes || ''
    }))
  } catch (error) {
    console.error('Error fetching daily expenses:', error)
    showError('حدث خطأ في جلب بيانات المصاريف')
  } finally {
    loading.value = false
  }
}

const addExpense = async () => {
  if (!newExpense.value.description || !newExpense.value.amount) return
  loading.value = true
  try {
    const expenseData = {
      name: newExpense.value.description,
      amount: Number(newExpense.value.amount),
      expenseDate: new Date().toISOString(),
      type: newExpense.value.category || null,
      status: newExpense.value.status || 'pending',
      notes: newExpense.value.notes || null,
      projectId: projectId.value ? Number(projectId.value) : null
    }

    await createExpense(expenseData)
    success('تم إضافة المصروف بنجاح')
    newExpense.value = { description: '', category: '', amount: '', status: 'pending', notes: '' }
    showAddExpenseModal.value = false
    await fetchDailyExpenses()
  } catch (error) {
    console.error('Error adding expense:', error)
    showError('حدث خطأ في إضافة المصروف')
  } finally {
    loading.value = false
  }
}

// Edit expense
const editExpense = (item) => {
  editingExpense.value = item
  editExpenseForm.value = {
    description: item.description || '',
    category: item.category || '',
    amount: item.amount || 0,
    status: item.status || 'pending',
    notes: item.notes || ''
  }
  showEditExpenseModal.value = true
}

const saveEditedExpense = async () => {
  if (!editExpenseForm.value.description || !editExpenseForm.value.amount || !editingExpense.value) return
  loading.value = true
  try {
    const expenseData = {
      name: editExpenseForm.value.description,
      amount: Number(editExpenseForm.value.amount),
      type: editExpenseForm.value.category || null,
      status: editExpenseForm.value.status || 'pending',
      notes: editExpenseForm.value.notes || null,
      projectId: projectId.value ? Number(projectId.value) : null
    }

    await updateExpense(editingExpense.value.id, expenseData)
    success('تم تحديث المصروف بنجاح')
    closeEditExpenseModal()
    await fetchDailyExpenses()
  } catch (error) {
    console.error('Error updating expense:', error)
    showError('حدث خطأ في تحديث المصروف')
  } finally {
    loading.value = false
  }
}

const closeEditExpenseModal = () => {
  showEditExpenseModal.value = false
  editingExpense.value = null
  editExpenseForm.value = { description: '', category: '', amount: '', status: 'pending', notes: '' }
}

const deleteExpense = async (item) => {
  if (!confirm('هل أنت متأكد من حذف هذا المصروف؟')) return
  loading.value = true
  try {
    await apiDeleteExpense(item.id)
    success('تم حذف المصروف بنجاح')
    await fetchDailyExpenses()
  } catch (error) {
    console.error('Error deleting expense:', error)
    showError('حدث خطأ في حذف المصروف')
  } finally {
    loading.value = false
  }
}

// Fetch data on mount
onMounted(async () => {
  await fetchWorkday()
  fetchMaterials()
  fetchDailyExpenses()
})
</script>

<style scoped>
.materials-expenses-container {
  padding: 32px;
  max-width: 1600px;
  margin: 0 auto;
}

/* Header */
.materials-expenses-header {
  background: linear-gradient(135deg, #018790 0%, #005461 100%) !important;
}

.materials-expenses-header::before {
  background: linear-gradient(135deg, #06b6d4 0%, #10b981 50%, #06b6d4 100%) !important;
}

/* Section Header */
.section-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-top: 32px;
  margin-bottom: 16px;
  padding: 16px 24px;
  background: linear-gradient(135deg, #0a3d42 0%, #052428 100%);
  border-radius: 12px;
  border: 2px solid rgba(6, 182, 212, 0.3);
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
  color: #06b6d4;
}

.add-btn {
  display: flex;
  align-items: center;
  gap: 8px;
  padding: 10px 20px;
  background: linear-gradient(135deg, #06b6d4 0%, #10b981 100%);
  border: none;
  border-radius: 10px;
  color: white;
  font-size: 14px;
  font-weight: 600;
  cursor: pointer;
  transition: all 0.3s ease;
  box-shadow: 0 4px 15px rgba(6, 182, 212, 0.3);
}

.add-btn.expense {
  background: linear-gradient(135deg, #ef4444 0%, #dc2626 100%);
  box-shadow: 0 4px 15px rgba(239, 68, 68, 0.3);
}

.add-btn:hover {
  transform: translateY(-2px);
}

/* Data Card */
.data-card {
  border-radius: 16px !important;
  background: linear-gradient(135deg, #0a3d42 0%, #052428 100%) !important;
  border: 2px solid rgba(6, 182, 212, 0.3) !important;
  overflow: hidden;
}

.cost-value {
  font-weight: 700;
  color: #10b981;
}

.cost-value.expense {
  color: #ef4444;
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
  color: #06b6d4;
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
  border: 2px solid rgba(6, 182, 212, 0.3);
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

.materials-summary .summary-icon {
  background: linear-gradient(135deg, #06b6d4 0%, #0891b2 100%);
}

.expenses-summary .summary-icon {
  background: linear-gradient(135deg, #ef4444 0%, #dc2626 100%);
}

.total-summary .summary-icon {
  background: linear-gradient(135deg, #10b981 0%, #059669 100%);
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
  background: linear-gradient(135deg, #06b6d4 0%, #10b981 100%);
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
  border: 2px solid rgba(6, 182, 212, 0.4);
  width: 90%;
  max-width: 500px;
  max-height: 90vh;
  overflow-y: auto;
}

.expense-modal {
  border-color: rgba(239, 68, 68, 0.4);
}

.modal-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 20px 24px;
  border-bottom: 1px solid rgba(6, 182, 212, 0.2);
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
  color: #06b6d4;
}

.expense-modal .modal-header h3 i {
  color: #ef4444;
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
  border: 1px solid rgba(6, 182, 212, 0.3);
  border-radius: 10px;
  padding: 12px 16px;
  color: #fff;
  font-size: 14px;
  outline: none;
}

.form-group input:focus,
.form-group select:focus,
.form-group textarea:focus {
  border-color: #06b6d4;
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
  border-top: 1px solid rgba(6, 182, 212, 0.2);
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
  background: linear-gradient(135deg, #06b6d4 0%, #10b981 100%);
  color: white;
}

.btn-save.expense {
  background: linear-gradient(135deg, #ef4444 0%, #dc2626 100%);
}

/* Responsive */
@media (max-width: 1024px) {
  .summary-section {
    grid-template-columns: 1fr;
  }
}
</style>
