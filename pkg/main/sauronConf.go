package main

import (
	"io"
	"log"
	"os"
	
	"github.com/step/sauron_go/pkg/sauron"
)

func getLogger(file *os.File) sauron.SauronLogger {
	multiWriter := io.MultiWriter(file, os.Stdout)

	actualLogger := log.New(multiWriter, "--> ", log.LstdFlags)
	return sauron.SauronLogger{Logger: actualLogger}
}