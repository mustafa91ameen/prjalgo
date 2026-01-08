<template>
  <div class="team-container">
    <!-- Page Header Component -->
    <PageHeader
      title="إدارة أعضاء الفريق"
      subtitle="إدارة وتنظيم جميع أعضاء فريق العمل"
      badge="أعضاء الفريق"
      badgeType="success"
      class="team-header"
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
      <!-- Total Members Card -->
      <v-card class="stat-card" elevation="0">
        <div class="stat-card-content">
          <div class="stat-icon team">
            <i class="mdi mdi-account-group"></i>
          </div>
          <div class="stat-info">
            <div class="stat-label">إجمالي الأعضاء</div>
            <div class="stat-value">{{ totalMembers }}</div>
            <div class="stat-change positive">
              <i class="mdi mdi-trending-up"></i>
              <span>+6 هذا الشهر</span>
            </div>
          </div>
        </div>
      </v-card>

      <!-- Unique Users Card -->
      <v-card class="stat-card" elevation="0">
        <div class="stat-card-content">
          <div class="stat-icon unique">
            <i class="mdi mdi-account-star"></i>
          </div>
          <div class="stat-info">
            <div class="stat-label">المستخدمين الفريدين</div>
            <div class="stat-value">{{ uniqueUsers }}</div>
            <div class="stat-change positive">
              <i class="mdi mdi-trending-up"></i>
              <span>+4 هذا الشهر</span>
            </div>
          </div>
        </div>
      </v-card>

      <!-- Average Members per Project Card -->
      <v-card class="stat-card" elevation="0">
        <div class="stat-card-content">
          <div class="stat-icon average">
            <i class="mdi mdi-chart-line"></i>
          </div>
          <div class="stat-info">
            <div class="stat-label">متوسط الأعضاء/مشروع</div>
            <div class="stat-value">{{ averageMembersPerProject }}</div>
            <div class="stat-change positive">
              <i class="mdi mdi-trending-up"></i>
              <span>+0.5 هذا الشهر</span>
            </div>
          </div>
        </div>
      </v-card>
    </div>

    <!-- Team Members List Header -->
    <div class="team-list-header">
      <div class="list-header-content">
        <div class="list-header-info">
          <h2 class="list-header-title">
            <i class="mdi mdi-format-list-bulleted"></i>
            قائمة أعضاء الفريق
          </h2>
          <p class="list-header-subtitle">عرض جميع أعضاء الفريق والمشاريع المرتبطة</p>
        </div>
        <div class="list-header-actions">
          <button v-if="canCreate" class="list-action-btn primary" @click="openAddDialog">
            <i class="mdi mdi-plus"></i>
            إضافة عضو جديد
          </button>
        </div>
      </div>
    </div>

    <!-- Team Members Table -->
    <v-card class="mb-6">
      <v-data-table
        :headers="headers"
        :items="teamMembers"
        :loading="loading"
        class="elevation-1"
      >
        <template v-slot:item.projects="{ item }">
          <v-chip
            v-for="(project, index) in item.projects.slice(0, 2)"
            :key="index"
            size="small"
            class="me-1"
            color="primary"
          >
            {{ project }}
          </v-chip>
          <v-chip
            v-if="item.projects.length > 2"
            size="small"
            color="grey"
          >
            +{{ item.projects.length - 2 }}
          </v-chip>
        </template>
        <template v-slot:item.actions="{ item }">
          <v-btn
            v-if="canDelete"
            size="small"
            color="error"
            @click="deleteMember(item)"
          >
            <i class="mdi mdi-delete"></i>
          </v-btn>
        </template>
      </v-data-table>
    </v-card>

    <!-- Add Team Member Dialog -->
    <v-dialog v-model="showAddDialog" max-width="700" persistent>
      <v-card class="member-dialog">
        <v-card-title class="dialog-header">
          <i class="mdi mdi-account-plus"></i>
          إضافة عضو جديد للفريق
        </v-card-title>
        
        <v-card-text class="dialog-content">
          <v-form ref="memberForm" v-model="formValid">
            <v-row>
              <!-- المشروع -->
              <v-col cols="12">
                <v-select
                  v-model="newMember.project"
                  :items="availableProjects"
                  label="المشروع"
                  variant="outlined"
                  density="comfortable"
                  :rules="[v => !!v || 'المشروع مطلوب']"
                  required
                  autofocus
                ></v-select>
              </v-col>

              <!-- المستخدم -->
              <v-col cols="12">
                <v-select
                  v-model="newMember.user"
                  :items="availableUsers"
                  item-title="name"
                  item-value="id"
                  label="المستخدم"
                  variant="outlined"
                  density="comfortable"
                  :rules="[v => !!v || 'المستخدم مطلوب']"
                  required
                >
                  <template v-slot:item="{ props, item }">
                    <v-list-item v-bind="props">
                      <template v-slot:prepend>
                        <v-avatar color="primary" size="32">
                          <span class="text-white">{{ item.raw.name.charAt(0) }}</span>
                        </v-avatar>
                      </template>
                      <v-list-item-title>{{ item.raw.name }}</v-list-item-title>
                      <v-list-item-subtitle>{{ item.raw.role }}</v-list-item-subtitle>
                    </v-list-item>
                  </template>
                </v-select>
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
          <button class="dialog-btn save" @click="saveMember" :disabled="!formValid">
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
import { listTeamMembers, createTeamMember, deleteTeamMember as apiDeleteTeamMember, getTeamMemberStats } from '@/api/teamMembers'
import { getProjectsDropdown } from '@/api/projects'
import { getUsersDropdown } from '@/api/users'
import { usePermissions } from '@/composables/usePermissions'
import { useToast } from '@/composables/useToast'
import { DEFAULT_PAGE, DEFAULT_LIMIT } from '@/constants/pagination'

