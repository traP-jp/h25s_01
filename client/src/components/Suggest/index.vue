<template>
  <div>
    <v-container>
      <p class="text-primary">Suggest</p>
      <v-divider
        class="border-opacity-100 rounded-pill"
        color="primary"
        length="7rem"
        :thickness="3"
      />

      <div class="pl-4 pr-4">
        <Showcase class="pt-4" />
      </div>
    </v-container>
  </div>
</template>

<script setup>
  import Showcase from './detail.vue'
</script>

<style lang="sass"></style>
