package main

import (
	"log"
	"net/http"
	"songchord-api/driver"
	"songchord-api/routes"

	"github.com/joho/godotenv"
)

func main() {
	r := routes.RegisterRoutes()
	log.Println("Server ready at 8081")
	log.Fatal(http.ListenAndServe(":8081", r))
}

func init() {
	loadEnv()
	driver.ConnectDatabase()
}

func loadEnv() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}
