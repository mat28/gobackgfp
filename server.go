package main


import (
	"gobackgfp/routers"
	"github.com/codegangsta/negroni"
	"net/http"
	"github.com/rs/cors"
)

func main(){
	c := cors.New(cors.Options{
		AllowedMethods: []string{"GET", "POST", "DELETE","PUT", "OPTIONS","PATCH"},
	})

	router := routers.InitRoutes()
	n := negroni.Classic()
	n.Use(c)
	n.UseHandler(router)
	http.ListenAndServe(":5000",n)
}
