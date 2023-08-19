import { defineStore } from 'pinia'
import {allSearchApi} from '../services/services.js'


export const useEmailStore =defineStore('currentMail',{
  state:()=>{
    return {
      currentMail:{}, 
      listOfEmails:[],
      currentPage:  1,
      numberOfResults: 0,
      searchQuery:"",
      filters:{
        MessageId:"",
        Date:"",
        From:"",
        To:"",
        CC:"",
        Subject:"",
        Body:""
      } 
    }
  },

  actions:{

     async initializeState(){
      const result= await allSearchApi(1)
      this.$patch({...this.state, listOfEmails:result.results,numberOfResults:result.number_results})
    },
    updateFilter(key,value){
      this.$patch({...this.state,filters:{...this.state.filters,key:value}})
    },
    update(newMail){
      this.$patch({...this.state,currentMail:newMail})
    },
    updateListOfEmails(emails){
      this.$patch({...this.state,listOfEmails:emails})
    },
    updatePage(page){
      this.$patch({...this.state,currentPage:page})
    },
    updateNumberOfResults(number){
      this.$patch({... this.state, numberOfResults:number})
    }

  }
})
