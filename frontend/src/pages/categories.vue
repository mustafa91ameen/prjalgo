<template>
  <div class="categories-container">
    <!-- Page Header Component -->
    <PageHeader
      title="إدارة التصنيفات"
      subtitle="إدارة وتنظيم جميع تصنيفات المشاريع"
      badge="التصنيفات"
      badgeType="secondary"
      class="categories-header"
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
      <!-- Total Categories Card -->
      <v-card class="stat-card" elevation="0">
        <div class="stat-card-content">
          <div class="stat-icon category">
            <i class="mdi mdi-shape"></i>
          </div>
          <div class="stat-info">
            <div class="stat-label">إجمالي التصنيفات</div>
            <div class="stat-value">{{ totalCategories }}</div>
          </div>
        </div>
      </v-card>

      <!-- Active Categories Card -->
      <v-card class="stat-card" elevation="0">
        <div class="stat-card-content">
          <div class="stat-icon success">
            <i class="mdi mdi-check-circle"></i>
          </div>
          <div class="stat-info">
            <div class="stat-label">تصنيفات نشطة</div>
            <div class="stat-value">{{ activeCategories }}</div>
          </div>
        </div>
      </v-card>

      <!-- Disabled Categories Card -->
      <v-card class="stat-card" elevation="0">
        <div class="stat-card-content">
          <div class="stat-icon warning">
            <i class="mdi mdi-cancel"></i>
          </div>
          <div class="stat-info">
            <div class="stat-label">تصنيفات معطلة</div>
            <div class="stat-value">{{ disabledCategories }}</div>
          </div>
        </div>
      </v-card>

      <!-- Total SubCategories Card -->
      <v-card class="stat-card" elevation="0">
        <div class="stat-card-content">
          <div class="stat-icon primary">
            <i class="mdi mdi-tag-multiple"></i>
          </div>
          <div class="stat-info">
            <div class="stat-label">إجمالي الأصناف الثانوية</div>
            <div class="stat-value">{{ totalSubCategoriesCount }}</div>
          </div>
        </div>
      </v-card>
    </div>

    <!-- Categories List Header -->
    <div class="categories-list-header">
      <div class="list-header-content">
        <div class="list-header-info">
          <h2 class="list-header-title">
            <i class="mdi mdi-format-list-bulleted"></i>
            قائمة التصنيفات
          </h2>
          <p class="list-header-subtitle">عرض جميع التصنيفات المتاحة</p>
        </div>
        <div class="list-header-actions">
          <button v-if="canCreate" class="list-action-btn primary" @click="openAddDialog">
            <i class="mdi mdi-plus"></i>
            إضافة تصنيف جديد
          </button>
        </div>
      </div>
    </div>

    <!-- Categories Table -->
    <v-card class="mb-6">
      <v-data-table
        :headers="headers"
        :items="categories"
        :loading="loading"
        class="elevation-1"
      >
        <template v-slot:item.subCategoriesCount="{ item }">
          <v-chip color="info" size="small" variant="flat">
            {{ getSubCategoriesCount(item.id) }} صنف
          </v-chip>
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
            size="small"
            color="info"
            variant="tonal"
            icon
            class="me-1"
            @click="openSubCategoryDialog(item)"
            title="الأصناف الثانوية"
          >
            <i class="mdi mdi-format-list-group"></i>
          </v-btn>
          <v-btn
            v-if="canUpdate"
            size="small"
            color="primary"
            variant="tonal"
            icon
            class="me-1"
            @click="editCategory(item)"
            title="تعديل"
          >
            <i class="mdi mdi-pencil"></i>
          </v-btn>
          <v-btn
            v-if="canDelete"
            size="small"
            color="error"
            variant="tonal"
            icon
            @click="confirmDeleteCategory(item)"
            title="حذف"
          >
            <i class="mdi mdi-delete"></i>
          </v-btn>
        </template>
      </v-data-table>
    </v-card>

    <!-- Add Category Dialog -->
    <v-dialog v-model="showAddDialog" max-width="600" persistent>
      <v-card class="category-dialog">
        <v-card-title class="dialog-header">
          <i class="mdi mdi-shape-plus"></i>
          إضافة تصنيف جديد
        </v-card-title>
        
        <v-card-text class="dialog-content">
          <v-form ref="categoryForm" v-model="formValid">
            <v-row>
              <!-- اسم القسم -->
              <v-col cols="12">
                <v-text-field
                  v-model="newCategory.name"
                  label="اسم القسم"
                  variant="outlined"
                  density="comfortable"
                  :rules="[v => !!v || 'اسم القسم مطلوب']"
                  required
                  autofocus
                ></v-text-field>
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
          <button class="dialog-btn save" @click="saveCategory" :disabled="!formValid">
            <i class="mdi mdi-content-save"></i>
            حفظ
          </button>
        </v-card-actions>
      </v-card>
    </v-dialog>

    <!-- Sub-Categories Dialog -->
    <v-dialog v-model="showSubCategoryDialog" max-width="800" persistent>
      <v-card class="category-dialog">
        <v-card-title class="dialog-header">
          <i class="mdi mdi-format-list-group"></i>
          الأصناف الثانوية - {{ selectedCategory?.name }}
        </v-card-title>

        <v-card-text class="dialog-content">
          <!-- Add New Sub-Category Form -->
          <div class="sub-category-form">
            <v-row align="center">
              <v-col cols="12" md="6">
                <v-text-field
                  v-model="newSubCategoryName"
                  label="اسم الصنف الثانوي"
                  variant="outlined"
                  density="comfortable"
                  placeholder="أدخل اسم الصنف الثانوي"
                  hide-details
                ></v-text-field>
              </v-col>
              <v-col cols="12" md="4">
                <v-text-field
                  v-model="newSubCategoryPercentage"
                  label="النسبة المئوية %"
                  variant="outlined"
                  density="comfortable"
                  type="number"
                  min="0"
                  max="100"
                  step="0.01"
                  placeholder="0"
                  hide-details
                ></v-text-field>
              </v-col>
              <v-col cols="12" md="2">
                <v-btn
                  color="teal"
                  variant="flat"
                  class="add-sub-btn"
                  @click="addSubCategory"
                  :disabled="!canAddSubCategory"
                  block
                >
                  <v-icon start>mdi-plus</v-icon>
                  إضافة
                </v-btn>
              </v-col>
            </v-row>
          </div>

          <!-- Sub-Categories List -->
          <div class="sub-categories-list">
            <div class="sub-list-header">
              <i class="mdi mdi-format-list-bulleted"></i>
              قائمة الأصناف الثانوية
            </div>

            <div v-if="currentSubCategories.length === 0" class="empty-sub-list">
              <i class="mdi mdi-folder-open-outline"></i>
              <p>لا توجد أصناف ثانوية</p>
            </div>

            <div v-else class="sub-items-container">
              <div
                v-for="(sub, index) in currentSubCategories"
                :key="sub.id"
                class="sub-category-item"
              >
                <div class="sub-item-info">
                  <i class="mdi mdi-tag-outline"></i>
                  <span class="sub-item-name">{{ sub.name }}</span>
                  <v-chip
                    v-if="sub.percentage !== null && sub.percentage !== undefined"
                    color="primary"
                    size="small"
                    variant="tonal"
                    class="ms-2"
                  >
                    {{ sub.percentage }}%
                  </v-chip>
                </div>
                <div class="sub-item-actions">
                  <button class="sub-action-btn edit" @click="openEditSubCategoryDialog(sub, index)" title="تعديل">
                    <i class="mdi mdi-pencil"></i>
                  </button>
                  <button class="sub-action-btn delete" @click="deleteSubCategory(index)" title="حذف">
                    <i class="mdi mdi-delete"></i>
                  </button>
                </div>
              </div>
            </div>
          </div>
        </v-card-text>

        <v-card-actions class="dialog-actions">
          <v-spacer></v-spacer>
          <button class="dialog-btn cancel" @click="closeSubCategoryDialog">
            <i class="mdi mdi-close"></i>
            إغلاق
          </button>
        </v-card-actions>
      </v-card>
    </v-dialog>

    <!-- Edit Sub-Category Dialog -->
    <v-dialog v-model="showEditSubCategoryDialog" max-width="500" persistent>
      <v-card class="category-dialog">
        <v-card-title class="dialog-header">
          <i class="mdi mdi-pencil"></i>
          تعديل الصنف الثانوي
        </v-card-title>

        <v-card-text class="dialog-content">
          <v-row>
            <v-col cols="12">
              <v-text-field
                v-model="editSubCategoryData.name"
                label="اسم الصنف الثانوي"
                variant="outlined"
                density="comfortable"
                :rules="[v => !!v || 'الاسم مطلوب']"
              ></v-text-field>
            </v-col>
            <v-col cols="12">
              <v-text-field
                v-model.number="editSubCategoryData.percentage"
                label="النسبة المئوية %"
                variant="outlined"
                density="comfortable"
                type="number"
                min="0"
                max="100"
                step="0.01"
                placeholder="0"
              ></v-text-field>
            </v-col>
          </v-row>
        </v-card-text>

        <v-card-actions class="dialog-actions">
          <v-spacer></v-spacer>
          <button class="dialog-btn cancel" @click="closeEditSubCategoryDialog">
            <i class="mdi mdi-close"></i>
            إلغاء
          </button>
          <button class="dialog-btn save" @click="saveEditSubCategory" :disabled="!editSubCategoryData.name || savingSubCategory">
            <i class="mdi mdi-content-save"></i>
            {{ savingSubCategory ? 'جاري الحفظ...' : 'حفظ' }}
          </button>
        </v-card-actions>
      </v-card>
    </v-dialog>

    <!-- Delete Category Confirmation Dialog -->
    <v-dialog v-model="showDeleteDialog" max-width="500">
      <v-card class="category-dialog">
        <v-card-title class="dialog-header" style="background: linear-gradient(135deg, #dc2626 0%, #b91c1c 100%) !important;">
          <i class="mdi mdi-delete-alert"></i>
          تأكيد حذف التصنيف
        </v-card-title>

        <v-card-text v-if="editingCategory" class="dialog-content text-center">
          <v-avatar size="60" color="error" class="mb-3">
            <i class="mdi mdi-shape" style="color: white; font-size: 30px;"></i>
          </v-avatar>
          <h4 style="color: rgba(255,255,255,0.95);">{{ editingCategory.name }}</h4>

          <v-alert v-if="getSubCategoriesCount(editingCategory.id) > 0" type="warning" variant="tonal" class="my-4" style="text-align: right; direction: rtl;">
            <strong>تنبيه:</strong> هذا التصنيف يحتوي على {{ getSubCategoriesCount(editingCategory.id) }} أصناف ثانوية.
            <br>يجب حذف جميع الأصناف الثانوية أولاً قبل حذف التصنيف.
          </v-alert>

          <v-alert type="error" variant="tonal" class="my-4" style="text-align: right; direction: rtl;">
            تحذير: هذا الإجراء لا يمكن التراجع عنه!
          </v-alert>

          <p style="color: rgba(255,255,255,0.7);">
            هل أنت متأكد من حذف هذا التصنيف نهائياً؟
          </p>
        </v-card-text>

        <v-card-actions class="dialog-actions">
          <v-spacer></v-spacer>
          <button class="dialog-btn cancel" @click="showDeleteDialog = false">
            إلغاء
          </button>
          <button class="dialog-btn delete" @click="deleteCategoryItem" :disabled="deleting || getSubCategoriesCount(editingCategory?.id) > 0">
            <i class="mdi mdi-delete"></i>
            {{ deleting ? 'جاري الحذف...' : 'حذف نهائي' }}
          </button>
        </v-card-actions>
      </v-card>
    </v-dialog>

    <!-- Delete SubCategory Confirmation Dialog -->
    <v-dialog v-model="showDeleteSubCategoryDialog" max-width="500">
      <v-card class="category-dialog">
        <v-card-title class="dialog-header" style="background: linear-gradient(135deg, #dc2626 0%, #b91c1c 100%) !important;">
          <i class="mdi mdi-delete-alert"></i>
          تأكيد حذف الصنف الثانوي
        </v-card-title>

        <v-card-text v-if="deletingSubCategory" class="dialog-content text-center">
          <v-avatar size="60" color="error" class="mb-3">
            <i class="mdi mdi-tag-outline" style="color: white; font-size: 30px;"></i>
          </v-avatar>
          <h4 style="color: rgba(255,255,255,0.95);">{{ deletingSubCategory.name }}</h4>
          <p v-if="deletingSubCategory.percentage" style="color: rgba(255,255,255,0.7); margin-top: 8px;">
            النسبة: {{ deletingSubCategory.percentage }}%
          </p>

          <v-alert type="error" variant="tonal" class="my-4" style="text-align: right; direction: rtl;">
            تحذير: هذا الإجراء لا يمكن التراجع عنه!
          </v-alert>

          <p style="color: rgba(255,255,255,0.7);">
            هل أنت متأكد من حذف هذا الصنف الثانوي نهائياً؟
          </p>
        </v-card-text>

        <v-card-actions class="dialog-actions">
          <v-spacer></v-spacer>
          <button class="dialog-btn cancel" @click="closeDeleteSubCategoryDialog">
            إلغاء
          </button>
          <button class="dialog-btn delete" @click="confirmDeleteSubCategory" :disabled="deletingSubCategoryLoading">
            <i class="mdi mdi-delete"></i>
            {{ deletingSubCategoryLoading ? 'جاري الحذف...' : 'حذف نهائي' }}
          </button>
        </v-card-actions>
      </v-card>
    </v-dialog>

  </div>
