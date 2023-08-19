
<template>
<div >
  <section class="background-radial-gradient mt-15 min-h-screen flex flex-col">
    <div class="grid mt-15 justify-center items-center">
      <Search/>
      <Pagination       
      :currentPage="currentPage"
      :totalPages="totalPages"
      :changePage="changePage"
      :pagItems="items"/>
      <div class="grid grid-cols-2 gap-4  whitespace-normal break-words">
        <List/>
        <EmailInfo/>
      </div>
      <Pagination       
      :currentPage="currentPage"
      :totalPages="totalPages"
      :changePage="changePage"
      :pagItems="items"/>
    </div>
  </section>
</div>
</template>

<script>
import List from '../components/Emails/List.vue'
import EmailInfo from '../components/Emails/EmailInfo.vue'
import Search from '../components/Emails/Search.vue'
import Pagination from '../components/Emails/Pagination.vue'

import {useEmailStore} from  '../store/emailStore.js'
import {storeToRefs} from 'pinia'

import {getUpdatedEmails} from '../services/utils.js'

export default{
  components: {List,EmailInfo,Search,Pagination},
  data() {
    return {
      itemsPerPage: 10,
      storeData:storeToRefs(useEmailStore())
    };
  },
  computed:{
    totalPages() {
      return Math.ceil(this.storeData.numberOfResults / 10);
    },
    currentPage(){
      return  this.storeData.currentPage
    },
    items(){
      if (this.storeData.currentPage==1){
        return [1,2,3]
      }else if (this.storeData.currentPage==this.totalPages){
          
        return [this.totalPages-2,this.totalPages-1,this.totalPages]
      }else{
        return [this.storeData.currentPage-1,this.storeData.currentPage,this.storeData.currentPage+1]
      }
    }
  },
  methods:{
    async changePage(pageNumber) {
      useEmailStore().updatePage(pageNumber)
      const result=await getUpdatedEmails(this.storeData.currentPage)
      useEmailStore().updateListOfEmails(result.results)
    }
  }
}
</script>

<style> 
  .background-radial-gradient {
    background-color: hsl(218, 41%, 15%);
    background-image: radial-gradient(650px circle at 0% 0%,
        hsl(218, 41%, 35%) 15%,
        hsl(218, 41%, 30%) 35%,
        hsl(218, 41%, 20%) 75%,
        hsl(218, 41%, 19%) 80%,
        transparent 100%),
      radial-gradient(1250px circle at 100% 100%,
        hsl(218, 41%, 45%) 15%,
        hsl(218, 41%, 30%) 35%,
        hsl(218, 41%, 20%) 75%,
        hsl(218, 41%, 19%) 80%,
        transparent 100%);
  }

</style>
