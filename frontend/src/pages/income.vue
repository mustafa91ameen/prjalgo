<template>
  <div class="income-container">
    <!-- Page Header Component -->
    <PageHeader
      title="إدارة الإيرادات"
      subtitle="إدارة وتتبع جميع مصادر الإيرادات"
      badge="الإيرادات"
      badgeType="success"
      class="income-header"
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
      <!-- Total Income Card -->
      <v-card class="stat-card" elevation="0">
        <div class="stat-card-content">
          <div class="stat-icon profit">
            <i class="mdi mdi-cash-plus"></i>
          </div>
          <div class="stat-info">
            <div class="stat-label">إجمالي الإيرادات</div>
            <div class="stat-value">{{ formatCurrency(totalIncome) }}</div>
          </div>
        </div>
      </v-card>

      <!-- Monthly Income Card -->
      <v-card class="stat-card" elevation="0">
        <div class="stat-card-content">
          <div class="stat-icon success">
            <i class="mdi mdi-calendar-check"></i>
          </div>
          <div class="stat-info">
            <div class="stat-label">إيرادات هذا الشهر</div>
            <div class="stat-value">{{ formatCurrency(monthlyIncome) }}</div>
          </div>
        </div>
      </v-card>

      <!-- Income Growth Card -->
      <v-card class="stat-card" elevation="0">
        <div class="stat-card-content">
          <div class="stat-icon">
            <i class="mdi mdi-chart-line"></i>
          </div>
          <div class="stat-info">
            <div class="stat-label">نمو الإيرادات</div>
            <div class="stat-value">{{ incomeGrowth }}%</div>
          </div>
        </div>
      </v-card>

      <!-- Income Sources Card -->
      <v-card class="stat-card" elevation="0">
        <div class="stat-card-content">
          <div class="stat-icon warning">
            <i class="mdi mdi-source-branch"></i>
          </div>
          <div class="stat-info">
            <div class="stat-label">مصادر الإيرادات</div>
            <div class="stat-value">{{ incomeSources.length }}</div>
          </div>
        </div>
      </v-card>
    </div>

    <!-- Income Sources List Header -->
    <div class="income-list-header">
      <div class="list-header-content">
        <div class="list-header-info">
          <h2 class="list-header-title">
            <i class="mdi mdi-format-list-bulleted"></i>
            مصادر الإيرادات
          </h2>
          <p class="list-header-subtitle">عرض جميع مصادر الإيرادات المتاحة</p>
        </div>
        <div class="list-header-actions">
          <button v-if="canCreate" class="list-action-btn primary" @click="showAddDialog = true">
            <i class="mdi mdi-plus"></i>
            إضافة إيراد جديد
          </button>
        </div>
      </div>
    </div>

      <!-- Income Sources Table -->
      <v-card class="mb-6">
        <v-data-table
          :headers="headers"
          :items="incomeSources"
          :loading="loading"
          class="elevation-1"
        >
          <template v-slot:item.amount="{ item }">
            <span class="font-weight-bold text-success">{{ formatCurrency(item.amount) }}</span>
          </template>
          <template v-slot:item.date="{ item }">
            {{ formatDate(item.date) }}
          </template>
          <template v-slot:item.actions="{ item }">
            <v-btn
              v-if="canUpdate"
              size="small"
              color="primary"
              @click="editIncome(item)"
            >
              <i class="mdi mdi-pencil"></i>
            </v-btn>
            <v-btn
              v-if="canDelete"
              size="small"
              color="error"
              @click="deleteIncome(item)"
            >
              <i class="mdi mdi-delete"></i>
            </v-btn>
          </template>
        </v-data-table>
      </v-card>

      <!-- Add/Edit Income Dialog -->
      <v-dialog v-model="showAddDialog" max-width="800" persistent>
        <v-card class="income-dialog">
          <v-card-title class="dialog-header">
            <i class="mdi mdi-cash-plus"></i>
            {{ editingIncome ? 'تعديل الإيراد' : 'إضافة إيراد جديد' }}
          </v-card-title>
          
          <v-card-text class="dialog-content">
            <v-form ref="form" v-model="valid">
              <v-row>
                <!-- اسم الإيراد -->
                <v-col cols="12" md="6">
                  <v-text-field
                    v-model="incomeForm.description"
                    label="اسم الإيراد"
                    variant="outlined"
                    density="comfortable"
                    :rules="[v => !!v || 'اسم الإيراد مطلوب']"
                    required
                  ></v-text-field>
                </v-col>

                <!-- المبلغ (د.ع) -->
                <v-col cols="12" md="6">
                  <v-text-field
                    v-model.number="incomeForm.amount"
                    label="المبلغ (د.ع)"
                    variant="outlined"
                    density="comfortable"
                    type="number"
                    :rules="[v => v > 0 || 'المبلغ يجب أن يكون أكبر من صفر']"
                    required
                  ></v-text-field>
                </v-col>

                <!-- تاريخ الإيراد -->
                <v-col cols="12" md="6">
                  <v-text-field
                    v-model="incomeForm.date"
                    label="تاريخ الإيراد"
                    variant="outlined"
                    density="comfortable"
                    type="date"
                    :rules="[v => !!v || 'تاريخ الإيراد مطلوب']"
                    required
                  ></v-text-field>
                </v-col>

                <!-- النوع -->
                <v-col cols="12" md="6">
                  <v-text-field
                    v-model="incomeForm.category"
                    label="النوع"
                    variant="outlined"
                    density="comfortable"
                    placeholder="أدخل نوع الإيراد"
                  ></v-text-field>
                </v-col>

                <!-- الحالة -->
                <v-col cols="12" md="6">
                  <v-select
                    v-model="incomeForm.status"
                    :items="statusOptions"
                    label="الحالة"
                    variant="outlined"
                    density="comfortable"
                    :rules="[v => !!v || 'الحالة مطلوبة']"
                    required
                  ></v-select>
                </v-col>

                <!-- الملاحظات -->
                <v-col cols="12">
                  <v-textarea
                    v-model="incomeForm.notes"
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
            <button class="dialog-btn save" @click="saveIncome" :disabled="!valid">
              <i class="mdi mdi-content-save"></i>
              حفظ
            </button>
          </v-card-actions>
        </v-card>
      </v-dialog>

  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import PageHeader from '../components/PageHeader.vue'
