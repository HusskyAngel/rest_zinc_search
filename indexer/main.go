package main

import (
	"encoding/json"
	"fmt"

	"github.com/HusskyAngel/ZincIndexer/pkg/mreader"
)

func main() {

  const path string="/home/HusskyAngel/Descargas/enron_mail_20110402/maildir/buy-r/sent/1."

  result:=mreader.EReader(path)
  json,err:=json.Marshal(result)
  if err!=nil{
    fmt.Println("error transformando a json")
  }
  fmt.Println(string(json))
  
}
