package main

import (
	"flag"
)

var queueName string
var numberOfMessages int

func init() {
	flag.StringVar(&queueName, "q", "angmar", "`queue` to push jobs to")
	flag.IntVar(&numberOfMessages, "n", 1, "`number of jobs` to push")
}
