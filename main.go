package main

import (
	"TaskJWT/controller"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/login", controller.Login)
	http.HandleFunc("/refresh", controller.GetUser)

	log.Fatal(http.ListenAndServe(":8090", nil))
}