</template>

<script setup>
import { ref, shallowRef, computed, onMounted } from 'vue'
import PageHeader from '../components/PageHeader.vue'
import { listCategories, createCategory, updateCategory, deleteCategory as apiDeleteCategory, listSubCategories, createSubCategory, updateSubCategory, deleteSubCategory as apiDeleteSubCategory } from '@/api/categories'
import { usePermissions } from '@/composables/usePermissions'
import { useToast } from '@/composables/useToast'
import { DEFAULT_PAGE, DEFAULT_LIMIT } from '@/constants/pagination'

const { canCreate, canUpdate, canDelete } = usePermissions('/categories')
const { success, error: showError, warning } = useToast()

const loading = ref(false)
const showAddDialog = ref(false)
const formValid = ref(false)
const categoryForm = ref(null)
const editingCategory = ref(null)

// Pagination
const page = ref(DEFAULT_PAGE)
const limit = ref(DEFAULT_LIMIT)
const total = ref(0)

// New category data
const newCategory = ref({
  name: ''
})

const categories = ref([])

// All subcategories (fetched once)
const allSubCategories = ref([])

// Fetch categories from API
const fetchCategories = async () => {
  loading.value = true
  try {
    const [categoriesResponse, subCategoriesResponse] = await Promise.all([
      listCategories({ page: page.value, limit: limit.value }),
      listSubCategories({ limit: 500 })
    ])

    if (categoriesResponse.success) {
      categories.value = (categoriesResponse.data.data || []).map(c => ({
        id: c.id,
        name: c.name,
        description: c.description || '',
        status: c.status || 'active'
      }))
      total.value = categoriesResponse.data.total || 0
    }

    if (subCategoriesResponse.success) {
      allSubCategories.value = subCategoriesResponse.data.data || []
      // Group subcategories by categoryId
      subCategories.value = {}
      allSubCategories.value.forEach(sub => {
        if (!subCategories.value[sub.categoryId]) {
          subCategories.value[sub.categoryId] = []
        }
        subCategories.value[sub.categoryId].push(sub)
      })
    }
  } catch (error) {
    console.error('Error fetching categories:', error)
    showError('حدث خطأ في جلب التصنيفات')
  } finally {
    loading.value = false
  }
}

