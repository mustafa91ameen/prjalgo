<template>
  <div class="projects-container">
    <!-- Page Header Component -->
    <PageHeader
      title="إدارة المشاريع"
      subtitle="عرض وإدارة جميع المشاريع في النظام"
      :badge="`${projectStats.total} مشروع`"
      badgeType="info"
      class="projects-header"
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
      <!-- Total Projects Card -->
      <v-card class="stat-card" elevation="0">
        <div class="stat-card-content">
          <div class="stat-icon">
            <i class="mdi mdi-folder-multiple"></i>
          </div>
          <div class="stat-info">
            <div class="stat-label">إجمالي المشاريع</div>
            <div class="stat-value">{{ projectStats.total }}</div>
          </div>
        </div>
      </v-card>

      <!-- Active Projects Card -->
      <v-card class="stat-card" elevation="0">
        <div class="stat-card-content">
          <div class="stat-icon success">
            <i class="mdi mdi-briefcase"></i>
          </div>
          <div class="stat-info">
            <div class="stat-label">مشاريع نشطة</div>
            <div class="stat-value">{{ projectStats.active }}</div>
          </div>
        </div>
      </v-card>

      <!-- Pending Projects Card -->
      <v-card class="stat-card" elevation="0">
        <div class="stat-card-content">
          <div class="stat-icon warning">
            <i class="mdi mdi-clock-alert"></i>
          </div>
          <div class="stat-info">
            <div class="stat-label">في الانتظار</div>
            <div class="stat-value">{{ projectStats.pending }}</div>
          </div>
        </div>
      </v-card>

      <!-- Total Budget Card -->
      <v-card class="stat-card" elevation="0">
        <div class="stat-card-content">
          <div class="stat-icon profit">
            <i class="mdi mdi-cash-multiple"></i>
          </div>
          <div class="stat-info">
            <div class="stat-label">إجمالي الميزانية</div>
            <div class="stat-value">{{ formatCurrency(projectStats.totalBudget) }}</div>
          </div>
        </div>
      </v-card>

      <!-- Average Progress Card -->
      <v-card class="stat-card" elevation="0">
        <div class="stat-card-content">
          <div class="stat-icon">
            <i class="mdi mdi-chart-line"></i>
          </div>
          <div class="stat-info">
            <div class="stat-label">متوسط التقدم</div>
            <div class="stat-value">{{ projectStats.avgProgress }}%</div>
          </div>
        </div>
      </v-card>
    </div>

    <!-- Projects List Header -->
    <div class="projects-list-header">
      <div class="list-header-content">
        <div class="list-header-info">
          <h2 class="list-header-title">
            <i class="mdi mdi-format-list-bulleted"></i>
            قائمة المشاريع
          </h2>
          <p class="list-header-subtitle">عرض جميع المشاريع المتاحة</p>
        </div>
        <div class="list-header-actions">
          <button v-if="canCreate" class="list-action-btn primary" @click="openAddDialog">
            <i class="mdi mdi-plus"></i>
            مشروع جديد
          </button>
        </div>
      </div>
    </div>

    <!-- Projects Grid -->
    <div class="projects-grid">
      <v-card
        v-for="project in projects"
        :key="project.id"
        class="project-card"
        elevation="2"
      >
        <div class="project-card-header">
          <div class="project-status-badge" :class="project.statusClass">
            {{ project.status }}
          </div>
        </div>
        
        <v-card-text class="project-card-body">
          <h3 class="project-title">{{ project.name }}</h3>
          <p class="project-description">{{ project.description }}</p>
          
          <div class="project-meta">
            <div class="meta-item">
              <i class="mdi mdi-map-marker"></i>
              <span>{{ project.location }}</span>
            </div>
            <div class="meta-item">
              <i class="mdi mdi-calendar"></i>
              <span>{{ project.date }}</span>
            </div>
          </div>

          <div class="project-progress">
            <div class="progress-header">
              <span>نسبة الإنجاز</span>
              <span class="progress-value">{{ project.progress }}%</span>
            </div>
            <v-progress-linear
              :model-value="project.progress"
              :color="getProgressColor(project.progress)"
              height="8"
              rounded
            ></v-progress-linear>
          </div>

          <div class="project-footer">
            <div class="project-team">
              <i class="mdi mdi-account-group"></i>
              <span>{{ project.teamSize }} أعضاء</span>
            </div>
            <div class="project-budget">
              <i class="mdi mdi-cash"></i>
              <span>{{ formatCurrency(project.budget) }}</span>
            </div>
          </div>
        </v-card-text>

        <v-card-actions class="project-actions">
          <v-btn variant="text" color="#06b6d4" size="small" @click="$router.push({ path: '/work-days', query: { project_id: project.id } })">
            <i class="mdi mdi-eye"></i>
            عرض التفاصيل
          </v-btn>
        </v-card-actions>
      </v-card>
    </div>

    <!-- Add New Project Dialog -->
    <v-dialog v-model="showAddDialog" max-width="800" persistent>
      <v-card class="project-dialog">
        <v-card-title class="dialog-header">
          <i class="mdi mdi-folder-plus"></i>
          معلومات المشروع
        </v-card-title>
        
        <v-card-text class="dialog-content">
          <v-form ref="projectForm" v-model="formValid">
            <v-row>
              <!-- اسم المشروع -->
              <v-col cols="12" md="6">
                <v-text-field
                  v-model="newProject.name"
                  label="اسم المشروع"
                  variant="outlined"
                  density="comfortable"
                  :rules="[v => !!v || 'اسم المشروع مطلوب']"
                  required
                ></v-text-field>
              </v-col>

              <!-- مكان المشروع -->
              <v-col cols="12" md="6">
                <v-text-field
                  v-model="newProject.location"
                  label="مكان المشروع"
                  variant="outlined"
                  density="comfortable"
                  :rules="[v => !!v || 'مكان المشروع مطلوب']"
                  required
                ></v-text-field>
              </v-col>

              <!-- حد التنبيه (د.ع) -->
              <v-col cols="12" md="6">
                <v-text-field
                  v-model.number="newProject.alertLimit"
                  label="حد التنبيه (د.ع)"
                  variant="outlined"
                  density="comfortable"
                  type="number"
                  :rules="[v => v >= 0 || 'يجب أن يكون رقم موجب']"
                ></v-text-field>
              </v-col>

              <!-- مدة المشروع (أيام) -->
              <v-col cols="12" md="6">
                <v-text-field
                  v-model.number="newProject.duration"
                  label="مدة المشروع (أيام)"
                  variant="outlined"
                  density="comfortable"
                  type="number"
                  :rules="[v => v > 0 || 'يجب أن يكون أكبر من صفر']"
                  required
                ></v-text-field>
              </v-col>

              <!-- التكلفة الإجمالية -->
              <v-col cols="12" md="6">
                <v-text-field
                  v-model.number="newProject.totalCost"
                  label="التكلفة الإجمالية"
                  variant="outlined"
                  density="comfortable"
                  type="number"
                  :rules="[v => v > 0 || 'يجب أن يكون أكبر من صفر']"
                  required
                ></v-text-field>
              </v-col>

              <!-- تاريخ البدء -->
              <v-col cols="12" md="6">
                <v-text-field
                  v-model="newProject.startDate"
                  label="تاريخ البدء"
                  variant="outlined"
                  density="comfortable"
                  type="date"
                  :rules="[v => !!v || 'تاريخ البدء مطلوب']"
                  required
                ></v-text-field>
              </v-col>

              <!-- رقم العميل -->
              <v-col cols="12" md="6">
                <v-text-field
                  v-model="newProject.clientNumber"
                  label="رقم العميل"
                  variant="outlined"
                  density="comfortable"
                  :rules="[v => !!v || 'رقم العميل مطلوب']"
                  required
                ></v-text-field>
              </v-col>

              <!-- وصف المشروع -->
              <v-col cols="12">
                <v-textarea
                  v-model="newProject.description"
                  label="وصف المشروع"
                  variant="outlined"
                  density="comfortable"
                  rows="3"
                  :rules="[v => !!v || 'وصف المشروع مطلوب']"
                  required
                ></v-textarea>
              </v-col>

              <!-- ملاحظات -->
              <v-col cols="12">
                <v-textarea
                  v-model="newProject.notes"
                  label="ملاحظات"
                  variant="outlined"
                  density="comfortable"
                  rows="2"
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
          <button class="dialog-btn save" @click="saveProject" :disabled="!formValid">
            <i class="mdi mdi-content-save"></i>
            حفظ
          </button>
        </v-card-actions>
      </v-card>
    </v-dialog>

  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import PageHeader from '../components/PageHeader.vue'
