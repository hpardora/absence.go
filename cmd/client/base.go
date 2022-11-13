package client

import "github.com/spf13/cobra"

func NewCMDCli() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "Absence.io",
		Short: "absence Root command",
	}

	cmd.AddCommand(NewCmdPrinter())
	cmd.AddCommand(NewCmdScheduler())
	return cmd
}
