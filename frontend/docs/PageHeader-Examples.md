# أمثلة استخدام مكون PageHeader

## مثال 1: Dashboard (لوحة التحكم)
```vue
<PageHeader
  title="لوحة التحكم"
  subtitle="مرحباً بك في نظام إدارة المشاريع والبناء"
  badge="مباشر"
  badgeType="success"
  :breadcrumbs="[
    { label: 'الرئيسية', to: '/', icon: 'mdi mdi-home' },
    { label: 'لوحة التحكم', to: '/' }
  ]"
>
  <template #actions>
    <button class="page-action-btn secondary">
      <i class="mdi mdi-refresh"></i>
      تحديث البيانات
    </button>
    <button class="page-action-btn primary">
      <i class="mdi mdi-plus"></i>
      مشروع جديد
    </button>
    <button class="page-icon-btn">
      <i class="mdi mdi-filter-variant"></i>
    </button>
    <button class="page-icon-btn">
      <i class="mdi mdi-dots-vertical"></i>
    </button>
  </template>
</PageHeader>
```

## مثال 2: Projects (المشاريع)
```vue
<PageHeader
  title="إدارة المشاريع"
  subtitle="عرض وإدارة جميع المشاريع في النظام"
  badge="245 مشروع"
  badgeType="info"
  :breadcrumbs="[
    { label: 'الرئيسية', to: '/', icon: 'mdi mdi-home' },
    { label: 'المشاريع', to: '/projects' }
  ]"
>
  <template #actions>
    <button class="page-action-btn secondary">
      <i class="mdi mdi-export"></i>
      تصدير
    </button>
    <button class="page-action-btn primary">
      <i class="mdi mdi-plus"></i>
      مشروع جديد
    </button>
    <button class="page-icon-btn">
      <i class="mdi mdi-dots-vertical"></i>
    </button>
  </template>
</PageHeader>
```

## مثال 3: Users (المستخدمون)
```vue
<PageHeader
  title="إدارة المستخدمين"
  subtitle="إضافة وتعديل صلاحيات المستخدمين"
  badge="نشط"
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
    <button class="page-action-btn primary">
      <i class="mdi mdi-account-plus"></i>
      إضافة مستخدم
    </button>
  </template>
</PageHeader>
```

## مثال 4: Expenses (المصروفات)
```vue
<PageHeader
  title="المصروفات"
  subtitle="تتبع وإدارة جميع المصروفات"
  badge="تحذير"
  badgeType="warning"
  :breadcrumbs="[
    { label: 'الرئيسية', to: '/', icon: 'mdi mdi-home' },
    { label: 'الإدارة المالية', to: '/finance' },
    { label: 'المصروفات', to: '/expenses' }
  ]"
>
  <template #actions>
    <button class="page-action-btn secondary">
      <i class="mdi mdi-calendar"></i>
      تصفية بالتاريخ
    </button>
    <button class="page-action-btn secondary">
      <i class="mdi mdi-download"></i>
      تحميل التقرير
    </button>
    <button class="page-action-btn primary">
      <i class="mdi mdi-cash-minus"></i>
      تسجيل مصروف
    </button>
  </template>
</PageHeader>
```

## مثال 5: Settings (الإعدادات)
```vue
<PageHeader
  title="الإعدادات"
  subtitle="إدارة إعدادات النظام والتفضيلات"
  :breadcrumbs="[
    { label: 'الرئيسية', to: '/', icon: 'mdi mdi-home' },
    { label: 'الإعدادات', to: '/settings' }
  ]"
>
  <template #actions>
    <button class="page-action-btn secondary">
      <i class="mdi mdi-restore"></i>
      استعادة الافتراضي
    </button>
    <button class="page-action-btn primary">
      <i class="mdi mdi-content-save"></i>
      حفظ التغييرات
    </button>
  </template>
</PageHeader>
```

