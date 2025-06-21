<template>
  <v-combobox
    clear-icon="mdi-close"
    clearable
    :items="items"
    label="最寄り駅"
    menu-icon="mdi-chevron-down"
    variant="underlined"
    max-width="70%"
  >
    <template #append>
      <v-btn class="pl-0" prepend-icon="mdi-magnify">
        <p class="hidden-xs">さがす</p></v-btn
      >
    </template>
  </v-combobox>
</template>

<script setup></script>
