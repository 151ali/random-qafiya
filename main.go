package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-redis/redis"

	"github.com/gorilla/mux"
)

var port = os.Getenv("PORT")

// RedisClient :
var RedisClient *redis.Client

func main() {

	log.Println("Server running on port :", port, "...")
	router := mux.NewRouter()
	RedisClient, _ = ConnectRedis()

	AddRoutes(router)

	err := http.ListenAndServe(":"+port, router)

	if err != nil {
		log.Fatal("las said : ", err)
	}

}
