<template>
  <div class="equipment-details-page">
    <!-- Header Section -->
    <div class="page-header glass-effect gradient-animation">
      <div class="header-content">
        <div class="header-text">
          <h1 class="page-title text-glow fade-in">تفاصيل الآليات</h1>
          <p class="page-subtitle fade-in">إدارة المعدات والآلات المستخدمة في المشروع</p>
        </div>
        <v-btn 
          icon="mdi-arrow-right" 
          @click="goBack" 
          class="back-btn"
          size="small"
        >
          <v-icon size="20" color="white">mdi-arrow-right</v-icon>
        </v-btn>
      </div>
    </div>

    <!-- Content Container -->
    <div class="content-container">

    <!-- Equipment Section -->
    <v-card class="section-card mb-4" elevation="2">
      <v-card-title class="section-title">
        <v-icon class="me-2" color="white">mdi-truck</v-icon>
        الآليات والمعدات
      </v-card-title>
      
      <!-- Search Bar and Add Button -->
      <v-card-text class="pb-0">
        <div class="search-container">
          <div class="search-box">
            <v-text-field
              v-model="equipmentSearch"
              label="البحث في الآليات..."
              prepend-inner-icon="mdi-magnify"
              variant="outlined"
              density="comfortable"
              clearable
              hide-details
              class="search-field"
              @keyup.enter="searchEquipment"
            />
            <v-btn 
              color="primary" 
              variant="elevated" 
              class="search-btn"
              @click="searchEquipment"
              size="large"
            >
              <v-icon class="me-2">mdi-magnify</v-icon>
              بحث
            </v-btn>
            <v-btn 
              color="indigo" 
              variant="elevated" 
              class="add-btn"
              @click="openAddDialog"
              size="large"
            >
              <v-icon class="me-2">mdi-plus</v-icon>
              إضافة آلة
            </v-btn>
          </div>
        </div>
      </v-card-text>

      <!-- Equipment Table -->
      <v-data-table
        :headers="equipmentHeaders"
        :items="equipmentData"
        :search="equipmentSearch"
        :items-per-page="10"
        class="data-table equipment-table"
        no-data-text="لا توجد بيانات متاحة"
        loading-text="جاري التحميل..."
        density="comfortable"
      >
        <template v-slot:item.actions="{ item }">
          <v-btn
            icon="mdi-delete"
            size="small"
            color="error"
            variant="text"
            @click="deleteEquipment(item)"
            title="حذف الآلة"
          />
        </template>
      </v-data-table>
    </v-card>

    <!-- Add New Equipment Dialog -->
    <v-dialog v-model="showAddDialog" max-width="600" persistent>
      <v-card class="dialog-card">
        <v-card-title class="dialog-title">
          <v-icon class="me-2">mdi-truck-plus</v-icon>
          إضافة آلة جديدة
        </v-card-title>
        <v-card-text class="dialog-content">
          <v-form ref="form" v-model="formValid">
            <v-row>
              <v-col cols="12" md="6">
                <v-text-field
                  v-model="newEquipment.name"
                  label="اسم الآلة"
                  variant="outlined"
                  :rules="[v => !!v || 'اسم الآلة مطلوب']"
                  required
                />
              </v-col>
              <v-col cols="12" md="6">
                <v-text-field
                  v-model="newEquipment.type"
                  label="نوع الآلة"
                  variant="outlined"
                  :rules="[v => !!v || 'نوع الآلة مطلوب']"
                  required
                />
              </v-col>
              <v-col cols="12" md="6">
                <v-text-field
                  v-model="newEquipment.model"
                  label="الموديل"
                  variant="outlined"
                  :rules="[v => !!v || 'الموديل مطلوب']"
                  required
                />
              </v-col>
              <v-col cols="12" md="6">
                <v-text-field
                  v-model="newEquipment.rentalCost"
                  label="تكلفة الإيجار اليومية"
                  variant="outlined"
                  type="number"
                  :rules="[v => !!v || 'تكلفة الإيجار مطلوبة']"
                  required
                />
              </v-col>
              <v-col cols="12" md="6">
                <v-text-field
                  v-model="newEquipment.operator"
                  label="اسم السائق/المشغل"
                  variant="outlined"
                  :rules="[v => !!v || 'اسم المشغل مطلوب']"
                  required
                />
              </v-col>
              <v-col cols="12" md="6">
                <v-text-field
                  v-model="newEquipment.status"
                  label="الحالة"
                  variant="outlined"
                  :rules="[v => !!v || 'الحالة مطلوبة']"
                  required
                />
              </v-col>
            </v-row>
          </v-form>
        </v-card-text>
        <v-card-actions class="dialog-actions">
          <v-spacer></v-spacer>
          <v-btn @click="closeAddDialog" color="grey">
            إلغاء
          </v-btn>
          <v-btn @click="saveEquipment" color="primary" :disabled="!formValid">
            حفظ
          </v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>

    <!-- Simple Dialog for Add Equipment -->
    <SimpleDialog
      :open="showAddDialog"
      title="إضافة آلة"
      save-text="حفظ"
      icon="mdi-truck-plus"
      @close="closeAddDialog"
      @save="saveEquipmentModern"
    >
      <EquipmentForm
        v-model="newEquipment"
        ref="equipmentFormRef"
      />
    </SimpleDialog>

    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import SimpleDialog from '@/components/SimpleDialog.vue'
