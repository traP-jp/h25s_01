<template>
  <v-infinite-scroll
    class="pt-2 pb-2"
    direction="horizontal"
    :items="items"
    style="-ms-overflow-style: none; scrollbar-width: none"
    @load="load"
  >
    <template v-for="item in items" :key="item">
      <div class="pr-1">
        <v-card class="mx-auto" color="surface-variant" max-width="340">
          <v-img
            class="align-end"
            cover
            gradient="to bottom, rgba(0,0,0,0.1), rgba(0,0,0,.8)"
            height="8.4rem"
            src="https://cdn.vuetifyjs.com/images/cards/house.jpg"
            width="14.4rem"
          >
            <v-card-title class="primary-text pb-0"> Shop Name </v-card-title>
            <v-card-subtitle
              class="secondary-text pt-0 pb-2"
              style="line-height: 1"
            >
              ○○駅周辺 / text text / text text
            </v-card-subtitle>
          </v-img>
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
        resolve(
          Array.from({ length: 10 }, (k, v) => v + items.value.at(-1) + 1),
        )
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