const { success, error: showError } = useToast()
const { canCreate, canUpdate, canDelete } = usePermissions('/teamMembers')

const loading = ref(false)
const showAddDialog = ref(false)
const formValid = ref(false)
const memberForm = ref(null)

// Pagination
const page = ref(DEFAULT_PAGE)
const limit = ref(DEFAULT_LIMIT)
const total = ref(0)

// Stats
const teamStats = ref({
  totalMembers: 0,
  uniqueUsers: 0,
  avgMembersPerProject: 0
})

// New member data
const newMember = ref({
  project: '',
  user: null
})

const availableProjects = ref([])
const availableUsers = ref([])
const teamMembers = ref([])

// Fetch team members from API
const fetchTeamMembers = async () => {
  loading.value = true
  try {
    const response = await listTeamMembers({ page: page.value, limit: limit.value })
    if (response.success) {
      // Group by user with their projects
      const memberMap = new Map()
      const items = response.data.items || []
      items.forEach(item => {
        const userId = item.user_id
        if (!memberMap.has(userId)) {
          memberMap.set(userId, {
            id: item.id,
            userId: userId,
            name: item.user_name || 'مستخدم',
            role: item.user_role || '',
            email: item.user_email || '',
            phone: item.user_phone || '',
            projects: []
          })
        }
        memberMap.get(userId).projects.push(item.project_name)
      })
      teamMembers.value = Array.from(memberMap.values())
      total.value = response.data.total || 0
    }
  } catch (error) {
    console.error('Error fetching team members:', error)
    showError('حدث خطأ في جلب أعضاء الفريق')
  } finally {
    loading.value = false
  }
}

// Fetch stats
const fetchStats = async () => {
  try {
    const response = await getTeamMemberStats()
    if (response.success) {
      teamStats.value = {
        totalMembers: response.data.total_members || 0,
        uniqueUsers: response.data.unique_users || 0,
        avgMembersPerProject: response.data.avg_members_per_project || 0
      }
    }
  } catch (error) {
    console.error('Error fetching stats:', error)
  }
}

// Fetch dropdowns
const fetchDropdowns = async () => {
  try {
    const [projectsRes, usersRes] = await Promise.all([
      getProjectsDropdown(),
      getUsersDropdown()
    ])
    if (projectsRes.success) {
      availableProjects.value = (projectsRes.data || []).map(p => ({
        title: p.name,
        value: p.id
      }))
    }
    if (usersRes.success) {
      availableUsers.value = (usersRes.data || []).map(u => ({
        id: u.id,
        name: u.name,
        role: u.role_name || ''
      }))
    }
  } catch (error) {
    console.error('Error fetching dropdowns:', error)
  }
}

const headers = [
  { title: 'الاسم', key: 'name', align: 'start' },
  { title: 'الدور', key: 'role', align: 'center' },
  { title: 'البريد الإلكتروني', key: 'email', align: 'center' },
  { title: 'الهاتف', key: 'phone', align: 'center' },
  { title: 'المشاريع', key: 'projects', align: 'center' },
  { title: 'الإجراءات', key: 'actions', align: 'center', sortable: false }
]

