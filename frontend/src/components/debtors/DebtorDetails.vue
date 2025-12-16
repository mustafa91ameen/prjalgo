<template>
  <v-dialog v-model="dialogModel" max-width="500px">
    <v-card class="debtor-details-card">
      <v-card-title class="details-header">
        <v-icon class="me-2" color="white">mdi-account-details</v-icon>
        تفاصيل المديون
      </v-card-title>
      <v-card-text v-if="debtor" class="pa-6">
        <v-row>
          <v-col cols="12">
            <div class="text-center mb-4">
              <v-avatar
                :color="getStatusColor(debtor.status)"
                size="64"
              >
                <span class="text-white text-h4 font-weight-bold">
                  {{ debtor.name?.charAt(0) }}
                </span>
              </v-avatar>
              <h3 class="text-h5 font-weight-bold mt-2">{{ debtor.name }}</h3>
            </div>
          </v-col>
          <v-col cols="6">
            <strong>البريد الإلكتروني:</strong>
            <p>{{ debtor.email }}</p>
          </v-col>
          <v-col cols="6">
            <strong>رقم الهاتف:</strong>
            <p>{{ debtor.phone || 'غير محدد' }}</p>
          </v-col>
          <v-col cols="6">
            <strong>المبلغ المطلوب:</strong>
            <p class="text-h6 font-weight-bold text-error">
              {{ formatCurrency(debtor.amount) }}
            </p>
          </v-col>
          <v-col cols="6">
            <strong>تاريخ الاستحقاق:</strong>
            <p>{{ formatDate(debtor.dueDate) }}</p>
          </v-col>
          <v-col cols="12">
            <strong>الحالة:</strong>
            <v-chip
              :color="getStatusColor(debtor.status)"
              :variant="debtor.status === 'paid' ? 'flat' : 'tonal'"
              class="mt-1"
            >
              {{ getStatusText(debtor.status) }}
            </v-chip>
          </v-col>
          <v-col cols="12" v-if="debtor.notes">
            <strong>ملاحظات:</strong>
            <p>{{ debtor.notes }}</p>
          </v-col>
        </v-row>
      </v-card-text>
      <v-card-actions class="pa-4">
        <v-spacer />
        <v-btn
          color="primary"
          variant="elevated"
          @click="closeDialog"
        >
          إغلاق
        </v-btn>
      </v-card-actions>
    </v-card>
  </v-dialog>
</template>

<script setup>
import { computed } from 'vue'

const props = defineProps({
  modelValue: {
    type: Boolean,
    default: false
  },
  debtor: {
    type: Object,
    default: null
  }
})

const emit = defineEmits(['update:modelValue', 'close'])

const dialogModel = computed({
  get: () => props.modelValue,
  set: (value) => emit('update:modelValue', value)
})

const closeDialog = () => {
  dialogModel.value = false
  emit('close')
}

const formatCurrency = (amount) => {
  return new Intl.NumberFormat('ar-SA', {
    style: 'currency',
    currency: 'IQD'
  }).format(amount)
}

const formatDate = (date) => {
  if (!date) return ''
  return new Date(date).toLocaleDateString('ar-SA')
}

const getStatusColor = (status) => {
  const colors = {
    'overdue': 'error',
    'pending': 'warning',
    'paid': 'success'
  }
  return colors[status] || 'grey'
}

const getStatusText = (status) => {
  const texts = {
    'overdue': 'متأخر',
    'pending': 'معلق',
    'paid': 'مدفوع'
  }
  return texts[status] || 'غير محدد'
}
</script>

<style scoped>
.debtor-details-card {
  border-radius: var(--radius-2xl) !important;
  overflow: hidden !important;
}

.details-header {
  background: linear-gradient(135deg, #3b82f6 0%, #2563eb 50%, #1d4ed8 100%) !important;
  color: white !important;
  font-weight: 700 !important;
  padding: var(--space-5) 24px !important;
}
</style>
