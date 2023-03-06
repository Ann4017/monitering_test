package main

import (
	"bufio"
	"fmt"
	"net/http"
)

type https struct {
}

func main() {
	t := https{}
	err := t.get_status_boby("https://www.google.com/?hl=ko")
	fmt.Println("error:", err)
}

func (t *https) get_status_boby(url string) error {
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	fmt.Println("status:", response.Status)

	header := bufio.NewScanner(response.Body)
	for i := 0; header.Scan() && i < 5; i++ {
		fmt.Println(header.Text())
	}
	if err := header.Err(); err != nil {
		panic(err)
	}
	return nil
}
