package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	"expvar"
	_ "net/http/pprof"
)

var hits = expvar.NewInt("hits")

func main() {
	http.HandleFunc("/fub", fub)

	http.HandleFunc("/debug", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(404)
		fmt.Fprintln(w, "Available handlers:")
		fmt.Fprintln(w, "/debug/pprof/")
		fmt.Fprintln(w, "/debug/pprof/cmdline")
		fmt.Fprintln(w, "/debug/pprof/profile")
		fmt.Fprintln(w, "/debug/pprof/symbol")
		fmt.Fprintln(w, "/debug/pprof/trace")
		fmt.Fprintln(w, "/debug/vars")
	})

	log.Fatal(http.ListenAndServe(":9999", nil))
}

func fub(w http.ResponseWriter, r *http.Request) {
	hits.Add(1)

	for i := 0; i < 50000000; i++ {
		now := time.Now()
		seconds := now.Hour()*3600 + now.Minute()*60 + now.Second()

		var buf bytes.Buffer

		buf.WriteString("seconds so far: ")
		buf.WriteString(strconv.Itoa(seconds))

		buf.WriteTo(ioutil.Discard)
	}
	w.Write([]byte("fub done\n"))
}