import { listProjects, createProject, updateProject, getProjectStats } from '@/api/projects'
import { usePermissions } from '@/composables/usePermissions'
import { useToast } from '@/composables/useToast'
import { DEFAULT_PAGE, DEFAULT_LIMIT } from '@/constants/pagination'

const { success, error: showError } = useToast()
const { canCreate, canUpdate, canDelete } = usePermissions('/projects')

const searchQuery = ref('')
const statusFilter = ref('الكل')
const categoryFilter = ref('الكل')
const loading = ref(false)

// Pagination
const page = ref(DEFAULT_PAGE)
const limit = ref(DEFAULT_LIMIT)
const total = ref(0)

// Stats
const projectStats = ref({
  total: 0,
  active: 0,
  pending: 0,
  completed: 0,
  totalBudget: 0,
  avgProgress: 0
})

const statusOptions = ['الكل', 'قيد التنفيذ', 'مكتمل', 'معلق', 'ملغي']
const categoryOptions = ['الكل', 'إنشاءات', 'صيانة', 'تطوير', 'استشارات']

// Dialog state
const showAddDialog = ref(false)
const formValid = ref(false)
const projectForm = ref(null)
const editingProject = ref(null)

// New project data
const newProject = ref({
  name: '',
  location: '',
  alertLimit: 0,
  duration: 0,
  totalCost: 0,
  startDate: '',
  clientNumber: '',
  description: '',
  notes: ''
})

