/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package main

import (
	"log"

	"github.com/ramadhanalfarisi/sql-excel-cli/cmd"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	cmd.Execute()
}
