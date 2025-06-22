<template>
  <v-infinite-scroll
    class="pt-2 pb-2"
    direction="horizontal"
    :items="items"
    style="-ms-overflow-style: none; scrollbar-width: none"
    @load="load"
  >
    <template v-for="item in items" :key="item.id">
      <div class="pr-1">
        <v-card class="mx-auto" color="surface-variant" max-width="340">
          <v-img
            class="align-end"
            cover
            gradient="to bottom, rgba(0,0,0,0.1), rgba(0,0,0,.8)"
            height="8.4rem"
            :src="getImageUrl(item.images?.[0])"
            width="14.4rem"
          >
            <v-card-title class="primary-text pb-0">{{
              item.name
            }}</v-card-title>
            <v-card-subtitle
              class="secondary-text pt-0 pb-2"
              style="line-height: 1"
            >
              {{ getStationNames(item.stations) }} / {{ item.address }}
            </v-card-subtitle>
          </v-img>
        </v-card>
      </div>
    </template>
  </v-infinite-scroll>
</template>

<script setup>
  import { ref, onMounted } from 'vue'

  const items = ref([])
  const stations = ref([])
  const allShops = ref([])
  const currentIndex = ref(0)
  const pageSize = 10

  // 駅データを取得
  async function fetchStations() {
    try {
      const res = await fetch('/api/v1/stations')
      if (!res.ok) throw new Error('Failed to fetch stations')
      stations.value = await res.json()
    } catch (error) {
      console.error('Error fetching stations:', error)
    }
  }

  // 店舗データを取得
  async function fetchShops() {
    try {
      const res = await fetch('/api/v1/shops')
      if (!res.ok) throw new Error('Failed to fetch shops')
      allShops.value = await res.json()

      // 初期表示分をitemsに設定
      items.value = allShops.value.slice(0, pageSize)
      currentIndex.value = pageSize
    } catch (error) {
      console.error('Error fetching shops:', error)
    }
  }

  // 画像URLを取得
  function getImageUrl(imageId) {
    if (!imageId) {
      return 'https://cdn.vuetifyjs.com/images/cards/house.jpg'
    }
    return `/api/v1/images/${imageId}`
  }

  // 駅名を取得
  function getStationNames(stationIds) {
    if (!stationIds || stationIds.length === 0) {
      return '○○駅周辺'
    }

    const stationNames = stationIds
      .map((id) => {
        const station = stations.value.find((s) => s.id === id)
        return station ? station.name : ''
      })
      .filter((name) => name)

    return stationNames.length > 0
      ? `${stationNames.join(', ')}周辺`
      : '○○駅周辺'
  }

  async function load({ done }) {
    // 次のページのデータを取得
    const nextItems = allShops.value.slice(
      currentIndex.value,
      currentIndex.value + pageSize
    )

    if (nextItems.length > 0) {
      items.value.push(...nextItems)
      currentIndex.value += pageSize
    }

    done('ok')
  }

  onMounted(async () => {
    await fetchStations()
    await fetchShops()
  })
</script>
