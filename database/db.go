package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Init(){
	var err error
	DB, err = sql.Open("mysql","root:@/godatabase")

	if err != nil{
		panic(err.Error())
	}
}