<template>
  <v-card class="search-card mb-4" elevation="2">
    <v-card-text class="pa-4">
      <v-row class="align-center">
        <v-col cols="12" md="4">
          <v-text-field
            :model-value="searchQuery"
            @update:model-value="$emit('update:searchQuery', $event)"
            label="البحث في المصاريف الإدارية..."
            prepend-inner-icon="mdi-magnify"
            variant="outlined"
            density="comfortable"
            clearable
            hide-details
            class="search-field"
            style="background: #f5f5f5;"
          />
        </v-col>
        <v-col cols="12" md="2">
          <v-select
            :model-value="expenseType"
            @update:model-value="$emit('update:expenseType', $event)"
            :items="expenseTypeOptions"
            label="نوع المصروف"
            variant="outlined"
            density="comfortable"
            clearable
            hide-details
            class="filter-field"
          />
        </v-col>
        <v-col cols="12" md="2">
          <v-select
            :model-value="status"
            @update:model-value="$emit('update:status', $event)"
            :items="statusOptions"
            label="الحالة"
            variant="outlined"
            density="comfortable"
            clearable
            hide-details
            class="filter-field"
          />
        </v-col>
        <v-col cols="12" md="2">
          <v-btn
            color="error"
            variant="elevated"
            size="large"
            class="search-btn"
            @click="$emit('search')"
            block
          >
            بحث
          </v-btn>
        </v-col>
        <v-col cols="12" md="2">
          <v-btn
            v-if="canAdd"
            color="success"
            variant="elevated"
            size="large"
            class="add-expense-btn"
            @click="$emit('add')"
            block
          >
            <v-icon class="me-2">mdi-plus</v-icon>
            إضافة صنف جديد
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
  expenseType: {
    type: String,
    default: ''
  },
  status: {
    type: String,
    default: ''
  },
  expenseTypeOptions: {
    type: Array,
    default: () => [
      'جميع الأنواع',
      'مصاريف إدارية',
      'مصاريف مشاريع',
      'مرتبات وأجور',
      'إيجار ومرافق',
      'معدات وتجهيزات',
      'نقل ومواصلات',
      'تدريب وتطوير',
      'صيانة وإصلاح',
      'مواد مكتبية',
      'تسويق وإعلان',
      'أخرى'
    ]
  },
  statusOptions: {
    type: Array,
    default: () => [
      'جميع الحالات',
      'معتمد',
      'معلق',
      'مرفوض',
      'مسودة'
    ]
  },
  canAdd: {
    type: Boolean,
    default: true
  }
})

defineEmits(['update:searchQuery', 'update:expenseType', 'update:status', 'search', 'add'])
</script>

<style>
@import './styles/expenses.css';
</style>

<style scoped>
.search-card {
  border-radius: 16px !important;
  background: linear-gradient(145deg, #ffffff 0%, #f8fafc 100%) !important;
}

.search-field {
  border-radius: 12px !important;
}

.filter-field {
  border-radius: 12px !important;
}

.search-btn {
  border-radius: 12px !important;
  font-weight: 600 !important;
  height: 48px !important;
}

.add-expense-btn {
  border-radius: 12px !important;
  font-weight: 600 !important;
  height: 48px !important;
}
</style>