import EquipmentForm from '@/components/EquipmentForm.vue'

const router = useRouter()

// Reactive data
const equipmentSearch = ref('')
const showAddDialog = ref(false)
const formValid = ref(false)

// Modern dialog data
const newEquipment = ref({
  type: '',
  dailyCost: '',
  daysUsed: '',
  totalCost: '',
  driverName: '',
  driverPhone: '',
  notes: ''
})

const equipmentFormRef = ref(null)

// Form data
const newEquipmentItem = ref({
  name: '',
  type: '',
  model: '',
  rentalCost: '',
  operator: '',
  status: ''
})

// Table headers
const equipmentHeaders = [
  { title: 'التسلسل', key: 'id', sortable: false },
  { title: 'اسم الآلة', key: 'name', sortable: true },
  { title: 'نوع الآلة', key: 'type', sortable: true },
  { title: 'الموديل', key: 'model', sortable: true },
  { title: 'تكلفة الإيجار اليومية', key: 'rentalCost', sortable: true },
  { title: 'المشغل', key: 'operator', sortable: true },
  { title: 'الحالة', key: 'status', sortable: true },
  { title: 'الإجراءات', key: 'actions', sortable: false }
]

// Sample data
const equipmentData = ref([
  {
    id: 1,
    name: 'حفار هيدروليكي',
    type: 'حفار',
    model: 'CAT 320D',
    rentalCost: 150000,
    operator: 'أحمد محمد',
    status: 'نشط'
  },
  {
    id: 2,
    name: 'رافعة برجية',
    type: 'رافعة',
    model: 'Potain MDT 178',
    rentalCost: 200000,
    operator: 'محمد علي',
    status: 'نشط'
  },
  {
    id: 3,
    name: 'خلاطة خرسانة',
    type: 'خلاطة',
    model: 'Schwing SP 1000',
    rentalCost: 80000,
    operator: 'علي حسن',
    status: 'نشط'
  },
  {
    id: 4,
    name: 'شاحنة نقل',
    type: 'شاحنة',
    model: 'Mercedes Actros',
    rentalCost: 60000,
    operator: 'حسن أحمد',
    status: 'نشط'
  },
  {
    id: 5,
    name: 'جرافة',
    type: 'جرافة',
    model: 'Caterpillar D6T',
    rentalCost: 120000,
    operator: 'أحمد علي',
    status: 'صيانة'
  }
])

// Methods
const goBack = () => {
  router.push('/work-day-details')
}

