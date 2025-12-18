<template>
  <v-app dir="rtl" class="arabic-app">
    <!-- الشريط الجانبي الأيمن -->
    <v-navigation-drawer
      v-model="drawer"
      location="right"
      permanent
      width="320"
      class="modern-sidebar rtl-sidebar"
    >
      <!-- الشعار والعنوان -->
      <div class="sidebar-header">
        <v-img
          class="sidebar-logo"
          height="56"
          width="56"
          src="@/assets/logo.png"
        />
        <h1 class="header-title">نظام إدارة المشاريع</h1>
      </div>

      <!-- قائمة التنقل الرئيسية -->
      <v-list class="pa-3" nav>
        <v-list-item
          v-for="item in mainMenuItems"
          :key="item.title"
          :to="item.to"
          class="menu-item mb-2"
          rounded="xl"
          :class="{ 'active-menu-item': item.active }"
          :ripple="false"
        >
          <template v-slot:prepend>
            <div class="icon-container ms-3">
              <v-icon :color="item.active ? undefined : 'grey-darken-2'" size="24">
                {{ item.icon }}
              </v-icon>
            </div>
          </template>
          <v-list-item-title class="text-body-1 font-weight-medium text-right" :class="item.active ? 'font-weight-bold' : 'text-grey-darken-3'">
            {{ item.title }}
          </v-list-item-title>
          <template v-slot:append v-if="item.badge">
            <v-chip
              :color="item.badgeColor"
              size="small"
              variant="flat"
              text-color="white"
            >
              {{ item.badge }}
            </v-chip>
          </template>
        </v-list-item>
      </v-list>

      <!-- قسم المميزات -->
      <!-- Features Section Removed -->

      <!-- Footer -->
      <template v-slot:append>
        <div class="pa-4 text-center">
                 <p class="text-caption text-secondary text-opacity-60">
                
                 </p>
                 <p class="text-caption text-secondary text-opacity-60">
                 </p>
        </div>
      </template>
    </v-navigation-drawer>

    <!-- المحتوى الرئيسي -->
    <v-main class="app-main-content">
      <router-view />
    </v-main>
  </v-app>
</template>

<script setup>
import { ref, computed, watch } from 'vue'
import { useRoute } from 'vue-router'

const route = useRoute()
const drawer = ref(true)

// قائمة التنقل الرئيسية
const mainMenuItems = ref([
  { title: 'الرئيسية', icon: 'mdi-view-dashboard', to: '/', active: false },
  { title: 'المشاريع', icon: 'mdi-folder-multiple', to: '/project-management', active: false },
  { title: 'إدارة المهام', icon: 'mdi-clipboard-list', to: '/task-management', active: false, badge: 'جديد', badgeColor: 'success' },
  { title: 'المهندسين', icon: 'mdi-account-hard-hat', to: '/engineers', active: false },
  { title: 'التصنيفات', icon: 'mdi-tag-multiple', to: '/categories', active: false },
  { title: 'المصروفات الإدارية', icon: 'mdi-chart-line', to: '/expenses', active: false },
  { title: 'أنواع المصروفات', icon: 'mdi-format-list-bulleted-type', to: '/expense-types', active: false },
  { title: 'الإيرادات', icon: 'mdi-trending-up', to: '/income', active: false },
  { title: 'المديونون', icon: 'mdi-credit-card', to: '/debtors', active: false, badge: '3', badgeColor: 'error' },
  { title: 'المستخدمين', icon: 'mdi-account-multiple', to: '/users', active: false },
  { title: 'الموارد البشرية', icon: 'mdi-account-group', to: '/human-resources', active: false },
  { title: 'بصمة الموظفين', icon: 'mdi-fingerprint', to: '/human-resources#fingerprint', active: false, badge: 'جديد', badgeColor: 'primary' }
])

const features = ref([
  {
    icon: 'mdi-chart-line',
    title: 'تقارير المشاريع',
    color: 'primary'
  },
  {
    icon: 'mdi-calculator',
    title: 'إدارة المالية',
    color: 'warning'
  },
  {
    icon: 'mdi-shield-check',
    title: 'أمان عالي',
    color: 'error'
  }
])

