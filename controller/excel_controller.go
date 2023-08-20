package controller

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"unicode/utf8"

	"github.com/ramadhanalfarisi/sql-excel-cli/db"
	"github.com/ramadhanalfarisi/sql-excel-cli/helper"
	"github.com/ramadhanalfarisi/sql-excel-cli/model"
	"github.com/spf13/cobra"
	"github.com/xuri/excelize/v2"
)

type ExcelController struct{}

func NewExcelController() ExcelControllerInterface {
	return &ExcelController{}
}

// CreateTable implements ExcelControllerInterface.
func (c *ExcelController) CreateTable(file string, dir string) (*helper.SqlFile, error) {
	connDB := db.ConnectDB()
	sqlHelper := helper.SqlNew()
	if len(file) > 0 && len(dir) > 0 {
		return nil, fmt.Errorf("just fill one (file or dir)")
	} else {
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
	err2 := sqlHelper.DropTables(connDB)
	if err2 != nil {
		return nil, err2
	}

	_, err3 := sqlHelper.Exec(connDB)
	if err3 != nil {
		return nil, err3
	}
	return sqlHelper, nil
}

// GetData implements ExcelControllerInterface.
func (c *ExcelController) GetData(file string, dir string) ([]model.DataTable, error) {
	sqlFile, err := c.CreateTable(file, dir)
	dataModel := model.DataModel{DB: db.ConnectDB()}
	if err != nil {
		return nil, err
	}
	chanData, err := dataModel.GetData(sqlFile.Tables)
	if err != nil {
		return nil, err
	}
	var dataTables []model.DataTable
	for data := range chanData {
		dataTables = append(dataTables, data)
	}
	return dataTables, nil
}

// InfoCommand implements ExcelControllerInterface.
func (c *ExcelController) ExcelCommand(cmd *cobra.Command, args []string) {
	sqlFlags, err := cmd.Flags().GetString("file")
	if err != nil {
		log.Fatal(err)
	}
	dirFlags, err := cmd.Flags().GetString("dir")
	if err != nil {
		log.Fatal(err)
	}
	destFlags, err := cmd.Flags().GetString("dest")
	if err != nil {
		log.Fatal(err)
	}
	tables, err := c.GetData(sqlFlags, dirFlags)
	if err != nil {
		log.Fatal(err)
	}
	err2 := c.ExcelGenerator(tables, destFlags)
	if err2 != nil {
		log.Fatal(err2)
	}
	fmt.Println("Generate excel successfully")
}

func (c *ExcelController) ExcelGenerator(tables []model.DataTable, dest string) error {
	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	for itab, table := range tables {
		if itab == 0 {
			err := f.SetSheetName("Sheet1", table.Name)
			if err != nil {
				log.Fatal(err)
			}
		} else {
			_, err := f.NewSheet(table.Name)
			if err != nil {
				log.Fatal(err)
			}
		}

		chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
		for icol, col := range table.Columns {
			var char string
			if icol <= 25 {
				char = string(chars[icol])
			} else {
				div := math.Floor(float64(icol) / 25)
				mod := icol % 25
				char = string(chars[int(div)])
				char += string(chars[mod])
			}
			f.SetCellValue(table.Name, char+"1", col.Name)
		}
		for iitem, item := range table.Data {
			for icol, col := range table.Columns {
				var char string
				if icol <= 25 {
					char = string(chars[icol])
				} else {
					div := math.Floor(float64(icol) / 25)
					mod := icol % 25
					char = string(chars[int(div)])
					char += string(chars[mod])
				}
				num := strconv.Itoa(iitem + 2)
				lengthCol := utf8.RuneCountInString(string(item[col.Name].([]byte)[:])) + 2
				if int64(lengthCol) > col.Length {
					table.Columns[icol].Length = int64(lengthCol)
				}
				f.SetCellValue(table.Name, char+num, item[col.Name].([]byte)[:])
			}
			helper.PrintProgressBar(iitem+1, len(table.Data), "Progress "+table.Name, "Complete", 25, "=")
		}
		for icol, col := range table.Columns {
			var char string
			if icol <= 25 {
				char = string(chars[icol])
			} else {
				div := math.Floor(float64(icol) / 25)
				mod := icol % 25
				char = string(chars[int(div)])
				char += string(chars[mod])
			}
			f.SetColWidth(table.Name, char, char, float64(col.Length))
		}
	}
	if err := f.SaveAs(dest); err != nil {
		log.Fatal(err)
	}
	return nil
}
