package querybuilder

import (
	"fmt"
	"log"
	"reflect"

	"github.com/HusskyAngel/BackendEmailSearcher/pkg/bodies"
)

func SpecificQuerySearch(pag string,filter bodies.BodyQuery) string{
  bodyValue:=reflect.ValueOf(filter)
  bodyType:=bodyValue.Type()
  
  var search string 
  var aggs string="{"

  for i :=0;i < bodyValue.NumField(); i++{
    fieldName:=bodyType.Field(i).Name
    fieldValue:=fmt.Sprintf("%v",bodyValue.Field(i).Interface())
    
    if fieldName=="Search"{
      search=fmt.Sprintf("%s",fieldValue)
    }else if fieldName=="MessageId" && fieldValue!=""{
      var agg string =`"%s":{"agg_type":"term","field": "Message-Id","size"= 10},`
      agg=fmt.Sprintf(agg,fieldValue)
      aggs+=agg
    }else if fieldValue!=""{
      var agg string =`"%s":{"agg_type":"term","field":"%s","size"= 10},`
      log.Println("entro2")
      agg=fmt.Sprintf(agg,fieldValue,fieldName)
      aggs+=agg
    }else{
      continue
    }
  }
  aggs+="}"

  var query string=`{
    "search_type": "match",
    "query": {
        "term": "%s",
        "field": "_all"
    },
    "sort_fields": ["-@timestamp"],
    "from": %s,
    "max_results": 10,
    "aggs":%s,
    "_source": [
    ]
}`
  query=fmt.Sprintf(query,search,pag,aggs)
  log.Println(query)

  return query
}


