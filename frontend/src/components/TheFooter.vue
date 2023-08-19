<script>
import FacebookAttach from './TheFooter/FacebookAttach.vue'
import GithubAttach from './TheFooter/GitHubAttach.vue' 
import LinkedinAttach from './TheFooter/LinkedinAttach.vue'

import CopyRight  from './TheFooter/CopyRight.vue'

export default {
  components: {
    CopyRight,FacebookAttach,GithubAttach,LinkedinAttach
  },
};

</script>

<template>
<!--Footer container-->
<footer class="bg-neutral-200 text-center text-white dark:bg-neutral-600">
    <div class="text-center mb-5 mt-5 flex justify-center">
      <FacebookAttach/>
      <LinkedinAttach/>
      <GithubAttach/>
    </div>

  <!--Copyright section-->
  <div
    class="bg-neutral-300 p-4 text-center text-neutral-700 dark:bg-neutral-700 dark:text-neutral-200">
    <CopyRight/>
  </div>
</footer>
</template>

<style> 
</style>
