package routers

import (
	"gobackgfp/controllers"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

func SetHelloRoutes(router *mux.Router) *mux.Router {
	router.Handle("/hello",
		negroni.New(
			negroni.HandlerFunc(controllers.HelloController),
		)).Methods("GET")
	return router
}