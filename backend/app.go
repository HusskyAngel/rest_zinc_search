package main

import (
	"log"
	"net/http"

	"github.com/HusskyAngel/BackendEmailSearcher/config"
	"github.com/HusskyAngel/BackendEmailSearcher/middleware"
	"github.com/HusskyAngel/BackendEmailSearcher/pkg/chirouter"
	"github.com/HusskyAngel/BackendEmailSearcher/routes"
)

func main() {
	
  middleware.StartMiddleware()
  routes.StartRoutes()

  log.Println("starting server in http://localhost:"+ config.GetConfig().Port)
	http.ListenAndServe(":"+ config.GetConfig().Port, chirouter.Router)
}
