package client

import (
	"github.com/hpardora/absence.go/internal/service/scheduler"
	"github.com/hpardora/absence.go/pkg/absence"
	"github.com/hpardora/absence.go/pkg/logger"
	"github.com/spf13/cobra"
	"os"
)

func NewCmdScheduler() *cobra.Command {
	log := logger.GetLogger()
	log.Infof("starting scheduling")
	cmd := &cobra.Command{
		Use:   "scheduler",
		Short: "auto Absence-Me!",
		RunE: func(cmd *cobra.Command, args []string) error {
			path := os.Getenv("ABSENCE_PATH")
			if len(path) == 0 {
				path = "/tmp/absence.yaml"
			}

			cConfig := absence.NewFromPath(path)
			client := absence.New(cConfig, log)

			config := scheduler.NewFromPath(path)
			schedulerSvc := scheduler.New(config, client, log)
			schedulerSvc.Process()
			return nil
		},
	}
	return cmd
}
