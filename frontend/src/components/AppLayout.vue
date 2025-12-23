<template>
  <div class="ds-app-layout">
    <!-- Sidebar -->
    <v-navigation-drawer
      v-model="drawer"
      location="right"
      :permanent="permanent"
      :width="sidebarWidth"
      class="ds-sidebar"
    >
      <div class="ds-sidebar-header">
        <div class="ds-sidebar-header-content">
          <v-img
            v-if="logo"
            class="ds-sidebar-logo"
            :height="logoSize"
            :width="logoSize"
            :src="logo"
          />
          <h1 v-if="title" class="ds-sidebar-title">{{ title }}</h1>
          <p v-if="subtitle" class="ds-sidebar-subtitle">{{ subtitle }}</p>
        </div>
      </div>

      <v-list class="ds-sidebar-menu" nav>
        <v-list-item
          v-for="item in menuItems"
          :key="item.title"
          :to="item.to"
          class="ds-sidebar-menu-item"
          rounded="xl"
          :class="{ 'ds-sidebar-menu-item-active': item.active }"
        >
          <template v-slot:prepend>
            <v-icon :color="item.active ? 'primary' : 'white'" class="ds-sidebar-menu-icon">
              {{ item.icon }}
            </v-icon>
          </template>
          <v-list-item-title class="ds-sidebar-menu-text">
            {{ item.title }}
          </v-list-item-title>
          <template v-slot:append v-if="item.badge">
            <v-chip
              :color="item.badgeColor"
              size="small"
              class="ds-sidebar-menu-badge"
            >
              {{ item.badge }}
            </v-chip>
          </template>
        </v-list-item>
      </v-list>

      <template v-slot:append>
        <div class="ds-sidebar-footer">
          <slot name="footer"></slot>
        </div>
      </template>
    </v-navigation-drawer>

    <!-- Main Content -->
    <v-main class="ds-main-content">
      <v-btn
        v-if="!drawer && showToggleButton"
        icon
        @click="drawer = true"
        class="ds-sidebar-toggle-btn"
        color="primary"
        size="large"
      >
        <v-icon>mdi-menu</v-icon>
      </v-btn>
      <slot></slot>
    </v-main>
  </div>
</template>

<script setup>
import { ref, watch } from 'vue'
import { useRoute } from 'vue-router'

const props = defineProps({
  logo: {
    type: String,
    default: ''
  },
  logoSize: {
    type: [String, Number],
    default: 80
  },
  title: {
    type: String,
    default: 'نظام إدارة المشاريع'
  },
  subtitle: {
    type: String,
    default: 'نظام شامل لإدارة المشاريع والمهندسين والمصاريف'
  },
  menuItems: {
    type: Array,
    default: () => []
  },
  permanent: {
    type: Boolean,
    default: true
  },
  sidebarWidth: {
    type: [String, Number],
    default: 320
  },
  showToggleButton: {
    type: Boolean,
    default: true
  }
})

const route = useRoute()
const drawer = ref(props.permanent)

// Update active menu items based on current route
watch(() => route.path, () => {
  props.menuItems.forEach(item => {
    item.active = item.to === route.path
  })
}, { immediate: true })
</script>

<style scoped>
.ds-app-layout {
  display: flex;
  min-height: 100vh;
  direction: rtl;
}

.ds-sidebar {
  background: linear-gradient(135deg, #1976d2 0%, #1565c0 100%);
  border-left: none;
  box-shadow: -4px 0 20px rgba(25, 118, 210, 0.3);
  z-index: var(--ds-z-base);
  direction: rtl;
  text-align: right;
}

.ds-sidebar-header {
  background: rgba(255, 255, 255, 0.1);
  backdrop-filter: blur(10px);
  border-radius: 0 0 20px 20px;
  margin: -16px -16px 16px -16px;
  padding: 12px;
}

.ds-sidebar-header-content {
  text-align: center;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
}

.ds-sidebar-logo {
  filter: drop-shadow(0 4px 8px rgba(0, 0, 0, 0.2));
  border-radius: 50%;
  background: rgba(255, 255, 255, 0.1);
  padding: 8px;
  margin-bottom: 8px;
}

.ds-sidebar-title {
  font-size: 1rem;
  font-weight: var(--ds-font-weight-bold);
  color: var(--ds-text-white);
  margin: 0 0 4px 0;
  text-align: center;
}

.ds-sidebar-subtitle {
  font-size: 0.8rem;
  color: rgba(255, 255, 255, 0.9);
  margin: 0;
  text-align: center;
}

.ds-sidebar-menu {
  padding: 8px;
  gap: 0;
}

.ds-sidebar-menu-item {
  transition: var(--ds-transition-normal);
  border-radius: var(--ds-radius-xl);
  margin-bottom: 1px;
  margin-top: 0;
  min-height: 32px;
  padding: 4px 8px;
}

.ds-sidebar-menu-item:hover {
  background: rgba(255, 255, 255, 0.2);
  transform: translateX(-4px);
}

.ds-sidebar-menu-item-active {
  background: linear-gradient(135deg, #1565c0 0%, #1976d2 100%);
  box-shadow: 0 4px 15px rgba(255, 255, 255, 0.4);
  transform: translateX(-4px);
}

.ds-sidebar-menu-item-active .ds-sidebar-menu-text {
  font-weight: var(--ds-font-weight-semibold);
}

.ds-sidebar-menu-icon {
  margin-inline-end: 12px;
}

.ds-sidebar-menu-text {
  font-size: 0.4rem;
  font-weight: var(--ds-font-weight-medium);
  color: var(--ds-text-white);
}

.ds-sidebar-menu-badge {
  color: var(--ds-text-white);
}

.ds-sidebar-footer {
  padding: 16px;
  text-align: center;
}

.ds-main-content {
  direction: rtl;
  flex: 1;
  background: var(--ds-bg-page);
}

.ds-sidebar-toggle-btn {
  position: fixed;
  top: 16px;
  right: 16px;
  z-index: var(--ds-z-sticky);
}
</style>

