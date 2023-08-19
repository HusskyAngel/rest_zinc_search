<template>
  <div class="pagination">
    <ul class="flex space-x-1">
      <li>
       <button
          @click="changePage(1)"
          class="
            hover:bg-blue-500 hover:text-white
            bg-gray-300 text-gray-700 px-3 py-1 rounded-md
          "
        >
          Inicio 
       </button>
      </li>
      <li v-for="pageNumber in pagItems" :key="pageNumber">
        <button  
          @click="changePage(pageNumber)"
          :class="{
            'bg-blue-500 text-white': pageNumber === currentPage,
            'bg-gray-300 text-gray-700': pageNumber !== currentPage
          }"
          class="px-3 py-1 rounded-md"
        >
          {{ pageNumber }}
        </button>
      </li>
         <button
          @click="changePage(this.numberOfResults)"
          class="
            hover:bg-blue-500 hover:text-white
            bg-gray-300 text-gray-700 px-3 py-1 rounded-md
          "
        >
          Final 
       </button>

    </ul>
  </div>
</template>

<script>
import {useEmailStore} from '../../store/emailStore.js'
import {storeToRefs} from 'pinia'


export default {
  data(){
    return {
      storeData:storeToRefs(useEmailStore()),
    }     
  },
  computed:{
    numberOfResults(){
      return Math.ceil(this.storeData.numberOfResults/10)
    }
  },
  props: {
    currentPage: {
      type: Number,
      required: true
    },
    totalPages: {
      type: Number,
      required: true
    },
    changePage: {
      type: Function,
      required: true
    },
    pagItems: {
      type: Array,
      required:true 
    }
  }
};
</script>

<style>
.pagination {
  display: flex;
  justify-content: center;
  margin-top: 1rem;
  margin-bottom: 1rem;
}
</style>
