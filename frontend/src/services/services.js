import axios from 'axios';

export async function allSearchApi(pageNumber){
   try {
//    const apiURL=process.env.VUE_APP_API_URL ;
    const response = await axios.get("http://localhost:4444/"+"all/"+(pageNumber-1).toString());
    return response.data

   } catch (error) {
    return {error, msg:"error in allSearchApi"} 
   }
}

export async function specificSearchApi(pageNumber,params){
  try {
    const url = `http://localhost:4444/specificsearch/${pageNumber}`;
    console.log(params.Search) 
    const data=await axios({
      method: 'post', 
      url: url,
      data: JSON.stringify(params),
      headers: {'Content-Type': 'application/json'}
    })
    return data.data
  } catch (error) {
    console.log(params)
    return { error, msg: 'error in specificSearchApi', params };
  }
}


export async function generalSearchApi(pageNumber, params) {
  try {
    const url = `http://localhost:4444/generalsearch/${pageNumber}`;
    console.log(params.Search) 
    const data=await axios({
      method: 'post', 
      url: url,
      data: JSON.stringify(params),
      headers: {'Content-Type': 'application/json'}
    })
    return data.data
  } catch (error) {
    console.log(params)
    return { error, msg: 'error in generalSearchApi', params };
  }
}


