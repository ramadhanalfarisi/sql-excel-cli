package controller

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

type EnvController struct{}

func NewEnvController() EnvControllerInterface {
	return &EnvController{}
}

// SetEnvironment implements EnvControllerInterface.
func (*EnvController) SetEnvironment(cmd *cobra.Command, args []string) {
	host, err := cmd.Flags().GetString("host")
	if err != nil {
		log.Fatal(err)
	}
	uname, err := cmd.Flags().GetString("username")
	if err != nil {
		log.Fatal(err)
	}
	pass, err := cmd.Flags().GetString("password")
	if err != nil {
		log.Fatal(err)
	}
	port, err := cmd.Flags().GetString("port")
	if err != nil {
		log.Fatal(err)
	}
	dbname, err := cmd.Flags().GetString("dbname")
	if err != nil {
		log.Fatal(err)
	}
	env := fmt.Sprintf(`SQL_EXCEL_CLI_HOST=%s
SQL_EXCEL_CLI_USERNAME=%s
SQL_EXCEL_CLI_PASSWORD=%s
SQL_EXCEL_CLI_PORT=%s
SQL_EXCEL_CLI_NAME=%s`, host, uname, pass, port, dbname)
	err2 := os.WriteFile(".env", []byte(env), 0644)
	if err != nil {
		log.Fatal(err2)
	}
}
