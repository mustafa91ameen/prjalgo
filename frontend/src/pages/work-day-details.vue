<template>
  <div class="work-day-details-container">
    <!-- Page Header -->
    <PageHeader
      title="تفاصيل يوم العمل"
      subtitle="عرض تفاصيل يوم العمل الكاملة"
      badge="تفاصيل"
      badgeType="info"
      class="details-header"
    >
      <template #actions>
        <button class="page-action-btn secondary" @click="$router.back()">
          <i class="mdi mdi-arrow-right"></i>
          رجوع
        </button>
        <button class="page-action-btn primary">
          <i class="mdi mdi-pencil"></i>
          تعديل
        </button>
      </template>
    </PageHeader>

    <!-- Management Sections Cards -->
    <div class="management-sections-grid">
      <!-- المواد والمصاريف اليومية -->
      <div class="section-card materials-card" :class="{ 'card-disabled': !materialsEnabled }">
        <div class="card-toggle">
          <label class="toggle-switch">
            <input type="checkbox" v-model="materialsEnabled">
            <span class="toggle-slider"></span>
          </label>
        </div>
        <div class="section-card-content">
          <div class="section-icon">
            <i class="mdi mdi-package-variant"></i>
          </div>
          <div class="section-info">
            <h3 class="section-title">المواد والمصاريف اليومية</h3>
            <p class="section-subtitle">إدارة المواد والمصروفات اليومية للمشروع</p>
          </div>
          <div class="section-stats">
            <div class="stat-item">
              <span class="stat-number">{{ materials.length }}</span>
              <span class="stat-label">مادة</span>
            </div>
          </div>
          <button class="section-btn" @click="goToMaterialsExpenses" :disabled="!materialsEnabled">
            <i class="mdi mdi-eye"></i>
            عرض التفاصيل
          </button>
        </div>
      </div>

      <!-- الأيدي العاملة -->
      <div class="section-card labor-card" :class="{ 'card-disabled': !laborEnabled }">
        <div class="card-toggle">
          <label class="toggle-switch">
            <input type="checkbox" v-model="laborEnabled">
            <span class="toggle-slider"></span>
          </label>
        </div>
        <div class="section-card-content">
          <div class="section-icon">
            <i class="mdi mdi-account-hard-hat"></i>
          </div>
          <div class="section-info">
            <h3 class="section-title">الأيدي العاملة</h3>
            <p class="section-subtitle">إدارة العمال والأجور اليومية</p>
          </div>
          <div class="section-stats">
            <div class="stat-item">
              <span class="stat-number">{{ labor.length }}</span>
              <span class="stat-label">عامل</span>
            </div>
          </div>
          <button class="section-btn" @click="goToLaborDetails" :disabled="!laborEnabled">
            <i class="mdi mdi-eye"></i>
            عرض التفاصيل
          </button>
        </div>
      </div>

      <!-- الآليات -->
      <div class="section-card equipment-card" :class="{ 'card-disabled': !equipmentEnabled }">
        <div class="card-toggle">
          <label class="toggle-switch">
            <input type="checkbox" v-model="equipmentEnabled">
            <span class="toggle-slider"></span>
          </label>
        </div>
        <div class="section-card-content">
          <div class="section-icon">
            <i class="mdi mdi-excavator"></i>
          </div>
          <div class="section-info">
            <h3 class="section-title">الآليات</h3>
            <p class="section-subtitle">إدارة المعدات والآليات المستخدمة</p>
          </div>
          <div class="section-stats">
            <div class="stat-item">
              <span class="stat-number">{{ equipment.length }}</span>
              <span class="stat-label">آلية</span>
            </div>
          </div>
          <button class="section-btn" @click="goToEquipmentDetails" :disabled="!equipmentEnabled">
            <i class="mdi mdi-eye"></i>
            عرض التفاصيل
          </button>
        </div>
      </div>
    </div>

    <!-- Sections Title Bar -->
    <div class="sections-title-bar">
      <div class="sections-title-content">
        <i class="mdi mdi-view-grid-outline"></i>
        <span>أقسام إدارة يوم العمل</span>
      </div>
    </div>

    <!-- Work Day Info Card -->
    <v-card class="details-card">
      <v-card-title class="card-header">
        <i class="mdi mdi-information"></i>
        معلومات يوم العمل
      </v-card-title>
      <v-card-text class="card-body">
        <div class="details-grid">
          <!-- مكان العمل -->
          <div class="detail-item">
            <div class="detail-label">
              <i class="mdi mdi-map-marker"></i>
              مكان العمل
            </div>
            <div class="detail-value">{{ workdayDetails.location || '-' }}</div>
          </div>

          <!-- رقم الاستمارة -->
          <div class="detail-item">
            <div class="detail-label">
              <i class="mdi mdi-file-document"></i>
              رقم الاستمارة
            </div>
            <div class="detail-value">{{ workdayDetails.formNumber || '-' }}</div>
          </div>

          <!-- التاريخ -->
          <div class="detail-item">
            <div class="detail-label">
              <i class="mdi mdi-calendar-today"></i>
              التاريخ
            </div>
            <div class="detail-value">{{ workdayDetails.date || '-' }}</div>
          </div>

          <!-- اليوم -->
          <div class="detail-item">
            <div class="detail-label">
              <i class="mdi mdi-calendar"></i>
              اليوم
            </div>
            <div class="detail-value">{{ workdayDetails.day || '-' }}</div>
          </div>

          <!-- فترة العمل من -->
          <div class="detail-item">
            <div class="detail-label">
              <i class="mdi mdi-clock-start"></i>
              فترة العمل من
            </div>
            <div class="detail-value">{{ workdayDetails.workPeriodFrom || '-' }}</div>
          </div>

          <!-- فترة العمل إلى -->
          <div class="detail-item">
            <div class="detail-label">
              <i class="mdi mdi-clock-end"></i>
              فترة العمل إلى
            </div>
            <div class="detail-value">{{ workdayDetails.workPeriodTo || '-' }}</div>
          </div>

          <!-- عدد العمال -->
          <div class="detail-item">
            <div class="detail-label">
              <i class="mdi mdi-account-group"></i>
              عدد العمال
            </div>
            <div class="detail-value">{{ labor.length }} عامل</div>
          </div>

          <!-- التكلفة الإجمالية -->
          <div class="detail-item">
            <div class="detail-label">
              <i class="mdi mdi-cash"></i>
              التكلفة الإجمالية
            </div>
            <div class="detail-value">{{ formatCurrency(workdayDetails.totalCost) }}</div>
          </div>

          <!-- الحالة -->
          <div class="detail-item">
            <div class="detail-label">
              <i class="mdi mdi-check-circle"></i>
              الحالة
            </div>
            <div class="detail-value">
              <v-chip :color="getStatusColor(workdayDetails.status)" size="small">{{ getStatusText(workdayDetails.status) }}</v-chip>
            </div>
          </div>

          <!-- الوصف -->
          <div class="detail-item full-width">
            <div class="detail-label">
              <i class="mdi mdi-text"></i>
              الوصف
            </div>
            <div class="detail-value">{{ workdayDetails.description || '-' }}</div>
          </div>
        </div>
      </v-card-text>
    </v-card>

    <!-- Workers List Card -->
    <v-card class="workers-card mt-6">
      <v-card-title class="card-header">
        <i class="mdi mdi-account-multiple"></i>
        قائمة العمال
      </v-card-title>
      <v-card-text class="card-body">
        <v-data-table
          :headers="laborHeaders"
          :items="labor"
          class="elevation-0"
        >
          <template v-slot:item.index="{ index }">
            {{ index + 1 }}
          </template>
          <template v-slot:item.salary="{ item }">
            {{ formatCurrency(item.salary) }}
          </template>
        </v-data-table>
      </v-card-text>
    </v-card>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import PageHeader from '../components/PageHeader.vue'
