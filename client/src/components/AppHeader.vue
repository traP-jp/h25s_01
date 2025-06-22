<template>
  <v-app-bar class="pl-2 w-screen" color="primary" :elevation="0">
    <template #prepend>
      <v-btn class="hidden-sm-and-up" icon href="/">
        <v-img height="2.5rem" rel="icon" src="@/assets/icon.svg" width="2.5rem" />
      </v-btn>

      <v-btn class="hidden-xs" href="/">
        <v-img height="2.5rem" rel="icon" src="@/assets/icon.svg" width="2.5rem" />
        <v-img height="2.5rem" rel="icon" src="@/assets/logo.png" width="8rem" />
      </v-btn>
    </template>

    <template #append>
      <v-app-bar-nav-icon variant="text" @click.stop="drawer = !drawer" />
    </template>

    <SearchSlot class="mt-4" :items="stations" />
  </v-app-bar>

  <v-navigation-drawer v-model="drawer" location="right" temporary>
    <v-list-item link title="Top" href="/" />
    <v-list-item link title="Timeline" href="/timeline" />
    <v-list-item link title="Shops" href="/shop" />
  </v-navigation-drawer>
</template>

<script setup>
  import { Icon } from '@iconify/vue'
  import { ref, watch } from 'vue'

  const drawer = ref(false)
  const group = ref(null)

  watch(group, () => {
    drawer.value = false
  })
</script>
