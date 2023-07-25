package controller

import (
	"github.com/ramadhanalfarisi/sql-excel-cli/helper"
	"github.com/ramadhanalfarisi/sql-excel-cli/model"
	"github.com/spf13/cobra"
)

type InfoControllerInterface interface {
	GetTableInfo(string, string) ([]model.InfoTable, error)
	CreateTable(string, string) (*helper.SqlFile, error)
	InfoCommand(cmd *cobra.Command, args []string)
}