const searchEquipment = () => {
  console.log('البحث في الآليات:', equipmentSearch.value)
}

const openAddDialog = () => {
  newEquipmentItem.value = {
    name: '',
    type: '',
    model: '',
    rentalCost: '',
    operator: '',
    status: ''
  }
  // Reset modern form data
  newEquipment.value = {
    type: '',
    dailyCost: '',
    daysUsed: '',
    totalCost: '',
    driverName: '',
    driverPhone: '',
    notes: ''
  }
  showAddDialog.value = true
}

const closeAddDialog = () => {
  showAddDialog.value = false
  newEquipment.value = {
    name: '',
    type: '',
    model: '',
    rentalCost: '',
    operator: '',
    status: ''
  }
}

const saveEquipment = () => {
  const newMachine = {
    id: equipmentData.value.length + 1,
    name: newEquipmentItem.value.name,
    type: newEquipmentItem.value.type,
    model: newEquipmentItem.value.model,
    rentalCost: parseInt(newEquipmentItem.value.rentalCost),
    operator: newEquipmentItem.value.operator,
    status: newEquipmentItem.value.status
  }
  equipmentData.value.push(newMachine)
  closeAddDialog()
}

// Modern dialog function for equipment
const saveEquipmentModern = () => {
  if (equipmentFormRef.value && equipmentFormRef.value.isValid()) {
    const newMachine = {
      id: equipmentData.value.length + 1,
      name: newEquipment.value.type,
      type: newEquipment.value.type,
      model: '-',
      rentalCost: parseInt(newEquipment.value.dailyCost),
      operator: newEquipment.value.driverName,
      status: 'نشط',
      daysUsed: parseInt(newEquipment.value.daysUsed),
      totalCost: parseInt(newEquipment.value.totalCost),
      driverPhone: newEquipment.value.driverPhone,
      notes: newEquipment.value.notes
    }
    
    equipmentData.value.push(newMachine)
    
    // Reset form
    equipmentFormRef.value.resetForm()
    newEquipment.value = {
      type: '',
      dailyCost: '',
      daysUsed: '',
      totalCost: '',
      driverName: '',
      driverPhone: '',
      notes: ''
    }
    
    showAddDialog.value = false
  }
}

const deleteEquipment = (item) => {
  if (confirm('هل أنت متأكد من حذف هذه الآلة؟')) {
    const index = equipmentData.value.findIndex(e => e.id === item.id)
    if (index > -1) {
      equipmentData.value.splice(index, 1)
    }
  }
}

// Lifecycle
onMounted(() => {
  console.log('✅ صفحة تفاصيل الآليات تم تحميلها بنجاح!')
})
</script>

<style scoped>
.equipment-details-page {
  min-height: 100vh;
  background: #f5f5f5;
  padding: 0 !important;
  padding-top: 0 !important;
  font-family: 'Cairo', 'Tajawal', 'Arial', sans-serif !important;
  position: relative;
  overflow-x: hidden;
  direction: rtl;
  margin: 0;
  width: 100%;
}

.content-container {
  padding: 0 2rem 2rem 2rem !important;
  max-width: 100%;
}

.page-header {
  background: linear-gradient(135deg, #60a5fa 0%, #3b82f6 50%, #2563eb 100%) !important;
  color: white;
  padding: 0.4rem 0.75rem !important;
  border-radius: 0 !important;
  margin: 0 !important;
  margin-bottom: 1rem !important;
  margin-top: 0 !important;
  position: relative;
  overflow: hidden;
  box-shadow: 0 2px 8px rgba(25, 118, 210, 0.15);
  border: none !important;
  border-bottom: 1px solid rgba(255, 255, 255, 0.3) !important;
  backdrop-filter: blur(15px);
  z-index: 1;
  width: 100% !important;
  max-width: 100% !important;
  box-sizing: border-box !important;
}

.header-content {
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 1rem;
  position: relative;
  z-index: 1;
}

.header-text {
  flex: 1;
  position: relative;
  z-index: 1;
}

.back-btn {
  background: rgba(255, 255, 255, 0.15) !important;
  backdrop-filter: blur(15px);
  border: 1px solid rgba(255, 255, 255, 0.2);
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
  transition: all 0.3s ease;
  position: relative;
  z-index: 1;
  min-width: 36px !important;
  height: 36px !important;
}

.back-btn:hover {
  background: rgba(255, 255, 255, 0.25) !important;
  transform: translateX(2px);
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.15);
}

