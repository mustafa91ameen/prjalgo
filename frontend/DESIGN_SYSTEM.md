# Design System Documentation
# توثيق نظام التصميم الموحد

## Overview - نظرة عامة

This design system provides a centralized, consistent approach to styling across the entire application. All styles are defined in CSS variables and reusable components.

يوفر نظام التصميم هذا نهجاً موحداً ومركزياً لتنسيق التطبيق بالكامل. يتم تعريف جميع الأنماط في متغيرات CSS ومكونات قابلة لإعادة الاستخدام.

## File Structure - هيكل الملفات

```
src/
├── styles/
│   ├── design-system.css    # Core design system variables
│   ├── main.css              # Main stylesheet (imports all)
│   └── ...
├── components/
│   ├── PageHeader.vue        # Reusable page header component
│   ├── DataTable.vue         # Reusable data table component
│   ├── AppLayout.vue         # Main layout with sidebar
│   └── ...
```

## CSS Variables - متغيرات CSS

All design tokens are defined in `design-system.css` using CSS custom properties (variables).

### Colors - الألوان

```css
--ds-color-primary: #60a5fa;
--ds-color-primary-light: #93c5fd;
--ds-color-primary-dark: #3b82f6;
--ds-color-primary-darker: #2563eb;

--ds-color-secondary: #10b981;
--ds-text-primary: #1a1a1a;
--ds-text-white: #ffffff;
--ds-bg-primary: #ffffff;
--ds-bg-page: #f5f5f5;
```

### Typography - الخطوط

```css
--ds-font-family: 'Cairo', 'Tajawal', 'Arial', sans-serif;
--ds-font-size-xs: 0.6rem;
--ds-font-size-sm: 0.75rem;
--ds-font-size-base: 0.875rem;
--ds-font-size-md: 0.95rem;
--ds-font-size-lg: 1.1rem;
```

### Spacing - المسافات

```css
--ds-space-xs: 0.25rem;   /* 4px */
--ds-space-sm: 0.5rem;    /* 8px */
--ds-space-md: 0.75rem;   /* 12px */
--ds-space-lg: 1rem;      /* 16px */
--ds-space-xl: 1.5rem;    /* 24px */
```

### Table Styles - أنماط الجداول

```css
--ds-table-header-bg: var(--ds-gradient-primary);
--ds-table-header-text: var(--ds-text-white);
--ds-table-header-font-size: var(--ds-font-size-md);
--ds-table-header-padding: 3px 5px;
--ds-table-cell-font-size: var(--ds-font-size-sm);
```

## Components - المكونات

### PageHeader Component

A reusable page header component with consistent styling.

```vue
<template>
  <PageHeader
    title="عنوان الصفحة"
    subtitle="وصف الصفحة"
    icon="mdi-folder"
    :show-back-button="true"
  />
</template>

<script setup>
import PageHeader from '@/components/PageHeader.vue'
</script>
```

**Props:**
- `title` (required): Page title
- `subtitle` (optional): Page subtitle
- `icon` (optional): Icon name (MDI)
- `iconSize` (optional): Icon size (default: 18)
- `showBackButton` (optional): Show back button (default: false)
- `backRoute` (optional): Custom back route

### DataTable Component

A reusable data table component with search, header, and consistent styling.

```vue
<template>
  <DataTable
    title="قائمة البيانات"
    subtitle="وصف الجدول"
    icon="mdi-table"
    :count="items.length"
    :headers="headers"
    :items="items"
    :show-search="true"
  />
</template>

<script setup>
import DataTable from '@/components/DataTable.vue'

const headers = [
  { title: 'الاسم', key: 'name' },
  { title: 'البريد', key: 'email' }
]

const items = [
  { name: 'أحمد', email: 'ahmed@example.com' }
]
</script>
```

**Props:**
- `title` (optional): Table title
- `subtitle` (optional): Table subtitle
- `icon` (optional): Icon name
- `count` (optional): Item count to display
- `headers` (required): Table headers array
- `items` (required): Table items array
- `showSearch` (optional): Show search bar (default: true)
- `headerColor` (optional): Header color theme ('primary', 'secondary', 'green', 'blue')

### AppLayout Component

Main layout component with sidebar navigation.

