package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main(){
	fmt.Println("Testing")
	godotenv.Load(".env")
	portString := os.Getenv("PORT")
	if portString==""{
		log.Fatal("Port not set");
	}
	
	router:=chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))
	v1Router:=chi.NewRouter()
	v1Router.HandleFunc("/ready",handlerReadiness)
	router.Mount("/v1",v1Router)
	server:= &http.Server{
			Handler:router,
			Addr: ":"+portString,
	}
	fmt.Println("Server listening on Port: ",portString)
	err:=server.ListenAndServe()
	if(err!=nil){
		log.Fatal(err)
	}

}