// Open dialog function
const openAddDialog = () => {
  showAddDialog.value = true
}

// Close dialog function
const closeDialog = () => {
  showAddDialog.value = false
  editingProject.value = null
  resetForm()
}

// Reset form
const resetForm = () => {
  newProject.value = {
    name: '',
    location: '',
    alertLimit: 0,
    duration: 0,
    totalCost: 0,
    startDate: '',
    clientNumber: '',
    description: '',
    notes: ''
  }
  if (projectForm.value) {
    projectForm.value.reset()
  }
}

// Save project function
const saveProject = async () => {
  if (!formValid.value) return

  loading.value = true
  try {
    const projectData = {
      name: newProject.value.name,
      location: newProject.value.location,
      warningCost: newProject.value.alertLimit || 0,
      duration: newProject.value.duration || 0,
      totalCost: newProject.value.totalCost || 0,
      startDate: newProject.value.startDate ? new Date(newProject.value.startDate).toISOString() : null,
      clientPhone: newProject.value.clientNumber,
      description: newProject.value.description,
      notes: newProject.value.notes
    }

    let response
    if (editingProject.value) {
      response = await updateProject(editingProject.value.id, projectData)
    } else {
      response = await createProject(projectData)
    }

    if (response.success) {
      success(editingProject.value ? 'تم تحديث المشروع بنجاح' : 'تم إضافة المشروع بنجاح')
      closeDialog()
      fetchProjects()
      fetchStats()
    } else {
      showError(response.message || 'حدث خطأ')
    }
  } catch (error) {
    console.error('Error saving project:', error)
    showError('حدث خطأ في حفظ المشروع')
  } finally {
    loading.value = false
  }
}

// Edit project
const editProject = (project) => {
  editingProject.value = project
  newProject.value = {
    name: project.name,
    location: project.location,
    alertLimit: project.alertLimit || 0,
    duration: project.duration || 0,
    totalCost: project.budget || 0,
    startDate: project.date,
    clientNumber: project.clientNumber || '',
    description: project.description,
    notes: ''
  }
  showAddDialog.value = true
}

// Initialize
onMounted(() => {
  fetchProjects()
  fetchStats()
})

const projects = ref([])

// Fetch projects from API
const fetchProjects = async () => {
  loading.value = true
  try {
    const params = {
      page: page.value,
      limit: limit.value
    }

    // Add status filter if not 'all'
    if (statusFilter.value !== 'الكل') {
      const statusMap = {
        'قيد التنفيذ': 'in_progress',
        'مكتمل': 'completed',
        'معلق': 'pending',
        'ملغي': 'cancelled'
      }
      params.status = statusMap[statusFilter.value]
    }

    const response = await listProjects(params)
    if (response.success) {
      projects.value = (response.data.items || []).map(p => ({
        id: p.id,
        name: p.name,
        description: p.description || p.notes || '',
        location: p.location,
        date: p.startDate,
        progress: p.progressPercentage || 0,
        status: getStatusText(p.status),
        statusClass: getStatusClass(p.status),
        teamSize: p.teamSize || 0,
        budget: p.totalCost || 0,
        alertLimit: p.warningCost,
        duration: p.duration,
        clientNumber: p.clientPhone
      }))
      total.value = response.data.total || 0
    }
  } catch (error) {
    console.error('Error fetching projects:', error)
    showError('حدث خطأ في جلب المشاريع')
  } finally {
    loading.value = false
  }
}

