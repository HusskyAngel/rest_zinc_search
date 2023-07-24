package post

import (
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/joho/godotenv"
)
  
func PostEmail(data string) {
  err := godotenv.Load()
  if err!=nil {
    log.Fatal("error loading .env file")
  }
  ZINC_FIRST_ADMIN_USER:=os.Getenv("ZINC_FIRST_ADMIN_USER")
  ZINC_FIRST_ADMIN_PASSWORD:=os.Getenv("ZINC_FIRST_ADMIN_PASSWORD")
  ZINC_URL:=os.Getenv("ZINC_URL")

  req, err := http.NewRequest("POST",ZINC_URL , strings.NewReader(data))
  if err !=nil{
    log.Fatal("fail doing the post request")
  }
  req.SetBasicAuth(ZINC_FIRST_ADMIN_USER,ZINC_FIRST_ADMIN_PASSWORD)
  req.Header.Set("Content-Type","application/json")
  req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Safari/537.36")

  resp,err:= http.DefaultClient.Do(req)
  if (err!= nil){
    log.Fatal(err)
  }
	defer resp.Body.Close()
  log.Println(resp.StatusCode)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Indexed ", body)
}
