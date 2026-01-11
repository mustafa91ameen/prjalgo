<template>
  <div class="debtors-container">
    <!-- Page Header Component -->
    <PageHeader
      title="إدارة المديونين"
      subtitle="إدارة وتتبع جميع المديونيات والمتأخرات"
      badge="المديونون"
      badgeType="warning"
      class="debtors-header"
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
      <!-- Total Debtors Card -->
      <v-card class="stat-card" elevation="0">
        <div class="stat-card-content">
          <div class="stat-icon warning">
            <i class="mdi mdi-account-multiple"></i>
          </div>
          <div class="stat-info">
            <div class="stat-label">إجمالي المديونين</div>
            <div class="stat-value">{{ totalDebtors }}</div>
          </div>
        </div>
      </v-card>

      <!-- Total Debt Card -->
      <v-card class="stat-card" elevation="0">
        <div class="stat-card-content">
          <div class="stat-icon expense">
            <i class="mdi mdi-cash-remove"></i>
          </div>
          <div class="stat-info">
            <div class="stat-label">إجمالي المديونية</div>
            <div class="stat-value">{{ formatCurrency(totalDebt) }}</div>
          </div>
        </div>
      </v-card>

      <!-- Overdue Card -->
      <v-card class="stat-card" elevation="0">
        <div class="stat-card-content">
          <div class="stat-icon danger">
            <i class="mdi mdi-alert-circle"></i>
          </div>
          <div class="stat-info">
            <div class="stat-label">متأخرين</div>
            <div class="stat-value">{{ overdueDebtors }}</div>
          </div>
        </div>
      </v-card>

      <!-- Paid Card -->
      <v-card class="stat-card" elevation="0">
        <div class="stat-card-content">
          <div class="stat-icon success">
            <i class="mdi mdi-check-circle"></i>
          </div>
          <div class="stat-info">
            <div class="stat-label">مدفوع</div>
            <div class="stat-value">{{ formatCurrency(paidAmount) }}</div>
          </div>
        </div>
      </v-card>
    </div>

    <!-- Debtors List Header -->
    <div class="debtors-list-header">
      <div class="list-header-content">
        <div class="list-header-info">
          <h2 class="list-header-title">
            <i class="mdi mdi-format-list-bulleted"></i>
            قائمة المديونين
          </h2>
          <p class="list-header-subtitle">عرض جميع المديونين والمتأخرات</p>
        </div>
        <div class="list-header-actions">
          <button v-if="canCreate" class="list-action-btn primary" @click="openAddDialog">
            <i class="mdi mdi-plus"></i>
            إضافة مديون جديد
          </button>
        </div>
      </div>
    </div>

    <!-- Debtors Table -->
    <v-card class="mb-6">
      <v-data-table
        :headers="headers"
        :items="debtors"
        :loading="loading"
        class="elevation-1 debtors-data-table"
      >
        <!-- عمود الاسم -->
        <template v-slot:item.name="{ item }">
          <div class="d-flex align-center">
            <v-avatar
              :color="getStatusColor(item.status)"
              size="32"
              class="me-3"
            >
              <span class="text-white font-weight-bold">
                {{ item.name ? item.name.charAt(0) : '?' }}
              </span>
            </v-avatar>
            <div>
              <div class="font-weight-medium">{{ item.name }}</div>
              <div class="text-caption text-grey">{{ item.phone }}</div>
            </div>
          </div>
        </template>

        <!-- عمود المبلغ المطلوب -->
        <template v-slot:item.totalDebt="{ item }">
          <div class="text-center">
            <div class="font-weight-bold">{{ formatCurrency(item.totalDebt) }}</div>
          </div>
        </template>

        <!-- عمود المسدد -->
        <template v-slot:item.paidAmount="{ item }">
          <div class="text-center">
            <div class="font-weight-bold text-success">{{ formatCurrency(item.paidAmount || 0) }}</div>
            <v-progress-linear
              v-if="item.totalDebt > 0"
              :model-value="getPaymentProgress(item)"
              color="success"
              height="6"
              rounded
              class="mt-1 mx-auto"
              style="max-width: 80px;"
            />
          </div>
        </template>

        <!-- عمود المتبقي -->
        <template v-slot:item.remainingAmount="{ item }">
          <div class="text-center">
            <div class="font-weight-bold" :class="item.remainingAmount > 0 ? 'text-error' : 'text-success'">
              {{ formatCurrency(item.remainingAmount || 0) }}
            </div>
            <div class="text-caption text-grey">
              {{ getPaymentProgress(item).toFixed(0) }}% مسدد
            </div>
          </div>
        </template>

        <!-- عمود تاريخ الاستحقاق -->
        <template v-slot:item.dueDate="{ item }">
          <div class="text-center">
            <div class="font-weight-medium">{{ formatDate(item.dueDate) }}</div>
            <v-chip
              :color="getDueDateColor(item.dueDate, item.status)"
              size="small"
              :variant="isDueDateWarning(item.dueDate, item.status) ? 'flat' : 'tonal'"
            >
              <v-icon v-if="isDueDateWarning(item.dueDate, item.status)" size="small" class="me-1">
                mdi-alert
              </v-icon>
              {{ getDueDateStatus(item.dueDate, item.status) }}
            </v-chip>
          </div>
        </template>

        <!-- عمود الحالة -->
        <template v-slot:item.status="{ item }">
          <v-chip
            :color="getStatusColor(item.status)"
            :variant="item.status === 'paid' ? 'flat' : 'tonal'"
            size="small"
          >
            {{ getStatusText(item.status) }}
          </v-chip>
        </template>

        <!-- عمود الملاحظات -->
        <template v-slot:item.notes="{ item }">
          <v-tooltip v-if="item.notes" location="top" max-width="300">
            <template v-slot:activator="{ props }">
              <span v-bind="props" class="notes-cell text-truncate">{{ item.notes }}</span>
            </template>
            <span>{{ item.notes }}</span>
          </v-tooltip>
          <span v-else class="notes-empty">-</span>
        </template>

        <!-- عمود الإجراءات -->
        <template v-slot:item.actions="{ item }">
          <div class="d-flex align-center gap-1">
            <v-btn
              icon="mdi-eye"
              size="small"
              variant="text"
              @click="viewDebtor(item)"
            />
            <v-btn
              v-if="canUpdate"
              icon="mdi-pencil"
              size="small"
              variant="text"
              @click="editDebtor(item)"
            />
            <v-btn
              v-if="canDelete"
              icon="mdi-delete"
              size="small"
              variant="text"
              color="error"
              @click="confirmDeleteDebtor(item)"
            />
          </div>
        </template>
      </v-data-table>
    </v-card>

    <!-- Add Debtor Dialog -->
    <v-dialog v-model="showAddDialog" max-width="800" persistent>
      <v-card class="debtor-dialog">
        <v-card-title class="dialog-header">
          <i class="mdi mdi-account-plus"></i>
          إضافة مديون جديد
        </v-card-title>
        
        <v-card-text class="dialog-content">
          <v-form ref="debtorForm" v-model="formValid">
            <v-row>
              <!-- الاسم الكامل -->
              <v-col cols="12" md="6">
                <v-text-field
                  v-model="newDebtor.fullName"
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
                  v-model="newDebtor.email"
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
                  v-model="newDebtor.phone"
                  label="رقم الهاتف"
                  variant="outlined"
                  density="comfortable"
                  :rules="[v => !!v || 'رقم الهاتف مطلوب']"
                  required
                ></v-text-field>
              </v-col>

              <!-- مبلغ المطلوب -->
              <v-col cols="12" md="6">
                <v-text-field
                  v-model.number="newDebtor.amount"
                  label="مبلغ المطلوب"
                  variant="outlined"
                  density="comfortable"
                  type="number"
                  :rules="[v => v > 0 || 'المبلغ يجب أن يكون أكبر من صفر']"
                  required
                ></v-text-field>
              </v-col>

              <!-- العملة -->
              <v-col cols="12" md="6">
                <v-select
                  v-model="newDebtor.currency"
                  :items="currencies"
                  label="العملة"
                  variant="outlined"
                  density="comfortable"
                  :rules="[v => !!v || 'العملة مطلوبة']"
                  required
                ></v-select>
              </v-col>

              <!-- تاريخ الاستحقاق -->
              <v-col cols="12" md="6">
                <v-text-field
                  v-model="newDebtor.dueDate"
                  label="تاريخ الاستحقاق"
                  variant="outlined"
                  density="comfortable"
                  type="date"
                  :rules="[v => !!v || 'تاريخ الاستحقاق مطلوب']"
                  required
                ></v-text-field>
              </v-col>

              <!-- الملاحظات -->
              <v-col cols="12">
                <v-textarea
                  v-model="newDebtor.notes"
                  label="الملاحظات"
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
          <button class="dialog-btn save" @click="saveDebtor" :disabled="!formValid">
            <i class="mdi mdi-content-save"></i>
            حفظ
          </button>
        </v-card-actions>
      </v-card>
    </v-dialog>

    <!-- View Debtor Details Dialog -->
    <v-dialog v-model="showViewDialog" max-width="600" persistent>
      <v-card class="debtor-dialog">
        <v-card-title class="dialog-header">
          <i class="mdi mdi-account-details"></i>
          تفاصيل المديون
        </v-card-title>

        <v-card-text class="dialog-content" v-if="selectedDebtor">
          <!-- Avatar and Name -->
          <div class="debtor-profile">
            <v-avatar
              :color="getStatusColor(selectedDebtor.status)"
              size="80"
              class="mb-3"
            >
              <span class="text-white text-h4 font-weight-bold">
                {{ selectedDebtor.name ? selectedDebtor.name.charAt(0) : '?' }}
              </span>
            </v-avatar>
            <h3 class="debtor-name">{{ selectedDebtor.name }}</h3>
            <v-chip
              :color="getStatusColor(selectedDebtor.status)"
              :variant="selectedDebtor.status === 'paid' ? 'flat' : 'tonal'"
              size="small"
              class="mt-2"
            >
              {{ getStatusText(selectedDebtor.status) }}
            </v-chip>
          </div>

          <!-- Stats Cards -->
          <div class="debtor-stats-row">
            <div class="debtor-stat-card error-card">
              <div class="stat-icon-small">
                <i class="mdi mdi-cash-multiple"></i>
              </div>
              <div class="stat-content">
                <span class="stat-label">إجمالي الدين</span>
                <span class="stat-value">{{ formatCurrency(selectedDebtor.totalDebt) }}</span>
              </div>
            </div>
            <div class="debtor-stat-card success-card">
              <div class="stat-icon-small">
                <i class="mdi mdi-check-circle"></i>
              </div>
              <div class="stat-content">
                <span class="stat-label">المسدد</span>
                <span class="stat-value">{{ formatCurrency(selectedDebtor.paidAmount || 0) }}</span>
              </div>
            </div>
            <div class="debtor-stat-card warning-card">
              <div class="stat-icon-small">
                <i class="mdi mdi-clock-alert"></i>
              </div>
              <div class="stat-content">
                <span class="stat-label">المتبقي</span>
                <span class="stat-value">{{ formatCurrency(selectedDebtor.remainingAmount || 0) }}</span>
              </div>
            </div>
          </div>

          <!-- Progress Bar -->
          <div class="payment-progress-section">
            <div class="progress-header">
              <span>نسبة السداد</span>
              <span class="progress-percentage">{{ getPaymentProgress(selectedDebtor).toFixed(0) }}%</span>
            </div>
            <v-progress-linear
              :model-value="getPaymentProgress(selectedDebtor)"
              color="success"
              height="10"
              rounded
            />
          </div>

          <!-- Contact Info -->
          <div class="info-section">
            <h4 class="section-title">
              <i class="mdi mdi-card-account-details"></i>
              معلومات الاتصال
            </h4>
            <div class="info-grid">
              <div class="info-item">
                <i class="mdi mdi-email"></i>
                <div>
                  <span class="info-label">البريد الإلكتروني</span>
                  <span class="info-value">{{ selectedDebtor.email || 'غير محدد' }}</span>
                </div>
              </div>
              <div class="info-item">
                <i class="mdi mdi-phone"></i>
                <div>
                  <span class="info-label">رقم الهاتف</span>
                  <span class="info-value">{{ selectedDebtor.phone || 'غير محدد' }}</span>
                </div>
              </div>
            </div>
          </div>

          <!-- Due Date Info -->
          <div class="info-section">
            <h4 class="section-title">
              <i class="mdi mdi-calendar"></i>
              تاريخ الاستحقاق
            </h4>
            <div class="due-date-display">
              <span class="due-date-value">{{ formatDate(selectedDebtor.dueDate) }}</span>
              <v-chip
                :color="getDueDateColor(selectedDebtor.dueDate, selectedDebtor.status)"
                size="small"
                :variant="isDueDateWarning(selectedDebtor.dueDate, selectedDebtor.status) ? 'flat' : 'tonal'"
              >
                <v-icon v-if="isDueDateWarning(selectedDebtor.dueDate, selectedDebtor.status)" size="small" class="me-1">
                  mdi-alert
                </v-icon>
                {{ getDueDateStatus(selectedDebtor.dueDate, selectedDebtor.status) }}
              </v-chip>
            </div>
          </div>

          <!-- Notes -->
          <div class="info-section" v-if="selectedDebtor.notes">
            <h4 class="section-title">
              <i class="mdi mdi-note-text"></i>
              ملاحظات
            </h4>
            <p class="notes-text">{{ selectedDebtor.notes }}</p>
          </div>
        </v-card-text>

        <v-card-actions class="dialog-actions">
          <v-spacer></v-spacer>
          <button class="dialog-btn cancel" @click="closeViewDialog">
            <i class="mdi mdi-close"></i>
            إغلاق
          </button>
          <button v-if="canUpdate" class="dialog-btn save" @click="editFromView">
            <i class="mdi mdi-pencil"></i>
            تعديل
          </button>
        </v-card-actions>
      </v-card>
    </v-dialog>

    <!-- Delete Confirmation Dialog -->
    <v-dialog v-model="showDeleteDialog" max-width="500">
      <v-card class="debtor-dialog delete-dialog">
        <v-card-title class="dialog-header delete-header">
          <i class="mdi mdi-delete-alert"></i>
          تأكيد حذف المديون
        </v-card-title>

        <v-card-text v-if="deletingDebtor" class="dialog-content text-center">
          <v-avatar size="60" color="error" class="mb-3">
            <span class="text-white text-h5 font-weight-bold">
              {{ deletingDebtor.name ? deletingDebtor.name.charAt(0) : '?' }}
            </span>
          </v-avatar>
          <h4 style="color: rgba(255,255,255,0.95);">{{ deletingDebtor.name }}</h4>
          <p style="color: rgba(255,255,255,0.7); margin-top: 8px;">
            المبلغ: {{ formatCurrency(deletingDebtor.totalDebt) }}
          </p>

          <v-alert type="error" variant="tonal" class="my-4" style="text-align: right; direction: rtl;">
            تحذير: هذا الإجراء لا يمكن التراجع عنه!
          </v-alert>

          <p style="color: rgba(255,255,255,0.7);">
            هل أنت متأكد من حذف هذا المديون نهائياً؟
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
import { ref, computed, onMounted } from 'vue'
import PageHeader from '../components/PageHeader.vue'
import { listDebtors, createDebtor, updateDebtor, deleteDebtor as apiDeleteDebtor, getDebtorStats } from '@/api/debtors'
import { usePermissions } from '@/composables/usePermissions'
import { useToast } from '@/composables/useToast'
import { DEFAULT_PAGE, DEFAULT_LIMIT } from '@/constants/pagination'

