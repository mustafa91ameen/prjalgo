<template>
  <div class="expenses-container">
    <!-- Page Header Component -->
    <PageHeader
      title="إدارة المصروفات"
      subtitle="إدارة وتتبع جميع المصروفات والنفقات"
      badge="المصروفات"
      badgeType="error"
      class="expenses-header"
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
      <!-- Total Expenses Card -->
      <v-card class="stat-card" elevation="0">
        <div class="stat-card-content">
          <div class="stat-icon expense">
            <i class="mdi mdi-cash-minus"></i>
          </div>
          <div class="stat-info">
            <div class="stat-label">إجمالي المصاريف</div>
            <div class="stat-value">{{ formatCurrency(totalExpenses) }}</div>
          </div>
        </div>
      </v-card>

      <!-- Active Expenses Card -->
      <v-card class="stat-card" elevation="0">
        <div class="stat-card-content">
          <div class="stat-icon success">
            <i class="mdi mdi-check-circle"></i>
          </div>
          <div class="stat-info">
            <div class="stat-label">مصاريف نشطة</div>
            <div class="stat-value">{{ activeExpensesCount }}</div>
          </div>
        </div>
      </v-card>

      <!-- Pending Expenses Card -->
      <v-card class="stat-card" elevation="0">
        <div class="stat-card-content">
          <div class="stat-icon warning">
            <i class="mdi mdi-clock-alert"></i>
          </div>
          <div class="stat-info">
            <div class="stat-label">في الانتظار</div>
            <div class="stat-value">{{ pendingExpensesCount }}</div>
          </div>
        </div>
      </v-card>

      <!-- Total Cost Card -->
      <v-card class="stat-card" elevation="0">
        <div class="stat-card-content">
          <div class="stat-icon profit">
            <i class="mdi mdi-cash-multiple"></i>
          </div>
          <div class="stat-info">
            <div class="stat-label">إجمالي التكلفة</div>
            <div class="stat-value">{{ formatCurrency(monthlyExpenses) }}</div>
          </div>
        </div>
      </v-card>
    </div>

    <!-- Expenses List Header -->
    <div class="expenses-list-header">
      <div class="list-header-content">
        <div class="list-header-info">
          <h2 class="list-header-title">
            <i class="mdi mdi-format-list-bulleted"></i>
            قائمة المصروفات
          </h2>
          <p class="list-header-subtitle">عرض جميع المصروفات والنفقات</p>
        </div>
        <div class="list-header-actions">
          <button v-if="canCreate" class="list-action-btn primary" @click="openAddExpenseModal">
            <i class="mdi mdi-plus"></i>
            إضافة مصروف جديد
          </button>
        </div>
      </div>
    </div>

      <!-- Expenses Table -->
      <v-card class="mb-6">
        <v-data-table
          :headers="headers"
          :items="expenses"
          :loading="loading"
          class="elevation-1"
        >
          <template v-slot:item.amount="{ item }">
            <span class="font-weight-bold text-error">{{ formatCurrency(item.amount) }}</span>
          </template>
          <template v-slot:item.date="{ item }">
            {{ formatDate(item.date) }}
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
            <v-btn
              v-if="canUpdate"
              size="small"
              color="primary"
              @click="editExpense(item)"
            >
              <i class="mdi mdi-pencil"></i>
            </v-btn>
            <v-btn
              v-if="canDelete"
              size="small"
              color="error"
              @click="deleteExpense(item)"
            >
              <i class="mdi mdi-delete"></i>
            </v-btn>
          </template>
        </v-data-table>
      </v-card>

      <!-- Add Expense Modal -->
    <div v-if="showAddExpenseModal" class="modal-overlay" @click.self="closeAddExpenseModal">
      <div class="modal-container expense-modal">
        <div class="modal-header">
          <h3>
            <i class="mdi mdi-cash-plus"></i>
            إضافة مصروف جديد
          </h3>
          <button class="modal-close" @click="closeAddExpenseModal">
            <i class="mdi mdi-close"></i>
          </button>
        </div>
        
        <div class="modal-body">
          <!-- اسم المصروف -->
          <div class="form-group">
            <label for="expenseName">اسم المصروف</label>
            <input
              type="text"
              id="expenseName"
              v-model="newExpense.name"
              placeholder="أدخل اسم المصروف"
            />
          </div>

          <!-- المبلغ -->
          <div class="form-group">
            <label for="amount">المبلغ (د.ع)</label>
            <input
              type="number"
              id="amount"
              v-model="newExpense.amount"
              placeholder="أدخل المبلغ"
            />
          </div>

          <!-- تاريخ المصروف -->
          <div class="form-group">
            <label for="expenseDate">تاريخ المصروف</label>
            <input
              type="date"
              id="expenseDate"
              v-model="newExpense.expenseDate"
            />
          </div>

          <!-- نوع المصروف -->
          <div class="form-group">
            <label for="expenseType">نوع المصروف</label>
            <input
              type="text"
              id="expenseType"
              v-model="newExpense.type"
              placeholder="أدخل نوع المصروف"
            />
          </div>

          <!-- المشروع -->
          <div class="form-group">
            <label for="projectId">المشروع (اختياري)</label>
            <select id="projectId" v-model="newExpense.projectId">
              <option :value="null">بدون مشروع</option>
              <option v-for="project in availableProjects" :key="project.id" :value="project.id">
                {{ project.name }}
              </option>
            </select>
          </div>

          <!-- الحالة -->
          <div class="form-group">
            <label for="expenseStatus">الحالة</label>
            <select id="expenseStatus" v-model="newExpense.status">
              <option value="pending">معلق</option>
              <option value="approved">موافق عليه</option>
              <option value="rejected">مرفوض</option>
            </select>
          </div>

          <!-- هل هذا لتسديد دين؟ -->
          <div class="form-group">
            <label class="debt-checkbox-label">
              <input
                type="checkbox"
                v-model="isDebtPayment"
                @change="onDebtPaymentChange"
              />
              <span>هل هذا المصروف لتسديد دين؟</span>
            </label>
          </div>

          <!-- اختيار المدين (يظهر فقط عند تحديد تسديد دين) -->
          <div v-if="isDebtPayment" class="form-group">
            <label for="debtorId">اختر المدين <span class="required-star">*</span></label>
            <select
              id="debtorId"
              v-model="selectedDebtorId"
              @change="onDebtorChange"
              :disabled="loadingDebtors"
            >
              <option :value="null">{{ loadingDebtors ? 'جاري التحميل...' : 'اختر المدين' }}</option>
              <option v-for="debtor in debtorItems" :key="debtor.value" :value="debtor.value">
                {{ debtor.title }}
              </option>
            </select>
          </div>

          <!-- معلومات المدين -->
          <div v-if="isDebtPayment && selectedDebtor" class="debtor-info-card">
            <div class="debtor-info-row">
              <span class="debtor-info-label">إجمالي الدين:</span>
              <span class="debtor-info-value">{{ formatCurrency(selectedDebtor.totalDebt) }}</span>
            </div>
            <div class="debtor-info-row">
              <span class="debtor-info-label">المسدد:</span>
              <span class="debtor-info-value paid">{{ formatCurrency(selectedDebtor.paidAmount) }}</span>
            </div>
            <div class="debtor-info-row">
              <span class="debtor-info-label">المتبقي:</span>
              <span class="debtor-info-value remaining">{{ formatCurrency(selectedDebtor.remainingDebt) }}</span>
            </div>
          </div>

          <!-- تحذير تجاوز المبلغ المتبقي -->
          <div v-if="amountExceedsRemaining" class="amount-warning">
            <i class="mdi mdi-alert"></i>
            المبلغ المدخل ({{ formatCurrency(newExpense.amount) }}) يتجاوز المبلغ المتبقي ({{ formatCurrency(selectedDebtor?.remainingDebt || 0) }})
          </div>

          <!-- الملاحظات -->
          <div class="form-group">
            <label for="notes">الملاحظات</label>
            <textarea
              id="notes"
              v-model="newExpense.notes"
              placeholder="أدخل الملاحظات"
              rows="3"
            ></textarea>
          </div>
        </div>

        <div class="modal-footer">
          <button class="btn-cancel" @click="closeAddExpenseModal">إلغاء</button>
          <button class="btn-save" @click="saveNewExpense">حفظ المصروف</button>
        </div>
      </div>
    </div>

      <!-- Edit Expense Dialog (Vuetify) -->
      <v-dialog v-model="showAddDialog" max-width="600">
        <v-card class="edit-expense-card">
          <v-card-title class="edit-expense-title">
            <i class="mdi mdi-pencil"></i>
            <span>{{ editingExpense ? 'تعديل المصروف' : 'إضافة مصروف جديد' }}</span>
          </v-card-title>
          <v-card-text>
            <v-form ref="form" v-model="valid">
              <v-text-field
                v-model="expenseForm.description"
                label="اسم المصروف"
                :rules="[v => !!v || 'الاسم مطلوب']"
                required
                variant="outlined"
                density="comfortable"
              />
              <v-text-field
                v-model.number="expenseForm.amount"
                label="المبلغ (د.ع)"
                type="number"
                :rules="[v => v > 0 || 'المبلغ يجب أن يكون أكبر من صفر']"
                required
                variant="outlined"
                density="comfortable"
              />
              <v-text-field
                v-model="expenseForm.expenseDate"
                label="تاريخ المصروف"
                type="date"
                :rules="[v => !!v || 'التاريخ مطلوب']"
                required
                variant="outlined"
                density="comfortable"
              />
              <v-text-field
                v-model="expenseForm.category"
                label="نوع المصروف"
                placeholder="أدخل نوع المصروف"
                variant="outlined"
                density="comfortable"
              />
              <v-select
                v-model="expenseForm.projectId"
                :items="projectItems"
                label="المشروع (اختياري)"
                variant="outlined"
                density="comfortable"
                clearable
              />
              <v-select
                v-model="expenseForm.status"
                :items="statusOptions"
                label="الحالة"
                :rules="[v => !!v || 'الحالة مطلوبة']"
                required
                variant="outlined"
                density="comfortable"
              />
              <v-textarea
                v-model="expenseForm.notes"
                label="ملاحظات"
                rows="3"
                variant="outlined"
                density="comfortable"
              />
            </v-form>
          </v-card-text>
          <v-card-actions class="edit-expense-actions">
            <v-spacer />
            <v-btn variant="outlined" @click="closeDialog">إلغاء</v-btn>
            <v-btn color="error" variant="flat" @click="saveExpense" :disabled="!valid">
              <i class="mdi mdi-content-save me-1"></i>
              حفظ التعديلات
            </v-btn>
          </v-card-actions>
        </v-card>
      </v-dialog>

    <!-- Delete Confirmation Dialog -->
    <v-dialog v-model="showDeleteDialog" max-width="420" content-class="delete-dialog-wrapper">
      <v-card class="delete-dialog" dir="rtl">
        <div class="delete-dialog-header">
          <div class="delete-icon-wrapper">
            <i class="mdi mdi-delete-alert"></i>
          </div>
          <h3>تأكيد حذف المصروف</h3>
        </div>

        <div v-if="deletingExpense" class="delete-dialog-body">
          <div class="delete-item-info">
            <span class="delete-item-label">المصروف:</span>
            <span class="delete-item-value">{{ deletingExpense.description }}</span>
          </div>
          <div class="delete-item-info">
            <span class="delete-item-label">المبلغ:</span>
            <span class="delete-item-value amount">{{ formatCurrency(deletingExpense.amount) }}</span>
          </div>

          <div class="delete-warning">
            <i class="mdi mdi-alert-circle"></i>
            <span>هذا الإجراء لا يمكن التراجع عنه!</span>
          </div>
        </div>

        <div class="delete-dialog-actions">
          <button class="delete-btn cancel" @click="showDeleteDialog = false">
            <i class="mdi mdi-close"></i>
            إلغاء
          </button>
          <button class="delete-btn confirm" @click="confirmDeleteExpense" :disabled="deleteLoading">
            <i class="mdi mdi-delete"></i>
            {{ deleteLoading ? 'جاري الحذف...' : 'تأكيد الحذف' }}
          </button>
        </div>
      </v-card>
    </v-dialog>

  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import PageHeader from '../components/PageHeader.vue'
