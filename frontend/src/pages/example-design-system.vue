<template>
  <!-- Example page using the design system -->
  <div class="ds-page-container">
    <!-- Page Header -->
    <PageHeader
      title="مثال على نظام التصميم"
      subtitle="صفحة توضيحية لاستخدام نظام التصميم الموحد"
      icon="mdi-palette"
      :show-back-button="true"
    />

    <!-- Content Container -->
    <div class="ds-content-container">
      <!-- Search Container -->
      <div class="ds-search-container">
        <div class="ds-search-box">
          <v-text-field
            v-model="searchQuery"
            placeholder="البحث..."
            prepend-inner-icon="mdi-magnify"
            variant="outlined"
            density="compact"
            hide-details
            class="ds-search-field"
          />
          <v-btn class="ds-search-btn">
            <v-icon class="me-2" size="18">mdi-magnify</v-icon>
            بحث
          </v-btn>
        </div>
      </div>

      <!-- Cards Grid -->
      <v-row>
        <v-col cols="12" md="4">
          <div class="ds-card">
            <h3 class="ds-text-primary">بطاقة 1</h3>
            <p class="ds-text-secondary">محتوى البطاقة</p>
          </div>
        </v-col>
        <v-col cols="12" md="4">
          <div class="ds-card">
            <h3 class="ds-text-primary">بطاقة 2</h3>
            <p class="ds-text-secondary">محتوى البطاقة</p>
          </div>
        </v-col>
        <v-col cols="12" md="4">
          <div class="ds-card">
            <h3 class="ds-text-primary">بطاقة 3</h3>
            <p class="ds-text-secondary">محتوى البطاقة</p>
          </div>
        </v-col>
      </v-row>

      <!-- Data Table -->
      <DataTable
        title="جدول البيانات"
        subtitle="مثال على استخدام مكون الجدول"
        icon="mdi-table"
        :count="tableItems.length"
        :headers="tableHeaders"
        :items="tableItems"
        header-color="primary"
      >
        <template v-slot:item.actions="{ item }">
          <div class="ds-actions-buttons">
            <v-btn icon class="ds-action-btn" size="small">
              <v-icon>mdi-pencil</v-icon>
            </v-btn>
            <v-btn icon class="ds-action-btn" size="small">
              <v-icon>mdi-delete</v-icon>
            </v-btn>
          </div>
        </template>
      </DataTable>

      <!-- Buttons Example -->
      <div class="ds-p-lg" style="display: flex; gap: 1rem; flex-wrap: wrap;">
        <v-btn class="ds-btn ds-btn-primary ds-btn-sm">زر صغير</v-btn>
        <v-btn class="ds-btn ds-btn-primary">زر عادي</v-btn>
        <v-btn class="ds-btn ds-btn-primary ds-btn-lg">زر كبير</v-btn>
      </div>

      <!-- Section Title Example -->
      <div class="ds-info-card">
        <div class="ds-section-title">
          <v-icon class="ds-section-title-icon">mdi-information</v-icon>
          <span>عنوان القسم</span>
        </div>
        <div class="ds-info-item">
          <div class="ds-info-label">التسمية</div>
          <div class="ds-info-value">القيمة</div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import PageHeader from '@/components/PageHeader.vue'
import DataTable from '@/components/DataTable.vue'

const searchQuery = ref('')

const tableHeaders = [
  { title: 'التسلسل', key: 'id', sortable: false },
  { title: 'الاسم', key: 'name', sortable: true },
  { title: 'البريد', key: 'email', sortable: true },
  { title: 'الإجراءات', key: 'actions', sortable: false }
]

const tableItems = [
  { id: 1, name: 'أحمد محمد', email: 'ahmed@example.com' },
  { id: 2, name: 'فاطمة علي', email: 'fatima@example.com' },
  { id: 3, name: 'محمد حسن', email: 'mohammed@example.com' }
]
</script>

<style scoped>
/* Additional custom styles can be added here */
/* Most styles are handled by the design system */
</style>

