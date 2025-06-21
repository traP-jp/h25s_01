<template>
  <v-app-bar class="pl-2" color="primary" :elevation="0" height="70">
    <template #prepend>
      <Icon height="80%" icon="academicons:ceur-square" :ssr="true" />
    </template>

    <template #append>
      <v-app-bar-nav-icon />
    </template>

    <v-container>
      <v-row>
        <v-col class="d-flex flex-row justify-center" cols="0">
          <v-spacer />
          <SearchSlot class="mt-4" :items="['test']" />
          <v-spacer />
        </v-col>
      </v-row>
    </v-container>
  </v-app-bar>
</template>
test
<script setup>
  import { Icon } from '@iconify/vue'
</script>
