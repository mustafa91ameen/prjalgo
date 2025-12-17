<template>
  <v-row class="mb-6 stats-row" no-gutters>
    <v-col cols="12" sm="6" md="2" lg="2" xl="2" class="pa-2">
      <div class="modern-stat-card stat-card-info hover-lift">
        <div class="stat-card-background"></div>
        <div class="project-stat-content">
          <div class="project-icon-wrapper">
            <v-icon size="32" class="project-stat-icon">mdi-folder-multiple</v-icon>
          </div>
          <div class="project-stat-value">{{ totalProjects || 0 }}</div>
          <div class="project-stat-label">إجمالي المشاريع</div>
        </div>
      </div>
    </v-col>
    
    <v-col cols="12" sm="6" md="2" lg="2" xl="2" class="pa-2">
      <div class="modern-stat-card stat-card-success hover-lift">
        <div class="stat-card-background"></div>
        <div class="project-stat-content">
          <div class="project-icon-wrapper">
            <v-icon size="32" class="project-stat-icon">mdi-check-circle</v-icon>
          </div>
          <div class="project-stat-value">{{ activeProjects || 0 }}</div>
          <div class="project-stat-label">مشاريع نشطة</div>
        </div>
      </div>
    </v-col>
    
    <v-col cols="12" sm="6" md="2" lg="2" xl="2" class="pa-2">
      <div class="modern-stat-card stat-card-warning hover-lift">
        <div class="stat-card-background"></div>
        <div class="project-stat-content">
          <div class="project-icon-wrapper">
            <v-icon size="32" class="project-stat-icon">mdi-clock-alert</v-icon>
          </div>
          <div class="project-stat-value">{{ pendingProjects || 0 }}</div>
          <div class="project-stat-label">في الانتظار</div>
        </div>
      </div>
    </v-col>
    
    <v-col cols="12" sm="6" md="3" lg="3" xl="3" class="pa-2">
      <div class="modern-stat-card stat-card-primary hover-lift">
        <div class="stat-card-background"></div>
        <div class="project-stat-content">
          <div class="project-icon-wrapper">
            <v-icon size="32" class="project-stat-icon">mdi-currency-usd</v-icon>
          </div>
          <div class="project-stat-value">{{ formattedBudget }}</div>
          <div class="project-stat-label">إجمالي الميزانية</div>
        </div>
      </div>
    </v-col>
    
    <v-col cols="12" sm="6" md="3" lg="3" xl="3" class="pa-2">
      <div class="modern-stat-card stat-card-info hover-lift">
        <div class="stat-card-background"></div>
        <div class="project-stat-content">
          <div class="project-icon-wrapper">
            <v-icon size="32" class="project-stat-icon">mdi-chart-line</v-icon>
          </div>
          <div class="project-stat-value">{{ averageProgress || 0 }}%</div>
          <div class="project-stat-label">متوسط التقدم</div>
        </div>
      </div>
    </v-col>
  </v-row>
</template>

<script setup>
import { computed } from 'vue'
import { formatCurrencyFull } from '@/utils/formatters'

const props = defineProps({
  totalProjects: {
    type: Number,
    default: 0
  },
  activeProjects: {
    type: Number,
    default: 0
  },
  pendingProjects: {
    type: Number,
    default: 0
  },
  totalBudget: {
    type: Number,
    default: 0
  },
  averageProgress: {
    type: Number,
    default: 0
  }
})

const formattedBudget = computed(() => {
  return formatCurrencyFull(props.totalBudget, 'IQD', 'en-US') || '0 د.ع'
})
</script>

<style>
/* Import shared project styles */
@import './styles/projects.css';
</style>

<style scoped>
/* Component-specific styles for ProjectStats */

/* Stats row layout */
.stats-row {
  width: 100%;
  justify-content: center;
  margin: 0 auto;
  gap: 0;
  display: flex;
  flex-wrap: wrap;
  align-items: stretch;
  padding: 0;
}

.stats-row .v-col {
  margin: 0 !important;
  padding: var(--space-2) !important;
  box-sizing: border-box !important;
}

/* Responsive adjustments specific to stats row */
@media (max-width: 960px) {
  .stats-row {
    gap: 0;
    padding: 0;
  }

  .stats-row .v-col {
    padding: var(--space-1-5) !important;
  }
}

@media (max-width: 600px) {
  .stats-row .v-col {
    padding: var(--space-1) !important;
  }
}
</style>
