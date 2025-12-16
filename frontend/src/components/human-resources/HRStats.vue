<template>
  <v-row class="mb-6 stats-row full-width">
    <v-col cols="12" sm="6" md="4" lg="4" xl="4">
      <v-card class="stat-card pa-3 pb-4 text-center hover-lift card-glow smooth-transition" elevation="2">
        <div class="stat-icon mb-2 icon-glow">
          <v-icon size="48" color="primary">mdi-account-group</v-icon>
        </div>
        <h3 class="text-h3 font-weight-bold text-primary mb-2 stat-number stat-number-ltr">{{ formatNumber(totalEmployees) }}</h3>
        <p class="text-subtitle-1 text-primary mb-0">إجمالي الموظفين</p>
      </v-card>
    </v-col>
    <v-col cols="12" sm="6" md="4" lg="4" xl="4">
      <v-card class="stat-card pa-3 pb-4 text-center hover-lift card-glow smooth-transition" elevation="2">
        <div class="stat-icon mb-2 icon-glow">
          <v-icon size="48" color="success">mdi-account-check</v-icon>
        </div>
        <h3 class="text-h3 font-weight-bold text-success mb-2 stat-number stat-number-ltr">{{ formatNumber(activeEmployees) }}</h3>
        <p class="text-subtitle-1 text-success mb-0">موظفين نشطين</p>
      </v-card>
    </v-col>
    <v-col cols="12" sm="6" md="4" lg="4" xl="4">
      <v-card class="stat-card pa-3 pb-4 text-center hover-lift card-glow smooth-transition" elevation="2">
        <div class="stat-icon mb-2 icon-glow">
          <v-icon size="48" color="info">mdi-office-building</v-icon>
        </div>
        <h3 class="text-h3 font-weight-bold text-info mb-2 stat-number stat-number-ltr">{{ formatNumber(totalDepartments) }}</h3>
        <p class="text-subtitle-1 text-info mb-0">إجمالي الأقسام</p>
      </v-card>
    </v-col>
    <v-col cols="12" sm="6" md="4" lg="4" xl="4">
      <v-card class="stat-card pa-3 pb-4 text-center hover-lift card-glow smooth-transition" elevation="2">
        <div class="stat-icon mb-2 icon-glow">
          <v-icon size="48" color="warning">mdi-calendar-clock</v-icon>
        </div>
        <h3 class="text-h3 font-weight-bold text-warning mb-2 stat-number stat-number-ltr">{{ formatNumber(onLeave) }}</h3>
        <p class="text-subtitle-1 text-warning mb-0">في إجازة</p>
      </v-card>
    </v-col>
    <v-col cols="12" sm="6" md="4" lg="4" xl="4">
      <v-card class="stat-card pa-3 pb-4 text-center hover-lift card-glow smooth-transition" elevation="2">
        <div class="stat-icon mb-2 icon-glow">
          <v-icon size="48" color="error">mdi-currency-usd</v-icon>
        </div>
        <h3 class="text-h3 font-weight-bold text-error mb-2 stat-number stat-number-ltr text-base-plus">{{ formatCurrency(totalSalaries) }}</h3>
        <p class="text-subtitle-1 text-error mb-0">إجمالي الرواتب</p>
      </v-card>
    </v-col>
    <v-col cols="12" sm="6" md="4" lg="4" xl="4">
      <v-card class="stat-card pa-3 pb-4 text-center hover-lift card-glow smooth-transition" elevation="2">
        <div class="stat-icon mb-2 icon-glow average-salary-icon">
          <v-icon size="48" class="text-purple">mdi-chart-line</v-icon>
        </div>
        <h3 class="text-h3 font-weight-bold mb-2 stat-number average-salary-value stat-number-ltr text-base-plus">{{ formatCurrency(averageSalary) }}</h3>
        <p class="text-subtitle-1 mb-0 average-salary-label">متوسط الراتب</p>
      </v-card>
    </v-col>
  </v-row>
</template>

<script setup>
import { formatCurrency, formatNumber } from '@/utils/formatters'

defineProps({
  totalEmployees: {
    type: Number,
    default: 0
  },
  activeEmployees: {
    type: Number,
    default: 0
  },
  totalDepartments: {
    type: Number,
    default: 0
  },
  onLeave: {
    type: Number,
    default: 0
  },
  totalSalaries: {
    type: Number,
    default: 0
  },
  averageSalary: {
    type: Number,
    default: 0
  }
})
</script>

<style scoped>
@import './styles/human-resources.css';
</style>
