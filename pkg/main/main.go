package main

import (
	"fmt"
	"github.com/step/sauron_go/pkg/flowidgenerator"
	"net/http"

	"github.com/spf13/viper"
	"github.com/step/sauron_go/pkg/sauron"
)

func main() {

	viperInst := viper.New()
	viperInst.SetConfigName("config")
	viperInst.AddConfigPath(".")
	err := viperInst.ReadInConfig()

	if err != nil {
		fmt.Printf("Unable to read config \n%s", err)
	}

	file := getLogfile()

	logger := getLogger(file)
	redisClient := GetRedisClient()
	uuidGenerator := flowidgenerator.NewUUIDGenerator()
	sauron := sauron.Sauron{
		Queue:           "angmar",
		QueueClient:     redisClient,
		StreamClient:    redisClient,
		Flowidgenerator: uuidGenerator,
		GithubSecret:    "test",
		Logger:          logger,
	}
	listener := sauron.Listener(viperInst)
	http.HandleFunc("/", listener)
	http.ListenAndServe(":3333", nil)
}
