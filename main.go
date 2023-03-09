package main

import (
	"database/sql"
	"fmt"
	"log"
)

func main() {
	info := &DB_info{}

	err := info.Get_ini()
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}

	dataSource := info.s_user + ":" + info.s_pwd + "@tcp(" + info.s_host + ":" + info.s_port + ")/" + info.s_database
	db, err := sql.Open(info.s_engine, dataSource)
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
	defer db.Close()

	fmt.Println(dataSource)
}
