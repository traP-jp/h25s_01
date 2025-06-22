<template>
  <v-card class="ml-10 mr-10 text-center" :elevation="0">
    <v-card-title>
      <p>{{ shop.name || 'お店の名前' }}</p>
    </v-card-title>
    <v-card-subtitle>
      <p>{{ getShopDescription() }}</p>
    </v-card-subtitle>

    <v-container class="d-flex flex-row justify-center">
      <v-img
        cover
        gradient="to bottom, rgba(0,0,0,0.1), rgba(0,0,0,.8)"
        max-height="20rem"
        max-width="30rem"
        rounded="lg"
        :src="getMainImageUrl()"
      />
    </v-container>

    <!-- 店舗詳細情報 -->
    <v-container v-if="shop.id" class="mt-4">
      <v-row>
        <v-col cols="12" md="6">
          <v-card variant="outlined">
            <v-card-title class="text-h6">基本情報</v-card-title>
            <v-card-text>
              <div class="mb-2">
                <strong>住所:</strong> {{ shop.address || '情報なし' }}
              </div>
              <div class="mb-2">
                <strong>郵便番号:</strong> {{ shop.post_code || '情報なし' }}
              </div>
              <div class="mb-2">
                <strong>登録者:</strong> {{ shop.registerer || '情報なし' }}
              </div>
            </v-card-text>
          </v-card>
        </v-col>
        <v-col cols="12" md="6">
          <v-card variant="outlined">
            <v-card-title class="text-h6">支払い方法</v-card-title>
            <v-card-text>
              <div
                v-if="shop.payment_methods && shop.payment_methods.length > 0"
              >
                <v-chip
                  v-for="method in shop.payment_methods"
                  :key="method"
                  class="ma-1"
                  size="small"
                  color="primary"
                  variant="outlined"
                >
                  {{ method }}
                </v-chip>
              </div>
              <div v-else class="text-grey">支払い方法の情報がありません</div>
            </v-card-text>
          </v-card>
        </v-col>
      </v-row>

      <!-- 画像ギャラリー -->
      <v-row v-if="shop.images && shop.images.length > 1" class="mt-4">
        <v-col cols="12">
          <v-card variant="outlined">
            <v-card-title class="text-h6">画像ギャラリー</v-card-title>
            <v-card-text>
              <div class="d-flex flex-wrap ga-2">
                <v-img
                  v-for="imageId in shop.images"
                  :key="imageId"
                  :src="getImageUrl(imageId)"
                  cover
                  rounded="lg"
                  max-width="150"
                  max-height="150"
                  class="cursor-pointer"
                  @click="openImageModal(imageId)"
                />
              </div>
            </v-card-text>
          </v-card>
        </v-col>
      </v-row>
    </v-container>

    <!-- ローディング状態 -->
    <v-container v-else class="text-center">
      <v-progress-circular indeterminate color="primary" />
      <p class="mt-2">店舗情報を読み込み中...</p>
    </v-container>
  </v-card>
</template>

<script setup>
  import { ref, onMounted } from 'vue'
  import { useRoute } from 'vue-router'

  const route = useRoute()
  const shop = ref({})
  const stations = ref([])

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
  async function fetchShop(shopId) {
    try {
      const res = await fetch(`/api/v1/shops/${shopId}`)
      if (!res.ok) throw new Error('Failed to fetch shop')
      shop.value = await res.json()
    } catch (error) {
      console.error('Error fetching shop:', error)
    }
  }

  // 店舗の説明文を生成
  function getShopDescription() {
    if (!shop.value.id) return '○○駅周辺 / text text / text text'

    const stationNames = getStationNames(shop.value.stations)
    const address = shop.value.address || 'text text'
    return `${stationNames} / ${address}`
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

  // メイン画像URLを取得
  function getMainImageUrl() {
    if (!shop.value.images || shop.value.images.length === 0) {
      return 'https://cdn.vuetifyjs.com/images/cards/house.jpg'
    }
    return `/api/v1/images/${shop.value.images[0]}`
  }

  // 画像URLを取得
  function getImageUrl(imageId) {
    if (!imageId) return ''
    return `/api/v1/images/${imageId}`
  }

  // 画像モーダルを開く（将来的な機能拡張用）
  function openImageModal(imageId) {
    console.log('Image clicked:', imageId)
  }

  onMounted(async () => {
    await fetchStations()

    // URLパラメータから店舗IDを取得
    const shopId = route.params.id
    if (shopId) {
      await fetchShop(shopId)
    }
  })
</script>
