package cmd

import (
	"github.com/ramadhanalfarisi/sql-excel-cli/controller"
	"github.com/ramadhanalfarisi/sql-excel-cli/model"
	"github.com/spf13/cobra"
)

var infoModel = model.InfoModel{DB: Dbroot}
var infoController = controller.NewInfoController(infoModel)
var infoCommand = &cobra.Command{
	Use:   "info",
	Short: "This command used to get the information table",
	Long:  "This command used to get the information of all tables in sql file",
	Run:   infoController.InfoCommand,
}
