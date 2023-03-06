package db

import (
	"bufio"
	"fmt"
	"os"
)

type DB_info struct {
	s_id       string
	s_pwd      string
	s_url      string
	s_engine   string
	s_database string
}

func (t *DB_info) Get_db_info() {
	file, err := os.Open("db_info.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	text := bufio.NewScanner(file)
	for text.Scan() {
		fmt.Println(text.Text())
	}
	if err := text.Err(); err != nil {
		panic(err)
	}
}
