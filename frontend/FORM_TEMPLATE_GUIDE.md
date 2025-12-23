# دليل استخدام النموذج الموحد (Profile Form Template)

## نظرة عامة
تم إنشاء نموذج موحد لجميع نماذج الإضافة في التطبيق لضمان التنسيق المتسق. هذا النموذج يستخدم فئات CSS محددة يمكن تطبيقها على أي نموذج إضافة.

## البنية الأساسية

```vue
<v-dialog v-model="dialog" max-width="900" scrollable persistent>
  <v-card class="task-dialog-card profile-form-card">
    <!-- Header Section -->
    <v-card-title class="task-dialog-header profile-form-header">
      <v-btn icon="mdi-close" @click="closeDialog" variant="text" class="close-btn">
        <v-icon>mdi-close</v-icon>
      </v-btn>
      <h2 class="profile-form-title">
        {{ isEditing ? 'تعديل العنصر' : 'معلومات العنصر' }}
      </h2>
    </v-card-title>

    <!-- Form Content -->
    <v-card-text class="profile-form-content">
      <p class="profile-form-instruction">
        لإتمام {{ isEditing ? 'تعديل' : 'إضافة' }} العنصر، يرجى توفير المعلومات التالية. 
        يرجى ملاحظة أن جميع الحقول المميزة بعلامة النجمة (*) مطلوبة.
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

    <!-- Footer Actions -->
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
        {{ isEditing ? 'تحديث' : 'حفظ' }}
      </v-btn>
    </v-card-actions>
  </v-card>
</v-dialog>
```

## الفئات المطلوبة

### الفئات الأساسية:
- `profile-form-card` - للبطاقة الرئيسية
- `profile-form-header` - لرأس النموذج
- `profile-form-title` - لعنوان النموذج
- `profile-form-content` - لمحتوى النموذج
- `profile-form-instruction` - لنص التعليمات
- `profile-form-row` - للصفوف
- `profile-form-column` - للأعمدة
- `profile-form-field-wrapper` - لغطاء الحقل
- `profile-form-label` - للتسميات
- `profile-form-input` - للحقول
- `profile-form-actions` - لأزرار الإجراءات
- `profile-form-cancel-btn` - لزر الإلغاء
- `profile-form-continue-btn` - لزر الحفظ

### الفئات الخاصة:
- `required-star` - للنجمة الحمراء للحقول المطلوبة
- `profile-form-label-row` - لتسمية مع قيمة (مثل نسبة الإنجاز)
- `progress-value` - لعرض قيمة النسبة
- `profile-form-slider` - للشريط المنزلق

## أمثلة الاستخدام

### مثال 1: نموذج بسيط
```vue
<v-col cols="12" md="4" class="profile-form-column">
  <div class="profile-form-field-wrapper">
    <label class="profile-form-label">
      الاسم <span class="required-star">*</span>
    </label>
    <v-text-field
      v-model="form.name"
      variant="outlined"
      density="comfortable"
      :rules="[v => !!v || 'الاسم مطلوب']"
      required
      class="profile-form-input"
    />
  </div>
</v-col>
```

### مثال 2: حقل مع شريط منزلق
```vue
<v-col cols="12" md="6" class="profile-form-column">
  <div class="profile-form-field-wrapper">
    <div class="profile-form-label-row">
      <label class="profile-form-label">نسبة الإنجاز</label>
      <span class="progress-value">{{ form.progress }}</span>
    </div>
    <v-slider
      v-model="form.progress"
      min="0"
      max="100"
      step="5"
      thumb-label
      color="primary"
      class="profile-form-slider"
    />
  </div>
</v-col>
```

### مثال 3: حقل نصي طويل
```vue
<v-col cols="12" class="profile-form-column">
  <div class="profile-form-field-wrapper">
    <label class="profile-form-label">ملاحظات</label>
    <v-textarea
      v-model="form.notes"
      variant="outlined"
      rows="4"
      density="comfortable"
      placeholder="أدخل ملاحظات إضافية"
      class="profile-form-input"
    />
  </div>
</v-col>
```

## الصفحات التي يجب تحديثها

1. ✅ `project-management.vue` - تم التطبيق
2. ✅ `work-days.vue` - يستخدم النموذج
3. ✅ `engineers.vue` - يستخدم النموذج
4. ✅ `task-management.vue` - يستخدم النموذج
5. ⚠️ `project-expenses.vue` - يحتاج تحديث
6. ⚠️ `expenses.vue` - يحتاج تحديث
7. ⚠️ `materials-expenses-details.vue` - يستخدم SimpleDialog (يحتاج مراجعة)
8. ⚠️ `labor-details.vue` - يستخدم SimpleDialog (يحتاج مراجعة)
9. ⚠️ `equipment-details.vue` - يحتاج تحديث

## ملاحظات مهمة

1. **الاستيراد**: تم إضافة ملف `profile-form.css` إلى `main.css` تلقائياً
2. **التوافق**: النموذج متوافق مع SimpleDialog و ModernDialog الموجودين
3. **RTL**: جميع العناصر تدعم الاتجاه من اليمين لليسار
4. **الاستجابة**: النموذج متجاوب ويعمل على جميع الأحجام

## التحقق من التطبيق

للتأكد من أن النموذج يعمل بشكل صحيح:
1. افتح أي صفحة تحتوي على نموذج إضافة
2. اضغط على زر الإضافة
3. تحقق من أن التنسيق يطابق نموذج "معلومات المشروع"
4. تحقق من أن جميع الحقول والأزرار تعمل بشكل صحيح




