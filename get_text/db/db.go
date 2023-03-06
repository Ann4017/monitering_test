package db

import (
	"fmt"
	"io/ioutil"
)

type DB_info struct {
	s_uesr     string
	s_pwd      string
	s_host     string
	s_port     string
	s_database string
	s_engine   string
}

func (db *DB_info) Get_text() {
	file, err := ioutil.ReadFile("db_info.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(file))
}
