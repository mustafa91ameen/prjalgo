import { ref, computed } from 'vue'
import { toast } from 'vue3-toastify'
import { teamsApi } from '@/api/teams'

/**
 * Composable for team management logic
 * Handles team members CRUD operations, dialog states, and computed statistics
 *
 * @returns {Object} Team management utilities
 */
export function useTeamManagement() {
  // State
  const teamMembers = ref([])
  const loading = ref(false)
  const error = ref(null)
  const showTeamSection = ref(false)
  const showAddMemberDialog = ref(false)
  const showEditMemberDialog = ref(false)
  const showDeleteMemberDialog = ref(false)
  const selectedMember = ref(null)
  const memberFormValid = ref(false)

  // Form state
  const memberForm = ref({
    name: '',
    email: '',
    phone: '',
    department: '',
    role: '',
    status: 'نشط',
  })

  // Options
  const departmentOptions = [
    'الهندسة المدنية',
    'الهندسة المعمارية',
    'الهندسة الكهربائية',
    'الهندسة الميكانيكية',
    'إدارة المشاريع',
    'الموارد البشرية',
    'المحاسبة',
    'التسويق',
  ]

  const roleOptions = [
    'مدير المشروع',
    'مهندس أول',
    'مهندس',
    'مساعد مهندس',
    'فني',
    'مساعد إداري',
    'محاسب',
    'مستشار',
  ]

  const statusOptions = ['نشط', 'غير نشط', 'في إجازة', 'مستقيل']

  // Table headers
  const teamHeaders = [
    { title: 'الصورة', key: 'avatar', sortable: false, width: '60px' },
    { title: 'الاسم', key: 'name', sortable: true },
    { title: 'القسم', key: 'department', sortable: true },
    { title: 'الدور', key: 'role', sortable: true },
    { title: 'الحالة', key: 'status', sortable: true },
    { title: 'المهام', key: 'tasks', sortable: true },
    { title: 'الإجراءات', key: 'actions', sortable: false, width: '120px' },
  ]

  // Computed statistics
  const activeTeamMembers = computed(() => teamMembers.value.filter((member) => member.status === 'نشط').length)

  const teamDepartments = computed(() => new Set(teamMembers.value.map((member) => member.department)).size)

  const totalTeamTasks = computed(() => teamMembers.value.reduce((sum, member) => sum + (member.tasksCount || 0), 0))

  const teamMemberCount = computed(() => teamMembers.value.length)

  // Helper functions
  function getDepartmentColor(department) {
    const colors = {
      'الهندسة المدنية': 'blue',
      'الهندسة المعمارية': 'purple',
      'الهندسة الكهربائية': 'orange',
      'الهندسة الميكانيكية': 'green',
      'إدارة المشاريع': 'red',
      'الموارد البشرية': 'pink',
      المحاسبة: 'teal',
      التسويق: 'indigo',
    }
    return colors[department] || 'grey'
  }

  function getRoleColor(role) {
    if (role?.includes('مدير')) return 'red'
    if (role?.includes('أول')) return 'blue'
    if (role?.includes('مهندس')) return 'green'
    return 'grey'
  }

  // Form management
  function resetForm() {
    memberForm.value = {
      name: '',
      email: '',
      phone: '',
      department: '',
      role: '',
      status: 'نشط',
    }
    memberFormValid.value = false
  }

  function openAddDialog() {
    resetForm()
    showAddMemberDialog.value = true
  }

  function closeAddDialog() {
    showAddMemberDialog.value = false
    resetForm()
  }

  function openEditDialog(member) {
    selectedMember.value = member
    memberForm.value = { ...member }
    showEditMemberDialog.value = true
  }

  function closeEditDialog() {
    showEditMemberDialog.value = false
    selectedMember.value = null
    resetForm()
  }

  function openDeleteDialog(member) {
    selectedMember.value = member
    showDeleteMemberDialog.value = true
  }

  function closeDeleteDialog() {
    showDeleteMemberDialog.value = false
    selectedMember.value = null
  }

  // Toggle team section visibility
  function toggleTeamSection() {
    showTeamSection.value = !showTeamSection.value
  }

  // API Operations
  async function fetchTeamMembers(projectId = null) {
    loading.value = true
    error.value = null

    try {
      const params = projectId ? { projectId } : {}
      const response = await teamsApi.getAll(params)
      teamMembers.value = response.data ?? response ?? []
    } catch (err) {
      error.value = err.message
      teamMembers.value = []
    } finally {
      loading.value = false
    }
  }

  async function addTeamMember() {
    if (!memberFormValid.value) return null

    loading.value = true
    error.value = null

    try {
      const newMember = {
        ...memberForm.value,
        tasksCount: 0,
        avatar: null,
      }

      const response = await teamsApi.create(newMember)
      const createdMember = response.data ?? response

      teamMembers.value.push(createdMember)
      toast.success('تم إضافة العضو بنجاح')
      closeAddDialog()
      return createdMember
    } catch (err) {
      error.value = err.message
      toast.error(err.message || 'فشل في إضافة العضو')
      return null
    } finally {
      loading.value = false
    }
  }

  async function updateTeamMember() {
    if (!selectedMember.value) return null

    loading.value = true
    error.value = null

    try {
      const updatedData = { ...memberForm.value }
      // Note: teamsApi doesn't have update method, this is for local state
      const index = teamMembers.value.findIndex((m) => m.id === selectedMember.value.id)
      if (index !== -1) {
        teamMembers.value[index] = { ...teamMembers.value[index], ...updatedData }
      }

      toast.success('تم تحديث بيانات العضو بنجاح')
      closeEditDialog()
      return teamMembers.value[index]
    } catch (err) {
      error.value = err.message
      toast.error(err.message || 'فشل في تحديث بيانات العضو')
      return null
    } finally {
      loading.value = false
    }
  }

  async function deleteTeamMember() {
    if (!selectedMember.value) return false

    loading.value = true
    error.value = null

    try {
      await teamsApi.delete(selectedMember.value.id)

      teamMembers.value = teamMembers.value.filter((m) => m.id !== selectedMember.value.id)
      toast.success('تم حذف العضو بنجاح')
      closeDeleteDialog()
      return true
    } catch (err) {
      error.value = err.message
      toast.error(err.message || 'فشل في حذف العضو')
      return false
    } finally {
      loading.value = false
    }
  }

  // Local state management (for demo/sample data)
  function addLocalMember(member) {
    const newMember = {
      ...member,
      id: Date.now(),
      tasksCount: 0,
      avatar: null,
    }
    teamMembers.value.push(newMember)
    return newMember
  }

  function updateLocalMember(id, data) {
    const index = teamMembers.value.findIndex((m) => m.id === id)
    if (index !== -1) {
      teamMembers.value[index] = { ...teamMembers.value[index], ...data }
      return teamMembers.value[index]
    }
    return null
  }

  function deleteLocalMember(id) {
    const index = teamMembers.value.findIndex((m) => m.id === id)
    if (index !== -1) {
      teamMembers.value.splice(index, 1)
      return true
    }
    return false
  }

  function setTeamMembers(members) {
    teamMembers.value = members
  }

  return {
    // State
    teamMembers,
    loading,
    error,
    showTeamSection,
    showAddMemberDialog,
    showEditMemberDialog,
    showDeleteMemberDialog,
    selectedMember,
    memberForm,
    memberFormValid,
    // Options
    departmentOptions,
    roleOptions,
    statusOptions,
    teamHeaders,
    // Computed
    activeTeamMembers,
    teamDepartments,
    totalTeamTasks,
    teamMemberCount,
    // Helpers
    getDepartmentColor,
    getRoleColor,
    // Form management
    resetForm,
    openAddDialog,
    closeAddDialog,
    openEditDialog,
    closeEditDialog,
    openDeleteDialog,
    closeDeleteDialog,
    toggleTeamSection,
    // API operations
    fetchTeamMembers,
    addTeamMember,
    updateTeamMember,
    deleteTeamMember,
    // Local state management
    addLocalMember,
    updateLocalMember,
    deleteLocalMember,
    setTeamMembers,
  }
}

export default useTeamManagement
