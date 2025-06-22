<template>
  <v-container class="mt-2">
    <v-row>
      <v-col cols="12">
        <h1 class="text-h4 mb-4">店舗一覧</h1>
        <v-divider class="mb-4" />
      </v-col>
    </v-row>

    <v-row>
      <v-col
        v-for="shop in shops"
        :key="shop.id"
        cols="12"
        sm="6"
        md="4"
        lg="3"
      >
        <v-card
          class="mb-4"
          :elevation="2"
          @click="navigateToShop(shop.id)"
          style="cursor: pointer; height: 100%"
        >
          <v-img
            :src="getImageUrl(shop.images?.[0])"
            height="200"
            cover
            gradient="to bottom, rgba(0,0,0,0.1), rgba(0,0,0,0.4)"
          >
            <template v-slot:placeholder>
              <v-row class="fill-height ma-0" align="center" justify="center">
                <v-progress-circular indeterminate color="grey lighten-5" />
              </v-row>
            </template>
          </v-img>

          <v-card-title class="text-h6">
            {{ shop.name }}
          </v-card-title>

          <v-card-subtitle>
            {{ getStationNames(shop.stations) }}
          </v-card-subtitle>

          <v-card-text>
            <div class="text-body-2 text-grey-darken-1">
              {{ shop.address || '住所情報なし' }}
            </div>

            <!-- 支払い方法 -->
            <div
              v-if="shop.payment_methods && shop.payment_methods.length > 0"
              class="mt-2"
            >
              <v-chip
                v-for="method in shop.payment_methods.slice(0, 3)"
                :key="method"
                size="x-small"
                class="me-1 mb-1"
                color="primary"
                variant="outlined"
              >
                {{ method }}
              </v-chip>
              <v-chip
                v-if="shop.payment_methods.length > 3"
                size="x-small"
                color="grey"
                variant="outlined"
              >
                +{{ shop.payment_methods.length - 3 }}
              </v-chip>
            </div>
          </v-card-text>

          <v-card-actions>
            <v-btn
              variant="text"
              color="primary"
              @click.stop="navigateToShop(shop.id)"
            >
              詳細を見る
            </v-btn>
          </v-card-actions>
        </v-card>
      </v-col>
    </v-row>

    <!-- ローディング状態 -->
    <v-row v-if="loading">
      <v-col cols="12" class="text-center">
        <v-progress-circular indeterminate color="primary" />
        <p class="mt-2">店舗情報を読み込み中...</p>
      </v-col>
    </v-row>

    <!-- データがない場合 -->
    <v-row v-if="!loading && shops.length === 0">
      <v-col cols="12" class="text-center">
        <v-icon size="64" color="grey">mdi-store-off</v-icon>
        <p class="text-h6 mt-2">店舗が見つかりませんでした</p>
      </v-col>
    </v-row>
  </v-container>
</template>

<script setup>
  import { ref, onMounted } from 'vue'
  import { useRouter } from 'vue-router'

  const router = useRouter()
  const shops = ref([])
  const stations = ref([])
  const loading = ref(true)

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
      shops.value = await res.json()
    } catch (error) {
      console.error('Error fetching shops:', error)
    } finally {
      loading.value = false
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

  onMounted(async () => {
    await fetchStations()
    await fetchShops()
  })
</script>

<style scoped>
  .v-card {
    transition: transform 0.2s ease-in-out;
  }

  .v-card:hover {
    transform: translateY(-2px);
  }
</style>
