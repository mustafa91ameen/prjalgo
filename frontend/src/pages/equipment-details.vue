<template>
  <div class="equipment-details-container">
    <!-- Page Header -->
    <PageHeader
      title="تفاصيل الآليات"
      subtitle="إدارة المعدات والآليات المستخدمة"
      badge="الآليات"
      badgeType="warning"
      class="equipment-details-header"
    >
      <template #actions>
        <button class="page-action-btn secondary" @click="$router.back()">
          <i class="mdi mdi-arrow-right"></i>
          رجوع
        </button>
        <button v-if="canCreate" class="page-action-btn primary" @click="showAddEquipmentModal = true">
          <i class="mdi mdi-plus"></i>
          إضافة آلية
        </button>
      </template>
    </PageHeader>

    <!-- Summary Cards -->
    <div class="summary-section">
      <div class="summary-card equipment-count-summary">
        <div class="summary-icon">
          <i class="mdi mdi-excavator"></i>
        </div>
        <div class="summary-info">
          <span class="summary-label">إجمالي الآليات</span>
          <span class="summary-value">{{ equipment.length }}</span>
        </div>
      </div>

      <div class="summary-card quantity-summary">
        <div class="summary-icon">
          <i class="mdi mdi-counter"></i>
        </div>
        <div class="summary-info">
          <span class="summary-label">إجمالي الكمية</span>
          <span class="summary-value">{{ totalQuantity }}</span>
        </div>
      </div>

      <div class="summary-card cost-summary">
        <div class="summary-icon">
          <i class="mdi mdi-cash"></i>
        </div>
        <div class="summary-info">
          <span class="summary-label">إجمالي تكلفة التشغيل</span>
          <span class="summary-value">{{ formatCurrency(totalCost) }}</span>
        </div>
      </div>
    </div>

    <!-- Equipment Section Header -->
    <div class="section-header">
      <div class="section-header-content">
        <i class="mdi mdi-excavator"></i>
        <span>جدول الآليات والمعدات</span>
      </div>
      <div class="section-header-stats">
        <div class="header-stat">
          <span class="header-stat-value">{{ equipment.length }}</span>
          <span class="header-stat-label">آلية</span>
        </div>
        <div class="header-stat">
          <span class="header-stat-value">{{ formatCurrency(totalCost) }}</span>
          <span class="header-stat-label">إجمالي التكلفة</span>
        </div>
      </div>
    </div>

    <!-- Equipment Table -->
    <v-card class="data-card">
      <v-data-table
        :headers="equipmentHeaders"
        :items="equipment"
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
          <v-btn v-if="canDelete" size="small" color="error" @click="deleteEquipment(item)">
            <i class="mdi mdi-delete"></i>
          </v-btn>
        </template>
      </v-data-table>
    </v-card>

    <!-- Add Equipment Modal -->
    <div v-if="showAddEquipmentModal" class="modal-overlay" @click.self="showAddEquipmentModal = false">
      <div class="modal-container">
        <div class="modal-header">
          <h3>
            <i class="mdi mdi-excavator"></i>
            إضافة آلية جديدة
          </h3>
          <button class="modal-close" @click="showAddEquipmentModal = false">
            <i class="mdi mdi-close"></i>
          </button>
        </div>
        
        <div class="modal-body">
          <div class="form-group">
            <label>اسم الآلية</label>
            <input type="text" v-model="newEquipment.name" placeholder="أدخل اسم الآلية" />
          </div>
          <div class="form-group">
            <label>الكمية</label>
            <input type="number" v-model="newEquipment.quantity" placeholder="أدخل الكمية" />
          </div>
          <div class="form-group">
            <label>التكلفة (د.ع)</label>
            <input type="number" v-model="newEquipment.cost" placeholder="أدخل التكلفة" />
          </div>
          <div class="form-group">
            <label>ملاحظات</label>
            <textarea v-model="newEquipment.notes" placeholder="ملاحظات إضافية" rows="2"></textarea>
          </div>
        </div>

        <div class="modal-footer">
          <button class="btn-cancel" @click="showAddEquipmentModal = false">إلغاء</button>
          <button class="btn-save" @click="addEquipment">حفظ الآلية</button>
        </div>
      </div>
    </div>

  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { useRoute } from 'vue-router'
import PageHeader from '../components/PageHeader.vue'
import { listEquipmentByWorkDay, createEquipment, deleteEquipment as apiDeleteEquipment } from '@/api/materials'
import { usePermissions } from '@/composables/usePermissions'
import { useToast } from '@/composables/useToast'

const route = useRoute()
const { success, error: showError } = useToast()
const { canCreate, canUpdate, canDelete } = usePermissions('/workdayEquipment')

const loading = ref(false)
const showAddEquipmentModal = ref(false)

// Workday ID from route
const workdayId = ref(route.query.workday_id || null)

// Equipment data
const equipment = ref([])

