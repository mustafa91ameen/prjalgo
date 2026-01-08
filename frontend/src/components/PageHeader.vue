<template>
  <div class="page-header-wrapper">
    <div class="page-header-container">
      <!-- Right Section: Title & Subtitle -->
      <div class="page-header-info">
        <div class="page-header-top">
          <h1 class="page-header-title">{{ title }}</h1>
          <span v-if="badge" class="page-header-badge" :class="badgeType">
            {{ badge }}
          </span>
        </div>
        <p v-if="subtitle" class="page-header-subtitle">{{ subtitle }}</p>
        
        <!-- Breadcrumb -->
        <nav v-if="breadcrumbs && breadcrumbs.length" class="page-breadcrumb">
          <router-link 
            v-for="(crumb, index) in breadcrumbs" 
            :key="index"
            :to="crumb.to"
            class="breadcrumb-item"
            :class="{ 'active': index === breadcrumbs.length - 1 }"
          >
            <i v-if="crumb.icon" :class="crumb.icon" class="breadcrumb-icon"></i>
            {{ crumb.label }}
            <i v-if="index < breadcrumbs.length - 1" class="mdi mdi-chevron-left breadcrumb-separator"></i>
          </router-link>
        </nav>
      </div>

      <!-- Left Section: Actions -->
      <div v-if="$slots.actions" class="page-header-actions">
        <slot name="actions"></slot>
      </div>
    </div>

    <!-- Optional Stats Bar -->
    <div v-if="$slots.stats" class="page-header-stats">
      <slot name="stats"></slot>
    </div>
  </div>
</template>

<script setup>
import { defineProps } from 'vue'

const props = defineProps({
  title: {
    type: String,
    required: true
  },
  subtitle: {
    type: String,
    default: ''
  },
  badge: {
    type: String,
    default: ''
  },
  badgeType: {
    type: String,
    default: 'primary', // primary, success, warning, danger, info
    validator: (value) => ['primary', 'success', 'warning', 'danger', 'info'].includes(value)
  },
  breadcrumbs: {
    type: Array,
    default: () => []
    // Example: [{ label: 'الرئيسية', to: '/', icon: 'mdi mdi-home' }, { label: 'المشاريع', to: '/projects' }]
  }
})
</script>

<style scoped>
@import '../styles/page-header.css';
</style>
