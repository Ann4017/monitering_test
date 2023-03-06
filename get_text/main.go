package main

import (
	"database/sql"
	"fmt"
	"get_text/db"
	"log"
)

func main() {
	info := db.DB_info{}
	err := info.Get_text()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(info)

	dataSource := info.s_user + ":" + info.s_pwd + "@tcp(" + info.s_host + ":" + info.s_port + ")/" + info.s_database
	db, err := sql.Open(info.s_engine, dataSource)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
}
