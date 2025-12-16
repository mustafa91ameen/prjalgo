<template>
  <v-row class="mb-8">
    <v-col cols="12" md="3">
      <v-card class="stat-card pa-6 text-center" elevation="2">
        <div class="stat-icon mb-3">
          <v-icon size="64" color="info">mdi-account-group</v-icon>
        </div>
        <h3 class="text-h3 font-weight-bold text-info mb-2 stat-number-ltr">{{ totalDebtors }}</h3>
        <p class="text-subtitle-1 text-info mb-0">إجمالي المديونون</p>
      </v-card>
    </v-col>
    <v-col cols="12" md="3">
      <v-card class="stat-card pa-6 text-center" elevation="2">
        <div class="stat-icon mb-3">
          <v-icon size="64" color="error">mdi-currency-usd</v-icon>
        </div>
        <h3 class="text-h3 font-weight-bold text-error mb-2 stat-number-ltr">{{ formattedTotalDebt }}</h3>
        <p class="text-subtitle-1 text-error mb-0">إجمالي المديونية</p>
      </v-card>
    </v-col>
    <v-col cols="12" md="3">
      <v-card class="stat-card pa-6 text-center" elevation="2">
        <div class="stat-icon mb-3">
          <v-icon size="64" color="warning">mdi-clock-alert</v-icon>
        </div>
        <h3 class="text-h3 font-weight-bold text-warning mb-2 stat-number-ltr">{{ overdueCount }}</h3>
        <p class="text-subtitle-1 text-warning mb-0">متأخرين</p>
      </v-card>
    </v-col>
    <v-col cols="12" md="3">
      <v-card class="stat-card pa-6 text-center" elevation="2">
        <div class="stat-icon mb-3">
          <v-icon size="64" color="success">mdi-check-circle</v-icon>
        </div>
        <h3 class="text-h3 font-weight-bold text-success mb-2 stat-number-ltr">{{ paidCount }}</h3>
        <p class="text-subtitle-1 text-success mb-0">مدفوع</p>
      </v-card>
    </v-col>
  </v-row>
</template>

<script setup>
import { computed } from 'vue'
import { formatCurrency } from '@/utils/formatters'

const props = defineProps({
  totalDebtors: {
    type: Number,
    default: 0
  },
  totalDebt: {
    type: Number,
    default: 0
  },
  overdueCount: {
    type: Number,
    default: 0
  },
  paidCount: {
    type: Number,
    default: 0
  }
})

const formattedTotalDebt = computed(() => formatCurrency(props.totalDebt))
</script>

<style scoped>
.stat-card {
  position: relative;
  border-radius: var(--radius-3xl) !important;
  background: var(--gradient-card-white) !important;
  border: 1px solid var(--border-light);
  overflow: hidden;
  transition: all var(--transition-normal);
}

.stat-card:hover {
  transform: translateY(-8px);
  box-shadow: var(--shadow-xl) !important;
}

.stat-icon {
  display: flex;
  justify-content: center;
  align-items: center;
}

.stat-icon .v-icon {
  filter: drop-shadow(0 4px 8px var(--shadow-light));
}
</style>
