package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main(){
	fmt.Println("Testing")
	godotenv.Load(".env")
	portString := os.Getenv("PORT")
	if portString==""{
		log.Fatal("Port not set");
	}
	fmt.Println("Port:",portString)


}