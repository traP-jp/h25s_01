<template>
  <v-app>
    <header>
      <AppHeader />
    </header>

    <main>
      <router-view />
    </main>
  </v-app>
</template>

<script lang="ts" setup>
  import { RouterView } from 'vue-router'
</script>
