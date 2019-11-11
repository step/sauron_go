package main

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/step/angmar/pkg/redisclient"

	"github.com/step/saurontypes"
)

type Repo struct {
	Archive_url string `json:archive_url`
	Name        string `json:name`
}

type Pusher struct {
	Name string `json:name`
}

type Payload struct {
	Ref        string `json:ref`
	After      string `json:after`
	Repository Repo
	Pusher     Pusher
}

func (payload *Payload) getArchiveUrl(format string) string {
	archiveUrl := payload.Repository.Archive_url
	archiveUrl = strings.Replace(archiveUrl, "{archive_format}", "tarball/", 1)
	archiveUrl = strings.Replace(archiveUrl, "{/ref}", payload.Ref, 1)
	return archiveUrl
}

func getJSON(body string) Payload {
	payload := new(Payload)
	bodyReader := strings.NewReader(body)
	decoder := json.NewDecoder(bodyReader)
	err := decoder.Decode(payload)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(payload)
	return *payload
}

func isFromGithub(signature string, body []byte) bool {
	hexDecodedSignature, _ := hex.DecodeString(signature)
	return VerifyMessage(string(body), "test", hexDecodedSignature)
	// if isValidMessage {
	// 	v := getJSON(string(body))
	// 	fmt.Println(v)
	// 	archiveUrl := v.getArchiveUrl("tarball")
	// 	fmt.Println(archiveUrl)
	// }
}

// type AngmarMessage struct {
// 	Url       string   `json:"url"`
// 	SHA       string   `json:"sha"`
// 	Pusher    string   `json:"pusher"`
// 	Project   string   `json:"project"`
// 	ImageName string   `json:"imageName"`
// 	Tasks     []string `json:"tasks"`
// }

func handleGithubEvent(w http.ResponseWriter, r *http.Request) {
	signature := r.Header.Get("X-Hub-Signature")[5:]
	body, _ := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if isFromGithub(signature, body) {
		message := getJSON(string(body))
		archiveUrl := message.getArchiveUrl("tarball")
		x := saurontypes.AngmarMessage{
			Url:       archiveUrl,
			SHA:       message.After,
			Pusher:    message.Pusher.Name,
			Project:   message.Repository.Name,
			ImageName: "orc_sample",
			Tasks:     []string{"test"},
		}
		rc := redisclient.NewDefaultClient(redisclient.RedisConf{
			Address:  "localhost:6379",
			Db:       2,
			Password: "",
		})
		msg, _ := json.Marshal(x)
		err := rc.Enqueue("my_queue", string(msg))
		if err != nil {
			fmt.Println("Unable to enqueue", string(msg), err)
		}
	}
}

func main() {
	http.HandleFunc("/", handleGithubEvent)
	http.ListenAndServe(":8080", nil)
}
