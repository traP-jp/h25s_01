<template>
  <div class="fab-container">
    <v-fab
      absolute
      app
      class="w-100"
      color="primary"
      :elevation="2"
      icon
      location="right bottom"
      size="large"
    >
      <v-icon color="primaryText" icon="mdi-plus" />
    </v-fab>
  </div>
</template>

<style scoped>
  .fab-container {
    /* position: relative; */
    margin: 0 30px 50px;
  }
</style>