// تحديث العنصر النشط حسب الصفحة الحالية
const updateActiveMenuItem = () => {
  mainMenuItems.value.forEach(item => {
    item.active = item.to === route.path
  })
}

// مراقبة تغيير الصفحة
watch(() => route.path, updateActiveMenuItem, { immediate: true })
</script>

<style>
/* دعم RTL للنظام العربي */
.arabic-app {
  direction: rtl;
  text-align: right;
  font-family: var(--font-primary);
  background-color: var(--background-secondary) !important; /* Force global background */
}

/* ========================================
   Main Content Area - Crystal Indigo
   ======================================== */
/* ========================================
   Main Content Area - Crystal Indigo
   ======================================== */
.app-main-content {
  background-color: var(--background-tertiary) !important; /* Slate 100 - Better Contrast */
  min-height: 100vh;
}

/* ========================================
   السايد بار العصري - Crystal Indigo Theme
   ======================================== */
/* ========================================
   السايد بار العصري - Dark Crystal Theme
   ======================================== */
/* ========================================
   السايد بار العصري - Crystal Indigo White
   ======================================== */
/* ========================================
   السايد بار العصري - Midnight Crystal (Hybrid)
   ======================================== */
/* ========================================
   السايد بار العصري - Modern Floating Crystal
   ======================================== */
/* ========================================
   السايد بار العصري - Midnight Integrated
   ======================================== */
.modern-sidebar,
.v-navigation-drawer.modern-sidebar,
.v-application .modern-sidebar,
.v-application .v-navigation-drawer.modern-sidebar,
body .modern-sidebar,
body .v-navigation-drawer.modern-sidebar {
  background: linear-gradient(180deg, var(--color-slate-900) 0%, var(--color-indigo-950) 100%) !important; /* Unified Gradient */
  border: none !important;
  border-right: none !important;
  border-left: none !important;
  margin-right: 0 !important;
  margin: 0 !important;
  padding: 0 !important;
  right: 0 !important;
  left: auto !important;
  border-top-right-radius: 0 !important; /* Force Square Edge */
  border-bottom-right-radius: 0 !important; /* Force Square Edge */
  box-shadow: 4px 0 30px rgba(0, 0, 0, 0.2) !important;
  color: #ffffff !important;
  z-index: 1004 !important;
}

/* Nuclear Option: Kill the explicit Vuetify border element */
.modern-sidebar .v-navigation-drawer__border,
.v-navigation-drawer.modern-sidebar .v-navigation-drawer__border {
  display: none !important;
  width: 0 !important;
  opacity: 0 !important;
}

/* Make inner list transparent */
.modern-sidebar .v-list,
.v-navigation-drawer.modern-sidebar .v-list {
  background: transparent !important;
}

/* Alignment Fix */
.modern-sidebar .v-list-item,
.v-navigation-drawer.modern-sidebar .v-list-item {
  display: flex !important;
  align-items: center !important;
  min-height: 48px; /* Ensure consistent height */
  padding-inline-start: 16px !important;
  padding-inline-end: 16px !important;
}

/* Ensure prepend container alignment and transparency */
.modern-sidebar .v-list-item__prepend,
.v-navigation-drawer.modern-sidebar .v-list-item__prepend {
  display: flex !important;
  align-items: center !important;
  justify-content: center !important;
  margin-inline-end: 16px !important; /* Proper RTL margin */
  background-color: transparent !important;
  background: transparent !important;
}

/* Fix Title Alignment */
.modern-sidebar .v-list-item-title,
.modern-sidebar .v-list-item__content {
  text-align: right !important;
  flex: 1;
}

/* Icon Container for center alignment */
.icon-container {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 24px;
}

/* Menu item text styling */
/* Menu item text styling */
/* Menu item text styling */
/* Menu item text styling */
/* Menu item text styling */
/* Menu item text styling */
/* Menu item text styling */
/* Menu item text styling */
.modern-sidebar .v-list-item-title,
.modern-sidebar .v-list-item__content {
  color: var(--color-slate-300) !important; /* Light Slate for Dark Bg */
  -webkit-text-fill-color: var(--color-slate-300) !important;
  font-family: var(--font-secondary) !important;
  font-size: 0.9rem !important;
  font-weight: 500 !important;
  letter-spacing: 0px;
  transition: all 0.2s ease;
}

