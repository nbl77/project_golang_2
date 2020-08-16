package db

import (
	"database/sql"
)

type Mysql struct {
	rows *sql.Rows
}
func SelectAll(query string) *sql.Rows{
	db := Connect()
	rows, err := db.Query(query)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	return rows
}

func Execute(query string, stmt ...interface{}) (int64, error) {
	db := Connect()
	exe,err := db.Exec(query, stmt...)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	return exe.RowsAffected()
}
func Select(query string) *sql.Row{
	db := Connect()
	row := db.QueryRow(query)
	defer db.Close()
	return row
}
func SelectParam(query string,stmt ...interface{}) *sql.Row{
	db := Connect()
	row := db.QueryRow(query,stmt...)
	defer db.Close()
	return row
}