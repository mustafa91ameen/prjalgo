<template>
  <div class="stats-container mb-6">
    <v-row>
      <v-col cols="12" md="3">
        <v-card class="pa-4 text-center" color="primary" variant="tonal">
          <v-icon size="48" color="primary" class="mb-2">mdi-account-multiple</v-icon>
          <h3 class="text-h4 font-weight-bold">{{ totalUsers }}</h3>
          <p class="text-subtitle-1">إجمالي المستخدمين</p>
        </v-card>
      </v-col>
      <v-col cols="12" md="3">
        <v-card class="pa-4 text-center" color="success" variant="tonal">
          <v-icon size="48" color="success" class="mb-2">mdi-check-circle</v-icon>
          <h3 class="text-h4 font-weight-bold">{{ activeUsers }}</h3>
          <p class="text-subtitle-1">مستخدمين نشطين</p>
        </v-card>
      </v-col>
      <v-col cols="12" md="3">
        <v-card class="pa-4 text-center" color="warning" variant="tonal">
          <v-icon size="48" color="warning" class="mb-2">mdi-account-clock</v-icon>
          <h3 class="text-h4 font-weight-bold">{{ pendingUsers }}</h3>
          <p class="text-subtitle-1">مستخدمين معلقين</p>
        </v-card>
      </v-col>
      <v-col cols="12" md="3">
        <v-card class="pa-4 text-center" color="info" variant="tonal">
          <v-icon size="48" color="info" class="mb-2">mdi-shield-account</v-icon>
          <h3 class="text-h4 font-weight-bold">{{ adminUsers }}</h3>
          <p class="text-subtitle-1">مدراء النظام</p>
        </v-card>
      </v-col>
    </v-row>
  </div>
</template>

<script setup>
defineProps({
  totalUsers: {
    type: Number,
    default: 0
  },
  activeUsers: {
    type: Number,
    default: 0
  },
  pendingUsers: {
    type: Number,
    default: 0
  },
  adminUsers: {
    type: Number,
    default: 0
  }
})
</script>

<style scoped>
@import './styles/users.css';
</style>
