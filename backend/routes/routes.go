package routes

import (
	"github.com/HusskyAngel/BackendEmailSearcher/controller"
	"github.com/HusskyAngel/BackendEmailSearcher/pkg/chirouter"
)

//initialize the endpoints in the chirouter.Router
func StartRoutes(){
	chirouter.Router.Get("/",controller.HelloWorld )
	chirouter.Router.Post("/generalsearch/{pag}",controller.GeneralSearch)
  chirouter.Router.Post("/specificsearch/{pag}",controller.SpecificSearch)
  chirouter.Router.Get("/all/{pag}",controller.AllSearch)
}