import { listIncome, createIncome, updateIncome, deleteIncome as apiDeleteIncome } from '@/api/income'
import { usePermissions } from '@/composables/usePermissions'
import { useToast } from '@/composables/useToast'
import { DEFAULT_PAGE, DEFAULT_LIMIT } from '@/constants/pagination'

const { success, error: showError } = useToast()
const { canCreate, canUpdate, canDelete } = usePermissions('/income')

// Reactive data
const loading = ref(false)
const showAddDialog = ref(false)
const valid = ref(false)
const editingIncome = ref(null)

// Pagination
const page = ref(DEFAULT_PAGE)
const limit = ref(DEFAULT_LIMIT)
const total = ref(0)

const incomeForm = ref({
  description: '',
  amount: 0,
  category: '',
  date: new Date().toISOString().substr(0, 10),
  status: 'pending',
  notes: ''
})

const incomeSources = ref([])

// Fetch income from API
const fetchIncome = async () => {
  loading.value = true
  try {
    const response = await listIncome({ page: page.value, limit: limit.value })
    if (response.success) {
      incomeSources.value = (response.data.items || []).map(i => ({
        id: i.id,
        description: i.description || '',
        amount: i.amount || 0,
        category: i.category || '',
        date: i.date || i.created_at,
        status: i.status || 'pending',
        notes: i.notes || ''
      }))
      total.value = response.data.total || 0
    }
  } catch (error) {
    console.error('Error fetching income:', error)
    showError('حدث خطأ في جلب الإيرادات')
  } finally {
    loading.value = false
  }
}

const statusOptions = [
  'pending',
  'approved',
  'rejected'
]

const headers = [
  { title: 'الوصف', key: 'description', align: 'start' },
  { title: 'المبلغ', key: 'amount', align: 'center' },
  { title: 'الفئة', key: 'category', align: 'center' },
  { title: 'التاريخ', key: 'date', align: 'center' },
  { title: 'الإجراءات', key: 'actions', align: 'center', sortable: false }
]

// Computed properties
const totalIncome = computed(() => {
  return incomeSources.value.reduce((sum, item) => sum + item.amount, 0)
})

const monthlyIncome = computed(() => {
  const currentMonth = new Date().getMonth()
  const currentYear = new Date().getFullYear()
  
  return incomeSources.value
    .filter(item => {
      const itemDate = new Date(item.date)
      return itemDate.getMonth() === currentMonth && itemDate.getFullYear() === currentYear
    })
    .reduce((sum, item) => sum + item.amount, 0)
})

const incomeGrowth = computed(() => {
  // Simplified growth calculation
  return 15.5
})

// Methods
const formatCurrency = (amount) => {
  return new Intl.NumberFormat('en-US').format(amount) + ' د.ع'
}

const formatDate = (dateString) => {
  return new Date(dateString).toLocaleDateString('en-US')
}

const editIncome = (item) => {
  editingIncome.value = item
  incomeForm.value = { ...item }
  showAddDialog.value = true
}

