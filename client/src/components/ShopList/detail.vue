<template>
  <v-infinite-scroll
    class="pt-2 pb-2 overflow-y-scroll"
    style="-ms-overflow-style: none; scrollbar-width: none"
    height="100%"
    :items="items"
    @load="load"
  >
    <template v-for="item in items" :key="item">
      <v-card
        class="pl-0"
        width="80vw"
        max-width="350px"
        :elevation="0"
        style="min-height: max-content"
      >
        <div
          style="
            display: grid;
            grid-template-columns: 2fr 1fr;
            grid-template-rows: 1fr 1fr;
            height: 170px;
            gap: 2px;
          "
        >
          <v-img
            cover
            height="100%"
            src="https://cdn.vuetifyjs.com/images/cards/house.jpg"
            style="grid-row: 1 / 3; grid-column: 1 / 2; object-fit: cover"
            width="100%"
          />
          <v-img
            cover
            height="100%"
            src="https://cdn.vuetifyjs.com/images/cards/house.jpg"
            style="grid-row: 1 / 2; grid-column: 2 / 3; object-fit: cover"
            width="100%"
          />
          <v-img
            cover
            height="100%"
            src="https://cdn.vuetifyjs.com/images/cards/house.jpg"
            style="grid-row: 2 / 3; grid-column: 2 / 3; object-fit: cover"
            width="100%"
          />
        </div>

        <v-card class="pl-0">
          <v-card-title>
            <p class="color text-text text-caption left">お店の名前</p>
          </v-card-title>
          <v-card-subtitle>
            <p class="color text-secondary text-caption left">
              ○○駅周辺 / text text / text text
            </p>
          </v-card-subtitle>
        </v-card>
      </v-card>
    </template>
  </v-infinite-scroll>
</template>

<script setup>
  import { ref } from 'vue'

  const items = ref(Array.from({ length: 2 }, (k, v) => v + 1))

  async function api() {
    return new Promise((resolve) => {
      setTimeout(() => {
        resolve(Array.from({ length: 5 }, (k, v) => v + items.value.at(-1) + 1))
      }, 1000)
    })
  }

  async function load({ done }) {
    // Perform API call
    const res = await api()
    items.value.push(...res)
    done('ok')
  }
</script>
