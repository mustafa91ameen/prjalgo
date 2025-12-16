<template>
  <v-row class="mb-6 stats-row full-width">
    <v-col cols="12" sm="6" md="3" lg="3" xl="3">
      <v-card class="stat-card pa-6 pb-8 text-center hover-lift card-glow smooth-transition" elevation="2">
        <div class="stat-icon mb-3 icon-glow">
          <v-icon size="64" color="error">mdi-currency-usd</v-icon>
        </div>
        <h3 class="text-h3 font-weight-bold text-error mb-2 stat-number-ltr">{{ totalExpenses }}</h3>
        <p class="text-subtitle-1 text-error mb-0">إجمالي المصاريف</p>
      </v-card>
    </v-col>
    <v-col cols="12" sm="6" md="3" lg="3" xl="3">
      <v-card class="stat-card pa-6 pb-8 text-center hover-lift card-glow smooth-transition" elevation="2">
        <div class="stat-icon mb-3 icon-glow">
          <v-icon size="64" color="success">mdi-check-circle</v-icon>
        </div>
        <h3 class="text-h3 font-weight-bold text-success mb-2 stat-number-ltr">{{ activeExpenses }}</h3>
        <p class="text-subtitle-1 text-success mb-0">مصاريف نشطة</p>
      </v-card>
    </v-col>
    <v-col cols="12" sm="6" md="3" lg="3" xl="3">
      <v-card class="stat-card pa-6 pb-8 text-center hover-lift card-glow smooth-transition" elevation="2">
        <div class="stat-icon mb-3 icon-glow">
          <v-icon size="64" color="warning">mdi-clock-alert</v-icon>
        </div>
        <h3 class="text-h3 font-weight-bold text-warning mb-2 stat-number-ltr">{{ pendingExpenses }}</h3>
        <p class="text-subtitle-1 text-warning mb-0">في الانتظار</p>
      </v-card>
    </v-col>
    <v-col cols="12" sm="6" md="3" lg="3" xl="3">
      <v-card class="stat-card pa-6 pb-8 text-center hover-lift card-glow smooth-transition" elevation="2">
        <div class="stat-icon mb-3 icon-glow">
          <v-icon size="64" color="info">mdi-chart-line</v-icon>
        </div>
        <h3 class="text-h6 font-weight-bold text-info mb-2 stat-number-ltr-nowrap">{{ formattedTotalCost }}</h3>
        <p class="text-subtitle-1 text-info mb-0">إجمالي التكلفة</p>
      </v-card>
    </v-col>
  </v-row>
</template>

<script setup>
import { computed } from 'vue'
import { formatCurrency } from '@/utils/formatters'

const props = defineProps({
  totalExpenses: {
    type: Number,
    default: 0
  },
  activeExpenses: {
    type: Number,
    default: 0
  },
  pendingExpenses: {
    type: Number,
    default: 0
  },
  totalCost: {
    type: Number,
    default: 0
  }
})

const formattedTotalCost = computed(() => {
  return formatCurrency(props.totalCost)
})
</script>

<style scoped>
.stat-card {
  position: relative;
  border-radius: var(--radius-3xl) !important;
  background: var(--gradient-card-white) !important;
  border: 1px solid var(--border-light);
  overflow: hidden;
}

.stat-card::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  height: var(--space-1);
  background: var(--gradient-error-dark);
  border-radius: var(--radius-3xl) 24px 0 0;
}

.stat-card:nth-child(2)::before {
  background: var(--gradient-success-dark);
}

.stat-card:nth-child(3)::before {
  background: var(--gradient-warning-dark);
}

.stat-card:nth-child(4)::before {
  background: var(--gradient-info-cyan);
}

.hover-lift {
  transition: all var(--transition-normal);
}

.hover-lift:hover {
  transform: translateY(-8px);
  box-shadow: var(--shadow-xl) !important;
}

.stat-icon {
  display: flex;
  justify-content: center;
  align-items: center;
}

.icon-glow .v-icon {
  filter: drop-shadow(0 4px 8px var(--shadow-light));
}
</style>
