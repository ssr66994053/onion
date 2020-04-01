package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ssr66994053/carrot"
)

func main() {
	fmt.Println("carrot starting ...")

	c := carrot.New()
	c.Get("/hello", func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("from /hello"))
	})
	c.Post("/p", func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("from /p"))
	})

	log.Fatal(c.Start(":3000"))
}
