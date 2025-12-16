<template>
  <div>
    <v-container fluid class="pa-0">
        <!-- Hero Section -->
        <div class="page-header glass-effect gradient-animation">
          <div class="header-top-content">
            <h1 class="page-title">مرحباً بك في نظام إدارة المشاريع</h1>
            <span class="page-icon">🏠</span>
          </div>
          <p class="page-subtitle">نظام متكامل لإدارة مشاريعك وفريقك ومواردك المالية</p>
        </div>

        <!-- Quick Actions Grid -->
        <div class="quick-actions-container">
            <div class="quick-actions-grid">
              <v-card
                v-for="action in quickActions"
                :key="action.title"
                class="quick-action-card"
                @click="$router.push(action.route)"
              >
                <v-card-text class="text-center pa-6">
                  <v-icon :color="action.color" size="72" class="mb-4">{{ action.icon }}</v-icon>
                  <h3 class="text-h6 font-weight-bold mb-2 text-primary">{{ action.title }}</h3>
                  <p class="text-body-2 text-secondary">{{ action.description }}</p>
                </v-card-text>
              </v-card>
            </div>
        </div>

        <!-- Statistics Section -->
        <div class="stats-section pa-6">
          <h2 class="text-h4 font-weight-bold text-center mb-6">إحصائيات المشاريع</h2>
          <v-row>
            <v-col cols="12" md="3">
              <v-card class="stat-card pa-6 text-center" elevation="2">
                <div class="stat-icon mb-3">
                  <v-icon size="84" color="error">mdi-currency-usd</v-icon>
                </div>
                <h3 class="text-h3 font-weight-bold text-error mb-2">{{ projectStats.totalBudget }}</h3>
                <p class="text-subtitle-1 text-error mb-0">إجمالي الميزانية</p>
              </v-card>
            </v-col>
            <v-col cols="12" md="3">
              <v-card class="stat-card pa-6 text-center" elevation="2">
                <div class="stat-icon mb-3">
                  <v-icon size="84" color="info">mdi-folder-multiple</v-icon>
                </div>
                <h3 class="text-h3 font-weight-bold text-info mb-2">{{ projectStats.totalProjects }}</h3>
                <p class="text-subtitle-1 text-info mb-0">إجمالي المشاريع</p>
              </v-card>
            </v-col>
            <v-col cols="12" md="3">
              <v-card class="stat-card pa-6 text-center" elevation="2">
                <div class="stat-icon mb-3">
                  <v-icon size="84" color="warning">mdi-check-circle</v-icon>
                </div>
                <h3 class="text-h3 font-weight-bold text-warning mb-2">{{ projectStats.completedProjects }}</h3>
                <p class="text-subtitle-1 text-warning mb-0">مشاريع مكتملة</p>
              </v-card>
            </v-col>
            <v-col cols="12" md="3">
              <v-card class="stat-card pa-6 text-center" elevation="2">
                <div class="stat-icon mb-3">
                  <v-icon size="84" color="success">mdi-clock-alert</v-icon>
                </div>
                <h3 class="text-h3 font-weight-bold text-success mb-2">{{ projectStats.activeProjects }}</h3>
                <p class="text-subtitle-1 text-success mb-0">مشاريع نشطة</p>
              </v-card>
            </v-col>
            <v-col cols="12" md="3">
              <v-card class="stat-card pa-6 text-center" elevation="2">
                <div class="stat-icon mb-3">
                  <v-icon size="84" color="primary">mdi-account-hard-hat</v-icon>
                </div>
                <h3 class="text-h3 font-weight-bold text-primary mb-2">{{ projectStats.totalEngineers }}</h3>
                <p class="text-subtitle-1 text-primary mb-0">المهندسين</p>
              </v-card>
            </v-col>
          </v-row>
        </div>
      </v-container>
  </div>
</template>

<script setup>
import { ref, computed } from 'vue'
import { formatCurrency } from '@/utils/formatters'

// Sample data - in real app, this would come from API
const incomeData = ref(175000) // Total income
const expenseData = ref(103000) // Total expenses

const projectStats = ref({
  totalProjects: 24,
  completedProjects: 18,
  activeProjects: 6,
  totalEngineers: 12
})

// الإجراءات السريعة
const quickActions = ref([
  {
    title: 'إدارة المشاريع',
    description: 'إدارة وتتبع جميع المشاريع والمهام',
    icon: 'mdi-folder-multiple',
    color: 'primary',
    route: '/project-management'
  },
  {
    title: 'المهندسين',
    description: 'إدارة فريق المهندسين والمطورين',
    icon: 'mdi-account-hard-hat',
    color: 'success',
    route: '/engineers'
  },
  {
    title: 'المصاريف الإدارية',
    description: 'إدارة وتتبع المصاريف الإدارية',
    icon: 'mdi-currency-usd',
    color: 'error',
    route: '/expenses'
  },
  {
    title: 'التصنيفات',
    description: 'إدارة وتنظيم تصنيفات المشاريع والمهام',
    icon: 'mdi-tag-multiple',
    color: 'info',
    route: '/categories'
  },
  {
    title: 'المديونون',
    description: 'إدارة حسابات المديونون والمستحقات',
    icon: 'mdi-credit-card',
    color: 'error',
    route: '/debtors'
  },
  {
    title: 'المخزون',
    description: 'إدارة وتتبع جميع عناصر المخزون والمواد',
    icon: 'mdi-package-variant',
    color: 'purple',
    route: '/inventory'
  },
  {
    title: 'المبيعات',
    description: 'إدارة وتتبع جميع عمليات البيع والمبيعات',
    icon: 'mdi-cash-multiple',
    color: 'success',
    route: '/sales'
  },
  {
    title: 'المشتريات',
    description: 'إدارة وتتبع جميع عمليات الشراء والمشتريات',
    icon: 'mdi-cart',
    color: 'error',
    route: '/purchases'
  },
  {
    title: 'الموارد البشرية',
    description: 'إدارة وتتبع جميع الموظفين والموارد البشرية',
    icon: 'mdi-account-group',
    color: 'info',
    route: '/human-resources'
  }
])


// Computed properties
const totalIncome = computed(() => incomeData.value)
const totalExpenses = computed(() => expenseData.value)
const netProfit = computed(() => totalIncome.value - totalExpenses.value)

</script>


<style scoped>
/* Import page styles - scoped to this component only */
@import './styles/index.css';
</style>
