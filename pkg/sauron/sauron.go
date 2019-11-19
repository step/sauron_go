package sauron

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/spf13/viper"

	"github.com/step/angmar/pkg/queueclient"
	"github.com/step/sauron_go/pkg/parser"
	"github.com/step/saurontypes"
)

// Repo is a custom type which will contain the ArchiveURL
// and the Name of the Repo
type Repo struct {
	ArchiveURL string `json:"archive_url"`
	Name       string `json:"name"`
}

// Pusher is a custom type which will contain the Name
// of the Pusher
type Pusher struct {
	Name string `json:"name"`
}

// Payload is a custom type which will contain the SHA,
// Timestamp and Repository and Pusher
type Payload struct {
	Ref        string `json:"ref"`
	After      string `json:"after"`
	Repository Repo
	Pusher     Pusher
}

func (payload *Payload) getArchiveURL(format string) string {
	archiveURL := payload.Repository.ArchiveURL
	archiveURL = strings.Replace(archiveURL, "{archive_format}", "tarball/", 1)
	archiveURL = strings.Replace(archiveURL, "{/ref}", payload.Ref, 1)
	return archiveURL
}

// Sauron is server which listens and responds to payload
// sent from github. And also parses the payload to make
// AngmarMessage and places on queue
type Sauron struct {
	Queue        string
	QueueClient  queueclient.QueueClient
	GithubSecret string
	Logger       SauronLogger
}

func (s Sauron) String() string {
	var builder strings.Builder
	builder.WriteString(s.QueueClient.String() + "\n")
	return builder.String()
}

// VerifyMessage is to verify the message if it is from github
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

func (s Sauron) getJSON(body string) Payload {
	payload := new(Payload)
	bodyReader := strings.NewReader(body)
	decoder := json.NewDecoder(bodyReader)
	err := decoder.Decode(payload)
	if err != nil {
		s.Logger.JSONDecodingError(err)
	}
	return *payload
}

func (s Sauron) getMessage(message Payload, sauronConfig saurontypes.SauronConfig) saurontypes.AngmarMessage {
	archiveURL := message.getArchiveURL("tarball")
	var tasks []saurontypes.Task

	for _, assignment := range sauronConfig.Assignments {
		assignmentName := strings.Split(message.Repository.Name, "-")[0]
		if assignment.Name == assignmentName {
			tasks = assignment.Tasks
			break
		}
	}
	angmarMessage := saurontypes.AngmarMessage{
		URL:     archiveURL,
		SHA:     message.After,
		Pusher:  message.Pusher.Name,
		Project: message.Repository.Name,
		Tasks:   tasks,
	}
	return angmarMessage
}

// Listener takes a viper instance to parse the config file
// and returns a listener for sauron
func (s Sauron) Listener(viperInst *viper.Viper) func(http.ResponseWriter, *http.Request) {
	s.Logger.StartSauron(s)
	return func(resp http.ResponseWriter, r *http.Request) {
		responseStatusCode := http.StatusOK
		signature := r.Header.Get("X-Hub-Signature")[5:]
		body, _ := ioutil.ReadAll(r.Body)
		defer r.Body.Close()

		sauronConfig := parser.ParseConfig(*viperInst)

		if isFromGithub(s.GithubSecret, signature, body) {
			message := s.getJSON(string(body))
			s.Logger.ReceivedMessage(message)
			angmarMessage := s.getMessage(message, sauronConfig)
			angmarMessageJSON, err := json.Marshal(angmarMessage)

			if err != nil {
				s.Logger.AngmarMessageMarshalingError(err)
			}

			err = s.QueueClient.Enqueue(s.Queue, string(angmarMessageJSON))

			if err != nil {
				s.Logger.EnqueueError(err)
			}
		} else {
			responseStatusCode = http.StatusForbidden
		}
		s.Logger.TaskPlacedOnQueue(s)
		resp.WriteHeader(responseStatusCode)
	}
}
