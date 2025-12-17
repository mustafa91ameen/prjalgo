<template>
  <v-card class="data-table-card" elevation="2">
    <v-card-title class="table-title">
      <span class="title-text">المشاريع</span>
      <v-chip class="ms-2" size="small" color="primary" variant="flat">
        {{ items.length }}
      </v-chip>
    </v-card-title>

    <v-data-table
      :headers="tableHeaders"
      :items="items"
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

      <!-- Project Name Column -->
      <template #item.projectName="{ item }">
        <span class="project-name">{{ item.projectName }}</span>
      </template>

      <!-- Start Date Column -->
      <template #item.startDate="{ item }">
        <span class="date-text">{{ item.startDate }}</span>
      </template>

      <!-- End Date Column -->
      <template #item.endDate="{ item }">
        <span class="date-text">{{ item.endDate }}</span>
      </template>

      <!-- Cost Column -->
      <template #item.cost="{ item }">
        <span class="cost-text">{{ item.cost }}</span>
      </template>

      <!-- Work Location Column -->
      <template #item.workLocation="{ item }">
        <span class="location-text">{{ item.workLocation }}</span>
      </template>

      <!-- Notes Column -->
      <template #item.notes="{ item }">
        <span class="notes-text">{{ item.notes || 'لايوجد' }}</span>
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
            color="warning"
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
  items: {
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

const tableHeaders = [
  { title: 'التسلسل', key: 'serial', sortable: false, align: 'center' },
  { title: 'اسم المشروع', key: 'projectName', sortable: true, align: 'right' },
  { title: 'تاريخ بدء المشروع', key: 'startDate', sortable: true, align: 'center' },
  { title: 'تاريخ انتهاء المشروع', key: 'endDate', sortable: true, align: 'center' },
  { title: 'التكلفة', key: 'cost', sortable: true, align: 'center' },
  { title: 'مكان العمل', key: 'workLocation', sortable: true, align: 'center' },
  { title: 'الملاحظات', key: 'notes', sortable: false, align: 'right' },
  { title: 'الاجراءات', key: 'actions', sortable: false, align: 'center' }
]
</script>

<style scoped>
.data-table-card {
  border-radius: var(--radius-2xl) !important;
  overflow: hidden;
}

.table-title {
  background: var(--background-surface) !important;
  color: var(--text-primary) !important;
  padding: 1.25rem 1.5rem !important;
  display: flex;
  align-items: center;
  border-bottom: 1px solid var(--border-light);
}

.title-text {
  font-weight: 700;
  font-size: var(--font-size-base-plus);
}

.serial-number {
  font-weight: 600;
  color: var(--color-slate-700);
}

.project-name {
  font-weight: 600;
  color: var(--color-blue-800);
}

.date-text {
  color: var(--color-slate-600);
}

.cost-text {
  font-weight: 600;
  color: var(--color-emerald-600);
}

.location-text {
  color: var(--color-slate-600);
}

.notes-text {
  color: var(--color-slate-500);
  font-size: var(--font-size-sm);
}

.action-buttons {
  display: flex;
  gap: var(--space-1);
  justify-content: center;
}

.action-btn {
  min-width: var(--space-8) !important;
  height: var(--space-8) !important;
}
</style>
