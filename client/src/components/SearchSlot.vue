<template>
  <v-combobox
    clear-icon="mdi-close"
    clearable
    :items="items"
    label="最寄り駅"
    max-width="70%"
    menu-icon="mdi-chevron-down"
    variant="underlined"
  >
    <template #append>
      <v-btn prepend-icon="mdi-magnify">
        <p class="hidden-xs">さがす</p>
      </v-btn>
    </template>
  </v-combobox>
</template>

<script setup></script>
