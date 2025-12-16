<template>
  <v-row class="mb-6 stats-row" no-gutters>
    <v-col cols="12" sm="6" md="2" lg="2" xl="2" class="pa-2">
      <v-card class="stat-card pa-4 pb-6 text-center hover-lift card-glow smooth-transition" elevation="2">
        <div class="stat-icon mb-3 icon-glow">
          <v-icon size="48" color="info">mdi-folder-multiple</v-icon>
        </div>
        <h3 class="text-h4 font-weight-bold text-info mb-2 stat-number-ltr">{{ totalProjects || 0 }}</h3>
        <p class="text-caption text-info mb-0">إجمالي المشاريع</p>
      </v-card>
    </v-col>
    <v-col cols="12" sm="6" md="2" lg="2" xl="2" class="pa-2">
      <v-card class="stat-card pa-4 pb-6 text-center hover-lift card-glow smooth-transition" elevation="2">
        <div class="stat-icon mb-3 icon-glow">
          <v-icon size="48" color="success">mdi-check-circle</v-icon>
        </div>
        <h3 class="text-h4 font-weight-bold text-success mb-2 stat-number-ltr">{{ activeProjects || 0 }}</h3>
        <p class="text-caption text-success mb-0">مشاريع نشطة</p>
      </v-card>
    </v-col>
    <v-col cols="12" sm="6" md="2" lg="2" xl="2" class="pa-2">
      <v-card class="stat-card pa-4 pb-6 text-center hover-lift card-glow smooth-transition" elevation="2">
        <div class="stat-icon mb-3 icon-glow">
          <v-icon size="48" color="warning">mdi-clock-alert</v-icon>
        </div>
        <h3 class="text-h4 font-weight-bold text-warning mb-2 stat-number-ltr">{{ pendingProjects || 0 }}</h3>
        <p class="text-caption text-warning mb-0">في الانتظار</p>
      </v-card>
    </v-col>
    <v-col cols="12" sm="6" md="3" lg="3" xl="3" class="pa-2">
      <v-card class="stat-card pa-4 pb-6 text-center hover-lift card-glow smooth-transition" elevation="2">
        <div class="stat-icon mb-3 icon-glow">
          <v-icon size="48" color="error">mdi-currency-usd</v-icon>
        </div>
        <h3 class="text-h6 font-weight-bold text-error mb-2 stat-number-ltr-nowrap">{{ formattedBudget }}</h3>
        <p class="text-caption text-error mb-0">إجمالي الميزانية</p>
      </v-card>
    </v-col>
    <v-col cols="12" sm="6" md="3" lg="3" xl="3" class="pa-2">
      <v-card class="stat-card pa-4 pb-6 text-center hover-lift card-glow smooth-transition" elevation="2">
        <div class="stat-icon mb-3 icon-glow">
          <v-icon size="48" color="primary">mdi-chart-line</v-icon>
        </div>
        <h3 class="text-h4 font-weight-bold text-primary mb-2 stat-number-ltr">{{ averageProgress || 0 }}%</h3>
        <p class="text-caption text-primary mb-0">متوسط التقدم</p>
      </v-card>
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

.stats-row .stat-card {
  width: 100% !important;
  margin: 0 !important;
  box-sizing: border-box !important;
}

/* Enhanced stat card hover effects */
.stat-card::after {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: radial-gradient(circle at 50% 0%, rgba(255, 255, 255, 0.1) 0%, transparent 70%);
  border-radius: var(--radius-3xl);
  z-index: 1;
  pointer-events: none;
}

.stat-card .v-icon {
  filter: drop-shadow(0 6px 12px var(--shadow-medium));
  transition: all var(--transition-normal);
  font-size: var(--font-size-6xl) !important;
  width: var(--space-14) !important;
  height: var(--space-14) !important;
}

.stat-card:hover .v-icon {
  transform: scale(1.15) rotate(8deg);
  filter: drop-shadow(0 12px 24px var(--shadow-dark));
}

/* Color-specific effects for each stat card */
.stat-card:nth-child(1)::before {
  background: var(--gradient-primary-deep);
}

.stat-card:nth-child(1):hover {
  box-shadow: 0 25px 50px var(--shadow-primary-glow), 0 8px 16px var(--shadow-info-glow) !important;
}

.stat-card:nth-child(2)::before {
  background: var(--gradient-success-deep);
}

.stat-card:nth-child(2):hover {
  box-shadow: 0 25px 50px var(--shadow-success-glow), 0 8px 16px var(--shadow-success-glow) !important;
}

.stat-card:nth-child(3)::before {
  background: var(--gradient-warning-deep);
}

.stat-card:nth-child(3):hover {
  box-shadow: 0 25px 50px var(--shadow-warning-glow), 0 8px 16px var(--shadow-warning-glow) !important;
}

.stat-card:nth-child(4)::before {
  background: var(--gradient-error-deep);
}

.stat-card:nth-child(4):hover {
  box-shadow: 0 25px 50px var(--shadow-error-glow), 0 8px 16px var(--shadow-error-glow) !important;
}

.stat-card:nth-child(5)::before {
  background: var(--gradient-info-deep);
}

.stat-card:nth-child(5):hover {
  box-shadow: 0 25px 50px var(--shadow-info-glow), 0 8px 16px var(--shadow-info-glow) !important;
}

/* Responsive adjustments specific to stats row */
@media (max-width: var(--space-96)) {
  .stats-row {
    gap: 0;
    padding: 0;
  }

  .stats-row .v-col {
    padding: var(--space-1-5) !important;
  }
}

@media (max-width: var(--space-96)) {
  .stats-row .v-col {
    padding: var(--space-1) !important;
  }

  .stat-card .v-icon {
    font-size: var(--font-size-4xl-plus) !important;
    width: var(--space-10) !important;
    height: var(--space-10) !important;
  }
}

@media (max-width: var(--space-96)) {
  .stats-row .v-col {
    padding: var(--space-0-5) !important;
  }

  .stat-card .v-icon {
    font-size: var(--font-size-3xl-plus) !important;
    width: var(--space-8) !important;
    height: var(--space-8) !important;
  }
}
</style>
