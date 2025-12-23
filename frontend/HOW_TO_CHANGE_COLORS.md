# كيفية تغيير الألوان في المنصة
# How to Change Colors in the Platform

## نظرة عامة - Overview

يمكنك تغيير ألوان المنصة بالكامل من خلال تعديل ملف واحد فقط: `src/styles/design-system.css`

You can change all platform colors by modifying a single file: `src/styles/design-system.css`

## الملف الرئيسي - Main File

**الملف:** `src/styles/design-system.css`

**File:** `src/styles/design-system.css`

## الألوان الحالية - Current Colors

### الألوان الأساسية - Primary Colors
```css
--ds-color-primary: rgb(150, 122, 129);          /* اللون الأساسي */
--ds-color-primary-light: rgb(160, 130, 138);   /* اللون الفاتح */
--ds-color-primary-dark: rgb(100, 70, 78);      /* اللون الداكن */
--ds-color-primary-darker: rgb(80, 55, 63);    /* اللون الأغمق */
```

### الألوان الثانوية - Secondary Colors
```css
--ds-color-secondary: rgb(132, 132, 132);        /* اللون الثانوي */
--ds-color-secondary-light: rgb(165, 165, 165); /* اللون الفاتح */
--ds-color-secondary-dark: rgb(105, 105, 105);   /* اللون الداكن */
--ds-color-secondary-darker: rgb(85, 85, 85);    /* اللون الأغمق */
```

## كيفية التغيير - How to Change

### خطوة 1: اختر الألوان الجديدة
**Step 1: Choose New Colors**

حدد الألوان التي تريد استخدامها. يمكنك استخدام:
- RGB: `rgb(128, 95, 103)`
- HEX: `#807f67`
- HSL: `hsl(345, 15%, 44%)`

### خطوة 2: حدّث الألوان الأساسية
**Step 2: Update Primary Colors**

افتح ملف `src/styles/design-system.css` وابحث عن:

```css
/* Primary Colors - الألوان الأساسية */
--ds-color-primary: rgb(128, 95, 103);
--ds-color-primary-light: rgb(160, 130, 138);
--ds-color-primary-dark: rgb(100, 70, 78);
--ds-color-primary-darker: rgb(80, 55, 63);
```

غيّر القيم إلى الألوان الجديدة. تأكد من إنشاء درجات فاتحة وداكنة متسقة.

### خطوة 3: حدّث الألوان الثانوية
**Step 3: Update Secondary Colors**

```css
/* Secondary Colors - الألوان الثانوية */
--ds-color-secondary: rgb(132, 132, 132);
--ds-color-secondary-light: rgb(165, 165, 165);
--ds-color-secondary-dark: rgb(105, 105, 105);
--ds-color-secondary-darker: rgb(85, 85, 85);
```

### خطوة 4: حدّث التدرجات
**Step 4: Update Gradients**

ابحث عن قسم "Gradients" وحدّث التدرجات:

```css
--ds-gradient-primary: linear-gradient(135deg, rgb(160, 130, 138) 0%, rgb(128, 95, 103) 50%, rgb(100, 70, 78) 100%);
--ds-gradient-secondary: linear-gradient(135deg, rgb(165, 165, 165) 0%, rgb(132, 132, 132) 50%, rgb(105, 105, 105) 100%);
--ds-gradient-header: linear-gradient(135deg, rgb(100, 70, 78) 0%, rgb(128, 95, 103) 25%, rgb(160, 130, 138) 50%, rgb(180, 155, 163) 75%, rgb(200, 180, 188) 100%);
```

استخدم الألوان الأساسية والثانوية الجديدة في التدرجات.

## مثال - Example

لنفترض أنك تريد تغيير اللون الأساسي إلى الأزرق:

**Before:**
```css
--ds-color-primary: rgb(128, 95, 103);
--ds-color-primary-light: rgb(160, 130, 138);
--ds-color-primary-dark: rgb(100, 70, 78);
--ds-color-primary-darker: rgb(80, 55, 63);
```

**After:**
```css
--ds-color-primary: rgb(96, 165, 250);        /* #60a5fa */
--ds-color-primary-light: rgb(147, 197, 253); /* #93c5fd */
--ds-color-primary-dark: rgb(59, 130, 246);    /* #3b82f6 */
--ds-color-primary-darker: rgb(37, 99, 235);  /* #2563eb */
```

ثم حدّث التدرجات:

```css
--ds-gradient-primary: linear-gradient(135deg, rgb(147, 197, 253) 0%, rgb(96, 165, 250) 50%, rgb(59, 130, 246) 100%);
--ds-gradient-header: linear-gradient(135deg, rgb(59, 130, 246) 0%, rgb(96, 165, 250) 25%, rgb(147, 197, 253) 50%, rgb(191, 219, 254) 75%, rgb(219, 234, 254) 100%);
```

## نصائح - Tips

1. **استخدم أدوات الألوان:** استخدم أدوات مثل [Coolors](https://coolors.co) أو [Adobe Color](https://color.adobe.com) لإنشاء لوحة ألوان متناسقة.

2. **أنشئ درجات متسقة:** تأكد من أن الألوان الفاتحة والداكنة متناسقة مع اللون الأساسي.

3. **اختبر التباين:** تأكد من أن النص واضح على الخلفيات الملونة.

4. **احفظ نسخة احتياطية:** احفظ نسخة من الملف قبل التعديل.

## الملفات المتأثرة - Affected Files

عند تغيير الألوان في `design-system.css`، سيتم تحديث:

- ✅ جميع الصفحات التي تستخدم CSS variables
- ✅ مكون PageHeader
- ✅ مكون DataTable
- ✅ جميع الأزرار والبطاقات
- ✅ جميع الجداول
- ✅ جميع التدرجات

## التحقق من التغييرات - Verify Changes

بعد التعديل:

1. احفظ الملف
2. أعد تحميل الصفحة في المتصفح
3. تحقق من أن الألوان تم تحديثها في جميع أنحاء المنصة

## استكشاف الأخطاء - Troubleshooting

إذا لم تظهر التغييرات:

1. تأكد من حفظ الملف
2. امسح ذاكرة التخزين المؤقت للمتصفح (Ctrl+Shift+R أو Cmd+Shift+R)
3. تحقق من أن الملف `main.css` يستورد `design-system.css`
4. تحقق من وحدة التحكم في المتصفح للأخطاء

## الدعم - Support

للمزيد من المعلومات:
- راجع `DESIGN_SYSTEM.md` للتوثيق الكامل
- راجع `DESIGN_SYSTEM_QUICK_REFERENCE.md` للمرجع السريع

---

**ملاحظة:** جميع التغييرات في ملف واحد فقط (`design-system.css`) ستطبق على المنصة بالكامل!

**Note:** All changes in a single file (`design-system.css`) will apply to the entire platform!