## مثال 6: Reports (التقارير)
```vue
<PageHeader
  title="التقارير والإحصائيات"
  subtitle="عرض التقارير المالية والإدارية"
  badge="محدث"
  badgeType="success"
  :breadcrumbs="[
    { label: 'الرئيسية', to: '/', icon: 'mdi mdi-home' },
    { label: 'التقارير', to: '/reports' }
  ]"
>
  <template #actions>
    <button class="page-action-btn secondary">
      <i class="mdi mdi-printer"></i>
      طباعة
    </button>
    <button class="page-action-btn secondary">
      <i class="mdi mdi-file-pdf"></i>
      تصدير PDF
    </button>
    <button class="page-action-btn primary">
      <i class="mdi mdi-chart-box"></i>
      تقرير جديد
    </button>
  </template>
</PageHeader>
```

## مثال 7: مع إحصائيات (Stats Slot)
```vue
<PageHeader
  title="نظرة عامة"
  subtitle="إحصائيات سريعة للنظام"
  badge="مباشر"
  badgeType="success"
>
  <template #actions>
    <button class="page-action-btn secondary">
      <i class="mdi mdi-refresh"></i>
      تحديث
    </button>
  </template>
  
  <template #stats>
    <div style="display: grid; grid-template-columns: repeat(auto-fit, minmax(150px, 1fr)); gap: 20px; padding-top: 20px;">
      <div style="text-align: center;">
        <div style="font-size: 28px; font-weight: 700; color: white;">245</div>
        <div style="font-size: 13px; color: rgba(255,255,255,0.6); margin-top: 4px;">إجمالي المشاريع</div>
      </div>
      <div style="text-align: center;">
        <div style="font-size: 28px; font-weight: 700; color: white;">182</div>
        <div style="font-size: 13px; color: rgba(255,255,255,0.6); margin-top: 4px;">المشاريع النشطة</div>
      </div>
      <div style="text-align: center;">
        <div style="font-size: 28px; font-weight: 700; color: white;">2.4M</div>
        <div style="font-size: 13px; color: rgba(255,255,255,0.6); margin-top: 4px;">الإيرادات</div>
      </div>
      <div style="text-align: center;">
        <div style="font-size: 28px; font-weight: 700; color: white;">600K</div>
        <div style="font-size: 13px; color: rgba(255,255,255,0.6); margin-top: 4px;">صافي الربح</div>
      </div>
    </div>
  </template>
</PageHeader>
```

## مثال 8: صفحة بسيطة بدون أزرار
```vue
<PageHeader
  title="سياسة الخصوصية"
  subtitle="اطلع على سياسة الخصوصية وشروط الاستخدام"
  :breadcrumbs="[
    { label: 'الرئيسية', to: '/', icon: 'mdi mdi-home' },
    { label: 'سياسة الخصوصية', to: '/privacy' }
  ]"
/>
```

## مثال 9: صفحة تفاصيل مع حالة
```vue
<PageHeader
  title="تفاصيل المشروع #1234"
  subtitle="مشروع المجمع السكني الحديث"
  badge="قيد التنفيذ"
  badgeType="info"
  :breadcrumbs="[
    { label: 'الرئيسية', to: '/', icon: 'mdi mdi-home' },
    { label: 'المشاريع', to: '/projects' },
    { label: 'تفاصيل المشروع', to: '/projects/1234' }
  ]"
>
  <template #actions>
    <button class="page-action-btn secondary">
      <i class="mdi mdi-share-variant"></i>
      مشاركة
    </button>
    <button class="page-action-btn secondary">
      <i class="mdi mdi-pencil"></i>
      تعديل
    </button>
    <button class="page-action-btn primary">
      <i class="mdi mdi-check"></i>
      إكمال المشروع
    </button>
  </template>
</PageHeader>
```

## مثال 10: صفحة مع تحذير
```vue
<PageHeader
  title="المشاريع المتأخرة"
  subtitle="المشاريع التي تجاوزت الموعد المحدد"
  badge="يتطلب انتباه"
  badgeType="danger"
  :breadcrumbs="[
    { label: 'الرئيسية', to: '/', icon: 'mdi mdi-home' },
    { label: 'المشاريع', to: '/projects' },
    { label: 'المتأخرة', to: '/projects/delayed' }
  ]"
>
  <template #actions>
    <button class="page-action-btn secondary">
      <i class="mdi mdi-email"></i>
      إرسال تنبيه
    </button>
    <button class="page-action-btn primary">
      <i class="mdi mdi-calendar-clock"></i>
      تمديد الموعد
    </button>
  </template>
</PageHeader>
```
