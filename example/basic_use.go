package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ssr66994053/onion"
)

func main() {
	fmt.Println("onion starting ...")

	c := onion.New(onion.DefaultConfig())
	c.GetFunc("/hello", func(w http.ResponseWriter, req *http.Request, params map[string]string) {
		w.Write([]byte("from /hello"))
	})
	c.PostFunc("/p", func(w http.ResponseWriter, req *http.Request, params map[string]string) {
		w.Write([]byte("from /p"))
	})

	log.Fatal(c.Start(":3000"))
}
