<template>
  <div class="pages-container">
    <!-- Page Header Component -->
    <PageHeader
      title="إدارة الصفحات"
      subtitle="إدارة وتنظيم جميع صفحات النظام"
      badge="الصفحات"
      badgeType="info"
      class="pages-header"
    >
      <template #actions>
        <button class="page-action-btn secondary">
          <i class="mdi mdi-export"></i>
          تصدير
        </button>
        <button class="page-icon-btn">
          <i class="mdi mdi-dots-vertical"></i>
        </button>
      </template>
    </PageHeader>

    <!-- Statistics Cards -->
    <div class="stats-grid">
      <!-- Total Pages Card -->
      <v-card class="stat-card" elevation="0">
        <div class="stat-card-content">
          <div class="stat-icon pages">
            <i class="mdi mdi-file-document-multiple"></i>
          </div>
          <div class="stat-info">
            <div class="stat-label">إجمالي الصفحات</div>
            <div class="stat-value">{{ totalPages }}</div>
          </div>
        </div>
      </v-card>

      <!-- Active Pages Card -->
      <v-card class="stat-card" elevation="0">
        <div class="stat-card-content">
          <div class="stat-icon success">
            <i class="mdi mdi-check-circle"></i>
          </div>
          <div class="stat-info">
            <div class="stat-label">صفحات نشطة</div>
            <div class="stat-value">{{ activePages }}</div>
          </div>
        </div>
      </v-card>

      <!-- Inactive Pages Card -->
      <v-card class="stat-card" elevation="0">
        <div class="stat-card-content">
          <div class="stat-icon warning">
            <i class="mdi mdi-cancel"></i>
          </div>
          <div class="stat-info">
            <div class="stat-label">صفحات غير نشطة</div>
            <div class="stat-value">{{ inactivePages }}</div>
          </div>
        </div>
      </v-card>
    </div>

    <!-- Pages List Header -->
    <div class="pages-list-header">
      <div class="list-header-content">
        <div class="list-header-info">
          <h2 class="list-header-title">
            <i class="mdi mdi-format-list-bulleted"></i>
            قائمة الصفحات
          </h2>
          <p class="list-header-subtitle">عرض جميع صفحات النظام</p>
        </div>
        <div class="list-header-actions">
          <button v-if="canCreate" class="list-action-btn primary" @click="openAddDialog">
            <i class="mdi mdi-plus"></i>
            إضافة صفحة جديدة
          </button>
        </div>
      </div>
    </div>

    <!-- Pages Table -->
    <v-card class="mb-6">
      <v-data-table
        :headers="headers"
        :items="pages"
        :loading="loading"
        class="elevation-1"
      >
        <template v-slot:item.icon="{ item }">
          <i :class="`mdi ${item.icon}`" style="font-size: 24px;"></i>
        </template>
        <template v-slot:item.status="{ item }">
          <v-chip
            :color="getStatusColor(item.status)"
            size="small"
          >
            {{ getStatusText(item.status) }}
          </v-chip>
        </template>
        <template v-slot:item.actions="{ item }">
          <v-btn
            v-if="canUpdate"
            size="small"
            color="primary"
            class="me-2"
            @click="editPageItem(item)"
          >
            <i class="mdi mdi-pencil"></i>
          </v-btn>
          <v-btn
            v-if="canDelete"
            size="small"
            color="error"
            @click="deletePageItem(item)"
          >
            <i class="mdi mdi-delete"></i>
          </v-btn>
        </template>
      </v-data-table>
    </v-card>

    <!-- Add Page Dialog -->
    <v-dialog v-model="showAddDialog" max-width="800" persistent>
      <v-card class="page-dialog">
        <v-card-title class="dialog-header">
          <i class="mdi mdi-file-document-plus"></i>
          إضافة صفحة جديدة
        </v-card-title>
        
        <v-card-text class="dialog-content">
          <v-form ref="pageForm" v-model="formValid">
            <v-row>
              <!-- اسم الصفحة -->
              <v-col cols="12" md="6">
                <v-text-field
                  v-model="newPage.name"
                  label="اسم الصفحة"
                  variant="outlined"
                  density="comfortable"
                  :rules="[v => !!v || 'اسم الصفحة مطلوب']"
                  required
                  autofocus
                ></v-text-field>
              </v-col>

              <!-- مسار الصفحة -->
              <v-col cols="12" md="6">
                <v-text-field
                  v-model="newPage.route"
                  label="مسار الصفحة"
                  variant="outlined"
                  density="comfortable"
                  placeholder="/example"
                  :rules="[
                    v => !!v || 'مسار الصفحة مطلوب',
                    v => v.startsWith('/') || 'المسار يجب أن يبدأ بـ /'
                  ]"
                  required
                ></v-text-field>
              </v-col>

              <!-- أيقونة الصفحة -->
              <v-col cols="12" md="6">
                <v-text-field
                  v-model="newPage.icon"
                  label="أيقونة الصفحة"
                  variant="outlined"
                  density="comfortable"
                  placeholder="mdi-home"
                  :rules="[v => !!v || 'أيقونة الصفحة مطلوبة']"
                  required
                >
                  <template #prepend-inner>
                    <i :class="`mdi ${newPage.icon}`" v-if="newPage.icon"></i>
                  </template>
                </v-text-field>
              </v-col>

              <!-- حالة الصفحة -->
              <v-col cols="12" md="6">
                <v-select
                  v-model="newPage.status"
                  :items="statusOptions"
                  label="حالة الصفحة"
                  variant="outlined"
                  density="comfortable"
                  :rules="[v => !!v || 'حالة الصفحة مطلوبة']"
                  required
                ></v-select>
              </v-col>
            </v-row>
          </v-form>
        </v-card-text>

        <v-card-actions class="dialog-actions">
          <v-spacer></v-spacer>
          <button class="dialog-btn cancel" @click="closeDialog">
            <i class="mdi mdi-close"></i>
            إلغاء
          </button>
          <button class="dialog-btn save" @click="savePage" :disabled="!formValid">
            <i class="mdi mdi-content-save"></i>
            حفظ
          </button>
        </v-card-actions>
      </v-card>
    </v-dialog>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import PageHeader from '../components/PageHeader.vue'