const { canCreate, canUpdate, canDelete } = usePermissions('/debtors')
const { success, error: showError } = useToast()

const loading = ref(false)
const showAddDialog = ref(false)
const showViewDialog = ref(false)
const formValid = ref(false)
const debtorForm = ref(null)
const editingDebtor = ref(null)
const selectedDebtor = ref(null)

// Delete dialog state
const showDeleteDialog = ref(false)
const deletingDebtor = ref(null)
const deleteLoading = ref(false)

// Pagination
const page = ref(DEFAULT_PAGE)
const limit = ref(DEFAULT_LIMIT)
const total = ref(0)

// Stats
const debtorStats = ref({
  totalDebtors: 0,
  totalDebt: 0,
  overdueCount: 0,
  paidAmount: 0
})

// New debtor data
const newDebtor = ref({
  fullName: '',
  email: '',
  phone: '',
  amount: 0,
  currency: 'IQD',
  dueDate: '',
  notes: ''
})

const currencies = [
  'IQD',
  'USD'
]

const debtors = ref([])

// Fetch debtors from API
const fetchDebtors = async () => {
  loading.value = true
  try {
    const response = await listDebtors({ page: page.value, limit: limit.value })
    // listDebtors returns { data, total, page, limit, totalPages }
    console.log('Debtors API response:', response)
    const items = response.data || []
    debtors.value = items.map(d => ({
      id: d.id,
      name: d.name || d.fullName || '',
      email: d.email || '',
      phone: d.phone || '',
      totalDebt: d.totalDebt || d.amount || 0,
      paidAmount: d.paidAmount || 0,
      remainingAmount: d.remainingAmount || (d.totalDebt - (d.paidAmount || 0)) || 0,
      dueDate: d.dueDate,
      status: d.status || 'pending',
      notes: d.notes || ''
    }))
    total.value = response.total || 0
  } catch (error) {
    console.error('Error fetching debtors:', error)
    showError('حدث خطأ في جلب المديونين')
  } finally {
    loading.value = false
  }
}

