<template>
  <v-infinite-scroll
    class="pt-2 pb-2 overflow-y-scroll d-inline-block"
    style="-ms-overflow-style: none; scrollbar-width: none"
    height="100%"
    width="100%"
    :items="items"
    @load="load"
  >
    <template v-for="item in items" :key="item.id">
      <v-card class="pl-0" :elevation="0" style="min-height: max-content">
        <v-container class="ml-2 mr-2 pl-0 pr-0">
          <v-row>
            <v-col class="v-col-auto pl-0 pr-0 pt-0">
              <v-img
                class="rounded-circle"
                cover
                :src="getUserIconUrl(item.author)"
                width="2.8rem"
              />
            </v-col>

            <v-col class="pr-0 pl-1 pt-0 pb-0" style="max-width: 80%">
              <v-card-subtitle class="pl-0">
                <div class="d-flex flex-row flex-wrap ga-1">
                  <p class="text-text text-caption center">{{ item.author }}</p>
                  <p class="text-secondary text-caption center">
                    {{ getRatingText(item.rating) }}
                  </p>
                  <p class="text-secondary text-caption center">
                    @{{ item.author }}
                  </p>
                  <p class="text-secondary text-caption center">
                    {{ formatTime(item.created_at) }}
                  </p>
                </div>
              </v-card-subtitle>

              <div>
                <v-card-text class="pl-0 pt-1 pr-5">
                  <p class="text-wrap" style="max-width: 100%">
                    {{ item.content || 'レビュー内容がありません' }}
                  </p>

                  <!-- レビュー画像表示 -->
                  <div
                    v-if="item.images && item.images.length > 0"
                    class="mt-2"
                  >
                    <div class="d-flex flex-wrap ga-2">
                      <v-img
                        v-for="imageId in item.images"
                        :key="imageId"
                        :src="getImageUrl(imageId)"
                        cover
                        rounded="lg"
                        max-width="120"
                        max-height="120"
                        class="cursor-pointer"
                        @click="openImageModal(imageId)"
                      />
                    </div>
                  </div>

                  <div v-if="getShopName(item.shop)" class="mt-2">
                    <v-chip size="small" color="primary" variant="outlined">
                      {{ getShopName(item.shop) }}
                    </v-chip>
                  </div>
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
  import { ref, onMounted } from 'vue'

  const items = ref([])
  const shops = ref([])
  const allReviews = ref([])
  const currentIndex = ref(0)
  const pageSize = 10

  // 店舗データを取得
  async function fetchShops() {
    try {
      const res = await fetch('/api/v1/shops')
      if (!res.ok) throw new Error('Failed to fetch shops')
      shops.value = await res.json()
    } catch (error) {
      console.error('Error fetching shops:', error)
    }
  }

  // レビューデータを取得
  async function fetchReviews() {
    try {
      const res = await fetch('/api/v1/reviews')
      if (!res.ok) throw new Error('Failed to fetch reviews')
      allReviews.value = await res.json()

      // 作成日時でソート（新しい順）
      allReviews.value.sort(
        (a, b) => new Date(b.created_at) - new Date(a.created_at)
      )

      // 初期表示分をitemsに設定
      items.value = allReviews.value.slice(0, pageSize)
      currentIndex.value = pageSize
    } catch (error) {
      console.error('Error fetching reviews:', error)
    }
  }

  // ユーザーアイコンURLを取得
  function getUserIconUrl(author) {
    return `https://q.trap.jp/api/v3/public/icon/${author}`
  }

  // 評価をテキストに変換
  function getRatingText(rating) {
    const ratingMap = {
      0: '',
      1: '★☆☆',
      2: '★★☆',
      3: '★★★',
    }
    return ratingMap[rating] || '☆☆☆☆'
  }

  // 時間をフォーマット
  function formatTime(dateString) {
    if (!dateString) return '00:00'

    const date = new Date(dateString)
    const now = new Date()
    const diffInMinutes = Math.floor((now - date) / (1000 * 60))

    if (diffInMinutes < 1) return '今'
    if (diffInMinutes < 60) return `${diffInMinutes}分前`
    if (diffInMinutes < 1440) return `${Math.floor(diffInMinutes / 60)}時間前`
    return `${Math.floor(diffInMinutes / 1440)}日前`
  }

  // 店舗名を取得
  function getShopName(shopId) {
    const shop = shops.value.find((s) => s.id === shopId)
    return shop ? shop.name : ''
  }

  // 画像URLを取得
  function getImageUrl(imageId) {
    if (!imageId) return ''
    return `/api/v1/images/${imageId}`
  }

  // 画像モーダルを開く（将来的な機能拡張用）
  function openImageModal(imageId) {
    // 画像の拡大表示などの機能を実装する場合はここに追加
    console.log('Image clicked:', imageId)
  }

  async function load({ done }) {
    // 次のページのデータを取得
    const nextItems = allReviews.value.slice(
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
    await fetchShops()
    await fetchReviews()
  })
</script>
