<template>
  <div class="dashboard-container">
    <!-- Page Header Component -->
    <PageHeader
      title="لوحة التحكم"
      subtitle="مرحباً بك في نظام إدارة المشاريع والبناء"
      badge="مباشر"
      badgeType="success"
      class="dashboard-header"
    >
      <template #actions>
        <button class="page-action-btn secondary" @click="refreshData" :disabled="loading">
          <i class="mdi mdi-refresh" :class="{ 'mdi-spin': loading }"></i>
          تحديث البيانات
        </button>
        <button class="page-icon-btn">
          <i class="mdi mdi-filter-variant"></i>
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
            <div class="stat-value">{{ stats.totalProjects }}</div>
            <div class="stat-change positive">
              <i class="mdi mdi-trending-up"></i>
              <span>+12% هذا الشهر</span>
            </div>
          </div>
        </div>
      </v-card>

      <!-- Active Projects Card -->
      <v-card class="stat-card" elevation="0">
        <div class="stat-card-content">
          <div class="stat-icon">
            <i class="mdi mdi-briefcase"></i>
          </div>
          <div class="stat-info">
            <div class="stat-label">المشاريع النشطة</div>
            <div class="stat-value">{{ stats.activeProjects }}</div>
            <div class="stat-change positive">
              <i class="mdi mdi-trending-up"></i>
              <span>+8% هذا الشهر</span>
            </div>
          </div>
        </div>
      </v-card>

      <!-- Total Revenue Card -->
      <v-card class="stat-card" elevation="0">
        <div class="stat-card-content">
          <div class="stat-icon">
            <i class="mdi mdi-cash-plus"></i>
          </div>
          <div class="stat-info">
            <div class="stat-label">إجمالي الإيرادات</div>
            <div class="stat-value">{{ formatCurrency(financialStats.totalIncome) }}</div>
            <div class="stat-change positive">
              <i class="mdi mdi-trending-up"></i>
              <span>+15% هذا الشهر</span>
            </div>
          </div>
        </div>
      </v-card>

      <!-- Total Expenses Card -->
      <v-card class="stat-card" elevation="0">
        <div class="stat-card-content">
          <div class="stat-icon expense">
            <i class="mdi mdi-cash-minus"></i>
          </div>
          <div class="stat-info">
            <div class="stat-label">إجمالي المصروفات</div>
            <div class="stat-value">{{ formatCurrency(financialStats.totalExpenses) }}</div>
            <div class="stat-change negative">
              <i class="mdi mdi-trending-down"></i>
              <span>-5% هذا الشهر</span>
            </div>
          </div>
        </div>
      </v-card>

      <!-- Completed Projects Card -->
      <v-card class="stat-card" elevation="0">
        <div class="stat-card-content">
          <div class="stat-icon success">
            <i class="mdi mdi-check-circle"></i>
          </div>
          <div class="stat-info">
            <div class="stat-label">مشاريع مكتملة</div>
            <div class="stat-value">{{ stats.completedProjects }}</div>
            <div class="stat-change positive">
              <i class="mdi mdi-trending-up"></i>
              <span>+18% هذا الشهر</span>
            </div>
          </div>
        </div>
      </v-card>

      <!-- Engineers Card -->
      <v-card class="stat-card" elevation="0">
        <div class="stat-card-content">
          <div class="stat-icon warning">
            <i class="mdi mdi-hard-hat"></i>
          </div>
          <div class="stat-info">
            <div class="stat-label">المهندسين</div>
            <div class="stat-value">{{ stats.engineersCount }}</div>
            <div class="stat-change positive">
              <i class="mdi mdi-trending-up"></i>
              <span>+3 هذا الشهر</span>
            </div>
          </div>
        </div>
      </v-card>

      <!-- Net Profit Card -->
      <v-card class="stat-card" elevation="0">
        <div class="stat-card-content">
          <div class="stat-icon profit">
            <i class="mdi mdi-chart-areaspline"></i>
          </div>
          <div class="stat-info">
            <div class="stat-label">صافي الربح</div>
            <div class="stat-value">{{ formatCurrency(financialStats.netProfit) }}</div>
            <div class="stat-change positive">
              <i class="mdi mdi-trending-up"></i>
              <span>+22% هذا الشهر</span>
            </div>
          </div>
        </div>
      </v-card>
    </div>

    <!-- Charts Section: Project Status & Tasks Summary -->
    <div class="charts-tasks-section">
      <!-- Project Status Chart -->
      <v-card class="chart-card" elevation="2">
        <v-card-title class="chart-title">
          <i class="mdi mdi-chart-donut"></i>
          حالة المشاريع
        </v-card-title>
        <v-card-text>
          <div class="chart-placeholder">
            <div class="donut-chart">
              <div class="donut-segment completed" :style="{ '--value': projectCompletedPercentage }"></div>
              <div class="chart-center">
                <span class="chart-center-value">{{ stats.totalProjects }}</span>
                <span class="chart-center-label">مشروع</span>
              </div>
            </div>
            <div class="chart-legend">
              <div class="legend-item">
                <span class="legend-color" style="background: #4CAF50;"></span>
                <span>مكتمل ({{ stats.completedProjects }})</span>
              </div>
              <div class="legend-item">
                <span class="legend-color" style="background: #2196F3;"></span>
                <span>قيد التنفيذ ({{ stats.activeProjects }})</span>
              </div>
              <div class="legend-item">
                <span class="legend-color" style="background: #FF9800;"></span>
                <span>معلق ({{ pendingProjects }})</span>
              </div>
            </div>
          </div>
        </v-card-text>
      </v-card>

      <!-- Tasks Summary Chart -->
      <v-card class="tasks-summary-chart" elevation="2">
        <v-card-title class="tasks-title">
          <i class="mdi mdi-clipboard-check-outline"></i>
          ملخص المهام
          <v-spacer></v-spacer>
          <v-chip color="primary" size="small">{{ tasks.length }} مهمة</v-chip>
        </v-card-title>
        <v-card-text>
          <div class="tasks-stats-container">
            <!-- Tasks Statistics Cards -->
            <div class="tasks-stats-grid">
              <div class="task-stat-card completed">
                <div class="task-stat-icon">
                  <i class="mdi mdi-check-circle"></i>
                </div>
                <div class="task-stat-info">
                  <div class="task-stat-number">{{ completedTasks }}</div>
                  <div class="task-stat-label">مكتملة</div>
                </div>
              </div>
              
              <div class="task-stat-card in-progress">
                <div class="task-stat-icon">
                  <i class="mdi mdi-progress-clock"></i>
                </div>
                <div class="task-stat-info">
                  <div class="task-stat-number">{{ inProgressTasks }}</div>
                  <div class="task-stat-label">قيد التنفيذ</div>
                </div>
              </div>
              
              <div class="task-stat-card pending">
                <div class="task-stat-icon">
                  <i class="mdi mdi-clock-outline"></i>
                </div>
                <div class="task-stat-info">
                  <div class="task-stat-number">{{ pendingTasks }}</div>
                  <div class="task-stat-label">معلقة</div>
                </div>
              </div>
            </div>

            <!-- Tasks Chart -->
            <div class="tasks-chart-wrapper">
              <div class="tasks-donut-chart">
                <svg viewBox="0 0 200 200" class="donut-svg">
                  <!-- Background circle -->
                  <circle cx="100" cy="100" r="80" fill="none" stroke="#e2e8f0" stroke-width="30"></circle>
                  
                  <!-- Completed tasks segment -->
                  <circle 
                    cx="100" 
                    cy="100" 
                    r="80" 
                    fill="none" 
                    stroke="#10b981" 
                    stroke-width="30"
                    :stroke-dasharray="`${completedPercentage * 5.024} 502.4`"
                    stroke-dashoffset="0"
                    transform="rotate(-90 100 100)"
                    class="donut-segment"
                  ></circle>
                  
                  <!-- In-progress tasks segment -->
                  <circle 
                    cx="100" 
                    cy="100" 
                    r="80" 
                    fill="none" 
                    stroke="#06b6d4" 
                    stroke-width="30"
                    :stroke-dasharray="`${inProgressPercentage * 5.024} 502.4`"
                    :stroke-dashoffset="`${-completedPercentage * 5.024}`"
                    transform="rotate(-90 100 100)"
                    class="donut-segment"
                  ></circle>
                  
                  <!-- Pending tasks segment -->
                  <circle 
                    cx="100" 
                    cy="100" 
                    r="80" 
                    fill="none" 
                    stroke="#f59e0b" 
                    stroke-width="30"
                    :stroke-dasharray="`${pendingPercentage * 5.024} 502.4`"
                    :stroke-dashoffset="`${-(completedPercentage + inProgressPercentage) * 5.024}`"
                    transform="rotate(-90 100 100)"
                    class="donut-segment"
                  ></circle>
                </svg>
                
                <div class="chart-center-content">
                  <div class="chart-total">{{ tasks.length }}</div>
                  <div class="chart-label">مهمة</div>
                </div>
              </div>
              
              <div class="tasks-legend">
                <div class="legend-item">
                  <span class="legend-dot completed"></span>
                  <span class="legend-text">مكتملة ({{ completedTasks }})</span>
                  <span class="legend-percentage">{{ completedPercentage }}%</span>
                </div>
                <div class="legend-item">
                  <span class="legend-dot in-progress"></span>
                  <span class="legend-text">قيد التنفيذ ({{ inProgressTasks }})</span>
                  <span class="legend-percentage">{{ inProgressPercentage }}%</span>
                </div>
                <div class="legend-item">
                  <span class="legend-dot pending"></span>
                  <span class="legend-text">معلقة ({{ pendingTasks }})</span>
                  <span class="legend-percentage">{{ pendingPercentage }}%</span>
                </div>
              </div>
            </div>
          </div>
        </v-card-text>
      </v-card>
    </div>

    <!-- Activity Section -->
    <div class="activity-section">
      <!-- Recent Activity -->
      <v-card class="recent-activity" elevation="2">
        <v-card-title class="activity-title">
          <i class="mdi mdi-history"></i>
          النشاطات الأخيرة
        </v-card-title>
        <v-card-text>
          <div class="activity-list">
            <div v-if="activities.length === 0" class="no-activities">
              <i class="mdi mdi-information-outline"></i>
              <span>لا توجد نشاطات حديثة</span>
            </div>
            <div
              v-for="activity in activities"
              :key="activity.id"
              class="activity-item"
            >
              <div class="activity-icon" :style="{ background: getActivityColor(activity.action) }">
                <i :class="getActivityIcon(activity.action)"></i>
              </div>
              <div class="activity-details">
                <div class="activity-text">{{ formatActivityText(activity) }}</div>
                <div class="activity-time">{{ activity.createdAt }}</div>
              </div>
            </div>
          </div>
          <!-- Activities Pagination -->
          <div v-if="activitiesPagination.totalPages > 1" class="activities-pagination">
            <v-pagination
              v-model="activitiesPagination.page"
              :length="activitiesPagination.totalPages"
              :total-visible="5"
              density="compact"
              @update:model-value="fetchActivities"
            ></v-pagination>
          </div>
        </v-card-text>
      </v-card>

    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import PageHeader from '../components/PageHeader.vue'