.page-title {
  font-size: 1.4rem !important;
  font-weight: 700 !important;
  margin: 0;
  padding: 0.4rem 0.8rem !important;
  text-shadow: 0 2px 6px rgba(0, 0, 0, 0.3), 0 1px 3px rgba(0, 0, 0, 0.2);
  letter-spacing: 0.3px !important;
  color: #ffffff !important;
  border-radius: 6px !important;
}

.page-subtitle {
  font-size: 0.9rem !important;
  opacity: 0.95;
  margin: 0.3rem 0 0 0 !important;
  font-weight: 400;
  text-shadow: 0 1px 3px rgba(0, 0, 0, 0.2);
  letter-spacing: 0.3px !important;
  color: rgba(255, 255, 255, 0.95);
}

.section-card {
  border-radius: 16px;
  overflow: hidden;
}

.section-title {
  background: linear-gradient(135deg, #60a5fa 0%, #3b82f6 100%);
  color: white;
  font-weight: 600 !important;
  font-size: 1rem !important;
  text-shadow: 0 2px 4px rgba(0, 0, 0, 0.2);
  padding: 0.6rem 1rem !important;
}

.section-title :deep(.v-icon) {
  font-size: 18px !important;
  color: #ffffff !important;
}

/* تحسين تصميم شريط البحث */
.search-container {
  padding: 0.75rem 1rem !important;
  background: linear-gradient(135deg, #eff6ff 0%, #dbeafe 50%, #bfdbfe 100%);
  border-radius: 8px !important;
  margin-bottom: 0.75rem !important;
  box-shadow: 0 2px 8px rgba(59, 130, 246, 0.15);
  border: 1px solid #93c5fd !important;
  position: relative;
}

.search-container::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  height: 2px !important;
  background: linear-gradient(90deg, #60a5fa, #3b82f6, #2563eb);
  border-radius: 8px 8px 0 0;
}

.search-box {
  display: flex;
  align-items: center;
  gap: 0.5rem !important;
  max-width: 600px;
  margin: 0 auto;
}

.search-field {
  flex: 1;
  margin-bottom: 0;
}

.search-field :deep(.v-field) {
  background-color: #ffffff !important;
  border-radius: 8px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.2) !important;
  border: 2px solid #3b82f6 !important;
  transition: all 0.3s ease;
  min-height: 36px !important;
}

.search-field :deep(.v-field__input) {
  background-color: #ffffff !important;
  border-radius: 8px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.2) !important;
  border: 2px solid #3b82f6 !important;
  transition: all 0.3s ease;
}

.search-field :deep(.v-field__outline) {
  border-color: #3b82f6 !important;
  border-width: 2px !important;
}

.search-field :deep(.v-field--focused .v-field__outline) {
  border-color: #2563eb !important;
  border-width: 3px !important;
  box-shadow: 0 0 0 3px rgba(59, 130, 246, 0.2) !important;
}

.search-field :deep(.v-field__input):focus {
  box-shadow: 0 4px 16px rgba(59, 130, 246, 0.4) !important;
  border-color: #2563eb !important;
  transform: translateY(-1px);
}

.search-field :deep(.v-label) {
  color: #1e40af !important;
  font-weight: 600 !important;
  font-size: 0.875rem !important;
  opacity: 1 !important;
}

.search-field :deep(.v-label--active),
.search-field :deep(.v-field__label--active) {
  color: #1e40af !important;
  font-weight: 700 !important;
  opacity: 1 !important;
}

