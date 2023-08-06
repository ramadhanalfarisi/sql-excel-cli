/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/ramadhanalfarisi/sql-excel-cli/controller"
	"github.com/spf13/cobra"
)

// envCmd represents the env command
var envController = controller.NewEnvController() 
var envCmd = &cobra.Command{
	Use:   "env",
	Short: "This command to set the application environment",
	Long: `This command is used to set database environment consisting of :
	- host
	- username
	- password
	- port
	- database name`,
	Run: envController.SetEnvironment,
}
