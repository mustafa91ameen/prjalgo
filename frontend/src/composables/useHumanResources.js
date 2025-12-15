import { ref, computed } from 'vue'
import { usePermissions } from './usePermissions'

/**
 * Composable for Human Resources management
 */
export function useHumanResources() {
  const { canRead, canWrite, canDelete } = usePermissions()

  // Permission checks
  const canReadHR = canRead('/human-resources')
  const canWriteHR = canWrite('/human-resources')
  const canDeleteHR = canDelete('/human-resources')

  // State
  const employees = ref([])
  const loading = ref(false)
  const error = ref(null)

  // Computed statistics
  const totalEmployees = computed(() => employees.value.length)

  const activeEmployees = computed(() => {
    return employees.value.filter(emp => emp.status === 'نشط').length
  })

  const totalDepartments = computed(() => {
    return new Set(employees.value.map(emp => emp.department)).size
  })

  const onLeave = computed(() => {
    return employees.value.filter(emp => emp.status === 'في إجازة').length
  })

  const totalSalaries = computed(() => {
    return employees.value.reduce((sum, emp) => sum + (emp.salary || 0), 0)
  })

  const averageSalary = computed(() => {
    if (employees.value.length === 0) return 0
    return totalSalaries.value / employees.value.length
  })

  // Department and position options from data
  const departmentOptions = computed(() => {
    const depts = [...new Set(employees.value.map(emp => emp.department))]
    return depts.map(dept => ({ title: dept, value: dept }))
  })

  const positionOptions = computed(() => {
    const positions = [...new Set(employees.value.map(emp => emp.position))]
    return positions.map(pos => ({ title: pos, value: pos }))
  })

  // Filter employees
  const filterEmployees = (filters) => {
    let filtered = [...employees.value]

    if (filters.department) {
      filtered = filtered.filter(emp => emp.department === filters.department)
    }

    if (filters.status) {
      filtered = filtered.filter(emp => emp.status === filters.status)
    }

    if (filters.position) {
      filtered = filtered.filter(emp => emp.position === filters.position)
    }

    return filtered
  }

  // CRUD operations
  const fetchEmployees = async () => {
    if (!canReadHR.value) {
      console.warn('No permission to read employees')
      return
    }
    loading.value = true
    try {
      // Will be replaced with API call
      // const response = await hrApi.getEmployees()
      // employees.value = response.data
    } catch (err) {
      error.value = err.message
    } finally {
      loading.value = false
    }
  }

  const createEmployee = async (data) => {
    if (!canWriteHR.value) {
      console.warn('No permission to create employees')
      return
    }
    const newEmployee = {
      ...data,
      id: Date.now(),
      salary: parseFloat(data.salary),
      leaves: [],
      attendance: [],
      evaluations: [],
      skills: [],
      certificates: [],
      salaryHistory: [],
      fingerprint: null,
      fingerprintDate: null
    }
    employees.value.push(newEmployee)
    return newEmployee
  }

  const updateEmployee = async (id, data) => {
    if (!canWriteHR.value) {
      console.warn('No permission to update employees')
      return
    }
    const index = employees.value.findIndex(e => e.id === id)
    if (index > -1) {
      employees.value[index] = {
        ...data,
        id,
        salary: parseFloat(data.salary),
        leaves: employees.value[index].leaves || [],
        attendance: employees.value[index].attendance || [],
        evaluations: employees.value[index].evaluations || [],
        skills: employees.value[index].skills || [],
        certificates: employees.value[index].certificates || [],
        salaryHistory: employees.value[index].salaryHistory || [],
        fingerprint: employees.value[index].fingerprint || null,
        fingerprintDate: employees.value[index].fingerprintDate || null
      }
      return employees.value[index]
    }
    return null
  }

  const deleteEmployee = async (id) => {
    if (!canDeleteHR.value) {
      console.warn('No permission to delete employees')
      return
    }
    const index = employees.value.findIndex(e => e.id === id)
    if (index > -1) {
      employees.value.splice(index, 1)
      return true
    }
    return false
  }

  // Employee details operations
  const addLeave = (employeeId, leaveData) => {
    const employee = employees.value.find(e => e.id === employeeId)
    if (employee) {
      if (!employee.leaves) employee.leaves = []
      const start = new Date(leaveData.startDate)
      const end = new Date(leaveData.endDate)
      const days = Math.ceil((end - start) / (1000 * 60 * 60 * 24)) + 1
      employee.leaves.push({
        id: Date.now(),
        ...leaveData,
        days
      })
    }
  }

  const addAttendance = (employeeId, attendanceData) => {
    const employee = employees.value.find(e => e.id === employeeId)
    if (employee) {
      if (!employee.attendance) employee.attendance = []
      const checkIn = new Date(`${attendanceData.date}T${attendanceData.checkIn}`)
      const checkOut = attendanceData.checkOut
        ? new Date(`${attendanceData.date}T${attendanceData.checkOut}`)
        : null
      const hours = checkOut ? ((checkOut - checkIn) / (1000 * 60 * 60)).toFixed(2) : 0
      employee.attendance.push({
        id: Date.now(),
        ...attendanceData,
        hours: parseFloat(hours)
      })
    }
  }

  const addEvaluation = (employeeId, evaluationData) => {
    const employee = employees.value.find(e => e.id === employeeId)
    if (employee) {
      if (!employee.evaluations) employee.evaluations = []
      employee.evaluations.push({
        id: Date.now(),
        ...evaluationData
      })
    }
  }

  const addSkill = (employeeId, skillName) => {
    const employee = employees.value.find(e => e.id === employeeId)
    if (employee) {
      if (!employee.skills) employee.skills = []
      if (!employee.skills.includes(skillName)) {
        employee.skills.push(skillName)
      }
    }
  }

  const deleteSkill = (employeeId, skillName) => {
    const employee = employees.value.find(e => e.id === employeeId)
    if (employee && employee.skills) {
      const index = employee.skills.indexOf(skillName)
      if (index > -1) {
        employee.skills.splice(index, 1)
      }
    }
  }

  const addCertificate = (employeeId, certificateData) => {
    const employee = employees.value.find(e => e.id === employeeId)
    if (employee) {
      if (!employee.certificates) employee.certificates = []
      employee.certificates.push({
        id: Date.now(),
        ...certificateData
      })
    }
  }

  const deleteCertificate = (employeeId, certId) => {
    const employee = employees.value.find(e => e.id === employeeId)
    if (employee && employee.certificates) {
      const index = employee.certificates.findIndex(c => c.id === certId)
      if (index > -1) {
        employee.certificates.splice(index, 1)
      }
    }
  }

  const addSalaryRecord = (employeeId, salaryData) => {
    const employee = employees.value.find(e => e.id === employeeId)
    if (employee) {
      if (!employee.salaryHistory) employee.salaryHistory = []
      employee.salaryHistory.push({
        id: Date.now(),
        ...salaryData
      })
    }
  }

  // Fingerprint operations
  const registerFingerprint = async (employeeId) => {
    const employee = employees.value.find(e => e.id === employeeId)
    if (employee) {
      // Simulate fingerprint registration
      await new Promise(resolve => setTimeout(resolve, 2000))
      const fingerprintId = 'FP_' + Date.now() + '_' + Math.random().toString(36).substr(2, 9)
      const fingerprintDate = new Date().toISOString().split('T')[0]
      employee.fingerprint = fingerprintId
      employee.fingerprintDate = fingerprintDate
      return { fingerprintId, fingerprintDate }
    }
    return null
  }

  const deleteFingerprint = (employeeId) => {
    const employee = employees.value.find(e => e.id === employeeId)
    if (employee) {
      employee.fingerprint = null
      employee.fingerprintDate = null
      return true
    }
    return false
  }

  return {
    // State
    employees,
    loading,
    error,

    // Computed statistics
    totalEmployees,
    activeEmployees,
    totalDepartments,
    onLeave,
    totalSalaries,
    averageSalary,
    departmentOptions,
    positionOptions,

    // Permissions
    canReadHR,
    canWriteHR,
    canDeleteHR,

    // Methods
    fetchEmployees,
    createEmployee,
    updateEmployee,
    deleteEmployee,
    filterEmployees,
    addLeave,
    addAttendance,
    addEvaluation,
    addSkill,
    deleteSkill,
    addCertificate,
    deleteCertificate,
    addSalaryRecord,
    registerFingerprint,
    deleteFingerprint
  }
}
