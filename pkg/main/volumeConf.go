package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

var sourceVolPath string
var logPath string
var logFilename string

func init() {
	flag.StringVar(&logPath, "log-path", "/tmp/sauron/log", "`location` where all source repositories are located")
	flag.StringVar(&logFilename, "log-filename", "sauron.log", "`filename` for logs")
}

func getLogfileName() string {
	return filepath.Join(logPath, logFilename)
}

func getLogfile() *os.File {
	file, err := os.OpenFile(getLogfileName(), os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		fmt.Println("here before")
		fmt.Println(err)
		fmt.Println("here")
		os.Exit(1)
	}
	return file
}
