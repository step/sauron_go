package main

import (
	"fmt"
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
	sauron := sauron.Sauron{"angmar", redisClient, "test", logger}
	listener := sauron.Listener(viperInst)
	http.HandleFunc("/", listener)
	http.ListenAndServe(":3333", nil)
}
