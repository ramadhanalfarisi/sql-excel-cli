package controller

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/ramadhanalfarisi/sql-excel-cli/helper"
	"github.com/ramadhanalfarisi/sql-excel-cli/model"
	"github.com/spf13/cobra"
)

type InfoController struct {
	Model model.InfoModel
}

func NewInfoController(model model.InfoModel) InfoControllerInterface {
	return &InfoController{Model: model}
}

func (c *InfoController) GetTableInfo(file string, dir string) ([]model.InfoTable, error) {
	info, err := c.CreateTable(file, dir)
	if err != nil {
		return nil, err
	}
	tables, err := c.Model.GetColumn(info.Tables)
	if err != nil {
		return nil, err
	}
	var infoTables []model.InfoTable
	for table := range tables {
		infoTables = append(infoTables, table)
	}
	return infoTables, nil
}

func (c *InfoController) CreateTable(file string, dir string) (*helper.SqlFile, error) {
	sqlHelper := helper.SqlNew()
	if len(file) > 0 && len(dir) > 0 {
		return nil, fmt.Errorf("Just fill one (file or dir)")
	}else{
		if len(file) > 0 {
			err := sqlHelper.Filex(file)
			if err != nil {
				return nil, err
			}
		} else if len(dir) > 0 {
			err := sqlHelper.Directory(dir)
			if err != nil {
				return nil, err
			}
		} else {
			return nil, fmt.Errorf("file or dir flag have to filled")
		}
	}
	
	err2 := sqlHelper.DropTables(c.Model.DB)
	if err2 != nil {
		return nil, err2
	}
	_, err3 := sqlHelper.Exec(c.Model.DB)
	if err3 != nil {
		return nil, err3
	}
	return sqlHelper, nil
}

func (c *InfoController) InfoCommand(cmd *cobra.Command, args []string) {
	sqlFlags, err := cmd.Flags().GetString("file")
	if err != nil {
		log.Fatal(err)
	}
	dirFlags, err := cmd.Flags().GetString("dir")
	if err != nil {
		log.Fatal(err)
	}
	tables, err := c.GetTableInfo(sqlFlags, dirFlags)
	if err != nil {
		log.Fatal(err)
	}
	json, err := json.MarshalIndent(tables, "", " ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(json))
}
