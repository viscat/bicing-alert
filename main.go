package main

import (
	"os"
	"fmt"
	"flag"
	"github.com/rifflock/lfshook"
	log "github.com/Sirupsen/logrus"
)

func main() {


	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Must specify a mode: update")
		os.Exit(1)
	}


	verbose := flag.Bool("v", false, "Show detailed information")
	update := flag.Bool("update", false, "Update bicing status")

	flag.Parse()

	initLogger(*verbose)

	switch {
		case *update:
			UpdateBicingStatus()
		default:
			fmt.Println("Must specify an action: update")


	}

}

func initLogger(verbosity bool) {
	if verbosity {
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.ErrorLevel)
	}

	log.SetOutput(os.Stdout)
	formatter := &log.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
		FullTimestamp:   true}
	log.SetFormatter(formatter)
	hook := lfshook.NewHook(lfshook.PathMap{
		log.InfoLevel:  "/var/log/info.log",
		log.ErrorLevel: "/var/log/error.log",
	})
	hook.SetFormatter(formatter)
	log.AddHook(hook)
}