import { listExpenses, createExpense, updateExpense, deleteExpense as apiDeleteExpense } from '@/api/expenses'
import { getProjectsDropdown } from '@/api/projects'
import { getActiveDebtorsWithRemaining } from '@/api/debtors'
import { usePermissions } from '@/composables/usePermissions'
import { useToast } from '@/composables/useToast'
import { DEFAULT_PAGE, DEFAULT_LIMIT } from '@/constants/pagination'

const { success, error: showError } = useToast()
const { canCreate, canUpdate, canDelete } = usePermissions('/expenses')

// Reactive data
const loading = ref(false)
const showAddDialog = ref(false)
const valid = ref(false)
const editingExpense = ref(null)
const showAddExpenseModal = ref(false)

// Delete dialog state
const showDeleteDialog = ref(false)
const deletingExpense = ref(null)
const deleteLoading = ref(false)

// Pagination
const page = ref(DEFAULT_PAGE)
const limit = ref(DEFAULT_LIMIT)
const total = ref(0)

// Available projects
const availableProjects = ref([])

// Debtor payment state
const isDebtPayment = ref(false)
const activeDebtors = ref([])
const selectedDebtorId = ref(null)
const loadingDebtors = ref(false)

// New Expense Form - aligned with backend CreateExpense DTO
const newExpense = ref({
  name: '',
  amount: '',
  expenseDate: '',
  type: '',
  projectId: null,
  status: 'pending',
  notes: ''
})

