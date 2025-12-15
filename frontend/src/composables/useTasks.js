import { ref, computed } from 'vue'

/**
 * Composable for project tasks/actions management
 * Handles task status, working days, and related utilities
 *
 * @returns {Object} Task management utilities
 */
export function useTasks() {
  // State
  const tasks = ref([])
  const workingDays = ref([])
  const loading = ref(false)
  const selectedTask = ref(null)

  // Task status options
  const taskStatuses = [
    { title: 'في الانتظار', value: 'pending' },
    { title: 'قيد التنفيذ', value: 'in-progress' },
    { title: 'مكتمل', value: 'completed' },
    { title: 'ملغي', value: 'cancelled' },
  ]

  // Default working days template
  const defaultWorkingDays = [
    { name: 'السبت', hours: '8 ساعات', isActive: true },
    { name: 'الأحد', hours: '8 ساعات', isActive: true },
    { name: 'الاثنين', hours: '8 ساعات', isActive: true },
    { name: 'الثلاثاء', hours: '8 ساعات', isActive: true },
    { name: 'الأربعاء', hours: '8 ساعات', isActive: true },
    { name: 'الخميس', hours: '4 ساعات', isActive: true },
    { name: 'الجمعة', hours: '0 ساعة', isActive: false },
  ]

  // Default tasks template
  const defaultTasks = [
    { id: 1, title: 'تحديد المتطلبات', description: 'تحديد متطلبات المشروع الأساسية', status: 'pending', dueDate: null },
    { id: 2, title: 'التخطيط', description: 'وضع خطة زمنية للمشروع', status: 'pending', dueDate: null },
    { id: 3, title: 'التنفيذ', description: 'بدء تنفيذ المشروع', status: 'pending', dueDate: null },
    { id: 4, title: 'الاختبار', description: 'اختبار المشروع والتأكد من جودته', status: 'pending', dueDate: null },
  ]

  // Computed
  const totalTasks = computed(() => tasks.value.length)

  const completedTasks = computed(() => tasks.value.filter((task) => task.status === 'completed').length)

  const pendingTasks = computed(() => tasks.value.filter((task) => task.status === 'pending').length)

  const inProgressTasks = computed(() => tasks.value.filter((task) => task.status === 'in-progress').length)

  const taskProgress = computed(() => {
    if (tasks.value.length === 0) return 0
    return Math.round((completedTasks.value / tasks.value.length) * 100)
  })

  const totalWorkingHours = computed(() => {
    return workingDays.value.reduce((total, day) => {
      if (day.isActive) {
        const hours = parseInt(day.hours) || 0
        return total + hours
      }
      return total
    }, 0)
  })

  const activeDaysCount = computed(() => workingDays.value.filter((day) => day.isActive).length)

  // Helper functions
  function getStatusText(status) {
    const statusTexts = {
      pending: 'في الانتظار',
      'in-progress': 'قيد التنفيذ',
      completed: 'مكتمل',
      cancelled: 'ملغي',
    }
    return statusTexts[status] || status
  }

  function getStatusColor(status) {
    const statusColors = {
      pending: 'grey',
      'in-progress': 'warning',
      completed: 'success',
      cancelled: 'error',
    }
    return statusColors[status] || 'grey'
  }

  function getStatusIcon(status) {
    const statusIcons = {
      pending: 'mdi-clock-outline',
      'in-progress': 'mdi-progress-clock',
      completed: 'mdi-check-circle',
      cancelled: 'mdi-cancel',
    }
    return statusIcons[status] || 'mdi-help-circle'
  }

  // Working days management
  function getDefaultWorkingDays() {
    return JSON.parse(JSON.stringify(defaultWorkingDays))
  }

  function setWorkingDays(days) {
    workingDays.value = days || getDefaultWorkingDays()
  }

  function updateWorkingDay(index, data) {
    if (workingDays.value[index]) {
      workingDays.value[index] = { ...workingDays.value[index], ...data }
    }
  }

  function toggleWorkingDay(index) {
    if (workingDays.value[index]) {
      workingDays.value[index].isActive = !workingDays.value[index].isActive
    }
  }

  // Task management
  function getDefaultTasks() {
    return JSON.parse(JSON.stringify(defaultTasks))
  }

  function setTasks(taskList) {
    tasks.value = taskList || getDefaultTasks()
  }

  function addTask(task) {
    const newTask = {
      ...task,
      id: Date.now(),
      status: task.status || 'pending',
    }
    tasks.value.push(newTask)
    return newTask
  }

  function updateTask(id, data) {
    const index = tasks.value.findIndex((t) => t.id === id)
    if (index !== -1) {
      tasks.value[index] = { ...tasks.value[index], ...data }
      return tasks.value[index]
    }
    return null
  }

  function deleteTask(id) {
    const index = tasks.value.findIndex((t) => t.id === id)
    if (index !== -1) {
      tasks.value.splice(index, 1)
      return true
    }
    return false
  }

  function updateTaskStatus(id, status) {
    return updateTask(id, { status })
  }

  function markTaskComplete(id) {
    return updateTaskStatus(id, 'completed')
  }

  function markTaskInProgress(id) {
    return updateTaskStatus(id, 'in-progress')
  }

  // Export working days to CSV
  function exportWorkingDaysToCSV(projectName = 'مشروع') {
    try {
      const days = workingDays.value.length > 0 ? workingDays.value : getDefaultWorkingDays()

      const excelData = [
        ['اسم المشروع', projectName],
        ['تاريخ التصدير', new Date().toLocaleDateString('ar-SA')],
        [''],
        ['يوم العمل', 'الحالة', 'عدد الساعات', 'ملاحظات'],
      ]

      days.forEach((day) => {
        excelData.push([day.name, day.isActive ? 'نشط' : 'غير نشط', day.hours, day.isActive ? 'يوم عمل عادي' : 'يوم إجازة'])
      })

      excelData.push(['', '', '', ''])
      excelData.push(['إجمالي ساعات العمل', '', totalWorkingHours.value + ' ساعة', ''])

      const csvContent = excelData.map((row) => row.map((cell) => `"${cell}"`).join(',')).join('\n')

      const BOM = '\uFEFF'
      const blob = new Blob([BOM + csvContent], { type: 'text/csv;charset=utf-8;' })

      const link = document.createElement('a')
      const url = URL.createObjectURL(blob)
      link.setAttribute('href', url)
      link.setAttribute('download', `أيام_العمل_${projectName.replace(/\s+/g, '_')}_${new Date().toISOString().split('T')[0]}.csv`)
      link.style.visibility = 'hidden'
      document.body.appendChild(link)
      link.click()
      document.body.removeChild(link)

      return true
    } catch {
      return false
    }
  }

  // Reset state
  function reset() {
    tasks.value = []
    workingDays.value = []
    selectedTask.value = null
  }

  return {
    // State
    tasks,
    workingDays,
    loading,
    selectedTask,
    // Options
    taskStatuses,
    // Computed
    totalTasks,
    completedTasks,
    pendingTasks,
    inProgressTasks,
    taskProgress,
    totalWorkingHours,
    activeDaysCount,
    // Helpers
    getStatusText,
    getStatusColor,
    getStatusIcon,
    // Working days
    getDefaultWorkingDays,
    setWorkingDays,
    updateWorkingDay,
    toggleWorkingDay,
    // Task management
    getDefaultTasks,
    setTasks,
    addTask,
    updateTask,
    deleteTask,
    updateTaskStatus,
    markTaskComplete,
    markTaskInProgress,
    // Export
    exportWorkingDaysToCSV,
    // Reset
    reset,
  }
}

export default useTasks