import { listPages, createPage, updatePage, deletePage as apiDeletePage } from '@/api/pages'
import { usePermissions } from '@/composables/usePermissions'
import { useToast } from '@/composables/useToast'
import { DEFAULT_PAGE, DEFAULT_LIMIT } from '@/constants/pagination'

const { success, error: showError } = useToast()
const { canCreate, canUpdate, canDelete } = usePermissions('/pages')

const loading = ref(false)
const showAddDialog = ref(false)
const formValid = ref(false)
const pageForm = ref(null)
const editingPage = ref(null)

// Pagination
const page = ref(DEFAULT_PAGE)
const limit = ref(DEFAULT_LIMIT)
const total = ref(0)

// New page data
const newPage = ref({
  name: '',
  route: '',
  icon: '',
  status: 'active'
})

const statusOptions = [
  { title: 'نشط', value: 'active' },
  { title: 'غير نشط', value: 'inactive' }
]

const pages = ref([])

// Fetch pages from API
const fetchPages = async () => {
  loading.value = true
  try {
    const response = await listPages({ page: page.value, limit: limit.value })
    if (response.success) {
      pages.value = (response.data.items || []).map(p => ({
        id: p.id,
        name: p.name,
        route: p.route,
        icon: p.icon || 'mdi-file-document',
        status: p.status || 'active',
        order: p.order || 0
      }))
      total.value = response.data.total || 0
    }
  } catch (error) {
    console.error('Error fetching pages:', error)
    showError('حدث خطأ في جلب الصفحات')
  } finally {
    loading.value = false
  }
}