// Computed properties
const totalMembers = computed(() => teamStats.value.totalMembers || total.value)

const uniqueUsers = computed(() => teamStats.value.uniqueUsers || teamMembers.value.length)

const averageMembersPerProject = computed(() => {
  return teamStats.value.avgMembersPerProject ? teamStats.value.avgMembersPerProject.toFixed(1) : '0'
})

// Methods
const openAddDialog = () => {
  showAddDialog.value = true
}

const closeDialog = () => {
  showAddDialog.value = false
  resetForm()
}

const resetForm = () => {
  newMember.value = {
    project: '',
    user: null
  }
  if (memberForm.value) {
    memberForm.value.reset()
  }
}

const saveMember = async () => {
  if (!formValid.value) return

  loading.value = true
  try {
    const memberData = {
      projectId: newMember.value.project,
      userId: newMember.value.user
    }

    const response = await createTeamMember(memberData)
    if (response.success) {
      success('تم إضافة العضو بنجاح')
      closeDialog()
      fetchTeamMembers()
      fetchStats()
    } else {
      showError(response.message || 'حدث خطأ')
    }
  } catch (error) {
    console.error('Error saving member:', error)
    showError('حدث خطأ في حفظ العضو')
  } finally {
    loading.value = false
  }
}

// Delete member
const deleteMember = async (member) => {
  if (!confirm('هل أنت متأكد من حذف هذا العضو؟')) return

  loading.value = true
  try {
    const response = await apiDeleteTeamMember(member.id)
    if (response.success) {
      success('تم حذف العضو بنجاح')
      fetchTeamMembers()
      fetchStats()
    } else {
      showError(response.message || 'حدث خطأ في الحذف')
    }
  } catch (error) {
    console.error('Error deleting member:', error)
    showError('حدث خطأ في حذف العضو')
  } finally {
    loading.value = false
  }
}

// Initialize
onMounted(() => {
  fetchTeamMembers()
  fetchStats()
  fetchDropdowns()
})
</script>

<style scoped>
.team-container {
  padding: 32px;
  max-width: 1600px;
  margin: 0 auto;
}

