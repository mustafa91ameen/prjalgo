<template>
  <v-dialog v-model="dialogModel" max-width="600px">
    <v-card class="debtor-form-card">
      <v-card-title class="form-header">
        <v-icon class="me-2" color="white">{{ isEdit ? 'mdi-pencil' : 'mdi-plus' }}</v-icon>
        {{ isEdit ? 'تعديل بيانات المديون' : 'إضافة مديون جديد' }}
      </v-card-title>
      <v-card-text class="pa-6">
        <v-form ref="formRef" v-model="formValid">
          <v-row>
            <v-col cols="12" md="6">
              <v-text-field
                v-model="form.name"
                label="الاسم الكامل"
                variant="outlined"
                :rules="[v => !!v || 'الاسم مطلوب']"
                required
              />
            </v-col>
            <v-col cols="12" md="6">
              <v-text-field
                v-model="form.email"
                label="البريد الإلكتروني"
                type="email"
                variant="outlined"
                :rules="[v => !!v || 'البريد الإلكتروني مطلوب']"
                required
              />
            </v-col>
            <v-col cols="12" md="6">
              <v-text-field
                v-model="form.phone"
                label="رقم الهاتف"
                variant="outlined"
              />
            </v-col>
            <v-col cols="12" md="6">
              <v-text-field
                v-model="form.amount"
                label="المبلغ المطلوب"
                type="number"
                variant="outlined"
                :rules="[v => !!v || 'المبلغ مطلوب']"
                required
              />
            </v-col>
            <v-col cols="12" md="6">
              <v-select
                v-model="form.currency"
                label="العملة"
                :items="currencyOptions"
                variant="outlined"
                required
              />
            </v-col>
            <v-col cols="12" md="6">
              <v-text-field
                v-model="form.dueDate"
                label="تاريخ الاستحقاق"
                type="date"
                variant="outlined"
                :rules="[v => !!v || 'تاريخ الاستحقاق مطلوب']"
                required
              />
            </v-col>
            <v-col cols="12" md="6">
              <v-select
                v-model="form.status"
                label="الحالة"
                :items="statusOptions"
                variant="outlined"
                required
              />
            </v-col>
            <v-col cols="12">
              <v-textarea
                v-model="form.notes"
                label="ملاحظات"
                variant="outlined"
                rows="3"
              />
            </v-col>
          </v-row>
        </v-form>
      </v-card-text>
      <v-card-actions class="pa-4">
        <v-spacer />
        <v-btn
          color="grey"
          variant="text"
          @click="closeDialog"
        >
          إلغاء
        </v-btn>
        <v-btn
          color="primary"
          variant="elevated"
          :disabled="!formValid"
          @click="saveDebtor"
        >
          {{ isEdit ? 'تحديث' : 'حفظ' }}
        </v-btn>
      </v-card-actions>
    </v-card>
  </v-dialog>
</template>

<script setup>
import { ref, computed, watch } from 'vue'

const props = defineProps({
  modelValue: {
    type: Boolean,
    default: false
  },
  debtor: {
    type: Object,
    default: null
  },
  isEdit: {
    type: Boolean,
    default: false
  }
})

const emit = defineEmits(['update:modelValue', 'save', 'close'])

const dialogModel = computed({
  get: () => props.modelValue,
  set: (value) => emit('update:modelValue', value)
})

const formRef = ref(null)
const formValid = ref(false)
const form = ref({
  name: '',
  email: '',
  phone: '',
  amount: '',
  currency: 'IQD',
  dueDate: '',
  status: 'pending',
  notes: ''
})

const currencyOptions = ['IQD', 'USD', 'EUR', 'SAR']
const statusOptions = [
  { title: 'قيد الانتظار', value: 'pending' },
  { title: 'مدفوع', value: 'paid' },
  { title: 'متأخر', value: 'overdue' }
]

watch(() => props.debtor, (newDebtor) => {
  if (newDebtor) {
    form.value = { ...newDebtor }
  } else {
    resetForm()
  }
}, { immediate: true })

const resetForm = () => {
  form.value = {
    name: '',
    email: '',
    phone: '',
    amount: '',
    currency: 'IQD',
    dueDate: '',
    status: 'pending',
    notes: ''
  }
}

const closeDialog = () => {
  dialogModel.value = false
  resetForm()
  emit('close')
}

const saveDebtor = () => {
  if (formValid.value) {
    emit('save', { ...form.value, amount: parseFloat(form.value.amount) })
    closeDialog()
  }
}
</script>

<style scoped>
.debtor-form-card {
  border-radius: 20px !important;
  overflow: hidden !important;
}

.form-header {
  background: linear-gradient(135deg, #3b82f6 0%, #2563eb 50%, #1d4ed8 100%) !important;
  color: white !important;
  font-weight: 700 !important;
  padding: 20px 24px !important;
}
</style>
