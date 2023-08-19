<script>
// Initialization for ES Users
import { Collapse, Dropdown, initTE } from "tw-elements"; 

import HamburguerButton from './TheHeader/HamburguerButton.vue'
import AppLogo from './TheHeader/AppLogo.vue'
import ViewChangue from './TheHeader/ViewChangue.vue'

initTE({ Collapse, Dropdown });

export default {
  components: {
    HamburguerButton,AppLogo,ViewChangue
  },
  data(){
    return{
      isMenuOpen: false,
    }
  },
  methods:{
    toggleMenu(){
      this.isMenuOpen=!this.isMenuOpen 
    }

  }
};

</script>

<template>
  <!-- Main navigation container -->
  <nav
    class="flex-no-wrap relative flex w-full items-center justify-between bg-neutral-100 py-2 shadow-md shadow-black/5 dark:bg-neutral-600 dark:shadow-black/10 lg:flex-wrap lg:justify-start lg:py-4"
    data-te-navbar-ref>
    <div class="flex w-full flex-wrap items-center justify-between px-3">
      <!-- Hamburguer Button for mobile view -->
      <HamburguerButton/>
      <!--  navigation container -->
      <div
        class="!visible hidden flex-grow basis-[100%] items-center lg:!flex lg:basis-auto"
        id="navbarSupportedContent1"
        data-te-collapse-item>
        <AppLogo/>
        <ViewChangue/>
      </div>
    </div>
  </nav>
</template>
<style>
</style>
