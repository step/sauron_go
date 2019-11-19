package sauron

import (
	"strings"
	"log"
)

// SauronLogger is a simple wrapper around a log.Logger
// and provides several convenience methods to log events
// specific to Sauron
type SauronLogger struct {
	Logger *log.Logger
}

// AngmarMessageMarshalingError is a type of error that happens 
// while marshaling AngmarMessage
func (l SauronLogger) AngmarMessageMarshalingError(err error)  {
	l.Logger.Printf("\nError => %s\n", err)
}

// EnqueueError is a type of error that happens during enqueueing
// something to queue
func (l SauronLogger) EnqueueError(err error)  {
	l.Logger.Printf("\nError => %s\n", err)
}

// JSONDecodingError is a type of error that happens while
// decoding the json payload
func (l SauronLogger) JSONDecodingError(err error)  {
	l.Logger.Printf("\nError => %s\n", err)
}

// ConfigReadingError is a type of error that happens while
// while reading the config for sauron
func (l SauronLogger) ConfigReadingError(err error)  {
	l.Logger.Printf("\nError => %s\n", err)
}

// StartSauron should be called when Sauron starts
// It logs the details of the sauron instance and
// the queue it is publishing to
func (l SauronLogger) StartSauron(s Sauron)  {
	var builder strings.Builder
	builder.WriteString("Starting Sauron...\n")
	builder.WriteString("---\n")
	builder.WriteString(s.String())
	builder.WriteString("Publishing to queue: " + s.Queue + "\n")
	builder.WriteString("---\n")

	l.Logger.Println(builder.String())
}

// ReceivedMessage should be called when sauron gets a
// payload and haven't started to execute it
func (l SauronLogger) ReceivedMessage(message Payload) {
	var builder strings.Builder
	builder.WriteString("Received Job...\n")
	builder.WriteString("Pusher name: " + message.Pusher.Name + "\n")
	builder.WriteString("Repo name: " + message.Repository.Name + "\n")
	builder.WriteString("SHA: " + message.Ref + "\n")
	builder.WriteString("Archieve Url: " + message.Repository.ArchiveURL + "\n")
	l.Logger.Println(builder.String())
}

// TaskPlacedOnQueue should be called when sauron has finished
// it job and has placed the AngmarMessage on queue
func (l SauronLogger) TaskPlacedOnQueue(s Sauron) {
	var builder strings.Builder
	builder.WriteString("Task placed on queue " + s.Queue + "\n")
	builder.WriteString(s.String())
	l.Logger.Println(builder.String())
}