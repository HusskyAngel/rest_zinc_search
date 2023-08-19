import {useEmailStore} from  '../store/emailStore.js'
import {storeToRefs} from 'pinia'

import {generalSearchApi, allSearchApi,specificSearchApi} from './services.js'


function filtersEmpty(filters ){
  Object.keys(filters).forEach(key => {
    if  (filters[key]!=""){
      return false 
    }
  });

  return true 
}

export async function getUpdatedEmails(pag){
  const filters=useEmailStore().filters
  const query=useEmailStore().searchQuery

  let vals={}

  if (filtersEmpty(filters) && query===""){// no search parameters
    vals =await allSearchApi(pag)
  }else if(filtersEmpty(filters) && query !=""){// just search parameters
    const params={
      Search:query
    } 
    vals=await generalSearchApi(pag,params)
  }else{ //filters and search parameters
    const params={
      Search:query,
      ...filters
    }
    vals=await specificSearchApi(pag,params)
  }

  return vals 
}  

