<template>
  <v-row class="mb-6 stats-row full-width">
    <v-col cols="12" sm="6" md="3" lg="3" xl="3">
      <v-card class="stat-card pa-6 pb-8 text-center hover-lift card-glow smooth-transition" elevation="2">
        <div class="stat-icon mb-3 icon-glow">
          <v-icon size="64" color="error">mdi-currency-usd</v-icon>
        </div>
        <h3 class="text-h3 font-weight-bold text-error mb-2" style="font-family: 'Arial', 'Helvetica', sans-serif; direction: ltr;">{{ totalExpenses }}</h3>
        <p class="text-subtitle-1 text-error mb-0">إجمالي المصاريف</p>
      </v-card>
    </v-col>
    <v-col cols="12" sm="6" md="3" lg="3" xl="3">
      <v-card class="stat-card pa-6 pb-8 text-center hover-lift card-glow smooth-transition" elevation="2">
        <div class="stat-icon mb-3 icon-glow">
          <v-icon size="64" color="success">mdi-check-circle</v-icon>
        </div>
        <h3 class="text-h3 font-weight-bold text-success mb-2" style="font-family: 'Arial', 'Helvetica', sans-serif; direction: ltr;">{{ activeExpenses }}</h3>
        <p class="text-subtitle-1 text-success mb-0">مصاريف نشطة</p>
      </v-card>
    </v-col>
    <v-col cols="12" sm="6" md="3" lg="3" xl="3">
      <v-card class="stat-card pa-6 pb-8 text-center hover-lift card-glow smooth-transition" elevation="2">
        <div class="stat-icon mb-3 icon-glow">
          <v-icon size="64" color="warning">mdi-clock-alert</v-icon>
        </div>
        <h3 class="text-h3 font-weight-bold text-warning mb-2" style="font-family: 'Arial', 'Helvetica', sans-serif; direction: ltr;">{{ pendingExpenses }}</h3>
        <p class="text-subtitle-1 text-warning mb-0">في الانتظار</p>
      </v-card>
    </v-col>
    <v-col cols="12" sm="6" md="3" lg="3" xl="3">
      <v-card class="stat-card pa-6 pb-8 text-center hover-lift card-glow smooth-transition" elevation="2">
        <div class="stat-icon mb-3 icon-glow">
          <v-icon size="64" color="info">mdi-chart-line</v-icon>
        </div>
        <h3 class="text-h6 font-weight-bold text-info mb-2" style="font-family: 'Arial', 'Helvetica', sans-serif; direction: ltr; white-space: nowrap;">{{ formattedTotalCost }}</h3>
        <p class="text-subtitle-1 text-info mb-0">إجمالي التكلفة</p>
      </v-card>
    </v-col>
  </v-row>
</template>

<script setup>
import { computed } from 'vue'

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

const formatCurrency = (amount) => {
  return new Intl.NumberFormat('ar-SA', {
    style: 'currency',
    currency: 'IQD',
    minimumFractionDigits: 0
  }).format(amount).replace('IQD', 'د.ع')
}

const formattedTotalCost = computed(() => {
  return formatCurrency(props.totalCost)
})
</script>

<style>
@import './styles/expenses.css';
</style>

<style scoped>
.stat-card {
  position: relative;
  border-radius: 24px !important;
  background: linear-gradient(145deg, #ffffff 0%, #f8fafc 100%) !important;
  border: 1px solid rgba(0, 0, 0, 0.05);
  overflow: hidden;
}

.stat-card::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  height: 4px;
  background: linear-gradient(90deg, #ef4444, #dc2626);
  border-radius: 24px 24px 0 0;
}

.stat-card:nth-child(2)::before {
  background: linear-gradient(90deg, #10b981, #059669);
}

.stat-card:nth-child(3)::before {
  background: linear-gradient(90deg, #f59e0b, #d97706);
}

.stat-card:nth-child(4)::before {
  background: linear-gradient(90deg, #3b82f6, #0ea5e9);
}

.hover-lift {
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
}

.hover-lift:hover {
  transform: translateY(-8px);
  box-shadow: 0 20px 40px rgba(0, 0, 0, 0.12) !important;
}

.stat-icon {
  display: flex;
  justify-content: center;
  align-items: center;
}

.icon-glow .v-icon {
  filter: drop-shadow(0 4px 8px rgba(0, 0, 0, 0.1));
}
</style>
