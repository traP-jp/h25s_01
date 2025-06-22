<template>
  <div class="pl-4 pr-4" style="max-width: 500px">
    <p class="text-primary">Timeline</p>

    <v-divider
      class="border-opacity-100 rounded-pill"
      color="primary"
      length="7rem"
      :thickness="3"
    />

    <Detail class="pt-4 ml-3 mr-3" />
  </div>
</template>

<script setup>
  import Detail from './detail.vue'
</script>
