# مكون شريط عنوان الصفحات (PageHeader Component)

## نظرة عامة
مكون قابل لإعادة الاستخدام لعرض شريط عنوان احترافي في جميع صفحات النظام.

## المميزات
- ✨ تصميم عصري مع تدرجات لونية وحدود متوهجة
- 📍 دعم مسار التنقل (Breadcrumbs)
- 🏷️ شارات (Badges) بألوان متعددة
- 🎯 منطقة مخصصة للأزرار والإجراءات
- 📊 منطقة اختيارية للإحصائيات
- 📱 تصميم متجاوب بالكامل

## الاستخدام الأساسي

```vue
<template>
  <PageHeader
    title="عنوان الصفحة"
    subtitle="وصف مختصر للصفحة"
  />
</template>

<script setup>
import PageHeader from '@/components/PageHeader.vue'
</script>
```

## الخصائص (Props)

### title (مطلوب)
- **النوع:** `String`
- **الوصف:** عنوان الصفحة الرئيسي
- **مثال:** `"لوحة التحكم"`

### subtitle (اختياري)
- **النوع:** `String`
- **القيمة الافتراضية:** `''`
- **الوصف:** وصف مختصر أسفل العنوان
- **مثال:** `"مرحباً بك في نظام إدارة المشاريع"`

### badge (اختياري)
- **النوع:** `String`
- **القيمة الافتراضية:** `''`
- **الوصف:** نص الشارة بجانب العنوان
- **مثال:** `"جديد"`, `"مباشر"`, `"5 عناصر"`

### badgeType (اختياري)
- **النوع:** `String`
- **القيمة الافتراضية:** `'primary'`
- **القيم المتاحة:** `'primary'`, `'success'`, `'warning'`, `'danger'`, `'info'`
- **الوصف:** نوع ولون الشارة

### breadcrumbs (اختياري)
- **النوع:** `Array`
- **القيمة الافتراضية:** `[]`
- **الوصف:** مسار التنقل
- **الهيكل:**
  ```javascript
  [
    { label: 'الرئيسية', to: '/', icon: 'mdi mdi-home' },
    { label: 'المشاريع', to: '/projects' },
    { label: 'تفاصيل المشروع', to: '/projects/123' }
  ]
  ```

## الـ Slots

### actions
منطقة مخصصة لإضافة الأزرار والإجراءات في الجانب الأيسر من الشريط.

```vue
<PageHeader title="المشاريع">
  <template #actions>
    <button class="page-action-btn primary">
      <i class="mdi mdi-plus"></i>
      إضافة مشروع
    </button>
    <button class="page-action-btn secondary">
      <i class="mdi mdi-export"></i>
      تصدير
    </button>
  </template>
</PageHeader>
```

### stats
منطقة اختيارية لعرض إحصائيات سريعة أسفل العنوان.

```vue
<PageHeader title="المشاريع">
  <template #stats>
    <div class="quick-stats">
      <div class="stat-item">
        <span class="stat-label">المشاريع النشطة</span>
        <span class="stat-value">24</span>
      </div>
      <div class="stat-item">
        <span class="stat-label">المكتملة</span>
        <span class="stat-value">156</span>
      </div>
    </div>
  </template>
</PageHeader>
```

## أمثلة متقدمة

### مثال 1: صفحة بسيطة
```vue
<PageHeader
  title="الإعدادات"
  subtitle="إدارة إعدادات النظام"
/>
```

### مثال 2: صفحة مع شارة ومسار تنقل
```vue
<PageHeader
  title="المشاريع النشطة"
  subtitle="عرض جميع المشاريع قيد التنفيذ"
  badge="24 مشروع"
  badgeType="info"
  :breadcrumbs="[
    { label: 'الرئيسية', to: '/', icon: 'mdi mdi-home' },
    { label: 'المشاريع', to: '/projects' },
    { label: 'النشطة', to: '/projects/active' }
  ]"
/>
```

### مثال 3: صفحة كاملة مع أزرار
```vue
<PageHeader
  title="إدارة المستخدمين"
  subtitle="إضافة وتعديل صلاحيات المستخدمين"
  badge="مباشر"
  badgeType="success"
  :breadcrumbs="[
    { label: 'الرئيسية', to: '/', icon: 'mdi mdi-home' },
    { label: 'إدارة النظام', to: '/admin' },
    { label: 'المستخدمون', to: '/users' }
  ]"
>
  <template #actions>
    <button class="page-action-btn secondary">
      <i class="mdi mdi-filter-variant"></i>
      تصفية
    </button>
    <button class="page-action-btn secondary">
      <i class="mdi mdi-export"></i>
      تصدير
    </button>
    <button class="page-action-btn primary">
      <i class="mdi mdi-account-plus"></i>
      إضافة مستخدم
    </button>
    <button class="page-icon-btn">
      <i class="mdi mdi-dots-vertical"></i>
    </button>
  </template>
</PageHeader>
```

