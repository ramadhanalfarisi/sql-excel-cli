package cmd

import (
	"github.com/ramadhanalfarisi/sql-excel-cli/controller"
	"github.com/ramadhanalfarisi/sql-excel-cli/model"
	"github.com/spf13/cobra"
)

var excelModel = model.DataModel{DB: Dbroot}
var excelController = controller.NewExcelController(excelModel)
var excelCommand = &cobra.Command{
	Use:   "excel",
	Short: "This command used to convert sql file to excel file",
	Long:  "This command used to convert all tables in sql file to excel file",
	Run:   excelController.ExcelCommand,
}
