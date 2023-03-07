package main

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/ini.v1"
)

type DB_info struct {
	s_user     string
	s_pwd      string
	s_host     string
	s_port     string
	s_database string
	s_engine   string
}

func (d *DB_info) Get_ini() error {
	_, err := os.Stat("config.ini")
	if err != nil {
		file_path := "config.ini"
		file, err := os.Create(file_path)
		if err != nil {
			fmt.Println(err)
			return err
		}
		defer file.Close()
	}

	file, err := ini.Load("config.ini")
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}

	sec := file.Section("dbinfo")

	d.s_user = sec.Key("user").String()
	d.s_pwd = sec.Key("pwd").String()
	d.s_host = sec.Key("host").String()
	d.s_port = sec.Key("port").String()
	d.s_database = sec.Key("database").String()
	d.s_engine = sec.Key("engine").String()

	return err
}
