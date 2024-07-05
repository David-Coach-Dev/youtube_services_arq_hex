package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"

)

func main() {
	// lee el archivo .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	// lee las variables de entorno port
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port= "7070"
	}

	// Compose all services
	youtubeAdapter := Compose()

	//youtubeApiAdapter := ComposeYT()
	router := CreateRouters(youtubeAdapter)

	server := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}

	log.Println("Server listening on http://0.0.0.0:"+port)
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