import { getDashboardStats, getFinancialStats, getTaskSummary, getActivities } from '@/api/dashboard'

// Loading state
const loading = ref(false)

// Add your reactive data and methods here
const stats = ref({
  totalProjects: 0,
  activeProjects: 0,
  completedProjects: 0,
  engineersCount: 0
})

const financialStats = ref({
  totalIncome: 0,
  totalExpenses: 0,
  netProfit: 0
})

const activities = ref([])
const activitiesPagination = ref({
  page: 1,
  limit: 10,
  total: 0,
  totalPages: 0
})

// Tasks data
const tasks = ref([])

// Project statistics
const pendingProjects = computed(() =>
  stats.value.totalProjects - stats.value.activeProjects - stats.value.completedProjects
)

const projectCompletedPercentage = computed(() =>
  stats.value.totalProjects > 0 ? Math.round((stats.value.completedProjects / stats.value.totalProjects) * 100) : 0
)

// Tasks statistics
const completedTasks = computed(() =>
  tasks.value.filter(t => t.status === 'completed').length
)

const inProgressTasks = computed(() =>
  tasks.value.filter(t => t.status === 'in-progress').length
)

const pendingTasks = computed(() =>
  tasks.value.filter(t => t.status === 'pending').length
)

const completedPercentage = computed(() =>
  tasks.value.length > 0 ? Math.round((completedTasks.value / tasks.value.length) * 100) : 0
)

