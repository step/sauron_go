package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/step/saurontypes"
	"io/ioutil"
	"net/http"
	"strings"
)

type Repo struct {
	Name string `json:name`
}

type Pusher struct {
	Name string `json:"name"`
}

type Payload struct {
	Repository Repo
	Pusher     Pusher
}

func getMessage(repoName string, pusher string) saurontypes.AngmarMessage {
	url := "https://api.github.com/repos/__NAME__/__REPO__/tarball/refs/heads/master"
	url = strings.Replace(url, "__NAME__", pusher, 1)
	url = strings.Replace(url, "__REPO__", repoName, 1)
	msg := saurontypes.AngmarMessage{
		Url:     url,
		SHA:     "master",
		Pusher:  pusher,
		Project: repoName,
		Tasks: []saurontypes.Task{
			{Queue: "test", ImageName: "mocha"},
			{Queue: "lint", ImageName: "eslint"},
		},
	}
	return msg
}

func githubPushListner(res http.ResponseWriter, r *http.Request) {
	flag.Parse()
	defer r.Body.Close()
	b, _ := ioutil.ReadAll(r.Body)
	var payload Payload
	json.Unmarshal(b, &payload)
	repoName := payload.Repository.Name
	pusher := payload.Pusher.Name

	redisClient := GetRedisClient()
	message, _ := json.Marshal(getMessage(repoName, pusher))

	redisClient.Enqueue(queueName, string(message))
	fmt.Println(string(message))

	res.WriteHeader(http.StatusOK)
}

func main() {
	http.HandleFunc("/", githubPushListner)
	http.ListenAndServe(":3333", nil)
}
