<template>
  <v-card class="ds-data-table-card" elevation="2">
    <v-card-title v-if="title" class="ds-data-table-header">
      <div class="ds-data-table-header-content">
        <div class="ds-data-table-header-left">
          <v-icon v-if="icon" class="ds-data-table-header-icon" :size="iconSize" :color="iconColor">
            {{ icon }}
          </v-icon>
          <div>
            <h2 class="ds-data-table-title">{{ title }}</h2>
            <p v-if="subtitle" class="ds-data-table-subtitle">{{ subtitle }}</p>
          </div>
          <v-chip v-if="count !== null" class="ds-data-table-count-chip" size="x-small" variant="elevated">
            {{ count }}
          </v-chip>
        </div>
        <div class="ds-data-table-header-right">
          <slot name="header-actions"></slot>
        </div>
      </div>
    </v-card-title>

    <v-card-text v-if="showSearch" class="ds-data-table-search-container">
      <div class="ds-data-table-search-box">
        <v-text-field
          v-model="searchQuery"
          :placeholder="searchPlaceholder"
          prepend-inner-icon="mdi-magnify"
          variant="outlined"
          density="compact"
          hide-details
          class="ds-data-table-search-field"
          @keyup.enter="handleSearch"
        />
        <v-btn
          v-if="showSearchButton"
          class="ds-data-table-search-btn"
          @click="handleSearch"
          size="small"
        >
          <v-icon class="me-2" size="18">mdi-magnify</v-icon>
          بحث
        </v-btn>
      </div>
    </v-card-text>

    <v-data-table
      :headers="headers"
      :items="items"
      :search="searchQuery"
      :items-per-page="itemsPerPage"
      :class="tableClass"
      :no-data-text="noDataText"
      :loading-text="loadingText"
      :density="density"
      v-bind="$attrs"
    >
      <template v-for="(_, slot) in $slots" v-slot:[slot]="scope">
        <slot :name="slot" v-bind="scope" />
      </template>
    </v-data-table>
  </v-card>
</template>

<script setup>
import { ref, computed } from 'vue'

const props = defineProps({
  title: {
    type: String,
    default: ''
  },
  subtitle: {
    type: String,
    default: ''
  },
  icon: {
    type: String,
    default: ''
  },
  iconSize: {
    type: [String, Number],
    default: 18
  },
  iconColor: {
    type: String,
    default: 'white'
  },
  count: {
    type: [Number, String],
    default: null
  },
  headers: {
    type: Array,
    required: true
  },
  items: {
    type: Array,
    required: true
  },
  showSearch: {
    type: Boolean,
    default: true
  },
  showSearchButton: {
    type: Boolean,
    default: false
  },
  searchPlaceholder: {
    type: String,
    default: 'البحث...'
  },
  itemsPerPage: {
    type: Number,
    default: 10
  },
  tableClass: {
    type: String,
    default: 'ds-data-table'
  },
  noDataText: {
    type: String,
    default: 'لا توجد بيانات متاحة'
  },
  loadingText: {
    type: String,
    default: 'جاري التحميل...'
  },
  density: {
    type: String,
    default: 'comfortable'
  },
  headerColor: {
    type: String,
    default: 'primary' // 'primary', 'secondary', 'green', 'blue'
  }
})

const emit = defineEmits(['search', 'update:search'])

const searchQuery = ref('')

const handleSearch = () => {
  emit('search', searchQuery.value)
}

const headerGradient = computed(() => {
  const gradients = {
    primary: 'linear-gradient(135deg, rgb(180, 155, 163) 0%, rgb(150, 122, 129) 50%, rgb(120, 95, 103) 100%)',
    secondary: 'linear-gradient(135deg, rgb(165, 165, 165) 0%, rgb(132, 132, 132) 50%, rgb(105, 105, 105) 100%)',
    green: 'linear-gradient(135deg, rgb(16, 185, 129) 0%, rgb(5, 150, 105) 50%, rgb(4, 120, 87) 100%)',
    blue: 'linear-gradient(135deg, rgb(180, 155, 163) 0%, rgb(150, 122, 129) 50%, rgb(120, 95, 103) 100%)'
  }
  return gradients[props.headerColor] || gradients.primary
})
</script>

<style scoped>
.ds-data-table-card {
  border-radius: var(--ds-radius-xl);
  overflow: hidden;
  margin-top: var(--ds-space-xl);
}

.ds-data-table-header {
  background: v-bind(headerGradient);
  color: var(--ds-text-white);
  padding: 8px 12px;
  border-radius: var(--ds-radius-lg) var(--ds-radius-lg) 0 0;
  box-shadow: var(--ds-shadow-md);
  margin: 0;
  font-weight: var(--ds-font-weight-semibold);
  font-size: var(--ds-font-size-md);
  text-align: right;
  display: flex;
  align-items: center;
  justify-content: space-between;
  min-height: 36px;
}

.ds-data-table-header-content {
  display: flex;
  align-items: center;
  justify-content: space-between;
  width: 100%;
  gap: 1rem;
}

