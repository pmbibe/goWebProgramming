package main

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func sqlxTest() {
	db, err := sqlx.Connect("mysql", connecttoDB("ducptm", "Password@Aa123", "192.168.141.241", "3306", "gwp"))
	if err != nil {
		log.Fatalln(err)
	}

	user := []Post{}

	db.Select(&user, "select * from posts")

	log.Println("users...")
	log.Println(user)

}