const expenseForm = ref({
  description: '',
  amount: 0,
  category: '',
  expenseDate: '',
  status: 'pending',
  notes: '',
  projectId: null,
  debtorId: null
})

const expenses = ref([])

// Fetch expenses from API
const fetchExpenses = async () => {
  loading.value = true
  try {
    const response = await listExpenses({ page: page.value, limit: limit.value })
    // listExpenses returns { data, total, page, limit, totalPages }
    console.log('Expenses API response:', response)
    const items = response.data || []
    expenses.value = items.map(e => ({
      id: e.id,
      description: e.name || e.description || '',
      amount: e.amount || 0,
      category: e.type || e.category || '',
      status: e.status || 'pending',
      date: e.expenseDate || e.date || e.createdAt,
      notes: e.notes || '',
      projectId: e.projectId,
      projectName: e.projectName,
      debtorId: e.debtorId,
      debtorName: e.debtorName
    }))
    total.value = response.total || 0
  } catch (error) {
    console.error('Error fetching expenses:', error)
    showError('حدث خطأ في جلب المصروفات')
  } finally {
    loading.value = false
  }
}

// Fetch projects dropdown
const fetchProjects = async () => {
  try {
    const response = await getProjectsDropdown()
    // getProjectsDropdown returns { data: [...] } or array directly
    console.log('Projects dropdown response:', response)
    availableProjects.value = response?.data || response || []
  } catch (error) {
    console.error('Error fetching projects:', error)
  }
}

// Load active debtors with remaining debt for dropdown
const loadActiveDebtors = async () => {
  loadingDebtors.value = true
  try {
    const result = await getActiveDebtorsWithRemaining()
    console.log('Active debtors response:', result)
    activeDebtors.value = Array.isArray(result) ? result : []
  } catch (err) {
    console.error('Error loading active debtors:', err)
    activeDebtors.value = []
  } finally {
    loadingDebtors.value = false
  }
}

// Get selected debtor info for validation
const selectedDebtor = computed(() => {
  if (!selectedDebtorId.value) return null
  return activeDebtors.value.find(d => d.id === selectedDebtorId.value)
})

// Check if amount exceeds remaining debt
const amountExceedsRemaining = computed(() => {
  if (!isDebtPayment.value || !selectedDebtor.value) return false
  return Number(newExpense.value.amount) > selectedDebtor.value.remainingDebt
})

// Format debtors for dropdown display
const debtorItems = computed(() => {
  return activeDebtors.value.map(d => ({
    value: d.id,
    title: `${d.name} - متبقي: ${formatCurrency(d.remainingDebt)}`
  }))
})

// Format projects for dropdown display in edit dialog
const projectItems = computed(() => {
  return availableProjects.value.map(p => ({
    value: p.id,
    title: p.name
  }))
})

