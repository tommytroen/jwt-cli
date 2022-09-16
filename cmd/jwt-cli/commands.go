package main

import (
	"github.com/spf13/cobra"
)

type Interface interface {
	// AddFlags adds this options' flags to the cobra command.
	AddFlags(cmd *cobra.Command)
}

func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:               "jwt",
		Short:             "jwt",
		DisableAutoGenTag: true,
		SilenceUsage:      true, // Don't show usage on errors
	}
	// Add sub-commands.
	cmd.AddCommand(Sign())
	cmd.AddCommand(Verify())
	return cmd
}