/* Header text */
/* Header text */
/* Header text */
/* Header text */
/* Header text */
.modern-sidebar h1,
.modern-sidebar h2,
.modern-sidebar h3 {
  color: #ffffff !important;
  -webkit-text-fill-color: #ffffff !important;
  font-family: var(--font-secondary) !important;
}

/* Subtitles and smaller text */
/* Subtitles and smaller text */
/* Subtitles and smaller text */
/* Subtitles and smaller text */
.modern-sidebar .text-caption,
.modern-sidebar .text-body-2,
.modern-sidebar .v-list-item-subtitle {
  color: var(--color-slate-400) !important;
  font-weight: 500 !important;
}

/* Icons - Elegant muted style */
/* Icons - Elegant muted style */
/* Icons - Elegant muted style */
/* Icons - Elegant muted style */
/* Icons - Elegant muted style */
/* Icons - Elegant muted style */
/* Icons - Elegant muted style */
/* Icons - Elegant muted style */
.modern-sidebar .v-icon,
.v-navigation-drawer.modern-sidebar .v-icon {
  color: var(--color-slate-400) !important;
  opacity: 1 !important;
  transition: all 0.2s cubic-bezier(0.4, 0, 0.2, 1);
}

/* Active Menu Item - Dark Theme Pop */
/* Active Menu Item - Pop-out Card Style */
/* Active Menu Item - Pop-out Card Style */
/* Active Menu Item - Modern Gradient Pill */
/* Active Menu Item - High Contrast White Pop */
.modern-sidebar .active-menu-item .v-list-item-title {
  color: var(--color-slate-900) !important; /* Unified Dark Slate */
  font-weight: 700 !important;
  -webkit-text-fill-color: var(--color-slate-900) !important;
  letter-spacing: 0;
}

.modern-sidebar .active-menu-item .v-icon {
  color: var(--color-slate-900) !important; /* Match Sidebar Background (Cutout Effect) */
  opacity: 1 !important;
}

/* White Pill on Dark Background */
/* White Pill on Dark Background - Focus Killer */
.modern-sidebar .active-menu-item,
.modern-sidebar .v-list-item--active,
.modern-sidebar .v-list-item:focus,
.modern-sidebar .active-menu-item:focus,
.modern-sidebar .v-list-item--active:focus,
.modern-sidebar .v-list-item:focus-visible {
  background: #ffffff !important; /* Pure White */
  /* Re-assert shadow - but allow no focus ring */
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1) !important;
  border: 0px solid transparent !important;
  border-width: 0 !important;
  outline: none !important; /* KILL FOCUS RING */
  --v-border-opacity: 0 !important;
  --v-theme-primary: #ffffff !important; /* Kill Primary color bleed */
  position: relative;
  border-radius: 8px !important;
  margin-bottom: 4px;
}

.modern-sidebar .active-menu-item::before,
.modern-sidebar .active-menu-item::after {
  display: none !important; /* Kill Pseudo borders */
  border: none !important;
  content: none !important; 
}

/* Remove side indicator (the whole pill is the indicator) */
.modern-sidebar .active-menu-item::before {
  display: none !important;
}

/* Hover State - smooth darker */
.modern-sidebar .menu-item:hover .v-list-item-title {
  color: #ffffff !important;
  -webkit-text-fill-color: #ffffff !important;
  transform: translateX(-4px);
}

/* Explicitly Prevent Hover changing the Active Item Text */
.modern-sidebar .active-menu-item:hover .v-list-item-title {
  color: var(--color-indigo-900) !important; /* Keep Text Dark */
  -webkit-text-fill-color: var(--color-indigo-900) !important;
  transform: none !important; /* Stop sliding */
}

.modern-sidebar .menu-item:hover .v-icon {
  color: #ffffff !important;
  opacity: 1 !important;
  transform: scale(1.05);
}

