<template>
  <v-app>
    <!-- Sidebar Navigation - Only show when not on login page -->
    <template v-if="!isLoginPage">
      <div class="sidebar-container" :class="{ 'mobile-open': mobileDrawer }">
        <!-- Logo/Header -->
        <div class="sidebar-header">
          <div class="sidebar-logo">
            <i class="mdi mdi-view-dashboard-variant sidebar-logo-icon"></i>
            <span class="sidebar-logo-text">نظام المشاريع</span>
          </div>
        </div>

        <!-- Navigation - Dynamic from user permissions -->
        <nav class="sidebar-nav">
          <!-- Dynamic menu items from auth store -->
          <div class="nav-section" v-for="(section, sectionIndex) in menuSections" :key="sectionIndex">
            <div class="nav-section-title">{{ section.title }}</div>

            <router-link
              v-for="item in section.items"
              :key="item.to"
              :to="item.to"
              class="nav-item"
              :class="{ 'active': isActive(item.to) }"
            >
              <i :class="['mdi', item.icon, 'nav-item-icon']"></i>
              <span class="nav-item-text">{{ item.title }}</span>
            </router-link>

            <div v-if="sectionIndex < menuSections.length - 1" class="nav-divider"></div>
          </div>

          <!-- Logout section -->
          <div class="nav-divider"></div>
          <div class="nav-section">
            <a
              href="#"
              class="nav-item"
              @click.prevent="handleLogout"
            >
              <i class="mdi mdi-logout nav-item-icon"></i>
              <span class="nav-item-text">تسجيل الخروج</span>
            </a>
          </div>
        </nav>

        <!-- User Profile Footer -->
        <div class="sidebar-footer">
          <div class="user-profile">
            <div class="user-avatar">
              <i class="mdi mdi-account"></i>
            </div>
            <div class="user-info">
              <div class="user-name">{{ currentUsername }}</div>
              <div class="user-role">مرحباً بك</div>
            </div>
            <i class="mdi mdi-dots-vertical" style="color: rgba(255, 255, 255, 0.7);"></i>
          </div>
        </div>
      </div>

      <!-- Mobile Menu Toggle Button -->
      <v-btn
        v-if="isMobile"
        icon
        class="mobile-menu-toggle"
        @click="mobileDrawer = !mobileDrawer"
        color="primary"
        size="large"
        elevation="4"
      >
        <i class="mdi mdi-menu"></i>
      </v-btn>
    </template>

    <!-- Main Content Area -->
    <v-main :class="{ 'main-content': !isLoginPage, 'login-main': isLoginPage }">
      <router-view />
    </v-main>
  </v-app>
</template>

<script setup>
import { ref, computed, onMounted, onUnmounted } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import { handleLogout as logoutUser, getCurrentUser } from '@/services/authService'
import { useAuthStore } from '@/stores/auth'

const router = useRouter()
const route = useRoute()
const authStore = useAuthStore()
const mobileDrawer = ref(false)
const windowWidth = ref(window.innerWidth)

const isMobile = computed(() => windowWidth.value <= 768)

// Check if we're on the login page
const isLoginPage = computed(() => {
  return route.path === '/login' || route.path === '/unauthorized'
})

// Get current username
const currentUsername = computed(() => {
  const user = getCurrentUser()
  return user?.username || localStorage.getItem('username') || 'المستخدم'
})

// Check if route is active
const isActive = (path) => {
  return route.path === path
}

// API-only routes that shouldn't appear in sidebar
const apiOnlyRoutes = [
  '/workdays',
  '/workdayLabor',
  '/workdayEquipment',
  '/workdayMaterials',
  '/rolePages',
  '/userRoles',
  '/workSubcategories'
]

// Organize menu items into sections
const menuSections = computed(() => {
  const pages = authStore.pages.filter(page => !apiOnlyRoutes.includes(page.route))

  // Define sections and their routes
  const sectionDefs = [
    {
      title: 'القائمة الرئيسية',
      routes: ['/', '/projects', '/tasks']
    },
    {
      title: 'الإدارة المالية',
      routes: ['/expenses', '/income', '/debtors']
    },
    {
      title: 'إدارة النظام',
      routes: ['/users', '/categories', '/roles', '/pages', '/team-members']
    }
  ]

  const sections = []

  for (const sectionDef of sectionDefs) {
    const items = pages
      .filter(page => sectionDef.routes.includes(page.route))
      .map(page => ({
        title: page.name,
        icon: page.icon || 'mdi-file-document',
        to: page.route
      }))

    if (items.length > 0) {
      sections.push({
        title: sectionDef.title,
        items
      })
    }
  }

  // Add any remaining pages to "Other" section
  const usedRoutes = sectionDefs.flatMap(s => s.routes)
  const otherItems = pages
    .filter(page => !usedRoutes.includes(page.route))
    .map(page => ({
      title: page.name,
      icon: page.icon || 'mdi-file-document',
      to: page.route
    }))

  if (otherItems.length > 0) {
    sections.push({
      title: 'أخرى',
      items: otherItems
    })
  }

  return sections
})

const handleResize = () => {
  windowWidth.value = window.innerWidth
}

const handleLogout = async () => {
  authStore.clearAuth()
  await logoutUser()
  router.push('/login')
}

onMounted(() => {
  window.addEventListener('resize', handleResize)
  // Load permissions from storage on mount
  authStore.loadFromStorage()
})

onUnmounted(() => {
  window.removeEventListener('resize', handleResize)
})
</script>

<style scoped>
/* Import sidebar styles */
@import '../styles/sidebar.css';

/* Mobile menu toggle button */
.mobile-menu-toggle {
  position: fixed;
  top: 16px;
  right: 16px;
  z-index: 1001;
  display: none;
}

@media (max-width: 768px) {
  .mobile-menu-toggle {
    display: flex;
  }
}

/* Login page main area */
.login-main {
  padding: 0 !important;
  margin: 0 !important;
}
</style>