.ds-data-table-header-left {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  flex: 1;
}

.ds-data-table-header-right {
  display: flex;
  align-items: center;
  gap: 0.5rem;
}

.ds-data-table-header-icon {
  color: var(--ds-text-white);
}

.ds-data-table-title {
  color: var(--ds-text-white);
  font-weight: var(--ds-font-weight-semibold);
  font-size: var(--ds-font-size-md);
  text-shadow: 0 1px 3px rgba(0, 0, 0, 0.25);
  margin: 0;
}

.ds-data-table-subtitle {
  color: rgba(255, 255, 255, 0.9);
  font-weight: var(--ds-font-weight-medium);
  font-size: var(--ds-font-size-sm);
  text-shadow: 0 1px 2px rgba(0, 0, 0, 0.25);
  margin: 0.25rem 0 0 0;
}

.ds-data-table-count-chip {
  background-color: rgba(255, 255, 255, 0.3);
  color: var(--ds-text-white);
  border: 1px solid rgba(255, 255, 255, 0.5);
  font-size: var(--ds-font-size-xs);
  height: 20px;
  font-weight: var(--ds-font-weight-semibold);
}

.ds-data-table-search-container {
  padding: var(--ds-search-container-padding);
  background: linear-gradient(135deg, rgb(250, 245, 247) 0%, rgb(240, 230, 235) 50%, rgb(230, 215, 223) 100%);
  border-radius: var(--ds-search-container-border-radius);
  margin-bottom: var(--ds-space-md);
  box-shadow: var(--ds-shadow-sm);
  border: 1px solid rgb(180, 155, 163);
  position: relative;
}

.ds-data-table-search-container::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  height: 2px;
  background: linear-gradient(90deg, rgb(180, 155, 163), rgb(150, 122, 129), rgb(120, 95, 103));
  border-radius: var(--ds-radius-lg) var(--ds-radius-lg) 0 0;
}

.ds-data-table-search-box {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  max-width: 600px;
  margin: 0 auto;
}

.ds-data-table-search-field {
  flex: 1;
  margin-bottom: 0;
}

.ds-data-table-search-field :deep(.v-field) {
  background-color: var(--ds-bg-primary);
  border-radius: var(--ds-radius-lg);
  box-shadow: var(--ds-shadow-sm);
  border: 2px solid rgb(120, 95, 103);
  transition: var(--ds-transition-normal);
  min-height: var(--ds-search-field-height);
}

.ds-data-table-search-field :deep(.v-label) {
  color: rgb(120, 95, 103);
  font-weight: var(--ds-font-weight-semibold);
  font-size: var(--ds-search-field-font-size);
  opacity: 1;
}

.ds-data-table-search-field :deep(.v-field__input input) {
  color: var(--ds-text-primary);
  font-weight: var(--ds-font-weight-medium);
  font-size: var(--ds-search-field-font-size);
}

.ds-data-table-search-field :deep(.v-field__prepend-inner .v-icon) {
  color: rgb(120, 95, 103);
  font-size: 18px;
}

.ds-data-table-search-btn {
  height: var(--ds-btn-height-md);
  min-width: 80px;
  border-radius: var(--ds-radius-lg);
  font-weight: var(--ds-font-weight-semibold);
  font-size: var(--ds-font-size-base);
  text-transform: none;
  box-shadow: var(--ds-shadow-sm);
  transition: var(--ds-transition-normal);
}

.ds-data-table-search-btn:hover {
  transform: translateY(-2px);
  box-shadow: var(--ds-shadow-md);
}
</style>

<style>
/* Global Data Table Styles */
.ds-data-table :deep(.v-data-table__th) {
  color: var(--ds-table-header-text);
  font-weight: var(--ds-table-header-font-weight);
  font-size: var(--ds-table-header-font-size);
  background: var(--ds-table-header-bg);
  text-shadow: 0 1px 2px rgba(0, 0, 0, 0.25);
  padding: var(--ds-table-header-padding);
  min-height: var(--ds-table-header-min-height);
  border-bottom: var(--ds-table-border-width) solid var(--ds-table-border-color);
  position: relative;
}

.ds-data-table :deep(.v-data-table__td) {
  color: var(--ds-table-cell-text);
  font-weight: var(--ds-table-cell-font-weight);
  font-size: var(--ds-table-cell-font-size);
  text-shadow: 0 1px 2px rgba(0, 0, 0, 0.1);
  padding: var(--ds-table-cell-padding);
  border-right: var(--ds-table-border-width) solid var(--ds-table-border-color);
  border-bottom: var(--ds-table-border-width) solid var(--ds-table-border-color);
}

.ds-data-table :deep(.v-data-table__wrapper table),
.ds-data-table :deep(.v-data-table__wrapper table td),
.ds-data-table :deep(.v-data-table__wrapper table th),
.ds-data-table :deep(.v-data-table__wrapper table tr) {
  border-color: var(--ds-table-border-color);
}
</style>

