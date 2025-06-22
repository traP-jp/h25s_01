<template>
  <v-app-bar
    class="pl-2 w-screen"
    color="primary"
    :elevation="0"
    :style="{
      minWidth: '0',
      width: '100vw',
      boxSizing: 'border-box',
      paddingLeft: isMobile ? '4px' : '8px',
      paddingRight: isMobile ? '4px' : '8px',
      height: isMobile ? '56px' : '64px',
    }"
  >
    <template #prepend>
      <Icon
        icon="academicons:ceur-square"
        :ssr="true"
        :style="{
          width: isMobile ? '32px' : '40px',
          height: isMobile ? '32px' : '40px'
        }"
      />
    </template>

    <template #append>
      <v-app-bar-nav-icon />
    </template>

    <SearchSlot
      class="mt-4"
      :items="stations"
      :style="{
        marginTop: isMobile ? '8px' : '16px'
      }"
    />
  </v-app-bar>
</template>

<script setup>
  import { Icon } from '@iconify/vue'
  import { onBeforeUnmount, onMounted, ref } from 'vue'

  const isMobile = ref(window.innerWidth < 600)

  function handleResize () {
    isMobile.value = window.innerWidth < 600
  }

  onMounted(() => {
    window.addEventListener('resize', handleResize)
  })

  onBeforeUnmount(() => {
    window.removeEventListener('resize', handleResize)
  })
</script>
