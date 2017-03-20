package main

import (
	"bicingalert/app"
	"flag"
	log "github.com/Sirupsen/logrus"
	"github.com/rifflock/lfshook"
	"os"
)

func main() {

	verbose := flag.Bool("v", false, "Show detailed information")
	flag.Parse()
	initLogger(*verbose)

	storage := app.Storage{Db: app.GetBicingDb()}
	storage.UpdateBicingStatus()
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