// Fetch equipment from API
const fetchEquipment = async () => {
  if (!workdayId.value) return

  loading.value = true
  try {
    const data = await listEquipmentByWorkDay(workdayId.value)
    equipment.value = (data || []).map(e => ({
      id: e.id,
      name: e.equipmentName,
      quantity: e.quantity || 0,
      cost: e.cost || 0,
      notes: e.notes || ''
    }))
  } catch (error) {
    console.error('Error fetching equipment:', error)
    showError('حدث خطأ في جلب بيانات الآليات')
  } finally {
    loading.value = false
  }
}

const equipmentHeaders = [
  { title: '#', key: 'index', align: 'center' },
  { title: 'اسم الآلية', key: 'name', align: 'start' },
  { title: 'الكمية', key: 'quantity', align: 'center' },
  { title: 'التكلفة', key: 'cost', align: 'center' },
  { title: 'ملاحظات', key: 'notes', align: 'center' },
  { title: 'الإجراءات', key: 'actions', align: 'center', sortable: false }
]

// New equipment form
const newEquipment = ref({
  name: '',
  quantity: '',
  cost: '',
  notes: ''
})

// Computed totals
const totalCost = computed(() => {
  return equipment.value.reduce((sum, item) => sum + item.cost, 0)
})

const totalQuantity = computed(() => {
  return equipment.value.reduce((sum, item) => sum + item.quantity, 0)
})

// Methods
const formatCurrency = (amount) => {
  return new Intl.NumberFormat('en-US').format(amount) + ' د.ع'
}

const addEquipment = async () => {
  if (!newEquipment.value.name || !newEquipment.value.cost || !workdayId.value) return
  loading.value = true
  try {
    const equipmentData = {
      workDayId: Number(workdayId.value),
      equipmentName: newEquipment.value.name,
      quantity: Number(newEquipment.value.quantity) || 1,
      cost: Number(newEquipment.value.cost),
      notes: newEquipment.value.notes || null
    }

    await createEquipment(equipmentData)
    success('تم إضافة الآلية بنجاح')
    newEquipment.value = { name: '', quantity: '', cost: '', notes: '' }
    showAddEquipmentModal.value = false
    await fetchEquipment()
  } catch (error) {
    console.error('Error adding equipment:', error)
    showError('حدث خطأ في إضافة الآلية')
  } finally {
    loading.value = false
  }
}

const deleteEquipment = async (item) => {
  if (!confirm('هل أنت متأكد من حذف هذه الآلية؟')) return
  loading.value = true
  try {
    await apiDeleteEquipment(item.id)
    success('تم حذف الآلية بنجاح')
    await fetchEquipment()
  } catch (error) {
    console.error('Error deleting equipment:', error)
    showError('حدث خطأ في حذف الآلية')
  } finally {
    loading.value = false
  }
}

// Fetch data on mount
onMounted(() => {
  fetchEquipment()
})
</script>

<style scoped>
.equipment-details-container {
  padding: 32px;
  max-width: 1600px;
  margin: 0 auto;
}

/* Header */
.equipment-details-header {
  background: linear-gradient(135deg, #018790 0%, #005461 100%) !important;
}

.equipment-details-header::before {
  background: linear-gradient(135deg, #14b8a6 0%, #0d9488 50%, #14b8a6 100%) !important;
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
  border: 2px solid rgba(20, 184, 166, 0.3);
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
  color: #14b8a6;
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
  background: linear-gradient(135deg, #14b8a6 0%, #06b6d4 100%);
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
  border: 2px solid rgba(20, 184, 166, 0.3) !important;
  overflow: hidden;
}

.cost-value {
  font-weight: 700;
  color: #14b8a6;
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
  color: #14b8a6;
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
  border: 2px solid rgba(20, 184, 166, 0.3);
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

.equipment-count-summary .summary-icon {
  background: linear-gradient(135deg, #14b8a6 0%, #0d9488 100%);
}

.quantity-summary .summary-icon {
  background: linear-gradient(135deg, #06b6d4 0%, #0891b2 100%);
}

.cost-summary .summary-icon {
  background: linear-gradient(135deg, #f59e0b 0%, #d97706 100%);
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
  background: linear-gradient(135deg, #14b8a6 0%, #06b6d4 100%);
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
  border: 2px solid rgba(20, 184, 166, 0.4);
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
  border-bottom: 1px solid rgba(20, 184, 166, 0.2);
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
  color: #14b8a6;
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
  border: 1px solid rgba(20, 184, 166, 0.3);
  border-radius: 10px;
  padding: 12px 16px;
  color: #fff;
  font-size: 14px;
  outline: none;
}

.form-group input:focus,
.form-group select:focus,
.form-group textarea:focus {
  border-color: #14b8a6;
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
  border-top: 1px solid rgba(20, 184, 166, 0.2);
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
  background: linear-gradient(135deg, #14b8a6 0%, #0d9488 100%);
  color: white;
  box-shadow: 0 4px 15px rgba(20, 184, 166, 0.3);
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
