# Design System Quick Reference
# مرجع سريع لنظام التصميم

## CSS Variables - متغيرات CSS

### Colors
```css
--ds-color-primary          /*rgb(128, 95, 103) */
--ds-color-primary-dark     /* #3b82f6 */
--ds-color-secondary         /*rgb(132, 132, 132) */
--ds-text-primary           /* #1a1a1a */
--ds-text-white             /* #ffffff */
--ds-bg-primary             /* #ffffff */
--ds-bg-page                /* #f5f5f5 */
```

### Typography
```css
--ds-font-size-xs           /* 0.6rem */
--ds-font-size-sm           /* 0.75rem */
--ds-font-size-base         /* 0.875rem */
--ds-font-size-md           /* 0.95rem */
--ds-font-size-lg           /* 1.1rem */
```

### Spacing
```css
--ds-space-xs               /* 0.25rem / 4px */
--ds-space-sm               /* 0.5rem / 8px */
--ds-space-md               /* 0.75rem / 12px */
--ds-space-lg               /* 1rem / 16px */
--ds-space-xl               /* 1.5rem / 24px */
--ds-space-2xl              /* 2rem / 32px */
```

### Border Radius
```css
--ds-radius-sm              /* 4px */
--ds-radius-md              /* 6px */
--ds-radius-lg              /* 8px */
--ds-radius-xl              /* 12px */
```

## Components - المكونات

### PageHeader
```vue
<PageHeader
  title="عنوان الصفحة"
  subtitle="وصف"
  icon="mdi-folder"
  :show-back-button="true"
/>
```

### DataTable
```vue
<DataTable
  title="قائمة البيانات"
  :count="items.length"
  :headers="headers"
  :items="items"
  header-color="primary"
/>
```

## Utility Classes - فئات مساعدة

```html
<!-- Backgrounds -->
<div class="ds-bg-primary">
<div class="ds-bg-secondary">
<div class="ds-bg-page">

<!-- Text Colors -->
<span class="ds-text-primary">
<span class="ds-text-white">

<!-- Spacing -->
<div class="ds-p-lg">
<div class="ds-m-md">

<!-- Border Radius -->
<div class="ds-rounded-lg">

<!-- Shadows -->
<div class="ds-shadow-md">
```

## Common Patterns - الأنماط الشائعة

### Page Container
```vue
<div class="ds-page-container">
  <PageHeader title="..." />
  <div class="ds-content-container">
    <!-- Content -->
  </div>
</div>
```

### Card
```vue
<div class="ds-card">
  <!-- Content -->
</div>
```

### Button
```vue
<v-btn class="ds-btn ds-btn-primary ds-btn-sm">
  Button Text
</v-btn>
```

### Search Container
```vue
<div class="ds-search-container">
  <div class="ds-search-box">
    <v-text-field class="ds-search-field" />
    <v-btn class="ds-search-btn">بحث</v-btn>
  </div>
</div>
```

### Table Title Bar
```vue
<div class="ds-table-title-bar">
  <div class="ds-table-title-bar-content">
    <v-icon class="ds-table-title-bar-icon">mdi-table</v-icon>
    <span class="ds-table-title-bar-text">عنوان الجدول</span>
    <v-chip class="ds-table-title-bar-chip">5</v-chip>
  </div>
</div>
```

