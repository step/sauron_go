package main

import (
	"flag"
	"fmt"
	"github.com/step/sauron_go/pkg/flowidgenerator"
	"net/http"

	"github.com/spf13/viper"
	"github.com/step/sauron_go/pkg/sauron"
)

func main() {
	flag.Parse()

	viperInst := viper.New()
	viperInst.SetConfigName(configFileName)
	viperInst.AddConfigPath(configFilePath)
	err := viperInst.ReadInConfig()

	if err != nil {
		fmt.Printf("Unable to read config \n%s", err)
	}

	file := getLogfile()

	logger := getLogger(file)
	redisClient := GetRedisClient()
	uuidGenerator := flowidgenerator.NewUUIDGenerator()
	sauron := sauron.Sauron{
		Queue:           queueName,
		QueueClient:     redisClient,
		Stream:          streamName,
		StreamClient:    redisClient,
		Flowidgenerator: uuidGenerator,
		GithubSecret:    githubSecret,
		Logger:          logger,
	}
	listener := sauron.Listener(viperInst)
	port := fmt.Sprintf(":%s", port)
	http.HandleFunc("/", listener)
	http.ListenAndServe(port, nil)
}
