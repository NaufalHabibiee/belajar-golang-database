package belajargolangdatabase

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func GetConnection() *sql.DB {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/belajar_golang_database?parseTime=true")
	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)                 // maksimal open koneksi nya 100
	db.SetConnMaxIdleTime(5 * time.Minute)  // saat 5 menit ada koneksi yang nganggur, koneksi tersebut akan di close
	db.SetConnMaxLifetime(60 * time.Minute) // kalau udah 60 menit koneksi apapun akan di perbarui dengan yang baru atau di ganti dengan yang baru

	return db
}