const headers = [
  { title: 'الأيقونة', key: 'icon', align: 'center' },
  { title: 'اسم الصفحة', key: 'name', align: 'start' },
  { title: 'المسار', key: 'route', align: 'center' },
  { title: 'الترتيب', key: 'order', align: 'center' },
  { title: 'الحالة', key: 'status', align: 'center' },
  { title: 'الإجراءات', key: 'actions', align: 'center', sortable: false }
]

// Computed properties
const totalPages = computed(() => total.value || pages.value.length)

const activePages = computed(() => {
  return pages.value.filter(p => p.status === 'active').length
})

const inactivePages = computed(() => {
  return pages.value.filter(p => p.status === 'inactive').length
})

// Methods
const openAddDialog = () => {
  showAddDialog.value = true
}

const closeDialog = () => {
  showAddDialog.value = false
  editingPage.value = null
  resetForm()
}

const resetForm = () => {
  newPage.value = {
    name: '',
    route: '',
    icon: '',
    status: 'active'
  }
  if (pageForm.value) {
    pageForm.value.reset()
  }
}

const savePage = async () => {
  if (!formValid.value) return

  loading.value = true
  try {
    const pageData = {
      name: newPage.value.name,
      route: newPage.value.route,
      icon: newPage.value.icon,
      status: newPage.value.status
    }

    let response
    if (editingPage.value) {
      response = await updatePage(editingPage.value.id, pageData)
    } else {
      response = await createPage(pageData)
    }

    if (response.success) {
      success(editingPage.value ? 'تم تحديث الصفحة بنجاح' : 'تم إضافة الصفحة بنجاح')
      closeDialog()
      fetchPages()
    } else {
      showError(response.message || 'حدث خطأ')
    }
  } catch (error) {
    console.error('Error saving page:', error)
    showError('حدث خطأ في حفظ الصفحة')
  } finally {
    loading.value = false
  }
}

// Edit page
const editPageItem = (pageItem) => {
  editingPage.value = pageItem
  newPage.value = {
    name: pageItem.name,
    route: pageItem.route,
    icon: pageItem.icon,
    status: pageItem.status
  }
  showAddDialog.value = true
}

// Delete page
const deletePageItem = async (pageItem) => {
  if (!confirm('هل أنت متأكد من حذف هذه الصفحة؟')) return

  loading.value = true
  try {
    const response = await apiDeletePage(pageItem.id)
    if (response.success) {
      success('تم حذف الصفحة بنجاح')
      fetchPages()
    } else {
      showError(response.message || 'حدث خطأ في الحذف')
    }
  } catch (error) {
    console.error('Error deleting page:', error)
    showError('حدث خطأ في حذف الصفحة')
  } finally {
    loading.value = false
  }
}

// Initialize
onMounted(() => {
  fetchPages()
})

const getStatusColor = (status) => {
  const colors = {
    active: 'success',
    inactive: 'warning'
  }
  return colors[status] || 'grey'
}

const getStatusText = (status) => {
  const texts = {
    active: 'نشط',
    inactive: 'غير نشط'
  }
  return texts[status] || status
}
</script>

<style scoped>
.pages-container {
  padding: 32px;
  max-width: 1600px;
  margin: 0 auto;
}

