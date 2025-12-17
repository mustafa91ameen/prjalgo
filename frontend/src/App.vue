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
              <v-icon :color="item.active ? 'primary' : 'grey-darken-2'" size="24">
                {{ item.icon }}
              </v-icon>
            </div>
          </template>
          <v-list-item-title class="text-body-1 font-weight-medium text-right" :class="item.active ? 'text-primary font-weight-bold' : 'text-grey-darken-3'">
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
      <div class="px-4 mt-4">
        <h3 class="text-h6 font-weight-bold text-secondary mb-3">المميزات الرئيسية</h3>
        <div class="feature-grid">
          <div
            v-for="feature in features"
            :key="feature.title"
            class="feature-item"
          >
            <v-icon :color="feature.color" size="20" class="mb-1">{{ feature.icon }}</v-icon>
            <div class="text-caption text-secondary">{{ feature.title }}</div>
          </div>
          <!-- رابط إدارة الفريق -->
          <router-link to="/team-management" class="feature-item feature-link">
            <v-icon color="success" size="20" class="mb-1">mdi-account-group</v-icon>
            <div class="text-caption text-secondary">إدارة الفريق</div>
            <v-chip size="x-small" color="success" class="mt-1 text-white">جديد</v-chip>
          </router-link>
        </div>
      </div>

      <!-- Footer -->
      <template v-slot:append>
        <div class="pa-4 text-center">
                 <p class="text-caption text-white text-opacity-60">
                
                 </p>
                 <p class="text-caption text-white text-opacity-60">
                 </p>
        </div>
      </template>
    </v-navigation-drawer>

    <!-- المحتوى الرئيسي -->
    <v-main>
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
  font-family: 'Cairo', 'Tajawal', 'Noto Sans Arabic', 'Arial', sans-serif;
}

/* ========================================
   السايد بار العصري - Enhanced White Theme
   ======================================== */