// Handle debt payment checkbox change
const onDebtPaymentChange = async () => {
  if (isDebtPayment.value && activeDebtors.value.length === 0) {
    await loadActiveDebtors()
  }
  if (!isDebtPayment.value) {
    selectedDebtorId.value = null
  }
}

// Handle debtor selection change
const onDebtorChange = () => {
  // Auto-set expense type to "سداد دين" when a debtor is selected
  if (selectedDebtorId.value) {
    newExpense.value.type = 'سداد دين'
  }
}

const statusOptions = [
  { value: 'pending', title: 'معلق' },
  { value: 'approved', title: 'موافق عليه' },
  { value: 'rejected', title: 'مرفوض' }
]

const headers = [
  { title: 'الوصف', key: 'description', align: 'start' },
  { title: 'المبلغ', key: 'amount', align: 'center' },
  { title: 'الفئة', key: 'category', align: 'center' },
  { title: 'الحالة', key: 'status', align: 'center' },
  { title: 'التاريخ', key: 'date', align: 'center' },
  { title: 'ملاحظات', key: 'notes', align: 'center' },
  { title: 'الإجراءات', key: 'actions', align: 'center', sortable: false }
]

// Computed properties
const activeExpensesCount = computed(() => {
  return expenses.value.filter(expense => expense.status === 'approved').length
})

const pendingExpensesCount = computed(() => {
  return expenses.value.filter(expense => expense.status === 'pending').length
})

const totalExpenses = computed(() => {
  return expenses.value.reduce((sum, item) => sum + item.amount, 0)
})

const monthlyExpenses = computed(() => {
  const currentMonth = new Date().getMonth()
  const currentYear = new Date().getFullYear()
  
  return expenses.value
    .filter(item => {
      const itemDate = new Date(item.date)
      return itemDate.getMonth() === currentMonth && itemDate.getFullYear() === currentYear
    })
    .reduce((sum, item) => sum + item.amount, 0)
})

const expenseGrowth = computed(() => {
  // Simplified growth calculation
  return 8.2
})

// Methods
const formatCurrency = (amount) => {
  return new Intl.NumberFormat('en-US').format(amount) + ' د.ع'
}

const formatDate = (dateString) => {
  return new Date(dateString).toLocaleDateString('en-US')
}

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

const editExpense = (item) => {
  editingExpense.value = item
  // Map frontend display fields to form fields
  expenseForm.value = {
    description: item.description || '',
    amount: item.amount || 0,
    category: item.category || '',
    expenseDate: item.date ? new Date(item.date).toISOString().split('T')[0] : '',
    status: item.status || 'pending',
    notes: item.notes || '',
    projectId: item.projectId || null,
    debtorId: item.debtorId || null
  }
  showAddDialog.value = true
}

const deleteExpense = (item) => {
  deletingExpense.value = item
  showDeleteDialog.value = true
}

const confirmDeleteExpense = async () => {
  if (!deletingExpense.value) return

  deleteLoading.value = true
  try {
    await apiDeleteExpense(deletingExpense.value.id)
    success('تم حذف المصروف بنجاح')
    showDeleteDialog.value = false
    deletingExpense.value = null
    fetchExpenses()
  } catch (error) {
    console.error('Error deleting expense:', error)
    showError(error.message || 'حدث خطأ في حذف المصروف')
  } finally {
    deleteLoading.value = false
  }
}

const saveExpense = async () => {
  if (!valid.value) return

  loading.value = true
  try {
    const expenseData = {
      name: expenseForm.value.description,
      amount: Number(expenseForm.value.amount),
      type: expenseForm.value.category || null,
      expenseDate: expenseForm.value.expenseDate ? new Date(expenseForm.value.expenseDate).toISOString() : new Date().toISOString(),
      status: expenseForm.value.status,
      notes: expenseForm.value.notes || null,
      projectId: expenseForm.value.projectId ? Number(expenseForm.value.projectId) : null,
      debtorId: expenseForm.value.debtorId ? Number(expenseForm.value.debtorId) : null
    }

    if (editingExpense.value) {
      await updateExpense(editingExpense.value.id, expenseData)
      success('تم تحديث المصروف بنجاح')
    } else {
      await createExpense(expenseData)
      success('تم إضافة المصروف بنجاح')
    }
    closeDialog()
    fetchExpenses()
  } catch (error) {
    console.error('Error saving expense:', error)
    showError(error.message || 'حدث خطأ في حفظ المصروف')
  } finally {
    loading.value = false
  }
}

const closeDialog = () => {
  showAddDialog.value = false
  editingExpense.value = null
  expenseForm.value = {
    description: '',
    amount: 0,
    category: '',
    expenseDate: '',
    status: 'pending',
    notes: '',
    projectId: null,
    debtorId: null
  }
}

// New Expense Modal Functions
const openAddExpenseModal = () => {
  showAddExpenseModal.value = true
}

const closeAddExpenseModal = () => {
  showAddExpenseModal.value = false
  newExpense.value = {
    name: '',
    amount: '',
    expenseDate: '',
    type: '',
    projectId: null,
    status: 'pending',
    notes: ''
  }
  // Reset debtor state
  isDebtPayment.value = false
  selectedDebtorId.value = null
}

