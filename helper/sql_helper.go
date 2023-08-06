package helper

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
)

type SqlFile struct {
	Files   []string
	Queries []string
	Tables  []string
}

func SqlNew() *SqlFile {
	return &SqlFile{}
}

func (s *SqlFile) FileString(file string) string {
	queries, tables, err := load(file)
	if err != nil {
		return ""
	}

	s.Files = append(s.Files, file)
	s.Queries = append(s.Queries, queries...)
	s.Tables = append(s.Tables, tables...)

	json, _ := json.Marshal(s.Queries)
	return string(json)
}

func (s *SqlFile) Filex(file string) error {
	queries, tables, err := load(file)
	if err != nil {
		return err
	}

	s.Files = append(s.Files, file)
	s.Queries = append(s.Queries, queries...)
	s.Tables = append(s.Tables, tables...)

	return nil
}

func (s *SqlFile) Filexs(files ...string) error {
	for _, file := range files {
		if err := s.Filex(file); err != nil {
			return err
		}
	}
	return nil
}

func (s *SqlFile) Directory(dir string) error {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return err
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		name := file.Name()
		if name[len(name)-3:] != "sql" {
			continue
		}

		if err := s.Filex(dir + "/" + name); err != nil {
			return err
		}
	}

	return nil
}

func (s *SqlFile) Exec(db *sql.DB) (res []sql.Result, err error) {
	tx, err := db.Begin()
	if err != nil {
		return res, err
	}
	defer saveTx(tx, &err)

	var rs []sql.Result
	for _, q := range s.Queries {
		r, err := tx.Exec(q)
		if err != nil {
			return res, fmt.Errorf(err.Error() + " : when executing > " + q)
		}
		rs = append(rs, r)
	}

	return rs, err
}

func (s *SqlFile) DropTables(db *sql.DB) (err error) {
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	for _, t := range s.Tables {
		query := fmt.Sprintf("DROP TABLE IF EXISTS %s", t)
		_, err := tx.Exec(query)
		if err != nil {
			return fmt.Errorf(err.Error() + " : when executing > " + query)
		}
	}

	return err
}

func load(path string) (qs []string, tbs []string, err error) {
	ls, err := readFileByLine(path)
	if err != nil {
		return qs, tbs, err
	}

	var ncls []string
	for _, l := range ls {
		if table := table(l); len(table) > 0 {
			tbs = append(tbs, table)
		}
		if ncl := isComment(l); !ncl {
			ncls = append(ncls, l)
		}
	}

	l := strings.Join(ncls, "")
	l = strings.ReplaceAll(l, "\t", "")
	l = strings.ReplaceAll(l, "\r", "")
	qs = strings.Split(l, ";")
	qs = qs[:len(qs)-1]

	return qs, tbs, nil
}

func readFileByLine(path string) (ls []string, err error) {
	f, err := ioutil.ReadFile(path)
	if err != nil {
		return ls, err
	}

	ls = strings.Split(string(f), "\n")
	return ls, nil
}

func isComment(line string) bool {
	d := "/*"
	c := "--"
	if len(line) > 1 {
		return line[:2] == d || line[:2] == c
	} else {
		return false
	}
}

func table(line string) string {
	str := "create table"
	line = strings.ReplaceAll(line, "/t", "")
	line = strings.ReplaceAll(line, "/r", "")
	line = strings.ReplaceAll(line, "`", "")
	line = strings.ToLower(line)
	if isContains := strings.Contains(line, str); isContains {
		arrLine := strings.Split(line, " ")
		return arrLine[2]
	}
	return ""
}

func saveTx(tx *sql.Tx, err *error) {
	if p := recover(); p != nil {
		tx.Rollback()
		panic(p)
	} else if *err != nil {
		tx.Rollback()
	} else {
		e := tx.Commit()
		err = &e
	}
}
