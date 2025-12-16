<template>
  <v-row class="mb-6 stats-row">
    <v-col cols="12" md="3">
      <v-card class="stat-card pa-6 text-center hover-lift card-glow smooth-transition" elevation="2">
        <div class="stat-icon mb-3 icon-glow">
          <v-icon size="64" color="primary">mdi-account-group</v-icon>
        </div>
        <h3 class="text-h3 font-weight-bold text-primary mb-2 stat-number-ltr">{{ totalEngineers }}</h3>
        <p class="text-subtitle-1 text-primary mb-0">إجمالي المهندسين</p>
      </v-card>
    </v-col>
    <v-col cols="12" md="3">
      <v-card class="stat-card pa-6 text-center hover-lift card-glow smooth-transition" elevation="2">
        <div class="stat-icon mb-3 icon-glow">
          <v-icon size="64" color="success">mdi-check-circle</v-icon>
        </div>
        <h3 class="text-h3 font-weight-bold text-success mb-2 stat-number-ltr">{{ activeEngineers }}</h3>
        <p class="text-subtitle-1 text-success mb-0">مهندسين نشطين</p>
      </v-card>
    </v-col>
    <v-col cols="12" md="3">
      <v-card class="stat-card pa-6 text-center hover-lift card-glow smooth-transition" elevation="2">
        <div class="stat-icon mb-3 icon-glow">
          <v-icon size="64" color="info">mdi-briefcase</v-icon>
        </div>
        <h3 class="text-h3 font-weight-bold text-info mb-2 stat-number-ltr">{{ totalProjects }}</h3>
        <p class="text-subtitle-1 text-info mb-0">إجمالي المشاريع</p>
      </v-card>
    </v-col>
    <v-col cols="12" md="3">
      <v-card class="stat-card pa-6 text-center hover-lift card-glow smooth-transition" elevation="2">
        <div class="stat-icon mb-3 icon-glow">
          <v-icon size="64" color="warning">mdi-star</v-icon>
        </div>
        <h3 class="text-h3 font-weight-bold text-warning mb-2 stat-number-ltr">{{ averageRating }}</h3>
        <p class="text-subtitle-1 text-warning mb-0">متوسط التقييم</p>
      </v-card>
    </v-col>
  </v-row>
</template>

<script setup>
defineProps({
  totalEngineers: {
    type: Number,
    default: 0
  },
  activeEngineers: {
    type: Number,
    default: 0
  },
  totalProjects: {
    type: Number,
    default: 0
  },
  averageRating: {
    type: [Number, String],
    default: 0
  }
})
</script>

<style scoped>
.stats-row {
  width: 100%;
  justify-content: center;
  margin: 0 auto;
}

.stat-card {
  position: relative;
  border-radius: var(--radius-3xl) !important;
  background: var(--gradient-card-white) !important;
  border: 1px solid var(--border-light);
  overflow: hidden;
}

.stat-card::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  height: var(--space-1);
  background: var(--gradient-primary-purple);
  border-radius: var(--radius-3xl) 24px 0 0;
}

.stat-card:nth-child(2)::before {
  background: var(--gradient-success-dark);
}

.stat-card:nth-child(3)::before {
  background: var(--gradient-info-cyan);
}

.stat-card:nth-child(4)::before {
  background: var(--gradient-warning-dark);
}

.hover-lift {
  transition: all var(--transition-normal);
}

.hover-lift:hover {
  transform: translateY(-8px);
  box-shadow: var(--shadow-xl) !important;
}

.stat-icon {
  display: flex;
  justify-content: center;
  align-items: center;
}

.icon-glow .v-icon {
  filter: drop-shadow(0 4px 8px var(--shadow-light));
}
</style>
