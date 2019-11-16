package sauron_test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/step/saurontypes"

	"github.com/step/angmar/pkg/queueclient"
	"github.com/step/sauron_go/pkg/sauron"
)

func TestSauron(t *testing.T) {
	content, err := ioutil.ReadFile("./payload")
	if err != nil {
		t.Error("Unable to read file 'payload'")
	}

	reader := bytes.NewReader(content)

	q := queueclient.NewDefaultClient()
	s := sauron.Sauron{"angmar", q, "test"}
	l := s.Listener()

	sauronServer := httptest.NewServer(http.HandlerFunc(l))
	request, err := http.NewRequest("POST", sauronServer.URL, reader)
	request.Header.Set("X-Hub-Signature", "sha1=4fa319856acc674327465391f682133675688aaa")
	if err != nil {
		t.Error("unable to create request")
	}
	response, err := sauronServer.Client().Do(request)

	if err != nil {
		t.Errorf("Response error %s", err.Error())
	}

	if response.StatusCode != http.StatusOK {
		t.Errorf("Wrong response code. Expected 200, got %d", response.StatusCode)
	}

	message, err := q.Dequeue("angmar")
	if err != nil {
		t.Errorf("Unable to dequeue from angmar %s\n", err.Error())
	}
	expectedAngmarMessage := saurontypes.AngmarMessage{
		Project: "sample-assignment",
		Pusher:  "craftybones",
		SHA:     "cc08dafb86c16562a8b876d195a31cd6d99feae9",
		Url:     "https://api.github.com/repos/craftybones/sample-assignment/tarball/refs/heads/master",
		Tasks: []saurontypes.Task{
			{Queue: "test", ImageName: "mocha"},
			{Queue: "lint", ImageName: "eslint"},
		},
	}
	expected, err := json.Marshal(expectedAngmarMessage)
	if err != nil {
		t.Errorf("Error marshalling dequeued message\n%s\n", err.Error())
	}

	if string(expected) != message {
		t.Errorf("expected %s\nactual %s\n", expected, message)
	}
}
