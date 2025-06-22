<template>
  <div id="container">
    <ShopList id="shop-list" class="mx-auto" :shops="shops" />
  </div>
</template>

<script setup>
  import { onMounted, ref } from 'vue'

  const shops = ref([])

  onMounted(async () => {
    try {
      const res = await fetch('/api/v1/shops')
      if (!res.ok) throw new Error('Failed to fetch shops')
      shops.value = await res.json()
    } catch (error) {
      console.error(error)
    }
  })
</script>

<style scoped>
  #container {
    display: flex;
    flex-direction: column;
    margin: 10px;
    height: 100%;
  }

  #shop-list {
    max-height: 100%;
  }
</style>
