package storage

import (
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
)

//SetDSN sets the DSN used by the storage package.
func SetDSN(dsn string) {
	connectionDSN = dsn
}

// MultiExec is a helper function for running multiple queries.
func MultiExec(e sqlx.Execer, query string) {
	stmts := strings.Split(query, ";\n")
	if len(strings.Trim(stmts[len(stmts)-1], " \n\t\r")) == 0 {
		stmts = stmts[:len(stmts)-1]
	}
	for _, s := range stmts {
		_, err := e.Exec(s)
		if err != nil {
			fmt.Println(err, s)
		}
	}
}
