package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-redis/redis"

	"github.com/gorilla/mux"
)

// RedisClient :
var RedisClient *redis.Client

func main() {

	var port = os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("Server running on port :", port, "...")
	router := mux.NewRouter()
	RedisClient, _ = ConnectRedis()
	RedisClient.SAdd("alf", `{"letter":"alf", "name":"مرحبا"}`)

	AddRoutes(router)

	err := http.ListenAndServe(":"+port, router)

	if err != nil {
		log.Fatal("las said : ", err)
	}

}