```vue
<template>
  <AppLayout
    :menu-items="menuItems"
    :logo="logoPath"
    title="نظام إدارة المشاريع"
  >
    <router-view />
  </AppLayout>
</template>

<script setup>
import AppLayout from '@/components/AppLayout.vue'

const menuItems = [
  { title: 'الرئيسية', icon: 'mdi-view-dashboard', to: '/' },
  { title: 'المشاريع', icon: 'mdi-folder-multiple', to: '/projects' }
]
</script>
```

## Usage Examples - أمثلة الاستخدام

### Example 1: Simple Page with Header

```vue
<template>
  <div class="ds-page-container">
    <PageHeader
      title="المشاريع"
      subtitle="إدارة جميع المشاريع"
      icon="mdi-folder-multiple"
    />
    
    <div class="ds-content-container">
      <!-- Your content here -->
    </div>
  </div>
</template>

<style scoped>
.ds-page-container {
  min-height: 100vh;
  background: var(--ds-bg-page);
  direction: rtl;
}

.ds-content-container {
  padding: 0 var(--ds-space-2xl) var(--ds-space-2xl);
}
</style>
```

### Example 2: Page with Table

```vue
<template>
  <div class="ds-page-container">
    <PageHeader
      title="قائمة المهندسين"
      icon="mdi-account-group"
      :show-back-button="true"
    />
    
    <div class="ds-content-container">
      <DataTable
        title="المهندسين"
        icon="mdi-account-group"
        :count="engineers.length"
        :headers="headers"
        :items="engineers"
        header-color="green"
      />
    </div>
  </div>
</template>
```

### Example 3: Using CSS Variables Directly

```vue
<template>
  <div class="custom-card">
    <h2 class="custom-title">عنوان مخصص</h2>
    <p class="custom-text">نص مخصص</p>
  </div>
</template>

<style scoped>
.custom-card {
  background: var(--ds-bg-primary);
  padding: var(--ds-space-lg);
  border-radius: var(--ds-radius-lg);
  box-shadow: var(--ds-shadow-md);
}

.custom-title {
  color: var(--ds-text-primary);
  font-size: var(--ds-font-size-lg);
  font-weight: var(--ds-font-weight-bold);
  margin-bottom: var(--ds-space-md);
}

.custom-text {
  color: var(--ds-text-secondary);
  font-size: var(--ds-font-size-base);
  line-height: var(--ds-line-height-normal);
}
</style>
```

## Migration Guide - دليل الترحيل

### Step 1: Replace Inline Styles

**Before:**
```vue
<div style="padding: 16px; background: #ffffff; border-radius: 8px;">
```

**After:**
```vue
<div class="ds-card">
```

Or use variables:
```vue
<div style="padding: var(--ds-space-lg); background: var(--ds-bg-primary); border-radius: var(--ds-radius-lg);">
```

### Step 2: Replace Page Headers

**Before:**
```vue
<div class="engineers-header-card">
  <div class="header-content">
    <h1>المشاريع</h1>
  </div>
</div>
```

**After:**
```vue
<PageHeader
  title="المشاريع"
  icon="mdi-folder-multiple"
/>
```

### Step 3: Replace Tables

**Before:**
```vue
<v-card>
  <v-card-title>قائمة البيانات</v-card-title>
  <v-data-table :headers="headers" :items="items" />
</v-card>
```

**After:**
```vue
<DataTable
  title="قائمة البيانات"
  :headers="headers"
  :items="items"
/>
```

## Best Practices - أفضل الممارسات

1. **Always use CSS variables** instead of hardcoded values
2. **Use components** (PageHeader, DataTable) instead of custom implementations
3. **Follow the spacing scale** (xs, sm, md, lg, xl, 2xl)
4. **Use semantic color names** (primary, secondary, success, error)
5. **Maintain consistency** across all pages

## Utility Classes - الفئات المساعدة

```css
/* Backgrounds */
.ds-bg-primary
.ds-bg-secondary
.ds-bg-page

/* Text Colors */
.ds-text-primary
.ds-text-secondary
.ds-text-white

/* Spacing */
.ds-p-sm, .ds-p-md, .ds-p-lg
.ds-m-sm, .ds-m-md, .ds-m-lg

/* Border Radius */
.ds-rounded-sm
.ds-rounded-md
.ds-rounded-lg

/* Shadows */
.ds-shadow-sm
.ds-shadow-md
.ds-shadow-lg
```

## Support - الدعم

For questions or issues, refer to:
- `src/styles/design-system.css` - All CSS variables
- `src/components/PageHeader.vue` - Header component
- `src/components/DataTable.vue` - Table component
- `src/components/AppLayout.vue` - Layout component