.search-field :deep(.v-field__input input) {
  color: #111827 !important;
  font-weight: 500 !important;
  font-size: 0.875rem !important;
}

.search-field :deep(.v-field__prepend-inner) {
  color: #3b82f6 !important;
}

.search-field :deep(.v-field__prepend-inner .v-icon) {
  color: #3b82f6 !important;
  font-size: 18px !important;
}

.search-field :deep(.v-field__prepend-inner i) {
  color: #3b82f6 !important;
}

.search-btn {
  height: 36px !important;
  min-width: 80px !important;
  border-radius: 8px;
  font-weight: 600 !important;
  font-size: 0.875rem !important;
  text-transform: none;
  box-shadow: 0 2px 8px rgba(25, 118, 210, 0.25);
  transition: all 0.3s ease;
}

.search-btn:hover {
  transform: translateY(-2px);
  box-shadow: 0 6px 16px rgba(25, 118, 210, 0.4);
}

/* أنماط أزرار الإضافة */
.add-btn {
  height: 36px !important;
  min-width: 100px !important;
  border-radius: 8px;
  font-weight: 600 !important;
  font-size: 0.875rem !important;
  text-transform: none;
  transition: all 0.3s ease;
  background: linear-gradient(135deg, #3f51b5 0%, #303f9f 100%) !important;
  box-shadow: 0 2px 8px rgba(63, 81, 181, 0.25);
}

.add-btn :deep(.v-icon) {
  font-size: 18px !important;
}

.add-btn:hover {
  transform: translateY(-2px);
  box-shadow: 0 6px 16px rgba(63, 81, 181, 0.4);
  background: linear-gradient(135deg, #303f9f 0%, #283593 100%) !important;
}

.data-table {
  margin-top: 1rem;
}

/* تحسين ألوان النصوص في الجداول */
.data-table :deep(.v-data-table__td) {
  color: #1a1a1a !important;
  font-weight: 500 !important;
  font-size: 0.6rem !important;
  text-shadow: 0 1px 2px rgba(0, 0, 0, 0.1);
  padding: 3px 4px !important;
  min-height: 32px !important;
}

.data-table :deep(.v-data-table__th) {
  color: #ffffff !important;
  font-weight: 500 !important;
  font-size: 0.55rem !important;
  background: linear-gradient(135deg, #1976d2 0%, #1565c0 100%) !important;
  text-shadow: 0 1px 2px rgba(0, 0, 0, 0.25);
  padding: 2px 4px !important;
  min-height: 20px !important;
}

.data-table :deep(.v-data-table__tr:hover) {
  background-color: #f8f9fa !important;
}

.data-table :deep(.v-data-table__tr:hover .v-data-table__td) {
  color: #000 !important;
}

/* تحسين ألوان النصوص في الجداول الفرعية */
.data-table :deep(.v-data-table__wrapper) {
  border: 2px solid #e8eaf6;
  border-radius: 12px;
  overflow: hidden;
  box-shadow: 0 4px 12px rgba(63, 81, 181, 0.1);
  background: white;
}

.data-table :deep(.v-data-table__wrapper table) {
  border-collapse: separate;
  border-spacing: 0;
}

.data-table :deep(.v-data-table__wrapper table td) {
  color: #333 !important;
  border-color: #e0e0e0 !important;
}

.data-table :deep(.v-data-table__wrapper table th) {
  color: #ffffff !important;
  background: linear-gradient(135deg, #3f51b5 0%, #303f9f 100%) !important;
  border-color: #283593 !important;
}

/* أنماط مخصصة لجداول الآليات */
.equipment-table :deep(.v-data-table__td) {
  color: #1a1a1a !important;
  font-weight: 500 !important;
  font-size: 0.6rem !important;
  text-shadow: 0 1px 2px rgba(0, 0, 0, 0.1);
  padding: 3px 4px !important;
  min-height: 32px !important;
}

.equipment-table :deep(.v-data-table__th) {
  color: #ffffff !important;
  font-weight: 500 !important;
  font-size: 0.55rem !important;
  background: linear-gradient(135deg, #1976d2 0%, #1565c0 100%) !important;
  text-shadow: 0 1px 2px rgba(0, 0, 0, 0.25);
  padding: 2px 4px !important;
  min-height: 20px !important;
  border-bottom: 1px solid #0d47a1 !important;
  position: relative;
}

/* تحسين ألوان الأزرار في الجداول */
.data-table :deep(.v-btn) {
  color: #dc3545 !important;
  transition: all 0.3s ease;
  min-width: 28px !important;
  height: 28px !important;
}

.data-table :deep(.v-btn .v-icon) {
  font-size: 16px !important;
}

.data-table :deep(.v-btn:hover) {
  background-color: #f8d7da !important;
  color: #721c24 !important;
  transform: scale(1.05);
  box-shadow: 0 2px 6px rgba(220, 53, 69, 0.25);
}

.dialog-card {
  border-radius: 16px;
  overflow: hidden;
}

.dialog-title {
  background: linear-gradient(135deg, #3f51b5 0%, #303f9f 100%);
  color: white;
  font-weight: bold;
  font-size: 1.3rem;
}

.dialog-content {
  padding: 2rem;
}

.dialog-actions {
  padding: 1rem 2rem;
  background: #f8f9fa;
}

/* Animations */
.glass-effect {
  backdrop-filter: blur(15px);
}

.page-header.glass-effect {
  background: linear-gradient(135deg, #60a5fa 0%, #3b82f6 50%, #2563eb 100%) !important;
}

.gradient-animation {
  background-size: 400% 400%;
  animation: gradientShift 20s ease infinite;
}

@keyframes gradientShift {
  0% { background-position: 0% 50%; }
  50% { background-position: 100% 50%; }
  100% { background-position: 0% 50%; }
}

.text-glow {
  text-shadow: 0 0 30px rgba(255, 255, 255, 0.6);
}

.fade-in {
  animation: fadeIn 1.2s ease-out;
}

@keyframes fadeIn {
  from { 
    opacity: 0; 
    transform: translateY(30px) scale(0.95); 
  }
  to { 
    opacity: 1; 
    transform: translateY(0) scale(1); 
  }
}

/* تأثيرات إضافية للرأس */
.page-header::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: linear-gradient(45deg, transparent 30%, rgba(255, 255, 255, 0.15) 50%, transparent 70%);
  animation: shimmer 3s ease-in-out infinite;
  pointer-events: none;
  border-radius: 20px;
}

@keyframes shimmer {
  0% { transform: translateX(-100%); }
  100% { transform: translateX(100%); }
}

.page-header::after {
  content: '';
  position: absolute;
  top: -50%;
  left: -50%;
  width: 200%;
  height: 200%;
  background: radial-gradient(circle, rgba(255, 255, 255, 0.15) 0%, rgba(66, 165, 245, 0.1) 50%, transparent 70%);
  animation: pulse 4s ease-in-out infinite;
  pointer-events: none;
  border-radius: 50%;
}

@keyframes pulse {
  0%, 100% { opacity: 0.3; transform: scale(0.8); }
  50% { opacity: 0.6; transform: scale(1.2); }
}

/* Responsive */
@media (max-width: 768px) {
  .equipment-details-page {
    padding: 1rem;
  }
  
  .page-title {
    font-size: 2rem;
  }
  
  .header-content {
    flex-direction: column;
    text-align: center;
  }
  
  .dialog-content {
    padding: 1rem;
  }
  
  .search-box {
    flex-direction: column;
    gap: 0.5rem;
  }
  
  .search-btn, .add-btn {
    width: 100%;
    min-width: auto;
  }
  
  .search-container {
    padding: 0.75rem;
  }
  
  .search-field {
    margin-bottom: 0.5rem;
  }
}
</style>
