<script>
  import { RouterLink } from 'vue-router'

  export default{
    components: {
      RouterLink
    },
  }
</script>

<template>
  <ul
    class="list-style-none ml-12 flex flex-col pl-0 lg:flex-row"
    data-te-navbar-nav-ref>
    <li class="mb-4 lg:mb-0 lg:pr-2" data-te-nav-item-ref>
      <RouterLink
        class="text-neutral-500 hover:text-neutral-700 focus:text-neutral-700 disabled:text-black/30 dark:text-neutral-200 dark:hover:text-neutral-300 dark:focus:text-neutral-300 lg:px-2 [&.active]:text-black/90 dark:[&.active]:text-zinc-400"
        to="/"
        data-te-nav-link-ref
        >Emails</RouterLink>
    </li>
    <!-- Team link -->
    <li class="mb-4 lg:mb-0 lg:pr-2" data-te-nav-item-ref>
      <RouterLink
        class="text-neutral-500 hover:text-neutral-700 focus:text-neutral-700 disabled:text-black/30 dark:text-neutral-200 dark:hover:text-neutral-300 dark:focus:text-neutral-300 lg:px-2 [&.active]:text-black/90 dark:[&.active]:text-neutral-400"
        to="/about"
        data-te-nav-link-ref
        >About</RouterLink>
    </li>
  </ul>
</template>

<style> 
</style>
