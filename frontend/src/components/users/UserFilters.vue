<template>
  <v-card class="filters-card mb-6" elevation="2">
    <v-card-text class="pa-4">
      <v-row>
        <v-col cols="12" md="3">
          <v-select
            :model-value="selectedRole"
            @update:model-value="$emit('update:selectedRole', $event)"
            :items="roleOptions"
            label="الدور"
            variant="outlined"
            density="compact"
            clearable
          />
        </v-col>
        <v-col cols="12" md="3">
          <v-select
            :model-value="selectedStatus"
            @update:model-value="$emit('update:selectedStatus', $event)"
            :items="statusOptions"
            label="الحالة"
            variant="outlined"
            density="compact"
            clearable
          />
        </v-col>
        <v-col cols="12" md="3">
          <v-select
            :model-value="selectedDepartment"
            @update:model-value="$emit('update:selectedDepartment', $event)"
            :items="departmentOptions"
            label="القسم"
            variant="outlined"
            density="compact"
            clearable
          />
        </v-col>
        <v-col cols="12" md="3">
          <v-btn
            color="primary"
            prepend-icon="mdi-filter"
            @click="$emit('apply-filters')"
            class="mt-2"
          >
            تطبيق الفلاتر
          </v-btn>
        </v-col>
      </v-row>
    </v-card-text>
  </v-card>
</template>

<script setup>
defineProps({
  selectedRole: {
    type: String,
    default: null
  },
  selectedStatus: {
    type: String,
    default: null
  },
  selectedDepartment: {
    type: String,
    default: null
  },
  roleOptions: {
    type: Array,
    default: () => []
  },
  statusOptions: {
    type: Array,
    default: () => []
  },
  departmentOptions: {
    type: Array,
    default: () => []
  }
})

defineEmits([
  'update:selectedRole',
  'update:selectedStatus',
  'update:selectedDepartment',
  'apply-filters'
])
</script>

<style scoped>
@import './styles/users.css';
</style>
