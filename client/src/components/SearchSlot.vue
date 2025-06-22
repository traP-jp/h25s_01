<template>
  <v-combobox
    clear-icon="mdi-close"
    clearable
    class="ml-3"
    :items="items"
    label="最寄り駅を入力"
    max-width="70%"
    menu-icon="mdi-chevron-down"
    variant="underlined"
  >
    <template #append>
      <v-btn
        class="hidden-sm-and-up"
        density="default"
        icon="mdi-magnify"
        size="md"
      />

      <v-btn
        class="hidden-xs"
        density="default"
        prepend-icon="mdi-magnify"
        size="md"
      >
        さがす
      </v-btn>
    </template>
  </v-combobox>
</template>

<script setup></script>
