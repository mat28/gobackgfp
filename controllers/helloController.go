package controllers


import (
	"net/http"
	"encoding/json"
	"log"
)


func HelloController(w http.ResponseWriter, r *http.Request, next http.HandlerFunc){
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(200)
	data := struct {
		Message string
	}{
		"Hello World",
	}
	if err := json.NewEncoder(w).Encode(data);err != nill {
		log.Println(err)
	}
}