.modern-sidebar,
.v-navigation-drawer.modern-sidebar,
.v-application .modern-sidebar,
.v-application .v-navigation-drawer.modern-sidebar,
body .modern-sidebar,
body .v-navigation-drawer.modern-sidebar {
  background: #ffffff !important;
  border-left: 1px solid #e2e8f0 !important;
  box-shadow: -6px 0 24px rgba(15, 23, 42, 0.06) !important;
  color: var(--text-primary) !important;
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
.modern-sidebar .v-list-item-title,
.modern-sidebar .v-list-item__content {
  color: #475569 !important; /* Slate 600 */
  -webkit-text-fill-color: #475569 !important;
  font-family: 'Tajawal', sans-serif !important;
  font-size: 0.95rem !important;
  font-weight: 500 !important;
  letter-spacing: 0.01em;
  transition: all 0.2s ease;
}

/* Header text */
.modern-sidebar h1,
.modern-sidebar h2,
.modern-sidebar h3 {
  color: #1e293b !important; /* Slate 800 */
  -webkit-text-fill-color: #1e293b !important;
  font-family: 'Tajawal', sans-serif !important;
}

/* Subtitles and smaller text */
.modern-sidebar .text-caption,
.modern-sidebar .text-body-2,
.modern-sidebar .v-list-item-subtitle {
  color: #94a3b8 !important; /* Slate 400 */
  font-weight: 500 !important;
}

/* Icons - Elegant muted style */
.modern-sidebar .v-icon,
.v-navigation-drawer.modern-sidebar .v-icon {
  color: #94a3b8 !important; /* Slate 400 */
  opacity: 1 !important;
  transition: all 0.25s cubic-bezier(0.4, 0, 0.2, 1);
}

/* Active Menu Item - Clean accent style */
.modern-sidebar .active-menu-item .v-list-item-title {
  color: #4338ca !important; /* Indigo 700 */
  -webkit-text-fill-color: #4338ca !important;
  font-weight: 600 !important;
}

.modern-sidebar .active-menu-item .v-icon {
  color: #4338ca !important; /* Indigo 700 */
}

.modern-sidebar .active-menu-item {
  background: linear-gradient(90deg, #eef2ff 0%, #f8fafc 100%) !important;
  border: none !important;
  border-right: 3px solid #4338ca !important;
  box-shadow: 0 2px 8px rgba(67, 56, 202, 0.08);
  position: relative;
}

/* Hover State - Smooth and subtle */
.modern-sidebar .menu-item:hover .v-list-item-title {
  color: #334155 !important; /* Slate 700 */
  -webkit-text-fill-color: #334155 !important;
}

.modern-sidebar .menu-item:hover .v-icon {
  color: #6366f1 !important; /* Indigo 500 */
  transform: scale(1.08);
}

.modern-sidebar .menu-item:hover {
  background: #f8fafc !important;
}


/* Features section title */
.modern-sidebar .px-4 h3 {
  color: #64748b !important;
  -webkit-text-fill-color: #64748b !important;
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
  border-bottom: 1px solid #f1f5f9;
  background: #ffffff !important;
}

.sidebar-logo {
  flex-shrink: 0;
  border-radius: 10px;
}

.header-title {
  font-size: 1.15rem !important;
  font-weight: 700 !important;
  color: #1e293b !important;
  -webkit-text-fill-color: #1e293b !important;
  margin: 0;
  line-height: 1.4;
  font-family: 'Tajawal', sans-serif !important;
}

/* Badge/Chip styling in sidebar - force white text (high specificity to override sidebar text rules) */
/* Badge/Chip styling in sidebar */
/* 1. Force white text on chip content */
.modern-sidebar .v-chip .v-chip__content,
.v-navigation-drawer.modern-sidebar .v-chip .v-chip__content {
  color: #ffffff !important;
  -webkit-text-fill-color: #ffffff !important;
}

/* 2. Force white icons in chips */
.modern-sidebar .v-chip .v-icon,
.v-navigation-drawer.modern-sidebar .v-chip .v-icon {
  color: #ffffff !important;
  -webkit-text-fill-color: #ffffff !important;
}

/* 3. Explicitly restore background colors for sidebar chips to prevent white-on-white */
/* Note: We target both bg- and text- classes as Vuetify can use either depending on version/variant */
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

/* عناصر القائمة */
.menu-item {
  transition: all 0.25s cubic-bezier(0.4, 0, 0.2, 1);
  border-radius: 12px !important;
  margin-bottom: 2px;
  border: none;
  min-height: 46px !important;
}

/* Features Grid */
.feature-grid {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 10px;
}

.feature-item {
  background: #f8fafc;
  border: 1px solid #e2e8f0;
  border-radius: 10px;
  padding: 14px 10px;
  text-align: center;
  transition: all 0.25s cubic-bezier(0.4, 0, 0.2, 1);
}

.feature-item:hover {
  background: #f1f5f9;
  border-color: #c7d2fe;
  transform: translateY(-2px);
  box-shadow: 0 4px 12px rgba(67, 56, 202, 0.08);
}

.feature-item .v-icon {
  color: #6366f1 !important;
}

.feature-item .text-caption {
  color: #64748b !important;
  font-weight: 600 !important;
  font-size: 0.75rem !important;
}

/* Feature Link */
.feature-link {
  text-decoration: none;
  display: flex;
  flex-direction: column;
  align-items: center;
  background: linear-gradient(135deg, #faf5ff 0%, #f8fafc 100%) !important;
  border: 1px dashed #c4b5fd !important;
}

.feature-link:hover {
  background: linear-gradient(135deg, #ede9fe 0%, #eef2ff 100%) !important;
  border-color: #a78bfa !important;
}

.feature-link .v-icon {
  color: #22c55e !important;
}

/* Common Fixes */
.v-list-item {
  text-align: right;
}

.v-card-title {
  text-align: right;
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
  font-family: 'Cairo', 'Tajawal', 'Noto Sans Arabic', 'Arial', sans-serif !important;
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