import { getWorkday } from '@/api/workdays'
import { listMaterialsByWorkDay, listLaborByWorkDay, listEquipmentByWorkDay } from '@/api/materials'
import { useToast } from '@/composables/useToast'

const { error: showError } = useToast()
const router = useRouter()
const route = useRoute()
const loading = ref(false)

// Work day ID from route
const workdayId = ref(route.query.id || null)

// Work day details
const workdayDetails = ref({
  location: '',
  formNumber: '',
  date: '',
  day: '',
  workPeriodFrom: '',
  workPeriodTo: '',
  workers: 0,
  totalCost: 0,
  status: 'pending',
  description: ''
})

// Fetch work day details
const fetchWorkdayDetails = async () => {
  if (!workdayId.value) return

  loading.value = true
  try {
    const response = await getWorkday(workdayId.value)
    if (response.success && response.data) {
      const w = response.data
      // Format date for display in English
      const workDate = w.workDate ? new Date(w.workDate) : null
      const dayNames = ['Sunday', 'Monday', 'Tuesday', 'Wednesday', 'Thursday', 'Friday', 'Saturday']

      workdayDetails.value = {
        location: w.notes || '',
        formNumber: w.id?.toString() || '',
        date: workDate ? workDate.toLocaleDateString('en-US', { year: 'numeric', month: '2-digit', day: '2-digit' }) : '',
        day: workDate ? dayNames[workDate.getDay()] : '',
        workPeriodFrom: '',
        workPeriodTo: '',
        workers: labor.value.length,
        totalCost: w.totalCost || 0,
        status: w.status || 'pending',
        description: w.description || ''
      }
    }
  } catch (error) {
    console.error('Error fetching workday details:', error)
    showError('حدث خطأ في جلب تفاصيل يوم العمل')
  } finally {
    loading.value = false
  }
}

