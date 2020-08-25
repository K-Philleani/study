package web

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func Fetch() {
	res, err := http.Get("https://books.studygolang.com/the-way-to-go_ZH_CN/15.3.html")
	if err != nil {
		log.Fatalf("Get : %v", err)
		return
	}
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("Got : %v", err)
		return
	}
	fmt.Printf("Got: %q", string(data))
}

