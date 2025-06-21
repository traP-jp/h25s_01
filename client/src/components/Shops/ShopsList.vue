<template>
  <v-infinite-scroll
    class="pt-2 pb-2"
    :items="items"
    style="-ms-overflow-style: none; scrollbar-width: none"
    @load="load"
  >
    <template
      v-for="item in items" :key="item"
    >
      <div
        style="
          width: 100%;
          max-width: 351px;
          border-radius: 6px;
          overflow: hidden;
          margin: 0 auto 16px auto;
          box-shadow: 0 2px 8px rgba(0,0,0,0.04);
        "
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
            style="grid-row: 1 / 3; grid-column: 1 / 2; object-fit: cover;"
            src="https://cdn.vuetifyjs.com/images/cards/house.jpg"
            width="100%"
            height="100%"
            cover
          ></v-img>
          <v-img
            style="grid-row: 1 / 2; grid-column: 2 / 3; object-fit: cover;"
            src="https://cdn.vuetifyjs.com/images/cards/house.jpg"
            width="100%"
            height="100%"
            cover
          ></v-img>
          <v-img
            style="grid-row: 2 / 3; grid-column: 2 / 3; object-fit: cover;"
            src="https://cdn.vuetifyjs.com/images/cards/house.jpg"
            width="100%"
            height="100%"
            cover
          ></v-img>
        </div>

        <v-card class="pl-0">
          <v-card-title>
            <p class="color text-text text-caption left">お店の名前</p>
          </v-card-title>
          <v-card-subtitle>
            <p class="color text-secondary text-caption left">○○駅周辺 / text text / text text</p>
          </v-card-subtitle>
        </v-card>
      </div>
    </template>
  </v-infinite-scroll>
</template>

<script setup>
import { ref } from 'vue'

const items = ref(Array.from({ length: 30 }, (k, v) => v + 1))

async function api () {
  return new Promise(resolve => {
    setTimeout(() => {
      resolve(Array.from({ length: 10 }, (k, v) => v + items.value.at(-1) + 1))
    }, 1000)
  })
}

async function load ({ done }) {
  // Perform API call
  const res = await api()
  items.value.push(...res)
  done('ok')
}
</script>

