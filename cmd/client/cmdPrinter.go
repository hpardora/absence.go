package client

import (
	"github.com/hpardora/absence.go/internal/service/printer"
	"github.com/hpardora/absence.go/pkg/absence"
	"github.com/hpardora/absence.go/pkg/logger"
	"github.com/spf13/cobra"
	"os"
)

func NewCmdPrinter() *cobra.Command {
	log := logger.GetLogger()
	log.Infof("starting printer")
	cmd := &cobra.Command{
		Use:   "print",
		Short: "retrieve my Absence data",
		RunE: func(cmd *cobra.Command, args []string) error {
			path := os.Getenv("ABSENCE_PATH")
			if len(path) == 0 {
				path = "/tmp/absence.yaml"
			}

			cConfig := absence.NewFromPath(path)
			client := absence.New(cConfig, log)

			printerSVC := printer.New(client, log)
			printerSVC.Process()
			return nil
		},
	}

	return cmd
}
