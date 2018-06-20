package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"github.com/julienschmidt/httprouter"
)

var (
	addr = flag.String("addr", ":8080", "http service address")
	data map[string]string
)

func main() {
	flag.Parse()
	data = map[string]string{}
	r := httprouter.New()
	r.GET("/list", show)
	r.GET("/list/:key", show)
	r.PUT("/list/:key/:value", update)
	err := http.ListenAndServe(*addr, r)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func show(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	k := p.ByName("key")
	if k == "" {
		for i, v := range data {
			fmt.Fprintf(w, "ID %s : %s \n", i, v)
		}
		return
	}
	fmt.Fprintf(w, "ID %s : %s", k, data[k])
}

func update(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	k := p.ByName("key")
	v := p.ByName("value")

	data[k] = v

	fmt.Fprintf(w, "[UPDATED] ID %s : %s", k, data[k])
}