const headers = [
  { title: 'الاسم', key: 'name', align: 'start' },
  { title: 'الوصف', key: 'description', align: 'center' },
  { title: 'عدد الأصناف الثانوية', key: 'subCategoriesCount', align: 'center' },
  { title: 'الحالة', key: 'status', align: 'center' },
  { title: 'الإجراءات', key: 'actions', align: 'center', sortable: false }
]

// Get subcategories count for a category
const getSubCategoriesCount = (categoryId) => {
  return subCategories.value[categoryId]?.length || 0
}

// Computed properties
const totalCategories = computed(() => total.value || categories.value.length)

const activeCategories = computed(() => categories.value.filter(cat => cat.status === 'active').length)

const disabledCategories = computed(() => categories.value.filter(cat => cat.status !== 'active').length)

const totalSubCategoriesCount = computed(() => allSubCategories.value.length)

// Delete dialog state
const showDeleteDialog = ref(false)
const deleting = ref(false)

// Methods
const openAddDialog = () => {
  editingCategory.value = null
  newCategory.value = { name: '' }
  showAddDialog.value = true
}

const editCategory = (category) => {
  editingCategory.value = category
  newCategory.value = {
    name: category.name
  }
  showAddDialog.value = true
}

const closeDialog = () => {
  showAddDialog.value = false
  editingCategory.value = null
  resetForm()
}

