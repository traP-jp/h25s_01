<template>
  <div class="pl-2">
    <v-container>
      <p class="text-primary">Shops</p>
      <v-divider
        class="border-opacity-100 rounded-pill"
        color="primary"
        length="7rem"
        :thickness="3"
      />

      <div class="pl-4">
        <ShopsList class="pt-4" />
      </div>
    </v-container>
  </div>
</template>

<script setup>
  import ShopsList from './ShopsList.vue'
</script>

<style lang="sass"></style>
