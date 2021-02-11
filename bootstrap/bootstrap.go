package bootstrap

import (
	"fmt"

	"github.com/rest_api/pkg/logger"
)

// RunApp :nodoc:
func RunApp() {
	initLog()
	initDatabases()
	initRepository()
	initService()

	restRouter := initREST()

	if err := restRouter.Run(fmt.Sprintf(":%s", cfg.App.Port)); err != nil {
		logger.Panic(
			logger.SetMessageFormat("Receiving error: %v", err),
		)
	}
}