// Fetch project stats
const fetchStats = async () => {
  try {
    const response = await getProjectStats()
    if (response.success) {
      projectStats.value = {
        total: response.data.total || 0,
        active: response.data.inProgress || 0,
        pending: response.data.pending || 0,
        completed: response.data.completed || 0,
        totalBudget: response.data.totalBudget || 0,
        avgProgress: response.data.averageProgress || 0
      }
    }
  } catch (error) {
    console.error('Error fetching stats:', error)
  }
}

// Helper functions
const getStatusText = (status) => {
  const statusMap = {
    'in_progress': 'قيد التنفيذ',
    'completed': 'مكتمل',
    'pending': 'معلق',
    'cancelled': 'ملغي'
  }
  return statusMap[status] || status
}

const getStatusClass = (status) => {
  const classMap = {
    'in_progress': 'status-active',
    'completed': 'status-completed',
    'pending': 'status-pending',
    'cancelled': 'status-cancelled'
  }
  return classMap[status] || 'status-pending'
}

const getProgressColor = (progress) => {
  if (progress >= 80) return '#4CAF50'
  if (progress >= 50) return '#2196F3'
  if (progress >= 30) return '#FF9800'
  return '#F44336'
}

const formatCurrency = (amount) => {
  return new Intl.NumberFormat('en-US').format(amount) + ' د.ع'
}
</script>

<style scoped>
.projects-container {
  padding: 32px;
  max-width: 1600px;
  margin: 0 auto;
}

/* Projects Header Custom Color */
.projects-header {
  background: linear-gradient(135deg, #018790 0%, #005461 100%) !important;
}

.projects-header::before {
  background: linear-gradient(135deg, #06b6d4 0%, #10b981 50%, #06b6d4 100%) !important;
}

/* Statistics Grid */
.stats-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(220px, 1fr));
  gap: 20px;
  margin-bottom: 32px;
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
  transform: translateY(-4px);
  box-shadow: 0 12px 24px rgba(6, 182, 212, 0.2);
}

.stat-card-content {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  padding: 24px 16px;
  text-align: center;
  position: relative;
  z-index: 1;
}

.stat-info {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 6px;
}

.stat-label {
  font-size: 13px;
  color: rgba(255, 255, 255, 0.7);
  margin-bottom: 4px;
  font-weight: 500;
  letter-spacing: 0.3px;
}

