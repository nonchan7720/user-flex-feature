package cmd

import "github.com/spf13/cobra"

func rootCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:           "user-flex-feature",
		SilenceErrors: true,
		SilenceUsage:  true,
	}
	cmd.AddCommand(serverCommand())
	return cmd
}
