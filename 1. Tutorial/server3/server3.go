package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"go-exercises/lissajousFunc"
	"strconv"
)

var mu sync.Mutex
var count int

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			log.Print(err)
		}

		c, _ := strconv.ParseFloat(r.Form.Get("c"), 64)
		res, _ := strconv.ParseFloat(r.Form.Get("r"), 64)
		s, _ := strconv.Atoi(r.Form.Get("s"))
		n, _:= strconv.Atoi(r.Form.Get("n"))
		d, _ := strconv.Atoi(r.Form.Get("d"))
		lissajousFunc.Lissajous(w, lissajousFunc.LissajousOpts{Cycles: c, Res: res, Size: s, Nframes: n, Delay: d})
	})
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q", k, v)
	}
	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}
}