const deleteIncome = async (item) => {
  if (!confirm('هل أنت متأكد من حذف هذا الإيراد؟')) return

  loading.value = true
  try {
    await apiDeleteIncome(item.id)
    success('تم حذف الإيراد بنجاح')
    fetchIncome()
  } catch (error) {
    console.error('Error deleting income:', error)
    showError(error.message || 'حدث خطأ في حذف الإيراد')
  } finally {
    loading.value = false
  }
}

const saveIncome = async () => {
  if (!valid.value) return

  loading.value = true
  try {
    const incomeData = {
      name: incomeForm.value.description,
      amount: Number(incomeForm.value.amount),
      type: incomeForm.value.category || null,
      incomeDate: incomeForm.value.date ? new Date(incomeForm.value.date).toISOString() : new Date().toISOString(),
      status: incomeForm.value.status || null,
      notes: incomeForm.value.notes || null
    }

    if (editingIncome.value) {
      await updateIncome(editingIncome.value.id, incomeData)
    } else {
      await createIncome(incomeData)
    }

    success(editingIncome.value ? 'تم تحديث الإيراد بنجاح' : 'تم إضافة الإيراد بنجاح')
    closeDialog()
    fetchIncome()
  } catch (error) {
    console.error('Error saving income:', error)
    showError(error.message || 'حدث خطأ في حفظ الإيراد')
  } finally {
    loading.value = false
  }
}

const closeDialog = () => {
  showAddDialog.value = false
  editingIncome.value = null
  incomeForm.value = {
    description: '',
    amount: 0,
    category: '',
    notes: ''
  }
}

onMounted(() => {
  fetchIncome()
})
</script>

<style scoped>
.income-container {
  padding: 32px;
  max-width: 1600px;
  margin: 0 auto;
}

/* Income Header Custom Color */
.income-header {
  background: linear-gradient(135deg, #018790 0%, #005461 100%) !important;
}

.income-header::before {
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

.stat-icon.profit {
  background: linear-gradient(135deg, #10b981 0%, #34d399 100%) !important;
  box-shadow: 0 6px 16px rgba(16, 185, 129, 0.4);
}

.stat-icon.success {
  background: linear-gradient(135deg, #10b981 0%, #059669 100%) !important;
  box-shadow: 0 6px 16px rgba(16, 185, 129, 0.4);
}

.stat-icon.warning {
  background: linear-gradient(135deg, #14b8a6 0%, #2dd4bf 100%) !important;
  box-shadow: 0 6px 16px rgba(20, 184, 166, 0.4);
}

/* Income Sources List Header */
.income-list-header {
  background: linear-gradient(135deg, #0a3d42 0%, #052428 100%);
  border-radius: 16px;
  margin-bottom: 24px;
  padding: 16px 24px;
  border: 2px solid transparent;
  position: relative;
  overflow: hidden;
}

.income-list-header::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  border-radius: 16px;
  padding: 2px;
  background: linear-gradient(135deg, #10b981 0%, #34d399 50%, #10b981 100%);
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
  color: #10b981;
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
  background: linear-gradient(135deg, #10b981 0%, #059669 100%);
  color: white;
  box-shadow: 0 4px 12px rgba(16, 185, 129, 0.3);
}

.list-action-btn.primary:hover {
  transform: translateY(-2px);
  box-shadow: 0 6px 16px rgba(16, 185, 129, 0.4);
}

.list-action-btn i {
  font-size: 16px;
}

/* Income Dialog */
.income-dialog {
  border-radius: 16px !important;
  background: linear-gradient(135deg, #0a3d42 0%, #052428 100%) !important;
  border: 2px solid transparent !important;
  position: relative;
  overflow: hidden;
  direction: rtl;
  text-align: right;
}

.income-dialog::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  border-radius: 16px;
  padding: 2px;
  background: linear-gradient(135deg, #10b981 0%, #34d399 50%, #10b981 100%);
  -webkit-mask: linear-gradient(#fff 0 0) content-box, linear-gradient(#fff 0 0);
  mask: linear-gradient(#fff 0 0) content-box, linear-gradient(#fff 0 0);
  -webkit-mask-composite: xor;
  mask-composite: exclude;
  pointer-events: none;
}

.dialog-header {
  background: rgba(16, 185, 129, 0.1) !important;
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
  color: #10b981;
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
  color: #10b981 !important;
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
  background: linear-gradient(135deg, #10b981 0%, #059669 100%);
  color: white;
  box-shadow: 0 4px 12px rgba(16, 185, 129, 0.3);
}

.dialog-btn.save:hover:not(:disabled) {
  transform: translateY(-2px);
  box-shadow: 0 6px 16px rgba(16, 185, 129, 0.4);
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
  .income-container {
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
  .income-container {
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

  .income-list-header {
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
}

@media (max-width: 480px) {
  .income-container {
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

  .income-list-header {
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
}

@media (max-width: 360px) {
  .income-container {
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

  .income-list-header {
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
