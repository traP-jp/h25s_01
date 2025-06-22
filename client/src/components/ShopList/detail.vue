<template>
  <v-infinite-scroll
    class="pt-2 pb-2 overflow-y-scroll d-inline-block"
    style="-ms-overflow-style: none; scrollbar-width: none"
    height="100%"
    :items="items"
    @load="load"
  >
    <template v-for="item in items" :key="item.id">
      <v-card
        class="pl-0"
        width="80vw"
        max-width="350px"
        :elevation="0"
        style="min-height: max-content"
        @click="navigateToShop(item.id)"
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
            :src="getImageUrl(item.images?.[0])"
            style="grid-row: 1 / 3; grid-column: 1 / 2; object-fit: cover"
            width="100%"
          />
          <v-img
            cover
            height="100%"
            :src="getImageUrl(item.images?.[1])"
            style="grid-row: 1 / 2; grid-column: 2 / 3; object-fit: cover"
            width="100%"
          />
          <v-img
            cover
            height="100%"
            :src="getImageUrl(item.images?.[2])"
            style="grid-row: 2 / 3; grid-column: 2 / 3; object-fit: cover"
            width="100%"
          />
        </div>

        <v-card class="pl-0">
          <v-card-title>
            <p class="color text-text text-caption left">{{ item.name }}</p>
          </v-card-title>
          <v-card-subtitle>
            <p class="color text-secondary text-caption left">
              {{ getStationNames(item.stations) }} / {{ item.address }}
            </p>
          </v-card-subtitle>
        </v-card>
      </v-card>
    </template>
  </v-infinite-scroll>
</template>

<script setup>
  import { ref, onMounted } from 'vue'
  import { useRouter } from 'vue-router'

  const router = useRouter()
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

  // 店舗詳細ページに遷移
  function navigateToShop(shopId) {
    router.push(`/shop/${shopId}`)
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