const resetForm = () => {
  newCategory.value = {
    name: ''
  }
  if (categoryForm.value) {
    categoryForm.value.reset()
  }
}

const confirmDeleteCategory = (category) => {
  editingCategory.value = category
  showDeleteDialog.value = true
}

const deleteCategoryItem = async () => {
  if (!editingCategory.value) return

  // Check if category has subcategories
  if (getSubCategoriesCount(editingCategory.value.id) > 0) {
    showError('لا يمكن حذف التصنيف لوجود أصناف ثانوية مرتبطة به')
    return
  }

  deleting.value = true
  try {
    const response = await apiDeleteCategory(editingCategory.value.id)
    if (response.success) {
      success('تم حذف التصنيف بنجاح')
      showDeleteDialog.value = false
      editingCategory.value = null
      fetchCategories()
    } else {
      // Check if error is due to related records
      const errorMsg = response.message || 'حدث خطأ في الحذف'
      if (errorMsg.includes('foreign key') || errorMsg.includes('constraint') || errorMsg.includes('referenced')) {
        showError('لا يمكن حذف التصنيف لوجود سجلات مرتبطة به')
      } else {
        showError(errorMsg)
      }
    }
  } catch (error) {
    console.error('Error deleting category:', error)
    const errorMsg = error.message || ''
    if (errorMsg.includes('foreign key') || errorMsg.includes('constraint') || errorMsg.includes('referenced') || errorMsg.includes('violates')) {
      showError('لا يمكن حذف التصنيف لوجود سجلات مرتبطة به')
    } else {
      showError('حدث خطأ في حذف التصنيف')
    }
  } finally {
    deleting.value = false
  }
}

