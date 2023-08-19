package zsrequest

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/HusskyAngel/BackendEmailSearcher/config"
	"github.com/HusskyAngel/BackendEmailSearcher/pkg/bodies"
)

func ZSRequest(query string) (bodies.ZSBody,error){
	req, errRB:= http.NewRequest("POST",config.GetConfig().ZSUrl ,strings.NewReader(query))
  if errRB!= nil {
    return bodies.ZSBody{},errRB  
  }

  req.SetBasicAuth(config.GetConfig().ZSAdmin, config.GetConfig().ZSPassword)
  req.Header.Set("Content-Type", "application/json")
  req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Safari/537.36")

  resp, errD := http.DefaultClient.Do(req)
  if errD != nil {
      return bodies.ZSBody{},errD
  }
  defer resp.Body.Close()

  log.Println(resp.StatusCode)
  body, errB := io.ReadAll(resp.Body)
  if errB != nil {
    return bodies.ZSBody{},errB 
  }
  
  var zsbody bodies.ZSBody
  errJ:=json.Unmarshal([]byte(string(body)),&zsbody)
  if errJ != nil{
    return bodies.ZSBody{},errJ
  }
  return zsbody,nil 
}
