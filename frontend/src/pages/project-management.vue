<template>
  <v-container class="fill-height data-page" fluid>
    <div class="fullscreen-content">
      <!-- Header Section - Compact rectangle with animation -->
      <div class="page-header glass-effect gradient-animation">
        <span class="page-icon star-twinkle">🏗️</span>
        <h1 class="page-title text-glow fade-in">
          <span class="title-icon">🏗️</span>
          <span class="title-text">إدارة المشاريع الهندسية</span>
          <span class="title-decoration"></span>
        </h1>
        <p class="page-subtitle fade-in">نظام متكامل لإدارة وتتبع جميع المشاريع والمهام الهندسية</p>
        <div class="page-badge">
          <v-chip color="primary" variant="elevated" size="small">
            <v-icon start>mdi-star</v-icon>
            نظام إدارة متقدم
          </v-chip>
        </div>
      </div>


      <!-- Statistics Cards Component -->
      <ProjectStats
        :total-projects="totalProjects"
        :active-projects="activeProjects"
        :pending-projects="pendingProjects"
        :total-budget="totalBudget"
        :average-progress="averageProgress"
      />

      <!-- Projects Section -->
      <v-card class="data-table-card card-glow smooth-transition centered-table" elevation="2">
        <v-card-title class="d-flex align-center justify-space-between">
          <div class="d-flex align-center">
            <v-icon class="me-2" style="color: #4338ca;" size="28">mdi-folder-multiple</v-icon>
            <span class="text-h4 font-weight-black" style="color: #ffffff; font-family: 'Arial', 'Helvetica', sans-serif; text-shadow: 0 3px 6px rgba(0, 0, 0, 0.3), 0 1px 3px rgba(0, 0, 0, 0.2); letter-spacing: 0.5px;">قائمة المشاريع</span>
            <v-chip class="ms-3" color="primary" size="small" variant="elevated">{{ projects.length }}</v-chip>
          </div>
          <v-btn
            class="add-button btn-glow light-sweep smooth-transition"
            @click="openAddProjectDialog"
            elevation="2"
          >
            <span>إضافة مشروع جديد <v-icon class="icon-glow inline-icon">mdi-plus</v-icon></span>
          </v-btn>
        </v-card-title>

        <!-- Loading State -->
        <v-progress-linear v-if="loading" indeterminate color="primary" />

        <!-- Projects Grid using ProjectCard Component -->
        <v-row v-if="projects.length > 0" class="projects-grid-row">
          <v-col
            v-for="project in projects"
            :key="project.id"
            cols="12"
            sm="6"
            md="6"
            lg="4"
            xl="4"
          >
            <ProjectCard
              :project="project"
              @view-details="viewProjectDetails"
            />
          </v-col>
        </v-row>

        <!-- No Projects Message -->
        <v-card v-else-if="!loading" class="no-projects-card" elevation="2">
          <v-card-text class="text-center py-8">
            <v-icon size="4rem" color="grey-lighten-1">mdi-folder-open-outline</v-icon>
            <h3 class="text-h5 text-grey-lighten-1 mt-4">لا يوجد مشاريع</h3>
            <p class="text-body-1 text-grey-lighten-1">لم يتم العثور على أي مشاريع</p>
          </v-card-text>
        </v-card>
      </v-card>

      <!-- Add/Edit Project Dialog Component -->
      <ProjectForm
        v-model="dialog"
        :project="selectedProject"
        :is-editing="isEditing"
        @save="saveProject"
        @close="closeDialog"
      />

      <!-- Delete Confirmation Dialog Component -->
      <DeleteConfirmDialog
        v-model="deleteDialog"
        title="تأكيد الحذف"
        message="هل أنت متأكد من حذف المشروع؟"
        :item-name="selectedProject?.name"
        :loading="loading"
        @confirm="confirmDelete"
        @cancel="deleteDialog = false"
      />
    </div>
  </v-container>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { storeToRefs } from 'pinia'

// Import stores
import { useProjectsStore } from '@/stores/projects'

// Import components
import {
  ProjectStats,
  ProjectCard,
  ProjectForm,
  DeleteConfirmDialog
} from '@/components/projects'

// Page title
document.title = 'إدارة المشاريع الهندسية - نظام إدارة المشاريع'

const router = useRouter()

// Initialize store
const projectsStore = useProjectsStore()

// Get reactive state from store
const { projects, loading, error } = storeToRefs(projectsStore)

// Local UI state
const dialog = ref(false)
const deleteDialog = ref(false)
const isEditing = ref(false)
const selectedProject = ref(null)


// Computed properties from projects
const totalProjects = computed(() => projects.value.length)
const activeProjects = computed(() => projects.value.filter(p => p.status === 'active').length)
const pendingProjects = computed(() => projects.value.filter(p => p.status === 'pending').length)
const totalBudget = computed(() => projects.value.reduce((sum, p) => sum + (p.initialCost || 0), 0))
const averageProgress = computed(() => {
  if (projects.value.length === 0) return 0
  const total = projects.value.reduce((sum, p) => sum + (p.progress || 0), 0)
  return Math.round(total / projects.value.length)
})

// Fetch projects on mount
onMounted(async () => {
  await projectsStore.fetchProjects()
  // TODO: Fetch team members from API
})

// Project dialog functions
const openAddProjectDialog = () => {
  isEditing.value = false
  selectedProject.value = null
  dialog.value = true
}

const closeDialog = () => {
  dialog.value = false
  isEditing.value = false
  selectedProject.value = null
}

const saveProject = async (projectData) => {
  if (isEditing.value && selectedProject.value) {
    await projectsStore.updateProject(selectedProject.value.id, projectData)
  } else {
    await projectsStore.createProject(projectData)
  }
  closeDialog()
}

const viewProjectDetails = (project) => {
  router.push(`/project-details?id=${project.id}`)
}

const editProject = (project) => {
  isEditing.value = true
  selectedProject.value = project
  dialog.value = true
}

const deleteProject = (project) => {
  selectedProject.value = project
  deleteDialog.value = true
}

const confirmDelete = async () => {
  if (selectedProject.value) {
    await projectsStore.deleteProject(selectedProject.value.id)
  }
  deleteDialog.value = false
  selectedProject.value = null
}

// Team management functions (placeholder - will connect to API later)
const viewTeamMember = (member) => {
  // TODO: Implement view member details
}

const editTeamMember = (member) => {
  // TODO: Implement edit member
}

const deleteTeamMember = (member) => {
  // TODO: Implement delete member
}

const addTeamMember = (newMember) => {
  teamMembers.value.push(newMember)
}
</script>

<style scoped>
/* Import page styles - scoped to this component only */
@import './styles/project-management.css';
</style>
