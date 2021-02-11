package bootstrap

import (
	"github.com/rest_api/pkg/logger"

	log "github.com/sirupsen/logrus"
)

func initLog() {
	log.SetFormatter(&logger.Formatter{
		ChildFormatter: &log.JSONFormatter{
			FieldMap: log.FieldMap{
				log.FieldKeyMsg: "message",
			},
		},
		Line:         true,
		Package:      false,
		File:         true,
		BaseNameOnly: false,
	})

	h, e := logger.NewSentryHook(cfg)

	if e != nil {
		logger.Fatal("log sentry failed to initialize Sentry")
	}

	if h != nil {
		log.AddHook(h)
	}

	log.SetLevel(log.InfoLevel)
}
