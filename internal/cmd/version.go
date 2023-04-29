package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewCmdVersion(version string) *cobra.Command {
	cmd := &cobra.Command{
		Use: "version",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Fprint(cmd.OutOrStdout(), version)
		},
	}
	return cmd
}
