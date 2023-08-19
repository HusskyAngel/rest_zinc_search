package middleware

import (
	"github.com/HusskyAngel/BackendEmailSearcher/pkg/chirouter"
	"github.com/go-chi/chi/v5/middleware"
)

//Here yo should initialize all middleware that do you want
func StartMiddleware(){
  chirouter.Router.Use(middleware.Logger)
  chirouter.Router.Use(corsConfig)
}