/* Team Header Custom Color */
.team-header {
  background: linear-gradient(135deg, #018790 0%, #005461 100%) !important;
}

.team-header::before {
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

.stat-icon.team {
  background: linear-gradient(135deg, #10b981 0%, #059669 100%) !important;
  box-shadow: 0 6px 16px rgba(16, 185, 129, 0.4);
}

.stat-icon.unique {
  background: linear-gradient(135deg, #f59e0b 0%, #d97706 100%) !important;
  box-shadow: 0 6px 16px rgba(245, 158, 11, 0.4);
}

.stat-icon.average {
  background: linear-gradient(135deg, #06b6d4 0%, #0891b2 100%) !important;
  box-shadow: 0 6px 16px rgba(6, 182, 212, 0.4);
}

/* Team List Header */
.team-list-header {
  background: linear-gradient(135deg, #0a3d42 0%, #052428 100%);
  border-radius: 16px;
  margin-bottom: 24px;
  padding: 16px 24px;
  border: 2px solid transparent;
  position: relative;
  overflow: hidden;
}

.team-list-header::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  border-radius: 16px;
  padding: 2px;
  background: linear-gradient(135deg, #10b981 0%, #059669 50%, #10b981 100%);
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
  color: #10b981;
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
  background: linear-gradient(135deg, #10b981 0%, #059669 100%);
  color: white;
  box-shadow: 0 4px 12px rgba(16, 185, 129, 0.3);
}

.list-action-btn.primary:hover {
  transform: translateY(-2px);
  box-shadow: 0 6px 16px rgba(16, 185, 129, 0.4);
}

.list-action-btn i {
  font-size: 16px;
}

/* Member Dialog */
.member-dialog {
  border-radius: 16px !important;
  background: linear-gradient(135deg, #0a3d42 0%, #052428 100%) !important;
  border: 2px solid transparent !important;
  position: relative;
  overflow: hidden;
  direction: rtl;
  text-align: right;
}

.member-dialog::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  border-radius: 16px;
  padding: 2px;
  background: linear-gradient(135deg, #10b981 0%, #059669 50%, #10b981 100%);
  -webkit-mask: linear-gradient(#fff 0 0) content-box, linear-gradient(#fff 0 0);
  mask: linear-gradient(#fff 0 0) content-box, linear-gradient(#fff 0 0);
  -webkit-mask-composite: xor;
  mask-composite: exclude;
  pointer-events: none;
}

.dialog-header {
  background: rgba(16, 185, 129, 0.1) !important;
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
  color: #10b981;
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
  color: #10b981 !important;
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
  background: linear-gradient(135deg, #10b981 0%, #059669 100%);
  color: white;
  box-shadow: 0 4px 12px rgba(16, 185, 129, 0.3);
}

.dialog-btn.save:hover:not(:disabled) {
  transform: translateY(-2px);
  box-shadow: 0 6px 16px rgba(16, 185, 129, 0.4);
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
  .team-container {
    padding: 24px;
  }

  .stats-grid {
    grid-template-columns: repeat(3, 1fr);
  }

  .list-header-content {
    flex-direction: column;
    align-items: flex-start;
    gap: 16px;
  }

  .list-action-btn.primary {
    width: 100%;
    justify-content: center;
  }
}

@media (max-width: 768px) {
  .team-container {
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

  .stat-value {
    font-size: 24px;
  }

  .team-list-header {
    padding: 14px 18px;
  }

  .list-header-title {
    font-size: 16px;
  }

  .list-action-btn {
    padding: 8px 14px;
    font-size: 12px;
  }

  /* Dialog responsive */
  :deep(.v-dialog) {
    margin: 16px !important;
  }

  .dialog-header {
    font-size: 18px !important;
    padding: 16px 20px !important;
  }

  .dialog-content {
    padding: 20px !important;
  }

  .dialog-actions {
    padding: 14px 20px !important;
  }

  .dialog-btn {
    padding: 8px 18px;
    font-size: 13px;
  }

  /* Table responsive */
  .v-data-table :deep(.v-table__wrapper) {
    overflow-x: auto;
  }

  .v-data-table :deep(th),
  .v-data-table :deep(td) {
    white-space: nowrap;
    font-size: 13px;
    padding: 10px 12px !important;
  }
}

@media (max-width: 480px) {
  .team-container {
    padding: 12px;
  }

  .stat-card {
    border-radius: 12px !important;
  }

  .stat-icon {
    width: 42px;
    height: 42px;
    font-size: 20px;
  }

  .stat-value {
    font-size: 20px;
  }

  .stat-label {
    font-size: 12px;
  }

  .stat-change {
    font-size: 10px;
    padding: 3px 8px;
  }

  .team-list-header {
    padding: 12px 14px;
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

  .list-action-btn i {
    font-size: 14px;
  }

  /* Dialog full width */
  :deep(.v-dialog) {
    margin: 8px !important;
  }

  :deep(.v-dialog > .v-overlay__content) {
    max-width: calc(100% - 16px) !important;
  }

  .dialog-header {
    font-size: 16px !important;
    padding: 14px 16px !important;
  }

  .dialog-header i {
    font-size: 20px;
  }

  .dialog-content {
    padding: 16px !important;
  }

  .dialog-actions {
    padding: 12px 16px !important;
    flex-wrap: wrap;
    gap: 8px;
  }

  .dialog-btn {
    padding: 8px 14px;
    font-size: 12px;
    flex: 1;
    min-width: 100px;
    justify-content: center;
  }

  /* Table horizontal scroll */
  .v-data-table :deep(.v-table__wrapper) {
    overflow-x: auto;
    -webkit-overflow-scrolling: touch;
  }

  .v-data-table :deep(th),
  .v-data-table :deep(td) {
    font-size: 12px;
    padding: 8px 10px !important;
  }

  /* Project chips */
  .v-data-table :deep(.v-chip) {
    font-size: 10px !important;
    height: 22px !important;
  }
}

@media (max-width: 360px) {
  .team-container {
    padding: 8px;
  }

  .stat-card-content {
    padding: 12px;
    gap: 12px;
  }

  .stat-icon {
    width: 38px;
    height: 38px;
    font-size: 18px;
  }

  .stat-value {
    font-size: 18px;
  }

  .team-list-header {
    padding: 10px 12px;
  }

  .list-header-title {
    font-size: 13px;
  }

  .dialog-header {
    font-size: 15px !important;
    padding: 12px 14px !important;
  }

  .dialog-content {
    padding: 12px !important;
  }

  .dialog-actions {
    padding: 10px 12px !important;
  }

  .dialog-btn {
    padding: 6px 12px;
    font-size: 11px;
  }
}
</style>
