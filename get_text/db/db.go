package db

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type DB_info struct {
	s_user     string
	s_pwd      string
	s_host     string
	s_port     string
	s_database string
	s_engine   string
}

func (d *DB_info) Get_text() error {
	file, err := ioutil.ReadFile("db_info.txt")
	if err != nil {
		return err
	}

	txt := string(file)
	sliceTxt := strings.Split(txt, "\n")

	for _, v := range sliceTxt {
		slice := strings.Split(v, "=")
		key := strings.TrimSpace(slice[0])
		value := strings.TrimSpace(slice[1])
		fmt.Printf("%s, %s\n", key, value)

		switch key {
		case "user":
			d.s_user = value
		case "pwd":
			d.s_pwd = value
		case "host":
			d.s_host = value
		case "port":
			d.s_port = value
		case "database":
			d.s_database = value
		case "engine":
			d.s_engine = value
		}
	}
	return nil
}