// Fetch stats
const fetchStats = async () => {
  try {
    const response = await getDebtorStats()
    // getDebtorStats returns the stats object directly (or wrapped in data)
    console.log('Debtor stats response:', response)
    const stats = response?.data || response || {}
    debtorStats.value = {
      totalDebtors: stats.totalDebtors || stats.total_debtors || 0,
      totalDebt: stats.totalDebt || stats.total_debt || 0,
      overdueCount: stats.overdueCount || stats.overdue_count || 0,
      paidAmount: stats.paidAmount || stats.paid_amount || 0
    }
  } catch (error) {
    console.error('Error fetching stats:', error)
  }
}


const headers = [
  { title: 'الاسم', key: 'name', align: 'start' },
  { title: 'المبلغ المطلوب', key: 'totalDebt', align: 'center' },
  { title: 'المسدد', key: 'paidAmount', align: 'center' },
  { title: 'المتبقي', key: 'remainingAmount', align: 'center' },
  { title: 'تاريخ الاستحقاق', key: 'dueDate', align: 'center' },
  { title: 'الحالة', key: 'status', align: 'center' },
  { title: 'ملاحظات', key: 'notes', align: 'center' },
  { title: 'الإجراءات', key: 'actions', align: 'center', sortable: false }
]

// Computed properties
const totalDebtors = computed(() => debtorStats.value.totalDebtors || total.value)

