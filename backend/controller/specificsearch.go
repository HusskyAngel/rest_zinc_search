package controller

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/HusskyAngel/BackendEmailSearcher/pkg/bodies"
	"github.com/HusskyAngel/BackendEmailSearcher/pkg/querybuilder"
	"github.com/HusskyAngel/BackendEmailSearcher/pkg/zsrequest"
	"github.com/go-chi/chi"
)

func SpecificSearch(w http.ResponseWriter, r *http.Request){
  pagParam := chi.URLParam(r, "pag")
  body,err :=ioutil.ReadAll(r.Body)
  if err!=nil{
     http.Error(w, "Request body is missing", http.StatusBadRequest)
  }
  defer r.Body.Close()
  var info bodies.BodyQuery
  errJ:=json.Unmarshal(body,&info)
  	if errJ != nil {
     log.Println(errJ)
     log.Println("error")
     http.Error(w, "reading body", http.StatusBadRequest)
	} 
  query:=querybuilder.SpecificQuerySearch(pagParam,info) 
  zsObject,errZS:=zsrequest.ZSRequest(query)
  if errZS!= nil {
    log.Println(errZS)
    http.Error(w, errZS.Error(), http.StatusInternalServerError)
    return
  }
  var responseObject= bodies.ResponseBody{NumberResults: zsObject.Hits.Total.Value ,Results: zsObject.Hits.Hits}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

  errNJ:=json.NewEncoder(w).Encode(responseObject)
  if errNJ != nil {
    http.Error(w, errNJ.Error(), http.StatusInternalServerError)
	}
}
