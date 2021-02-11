package main

import (
	"os"

	"github.com/rest_api/bootstrap"
	mongodb "github.com/rest_api/db/mongo"

	"github.com/prometheus/common/log"
)

func main() {
	cmdString := command()
	if cmdString == "migrate" {
		log.Infoln("Starting migration")
		err := mongodb.Migrate()
		if err != nil {
			log.Fatal("Failed to run migration: ", err)
		}
		log.Info("Applied migrations successfully")
	} else {
		bootstrap.RunApp()
	}
}

func command() string {
	args := os.Args[1:]

	if len(args) > 0 {
		return args[0]
	}
	return ""
}