const saveNewExpense = async () => {
  if (!newExpense.value.name || !newExpense.value.amount || !newExpense.value.expenseDate) {
    showError('يرجى ملء الحقول المطلوبة (الاسم، المبلغ، التاريخ)')
    return
  }

  // Validate debtor selection if debt payment is checked
  if (isDebtPayment.value && !selectedDebtorId.value) {
    showError('يرجى اختيار المدين')
    return
  }

  // Validate amount doesn't exceed remaining debt
  if (amountExceedsRemaining.value) {
    showError('المبلغ المدخل يتجاوز المبلغ المتبقي للمدين')
    return
  }

  loading.value = true
  try {
    const expenseData = {
      name: newExpense.value.name,
      amount: Number(newExpense.value.amount),
      expenseDate: new Date(newExpense.value.expenseDate).toISOString(),
      type: newExpense.value.type || null,
      projectId: newExpense.value.projectId ? Number(newExpense.value.projectId) : null,
      status: newExpense.value.status || null,
      notes: newExpense.value.notes || null,
      debtorId: isDebtPayment.value && selectedDebtorId.value ? Number(selectedDebtorId.value) : null
    }

    await createExpense(expenseData)
    success('تم إضافة المصروف بنجاح')
    closeAddExpenseModal()
    fetchExpenses()
  } catch (error) {
    console.error('Error creating expense:', error)
    showError(error.message || 'حدث خطأ في إضافة المصروف')
  } finally {
    loading.value = false
  }
}

onMounted(() => {
  fetchExpenses()
  fetchProjects()
})
</script>

<style scoped>
.expenses-container {
  padding: 32px;
  max-width: 1600px;
  margin: 0 auto;
}

