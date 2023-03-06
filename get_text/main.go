package main

import (
	"monitering/db"
	"monitering/http"
)

func main() {
	http.Test1()
	db := db.DB_info{}
}
