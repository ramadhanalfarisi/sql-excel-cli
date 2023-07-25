package model

import (
	"database/sql"
	"log"
	"sync"
)

type DataTable struct {
	Name    string
	Record  int
	Columns []ColumnTable
	Data    []map[string]interface{}
}

type DataModel struct {
	DB *sql.DB
}

func (m *DataModel) GetData(tables []string) (chan DataTable, error) {
	var wg sync.WaitGroup
	chanTables := make(chan DataTable)
	for _, table := range tables {
		wg.Add(1)
		go func(table string) {
			var dataTable DataTable
			dataTable.Name = table
			rows, err := m.DB.Query("SELECT * FROM "+table)
			if err != nil {
				log.Fatal(err)
			}
			columns, err := rows.ColumnTypes()
			if err != nil {
				log.Fatal(err)
			}
			for _, column := range columns {
				var columnTable ColumnTable
				columnTable.Name = column.Name()
				columnTable.TypeData = column.DatabaseTypeName()
				if lengthType, ok:= column.Length(); ok {
					columnTable.Length = lengthType
				}else {
					columnTable.Length = 0
				}
				dataTable.Columns = append(dataTable.Columns, columnTable)
			}
			for rows.Next() {
				cols := make([]interface{}, len(columns))
				colPointer := make([]interface{}, len(columns))
				for i, _ := range cols {
					colPointer[i] = &cols[i]
				}

				err := rows.Scan(colPointer...)
				if err != nil {
					log.Fatal(err)
				}

				mp := make(map[string]interface{})
				for i, col := range dataTable.Columns {
					value := colPointer[i].(*interface{})
					mp[col.Name] = *value
				}

				dataTable.Data = append(dataTable.Data, mp)
			}
			dataTable.Record = len(dataTable.Data)
			chanTables <- dataTable
			wg.Done()
		}(table)
	}

	go func() {
		wg.Wait()
		close(chanTables)
	}()

	return chanTables, nil
}
