package cmd

import (
	"github.com/ramadhanalfarisi/sql-excel-cli/controller"
	"github.com/spf13/cobra"
)

var excelController = controller.NewExcelController()
var excelCommand = &cobra.Command{
	Use:   "excel",
	Short: "This command used to convert sql file to excel file",
	Long:  "This command used to convert all tables in sql file to excel file",
	Run:   excelController.ExcelCommand,
}
