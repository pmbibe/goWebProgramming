package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func connecttoDB(user, password, hostname, port, database string) string {
	return user + ":" + password + "@" + "tcp(" + hostname + ":" + port + ")/" + database + "?parseTime=true"

}
func connect() {
	Db, _ = gorm.Open(mysql.Open(connecttoDB("ducptm", "Password@Aa123", "192.168.141.241", "3306", "gorm")), &gorm.Config{})

}