/* Pages Header Custom Color */
.pages-header {
  background: linear-gradient(135deg, #018790 0%, #005461 100%) !important;
}

.pages-header::before {
  background: linear-gradient(135deg, #06b6d4 0%, #10b981 50%, #06b6d4 100%) !important;
}

/* Statistics Grid - 3 cards */
.stats-grid {
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  gap: 16px;
  margin-bottom: 24px;
}

.stat-card {
  border-radius: 16px !important;
  overflow: hidden;
  transition: all 0.4s cubic-bezier(0.4, 0, 0.2, 1);
  background: linear-gradient(135deg, #0a3d42 0%, #052428 100%) !important;
  border: 2px solid transparent !important;
  position: relative;
}

.stat-card::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  border-radius: 16px;
  padding: 2px;
  background: linear-gradient(135deg, #06b6d4 0%, #10b981 50%, #14b8a6 100%);
  -webkit-mask: linear-gradient(#fff 0 0) content-box, linear-gradient(#fff 0 0);
  mask: linear-gradient(#fff 0 0) content-box, linear-gradient(#fff 0 0);
  -webkit-mask-composite: xor;
  mask-composite: exclude;
  pointer-events: none;
}

.stat-card:hover {
  transform: translateY(-4px) scale(1.01);
  box-shadow: 0 12px 24px rgba(6, 182, 212, 0.3),
              0 0 40px rgba(16, 185, 129, 0.2) !important;
}

.stat-card-content {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  padding: 20px 16px;
  text-align: center;
  position: relative;
  z-index: 1;
}

.stat-info {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 4px;
}

.stat-label {
  font-size: 13px;
  color: rgba(255, 255, 255, 0.7);
  margin-bottom: 2px;
  font-weight: 500;
  letter-spacing: 0.3px;
}

.stat-value {
  font-size: 28px;
  font-weight: 800;
  background: linear-gradient(135deg, #ffffff 0%, #e0e7ff 100%);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background-clip: text;
  margin-bottom: 4px;
}

.stat-change {
  display: flex;
  align-items: center;
  gap: 4px;
  font-size: 11px;
  font-weight: 600;
  padding: 4px 10px;
  border-radius: 12px;
  background: rgba(255, 255, 255, 0.1);
  backdrop-filter: blur(10px);
}

.stat-change.positive {
  color: #34d399;
}

.stat-change.negative {
  color: #f87171;
}

.stat-icon {
  width: 48px;
  height: 48px;
  border-radius: 12px;
  display: flex;
  align-items: center;
  justify-content: center;
  color: white;
  font-size: 24px;
  margin-bottom: 12px;
  background: linear-gradient(135deg, #06b6d4 0%, #0891b2 100%) !important;
  box-shadow: 0 6px 16px rgba(6, 182, 212, 0.4);
}

.stat-icon.pages {
  background: linear-gradient(135deg, #06b6d4 0%, #0891b2 100%) !important;
  box-shadow: 0 6px 16px rgba(6, 182, 212, 0.4);
}

.stat-icon.success {
  background: linear-gradient(135deg, #10b981 0%, #059669 100%) !important;
  box-shadow: 0 6px 16px rgba(16, 185, 129, 0.4);
}

.stat-icon.warning {
  background: linear-gradient(135deg, #f59e0b 0%, #d97706 100%) !important;
  box-shadow: 0 6px 16px rgba(245, 158, 11, 0.4);
}

/* Pages List Header */
.pages-list-header {
  background: linear-gradient(135deg, #0a3d42 0%, #052428 100%);
  border-radius: 16px;
  margin-bottom: 24px;
  padding: 16px 24px;
  border: 2px solid transparent;
  position: relative;
  overflow: hidden;
}

.pages-list-header::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  border-radius: 16px;
  padding: 2px;
  background: linear-gradient(135deg, #06b6d4 0%, #0891b2 50%, #06b6d4 100%);
  -webkit-mask: linear-gradient(#fff 0 0) content-box, linear-gradient(#fff 0 0);
  mask: linear-gradient(#fff 0 0) content-box, linear-gradient(#fff 0 0);
  -webkit-mask-composite: xor;
  mask-composite: exclude;
  pointer-events: none;
}

.list-header-content {
  display: flex;
  justify-content: space-between;
  align-items: center;
  position: relative;
  z-index: 1;
}

.list-header-info {
  flex: 1;
}

.list-header-title {
  font-size: 18px;
  font-weight: 700;
  color: rgba(255, 255, 255, 0.95);
  margin: 0 0 2px 0;
  display: flex;
  align-items: center;
  gap: 8px;
}

.list-header-title i {
  color: #06b6d4;
  font-size: 20px;
}

.list-header-subtitle {
  font-size: 12px;
  color: rgba(255, 255, 255, 0.6);
  margin: 0;
}

.list-header-actions {
  display: flex;
  gap: 12px;
}

.list-action-btn {
  padding: 8px 18px;
  border-radius: 12px;
  border: none;
  font-size: 13px;
  font-weight: 700;
  cursor: pointer;
  display: inline-flex;
  align-items: center;
  gap: 8px;
  transition: all 0.3s ease;
}

.list-action-btn.primary {
  background: linear-gradient(135deg, #06b6d4 0%, #0891b2 100%);
  color: white;
  box-shadow: 0 4px 12px rgba(6, 182, 212, 0.3);
}

.list-action-btn.primary:hover {
  transform: translateY(-2px);
  box-shadow: 0 6px 16px rgba(6, 182, 212, 0.4);
}

.list-action-btn i {
  font-size: 16px;
}

/* Page Dialog */
.page-dialog {
  border-radius: 16px !important;
  background: linear-gradient(135deg, #0a3d42 0%, #052428 100%) !important;
  border: 2px solid transparent !important;
  position: relative;
  overflow: hidden;
  direction: rtl;
  text-align: right;
}

.page-dialog::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  border-radius: 16px;
  padding: 2px;
  background: linear-gradient(135deg, #06b6d4 0%, #0891b2 50%, #06b6d4 100%);
  -webkit-mask: linear-gradient(#fff 0 0) content-box, linear-gradient(#fff 0 0);
  mask: linear-gradient(#fff 0 0) content-box, linear-gradient(#fff 0 0);
  -webkit-mask-composite: xor;
  mask-composite: exclude;
  pointer-events: none;
}

.dialog-header {
  background: rgba(6, 182, 212, 0.1) !important;
  color: rgba(255, 255, 255, 0.95) !important;
  font-size: 20px !important;
  font-weight: 700 !important;
  padding: 20px 24px !important;
  display: flex;
  align-items: center;
  gap: 10px;
  position: relative;
  z-index: 1;
  direction: rtl;
  text-align: right;
}

.dialog-header i {
  color: #06b6d4;
  font-size: 24px;
}

.dialog-content {
  padding: 24px !important;
  position: relative;
  z-index: 1;
  direction: rtl;
  text-align: right;
}

.dialog-content :deep(.v-field) {
  background: rgba(255, 255, 255, 0.08) !important;
  border-radius: 12px;
  direction: rtl;
  text-align: right;
}

.dialog-content :deep(.v-field__outline) {
  color: rgba(255, 255, 255, 0.2) !important;
}

.dialog-content :deep(.v-field--focused .v-field__outline) {
  color: #06b6d4 !important;
}

.dialog-content :deep(.v-label) {
  color: rgba(255, 255, 255, 0.7) !important;
  right: 12px !important;
  left: auto !important;
}

.dialog-content :deep(.v-field__input) {
  color: rgba(255, 255, 255, 0.95) !important;
  text-align: right;
  direction: rtl;
}

.dialog-content :deep(.v-input__details) {
  direction: rtl;
  text-align: right;
}

.dialog-actions {
  padding: 16px 24px !important;
  border-top: 1px solid rgba(255, 255, 255, 0.1);
  position: relative;
  z-index: 1;
  direction: rtl;
  text-align: right;
}

.dialog-btn {
  padding: 10px 24px;
  border-radius: 12px;
  border: none;
  font-size: 14px;
  font-weight: 700;
  cursor: pointer;
  display: inline-flex;
  align-items: center;
  gap: 8px;
  transition: all 0.3s ease;
}

.dialog-btn.cancel {
  background: rgba(255, 255, 255, 0.1);
  color: rgba(255, 255, 255, 0.8);
  border: 1px solid rgba(255, 255, 255, 0.2);
}

.dialog-btn.cancel:hover {
  background: rgba(255, 255, 255, 0.15);
  transform: translateY(-2px);
}

.dialog-btn.save {
  background: linear-gradient(135deg, #06b6d4 0%, #0891b2 100%);
  color: white;
  box-shadow: 0 4px 12px rgba(6, 182, 212, 0.3);
}

.dialog-btn.save:hover:not(:disabled) {
  transform: translateY(-2px);
  box-shadow: 0 6px 16px rgba(6, 182, 212, 0.4);
}

.dialog-btn.save:disabled {
  opacity: 0.5;
  cursor: not-allowed;
}

.dialog-btn i {
  font-size: 18px;
}

.v-card {
  border-radius: 12px;
}

/* Responsive Styles */
@media (max-width: 1024px) {
  .pages-container {
    padding: 24px;
  }

  .stats-grid {
    grid-template-columns: repeat(2, 1fr);
  }

  .list-header-content {
    flex-direction: column;
    align-items: flex-start;
    gap: 16px;
  }

  .list-header-actions {
    width: 100%;
  }

  .list-action-btn.primary {
    width: 100%;
    justify-content: center;
  }
}

@media (max-width: 768px) {
  .pages-container {
    padding: 16px;
  }

  .stats-grid {
    grid-template-columns: 1fr;
    gap: 12px;
  }

  .stat-card-content {
    flex-direction: row;
    justify-content: flex-start;
    gap: 16px;
    padding: 16px;
    text-align: right;
  }

  .stat-icon {
    margin-bottom: 0;
  }

  .stat-info {
    align-items: flex-start;
  }

  .pages-list-header {
    padding: 12px 16px;
  }

  .list-header-title {
    font-size: 16px;
  }

  .list-action-btn {
    padding: 10px 16px;
    font-size: 12px;
  }

  /* Dialog responsive */
  .page-dialog .dialog-header {
    padding: 16px !important;
    font-size: 18px !important;
  }

  .page-dialog .dialog-content {
    padding: 16px !important;
  }

  .page-dialog .dialog-actions {
    padding: 12px 16px !important;
    flex-wrap: wrap;
    gap: 8px;
  }

  .dialog-btn {
    padding: 8px 16px;
    font-size: 13px;
  }

  /* Table responsive */
  .v-data-table {
    font-size: 13px;
  }

  .v-data-table :deep(th),
  .v-data-table :deep(td) {
    padding: 8px 12px !important;
  }
}

@media (max-width: 480px) {
  .pages-container {
    padding: 12px;
  }

  .stat-value {
    font-size: 24px;
  }

  .stat-label {
    font-size: 12px;
  }

  .stat-icon {
    width: 40px;
    height: 40px;
    font-size: 20px;
  }

  .pages-list-header {
    padding: 12px;
    border-radius: 12px;
  }

  .list-header-title {
    font-size: 14px;
  }

  .list-header-subtitle {
    font-size: 11px;
  }

  .list-action-btn {
    padding: 8px 12px;
    font-size: 11px;
    border-radius: 10px;
  }

  /* Dialog responsive - full width on mobile */
  :deep(.v-dialog) {
    margin: 8px !important;
  }

  .page-dialog {
    border-radius: 12px !important;
  }

  .page-dialog .dialog-header {
    padding: 12px 16px !important;
    font-size: 16px !important;
    gap: 8px;
  }

  .page-dialog .dialog-header i {
    font-size: 20px;
  }

  .page-dialog .dialog-content {
    padding: 12px !important;
    max-height: 50vh;
  }

  .dialog-btn {
    padding: 8px 12px;
    font-size: 12px;
    flex: 1;
  }

  /* Table overflow handling */
  .v-card.mb-6 {
    overflow-x: auto;
  }

  .v-data-table :deep(.v-table__wrapper) {
    overflow-x: auto;
  }

  .v-data-table :deep(th),
  .v-data-table :deep(td) {
    padding: 6px 8px !important;
    font-size: 12px;
    white-space: nowrap;
  }

  .v-btn.me-2 {
    margin-inline-end: 4px !important;
  }
}
</style>
