<template>
  <div class="page-header">
    <div class="page-header-container">
      <div class="page-header-start">
        <!-- Breadcrumbs -->
        <div v-if="breadcrumbs.length > 0" class="page-breadcrumbs">
          <template v-for="(crumb, index) in breadcrumbs" :key="index">
            <span v-if="index > 0" class="breadcrumb-separator">/</span>
            <span :class="{'breadcrumb-active': index === breadcrumbs.length - 1, 'breadcrumb-item': true}">
              {{ crumb.title }}
            </span>
          </template>
        </div>

        <!-- Title Area -->
        <div class="page-header-content">
          <span v-if="icon" class="page-header-icon">{{ icon }}</span>
          <v-icon v-else-if="mdiIcon" :size="iconSize" class="page-header-icon">{{ mdiIcon }}</v-icon>
          <div class="page-header-text">
            <h1 class="page-header-title">{{ title }}</h1>
            <p v-if="subtitle" class="page-header-subtitle">{{ subtitle }}</p>
          </div>
        </div>
      </div>

      <!-- Actions Area -->
      <div class="page-header-actions">
        <slot name="actions"></slot>
      </div>
    </div>
  </div>
</template>

<script setup>
/**
 * PageHeader Component
 * Centralized page header for consistent styling and organization
 * Following FRONTEND_RULES.md - Smart Pages, Dumb Components
 */

defineProps({
  title: {
    type: String,
    required: true
  },
  subtitle: {
    type: String,
    default: ''
  },
  icon: {
    type: String,
    default: ''
  },
  mdiIcon: {
    type: String,
    default: ''
  },
  iconSize: {
    type: [Number, String],
    default: 48
  },
  breadcrumbs: {
    type: Array,
    default: () => []
  }
})
</script>

<style scoped>
@import './styles/shared.css';
</style>
