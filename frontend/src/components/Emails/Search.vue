<template>
  <div class="w-full  mx-auto">
    <div class="grid grid-rows-2 justify-items-center items-center">
      <div>
        <input
          type="text"
          class="border-2 mt-5 border-gray-300 rounded-md px-40 py-2 focus:outline-none focus:border-blue-500 mx-5"
          v-model="searchQuery"
          placeholder="Search by word 🔎..."
        />
        <SearchButton :customFunction="handleSearch"/>
        <ResetButton :customFunction="handleReset"/>
      </div> 
        <ShowFiltersButton :customFunction="toggleHiddenElements" :hidden="showHidden"/>
        <div v-if="showHidden" class="container grid grid-cols-2 gap-6 justify-items-center items-center max-w-md">
          <div v-for="(value, key) in filters" :key="key" >
            <input
              type="text"
              :value="value"
              @input="updateFilter(key, $event.target.value)"
              class="border-2 mt-5 mr-10 ml-10 px-2 border-gray-300 rounded-md py-2 focus:outline-none focus:border-blue-500 max-w-xs"
              :placeholder="key"
            />
          </div>
        </div>
    </div>
  </div>
</template>

<script>
import SearchButton from './buttons/SearchButton.vue'
import ResetButton from './buttons/ResetButton.vue'
import ShowFiltersButton from './buttons/ShowFiltersButton.vue';

import {useEmailStore} from  '../../store/emailStore.js'
import {storeToRefs} from 'pinia'
 
import {getUpdatedEmails} from '../../services/utils.js'

export default {
  components: {SearchButton,ResetButton,ShowFiltersButton},
  methods:{
    updateFilter(key, value) {
     useEmailStore().updateFilter(key,value) 
    },
    toggleHiddenElements() {
      this.showHidden = !this.showHidden;
    },
    async handleSearch(){
      useEmailStore().updatePage(1)
      const vals= await getUpdatedEmails(1)
      useEmailStore().updateListOfEmails(vals.results)
    },
    handleReset(){
      for (let key in this.filters) {
        this.filters[key] = '';
      }
      this.searchQuery='';
    }
  } ,
  data() {
    return {
      searchQuery:storeToRefs(useEmailStore()).searchQuery, 
      showHidden: false,
      filters:storeToRefs(useEmailStore()).filters,
    };
  },
};
</script>

<style>

</style>
