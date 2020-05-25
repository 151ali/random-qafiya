package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Index :
func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//
	randomKey, _ := RedisClient.RandomKey().Result()
	log.Println("RandomKey generated : ", randomKey)
	//
	res, _ := RedisClient.SRandMember(randomKey).Result()
	fmt.Fprintln(w, res)

}

// AddQafiya :
func addQafiya(w http.ResponseWriter, r *http.Request) {
	var NewQafiya Qafiya

	err := json.NewDecoder(r.Body).Decode(&NewQafiya)
	if err != nil {

		log.Println("json Decoder said :", err)
	}

	v, err := json.Marshal(NewQafiya)

	if err != nil {
		log.Println("json Marshal said :", err)
	}
	if NewQafiya.Letter != "" && NewQafiya.Name != "" {

		RedisClient.SAdd(NewQafiya.Letter, v)

	}

}
