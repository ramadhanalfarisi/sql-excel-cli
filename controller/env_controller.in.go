package controller

import "github.com/spf13/cobra"

type EnvControllerInterface interface {
	SetEnvironment(cmd *cobra.Command, args []string)
}