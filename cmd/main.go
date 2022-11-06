package main

import (
	"github.com/hpardora/absence.go/cmd/client"
	"github.com/hpardora/absence.go/pkg/logger"
)

func main() {
	log := logger.GetLogger()
	log.Infof("starting Absence.io client")

	rootCmd := client.NewCMDCli()
	if err := rootCmd.Execute(); err != nil {
		log.WithError(err).Fatalf("error starting HATT")
	}
}
