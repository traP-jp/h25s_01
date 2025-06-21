<template>
  <v-combobox
    clearable
    :items="items"
    label="最寄り駅"
    variant="underlined"
    width="80%"
  >
    <template #append>
      <v-btn prepend-icon="mdi-magnify">さがす</v-btn>
    </template>
  </v-combobox>
</template>

<script setup></script>