const inProgressPercentage = computed(() =>
  tasks.value.length > 0 ? Math.round((inProgressTasks.value / tasks.value.length) * 100) : 0
)

const pendingPercentage = computed(() =>
  tasks.value.length > 0 ? Math.round((pendingTasks.value / tasks.value.length) * 100) : 0
)

// Fetch dashboard data
const fetchDashboardData = async () => {
  loading.value = true
  try {
    // Fetch stats
    const statsResponse = await getDashboardStats()
    if (statsResponse.success) {
      stats.value = {
        totalProjects: statsResponse.data.totalProjects || 0,
        activeProjects: statsResponse.data.activeProjects || 0,
        completedProjects: statsResponse.data.completedProjects || 0,
        engineersCount: statsResponse.data.totalEngineers || 0
      }
    }

    // Fetch financial stats
    const financialResponse = await getFinancialStats({ period: 'all' })
    if (financialResponse.success) {
      financialStats.value = {
        totalIncome: financialResponse.data.totalIncome || 0,
        totalExpenses: financialResponse.data.totalExpenses || 0,
        netProfit: (financialResponse.data.totalIncome || 0) - (financialResponse.data.totalExpenses || 0)
      }
    }

    // Fetch task summary for current month
    const currentMonth = new Date().toISOString().slice(0, 7)
    const taskResponse = await getTaskSummary(currentMonth)
    if (taskResponse.success && taskResponse.data) {
      // Update task stats from API
      const taskData = taskResponse.data
      tasks.value = taskData.tasks || tasks.value
    }

    // Fetch recent activities
    await fetchActivities()
  } catch (error) {
    console.error('Error fetching dashboard data:', error)
  } finally {
    loading.value = false
  }
}