/* Explicitly Prevent Hover changing the Active Item Icon */
.modern-sidebar .active-menu-item:hover .v-icon {
  color: var(--color-slate-900) !important; /* Keep Icon Dark */
  transform: none !important;
}

.modern-sidebar .menu-item:hover {
  background: rgba(255, 255, 255, 0.1) !important; /* White Glass */
  border-radius: 8px !important;
}

/* Active Hover Override - STAY WHITE & SOLID */
.modern-sidebar .active-menu-item:hover {
  background: #ffffff !important;
  border: none !important;
  outline: none !important;
  opacity: 1 !important;
}


/* Features section title */
.modern-sidebar .px-4 h3 {
  color: var(--text-tertiary) !important;
  -webkit-text-fill-color: var(--text-tertiary) !important;
  font-size: 0.85rem !important;
  font-weight: 600 !important;
  text-transform: uppercase;
  letter-spacing: 0.05em;
}

.rtl-sidebar {
  direction: rtl;
  text-align: right;
}

.sidebar-header {
  display: flex;
  align-items: center;
  gap: 12px;
  padding: 20px;
  border-bottom: 1px solid rgba(255, 255, 255, 0.1) !important; /* Subtle Dark Theme Border */
  background: transparent !important; /* Blend with Gradient */
}

.sidebar-logo {
  flex-shrink: 0;
  border-radius: 10px;
}

.header-title {
  font-size: var(--font-size-lg) !important;
  font-weight: var(--font-weight-bold) !important;
  color: #ffffff !important;
  -webkit-text-fill-color: #ffffff !important;
  margin: 0;
  line-height: var(--line-height-snug);
  font-family: var(--font-secondary) !important;
}

/* Badge/Chip styling in sidebar */
/* 1. Force white text on chip content */
.modern-sidebar .v-chip .v-chip__content,
.v-navigation-drawer.modern-sidebar .v-chip .v-chip__content {
  color: var(--text-white) !important;
  -webkit-text-fill-color: var(--text-white) !important;
}

/* 2. Force white icons in chips */
.modern-sidebar .v-chip .v-icon,
.v-navigation-drawer.modern-sidebar .v-chip .v-icon {
  color: var(--text-white) !important;
  -webkit-text-fill-color: var(--text-white) !important;
}

/* 3. Explicitly restore background colors for sidebar chips */
.modern-sidebar .v-chip.bg-success,
.modern-sidebar .v-chip.text-success,
.v-navigation-drawer.modern-sidebar .v-chip.bg-success,
.v-navigation-drawer.modern-sidebar .v-chip.text-success {
  background-color: var(--color-success) !important;
  border-color: var(--color-success) !important;
}

.modern-sidebar .v-chip.bg-error,
.modern-sidebar .v-chip.text-error,
.v-navigation-drawer.modern-sidebar .v-chip.bg-error,
.v-navigation-drawer.modern-sidebar .v-chip.text-error {
  background-color: var(--color-error) !important;
  border-color: var(--color-error) !important;
}

.modern-sidebar .v-chip.bg-primary,
.modern-sidebar .v-chip.text-primary,
.v-navigation-drawer.modern-sidebar .v-chip.bg-primary,
.v-navigation-drawer.modern-sidebar .v-chip.text-primary {
  background-color: var(--color-primary) !important;
  border-color: var(--color-primary) !important;
}

/* عناصر القائمة - Menu Item Container */
.menu-item {
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
  border-radius: 12px !important;
  margin: 0 8px 6px 8px !important; /* Floating margins */
  border: none;
  min-height: 48px !important;
}

/* Features Grid */
.feature-grid {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 10px;
}

.feature-item {
  background: var(--background-secondary);
  border: 1px solid var(--border-light);
  border-radius: 10px;
  padding: 14px 10px;
  text-align: center;
  transition: all 0.25s cubic-bezier(0.4, 0, 0.2, 1);
}

.feature-item:hover {
  background: var(--background-tertiary);
  border-color: var(--color-indigo-200);
  transform: translateY(-2px);
  box-shadow: var(--shadow-light);
}

