package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func connecttoDB(user, password, hostname, port, database string) string {
	return user + ":" + password + "@" + "tcp(" + hostname + ":" + port + ")/" + database

}
func connect() *sql.DB {
	db, err := sql.Open("mysql", connecttoDB("ducptm", "Password@Aa123", "192.168.141.241", "3306", "gwp"))
	errHandler(err)
	return db
}
