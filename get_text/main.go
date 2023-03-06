package main

import (
	"fmt"
	"get_text/db"
)

func main() {
	info := db.DB_info{}
	err := info.Get_text()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(info)
}
