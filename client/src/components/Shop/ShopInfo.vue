<template>
  <v-card class="ml-10 mr-10 text-center" :elevation="0">
    <v-card-title>
      <p>お店の名前</p>
    </v-card-title>
    <v-card-subtitle>
      <p>○○駅周辺 / text text / text text</p>
    </v-card-subtitle>

    <v-container class="d-flex flex-row justify-center">
      <v-img
        cover
        gradient="to bottom, rgba(0,0,0,0.1), rgba(0,0,0,.8)"
        max-height="20rem "
        max-width="30rem"
        rounded="lg"
        src="https://cdn.vuetifyjs.com/images/cards/house.jpg"
      />
    </v-container>
  </v-card>
</template>
