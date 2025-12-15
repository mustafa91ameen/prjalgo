import { ref, computed } from 'vue'
import { toast } from 'vue3-toastify'
import { expensesApi } from '@/api/expenses'

/**
 * Composable for project expenses management logic
 * Handles administrative expenses CRUD, filtering, and dialog states
 *
 * @returns {Object} Expenses management utilities
 */
export function useProjectExpenses() {
  // State
  const expenses = ref([])
  const loading = ref(false)
  const error = ref(null)
  const showExpenseDialog = ref(false)
  const showDeleteDialog = ref(false)
  const isEditing = ref(false)
  const selectedExpense = ref(null)
  const expenseFormValid = ref(false)

  // Filter state
  const searchQuery = ref('')
  const selectedProjectFilter = ref('')
  const selectedCostRange = ref('')

  // Form state
  const expenseForm = ref({
    projectName: '',
    startDate: '',
    endDate: '',
    cost: '',
    workLocation: '',
    expenseType: '',
    notes: '',
  })

  // Options
  const expenseTypes = ['تطوير', 'تحديث', 'بناء', 'أمن', 'صيانة', 'تدريب', 'أخرى']

  const costRangeOptions = [
    { title: 'أقل من 50,000 د.ع', value: 'low' },
    { title: '50,000 - 100,000 د.ع', value: 'medium' },
    { title: 'أكثر من 100,000 د.ع', value: 'high' },
  ]

  // Table headers
  const expenseHeaders = [
    { title: 'اسم المشروع', key: 'projectName', sortable: true },
    { title: 'تاريخ البداية', key: 'startDate', sortable: true, width: '120px' },
    { title: 'تاريخ الانتهاء', key: 'endDate', sortable: true, width: '120px' },
    { title: 'التكلفة', key: 'cost', sortable: true, width: '120px' },
    { title: 'مكان العمل', key: 'workLocation', sortable: true, width: '120px' },
    { title: 'الملاحظات', key: 'notes', sortable: false, width: '200px' },
    { title: 'الإجراءات', key: 'actions', sortable: false, width: '120px' },
  ]

  // Computed
  const projectFilterOptions = computed(() => {
    const projects = [...new Set(expenses.value.map((expense) => expense.projectName))]
    return projects.map((project) => ({ title: project, value: project }))
  })

  const filteredExpenses = computed(() => {
    let filtered = expenses.value

    // Filter by search query
    if (searchQuery.value) {
      const query = searchQuery.value.toLowerCase()
      filtered = filtered.filter(
        (expense) =>
          expense.projectName?.toLowerCase().includes(query) ||
          expense.workLocation?.toLowerCase().includes(query) ||
          expense.notes?.toLowerCase().includes(query)
      )
    }

    // Filter by project
    if (selectedProjectFilter.value) {
      filtered = filtered.filter((expense) => expense.projectName === selectedProjectFilter.value)
    }

    // Filter by cost range
    if (selectedCostRange.value) {
      filtered = filtered.filter((expense) => {
        const cost = expense.cost
        switch (selectedCostRange.value) {
          case 'low':
            return cost < 50000
          case 'medium':
            return cost >= 50000 && cost <= 100000
          case 'high':
            return cost > 100000
          default:
            return true
        }
      })
    }

    return filtered
  })

  const totalExpenses = computed(() => expenses.value.reduce((sum, expense) => sum + (expense.cost || 0), 0))

  const expenseCount = computed(() => expenses.value.length)

  // Form management
  function resetForm() {
    expenseForm.value = {
      projectName: '',
      startDate: '',
      endDate: '',
      cost: '',
      workLocation: '',
      expenseType: '',
      notes: '',
    }
    expenseFormValid.value = false
    isEditing.value = false
    selectedExpense.value = null
  }

  function openAddDialog() {
    resetForm()
    showExpenseDialog.value = true
  }

  function closeExpenseDialog() {
    showExpenseDialog.value = false
    resetForm()
  }

  function openEditDialog(expense) {
    selectedExpense.value = expense
    expenseForm.value = { ...expense }
    isEditing.value = true
    showExpenseDialog.value = true
  }

  function openDeleteDialog(expense) {
    selectedExpense.value = expense
    showDeleteDialog.value = true
  }

  function closeDeleteDialog() {
    showDeleteDialog.value = false
    selectedExpense.value = null
  }

  // Reset filters
  function resetFilters() {
    searchQuery.value = ''
    selectedProjectFilter.value = ''
    selectedCostRange.value = ''
  }

  // API Operations
  async function fetchExpenses(params = {}) {
    loading.value = true
    error.value = null

    try {
      const response = await expensesApi.getAll(params)
      expenses.value = response.data ?? response ?? []
    } catch (err) {
      error.value = err.message
      expenses.value = []
    } finally {
      loading.value = false
    }
  }

  async function createExpense() {
    loading.value = true
    error.value = null

    try {
      const newExpense = {
        ...expenseForm.value,
        cost: parseFloat(expenseForm.value.cost),
      }

      const response = await expensesApi.create(newExpense)
      const createdExpense = response.data ?? response

      expenses.value.push(createdExpense)
      toast.success('تم إضافة المصروف بنجاح')
      closeExpenseDialog()
      return createdExpense
    } catch (err) {
      error.value = err.message
      toast.error(err.message || 'فشل في إضافة المصروف')
      return null
    } finally {
      loading.value = false
    }
  }

  async function updateExpense() {
    if (!selectedExpense.value) return null

    loading.value = true
    error.value = null

    try {
      const updatedData = {
        ...expenseForm.value,
        cost: parseFloat(expenseForm.value.cost),
      }

      const response = await expensesApi.update(selectedExpense.value.id, updatedData)
      const updatedExpense = response.data ?? response

      const index = expenses.value.findIndex((e) => e.id === selectedExpense.value.id)
      if (index !== -1) {
        expenses.value[index] = updatedExpense
      }

      toast.success('تم تحديث المصروف بنجاح')
      closeExpenseDialog()
      return updatedExpense
    } catch (err) {
      error.value = err.message
      toast.error(err.message || 'فشل في تحديث المصروف')
      return null
    } finally {
      loading.value = false
    }
  }

  async function deleteExpense() {
    if (!selectedExpense.value) return false

    loading.value = true
    error.value = null

    try {
      await expensesApi.delete(selectedExpense.value.id)

      expenses.value = expenses.value.filter((e) => e.id !== selectedExpense.value.id)
      toast.success('تم حذف المصروف بنجاح')
      closeDeleteDialog()
      return true
    } catch (err) {
      error.value = err.message
      toast.error(err.message || 'فشل في حذف المصروف')
      return false
    } finally {
      loading.value = false
    }
  }

  async function saveExpense() {
    if (isEditing.value) {
      return updateExpense()
    }
    return createExpense()
  }

  // Local state management (for demo/sample data)
  function addLocalExpense(expense) {
    const newExpense = {
      ...expense,
      id: Date.now(),
      cost: parseFloat(expense.cost),
    }
    expenses.value.push(newExpense)
    return newExpense
  }

  function updateLocalExpense(id, data) {
    const index = expenses.value.findIndex((e) => e.id === id)
    if (index !== -1) {
      expenses.value[index] = {
        ...expenses.value[index],
        ...data,
        cost: parseFloat(data.cost),
      }
      return expenses.value[index]
    }
    return null
  }

  function deleteLocalExpense(id) {
    const index = expenses.value.findIndex((e) => e.id === id)
    if (index !== -1) {
      expenses.value.splice(index, 1)
      return true
    }
    return false
  }

  function setExpenses(data) {
    expenses.value = data
  }

  return {
    // State
    expenses,
    loading,
    error,
    showExpenseDialog,
    showDeleteDialog,
    isEditing,
    selectedExpense,
    expenseForm,
    expenseFormValid,
    // Filter state
    searchQuery,
    selectedProjectFilter,
    selectedCostRange,
    // Options
    expenseTypes,
    costRangeOptions,
    expenseHeaders,
    // Computed
    projectFilterOptions,
    filteredExpenses,
    totalExpenses,
    expenseCount,
    // Form management
    resetForm,
    openAddDialog,
    closeExpenseDialog,
    openEditDialog,
    openDeleteDialog,
    closeDeleteDialog,
    resetFilters,
    // API operations
    fetchExpenses,
    createExpense,
    updateExpense,
    deleteExpense,
    saveExpense,
    // Local state management
    addLocalExpense,
    updateLocalExpense,
    deleteLocalExpense,
    setExpenses,
  }
}

export default useProjectExpenses