const saveCategory = async () => {
  if (!formValid.value) return

  loading.value = true
  try {
    const categoryData = {
      name: newCategory.value.name
    }

    let response
    if (editingCategory.value) {
      response = await updateCategory(editingCategory.value.id, categoryData)
    } else {
      response = await createCategory(categoryData)
    }

    if (response.success) {
      success(editingCategory.value ? 'تم تحديث التصنيف بنجاح' : 'تم إضافة التصنيف بنجاح')
      closeDialog()
      fetchCategories()
    } else {
      showError(response.message || 'حدث خطأ')
    }
  } catch (error) {
    console.error('Error saving category:', error)
    showError('حدث خطأ في حفظ التصنيف')
  } finally {
    loading.value = false
  }
}

// Initialize
onMounted(() => {
  fetchCategories()
})

const getStatusColor = (status) => {
  const colors = {
    active: 'success',
    disabled: 'warning'
  }
  return colors[status] || 'grey'
}

const getStatusText = (status) => {
  const texts = {
    active: 'نشط',
    disabled: 'معطل'
  }
  return texts[status] || status
}

// Sub-Categories State
const showSubCategoryDialog = ref(false)
const selectedCategory = ref(null)
const subCategories = ref({})

// New subcategory data with percentage - using shallowRef for Vuetify reactivity
const newSubCategoryName = shallowRef('')
const newSubCategoryPercentage = shallowRef(null)

// Edit subcategory state
const showEditSubCategoryDialog = ref(false)
const editSubCategoryData = ref({
  id: null,
  name: '',
  percentage: null,
  index: null
})
const savingSubCategory = ref(false)

// Delete subcategory state
const showDeleteSubCategoryDialog = ref(false)
const deletingSubCategory = ref(null)
const deletingSubCategoryIndex = ref(null)
const deletingSubCategoryLoading = ref(false)

// Computed for current sub-categories
const currentSubCategories = computed(() => {
  if (!selectedCategory.value) return []
  return subCategories.value[selectedCategory.value.id] || []
})

// Computed to check if add subcategory button should be disabled
const canAddSubCategory = computed(() => {
  const name = newSubCategoryName.value?.trim()
  const percentage = newSubCategoryPercentage.value
  const hasValidName = !!(name && name.length > 0)
  const hasValidPercentage = percentage !== null && percentage !== '' && Number(percentage) >= 0 && Number(percentage) <= 100
  return hasValidName && hasValidPercentage
})

// Sub-Categories Methods
const openSubCategoryDialog = (category) => {
  selectedCategory.value = category
  if (!subCategories.value[category.id]) {
    subCategories.value[category.id] = []
  }
  // Reset form
  newSubCategoryName.value = ''
  newSubCategoryPercentage.value = null
  showSubCategoryDialog.value = true
}

const closeSubCategoryDialog = () => {
  showSubCategoryDialog.value = false
  selectedCategory.value = null
  newSubCategoryName.value = ''
  newSubCategoryPercentage.value = null
}

const addSubCategory = async () => {
  console.log('addSubCategory called, name:', newSubCategoryName.value, 'category:', selectedCategory.value)
  const name = newSubCategoryName.value?.trim()
  if (!name) {
    warning('يرجى إدخال اسم الصنف الثانوي')
    return
  }
  if (!selectedCategory.value) {
    warning('يرجى اختيار تصنيف أولاً')
    return
  }

  try {
    const percentageValue = newSubCategoryPercentage.value !== null && newSubCategoryPercentage.value !== ''
      ? Number(newSubCategoryPercentage.value)
      : null
    const response = await createSubCategory({
      categoryId: selectedCategory.value.id,
      name: name,
      percentage: percentageValue
    })
    if (response.success) {
      // Add to local list
      if (!subCategories.value[selectedCategory.value.id]) {
        subCategories.value[selectedCategory.value.id] = []
      }
      const newSub = response.data || {
        name: name,
        percentage: percentageValue,
        id: Date.now(),
        categoryId: selectedCategory.value.id
      }
      subCategories.value[selectedCategory.value.id].push(newSub)
      allSubCategories.value.push(newSub)
      newSubCategoryName.value = ''
      newSubCategoryPercentage.value = null
      success('تم إضافة الصنف الثانوي بنجاح')
    } else {
      showError(response.message || 'حدث خطأ')
    }
  } catch (error) {
    console.error('Error adding sub category:', error)
    showError('حدث خطأ في إضافة الصنف الثانوي')
  }
}

