<template>
<div class="overflow-hidden container max-w-md whitespace-normal break-words rounded-lg p-4 mt-10 mb-10 bg-neutral-100 shadow-md shadow-black/5 dark:bg-neutral-600 dark:shadow-black/10">    
    <div class="py-6">
      <div >
        <div v-for="email in emails" :key="email.id" class="justify-items-center items-center overflow-hidden w-full" >
          <button @click="this.updateMail(email)" class="bg-white rounded-lg shadow-md pl-24 pr-24  mb-2 hover:bg-blue-400 whitespace-normal break-words w-full">
            <h2 class="text-lg font-semibold mb-2 ">{{ email._source.Subject }}</h2>
            <span class="text-gray-500">{{ email._source.From }}</span>
          </button>
        </div>
      </div>
    </div>
</div>
</template>

<script >
import {useEmailStore }  from '../../store/emailStore.js'
import {storeToRefs} from 'pinia'
 
export default {
  data() {
    return {
      emails:storeToRefs(useEmailStore()).listOfEmails, 
      emailStore: useEmailStore(),
    };
  },
  mounted(){
    this.emailStore.initializeState()
  },
  methods: {
    formatDate(date) {
      return new Date(date).toLocaleDateString();
    },
    updateMail(mail){
      this.emailStore.update(mail)
    }
  },
};
</script>

<style>
</style>