// Fetch activities with pagination
const fetchActivities = async () => {
  try {
    const activitiesResponse = await getActivities({
      page: activitiesPagination.value.page,
      limit: activitiesPagination.value.limit
    })
    if (activitiesResponse.success && activitiesResponse.data) {
      activities.value = activitiesResponse.data.data || []
      activitiesPagination.value.total = activitiesResponse.data.total || 0
      activitiesPagination.value.totalPages = activitiesResponse.data.totalPages || 0
    }
  } catch (error) {
    console.error('Error fetching activities:', error)
  }
}

// Refresh data
const refreshData = () => {
  fetchDashboardData()
}

// Format currency for display
const formatCurrency = (amount) => {
  if (amount >= 1000000) {
    return (amount / 1000000).toFixed(1) + 'M'
  } else if (amount >= 1000) {
    return (amount / 1000).toFixed(0) + 'K'
  }
  return amount.toString()
}

// Activity helper functions
const getActivityColor = (action) => {
  const colors = {
    'create': '#2196F3',
    'update': '#FF9800',
    'delete': '#F44336',
    'complete': '#4CAF50',
    'add': '#2196F3'
  }
  return colors[action] || '#9E9E9E'
}

const getActivityIcon = (action) => {
  const icons = {
    'create': 'mdi mdi-plus',
    'update': 'mdi mdi-pencil',
    'delete': 'mdi mdi-delete',
    'complete': 'mdi mdi-check',
    'add': 'mdi mdi-plus'
  }
  return icons[action] || 'mdi mdi-information'
}