.feature-item .v-icon {
  color: var(--color-primary) !important;
}

.feature-item .text-caption {
  color: var(--text-tertiary) !important;
  font-weight: 600 !important;
  font-size: 0.75rem !important;
}

/* Feature Link */
.feature-link {
  text-decoration: none;
  display: flex;
  flex-direction: column;
  align-items: center;
  background: linear-gradient(135deg, var(--background-surface) 0%, var(--background-secondary) 100%) !important;
  border: 1px dashed var(--color-indigo-200) !important;
}

.feature-link:hover {
  background: linear-gradient(135deg, var(--background-indigo-light) 0%, var(--background-indigo-medium) 100%) !important;
  border-color: var(--color-indigo-300) !important;
}

.feature-link .v-icon {
  color: var(--color-success) !important;
}

/* Common Fixes */
.v-list-item {
  text-align: right;
}

.v-card-title {
  text-align: right;
  direction: rtl;
}

/* Fix Vuetify v-card-title flex direction for RTL */
.v-card-title.v-card-item__content,
.v-card > .v-card-title {
  display: flex;
  flex-direction: row;
  justify-content: flex-start;
  direction: rtl;
}

.v-card-text {
  text-align: right;
}

/* إصلاح اتجاه الأيقونات الاتجاهية فقط (السهام والشيفرونات) */
/* Note: Don't flip all icons - only directional ones need flipping for RTL */
.v-icon.mdi-chevron-left,
.v-icon.mdi-chevron-right,
.v-icon.mdi-arrow-left,
.v-icon.mdi-arrow-right,
.v-icon.mdi-menu-left,
.v-icon.mdi-menu-right {
  transform: scaleX(-1);
}

/* إصلاح اتجاه الأزرار */
.v-btn {
  direction: rtl;
}

/* إصلاح اتجاه النصوص */
h1, h2, h3, h4, h5, h6, p, span, div {
  text-align: right;
}

/* إصلاح اتجاه الجداول */
.v-data-table {
  direction: rtl;
}

.v-data-table th,
.v-data-table td {
  text-align: right;
}

/* إصلاح اتجاه النماذج */
.v-text-field input {
  text-align: right;
}

.v-select .v-field__input {
  text-align: right;
}

/* إصلاح اتجاه القوائم */
.v-menu .v-list {
  direction: rtl;
}

/* إصلاح اتجاه الشريط الجانبي */
.v-navigation-drawer {
  direction: rtl;
}

/* إصلاح اتجاه المحتوى */
.v-main {
  direction: rtl;
}

/* إصلاح اتجاه الحاويات */
.v-container {
  direction: rtl;
}

/* إصلاح اتجاه الصفوف والأعمدة */
.v-row {
  direction: rtl;
}

.v-col {
  direction: rtl;
}

/* ========================================
   Main Content Area - Ensure proper colors
   ======================================== */

/* Ensure page content has proper colors */
.v-main .data-page,
.v-main .project-card,
.v-main .stat-card {
  color: inherit;
}

/* Allow page-specific text colors to work */
.v-main .data-page .page-content-section p,
.v-main .data-page .page-content-section span,
.v-main .data-page .page-content-section h1,
.v-main .data-page .page-content-section h2,
.v-main .data-page .page-content-section h3,
.v-main .data-page .page-content-section h4,
.v-main .data-page .page-content-section h5,
.v-main .data-page .page-content-section h6 {
  -webkit-text-fill-color: unset;
}

/* ========================================
   Toast Notifications RTL Support
   ======================================== */
.Toastify__toast-container {
  direction: rtl !important;
  z-index: 9999 !important;
}

.Toastify__toast {
  direction: rtl !important;
  font-family: var(--font-primary) !important;
}

.Toastify__toast-body {
  direction: rtl !important;
  text-align: right !important;
}

/* Position toast on left for RTL */
.Toastify__toast-container--top-right {
  right: auto !important;
  left: 1em !important;
}

.Toastify__toast-container--bottom-right {
  right: auto !important;
  left: 1em !important;
}
</style>
