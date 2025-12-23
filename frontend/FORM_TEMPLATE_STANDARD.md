# Template موحد لجميع نماذج الإضافة - Standard Form Template

## 📋 نظرة عامة

هذا التنسيق هو **المعيار الموحد** لجميع نماذج الإضافة في المنصة. يجب استخدامه في جميع الصفحات عند إضافة عناصر جديدة.

## 🎨 الملف الأساسي

جميع الأنماط محفوظة في: `/src/styles/profile-form.css`

يتم استيرادها في: `/src/styles/main.css`

## 📐 البنية القياسية

### 1. Dialog Container
```vue
<v-dialog v-model="dialog" max-width="900" scrollable persistent>
  <v-card class="task-dialog-card profile-form-card">
```

### 2. Header Section
```vue
<v-card-title class="task-dialog-header profile-form-header">
  <h2 class="profile-form-title">
    {{ isEditing ? 'تعديل العنصر' : 'معلومات العنصر' }}
  </h2>
</v-card-title>
```

### 3. Content Section
```vue
<v-card-text class="profile-form-content">
  <p class="profile-form-instruction">
    لإتمام {{ isEditing ? 'تعديل' : 'إضافة' }} العنصر، يرجى توفير المعلومات التالية. يرجى ملاحظة أن جميع الحقول المميزة بعلامة النجمة (*) مطلوبة.
  </p>

  <v-form ref="form" v-model="formValid">
    <!-- الصف الأول -->
    <v-row class="profile-form-row">
      <v-col cols="12" md="4" class="profile-form-column">
        <div class="profile-form-field-wrapper">
          <label class="profile-form-label">
            اسم الحقل <span class="required-star">*</span>
          </label>
          <v-text-field
            v-model="formData.fieldName"
            variant="outlined"
            density="comfortable"
            placeholder="أدخل القيمة"
            :rules="[v => !!v || 'هذا الحقل مطلوب']"
            required
            hide-details="auto"
            class="profile-form-input"
          />
        </div>
      </v-col>
    </v-row>
  </v-form>
</v-card-text>
```

### 4. Footer Actions
```vue
<v-card-actions class="profile-form-actions">
  <v-spacer />
  <v-btn
    class="profile-form-cancel-btn"
    variant="outlined"
    @click="closeDialog"
  >
    إلغاء
  </v-btn>
  <v-btn
    class="profile-form-continue-btn"
    variant="elevated"
    :disabled="!formValid"
    @click="saveItem"
  >
    {{ isEditing ? 'تحديث العنصر' : 'حفظ العنصر' }}
  </v-btn>
</v-card-actions>
```

## 🎯 الفئات المستخدمة (CSS Classes)

### Container Classes
- `profile-form-card` - البطاقة الرئيسية
- `task-dialog-card` - للتوافق مع الأنظمة القديمة

### Header Classes
- `profile-form-header` - رأس النموذج
- `task-dialog-header` - للتوافق
- `profile-form-title` - عنوان النموذج

### Content Classes
- `profile-form-content` - محتوى النموذج
- `profile-form-instruction` - نص التعليمات

### Form Classes
- `profile-form-row` - صف النموذج
- `profile-form-column` - عمود النموذج
- `profile-form-field-wrapper` - غلاف الحقل
- `profile-form-label` - تسمية الحقل
- `profile-form-input` - حقل الإدخال
- `required-star` - علامة الحقل المطلوب

### Footer Classes
- `profile-form-actions` - أزرار الإجراءات
- `profile-form-cancel-btn` - زر الإلغاء
- `profile-form-continue-btn` - زر الحفظ

## 📝 أمثلة الاستخدام

### مثال 1: نموذج إضافة مشروع
```vue
<!-- موجود في: src/pages/project-management.vue -->
```

### مثال 2: نموذج إضافة مادة
```vue
<!-- موجود في: src/pages/materials-expenses-details.vue -->
```

### مثال 3: نموذج إضافة مهندس
```vue
<!-- موجود في: src/pages/engineers.vue -->
```

### مثال 4: نموذج إضافة مهمة
```vue
<!-- موجود في: src/pages/task-management.vue -->
```

### مثال 5: نموذج إضافة يوم عمل
```vue
<!-- موجود في: src/pages/work-days.vue -->
```

### مثال 6: نموذج إضافة مصروف
```vue
<!-- موجود في: src/pages/project-expenses.vue -->
```

## ✅ قائمة الصفحات التي تستخدم التنسيق الموحد

- ✅ `/src/pages/project-management.vue` - إدارة المشاريع
- ✅ `/src/pages/materials-expenses-details.vue` - المواد والمصروفات
- ✅ `/src/pages/engineers.vue` - المهندسين
- ✅ `/src/pages/task-management.vue` - إدارة المهام
- ✅ `/src/pages/work-days.vue` - أيام العمل
- ✅ `/src/pages/project-expenses.vue` - مصروفات المشاريع

## 🔧 التعديلات المطلوبة

عند إنشاء نموذج جديد أو تعديل نموذج موجود:

1. **استخدم الفئات المذكورة أعلاه**
2. **اتبع البنية القياسية**
3. **لا تضيف أنماط مخصصة** - استخدم الفئات الموجودة
4. **راجع الملف `/src/styles/profile-form.css`** للتأكد من الفئات المتاحة

## 📌 ملاحظات مهمة

- **لا تغير ملف `profile-form.css`** إلا بعد التنسيق مع الفريق
- **استخدم `max-width="900"`** للـ dialog
- **استخدم `scrollable persistent`** للـ dialog
- **جميع الحقول المطلوبة** يجب أن تحتوي على `<span class="required-star">*</span>`
- **استخدم `formValid`** للتحكم في تفعيل زر الحفظ

## 🎨 المواصفات الفنية

- **Border Radius**: 8px
- **Padding**: 24px للمحتوى، 20px 24px للرأس
- **Font Size**: 
  - العنوان: 1.35rem
  - التعليمات: 0.85rem
  - التسميات: 0.8125rem
  - الحقول: 0.875rem
- **Colors**:
  - الخلفية: #ffffff
  - النص: #1e293b
  - الحدود: #d1d5db
  - الحقل المطلوب: #ef4444 (للنجمة)

---

**آخر تحديث**: تم حفظ التنسيق كمعيار موحد لجميع النماذج