// Edit subcategory methods
const openEditSubCategoryDialog = (sub, index) => {
  editSubCategoryData.value = {
    id: sub.id,
    name: sub.name,
    percentage: sub.percentage,
    index: index
  }
  showEditSubCategoryDialog.value = true
}

const closeEditSubCategoryDialog = () => {
  showEditSubCategoryDialog.value = false
  editSubCategoryData.value = { id: null, name: '', percentage: null, index: null }
}

const saveEditSubCategory = async () => {
  if (!editSubCategoryData.value.name || !editSubCategoryData.value.id) return

  savingSubCategory.value = true
  try {
    const response = await updateSubCategory(editSubCategoryData.value.id, {
      name: editSubCategoryData.value.name,
      percentage: editSubCategoryData.value.percentage || null
    })
    if (response.success) {
      // Update local list
      const idx = editSubCategoryData.value.index
      if (idx !== null && selectedCategory.value) {
        subCategories.value[selectedCategory.value.id][idx] = {
          ...subCategories.value[selectedCategory.value.id][idx],
          name: editSubCategoryData.value.name,
          percentage: editSubCategoryData.value.percentage
        }
        // Also update allSubCategories
        const allIdx = allSubCategories.value.findIndex(s => s.id === editSubCategoryData.value.id)
        if (allIdx !== -1) {
          allSubCategories.value[allIdx] = {
            ...allSubCategories.value[allIdx],
            name: editSubCategoryData.value.name,
            percentage: editSubCategoryData.value.percentage
          }
        }
      }
      success('تم تحديث الصنف الثانوي بنجاح')
      closeEditSubCategoryDialog()
    } else {
      showError(response.message || 'حدث خطأ')
    }
  } catch (error) {
    console.error('Error updating sub category:', error)
    showError('حدث خطأ في تحديث الصنف الثانوي')
  } finally {
    savingSubCategory.value = false
  }
}

// Open delete subcategory confirmation dialog
const deleteSubCategory = (index) => {
  if (!selectedCategory.value) return
  const subCat = subCategories.value[selectedCategory.value.id][index]
  deletingSubCategory.value = subCat
  deletingSubCategoryIndex.value = index
  showDeleteSubCategoryDialog.value = true
}

// Close delete subcategory dialog
const closeDeleteSubCategoryDialog = () => {
  showDeleteSubCategoryDialog.value = false
  deletingSubCategory.value = null
  deletingSubCategoryIndex.value = null
}

// Confirm and execute subcategory deletion
const confirmDeleteSubCategory = async () => {
  if (!selectedCategory.value || !deletingSubCategory.value) return

  deletingSubCategoryLoading.value = true
  try {
    const response = await apiDeleteSubCategory(deletingSubCategory.value.id)
    if (response.success) {
      subCategories.value[selectedCategory.value.id].splice(deletingSubCategoryIndex.value, 1)
      // Also remove from allSubCategories
      const allIdx = allSubCategories.value.findIndex(s => s.id === deletingSubCategory.value.id)
      if (allIdx !== -1) {
        allSubCategories.value.splice(allIdx, 1)
      }
      success('تم حذف الصنف الثانوي بنجاح')
      closeDeleteSubCategoryDialog()
    } else {
      // Check if error is due to related records
      const errorMsg = response.message || 'حدث خطأ'
      if (errorMsg.includes('foreign key') || errorMsg.includes('constraint') || errorMsg.includes('referenced')) {
        showError('لا يمكن حذف هذا الصنف الثانوي لوجود سجلات مرتبطة به')
      } else {
        showError(errorMsg)
      }
    }
  } catch (error) {
    console.error('Error deleting sub category:', error)
    const errorMsg = error.message || ''
    if (errorMsg.includes('foreign key') || errorMsg.includes('constraint') || errorMsg.includes('referenced') || errorMsg.includes('violates')) {
      showError('لا يمكن حذف هذا الصنف الثانوي لوجود سجلات مرتبطة به')
    } else {
      showError('حدث خطأ في حذف الصنف الثانوي')
    }
  } finally {
    deletingSubCategoryLoading.value = false
  }
}
</script>

<style scoped>
.categories-container {
  padding: 32px;
  max-width: 1600px;
  margin: 0 auto;
}

