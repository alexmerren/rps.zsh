package cmd

import "github.com/spf13/cobra"

func NewCmdRoot(version string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "rps",
		Short: "Repo Selector CLI",
		Long:  `Clone your favourite repositories with ease`,
	}

	versionCommand := NewCmdVersion(version)
	menuCommand := NewCmdMenu()
	cmd.AddCommand(versionCommand)
	cmd.AddCommand(menuCommand)
	return cmd
}
