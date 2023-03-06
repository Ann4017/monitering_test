package main

import (
	"get_text/db"
)

func main() {
	info := db.DB_info{}
	info.Get_text()
}
