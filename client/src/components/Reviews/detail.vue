<template>
  <v-infinite-scroll
    class="pt-2 pb-2"
    :items="items"
    style="-ms-overflow-style: none; scrollbar-width: none"
    @load="load"
  >
    <template v-for="item in items" :key="item">
      <v-card :elevation="0" class="pl-0">
        <v-container class="ml-2 mr-2 pl-0 pr-0">
          <v-row>
            <v-col class="v-col-auto pl-0 pr-0 pt-0">
              <v-img
                class="rounded-circle"
                width="2.8rem"
                cover
                src="https://q.trap.jp/api/v3/public/icon/howard127"
              />
            </v-col>

            <v-col class="pr-0 pl-1 pt-0 pb-0" style="max-width: 80%">
              <v-card-subtitle class="pl-0">
                <div class="d-flex flex-row flex-wrap ga-1">
                  <p class="text-text text-caption center">Display Name</p>
                  <p class="text-secondary text-caption center">25B</p>
                  <p class="text-secondary text-caption center">@username</p>
                  <p class="text-secondary text-caption center">00:00</p>
                </div>
              </v-card-subtitle>

              <div>
                <v-card-text class="pl-0 pt-1 pr-5">
                  <p class="text-wrap" style="max-width: 100%">
                    GO って書いてて楽しくない
                    AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
                  </p>
                </v-card-text>
              </div>
            </v-col>
          </v-row>
        </v-container>
      </v-card>
    </template>
  </v-infinite-scroll>
</template>

<script setup>
  import { ref } from 'vue'

  const items = ref(Array.from({ length: 30 }, (k, v) => v + 1))

  async function api() {
    return new Promise((resolve) => {
      setTimeout(() => {
        resolve(
          Array.from({ length: 10 }, (k, v) => v + items.value.at(-1) + 1)
        )
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
