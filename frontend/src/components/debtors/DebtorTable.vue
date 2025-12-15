<template>
  <v-card class="data-table-card" elevation="2">
    <v-card-title class="d-flex align-center justify-space-between pa-6">
      <div class="d-flex align-center">
        <v-icon class="me-2" color="primary">mdi-table</v-icon>
        <span class="text-h6 font-weight-bold">قائمة المديونون</span>
        <v-chip class="ms-3" color="primary" variant="tonal">
          {{ debtors.length }} عنصر
        </v-chip>
      </div>
      <div class="d-flex align-center gap-2">
        <v-btn
          v-if="canAdd"
          color="primary"
          size="default"
          prepend-icon="mdi-plus"
          class="add-button"
          @click="$emit('add')"
        >
          إضافة مديون جديد
        </v-btn>
        <v-btn
          icon="mdi-refresh"
          variant="text"
          color="primary"
          @click="$emit('refresh')"
          class="action-button"
        />
        <v-btn
          icon="mdi-download"
          variant="text"
          color="success"
          @click="$emit('export')"
          class="action-button"
        />
      </div>
    </v-card-title>
    <v-data-table
      :headers="tableHeaders"
      :items="debtors"
      :loading="loading"
      class="elevation-0"
      no-data-text="لا توجد بيانات"
      loading-text="جاري التحميل..."
    >
      <!-- Name Column -->
      <template v-slot:item.name="{ item }">
        <div class="d-flex align-center" style="cursor: pointer;" @click="$emit('view-debts', item)">
          <v-avatar
            :color="getStatusColor(item.status)"
            size="32"
            class="me-3"
          >
            <span class="text-white font-weight-bold">
              {{ item.name?.charAt(0) || '?' }}
            </span>
          </v-avatar>
          <div>
            <div class="font-weight-medium text-primary">{{ item.name }}</div>
            <div class="text-caption text-grey-darken-1">{{ item.email }}</div>
            <div class="text-caption text-primary">انقر لعرض الديون والتسديدات</div>
          </div>
        </div>
      </template>

      <!-- Amount Column -->
      <template v-slot:item.amount="{ item }">
        <div class="text-right">
          <div class="font-weight-bold text-h6">{{ formatCurrency(item.amount) }}</div>
          <div class="text-caption text-grey-darken-1">
            {{ item.currency || 'د.ع' }}
          </div>
        </div>
      </template>

      <!-- Due Date Column -->
      <template v-slot:item.dueDate="{ item }">
        <div class="text-center">
          <div class="font-weight-medium">{{ formatDate(item.dueDate) }}</div>
          <v-chip
            :color="getDueDateColor(item.dueDate)"
            size="small"
            variant="tonal"
          >
            {{ getDueDateStatus(item.dueDate) }}
          </v-chip>
        </div>
      </template>

      <!-- Status Column -->
      <template v-slot:item.status="{ item }">
        <v-chip
          :color="getStatusColor(item.status)"
          :variant="item.status === 'paid' ? 'flat' : 'tonal'"
          size="small"
        >
          {{ getStatusText(item.status) }}
        </v-chip>
      </template>

      <!-- Actions Column -->
      <template v-slot:item.actions="{ item }">
        <div class="d-flex align-center gap-1">
          <v-btn
            icon="mdi-eye"
            size="small"
            variant="text"
            @click="$emit('view', item)"
          />
          <v-btn
            v-if="canEdit"
            icon="mdi-pencil"
            size="small"
            variant="text"
            @click="$emit('edit', item)"
          />
          <v-btn
            v-if="canEdit && item.status !== 'paid'"
            icon="mdi-credit-card"
            size="small"
            variant="text"
            color="success"
            @click="$emit('mark-paid', item)"
          />
          <v-btn
            v-if="canDelete"
            icon="mdi-delete"
            size="small"
            variant="text"
            color="error"
            @click="$emit('delete', item)"
          />
        </div>
      </template>
    </v-data-table>
  </v-card>
</template>

<script setup>
defineProps({
  debtors: {
    type: Array,
    default: () => []
  },
  loading: {
    type: Boolean,
    default: false
  },
  canAdd: {
    type: Boolean,
    default: true
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

defineEmits(['add', 'view', 'edit', 'delete', 'mark-paid', 'view-debts', 'refresh', 'export'])

const tableHeaders = [
  { title: 'المديون', key: 'name', sortable: true },
  { title: 'المبلغ', key: 'amount', sortable: true, align: 'center' },
  { title: 'تاريخ الاستحقاق', key: 'dueDate', sortable: true, align: 'center' },
  { title: 'الحالة', key: 'status', sortable: true, align: 'center' },
  { title: 'الإجراءات', key: 'actions', sortable: false, align: 'center' }
]

const formatCurrency = (amount) => {
  return new Intl.NumberFormat('ar-SA', {
    style: 'currency',
    currency: 'IQD',
    minimumFractionDigits: 0
  }).format(amount || 0).replace('IQD', 'د.ع')
}

const formatDate = (dateString) => {
  if (!dateString) return 'غير محدد'
  const date = new Date(dateString)
  return date.toLocaleDateString('ar-SA')
}

const getStatusColor = (status) => {
  const colors = {
    paid: 'success',
    pending: 'warning',
    overdue: 'error'
  }
  return colors[status] || 'grey'
}

const getStatusText = (status) => {
  const texts = {
    paid: 'مدفوع',
    pending: 'قيد الانتظار',
    overdue: 'متأخر'
  }
  return texts[status] || 'غير محدد'
}

const getDueDateColor = (dueDate) => {
  if (!dueDate) return 'grey'
  const today = new Date()
  const due = new Date(dueDate)
  const diffDays = Math.ceil((due - today) / (1000 * 60 * 60 * 24))

  if (diffDays < 0) return 'error'
  if (diffDays <= 7) return 'warning'
  return 'success'
}

const getDueDateStatus = (dueDate) => {
  if (!dueDate) return 'غير محدد'
  const today = new Date()
  const due = new Date(dueDate)
  const diffDays = Math.ceil((due - today) / (1000 * 60 * 60 * 24))

  if (diffDays < 0) return `متأخر ${Math.abs(diffDays)} يوم`
  if (diffDays === 0) return 'اليوم'
  if (diffDays === 1) return 'غداً'
  return `بعد ${diffDays} يوم`
}
</script>

<style scoped>
.data-table-card {
  border-radius: 16px !important;
  overflow: hidden;
}

.add-button {
  border-radius: 12px !important;
}

.action-button {
  border-radius: 8px !important;
}
</style>
