package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/HusskyAngel/BackendEmailSearcher/pkg/bodies"
	"github.com/HusskyAngel/BackendEmailSearcher/pkg/querybuilder"
	"github.com/HusskyAngel/BackendEmailSearcher/pkg/zsrequest"
	"github.com/go-chi/chi"
)

//end point for return all emails
func AllSearch(w http.ResponseWriter, r *http.Request){
  pagParam := chi.URLParam(r, "pag")

  query:=querybuilder.AllQuerySearch(pagParam)
  zsObject,err:=zsrequest.ZSRequest(query)
  if err!= nil {
    log.Println(err)
    return
  }
  var responseObject= bodies.ResponseBody{NumberResults: zsObject.Hits.Total.Value ,Results: zsObject.Hits.Hits}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

  err=json.NewEncoder(w).Encode(responseObject)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
