<template>
  <v-app-bar class="pl-2 w-screen" color="primary" :elevation="0">
    <template #prepend>
      <Icon height="80%" icon="academicons:ceur-square" :ssr="true" />
    </template>

    <template #append>
      <v-app-bar-nav-icon />
    </template>

    <SearchSlot class="mt-4" :items="stations" />
  </v-app-bar>
</template>

<script setup>
  import { Icon } from '@iconify/vue'
</script>