const totalDebt = computed(() => debtorStats.value.totalDebt || debtors.value.reduce((sum, debtor) => sum + debtor.amount, 0))

const overdueDebtors = computed(() => debtorStats.value.overdueCount || debtors.value.filter(debtor => debtor.status === 'overdue').length)

const paidAmount = computed(() => debtorStats.value.paidAmount || 0)

// Methods
const openAddDialog = () => {
  showAddDialog.value = true
}

const closeDialog = () => {
  showAddDialog.value = false
  resetForm()
}

const resetForm = () => {
  newDebtor.value = {
    fullName: '',
    email: '',
    phone: '',
    amount: 0,
    currency: 'IQD',
    dueDate: '',
    notes: ''
  }
  if (debtorForm.value) {
    debtorForm.value.reset()
  }
}

const saveDebtor = async () => {
  if (!formValid.value) return

  loading.value = true
  try {
    // Convert date to RFC3339 format
    const dueDateRFC3339 = newDebtor.value.dueDate ? new Date(newDebtor.value.dueDate).toISOString() : null

    const debtorData = {
      name: newDebtor.value.fullName,
      email: newDebtor.value.email,
      phone: newDebtor.value.phone,
      totalDebt: newDebtor.value.amount,
      currency: newDebtor.value.currency,
      dueDate: dueDateRFC3339,
      notes: newDebtor.value.notes
    }

    if (editingDebtor.value) {
      await updateDebtor(editingDebtor.value.id, debtorData)
    } else {
      await createDebtor(debtorData)
    }

    // If we reach here, the API call succeeded (apiFetch throws on error)
    success(editingDebtor.value ? 'تم تحديث المديون بنجاح' : 'تم إضافة المديون بنجاح')
    closeDialog()
    fetchDebtors()
    fetchStats()
  } catch (error) {
    console.error('Error saving debtor:', error)
    showError(error.message || 'حدث خطأ في حفظ المديون')
  } finally {
    loading.value = false
  }
}

