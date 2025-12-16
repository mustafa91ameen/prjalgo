<template>
  <v-card class="users-table" elevation="2">
    <v-card-title class="d-flex align-center justify-space-between">
      <span class="text-h5 font-weight-bold">قائمة المستخدمين</span>
      <div class="d-flex gap-2">
        <v-btn
          color="success"
          prepend-icon="mdi-download"
          variant="outlined"
          @click="$emit('export')"
        >
          تصدير البيانات
        </v-btn>
        <v-btn
          color="info"
          prepend-icon="mdi-account-plus"
          variant="outlined"
          @click="$emit('bulk-add')"
        >
          إضافة متعدد
        </v-btn>
      </div>
    </v-card-title>

    <v-data-table
      :headers="headers"
      :items="users"
      :search="searchQuery"
      :loading="loading"
      class="elevation-0"
      :items-per-page="10"
    >
      <template v-slot:item.user="{ item }">
        <div class="d-flex align-center">
          <v-avatar size="40" class="me-3">
            <v-img :src="item.avatar || defaultAvatar" />
          </v-avatar>
          <div>
            <div class="font-weight-medium">{{ item.name }}</div>
            <div class="text-caption text-grey">{{ item.email }}</div>
          </div>
        </div>
      </template>

      <template v-slot:item.role="{ item }">
        <v-chip
          :color="getRoleColor(item.role)"
          size="small"
          variant="flat"
        >
          {{ getRoleText(item.role) }}
        </v-chip>
      </template>

      <template v-slot:item.status="{ item }">
        <v-chip
          :color="getStatusColor(item.status)"
          size="small"
          variant="flat"
        >
          {{ getStatusText(item.status) }}
        </v-chip>
      </template>

      <template v-slot:item.lastLogin="{ item }">
        <span class="text-body-2">{{ formatDate(item.lastLogin) }}</span>
      </template>

      <template v-slot:item.actions="{ item }">
        <v-btn
          icon="mdi-eye"
          size="small"
          variant="elevated"
          class="view-btn"
          data-action="view"
          @click="$emit('view', item)"
        />
        <v-btn
          icon="mdi-pencil"
          size="small"
          variant="elevated"
          class="edit-btn"
          data-action="edit"
          @click="$emit('edit', item)"
        />
        <v-btn
          icon="mdi-key"
          size="small"
          variant="elevated"
          class="reset-btn"
          data-action="reset"
          @click="$emit('reset-password', item)"
        />
        <v-btn
          icon="mdi-delete"
          size="small"
          variant="elevated"
          class="delete-btn"
          data-action="delete"
          @click="$emit('delete', item)"
        />
      </template>
    </v-data-table>
  </v-card>
</template>

<script setup>
import { formatDateTime } from '@/utils/formatters'

defineProps({
  users: {
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
  }
})

defineEmits(['view', 'edit', 'reset-password', 'delete', 'export', 'bulk-add'])

const defaultAvatar = 'https://ui-avatars.com/api/?name=User&background=667eea&color=fff&size=128'

// Use centralized formatter with custom fallback
const formatDate = (date) => {
  if (!date) return 'لم يسجل دخول'
  return formatDateTime(date)
}

const headers = [
  { title: 'المستخدم', key: 'user', sortable: true },
  { title: 'الدور', key: 'role', sortable: true },
  { title: 'القسم', key: 'department', sortable: true },
  { title: 'الحالة', key: 'status', sortable: true },
  { title: 'آخر دخول', key: 'lastLogin', sortable: true },
  { title: 'الإجراءات', key: 'actions', sortable: false }
]

const getRoleColor = (role) => {
  const colors = {
    'admin': 'red',
    'project_manager': 'blue',
    'engineer': 'green',
    'accountant': 'purple',
    'employee': 'orange',
    'reviewer': 'teal',
    'user': 'grey'
  }
  return colors[role] || 'grey'
}

const getRoleText = (role) => {
  const texts = {
    'admin': 'مدير النظام',
    'project_manager': 'مدير المشاريع',
    'engineer': 'مهندس',
    'accountant': 'محاسب',
    'employee': 'موظف إداري',
    'reviewer': 'مراجع',
    'user': 'مستخدم عادي'
  }
  return texts[role] || 'غير محدد'
}

const getStatusColor = (status) => {
  const colors = {
    'active': 'success',
    'inactive': 'error',
    'pending': 'warning',
    'banned': 'grey'
  }
  return colors[status] || 'grey'
}

const getStatusText = (status) => {
  const texts = {
    'active': 'نشط',
    'inactive': 'غير نشط',
    'pending': 'معلق',
    'banned': 'محظور'
  }
  return texts[status] || 'غير محدد'
}
</script>

<style scoped>
@import './styles/users.css';
</style>
