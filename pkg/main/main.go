package main

import (
	"net/http"
	
	"github.com/step/sauron_go/pkg/sauron"
)

func main() {
	redisClient := GetRedisClient()
	sauron := sauron.Sauron{"angmar", redisClient, "test"}
	listener := sauron.Listener()
	http.HandleFunc("/", listener)
	http.ListenAndServe(":3333", nil)
}