/* Expenses Header Custom Color */
.expenses-header {
  background: linear-gradient(135deg, #018790 0%, #005461 100%) !important;
}

.expenses-header::before {
  background: linear-gradient(135deg, #06b6d4 0%, #10b981 50%, #06b6d4 100%) !important;
}

/* Statistics Grid - Same as projects page */
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

.stat-change {
  display: flex;
  align-items: center;
  gap: 4px;
  font-size: 11px;
  font-weight: 600;
  padding: 4px 10px;
  border-radius: 12px;
  background: rgba(255, 255, 255, 0.1);
  backdrop-filter: blur(10px);
}

.stat-change.positive {
  color: #34d399;
}

.stat-change.negative {
  color: #f87171;
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

.stat-icon.expense {
  background: linear-gradient(135deg, #ef4444 0%, #dc2626 100%) !important;
  box-shadow: 0 6px 16px rgba(239, 68, 68, 0.4);
}

.stat-icon.success {
  background: linear-gradient(135deg, #10b981 0%, #059669 100%) !important;
  box-shadow: 0 6px 16px rgba(16, 185, 129, 0.4);
}

.stat-icon.warning {
  background: linear-gradient(135deg, #14b8a6 0%, #2dd4bf 100%) !important;
  box-shadow: 0 6px 16px rgba(20, 184, 166, 0.4);
}

.stat-icon.profit {
  background: linear-gradient(135deg, #10b981 0%, #34d399 100%) !important;
  box-shadow: 0 6px 16px rgba(16, 185, 129, 0.4);
}

/* Expenses List Header */
.expenses-list-header {
  background: linear-gradient(135deg, #0a3d42 0%, #052428 100%);
  border-radius: 16px;
  margin-bottom: 24px;
  padding: 16px 24px;
  border: 2px solid transparent;
  position: relative;
  overflow: hidden;
}

.expenses-list-header::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  border-radius: 16px;
  padding: 2px;
  background: linear-gradient(135deg, #ef4444 0%, #dc2626 50%, #ef4444 100%);
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
  color: #ef4444;
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
  background: linear-gradient(135deg, #ef4444 0%, #dc2626 100%);
  color: white;
  box-shadow: 0 4px 12px rgba(239, 68, 68, 0.3);
}

.list-action-btn.primary:hover {
  transform: translateY(-2px);
  box-shadow: 0 6px 16px rgba(239, 68, 68, 0.4);
}

.list-action-btn i {
  font-size: 16px;
}

.v-card {
  border-radius: 12px;
}

/* Notes cell styling */
.notes-cell {
  display: inline-block;
  max-width: 200px;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
  cursor: pointer;
  font-size: 13px;
  color: rgba(255, 255, 255, 0.8);
}

.notes-cell:hover {
  color: #ef4444;
}

.notes-empty {
  color: rgba(255, 255, 255, 0.4);
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
  animation: fadeIn 0.3s ease;
}

@keyframes fadeIn {
  from { opacity: 0; }
  to { opacity: 1; }
}

.modal-container {
  background: linear-gradient(135deg, #0a3d42 0%, #052428 100%);
  border-radius: 20px;
  border: 2px solid rgba(239, 68, 68, 0.4);
  width: 90%;
  max-width: 600px;
  max-height: 90vh;
  overflow-y: auto;
  box-shadow: 0 25px 60px rgba(0, 0, 0, 0.5);
  animation: slideUp 0.3s ease;
}

@keyframes slideUp {
  from { 
    opacity: 0;
    transform: translateY(30px);
  }
  to { 
    opacity: 1;
    transform: translateY(0);
  }
}

.modal-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 20px 24px;
  border-bottom: 1px solid rgba(239, 68, 68, 0.2);
}

.modal-header h3 {
  margin: 0;
  color: #fff;
  font-size: 20px;
  font-weight: 600;
  display: flex;
  align-items: center;
  gap: 10px;
}

.modal-header h3 i {
  color: #ef4444;
  font-size: 24px;
}

.modal-close {
  background: rgba(255, 255, 255, 0.1);
  border: none;
  width: 36px;
  height: 36px;
  border-radius: 10px;
  color: rgba(255, 255, 255, 0.7);
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: center;
  transition: all 0.3s ease;
}

.modal-close:hover {
  background: rgba(239, 68, 68, 0.3);
  color: #ef4444;
}

.modal-close i {
  font-size: 20px;
}

.modal-body {
  padding: 24px;
  display: flex;
  flex-direction: column;
  gap: 20px;
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
  border: 1px solid rgba(239, 68, 68, 0.3);
  border-radius: 10px;
  padding: 12px 16px;
  color: #fff;
  font-size: 14px;
  transition: all 0.3s ease;
  outline: none;
}

.form-group input::placeholder,
.form-group textarea::placeholder {
  color: rgba(255, 255, 255, 0.4);
}

.form-group input:focus,
.form-group select:focus,
.form-group textarea:focus {
  border-color: #ef4444;
  background: rgba(239, 68, 68, 0.1);
  box-shadow: 0 0 0 3px rgba(239, 68, 68, 0.1);
}

.form-group select {
  cursor: pointer;
}

.form-group select option {
  background: #0a3d42;
  color: #fff;
}

.form-group textarea {
  resize: vertical;
  min-height: 80px;
}

.modal-footer {
  display: flex;
  justify-content: flex-end;
  gap: 12px;
  padding: 20px 24px;
  border-top: 1px solid rgba(239, 68, 68, 0.2);
}

.btn-cancel,
.btn-save {
  padding: 12px 24px;
  border-radius: 10px;
  font-size: 14px;
  font-weight: 600;
  cursor: pointer;
  transition: all 0.3s ease;
  border: none;
}

.btn-cancel {
  background: rgba(255, 255, 255, 0.1);
  color: rgba(255, 255, 255, 0.8);
}

.btn-cancel:hover {
  background: rgba(255, 255, 255, 0.2);
}

.btn-save {
  background: linear-gradient(135deg, #ef4444 0%, #dc2626 100%);
  color: white;
  box-shadow: 0 4px 15px rgba(239, 68, 68, 0.3);
}

.btn-save:hover {
  transform: translateY(-2px);
  box-shadow: 0 6px 20px rgba(239, 68, 68, 0.4);
}

/* Debtor Payment Styles */
.debt-checkbox-label {
  display: flex;
  align-items: center;
  gap: 10px;
  cursor: pointer;
  color: rgba(255, 255, 255, 0.9);
  font-size: 14px;
}

.debt-checkbox-label input[type="checkbox"] {
  width: 18px;
  height: 18px;
  cursor: pointer;
  accent-color: #ef4444;
}

.required-star {
  color: #ef4444;
}

.debtor-info-card {
  background: rgba(239, 68, 68, 0.1);
  border: 1px solid rgba(239, 68, 68, 0.3);
  border-radius: 10px;
  padding: 16px;
  margin-top: 8px;
}

.debtor-info-row {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 8px 0;
  border-bottom: 1px solid rgba(255, 255, 255, 0.1);
}

.debtor-info-row:last-child {
  border-bottom: none;
}

.debtor-info-label {
  color: rgba(255, 255, 255, 0.7);
  font-size: 13px;
}

.debtor-info-value {
  color: rgba(255, 255, 255, 0.95);
  font-weight: 600;
  font-size: 14px;
}

.debtor-info-value.paid {
  color: #10b981;
}

.debtor-info-value.remaining {
  color: #ef4444;
}

.amount-warning {
  background: rgba(239, 68, 68, 0.2);
  border: 1px solid rgba(239, 68, 68, 0.5);
  border-radius: 10px;
  padding: 12px 16px;
  color: #fca5a5;
  font-size: 13px;
  display: flex;
  align-items: center;
  gap: 8px;
  margin-top: 8px;
}

.amount-warning i {
  font-size: 18px;
  color: #ef4444;
}

/* Edit Expense Dialog Styles */
.edit-expense-card {
  background: linear-gradient(135deg, #0a3d42 0%, #052428 100%) !important;
  border: 2px solid rgba(239, 68, 68, 0.4) !important;
  border-radius: 16px !important;
  direction: rtl !important;
  text-align: right !important;
}

.edit-expense-title {
  display: flex !important;
  align-items: center;
  gap: 10px;
  color: #fff !important;
  font-size: 20px !important;
  font-weight: 600 !important;
  padding: 20px 24px !important;
  border-bottom: 1px solid rgba(239, 68, 68, 0.2);
  direction: rtl !important;
}

.edit-expense-title i {
  color: #ef4444;
  font-size: 24px;
}

.edit-expense-card :deep(.v-card-text) {
  padding: 24px !important;
  direction: rtl !important;
}

.edit-expense-card :deep(.v-form) {
  direction: rtl !important;
}

.edit-expense-card :deep(.v-text-field),
.edit-expense-card :deep(.v-select),
.edit-expense-card :deep(.v-textarea) {
  margin-bottom: 16px;
  direction: rtl !important;
}

.edit-expense-card :deep(.v-field) {
  background: rgba(255, 255, 255, 0.08) !important;
  border-color: rgba(239, 68, 68, 0.3) !important;
  direction: rtl !important;
}

.edit-expense-card :deep(.v-field__input) {
  direction: rtl !important;
  text-align: right !important;
}

.edit-expense-card :deep(.v-field__outline) {
  direction: rtl !important;
}

.edit-expense-card :deep(.v-field-label) {
  right: 12px !important;
  left: auto !important;
  transform-origin: right !important;
}

.edit-expense-card :deep(.v-field--variant-outlined .v-field-label) {
  margin-inline-start: 0 !important;
  margin-inline-end: 12px !important;
}

.edit-expense-card :deep(.v-field:hover) {
  border-color: rgba(239, 68, 68, 0.5) !important;
}

.edit-expense-card :deep(.v-field--focused) {
  border-color: #ef4444 !important;
  box-shadow: 0 0 0 3px rgba(239, 68, 68, 0.1) !important;
}

.edit-expense-card :deep(.v-field__input),
.edit-expense-card :deep(.v-field-label),
.edit-expense-card :deep(.v-select__selection-text) {
  color: rgba(255, 255, 255, 0.9) !important;
}

.edit-expense-card :deep(.v-field-label--floating) {
  color: rgba(255, 255, 255, 0.7) !important;
}

.edit-expense-card :deep(.v-select__selection) {
  direction: rtl !important;
  text-align: right !important;
}

.edit-expense-card :deep(.v-input__append),
.edit-expense-card :deep(.v-field__append-inner) {
  margin-inline-start: 0 !important;
  margin-inline-end: auto !important;
}

.edit-expense-card :deep(.v-field__clearable) {
  margin-inline-start: 4px !important;
  margin-inline-end: 0 !important;
}

.edit-expense-card :deep(.v-messages) {
  direction: rtl !important;
  text-align: right !important;
}

.edit-expense-actions {
  padding: 16px 24px !important;
  border-top: 1px solid rgba(239, 68, 68, 0.2);
  direction: rtl !important;
}

.edit-expense-actions :deep(.v-btn) {
  margin-inline-start: 8px !important;
  margin-inline-end: 0 !important;
}

/* Responsive Styles */
@media (max-width: 1024px) {
  .expenses-container {
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

  .list-action-btn.primary {
    width: 100%;
    justify-content: center;
  }
}

@media (max-width: 768px) {
  .expenses-container {
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

  .stat-value {
    font-size: 24px;
  }

  .expenses-list-header {
    padding: 14px 18px;
  }

  .list-header-title {
    font-size: 16px;
  }

  .list-action-btn {
    padding: 8px 14px;
    font-size: 12px;
  }

  /* Modal responsive */
  .modal-container {
    width: 95%;
    max-height: 85vh;
    margin: 16px;
  }

  .modal-header {
    padding: 16px 20px;
  }

  .modal-header h3 {
    font-size: 18px;
  }

  .modal-body {
    padding: 20px;
    gap: 16px;
  }

  .form-row {
    grid-template-columns: 1fr;
  }

  .modal-footer {
    padding: 16px 20px;
  }

  /* Vuetify dialog responsive */
  :deep(.v-dialog) {
    margin: 16px !important;
  }

  .edit-expense-title {
    font-size: 18px !important;
    padding: 16px 20px !important;
  }

  .edit-expense-card :deep(.v-card-text) {
    padding: 20px !important;
  }

  .edit-expense-actions {
    padding: 14px 20px !important;
  }

  /* Table responsive */
  .v-data-table :deep(.v-table__wrapper) {
    overflow-x: auto;
  }

  .v-data-table :deep(th),
  .v-data-table :deep(td) {
    white-space: nowrap;
    font-size: 13px;
    padding: 10px 12px !important;
  }
}

@media (max-width: 480px) {
  .expenses-container {
    padding: 12px;
  }

  .stat-card {
    border-radius: 12px !important;
  }

  .stat-icon {
    width: 42px;
    height: 42px;
    font-size: 20px;
  }

  .stat-value {
    font-size: 20px;
  }

  .stat-label {
    font-size: 12px;
  }

  .expenses-list-header {
    padding: 12px 14px;
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

  .list-action-btn i {
    font-size: 14px;
  }

  /* Modal full width */
  .modal-container {
    width: calc(100% - 16px);
    max-width: none;
    max-height: 90vh;
    margin: 8px;
    border-radius: 16px;
  }

  .modal-header {
    padding: 14px 16px;
  }

  .modal-header h3 {
    font-size: 16px;
    gap: 8px;
  }

  .modal-header h3 i {
    font-size: 20px;
  }

  .modal-close {
    width: 32px;
    height: 32px;
    border-radius: 8px;
  }

  .modal-body {
    padding: 16px;
    gap: 14px;
  }

  .form-group label {
    font-size: 13px;
  }

  .form-group input,
  .form-group select,
  .form-group textarea {
    padding: 10px 14px;
    font-size: 13px;
    border-radius: 8px;
  }

  .modal-footer {
    padding: 14px 16px;
    flex-wrap: wrap;
    gap: 8px;
  }

  .btn-cancel,
  .btn-save {
    padding: 10px 18px;
    font-size: 13px;
    flex: 1;
    min-width: 100px;
    text-align: center;
    justify-content: center;
  }

  /* Debtor info card */
  .debtor-info-card {
    padding: 12px;
  }

  .debtor-info-label {
    font-size: 12px;
  }

  .debtor-info-value {
    font-size: 13px;
  }

  .amount-warning {
    padding: 10px 12px;
    font-size: 12px;
  }

  /* Vuetify dialog full width */
  :deep(.v-dialog) {
    margin: 8px !important;
  }

  :deep(.v-dialog > .v-overlay__content) {
    max-width: calc(100% - 16px) !important;
  }

  .edit-expense-title {
    font-size: 16px !important;
    padding: 14px 16px !important;
  }

  .edit-expense-title i {
    font-size: 20px;
  }

  .edit-expense-card :deep(.v-card-text) {
    padding: 16px !important;
  }

  .edit-expense-actions {
    padding: 12px 16px !important;
    flex-wrap: wrap;
    gap: 8px;
  }

  /* Table horizontal scroll */
  .v-data-table :deep(.v-table__wrapper) {
    overflow-x: auto;
    -webkit-overflow-scrolling: touch;
  }

  .v-data-table :deep(th),
  .v-data-table :deep(td) {
    font-size: 12px;
    padding: 8px 10px !important;
  }

  /* Notes cell */
  .notes-cell {
    max-width: 120px;
    font-size: 12px;
  }
}

@media (max-width: 360px) {
  .expenses-container {
    padding: 8px;
  }

  .stat-card-content {
    padding: 12px;
    gap: 12px;
  }

  .stat-icon {
    width: 38px;
    height: 38px;
    font-size: 18px;
  }

  .stat-value {
    font-size: 18px;
  }

  .expenses-list-header {
    padding: 10px 12px;
  }

  .list-header-title {
    font-size: 13px;
  }

  .modal-header {
    padding: 12px 14px;
  }

  .modal-header h3 {
    font-size: 15px;
  }

  .modal-body {
    padding: 12px;
  }

  .modal-footer {
    padding: 12px 14px;
  }

  .btn-cancel,
  .btn-save {
    padding: 8px 14px;
    font-size: 12px;
  }

  .edit-expense-title {
    font-size: 15px !important;
    padding: 12px 14px !important;
  }

  .edit-expense-card :deep(.v-card-text) {
    padding: 12px !important;
  }

  .edit-expense-actions {
    padding: 10px 12px !important;
  }
}

/* Delete Dialog */
.delete-dialog {
  background: linear-gradient(135deg, #0a3d42 0%, #052428 100%) !important;
  border-radius: 16px !important;
  border: 2px solid transparent !important;
  position: relative;
  overflow: hidden;
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

.delete-dialog-header {
  background: linear-gradient(135deg, #dc2626 0%, #b91c1c 100%);
  padding: 20px 24px;
  display: flex;
  align-items: center;
  gap: 14px;
  position: relative;
  z-index: 1;
}

.delete-icon-wrapper {
  width: 44px;
  height: 44px;
  background: rgba(255, 255, 255, 0.2);
  border-radius: 12px;
  display: flex;
  align-items: center;
  justify-content: center;
}

.delete-icon-wrapper i {
  font-size: 24px;
  color: white;
}

.delete-dialog-header h3 {
  margin: 0;
  font-size: 18px;
  font-weight: 600;
  color: white;
}

.delete-dialog-body {
  padding: 24px;
  position: relative;
  z-index: 1;
}

.delete-item-info {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 12px 16px;
  background: rgba(255, 255, 255, 0.05);
  border-radius: 10px;
  margin-bottom: 12px;
}

.delete-item-label {
  font-size: 14px;
  color: rgba(255, 255, 255, 0.7);
  font-weight: 500;
}

.delete-item-value {
  font-size: 14px;
  color: rgba(255, 255, 255, 0.95);
  font-weight: 600;
  text-align: left;
  max-width: 60%;
  word-break: break-word;
}

.delete-item-value.amount {
  color: #f87171;
  font-size: 15px;
}

.delete-warning {
  display: flex;
  align-items: center;
  gap: 10px;
  padding: 14px 16px;
  background: rgba(220, 38, 38, 0.15);
  border: 1px solid rgba(220, 38, 38, 0.3);
  border-radius: 10px;
  margin-top: 16px;
}

.delete-warning i {
  font-size: 20px;
  color: #f87171;
  flex-shrink: 0;
}

.delete-warning span {
  font-size: 13px;
  color: #fca5a5;
  font-weight: 500;
}

.delete-dialog-actions {
  display: flex;
  justify-content: flex-end;
  gap: 12px;
  padding: 16px 24px;
  background: rgba(0, 0, 0, 0.2);
  position: relative;
  z-index: 1;
}

.delete-btn {
  display: inline-flex;
  align-items: center;
  justify-content: center;
  gap: 8px;
  padding: 10px 20px;
  border-radius: 10px;
  font-size: 14px;
  font-weight: 600;
  cursor: pointer;
  transition: all 0.3s ease;
  border: none;
  min-width: 110px;
}

.delete-btn i {
  font-size: 18px;
}

.delete-btn.cancel {
  background: rgba(255, 255, 255, 0.1);
  color: rgba(255, 255, 255, 0.8);
  border: 1px solid rgba(255, 255, 255, 0.2);
}

.delete-btn.cancel:hover {
  background: rgba(255, 255, 255, 0.15);
  color: white;
}

.delete-btn.confirm {
  background: linear-gradient(135deg, #dc2626 0%, #b91c1c 100%);
  color: white;
  box-shadow: 0 4px 12px rgba(220, 38, 38, 0.3);
}

.delete-btn.confirm:hover:not(:disabled) {
  transform: translateY(-2px);
  box-shadow: 0 6px 16px rgba(220, 38, 38, 0.4);
}

.delete-btn.confirm:disabled {
  opacity: 0.6;
  cursor: not-allowed;
  transform: none;
}

/* Delete dialog responsive */
@media (max-width: 480px) {
  .delete-dialog-header {
    padding: 16px 18px;
    gap: 12px;
  }

  .delete-icon-wrapper {
    width: 38px;
    height: 38px;
  }

  .delete-icon-wrapper i {
    font-size: 20px;
  }

  .delete-dialog-header h3 {
    font-size: 16px;
  }

  .delete-dialog-body {
    padding: 18px;
  }

  .delete-item-info {
    padding: 10px 14px;
    flex-direction: column;
    align-items: flex-start;
    gap: 4px;
  }

  .delete-item-value {
    max-width: 100%;
    text-align: right;
  }

  .delete-dialog-actions {
    padding: 14px 18px;
    flex-direction: column-reverse;
    gap: 10px;
  }

  .delete-btn {
    width: 100%;
    padding: 12px 16px;
  }
}
</style>
