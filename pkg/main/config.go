package main

import (
	"flag"
)

var queueName string
var streamName string
var configFileName string
var configFilePath string
var githubSecret string
var port string
var numberOfMessages int

func init() {
	flag.StringVar(&queueName, "q", "angmar", "`queue` to push jobs to")
	flag.StringVar(&streamName, "stream", "eventHub", "stream to publish to")
	flag.StringVar(&configFileName, "configFileName", "config", "config file name for sauron")
	flag.StringVar(&configFilePath, "configFilePath", ".", "config file path for sauron")
	flag.StringVar(&port, "port", "3333", "port to listen to")
	flag.StringVar(&githubSecret, "secret", "test", "github secret key")
	flag.IntVar(&numberOfMessages, "n", 1, "`number of jobs` to push")
}
