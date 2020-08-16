package db

import (
	"database/sql"
	_"github.com/go-sql-driver/mysql"
	"Inventory_Project/config"
)

func Connect() *sql.DB {
	db,err := sql.Open("mysql",config.DB_USER+":"+config.DB_PASS+"@tcp("+config.DB_HOST+":3306)/"+config.DB_NAME)
	if err != nil {
		panic(err.Error())
	}
	return db
}