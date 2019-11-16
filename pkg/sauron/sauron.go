package sauron

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/spf13/viper"

	"github.com/step/angmar/pkg/queueclient"
	"github.com/step/sauron_go/pkg/parser"
	"github.com/step/saurontypes"
)

type Repo struct {
	Archive_url string `json:archive_url`
	Name        string `json:name`
}

type Pusher struct {
	Name string `json:"name"`
}

type Payload struct {
	Ref        string `json:"ref"`
	After      string `json:"after"`
	Repository Repo
	Pusher     Pusher
}

func (payload *Payload) getArchiveUrl(format string) string {
	archiveURL := payload.Repository.Archive_url
	archiveURL = strings.Replace(archiveURL, "{archive_format}", "tarball/", 1)
	archiveURL = strings.Replace(archiveURL, "{/ref}", payload.Ref, 1)
	return archiveURL
}

type Sauron struct {
	Queue        string
	QueueClient  queueclient.QueueClient
	GithubSecret string
}

func VerifyMessage(message, key string, actualDigest []byte) bool {
	mac := hmac.New(sha1.New, []byte(key))
	mac.Write([]byte(message))
	expectedMac := mac.Sum(nil)
	return hmac.Equal(expectedMac, actualDigest)
}

func isFromGithub(key, signature string, body []byte) bool {
	hexDecodedSignature, _ := hex.DecodeString(signature)
	return VerifyMessage(string(body), key, hexDecodedSignature)
}

func getJSON(body string) Payload {
	payload := new(Payload)
	bodyReader := strings.NewReader(body)
	decoder := json.NewDecoder(bodyReader)
	err := decoder.Decode(payload)
	if err != nil {
		fmt.Println(err)
	}
	return *payload
}

func getMessage(body []byte, sauronConfig saurontypes.SauronConfig) saurontypes.AngmarMessage {
	message := getJSON(string(body))
	archiveURL := message.getArchiveUrl("tarball")
	var tasks []saurontypes.Task

	for _, assignment := range sauronConfig.Assignments {
		assignmentName := strings.Split(message.Repository.Name, "-")[0]
		if assignment.Name == assignmentName {
			tasks = assignment.Tasks
		}
	}
	angmarMessage := saurontypes.AngmarMessage{
		Url:     archiveURL,
		SHA:     message.After,
		Pusher:  message.Pusher.Name,
		Project: message.Repository.Name,
		Tasks:   tasks,
	}
	return angmarMessage
}

func (s Sauron) Listener() func(http.ResponseWriter, *http.Request) {
	return func(resp http.ResponseWriter, r *http.Request) {
		responseStatusCode := http.StatusOK
		signature := r.Header.Get("X-Hub-Signature")[5:]
		body, _ := ioutil.ReadAll(r.Body)
		defer r.Body.Close()

		viperInst := viper.New()
		viperInst.SetConfigName("config")
		viperInst.AddConfigPath(".")
		err := viperInst.ReadInConfig()

		if err != nil {
			fmt.Println("something happened")
		}

		sauronConfig := parser.ParseConfig(*viperInst)

		if isFromGithub(s.GithubSecret, signature, body) {
			angmarMessage := getMessage(body, sauronConfig)
			angmarMessageJSON, err := json.Marshal(angmarMessage)

			if err != nil {
				fmt.Printf("Error: %s\n", err)
			}

			err = s.QueueClient.Enqueue(s.Queue, string(angmarMessageJSON))

			if err != nil {
				fmt.Printf("Error: %s\n", err.Error())
			}
		} else {
			responseStatusCode = http.StatusForbidden
		}
		resp.WriteHeader(responseStatusCode)
	}
}
