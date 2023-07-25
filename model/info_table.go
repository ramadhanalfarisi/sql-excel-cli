package model

import (
	"database/sql"
	"log"
	"sync"
)

type InfoModel struct {
	DB *sql.DB
}

type InfoTable struct {
	Name   string        `json:"name"`
	Record int64         `json:"record"`
	Column []ColumnTable `json:"columns"`
}

type ColumnTable struct {
	Name     string `json:"name"`
	TypeData string `json:"type_data"`
	Length   int64  `json:"length"`
}

func (m *InfoModel) GetColumn(tables []string) (chan InfoTable, error) {
	chanTables := make(chan InfoTable)
	var wg sync.WaitGroup
	for _, tab := range tables {
		wg.Add(1)
		go func(tab string) {
			defer wg.Done()
			var infoTable InfoTable
			infoTable.Name = tab
			rows, err := m.DB.Query("SELECT * FROM " + tab + " LIMIT 1")
			if err != nil {
				log.Fatal(err)
			}
			columnTypes, err := rows.ColumnTypes()
			if err != nil {
				log.Fatal(err)
			}
			for _, col := range columnTypes {
				var colTable ColumnTable
				colTable.Name = col.Name()
				colTable.TypeData = col.DatabaseTypeName()
				if length, ok := col.Length(); ok {
					colTable.Length = length
				} else {
					colTable.Length = 0
				}
				infoTable.Column = append(infoTable.Column, colTable)
			}
			countRows, err := m.DB.Query("SELECT COUNT(*) FROM " + tab)
			if err != nil {
				log.Fatal(err)
			}
			for countRows.Next() {
				var count int64
				err := countRows.Scan(&count)
				if err != nil {
					log.Fatal(err)
				}
				infoTable.Record = count
			}
			chanTables <- infoTable
		}(tab)
	}
	go func() {
		wg.Wait()
		close(chanTables)
	}()
	return chanTables, nil
}
