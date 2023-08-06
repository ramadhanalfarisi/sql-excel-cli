package cmd

import (
	"github.com/ramadhanalfarisi/sql-excel-cli/controller"
	"github.com/spf13/cobra"
)

var infoController = controller.NewInfoController()
var infoCommand = &cobra.Command{
	Use:   "info",
	Short: "This command used to get the information table",
	Long:  "This command used to get the information of all tables in sql file",
	Run:   infoController.InfoCommand,
}
