package post

import (
	"log"
	"net/http"
	"strings"

	"github.com/HusskyAngel/ZincIndexer/config"
)
  
func PostEmail(data string) error {

  configAdminUser:=config.GetConfig().AdminUser
  configAdminPassword:=config.GetConfig().AdminPassword
  configUrl:=config.GetConfig().Url

  req, err := http.NewRequest("POST",configUrl, strings.NewReader(data))
  if err !=nil{
    log.Println("Error: fail creating the post request")
    return err
  }
  req.SetBasicAuth(configAdminUser,configAdminPassword)
  req.Header.Set("Content-Type","application/json")
  req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Safari/537.36")

  resp,err:= http.DefaultClient.Do(req)
  if (err!= nil){
    log.Println("Error: doing post request")
    return err
  }
	defer resp.Body.Close()
  log.Println(resp.StatusCode)

  return nil
}
