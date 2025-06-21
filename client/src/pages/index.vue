<template>
  <div class="pl-2 pr-2" id="container">
    <Suggest />
    <Reviews id="reviews" />
  </div>
</template>

<script lang="ts" setup />

<style scoped>
  #container {
    display: flex;
    flex-direction: column;
    margin: 10px;
    height: 100%;
  }

  #container > :last-of-type {
    flex: 1;
    min-height: 0;
  }

  #reviews {
    max-height: 100%;
    max-width: 100%;
  }
</style>