// Fetch materials for this workday
const fetchMaterials = async () => {
  if (!workdayId.value) return
  try {
    const data = await listMaterialsByWorkDay(workdayId.value)
    materials.value = (data || []).map(m => ({
      id: m.id,
      name: m.materialName,
      quantity: m.quantity,
      cost: m.cost || 0
    }))
  } catch (error) {
    console.error('Error fetching materials:', error)
  }
}

// Fetch labor for this workday
const fetchLabor = async () => {
  if (!workdayId.value) return
  try {
    const data = await listLaborByWorkDay(workdayId.value)
    labor.value = (data || []).map(l => ({
      id: l.id,
      name: l.workerName,
      role: l.jobTitle || '',
      salary: l.cost || 0
    }))
  } catch (error) {
    console.error('Error fetching labor:', error)
  }
}

// Fetch equipment for this workday
const fetchEquipment = async () => {
  if (!workdayId.value) return
  try {
    const data = await listEquipmentByWorkDay(workdayId.value)
    equipment.value = (data || []).map(e => ({
      id: e.id,
      name: e.equipmentName,
      hours: e.quantity || 0,
      cost: e.cost || 0
    }))
  } catch (error) {
    console.error('Error fetching equipment:', error)
  }
}

// Initialize
onMounted(() => {
  fetchWorkdayDetails()
  fetchMaterials()
  fetchLabor()
  fetchEquipment()
})

// Toggle states for cards
const materialsEnabled = ref(true)
const laborEnabled = ref(true)
const equipmentEnabled = ref(true)

const goToMaterialsExpenses = () => {
  router.push({ path: '/materials-expenses', query: { workday_id: workdayId.value } })
}

const goToLaborDetails = () => {
  router.push({ path: '/labor-details', query: { workday_id: workdayId.value } })
}

const goToEquipmentDetails = () => {
  router.push({ path: '/equipment-details', query: { workday_id: workdayId.value } })
}

// Materials data (fetched from API)
const materials = ref([])

// Labor data (fetched from API)
const labor = ref([])

const laborHeaders = [
  { title: '#', key: 'index', align: 'center' },
  { title: 'الاسم', key: 'name', align: 'start' },
  { title: 'الدور', key: 'role', align: 'center' },
  { title: 'الأجر اليومي', key: 'salary', align: 'center' }
]

// Equipment data (fetched from API)
const equipment = ref([])

// Helper functions
const formatCurrency = (amount) => {
  return new Intl.NumberFormat('en-US').format(amount || 0) + ' د.ع'
}

const getStatusColor = (status) => {
  const colors = {
    completed: 'success',
    in_progress: 'warning',
    pending: 'info'
  }
  return colors[status] || 'grey'
}

const getStatusText = (status) => {
  const texts = {
    completed: 'مكتمل',
    in_progress: 'قيد التنفيذ',
    pending: 'معلق'
  }
  return texts[status] || status
}

</script>

<style scoped>
.work-day-details-container {
  padding: 32px;
  max-width: 1400px;
  margin: 0 auto;
}

