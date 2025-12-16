<template>
  <v-card class="search-card mb-4" elevation="2">
    <v-card-text class="pa-4">
      <v-row class="align-center">
        <v-col cols="12" md="3">
          <v-text-field
            :model-value="searchQuery"
            @update:model-value="$emit('update:searchQuery', $event)"
            label="البحث في الموظفين..."
            prepend-inner-icon="mdi-magnify"
            variant="outlined"
            density="comfortable"
            clearable
            hide-details
            class="search-field bg-filter-input"
          />
        </v-col>
        <v-col cols="12" md="2">
          <v-select
            :model-value="selectedDepartment"
            @update:model-value="$emit('update:selectedDepartment', $event)"
            :items="departmentOptions"
            label="القسم"
            variant="outlined"
            density="comfortable"
            clearable
            hide-details
            class="black-list"
          />
        </v-col>
        <v-col cols="12" md="2">
          <v-select
            :model-value="selectedStatus"
            @update:model-value="$emit('update:selectedStatus', $event)"
            :items="statusOptions"
            label="الحالة"
            variant="outlined"
            density="comfortable"
            clearable
            hide-details
          />
        </v-col>
        <v-col cols="12" md="2">
          <v-select
            :model-value="selectedPosition"
            @update:model-value="$emit('update:selectedPosition', $event)"
            :items="positionOptions"
            label="المنصب"
            variant="outlined"
            density="comfortable"
            clearable
            hide-details
          />
        </v-col>
        <v-col cols="12" md="1">
          <v-btn
            color="primary"
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
            إضافة موظف
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
  selectedDepartment: {
    type: String,
    default: ''
  },
  selectedStatus: {
    type: String,
    default: ''
  },
  selectedPosition: {
    type: String,
    default: ''
  },
  departmentOptions: {
    type: Array,
    default: () => []
  },
  statusOptions: {
    type: Array,
    default: () => []
  },
  positionOptions: {
    type: Array,
    default: () => []
  },
  canAdd: {
    type: Boolean,
    default: true
  }
})

defineEmits([
  'update:searchQuery',
  'update:selectedDepartment',
  'update:selectedStatus',
  'update:selectedPosition',
  'search',
  'add'
])
</script>

<style scoped>
@import './styles/human-resources.css';
</style>