### مثال 4: صفحة مع إحصائيات
```vue
<PageHeader
  title="لوحة التحكم"
  subtitle="نظرة عامة على النظام"
  badge="محدث الآن"
  badgeType="success"
>
  <template #actions>
    <button class="page-action-btn secondary">
      <i class="mdi mdi-refresh"></i>
      تحديث
    </button>
  </template>
  
  <template #stats>
    <div style="display: grid; grid-template-columns: repeat(4, 1fr); gap: 20px;">
      <div style="text-align: center;">
        <div style="font-size: 24px; font-weight: 700; color: white;">245</div>
        <div style="font-size: 12px; color: rgba(255,255,255,0.6);">إجمالي المشاريع</div>
      </div>
      <div style="text-align: center;">
        <div style="font-size: 24px; font-weight: 700; color: white;">182</div>
        <div style="font-size: 12px; color: rgba(255,255,255,0.6);">المشاريع النشطة</div>
      </div>
      <div style="text-align: center;">
        <div style="font-size: 24px; font-weight: 700; color: white;">2.4M</div>
        <div style="font-size: 12px; color: rgba(255,255,255,0.6);">الإيرادات</div>
      </div>
      <div style="text-align: center;">
        <div style="font-size: 24px; font-weight: 700; color: white;">600K</div>
        <div style="font-size: 12px; color: rgba(255,255,255,0.6);">صافي الربح</div>
      </div>
    </div>
  </template>
</PageHeader>
```

## أنماط الأزرار المتاحة

### أزرار الإجراءات
```html
<!-- زر أساسي -->
<button class="page-action-btn primary">
  <i class="mdi mdi-plus"></i>
  إضافة
</button>

<!-- زر ثانوي -->
<button class="page-action-btn secondary">
  <i class="mdi mdi-export"></i>
  تصدير
</button>
```

### أزرار الأيقونات
```html
<button class="page-icon-btn">
  <i class="mdi mdi-filter-variant"></i>
</button>

<button class="page-icon-btn">
  <i class="mdi mdi-dots-vertical"></i>
</button>
```

## ألوان الشارات

| النوع | اللون | الاستخدام |
|------|-------|----------|
| `primary` | بنفسجي | افتراضي |
| `success` | أخضر | نجاح، مباشر، نشط |
| `warning` | برتقالي | تحذير، انتباه |
| `danger` | أحمر | خطر، خطأ |
| `info` | أزرق فاتح | معلومات، عدد العناصر |

## التصميم المتجاوب

المكون متجاوب بالكامل ويتكيف مع جميع أحجام الشاشات:
- **الشاشات الكبيرة:** العنوان والأزرار في صف واحد
- **الشاشات المتوسطة:** العنوان والأزرار في عمودين
- **الشاشات الصغيرة:** العنوان والأزرار في عمود واحد

## الملفات ذات الصلة

- **المكون:** `/src/components/PageHeader.vue`
- **الأنماط:** `/src/styles/page-header.css`
- **مثال الاستخدام:** `/src/pages/index.vue` (لوحة التحكم)

## ملاحظات التطوير

1. **الأيقونات:** يستخدم المكون Material Design Icons (mdi)
2. **التدرجات:** يتبع نفس نظام الألوان المستخدم في النظام
3. **الحدود المتوهجة:** تستخدم تقنية CSS mask-composite
4. **الأنيميشن:** جميع التفاعلات تحتوي على انتقالات سلسة

## التخصيص

يمكنك تخصيص الألوان والأنماط من خلال تعديل ملف `/src/styles/page-header.css`.

### مثال: تغيير لون الشارة الأساسية
```css
.page-header-badge.primary {
  background: linear-gradient(135deg, rgba(YOUR_COLOR_1, 0.2) 0%, rgba(YOUR_COLOR_2, 0.2) 100%);
  color: YOUR_TEXT_COLOR;
  border-color: rgba(YOUR_COLOR_1, 0.3);
}
```
