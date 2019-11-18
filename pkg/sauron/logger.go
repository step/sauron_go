package sauron

import (
	"strings"
	"log"
)

type SauronLogger struct {
	Logger *log.Logger
}

func (l SauronLogger) AngmarMessageMarshalingError(err error)  {
	l.Logger.Printf("\nError => %s\n", err)
}

func (l SauronLogger) EnqueueError(err error)  {
	l.Logger.Printf("\nError => %s\n", err)
}

func (l SauronLogger) JsonDecodingError(err error)  {
	l.Logger.Printf("\nError => %s\n", err)
}

func (l SauronLogger) ConfigReadingError(err error)  {
	l.Logger.Printf("\nError => %s\n", err)
}

func (l SauronLogger) StartSauron(s Sauron)  {
	var builder strings.Builder
	builder.WriteString("Starting Sauron...\n")
	builder.WriteString("---\n")
	builder.WriteString(s.String())
	builder.WriteString("Publishing to queue: " + s.Queue + "\n")
	builder.WriteString("---\n")

	l.Logger.Println(builder.String())
}

func (l SauronLogger) ReceivedMessage(message Payload) {
	var builder strings.Builder
	builder.WriteString("Received Job...\n")
	builder.WriteString("Pusher name: " + message.Pusher.Name + "\n")
	builder.WriteString("Repo name: " + message.Repository.Name + "\n")
	builder.WriteString("SHA: " + message.Ref + "\n")
	builder.WriteString("Archieve Url: " + message.Repository.Archive_url + "\n")
	l.Logger.Println(builder.String())
}

func (l SauronLogger) TaskPlacedOnQueue(s Sauron) {
	var builder strings.Builder
	builder.WriteString("Task placed on queue " + s.Queue + "\n")
	builder.WriteString(s.String())
	l.Logger.Println(builder.String())
}