.stat-value {
  font-size: 32px;
  font-weight: 800;
  background: linear-gradient(135deg, #ffffff 0%, #e0e7ff 100%);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background-clip: text;
  margin-bottom: 6px;
  line-height: 1;
}

.stat-change {
  display: flex;
  align-items: center;
  gap: 4px;
  font-size: 12px;
  font-weight: 600;
  padding: 6px 12px;
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
  width: 56px;
  height: 56px;
  border-radius: 14px;
  display: flex;
  align-items: center;
  justify-content: center;
  color: white;
  font-size: 28px;
  margin-bottom: 16px;
  background: linear-gradient(135deg, #06b6d4 0%, #0891b2 100%) !important;
  box-shadow: 0 8px 20px rgba(6, 182, 212, 0.4);
}

.stat-icon.success {
  background: linear-gradient(135deg, #10b981 0%, #059669 100%) !important;
  box-shadow: 0 8px 20px rgba(16, 185, 129, 0.4);
}

.stat-icon.warning {
  background: linear-gradient(135deg, #f59e0b 0%, #d97706 100%) !important;
  box-shadow: 0 8px 20px rgba(245, 158, 11, 0.4);
}

.stat-icon.profit {
  background: linear-gradient(135deg, #10b981 0%, #34d399 100%) !important;
  box-shadow: 0 8px 20px rgba(16, 185, 129, 0.4);
}

/* Projects List Header */
.projects-list-header {
  background: linear-gradient(135deg, #0a3d42 0%, #052428 100%);
  border-radius: 16px;
  margin-bottom: 24px;
  padding: 20px 28px;
  border: 2px solid transparent;
  position: relative;
  overflow: hidden;
}

.projects-list-header::before {
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
  font-size: 22px;
  font-weight: 700;
  color: rgba(255, 255, 255, 0.95);
  margin: 0 0 4px 0;
  display: flex;
  align-items: center;
  gap: 10px;
}

.list-header-title i {
  color: #06b6d4;
  font-size: 24px;
}

.list-header-subtitle {
  font-size: 13px;
  color: rgba(255, 255, 255, 0.6);
  margin: 0;
}

.list-header-actions {
  display: flex;
  gap: 12px;
}

.list-action-btn {
  padding: 10px 20px;
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
  font-size: 18px;
}

/* Project Dialog */
.project-dialog {
  border-radius: 16px !important;
  background: linear-gradient(135deg, #0a3d42 0%, #052428 100%) !important;
  border: 2px solid transparent !important;
  position: relative;
  overflow: hidden;
  direction: rtl;
  text-align: right;
}

.project-dialog::before {
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

.dialog-header {
  background: rgba(6, 182, 212, 0.1) !important;
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
  color: #06b6d4;
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
  color: #06b6d4 !important;
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

/* Filter Section */
.projects-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(350px, 1fr));
  gap: 24px;
}

.project-card {
  border-radius: 20px !important;
  overflow: hidden;
  transition: all 0.3s ease;
  background: linear-gradient(135deg, #0a3d42 0%, #052428 100%) !important;
  border: 2px solid transparent !important;
  position: relative;
}

.project-card::before {
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

.project-card:hover {
  transform: translateY(-4px);
  box-shadow: 0 12px 24px rgba(6, 182, 212, 0.3);
}

.project-card:hover::before {
  opacity: 1;
}

.project-card-header {
  height: 40px;
  position: relative;
  background: linear-gradient(135deg, #06b6d4 0%, #10b981 100%);
  display: flex;
  align-items: center;
}

.project-status-badge {
  position: absolute;
  top: 50%;
  right: 20px;
  transform: translateY(-50%);
  padding: 6px 14px;
  border-radius: 16px;
  font-size: 11px;
  font-weight: 600;
  color: white;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.2);
}

.status-active {
  background: linear-gradient(135deg, #2196F3 0%, #1976D2 100%);
}

.status-completed {
  background: linear-gradient(135deg, #4CAF50 0%, #388E3C 100%);
}

.status-pending {
  background: linear-gradient(135deg, #FF9800 0%, #F57C00 100%);
}

.project-card-body {
  padding: 24px !important;
  position: relative;
  z-index: 1;
}

.project-title {
  font-size: 18px;
  font-weight: 700;
  color: rgba(255, 255, 255, 0.95);
  margin-bottom: 8px;
  line-height: 1.4;
  display: -webkit-box;
  -webkit-line-clamp: 2;
  -webkit-box-orient: vertical;
  overflow: hidden;
}

.project-description {
  font-size: 14px;
  color: rgba(255, 255, 255, 0.7);
  margin-bottom: 16px;
  line-height: 1.6;
  display: -webkit-box;
  -webkit-line-clamp: 2;
  -webkit-box-orient: vertical;
  overflow: hidden;
}

.project-meta {
  display: flex;
  flex-direction: column;
  gap: 8px;
  margin-bottom: 16px;
}

.meta-item {
  display: flex;
  align-items: center;
  gap: 8px;
  font-size: 13px;
  color: rgba(255, 255, 255, 0.7);
  font-weight: 500;
}

.meta-item i {
  color: #06b6d4;
  font-size: 16px;
  width: 18px;
  text-align: center;
}

.project-progress {
  margin-bottom: 16px;
}

.progress-header {
  display: flex;
  justify-content: space-between;
  margin-bottom: 8px;
  font-size: 13px;
  color: rgba(255, 255, 255, 0.7);
  font-weight: 500;
}

.progress-value {
  color: rgba(255, 255, 255, 0.95);
  font-weight: 700;
  font-size: 14px;
}

.project-footer {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding-top: 16px;
  border-top: 1px solid rgba(255, 255, 255, 0.1);
}

.project-team,
.project-budget {
  display: flex;
  align-items: center;
  gap: 6px;
  font-size: 13px;
  color: rgba(255, 255, 255, 0.7);
  font-weight: 600;
}

.project-team i,
.project-budget i {
  color: #06b6d4;
  font-size: 16px;
}

.project-budget {
  color: #10b981;
  font-weight: 700;
}

.project-actions {
  padding: 12px 16px !important;
  border-top: 1px solid #f0f0f0;
  display: flex;
  justify-content: space-around;
}

/* Responsive Design */
@media (max-width: 768px) {
  .projects-container {
    padding: 16px;
  }

  .page-header {
    flex-direction: column;
    align-items: flex-start;
    gap: 16px;
  }

  .stats-row {
    grid-template-columns: 1fr;
  }

  .projects-grid {
    grid-template-columns: 1fr;
  }
}
</style>
