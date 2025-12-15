<template>
  <v-row class="mb-6 stats-row">
    <v-col cols="12" md="3">
      <v-card class="stat-card pa-6 text-center hover-lift card-glow smooth-transition" elevation="2">
        <div class="stat-icon mb-3 icon-glow">
          <v-icon size="64" color="primary">mdi-account-group</v-icon>
        </div>
        <h3 class="text-h3 font-weight-bold text-primary mb-2" style="font-family: 'Arial', 'Helvetica', sans-serif; direction: ltr;">{{ totalEngineers }}</h3>
        <p class="text-subtitle-1 text-primary mb-0">إجمالي المهندسين</p>
      </v-card>
    </v-col>
    <v-col cols="12" md="3">
      <v-card class="stat-card pa-6 text-center hover-lift card-glow smooth-transition" elevation="2">
        <div class="stat-icon mb-3 icon-glow">
          <v-icon size="64" color="success">mdi-check-circle</v-icon>
        </div>
        <h3 class="text-h3 font-weight-bold text-success mb-2" style="font-family: 'Arial', 'Helvetica', sans-serif; direction: ltr;">{{ activeEngineers }}</h3>
        <p class="text-subtitle-1 text-success mb-0">مهندسين نشطين</p>
      </v-card>
    </v-col>
    <v-col cols="12" md="3">
      <v-card class="stat-card pa-6 text-center hover-lift card-glow smooth-transition" elevation="2">
        <div class="stat-icon mb-3 icon-glow">
          <v-icon size="64" color="info">mdi-briefcase</v-icon>
        </div>
        <h3 class="text-h3 font-weight-bold text-info mb-2" style="font-family: 'Arial', 'Helvetica', sans-serif; direction: ltr;">{{ totalProjects }}</h3>
        <p class="text-subtitle-1 text-info mb-0">إجمالي المشاريع</p>
      </v-card>
    </v-col>
    <v-col cols="12" md="3">
      <v-card class="stat-card pa-6 text-center hover-lift card-glow smooth-transition" elevation="2">
        <div class="stat-icon mb-3 icon-glow">
          <v-icon size="64" color="warning">mdi-star</v-icon>
        </div>
        <h3 class="text-h3 font-weight-bold text-warning mb-2" style="font-family: 'Arial', 'Helvetica', sans-serif; direction: ltr;">{{ averageRating }}</h3>
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

<style>
@import './styles/engineers.css';
</style>

<style scoped>
.stats-row {
  width: 100%;
  justify-content: center;
  margin: 0 auto;
}

.stat-card {
  position: relative;
  border-radius: 24px !important;
  background: linear-gradient(145deg, #ffffff 0%, #f8fafc 100%) !important;
  border: 1px solid rgba(0, 0, 0, 0.05);
  overflow: hidden;
}

.stat-card::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  height: 4px;
  background: linear-gradient(90deg, #3b82f6, #8b5cf6);
  border-radius: 24px 24px 0 0;
}

.stat-card:nth-child(2)::before {
  background: linear-gradient(90deg, #10b981, #059669);
}

.stat-card:nth-child(3)::before {
  background: linear-gradient(90deg, #3b82f6, #0ea5e9);
}

.stat-card:nth-child(4)::before {
  background: linear-gradient(90deg, #f59e0b, #d97706);
}

.hover-lift {
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
}

.hover-lift:hover {
  transform: translateY(-8px);
  box-shadow: 0 20px 40px rgba(0, 0, 0, 0.12) !important;
}

.stat-icon {
  display: flex;
  justify-content: center;
  align-items: center;
}

.icon-glow .v-icon {
  filter: drop-shadow(0 4px 8px rgba(0, 0, 0, 0.1));
}
</style>
