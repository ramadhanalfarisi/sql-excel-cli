package controller

import (
	"github.com/ramadhanalfarisi/sql-excel-cli/helper"
	"github.com/ramadhanalfarisi/sql-excel-cli/model"
	"github.com/spf13/cobra"
)

type ExcelControllerInterface interface {
	GetData(string, string) ([]model.DataTable, error)
	CreateTable(string, string) (*helper.SqlFile, error)
	ExcelCommand(cmd *cobra.Command, args []string)
}