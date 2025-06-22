<template>
  <div id="container">
    <ShopList id="shop-list" class="mx-auto" />
  </div>
</template>

<style scoped>
  #container {
    display: flex;
    flex-direction: column;
    margin: 10px;
    height: 100%;
  }

  #shop-list {
    max-height: 100%;
  }
</style>