// Initialize
onMounted(() => {
  fetchDebtors()
  fetchStats()
})

const formatCurrency = (amount) => {
  return new Intl.NumberFormat('en-US').format(amount) + ' د.ع'
}

const getStatusColor = (status) => {
  const colors = {
    overdue: 'error',
    pending: 'warning',
    paid: 'success'
  }
  return colors[status] || 'grey'
}

const getStatusText = (status) => {
  const texts = {
    overdue: 'متأخر',
    pending: 'معلق',
    paid: 'مدفوع'
  }
  return texts[status] || status
}

// Calculate payment progress percentage
const getPaymentProgress = (debtor) => {
  if (!debtor.totalDebt || debtor.totalDebt === 0) return 0
  const paidAmount = debtor.paidAmount || 0
  const progress = (paidAmount / debtor.totalDebt) * 100
  return Math.min(progress, 100) // Cap at 100%
}

// Format date for display
const formatDate = (date) => {
  if (!date) return '-'
  return new Date(date).toLocaleDateString('en-US', {
    year: 'numeric',
    month: '2-digit',
    day: '2-digit'
  })
}

// Get due date color based on status and remaining time
const getDueDateColor = (dueDate, status) => {
  if (status === 'paid') return 'success'
  if (!dueDate) return 'grey'

  const today = new Date()
  today.setHours(0, 0, 0, 0)
  const due = new Date(dueDate)
  due.setHours(0, 0, 0, 0)
  const diffTime = due - today
  const diffDays = Math.ceil(diffTime / (1000 * 60 * 60 * 24))

  if (diffDays < 0) return 'error' // Overdue
  if (diffDays <= 7) return 'warning' // Due soon
  return 'success' // OK
}

