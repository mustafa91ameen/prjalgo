<template>
  <v-card class="data-table-card" elevation="2">
    <v-card-title class="table-title indigo-title">
      <span class="title-text">قائمة الموظفين</span>
    </v-card-title>

    <v-data-table
      :headers="headers"
      :items="employees"
      :search="searchQuery"
      class="project-table"
      :items-per-page="10"
      :loading="loading"
      hover
      no-data-text="لا توجد بيانات"
    >
      <!-- Serial Number Column -->
      <template #item.serial="{ index }">
        <span class="serial-number">{{ index + 1 }}</span>
      </template>

      <!-- Employee Name Column -->
      <template #item.name="{ item }">
        <span class="project-name">{{ item.name }}</span>
      </template>

      <!-- ID Number Column -->
      <template #item.idNumber="{ item }">
        <span class="date-text">{{ item.idNumber || '-' }}</span>
      </template>

      <!-- Department Column -->
      <template #item.department="{ item }">
        <v-chip class="category-chip" size="small">
          {{ item.department }}
        </v-chip>
      </template>

      <!-- Position Column -->
      <template #item.position="{ item }">
        <span class="project-name">{{ item.position }}</span>
      </template>

      <!-- Phone Column -->
      <template #item.phone="{ item }">
        <span class="date-text">{{ item.phone }}</span>
      </template>

      <!-- Email Column -->
      <template #item.email="{ item }">
        <span class="project-name">{{ item.email }}</span>
      </template>

      <!-- Salary Column -->
      <template #item.salary="{ item }">
        <span class="cost-text">{{ formatCurrency(item.salary) }}</span>
      </template>

      <!-- Hire Date Column -->
      <template #item.hireDate="{ item }">
        <span class="date-text">{{ item.hireDate }}</span>
      </template>

      <!-- Status Column -->
      <template #item.status="{ item }">
        <v-chip class="status-chip" size="small">
          {{ item.status }}
        </v-chip>
      </template>

      <!-- Actions Column -->
      <template #item.actions="{ item }">
        <div class="action-buttons">
          <v-btn
            size="small"
            color="primary"
            variant="text"
            @click="$emit('view', item)"
            icon
            class="action-btn details-btn"
            title="عرض التفاصيل"
          >
            <v-icon size="16">mdi-eye</v-icon>
          </v-btn>
          <v-btn
            v-if="canEdit"
            size="small"
            color="success"
            variant="text"
            @click="$emit('edit', item)"
            icon
            class="action-btn"
            title="تعديل"
          >
            <v-icon size="16">mdi-pencil</v-icon>
          </v-btn>
          <v-btn
            v-if="canDelete"
            size="small"
            color="error"
            variant="text"
            @click="$emit('delete', item)"
            icon
            class="action-btn"
            title="حذف"
          >
            <v-icon size="16">mdi-delete</v-icon>
          </v-btn>
        </div>
      </template>
    </v-data-table>
  </v-card>
</template>

<script setup>
defineProps({
  employees: {
    type: Array,
    default: () => []
  },
  searchQuery: {
    type: String,
    default: ''
  },
  loading: {
    type: Boolean,
    default: false
  },
  canEdit: {
    type: Boolean,
    default: true
  },
  canDelete: {
    type: Boolean,
    default: true
  }
})

defineEmits(['view', 'edit', 'delete'])

const headers = [
  { title: '#', key: 'serial', align: 'center', sortable: false },
  { title: 'اسم الموظف', key: 'name', align: 'right' },
  { title: 'رقم الهوية', key: 'idNumber', align: 'center' },
  { title: 'القسم', key: 'department', align: 'center' },
  { title: 'المنصب', key: 'position', align: 'right' },
  { title: 'رقم الهاتف', key: 'phone', align: 'center' },
  { title: 'البريد الإلكتروني', key: 'email', align: 'right' },
  { title: 'الراتب', key: 'salary', align: 'center' },
  { title: 'تاريخ التوظيف', key: 'hireDate', align: 'center' },
  { title: 'الحالة', key: 'status', align: 'center' },
  { title: 'الإجراءات', key: 'actions', align: 'center', sortable: false }
]

const formatCurrency = (amount) => {
  if (!amount) return '0 د.ع'
  const formatted = new Intl.NumberFormat('ar-IQ', {
    style: 'decimal',
    minimumFractionDigits: 0,
    maximumFractionDigits: 0
  }).format(amount)
  return formatted + ' د.ع'
}
</script>

<style scoped>
@import './styles/human-resources.css';
</style>
