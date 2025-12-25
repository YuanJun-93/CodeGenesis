<template>
  <a-layout class="basic-layout">
    <!-- Top Navigation -->
    <GlobalHeader />
    <!-- Main Content -->
    <a-layout-content class="main-content">
      <router-view />
    </a-layout-content>
    <!-- Footer Copyright -->
    <GlobalFooter />
  </a-layout>
</template>

<script setup lang="ts">
import GlobalHeader from '@/components/GlobalHeader.vue'
import GlobalFooter from '@/components/GlobalFooter.vue'
</script>

<style scoped>
.basic-layout {
  background: none;
}

.main-content {
  max-width: 1200px;
  width: 100%;
  padding: 24px;
  background: white;
  margin: 16px auto 56px;
}
</style>