// Get due date status text
const getDueDateStatus = (dueDate, status) => {
  if (status === 'paid') return 'مدفوع'
  if (!dueDate) return 'غير محدد'

  const today = new Date()
  today.setHours(0, 0, 0, 0)
  const due = new Date(dueDate)
  due.setHours(0, 0, 0, 0)
  const diffTime = due - today
  const diffDays = Math.ceil(diffTime / (1000 * 60 * 60 * 24))

  if (diffDays < 0) return 'متأخر ' + Math.abs(diffDays) + ' يوم'
  if (diffDays === 0) return 'اليوم'
  if (diffDays === 1) return 'غداً'
  if (diffDays <= 7) return 'خلال ' + diffDays + ' أيام'
  return 'في الموعد'
}

// Check if due date should show warning
const isDueDateWarning = (dueDate, status) => {
  if (status === 'paid' || !dueDate) return false

  const today = new Date()
  today.setHours(0, 0, 0, 0)
  const due = new Date(dueDate)
  due.setHours(0, 0, 0, 0)

  return due <= today
}

// View debtor details
const viewDebtor = (item) => {
  selectedDebtor.value = item
  showViewDialog.value = true
}

const closeViewDialog = () => {
  showViewDialog.value = false
  selectedDebtor.value = null
}

const editFromView = () => {
  const item = selectedDebtor.value
  closeViewDialog()
  editDebtor(item)
}

const editDebtor = (item) => {
  editingDebtor.value = item
  newDebtor.value = {
    fullName: item.name,
    email: item.email || '',
    phone: item.phone || '',
    amount: item.totalDebt,
    currency: 'IQD',
    dueDate: item.dueDate ? item.dueDate.split('T')[0] : '',
    notes: item.notes || ''
  }
  showAddDialog.value = true
}

// Open delete confirmation dialog
const confirmDeleteDebtor = (item) => {
  deletingDebtor.value = item
  showDeleteDialog.value = true
}

// Confirm delete
const confirmDelete = async () => {
  if (!deletingDebtor.value) return

  deleteLoading.value = true
  try {
    await apiDeleteDebtor(deletingDebtor.value.id)
    success('تم حذف المديون بنجاح')
    showDeleteDialog.value = false
    deletingDebtor.value = null
    fetchDebtors()
    fetchStats()
  } catch (error) {
    console.error('Error deleting debtor:', error)
    showError(error.message || 'حدث خطأ في حذف المديون')
  } finally {
    deleteLoading.value = false
  }
}
</script>

<style scoped>
.debtors-container {
  padding: 32px;
  max-width: 1600px;
  margin: 0 auto;
}

