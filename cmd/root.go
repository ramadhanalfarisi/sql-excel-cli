/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"os"

	"github.com/ramadhanalfarisi/sql-excel-cli/db"
	"github.com/spf13/cobra"
)

var Dbroot = db.ConnectDB()

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "sql-excel-cli",
	Short: "Hello, this is my first CLI project",
	Long:  "Hello mate, this is my first CLI project, this project used to convert sql file to excel",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	infoCommand.Flags().String("file", "", "Add SQL file")
	infoCommand.Flags().String("dir", "", "Add SQL directory")

	excelCommand.Flags().String("file", "", "Add SQL file")
	excelCommand.Flags().String("dir", "", "Add SQL directory")
	excelCommand.Flags().String("dest", "", "Destination excel file")
	rootCmd.AddCommand(infoCommand)
	rootCmd.AddCommand(excelCommand)
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.sql-excel-cli.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
