<template>
  
<div class="overflow-hidden container max-w-md whitespace-normal break-words rounded-lg p-4 mt-10 mb-10 bg-neutral-100 shadow-md shadow-black/5 dark:bg-neutral-600 dark:shadow-black/10">    
    <h1 class="text-2xl font-semibold mb-4 text-center">Subject </h1>
              <p class="text-gray-700 mb-6 text-center">
                {{this.getCurrentMail._source.Subject}} 
              </p>
    <h1 class="text-2xl font-semibold mb-4 text-center">From </h1>
              <p class="text-gray-700 mb-6 text-center">
                {{this.getCurrentMail._source.From}}
              </p>
    <h1 class="text-2xl font-semibold mb-4 text-center">To </h1>
              <p class="text-gray-700 mb-6 text-center">
                {{this.getCurrentMail._source.To}}
              </p>
    <h1 class="text-2xl font-semibold mb-4 text-center">Date</h1>
              <p class="text-gray-700 mb-6 text-center">
                {{this.getCurrentMail._source.Date}}
              </p>
    <h1 class="text-2xl font-semibold mb-4 text-center">CC </h1>
              <p class="text-gray-700 mb-6 text-center">
                {{this.getCurrentMail._source.CC}}
              </p>
    <h1 class="text-2xl font-semibold mb-4 text-center">Body </h1>
              <p class="text-gray-700 mb-6 text-center">
                {{this.getCurrentMail._source.Body}}
              </p>
</div>
</template>

<script>
import {storeToRefs} from 'pinia'
import {useEmailStore} from '../../store/emailStore.js'

export default{
  data(){
    return{ 
      emailStore: useEmailStore.currentMail,
      getCurrentMail: storeToRefs(useEmailStore()).currentMail
    }
  },
  computed: {
    getCurrentSubject(){
      return storeToRefs(useEmailStore()).getCurrentMail
    } 
  },
}

</script>