/* Debtors Header Custom Color */
.debtors-header {
  background: linear-gradient(135deg, #018790 0%, #005461 100%) !important;
}

.debtors-header::before {
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

.stat-icon.warning {
  background: linear-gradient(135deg, #f59e0b 0%, #d97706 100%) !important;
  box-shadow: 0 6px 16px rgba(245, 158, 11, 0.4);
}

.stat-icon.expense {
  background: linear-gradient(135deg, #ef4444 0%, #dc2626 100%) !important;
  box-shadow: 0 6px 16px rgba(239, 68, 68, 0.4);
}

.stat-icon.danger {
  background: linear-gradient(135deg, #dc2626 0%, #991b1b 100%) !important;
  box-shadow: 0 6px 16px rgba(220, 38, 38, 0.4);
}

.stat-icon.success {
  background: linear-gradient(135deg, #10b981 0%, #059669 100%) !important;
  box-shadow: 0 6px 16px rgba(16, 185, 129, 0.4);
}

/* Debtors List Header */
.debtors-list-header {
  background: linear-gradient(135deg, #0a3d42 0%, #052428 100%);
  border-radius: 16px;
  margin-bottom: 24px;
  padding: 16px 24px;
  border: 2px solid transparent;
  position: relative;
  overflow: hidden;
}

.debtors-list-header::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  border-radius: 16px;
  padding: 2px;
  background: linear-gradient(135deg, #f59e0b 0%, #d97706 50%, #f59e0b 100%);
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
  color: #f59e0b;
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
  background: linear-gradient(135deg, #f59e0b 0%, #d97706 100%);
  color: white;
  box-shadow: 0 4px 12px rgba(245, 158, 11, 0.3);
}

.list-action-btn.primary:hover {
  transform: translateY(-2px);
  box-shadow: 0 6px 16px rgba(245, 158, 11, 0.4);
}

.list-action-btn i {
  font-size: 16px;
}

/* Debtor Dialog */
.debtor-dialog {
  border-radius: 16px !important;
  background: linear-gradient(135deg, #0a3d42 0%, #052428 100%) !important;
  border: 2px solid transparent !important;
  position: relative;
  overflow: hidden;
  direction: rtl;
  text-align: right;
}

.debtor-dialog::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  border-radius: 16px;
  padding: 2px;
  background: linear-gradient(135deg, #f59e0b 0%, #d97706 50%, #f59e0b 100%);
  -webkit-mask: linear-gradient(#fff 0 0) content-box, linear-gradient(#fff 0 0);
  mask: linear-gradient(#fff 0 0) content-box, linear-gradient(#fff 0 0);
  -webkit-mask-composite: xor;
  mask-composite: exclude;
  pointer-events: none;
}

.dialog-header {
  background: rgba(245, 158, 11, 0.1) !important;
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
  color: #f59e0b;
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
  color: #f59e0b !important;
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
  background: linear-gradient(135deg, #f59e0b 0%, #d97706 100%);
  color: white;
  box-shadow: 0 4px 12px rgba(245, 158, 11, 0.3);
}

.dialog-btn.save:hover:not(:disabled) {
  transform: translateY(-2px);
  box-shadow: 0 6px 16px rgba(245, 158, 11, 0.4);
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

/* Delete Dialog */
.delete-dialog::before {
  background: linear-gradient(135deg, #dc2626 0%, #b91c1c 50%, #dc2626 100%) !important;
}

.delete-header {
  background: linear-gradient(135deg, #dc2626 0%, #b91c1c 100%) !important;
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
  color: #f59e0b;
}

.notes-empty {
  color: rgba(255, 255, 255, 0.4);
}

/* View Debtor Dialog Styles */
.debtor-profile {
  display: flex;
  flex-direction: column;
  align-items: center;
  text-align: center;
  padding: 20px 0;
  border-bottom: 1px solid rgba(255, 255, 255, 0.1);
  margin-bottom: 20px;
}

.debtor-name {
  font-size: 24px;
  font-weight: 700;
  color: rgba(255, 255, 255, 0.95);
  margin: 0;
}

.debtor-stats-row {
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  gap: 12px;
  margin-bottom: 20px;
}

.debtor-stat-card {
  display: flex;
  align-items: center;
  gap: 12px;
  padding: 14px;
  border-radius: 12px;
  background: rgba(255, 255, 255, 0.05);
}

.debtor-stat-card.error-card {
  border: 1px solid rgba(239, 68, 68, 0.3);
  background: rgba(239, 68, 68, 0.1);
}

.debtor-stat-card.success-card {
  border: 1px solid rgba(16, 185, 129, 0.3);
  background: rgba(16, 185, 129, 0.1);
}

.debtor-stat-card.warning-card {
  border: 1px solid rgba(245, 158, 11, 0.3);
  background: rgba(245, 158, 11, 0.1);
}

.stat-icon-small {
  width: 40px;
  height: 40px;
  border-radius: 10px;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 20px;
}

.error-card .stat-icon-small {
  background: rgba(239, 68, 68, 0.2);
  color: #ef4444;
}

.success-card .stat-icon-small {
  background: rgba(16, 185, 129, 0.2);
  color: #10b981;
}

.warning-card .stat-icon-small {
  background: rgba(245, 158, 11, 0.2);
  color: #f59e0b;
}

.stat-content {
  display: flex;
  flex-direction: column;
}

.stat-content .stat-label {
  font-size: 11px;
  color: rgba(255, 255, 255, 0.6);
}

.stat-content .stat-value {
  font-size: 14px;
  font-weight: 700;
  color: rgba(255, 255, 255, 0.95);
}

.payment-progress-section {
  margin-bottom: 20px;
  padding: 16px;
  background: rgba(255, 255, 255, 0.05);
  border-radius: 12px;
}

.progress-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 10px;
  font-size: 14px;
  color: rgba(255, 255, 255, 0.8);
}

.progress-percentage {
  font-weight: 700;
  color: #10b981;
}

.info-section {
  margin-bottom: 16px;
  padding: 16px;
  background: rgba(255, 255, 255, 0.05);
  border-radius: 12px;
}

.section-title {
  display: flex;
  align-items: center;
  gap: 8px;
  font-size: 14px;
  font-weight: 600;
  color: rgba(255, 255, 255, 0.8);
  margin: 0 0 12px 0;
}

.section-title i {
  color: #f59e0b;
  font-size: 18px;
}

.info-grid {
  display: grid;
  grid-template-columns: repeat(2, 1fr);
  gap: 16px;
}

.info-item {
  display: flex;
  align-items: flex-start;
  gap: 10px;
}

.info-item > i {
  color: rgba(255, 255, 255, 0.5);
  font-size: 18px;
  margin-top: 2px;
}

.info-item > div {
  display: flex;
  flex-direction: column;
}

.info-label {
  font-size: 11px;
  color: rgba(255, 255, 255, 0.5);
}

.info-value {
  font-size: 14px;
  color: rgba(255, 255, 255, 0.95);
}

.due-date-display {
  display: flex;
  align-items: center;
  gap: 12px;
}

.due-date-value {
  font-size: 16px;
  font-weight: 600;
  color: rgba(255, 255, 255, 0.95);
}

.notes-text {
  margin: 0;
  font-size: 14px;
  color: rgba(255, 255, 255, 0.8);
  line-height: 1.6;
}

/* Responsive Styles */
@media (max-width: 1024px) {
  .debtors-container {
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
  .debtors-container {
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

  .debtors-list-header {
    padding: 14px 18px;
  }

  .list-header-title {
    font-size: 16px;
  }

  .list-action-btn {
    padding: 8px 14px;
    font-size: 12px;
  }

  /* Dialog responsive */
  :deep(.v-dialog) {
    margin: 16px !important;
  }

  .dialog-header {
    font-size: 18px !important;
    padding: 16px 20px !important;
  }

  .dialog-content {
    padding: 20px !important;
  }

  .dialog-actions {
    padding: 14px 20px !important;
  }

  .dialog-btn {
    padding: 8px 18px;
    font-size: 13px;
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

  /* View dialog stats */
  .debtor-stats-row {
    grid-template-columns: 1fr;
  }

  .info-grid {
    grid-template-columns: 1fr;
  }
}

@media (max-width: 480px) {
  .debtors-container {
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

  .debtors-list-header {
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

  /* Dialog full width */
  :deep(.v-dialog) {
    margin: 8px !important;
  }

  :deep(.v-dialog > .v-overlay__content) {
    max-width: calc(100% - 16px) !important;
  }

  .dialog-header {
    font-size: 16px !important;
    padding: 14px 16px !important;
  }

  .dialog-header i {
    font-size: 20px;
  }

  .dialog-content {
    padding: 16px !important;
    max-height: 55vh;
  }

  .dialog-actions {
    padding: 12px 16px !important;
    flex-wrap: wrap;
    gap: 8px;
  }

  .dialog-btn {
    padding: 8px 14px;
    font-size: 12px;
    flex: 1;
    min-width: 100px;
    justify-content: center;
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

  /* View dialog */
  .debtor-profile {
    padding: 16px 0;
  }

  .debtor-name {
    font-size: 20px;
  }

  .debtor-stat-card {
    padding: 12px;
  }

  .stat-icon-small {
    width: 36px;
    height: 36px;
    font-size: 18px;
  }

  .stat-content .stat-value {
    font-size: 13px;
  }

  .payment-progress-section {
    padding: 12px;
  }

  .info-section {
    padding: 12px;
  }

  .section-title {
    font-size: 13px;
  }

  .due-date-display {
    flex-direction: column;
    align-items: flex-start;
    gap: 8px;
  }

  .due-date-value {
    font-size: 14px;
  }

  .notes-text {
    font-size: 13px;
  }

  /* Notes cell */
  .notes-cell {
    max-width: 120px;
    font-size: 12px;
  }
}

@media (max-width: 360px) {
  .debtors-container {
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

  .debtors-list-header {
    padding: 10px 12px;
  }

  .list-header-title {
    font-size: 13px;
  }

  .dialog-header {
    font-size: 15px !important;
    padding: 12px 14px !important;
  }

  .dialog-content {
    padding: 12px !important;
  }

  .dialog-actions {
    padding: 10px 12px !important;
  }

  .dialog-btn {
    padding: 6px 12px;
    font-size: 11px;
  }
}
</style>

<!-- Global styles for v-select dropdown menu -->
<style>
.v-overlay__content .v-list {
  background: linear-gradient(135deg, #0a3d42 0%, #052428 100%) !important;
  border: 1px solid rgba(6, 182, 212, 0.3) !important;
  border-radius: 12px !important;
}

.v-overlay__content .v-list-item {
  color: rgba(255, 255, 255, 0.9) !important;
}

.v-overlay__content .v-list-item:hover {
  background: rgba(6, 182, 212, 0.2) !important;
}

.v-overlay__content .v-list-item--active {
  background: rgba(6, 182, 212, 0.3) !important;
  color: #06b6d4 !important;
}

.v-overlay__content .v-list-item-title {
  color: rgba(255, 255, 255, 0.9) !important;
}
</style>
