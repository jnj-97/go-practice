package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/jnj-97/go-practice/internal/database"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type apiConfig struct {
		DB *database.Queries
	}

func main(){
	fmt.Println("Testing")
	godotenv.Load(".env")
	portString := os.Getenv("PORT")
	db_url:=os.Getenv("DB_URL")
	if portString==""{
		log.Fatal("Port not set");
	}
	if db_url==""{
		log.Fatal("DB_URL Not Found")
	}
	conn,err:=sql.Open("postgres",db_url)
	if err!=nil{
		log.Fatal("Unable to connect to Database",err)
	}
	queries:=database.New(conn)
	

	apiCfg:=apiConfig{
		DB: queries,
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
	v1Router.Get("/ready",handlerReadiness)
	v1Router.Get("/error",handlerError)
	v1Router.Post("/users",apiCfg.handleUser)
	v1Router.Get("/users",apiCfg.middlewareAuth(apiCfg.handleGetUser))
	v1Router.Post("/feeds",apiCfg.middlewareAuth(apiCfg.handleCreateFeed))
	v1Router.Get("/feeds",apiCfg.handleGetFeeds)
	v1Router.Post("/feed_follow",apiCfg.middlewareAuth(apiCfg.handleCreateFeedFollow))
	v1Router.Get("/feed_follows",apiCfg.middlewareAuth(apiCfg.handleGetFeedFollows))
	v1Router.Delete("/feed_follows/{feed_id}",apiCfg.middlewareAuth(apiCfg.handleDeleteFeedFollow))
	router.Mount("/v1",v1Router)
	server:= &http.Server{
			Handler:router,
			Addr: ":"+portString,
	}
	fmt.Println("Server listening on Port: ",portString)
	err=server.ListenAndServe()
	if(err!=nil){
		log.Fatal(err)
	}

}