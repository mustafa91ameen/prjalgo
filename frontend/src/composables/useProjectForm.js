import { ref, computed } from 'vue'
import { useRouter } from 'vue-router'

/**
 * Composable for project form management
 * Handles project add/edit dialogs, form state, and validation
 *
 * @returns {Object} Project form utilities
 */
export function useProjectForm() {
  const router = useRouter()

  // Dialog states
  const showFormDialog = ref(false)
  const showDeleteDialog = ref(false)
  const showDetailsDialog = ref(false)
  const isEditing = ref(false)
  const formValid = ref(false)

  // Selected project state
  const selectedProject = ref(null)
  const selectedProjectDetails = ref(null)

  // Default form state
  const defaultFormState = {
    name: '',
    type: '',
    location: '',
    initialCost: 0,
    duration: 0,
    startDate: '',
    criticalCost: 0,
    phone: '',
    user: '',
    status: 'pending',
    priority: 'medium',
    progress: 0,
    description: '',
    category: '',
  }

  // Form state
  const projectForm = ref({ ...defaultFormState })

  // Options
  const priorityOptions = [
    { title: 'منخفضة', value: 'low' },
    { title: 'متوسطة', value: 'medium' },
    { title: 'عالية', value: 'high' },
    { title: 'عاجلة', value: 'urgent' },
  ]

  const statusOptions = [
    { title: 'في الانتظار', value: 'pending' },
    { title: 'نشط', value: 'active' },
    { title: 'مكتمل', value: 'completed' },
    { title: 'ملغي', value: 'cancelled' },
  ]

  const categoryOptions = [
    { title: 'تطوير', value: 'تطوير' },
    { title: 'تحديث', value: 'تحديث' },
    { title: 'بناء', value: 'بناء' },
    { title: 'صيانة', value: 'صيانة' },
    { title: 'استشارات', value: 'استشارات' },
  ]

  // Computed
  const dialogTitle = computed(() => {
    return isEditing.value ? 'تعديل المشروع' : 'إضافة مشروع جديد'
  })

  const submitButtonText = computed(() => {
    return isEditing.value ? 'حفظ التغييرات' : 'إضافة المشروع'
  })

  // Form management functions
  function resetForm() {
    projectForm.value = { ...defaultFormState }
    formValid.value = false
    isEditing.value = false
    selectedProject.value = null
  }

  function openAddDialog() {
    resetForm()
    showFormDialog.value = true
  }

  function openEditDialog(project) {
    isEditing.value = true
    selectedProject.value = project
    projectForm.value = { ...project }
    showFormDialog.value = true
  }

  function closeFormDialog() {
    showFormDialog.value = false
    resetForm()
  }

  function openDeleteDialog(project) {
    selectedProject.value = project
    showDeleteDialog.value = true
  }

  function closeDeleteDialog() {
    showDeleteDialog.value = false
    selectedProject.value = null
  }

  function openDetailsDialog(project) {
    selectedProjectDetails.value = project
    showDetailsDialog.value = true
  }

  function closeDetailsDialog() {
    showDetailsDialog.value = false
    selectedProjectDetails.value = null
  }

  // Navigation to project details
  function viewProjectDetails(project) {
    router.push('/work-days')
  }

  function editProjectFromDetails() {
    if (selectedProjectDetails.value) {
      openEditDialog(selectedProjectDetails.value)
      closeDetailsDialog()
    }
  }

  // Prepare form data for submission
  function getFormData() {
    return {
      ...projectForm.value,
      initialCost: parseFloat(projectForm.value.initialCost) || 0,
      criticalCost: parseFloat(projectForm.value.criticalCost) || 0,
      duration: parseInt(projectForm.value.duration) || 0,
      progress: parseInt(projectForm.value.progress) || 0,
    }
  }

  // Set form data from an existing project
  function setFormData(project) {
    projectForm.value = {
      ...defaultFormState,
      ...project,
    }
  }

  return {
    // Dialog states
    showFormDialog,
    showDeleteDialog,
    showDetailsDialog,
    isEditing,
    formValid,
    // Selected states
    selectedProject,
    selectedProjectDetails,
    // Form state
    projectForm,
    // Options
    priorityOptions,
    statusOptions,
    categoryOptions,
    // Computed
    dialogTitle,
    submitButtonText,
    // Form management
    resetForm,
    openAddDialog,
    openEditDialog,
    closeFormDialog,
    openDeleteDialog,
    closeDeleteDialog,
    openDetailsDialog,
    closeDetailsDialog,
    viewProjectDetails,
    editProjectFromDetails,
    getFormData,
    setFormData,
  }
}

export default useProjectForm
