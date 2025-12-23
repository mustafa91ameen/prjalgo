<template>
  <v-form ref="form" v-model="formValid">
    <!-- الصف الأول: نوع المادة -->
    <v-row class="profile-form-row">
      <v-col cols="12" class="profile-form-column">
        <div class="profile-form-field-wrapper">
          <label class="profile-form-label">
            نوع المادة <span class="required-star">*</span>
          </label>
          <v-text-field
            v-model="formData.type"
            variant="outlined"
            density="comfortable"
            placeholder="أدخل نوع المادة"
            :rules="[v => !!v || 'نوع المادة مطلوب']"
            required
            hide-details="auto"
            class="profile-form-input"
          />
        </div>
      </v-col>
    </v-row>

    <!-- الصف الثاني: الكمية والسعر -->
    <v-row class="profile-form-row">
      <v-col cols="12" md="6" class="profile-form-column">
        <div class="profile-form-field-wrapper">
          <label class="profile-form-label">
            الكمية <span class="required-star">*</span>
          </label>
          <v-text-field
            v-model.number="formData.quantity"
            type="number"
            variant="outlined"
            density="comfortable"
            placeholder="0"
            :rules="[v => (v > 0) || 'الكمية يجب أن تكون أكبر من صفر']"
            required
            hide-details="auto"
            class="profile-form-input"
          />
        </div>
      </v-col>

      <v-col cols="12" md="6" class="profile-form-column">
        <div class="profile-form-field-wrapper">
          <label class="profile-form-label">
            السعر (د.ع) <span class="required-star">*</span>
          </label>
          <v-text-field
            v-model.number="formData.price"
            type="number"
            variant="outlined"
            density="comfortable"
            placeholder="0"
            :rules="[v => (v > 0) || 'السعر يجب أن يكون أكبر من صفر']"
            required
            hide-details="auto"
            class="profile-form-input"
          />
        </div>
      </v-col>
    </v-row>

    <!-- الصف الثالث: السعر الكلي واسم السائق -->
    <v-row class="profile-form-row">
      <v-col cols="12" md="6" class="profile-form-column">
        <div class="profile-form-field-wrapper">
          <label class="profile-form-label">
            السعر الكلي (د.ع)
          </label>
          <v-text-field
            :model-value="computedTotalPrice"
            type="number"
            variant="outlined"
            density="comfortable"
            placeholder="0"
            readonly
            hide-details="auto"
            class="profile-form-input"
          />
        </div>
      </v-col>

      <v-col cols="12" md="6" class="profile-form-column">
        <div class="profile-form-field-wrapper">
          <label class="profile-form-label">
            اسم السائق <span class="required-star">*</span>
          </label>
          <v-text-field
            v-model="formData.driver"
            variant="outlined"
            density="comfortable"
            placeholder="أدخل اسم السائق"
            :rules="[v => !!v || 'اسم السائق مطلوب']"
            required
            hide-details="auto"
            class="profile-form-input"
          />
        </div>
      </v-col>
    </v-row>

    <!-- الصف الرابع: التفاصيل -->
    <v-row class="profile-form-row">
      <v-col cols="12" class="profile-form-column">
        <div class="profile-form-field-wrapper">
          <label class="profile-form-label">التفاصيل</label>
          <v-textarea
            v-model="formData.details"
            variant="outlined"
            rows="4"
            density="comfortable"
            placeholder="أدخل التفاصيل الإضافية"
            hide-details="auto"
            class="profile-form-input"
          />
        </div>
      </v-col>
    </v-row>
  </v-form>
</template>

<script setup>
import { ref, watch, computed } from 'vue'

const props = defineProps({
  modelValue: {
    type: Object,
    default: () => ({})
  }
})

const emit = defineEmits(['update:modelValue'])

const form = ref(null)
const formValid = ref(false)

const formData = ref({
  type: '',
  quantity: '',
  price: '',
  totalPrice: '',
  driver: '',
  details: '',
  ...props.modelValue
})

// حساب السعر الكلي تلقائياً
const computedTotalPrice = computed(() => {
  if (formData.value.quantity && formData.value.price) {
    const total = formData.value.quantity * formData.value.price
    formData.value.totalPrice = total.toString()
    return total
  }
  return formData.value.totalPrice || 0
})

watch([() => formData.value.quantity, () => formData.value.price], ([quantity, price]) => {
  if (quantity && price) {
    formData.value.totalPrice = (quantity * price).toString()
  }
})

// مراقبة التغييرات وإرسالها للوالد
watch(formData, (newValue) => {
  emit('update:modelValue', newValue)
}, { deep: true })

// دالة لإعادة تعيين النموذج
const resetForm = () => {
  formData.value = {
    type: '',
    quantity: '',
    price: '',
    totalPrice: '',
    driver: '',
    details: ''
  }
  if (form.value) {
    form.value.reset()
  }
}

// دالة للتحقق من صحة البيانات
const isValid = () => {
  return formValid.value && formData.value.type && 
         formData.value.quantity && 
         formData.value.price && 
         formData.value.driver
}

defineExpose({
  resetForm,
  isValid,
  formData,
  form
})
</script>

<style scoped>
/* يستخدم التنسيق الموحد من profile-form.css */
</style>
