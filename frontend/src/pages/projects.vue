<template>
  <v-container fluid class="py-6">
    <v-card class="mb-4" elevation="2">
      <v-card-title class="d-flex align-center">
        <v-icon class="me-2" color="primary">mdi-folder-multiple</v-icon>
        <span class="text-h5 font-weight-bold">المشاريع (من الخادم)</span>
        <v-spacer />
        <v-btn size="small" color="primary" variant="elevated" @click="loadData" :loading="loading">
          تحديث
        </v-btn>
      </v-card-title>
      <v-card-text>
        <div class="d-flex flex-wrap ga-4">
          <v-chip color="primary" variant="elevated">الإجمالي: {{ stats.total ?? '...' }}</v-chip>
          <v-chip color="success" variant="elevated">قيد التنفيذ: {{ stats.inProgress ?? '...' }}</v-chip>
          <v-chip color="warning" variant="elevated">معلّق: {{ stats.pending ?? '...' }}</v-chip>
          <v-chip color="info" variant="elevated">مكتمل: {{ stats.completed ?? '...' }}</v-chip>
          <v-chip color="secondary" variant="elevated">
            متوسط التقدم: {{ stats.averageProgress != null ? `${stats.averageProgress}%` : '...' }}
          </v-chip>
        </div>
      </v-card-text>
    </v-card>

    <v-card elevation="2">
      <v-card-title class="d-flex align-center justify-space-between">
        <div class="d-flex align-center">
          <v-icon class="me-2" color="primary">mdi-format-list-bulleted</v-icon>
          <span class="text-h6 font-weight-bold">قائمة المشاريع</span>
        </div>
        <div class="d-flex ga-2 align-center">
          <v-text-field
            v-model="search"
            density="compact"
            variant="outlined"
            label="بحث بالاسم"
            hide-details
            clearable
            style="max-width: 240px;"
          />
          <v-select
            v-model="limit"
            :items="[10, 20, 50]"
            density="compact"
            variant="outlined"
            label="عدد الصفوف"
            hide-details
            style="max-width: 140px;"
          />
        </div>
      </v-card-title>
      <v-data-table
        :headers="headers"
        :items="filteredProjects"
        :items-per-page="limit"
        :loading="loading"
        :page.sync="page"
        class="elevation-0"
      >
        <template #item.status="{ item }">
          <v-chip :color="statusColor(item.status)" size="small" variant="elevated">
            {{ statusLabel(item.status) }}
          </v-chip>
        </template>
        <template #item.progressPercentage="{ item }">
          <div class="d-flex align-center ga-2">
            <v-progress-linear :model-value="item.progressPercentage || 0" height="6" rounded />
            <span>{{ item.progressPercentage ?? 0 }}%</span>
          </div>
        </template>
        <template #item.startDate="{ item }">
          {{ formatDate(item.startDate) }}
        </template>
        <template #item.totalCost="{ item }">
          {{ formatCurrency(item.totalCost) }}
        </template>
        <template #no-data>
          <div class="pa-6 text-center text-medium-emphasis">
            لا توجد بيانات مشاريع حالياً.
          </div>
        </template>
      </v-data-table>
    </v-card>
  </v-container>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { listProjects, getProjectStats } from '@/api/projects'

const loading = ref(false)
const error = ref('')
const projects = ref([])
const stats = ref({})
const page = ref(1)
const limit = ref(10)
const search = ref('')

const headers = [
  { title: 'الاسم', key: 'name', align: 'start' },
  { title: 'النوع', key: 'type' },
  { title: 'الموقع', key: 'location' },
  { title: 'تاريخ البدء', key: 'startDate' },
  { title: 'الحالة', key: 'status' },
  { title: 'النسبة', key: 'progressPercentage' },
  { title: 'التكلفة الكلية', key: 'totalCost' },
]

const filteredProjects = computed(() => {
  const term = search.value.trim().toLowerCase()
  if (!term) return projects.value
  return projects.value.filter(p => (p.name || '').toLowerCase().includes(term))
})

const statusMap = {
  pending: 'قيد الانتظار',
  in_progress: 'قيد التنفيذ',
  completed: 'مكتمل',
}

function statusLabel(status) {
  return statusMap[status] || status || 'غير محدد'
}

function statusColor(status) {
  if (status === 'completed') return 'success'
  if (status === 'in_progress') return 'info'
  if (status === 'pending') return 'warning'
  return 'default'
}

function formatDate(value) {
  if (!value) return '-'
  const d = new Date(value)
  if (Number.isNaN(d.getTime())) return value
  return d.toLocaleDateString('en-US', {
    year: 'numeric',
    month: '2-digit',
    day: '2-digit'
  })
}

function formatCurrency(value) {
  if (value == null) return '-'
  return new Intl.NumberFormat('en-US', {
    minimumFractionDigits: 0,
    maximumFractionDigits: 0
  }).format(value) + ' IQD'
}

async function loadData() {
  loading.value = true
  error.value = ''
  try {
    const [list, statsData] = await Promise.all([
      listProjects({ page: page.value, limit: limit.value }),
      getProjectStats({ period: 'all' }),
    ])
    projects.value = Array.isArray(list) ? list : []
    stats.value = statsData || {}
  } catch (err) {
    console.error('Projects load error', err)
    error.value = err?.message || 'فشل تحميل البيانات'
  } finally {
    loading.value = false
  }
}

onMounted(loadData)
</script>

<style scoped>
.ga-4 {
  gap: 16px;
}
.ga-2 {
  gap: 8px;
}
</style>

