package belajargolangdatabase

import (
	"testing"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func TestEmpty(t *testing.T) {

}

func TestOpenConnection(t *testing.T) {
	db, err := sql.Open("mysql", "root:@rootname(localhost:3306)/database_name")
if err != nil {
	panic(err)
}
defer db.Close() // akan menutup jika sudah selesai
}

