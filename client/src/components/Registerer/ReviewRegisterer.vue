<template>
  <v-dialog max-width="500">
    <template #activator="{ props: activatorProps }">
      <OpenButton v-bind="activatorProps" />
    </template>

    <template #default="{ isActive }">
      <v-card class="pt-1">
        <v-card-title class="pb-0 pl-3 pt-2 position-relative">
          <span class="text-primary" style="width: fit-content">
            口コミを投稿する
          </span>

          <v-fab
            absolute
            class="mt-1 mr-2"
            color="primaryText"
            :elevation="0"
            icon
            location="right top"
            size="tiny"
            @click="isActive.value = false"
          >
            <Icon
              class="text-warning"
              height="1rem"
              icon="akar-icons:cross"
              style="margin: 4px"
            />
          </v-fab>
        </v-card-title>

        <v-divider
          class="border-opacity-100 rounded-pill"
          color="primary"
          length="14rem"
          :thickness="3"
        />

        <v-card-text class="pb-0 mt-2">
          <form @submit.prevent="submit">
            <v-combobox
              color="secondary"
              :items="shops"
              label="お店の名前"
              variant="outlined"
            />

            <div class="d-flex flex-col ga-4">
              <v-checkbox v-model="doRate" label="評価を付ける" />

              <v-slider
                :min="1"
                :max="3"
                v-model="rating"
                :ticks="tickLabels"
                thumb-label="always"
                show-ticks="always"
                step="1"
                :disabled="!doRate"
                tick-size="2"
              >
                <template #thumb-label="{ modelValue }">
                  {{ satisfactionEmojis[modelValue - 1] }}
                </template>
              </v-slider>
            </div>

            <v-textarea color="secondary" label="投稿内容" variant="outlined" />

            <v-file-upload
              v-model="model"
              class="mb-3"
              clearable
              color="text-color"
              density="compact"
              height="4rem"
              multiple
              show-size
              variant="compact"
            >
              <template #title>
                <p class="text-secondaryText">画像を追加</p></template
              >

              <template #item="{ props: files }">
                <v-file-upload-item
                  density="compact"
                  v-bind="files"
                  lines="one"
                  nav
                >
                  <template #prepend>
                    <v-avatar rounded size="32" />
                  </template>

                  <template #clear="{ props: clear }">
                    <v-btn color="primary" v-bind="clear" />
                  </template>
                </v-file-upload-item>
              </template>
            </v-file-upload>
          </form>
        </v-card-text>

        <v-card-actions class="mb-2">
          <v-spacer />
          <v-btn class="me-4 bg-primary" rounded type="submit" width="8rem">
            投稿する
          </v-btn>
          <v-spacer />
        </v-card-actions>
      </v-card>
    </template>
  </v-dialog>
</template>

<script setup>
  import { Icon } from '@iconify/vue'
  import { ref, shallowRef } from 'vue'

  import { VFileUpload, VFileUploadItem } from 'vuetify/labs/VFileUpload'
  const model = shallowRef(null)

  const rating = ref(2)
  const doRate = ref(false)

  const satisfactionEmojis = ['\uD83D\uDE22', '\uD83D\uDE42', '\uD83D\uDE0A']
</script>
