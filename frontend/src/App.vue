  
<template>
  <TheHeader/> 
  <div class="min-h-screen flex flex-col">
      <RouterView/>
  </div>
  <TheFooter/>

</template>

<script setup>
  import TheHeader from './components/TheHeader.vue' 
  import TheFooter from './components/TheFooter.vue'

  import { RouterView } from 'vue-router'
</script>

<style scoped>
</style>
