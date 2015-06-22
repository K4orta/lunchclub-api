package main

import (
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/k4orta/lunchclub-api/api"
	"github.com/rs/cors"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/users", api.CreateUser).Methods("POST")
	router.HandleFunc("/users", api.ReadUsers)
	// router.HandleFunc("/users", api.UpdateUser).Methods("PUT")
	// router.HandleFunc("/users", api.DeleteUser).Methods("DELETE")

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
	})

	n := negroni.New()
	n.Use(c)
	n.UseHandler(router)

	n.Run(":8001")
}