const formatActivityText = (activity) => {
  const actionTexts = {
    'create': 'تم إنشاء',
    'update': 'تم تحديث',
    'delete': 'تم حذف',
    'complete': 'تم إكمال',
    'add': 'تم إضافة'
  }
  const actionText = actionTexts[activity.action] || activity.action
  return `${actionText} ${activity.targetType} "${activity.targetName}"`
}

onMounted(() => {
  fetchDashboardData()
})
</script>

<style scoped>
.dashboard-container {
  padding: 32px;
  max-width: 1600px;
  margin: 0 auto;
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

.stat-icon.expense {
  background: linear-gradient(135deg, #ef4444 0%, #f87171 100%) !important;
  box-shadow: 0 6px 16px rgba(239, 68, 68, 0.4);
}

.stat-icon.success {
  background: linear-gradient(135deg, #10b981 0%, #34d399 100%) !important;
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

/* Charts and Tasks Section - Side by Side */
.charts-tasks-section {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(450px, 1fr));
  gap: 24px;
  margin-bottom: 32px;
}

.chart-card,
.tasks-summary-chart {
  border-radius: 16px !important;
  background: linear-gradient(135deg, #0a3d42 0%, #052428 100%) !important;
  border: 2px solid transparent !important;
  position: relative;
  overflow: hidden;
}

.chart-card::before,
.tasks-summary-chart::before {
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

.chart-title,
.tasks-title {
  font-size: 18px;
  font-weight: 600;
  color: rgba(255, 255, 255, 0.95) !important;
  display: flex;
  align-items: center;
  gap: 8px;
}

.chart-title i,
.tasks-title i {
  color: #06b6d4;
}

.chart-placeholder {
  padding: 24px;
  position: relative;
  z-index: 1;
}

/* Donut Chart */
.donut-chart {
  width: 200px;
  height: 200px;
  border-radius: 50%;
  background: conic-gradient(
    #10b981 0deg 162deg,
    #06b6d4 162deg 288deg,
    #f59e0b 288deg 360deg
  );
  position: relative;
  margin: 0 auto 24px;
}

.chart-center {
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  width: 140px;
  height: 140px;
  background: linear-gradient(135deg, #0a3d42 0%, #052428 100%);
  border-radius: 50%;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
}

.chart-center-value {
  font-size: 36px;
  font-weight: 700;
  background: linear-gradient(135deg, #ffffff 0%, #e0e7ff 100%);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background-clip: text;
}

.chart-center-label {
  font-size: 14px;
  color: rgba(255, 255, 255, 0.7);
}

.chart-legend {
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.legend-item {
  display: flex;
  align-items: center;
  gap: 8px;
  font-size: 14px;
  color: rgba(255, 255, 255, 0.8);
}

.legend-color {
  width: 16px;
  height: 16px;
  border-radius: 4px;
}

/* Bar Chart */
.bar-chart {
  display: flex;
  justify-content: space-around;
  align-items: flex-end;
  height: 250px;
  gap: 12px;
  padding: 20px 0;
}

.bar-group {
  flex: 1;
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 8px;
}

.bar {
  width: 100%;
  background: linear-gradient(180deg, #2196F3 0%, #00BCD4 100%);
  border-radius: 8px 8px 0 0;
  transition: all 0.3s ease;
  position: relative;
}

.bar:hover {
  transform: scaleY(1.05);
  filter: brightness(1.1);
}

.bar-label {
  font-size: 12px;
  color: #666;
  font-weight: 500;
}

/* Recent Activity */
.activity-section {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(450px, 1fr));
  gap: 24px;
  margin-bottom: 32px;
}

.recent-activity {
  border-radius: 16px !important;
  background: linear-gradient(135deg, #0a3d42 0%, #052428 100%) !important;
  border: 2px solid transparent !important;
  position: relative;
  overflow: hidden;
}

.recent-activity::before {
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

.activity-title {
  font-size: 18px;
  font-weight: 600;
  color: rgba(255, 255, 255, 0.95) !important;
  display: flex;
  align-items: center;
  gap: 8px;
  position: relative;
  z-index: 1;
}

.activity-title i {
  color: #06b6d4;
}

.activity-list {
  display: flex;
  flex-direction: column;
  gap: 16px;
  position: relative;
  z-index: 1;
}

.no-activities {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 8px;
  padding: 24px;
  color: rgba(255, 255, 255, 0.5);
  font-size: 14px;
}

.no-activities i {
  font-size: 20px;
}

.activity-item {
  display: flex;
  align-items: center;
  gap: 16px;
  padding: 12px;
  border-radius: 12px;
  transition: background 0.2s ease;
  background: linear-gradient(135deg, rgba(255, 255, 255, 0.05) 0%, rgba(255, 255, 255, 0.02) 100%);
  border: 1px solid rgba(255, 255, 255, 0.08);
}

.activity-item:hover {
  background: linear-gradient(135deg, rgba(255, 255, 255, 0.1) 0%, rgba(255, 255, 255, 0.05) 100%);
  border-color: rgba(6, 182, 212, 0.3);
}

.activity-icon {
  width: 48px;
  height: 48px;
  border-radius: 12px;
  display: flex;
  align-items: center;
  justify-content: center;
  color: white;
  font-size: 24px;
}

.activity-details {
  flex: 1;
}

.activity-text {
  font-size: 14px;
  font-weight: 500;
  color: rgba(255, 255, 255, 0.9);
  margin-bottom: 4px;
}

.activity-time {
  font-size: 12px;
  color: rgba(255, 255, 255, 0.5);
}

.activities-pagination {
  margin-top: 16px;
  display: flex;
  justify-content: center;
}

.activities-pagination :deep(.v-pagination__item) {
  color: rgba(255, 255, 255, 0.8);
}

.activities-pagination :deep(.v-pagination__item--is-active) {
  background: linear-gradient(135deg, #06b6d4 0%, #10b981 100%);
}

.tasks-stats-container {
  display: flex;
  flex-direction: column;
  gap: 32px;
}

.tasks-stats-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
  gap: 20px;
}

.task-stat-card {
  background: linear-gradient(135deg, rgba(255, 255, 255, 0.08) 0%, rgba(255, 255, 255, 0.05) 100%);
  border-radius: 16px;
  padding: 24px;
  display: flex;
  align-items: center;
  gap: 16px;
  border: 2px solid rgba(255, 255, 255, 0.1);
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
  position: relative;
  overflow: hidden;
  backdrop-filter: blur(10px);
}

.task-stat-card::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  width: 4px;
  height: 100%;
  transition: width 0.3s ease;
}

.task-stat-card.completed::before {
  background: linear-gradient(180deg, #10b981 0%, #059669 100%);
}

.task-stat-card.in-progress::before {
  background: linear-gradient(180deg, #06b6d4 0%, #0891b2 100%);
}

.task-stat-card.pending::before {
  background: linear-gradient(180deg, #f59e0b 0%, #d97706 100%);
}

.task-stat-card:hover {
  transform: translateY(-4px);
  box-shadow: 0 12px 24px rgba(6, 182, 212, 0.2);
  border-color: rgba(6, 182, 212, 0.4);
}

.task-stat-card:hover::before {
  width: 6px;
}

.task-stat-icon {
  width: 56px;
  height: 56px;
  border-radius: 14px;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 28px;
  color: white;
  flex-shrink: 0;
  transition: all 0.3s ease;
}

.task-stat-card.completed .task-stat-icon {
  background: linear-gradient(135deg, #10b981 0%, #059669 100%);
  box-shadow: 0 4px 12px rgba(16, 185, 129, 0.3);
}

.task-stat-card.in-progress .task-stat-icon {
  background: linear-gradient(135deg, #06b6d4 0%, #0891b2 100%);
  box-shadow: 0 4px 12px rgba(6, 182, 212, 0.3);
}

.task-stat-card.pending .task-stat-icon {
  background: linear-gradient(135deg, #f59e0b 0%, #d97706 100%);
  box-shadow: 0 4px 12px rgba(245, 158, 11, 0.3);
}

.task-stat-card:hover .task-stat-icon {
  transform: scale(1.1) rotate(5deg);
}

.task-stat-info {
  flex: 1;
}

.task-stat-number {
  font-size: 32px;
  font-weight: 800;
  background: linear-gradient(135deg, #ffffff 0%, #e0e7ff 100%);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background-clip: text;
  line-height: 1;
  margin-bottom: 6px;
}

.task-stat-label {
  font-size: 14px;
  color: rgba(255, 255, 255, 0.7);
  font-weight: 600;
}

.tasks-chart-wrapper {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 48px;
  flex-wrap: wrap;
}

.tasks-donut-chart {
  position: relative;
  width: 240px;
  height: 240px;
}

.donut-svg {
  width: 100%;
  height: 100%;
  filter: drop-shadow(0 4px 12px rgba(0, 0, 0, 0.1));
}

.donut-segment {
  transition: all 0.6s cubic-bezier(0.4, 0, 0.2, 1);
}

.chart-center-content {
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  text-align: center;
}

.chart-total {
  font-size: 48px;
  font-weight: 900;
  background: linear-gradient(135deg, #ffffff 0%, #e0e7ff 100%);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background-clip: text;
  line-height: 1;
  margin-bottom: 4px;
}

.chart-label {
  font-size: 16px;
  color: rgba(255, 255, 255, 0.7);
  font-weight: 600;
}

.tasks-legend {
  display: flex;
  flex-direction: column;
  gap: 16px;
}

.legend-item {
  display: flex;
  align-items: center;
  gap: 12px;
  padding: 12px 16px;
  background: linear-gradient(135deg, rgba(255, 255, 255, 0.08) 0%, rgba(255, 255, 255, 0.05) 100%);
  border-radius: 12px;
  border: 2px solid rgba(255, 255, 255, 0.1);
  transition: all 0.3s ease;
  backdrop-filter: blur(10px);
}

.legend-item:hover {
  border-color: rgba(6, 182, 212, 0.4);
  transform: translateX(-4px);
  box-shadow: 0 4px 12px rgba(6, 182, 212, 0.15);
}

.legend-dot {
  width: 16px;
  height: 16px;
  border-radius: 50%;
  flex-shrink: 0;
}

.legend-dot.completed {
  background: linear-gradient(135deg, #10b981 0%, #059669 100%);
  box-shadow: 0 2px 8px rgba(16, 185, 129, 0.3);
}

.legend-dot.in-progress {
  background: linear-gradient(135deg, #06b6d4 0%, #0891b2 100%);
  box-shadow: 0 2px 8px rgba(6, 182, 212, 0.3);
}

.legend-dot.pending {
  background: linear-gradient(135deg, #f59e0b 0%, #d97706 100%);
  box-shadow: 0 2px 8px rgba(245, 158, 11, 0.3);
}

.legend-text {
  flex: 1;
  font-size: 14px;
  color: rgba(255, 255, 255, 0.8);
  font-weight: 600;
}

.legend-percentage {
  font-size: 16px;
  font-weight: 800;
  background: linear-gradient(135deg, #ffffff 0%, #e0e7ff 100%);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background-clip: text;
}

/* Dashboard Header Custom Color */
.dashboard-header {
  background: linear-gradient(135deg, #018790 0%, #005461 100%) !important;
}

.dashboard-header::before {
  background: linear-gradient(135deg, #06b6d4 0%, #10b981 50%, #06b6d4 100%) !important;
}

/* Responsive Design */
@media (max-width: 768px) {
  .dashboard-container {
    padding: 16px;
  }

  .page-header {
    flex-direction: column;
    align-items: flex-start;
    gap: 16px;
  }

  .stats-grid {
    grid-template-columns: 1fr;
  }

  .charts-section {
    grid-template-columns: 1fr;
  }
}
</style>