/* Details Header Custom Color */
.details-header {
  background: linear-gradient(135deg, #018790 0%, #005461 100%) !important;
}

.details-header::before {
  background: linear-gradient(135deg, #06b6d4 0%, #10b981 50%, #06b6d4 100%) !important;
}

/* Details Card */
.details-card,
.workers-card {
  border-radius: 16px !important;
  background: linear-gradient(135deg, #0a3d42 0%, #052428 100%) !important;
  border: 2px solid transparent !important;
  position: relative;
}

.details-card::before,
.workers-card::before {
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

.card-header {
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

.card-header i {
  color: #06b6d4;
  font-size: 28px;
}

.card-body {
  padding: 32px !important;
  position: relative;
  z-index: 1;
}

/* Details Grid */
.details-grid {
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  gap: 24px;
}

.detail-item {
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.detail-item.full-width {
  grid-column: 1 / -1;
}

.detail-label {
  display: flex;
  align-items: center;
  gap: 8px;
  font-size: 13px;
  color: rgba(255, 255, 255, 0.6);
  font-weight: 600;
  text-transform: uppercase;
  letter-spacing: 0.5px;
}

.detail-label i {
  color: #06b6d4;
  font-size: 18px;
}

.detail-value {
  font-size: 16px;
  color: rgba(255, 255, 255, 0.95);
  font-weight: 500;
  padding: 12px 16px;
  background: rgba(255, 255, 255, 0.05);
  border-radius: 12px;
  border: 1px solid rgba(255, 255, 255, 0.1);
}

/* Responsive */
@media (max-width: 1024px) {
  .details-grid {
    grid-template-columns: repeat(3, 1fr);
  }
}

@media (max-width: 768px) {
  .details-grid {
    grid-template-columns: repeat(3, 1fr);
  }
}

/* Sections Title Bar */
.sections-title-bar {
  margin-top: 32px;
  margin-bottom: 24px;
  padding: 20px 32px;
  background: linear-gradient(135deg, #0a3d42 0%, #052428 100%);
  border-radius: 16px;
  border: 2px solid transparent;
  position: relative;
  overflow: hidden;
}

.sections-title-bar::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  border-radius: 16px;
  padding: 2px;
  background: linear-gradient(135deg, #f59e0b 0%, #f97316 50%, #ef4444 100%);
  -webkit-mask: linear-gradient(#fff 0 0) content-box, linear-gradient(#fff 0 0);
  mask: linear-gradient(#fff 0 0) content-box, linear-gradient(#fff 0 0);
  -webkit-mask-composite: xor;
  mask-composite: exclude;
  pointer-events: none;
}

.sections-title-content {
  display: flex;
  align-items: center;
  gap: 16px;
  position: relative;
  z-index: 1;
}

.sections-title-content i {
  font-size: 32px;
  color: #f59e0b;
  background: linear-gradient(135deg, rgba(245, 158, 11, 0.2) 0%, rgba(249, 115, 22, 0.2) 100%);
  padding: 12px;
  border-radius: 12px;
}

.sections-title-content span {
  font-size: 22px;
  font-weight: 700;
  color: rgba(255, 255, 255, 0.95);
  letter-spacing: 0.5px;
}

/* Management Sections Grid */
.management-sections-grid {
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  gap: 24px;
  margin-top: 24px;
}

.section-card {
  border-radius: 20px;
  padding: 28px;
  position: relative;
  overflow: hidden;
  transition: all 0.4s cubic-bezier(0.4, 0, 0.2, 1);
}

.section-card::before {
  content: '';
  position: absolute;
  top: 0;
  right: 0;
  width: 150px;
  height: 150px;
  background: rgba(255, 255, 255, 0.1);
  border-radius: 50%;
  transform: translate(30%, -30%);
}

.section-card::after {
  content: '';
  position: absolute;
  bottom: 0;
  left: 0;
  width: 100px;
  height: 100px;
  background: rgba(255, 255, 255, 0.05);
  border-radius: 50%;
  transform: translate(-30%, 30%);
}

.section-card:hover {
  transform: translateY(-8px);
  box-shadow: 0 20px 40px rgba(0, 0, 0, 0.3);
}

/* Materials Card */
.materials-card {
  background: linear-gradient(135deg, #0a3d42 0%, #052428 100%);
  border: 2px solid rgba(6, 182, 212, 0.4);
}

.materials-card .section-icon {
  background: linear-gradient(135deg, #06b6d4 0%, #0891b2 100%);
}

/* Labor Card */
.labor-card {
  background: linear-gradient(135deg, #0a3d42 0%, #052428 100%);
  border: 2px solid rgba(16, 185, 129, 0.4);
}

.labor-card .section-icon {
  background: linear-gradient(135deg, #10b981 0%, #059669 100%);
}

/* Equipment Card */
.equipment-card {
  background: linear-gradient(135deg, #0a3d42 0%, #052428 100%);
  border: 2px solid rgba(20, 184, 166, 0.4);
}

.equipment-card .section-icon {
  background: linear-gradient(135deg, #14b8a6 0%, #0d9488 100%);
}

.section-card-content {
  position: relative;
  z-index: 1;
  display: flex;
  flex-direction: column;
  align-items: center;
  text-align: center;
}

.section-icon {
  width: 80px;
  height: 80px;
  background: rgba(255, 255, 255, 0.2);
  border-radius: 20px;
  display: flex;
  align-items: center;
  justify-content: center;
  margin-bottom: 20px;
  backdrop-filter: blur(10px);
}

.section-icon i {
  font-size: 40px;
  color: white;
}

.section-info {
  margin-bottom: 20px;
}

.section-title {
  font-size: 20px;
  font-weight: 700;
  color: rgba(255, 255, 255, 0.95);
  margin: 0 0 8px 0;
}

.section-subtitle {
  font-size: 13px;
  color: rgba(255, 255, 255, 0.6);
  margin: 0;
  line-height: 1.5;
}

.section-stats {
  margin-bottom: 20px;
}

.stat-item {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 4px;
}

.stat-number {
  font-size: 36px;
  font-weight: 800;
  background: linear-gradient(135deg, #06b6d4 0%, #10b981 100%);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background-clip: text;
}

.stat-label {
  font-size: 14px;
  color: rgba(255, 255, 255, 0.8);
  font-weight: 500;
}

.section-btn {
  display: flex;
  align-items: center;
  gap: 8px;
  padding: 12px 24px;
  background: linear-gradient(135deg, #06b6d4 0%, #10b981 100%);
  border: none;
  border-radius: 12px;
  color: white;
  font-size: 14px;
  font-weight: 600;
  cursor: pointer;
  transition: all 0.3s ease;
  box-shadow: 0 4px 15px rgba(6, 182, 212, 0.3);
}

.section-btn:hover {
  transform: translateY(-2px);
  box-shadow: 0 6px 20px rgba(6, 182, 212, 0.4);
}

.section-btn i {
  font-size: 18px;
}

/* Responsive for sections */
@media (max-width: 1024px) {
  .management-sections-grid {
    grid-template-columns: 1fr;
  }
}

@media (max-width: 768px) {
  .management-sections-grid {
    grid-template-columns: 1fr;
  }
  
  .section-card {
    padding: 20px;
  }
}

/* Toggle Switch Styles */
.card-toggle {
  position: absolute;
  top: 16px;
  left: 16px;
  z-index: 10;
}

.toggle-switch {
  position: relative;
  display: inline-block;
  width: 50px;
  height: 26px;
}

.toggle-switch input {
  opacity: 0;
  width: 0;
  height: 0;
}

.toggle-slider {
  position: absolute;
  cursor: pointer;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: rgba(255, 255, 255, 0.2);
  transition: all 0.3s ease;
  border-radius: 26px;
  border: 2px solid rgba(255, 255, 255, 0.3);
}

.toggle-slider:before {
  position: absolute;
  content: "";
  height: 18px;
  width: 18px;
  left: 2px;
  bottom: 2px;
  background: white;
  transition: all 0.3s ease;
  border-radius: 50%;
}

.toggle-switch input:checked + .toggle-slider {
  background: linear-gradient(135deg, #06b6d4 0%, #10b981 100%);
  border-color: transparent;
}

.toggle-switch input:checked + .toggle-slider:before {
  transform: translateX(24px);
}

/* Disabled Card State */
.section-card {
  position: relative;
}

.section-card.card-disabled {
  opacity: 0.5;
  filter: grayscale(50%);
}

.section-card.card-disabled .section-card-content {
  pointer-events: none;
}

.section-card.card-disabled .card-toggle {
  pointer-events: auto;
}

.section-btn:disabled {
  opacity: 0.5;
  cursor: not-allowed;
  transform: none !important;
}

.section-btn:disabled:hover {
  transform: none !important;
  box-shadow: 0 4px 15px rgba(6, 182, 212, 0.3);
}
</style>
