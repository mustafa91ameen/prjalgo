<template>
  <v-row class="mb-8">
    <v-col cols="12" md="3">
      <v-card class="stat-card pa-6 text-center" elevation="2">
        <div class="stat-icon mb-3">
          <v-icon size="64" color="info">mdi-account-group</v-icon>
        </div>
        <h3 class="text-h3 font-weight-bold text-info mb-2" style="font-family: 'Arial', 'Helvetica', sans-serif; direction: ltr;">{{ totalDebtors }}</h3>
        <p class="text-subtitle-1 text-info mb-0">إجمالي المديونون</p>
      </v-card>
    </v-col>
    <v-col cols="12" md="3">
      <v-card class="stat-card pa-6 text-center" elevation="2">
        <div class="stat-icon mb-3">
          <v-icon size="64" color="error">mdi-currency-usd</v-icon>
        </div>
        <h3 class="text-h3 font-weight-bold text-error mb-2" style="font-family: 'Arial', 'Helvetica', sans-serif; direction: ltr;">{{ formattedTotalDebt }}</h3>
        <p class="text-subtitle-1 text-error mb-0">إجمالي المديونية</p>
      </v-card>
    </v-col>
    <v-col cols="12" md="3">
      <v-card class="stat-card pa-6 text-center" elevation="2">
        <div class="stat-icon mb-3">
          <v-icon size="64" color="warning">mdi-clock-alert</v-icon>
        </div>
        <h3 class="text-h3 font-weight-bold text-warning mb-2" style="font-family: 'Arial', 'Helvetica', sans-serif; direction: ltr;">{{ overdueCount }}</h3>
        <p class="text-subtitle-1 text-warning mb-0">متأخرين</p>
      </v-card>
    </v-col>
    <v-col cols="12" md="3">
      <v-card class="stat-card pa-6 text-center" elevation="2">
        <div class="stat-icon mb-3">
          <v-icon size="64" color="success">mdi-check-circle</v-icon>
        </div>
        <h3 class="text-h3 font-weight-bold text-success mb-2" style="font-family: 'Arial', 'Helvetica', sans-serif; direction: ltr;">{{ paidCount }}</h3>
        <p class="text-subtitle-1 text-success mb-0">مدفوع</p>
      </v-card>
    </v-col>
  </v-row>
</template>

<script setup>
import { computed } from 'vue'

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

const formatCurrency = (amount) => {
  return new Intl.NumberFormat('ar-SA', {
    style: 'currency',
    currency: 'IQD',
    minimumFractionDigits: 0
  }).format(amount).replace('IQD', 'د.ع')
}

const formattedTotalDebt = computed(() => formatCurrency(props.totalDebt))
</script>

<style scoped>
.stat-card {
  position: relative;
  border-radius: 24px !important;
  background: linear-gradient(145deg, #ffffff 0%, #f8fafc 100%) !important;
  border: 1px solid rgba(0, 0, 0, 0.05);
  overflow: hidden;
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
}

.stat-card:hover {
  transform: translateY(-8px);
  box-shadow: 0 20px 40px rgba(0, 0, 0, 0.12) !important;
}

.stat-icon {
  display: flex;
  justify-content: center;
  align-items: center;
}

.stat-icon .v-icon {
  filter: drop-shadow(0 4px 8px rgba(0, 0, 0, 0.1));
}
</style>
