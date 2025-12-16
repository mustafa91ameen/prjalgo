<template>
  <v-card class="search-filter-card mb-8" elevation="2">
    <v-card-title class="d-flex align-center">
      <v-icon class="me-2" color="primary">mdi-filter</v-icon>
      البحث والفلترة
    </v-card-title>
    <v-card-text>
      <v-row>
        <v-col cols="12" md="4">
          <v-text-field
            :model-value="searchQuery"
            @update:model-value="$emit('update:searchQuery', $event)"
            label="البحث عن مديون"
            prepend-inner-icon="mdi-magnify"
            variant="outlined"
            clearable
            hide-details
            class="search-field"
          />
        </v-col>
        <v-col cols="12" md="3">
          <v-select
            :model-value="statusFilter"
            @update:model-value="$emit('update:statusFilter', $event)"
            label="حالة الدفع"
            :items="statusOptions"
            variant="outlined"
            hide-details
            clearable
            class="filter-field"
          />
        </v-col>
        <v-col cols="12" md="3">
          <v-select
            :model-value="amountFilter"
            @update:model-value="$emit('update:amountFilter', $event)"
            label="نطاق المبلغ"
            :items="amountOptions"
            variant="outlined"
            hide-details
            clearable
            class="filter-field"
          />
        </v-col>
        <v-col cols="12" md="2">
          <v-btn
            color="primary"
            variant="outlined"
            block
            class="reset-button"
            @click="$emit('reset')"
          >
            <v-icon class="me-1">mdi-refresh</v-icon>
            إعادة تعيين
          </v-btn>
        </v-col>
      </v-row>
    </v-card-text>
  </v-card>
</template>

<script setup>
defineProps({
  searchQuery: {
    type: String,
    default: ''
  },
  statusFilter: {
    type: String,
    default: ''
  },
  amountFilter: {
    type: String,
    default: ''
  },
  statusOptions: {
    type: Array,
    default: () => [
      { title: 'الكل', value: '' },
      { title: 'مدفوع', value: 'paid' },
      { title: 'قيد الانتظار', value: 'pending' },
      { title: 'متأخر', value: 'overdue' }
    ]
  },
  amountOptions: {
    type: Array,
    default: () => [
      { title: 'الكل', value: '' },
      { title: 'أقل من 1000', value: 'low' },
      { title: '1000 - 5000', value: 'medium' },
      { title: 'أكثر من 5000', value: 'high' }
    ]
  }
})

defineEmits(['update:searchQuery', 'update:statusFilter', 'update:amountFilter', 'reset'])
</script>

<style scoped>
.search-filter-card {
  border-radius: var(--radius-2xl) !important;
  background: var(--gradient-card-white) !important;
}

.search-field,
.filter-field {
  border-radius: var(--radius-xl) !important;
}

.reset-button {
  height: var(--space-14) !important;
  border-radius: var(--radius-xl) !important;
}
</style>