/* Categories Header Custom Color */
.categories-header {
  background: linear-gradient(135deg, #018790 0%, #005461 100%) !important;
}

.categories-header::before {
  background: linear-gradient(135deg, #06b6d4 0%, #10b981 50%, #06b6d4 100%) !important;
}

/* Statistics Grid */
.stats-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
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

.stat-icon.category {
  background: linear-gradient(135deg, #14b8a6 0%, #0d9488 100%) !important;
  box-shadow: 0 6px 16px rgba(20, 184, 166, 0.4);
}

.stat-icon.success {
  background: linear-gradient(135deg, #10b981 0%, #059669 100%) !important;
  box-shadow: 0 6px 16px rgba(16, 185, 129, 0.4);
}

.stat-icon.warning {
  background: linear-gradient(135deg, #f59e0b 0%, #d97706 100%) !important;
  box-shadow: 0 6px 16px rgba(245, 158, 11, 0.4);
}

.stat-icon.primary {
  background: linear-gradient(135deg, #3b82f6 0%, #2563eb 100%) !important;
  box-shadow: 0 6px 16px rgba(59, 130, 246, 0.4);
}

/* Categories List Header */
.categories-list-header {
  background: linear-gradient(135deg, #0a3d42 0%, #052428 100%);
  border-radius: 16px;
  margin-bottom: 24px;
  padding: 16px 24px;
  border: 2px solid transparent;
  position: relative;
  overflow: hidden;
}

.categories-list-header::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  border-radius: 16px;
  padding: 2px;
  background: linear-gradient(135deg, #14b8a6 0%, #0d9488 50%, #14b8a6 100%);
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
  color: #14b8a6;
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
  background: linear-gradient(135deg, #14b8a6 0%, #0d9488 100%);
  color: white;
  box-shadow: 0 4px 12px rgba(20, 184, 166, 0.3);
}

.list-action-btn.primary:hover {
  transform: translateY(-2px);
  box-shadow: 0 6px 16px rgba(20, 184, 166, 0.4);
}

.list-action-btn i {
  font-size: 16px;
}

/* Color Preview */
.color-preview {
  width: 32px;
  height: 32px;
  border-radius: 8px;
  margin: 0 auto;
  border: 2px solid rgba(255, 255, 255, 0.2);
}

/* Category Dialog */
.category-dialog {
  border-radius: 16px !important;
  background: linear-gradient(135deg, #0a3d42 0%, #052428 100%) !important;
  border: 2px solid transparent !important;
  position: relative;
  overflow: hidden;
  direction: rtl;
  text-align: right;
}

.category-dialog::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  border-radius: 16px;
  padding: 2px;
  background: linear-gradient(135deg, #14b8a6 0%, #0d9488 50%, #14b8a6 100%);
  -webkit-mask: linear-gradient(#fff 0 0) content-box, linear-gradient(#fff 0 0);
  mask: linear-gradient(#fff 0 0) content-box, linear-gradient(#fff 0 0);
  -webkit-mask-composite: xor;
  mask-composite: exclude;
  pointer-events: none;
}

.dialog-header {
  background: rgba(20, 184, 166, 0.1) !important;
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
  color: #14b8a6;
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
  color: #14b8a6 !important;
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
  background: linear-gradient(135deg, #14b8a6 0%, #0d9488 100%);
  color: white;
  box-shadow: 0 4px 12px rgba(20, 184, 166, 0.3);
}

.dialog-btn.save:hover:not(:disabled) {
  transform: translateY(-2px);
  box-shadow: 0 6px 16px rgba(20, 184, 166, 0.4);
}

.dialog-btn.save:disabled {
  opacity: 0.5;
  cursor: not-allowed;
}

.dialog-btn.delete {
  background: linear-gradient(135deg, #dc2626 0%, #b91c1c 100%);
  color: white;
  box-shadow: 0 4px 12px rgba(220, 38, 38, 0.3);
}

.dialog-btn.delete:hover:not(:disabled) {
  transform: translateY(-2px);
  box-shadow: 0 6px 16px rgba(220, 38, 38, 0.4);
}

.dialog-btn.delete:disabled {
  opacity: 0.5;
  cursor: not-allowed;
}

.dialog-btn i {
  font-size: 18px;
}

.v-card {
  border-radius: 12px;
}

/* Sub-Categories Styles */
.sub-category-form {
  margin-bottom: 24px;
  padding-bottom: 20px;
  border-bottom: 1px solid rgba(255, 255, 255, 0.1);
  background: rgba(255, 255, 255, 0.05);
  padding: 16px;
  border-radius: 12px;
}

.sub-category-form :deep(.v-field) {
  background: rgba(255, 255, 255, 0.1) !important;
}

.sub-category-form :deep(.v-field__outline) {
  color: rgba(255, 255, 255, 0.3) !important;
}

.sub-category-form :deep(.v-field--focused .v-field__outline) {
  color: #14b8a6 !important;
}

.sub-category-form :deep(.v-label) {
  color: rgba(255, 255, 255, 0.8) !important;
}

.sub-category-form :deep(.v-field__input) {
  color: white !important;
}

.sub-category-form :deep(.v-field__input::placeholder) {
  color: rgba(255, 255, 255, 0.5) !important;
}

.add-sub-btn {
  height: 42px !important;
  border-radius: 12px !important;
  font-weight: 600 !important;
  transition: all 0.3s ease !important;
}

.add-sub-btn:hover:not(:disabled) {
  transform: translateY(-2px);
  box-shadow: 0 6px 16px rgba(20, 184, 166, 0.4) !important;
}

.sub-categories-list {
  background: rgba(0, 0, 0, 0.2);
  border-radius: 12px;
  padding: 16px;
}

.sub-list-header {
  display: flex;
  align-items: center;
  gap: 8px;
  color: rgba(255, 255, 255, 0.9);
  font-size: 16px;
  font-weight: 600;
  margin-bottom: 16px;
  padding-bottom: 12px;
  border-bottom: 1px solid rgba(255, 255, 255, 0.1);
}

.sub-list-header i {
  color: #14b8a6;
  font-size: 20px;
}

.empty-sub-list {
  text-align: center;
  padding: 32px 16px;
}

.empty-sub-list i {
  font-size: 48px;
  color: rgba(255, 255, 255, 0.3);
  margin-bottom: 12px;
}

.empty-sub-list p {
  color: rgba(255, 255, 255, 0.5);
  font-size: 14px;
  margin: 0;
}

.sub-items-container {
  display: flex;
  flex-direction: column;
  gap: 10px;
  max-height: 300px;
  overflow-y: auto;
}

.sub-category-item {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 12px 16px;
  background: rgba(255, 255, 255, 0.05);
  border-radius: 10px;
  border: 1px solid rgba(255, 255, 255, 0.1);
  transition: all 0.3s ease;
}

.sub-category-item:hover {
  background: rgba(255, 255, 255, 0.1);
  border-color: rgba(20, 184, 166, 0.3);
}

.sub-item-info {
  display: flex;
  align-items: center;
  gap: 10px;
  color: rgba(255, 255, 255, 0.9);
  font-size: 14px;
  flex: 1;
}

.sub-item-info i {
  color: #14b8a6;
  font-size: 18px;
}

.sub-item-name {
  font-weight: 500;
}

.sub-item-actions {
  display: flex;
  gap: 8px;
}

.sub-action-btn {
  width: 32px;
  height: 32px;
  border: none;
  border-radius: 8px;
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: center;
  transition: all 0.3s ease;
}

.sub-action-btn i {
  font-size: 16px;
}

.sub-action-btn.edit {
  background: rgba(59, 130, 246, 0.2);
  color: #3b82f6;
}

.sub-action-btn.edit:hover {
  background: rgba(59, 130, 246, 0.4);
}

.sub-action-btn.delete {
  background: rgba(239, 68, 68, 0.2);
  color: #ef4444;
}

.sub-action-btn.delete:hover {
  background: rgba(239, 68, 68, 0.4);
}

/* Responsive Styles */
@media (max-width: 1024px) {
  .categories-container {
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
  .categories-container {
    padding: 16px;
  }

  .stats-grid {
    grid-template-columns: 1fr 1fr;
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

  .categories-list-header {
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
  .category-dialog .dialog-header {
    padding: 16px !important;
    font-size: 18px !important;
  }

  .category-dialog .dialog-content {
    padding: 16px !important;
  }

  .category-dialog .dialog-actions {
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

  /* Sub-category form */
  .sub-category-form .v-row {
    margin: 0;
  }

  .sub-category-form .v-col {
    padding: 6px !important;
  }
}

@media (max-width: 480px) {
  .categories-container {
    padding: 12px;
  }

  .stats-grid {
    grid-template-columns: 1fr;
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

  .categories-list-header {
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

  .category-dialog {
    border-radius: 12px !important;
  }

  .category-dialog .dialog-header {
    padding: 12px 16px !important;
    font-size: 16px !important;
    gap: 8px;
  }

  .category-dialog .dialog-header i {
    font-size: 20px;
  }

  .category-dialog .dialog-content {
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

  /* Sub-category form mobile */
  .sub-category-form {
    padding: 12px;
  }

  .sub-category-form .v-row {
    flex-direction: column;
  }

  .sub-category-form .v-col {
    max-width: 100% !important;
    flex: 0 0 100% !important;
    padding: 4px 0 !important;
  }

  .add-sub-btn {
    margin-top: 8px;
  }

  .sub-categories-list {
    padding: 12px;
  }

  .sub-list-header {
    font-size: 14px;
    margin-bottom: 12px;
    padding-bottom: 8px;
  }

  .sub-category-item {
    padding: 10px 12px;
  }

  .sub-item-info {
    font-size: 13px;
    gap: 8px;
  }

  .sub-item-info i {
    font-size: 16px;
  }

  .sub-action-btn {
    width: 28px;
    height: 28px;
  }

  .sub-action-btn i {
    font-size: 14px;
  }
}
</style>
