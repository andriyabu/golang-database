package golangdatabase

import (
	"database/sql"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestEmpty(t *testing.T) {

}

func TestOpenConnection(t *testing.T) {
	db, err := sql.Open("mysql", "abeleon:showmethecode@tcp(localhost:3306)/golangdatabase")

	if err != nil {
		panic(err)
	}

	defer db.Close()
}
