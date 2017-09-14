package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/fcgi"
	"os"
	"runtime"
	"strconv"
)

var app_addr string
var index int

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	app_addr = os.Getenv("APP_ADDR") // e.g. "0.0.0.0:8080" or ""
}

func ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var body = ""
	if b, err := ioutil.ReadAll(r.Body); err == nil {
		body = strconv.Quote(string(b))
	}

	headers := w.Header()
	headers.Add("Content-Type", "text/html")
	io.WriteString(w, "<html><head></head><body><p>Hello world from Go!</p><table>")
	io.WriteString(w, fmt.Sprintf("<tr><td>Method</td><td>%s</td></tr>", r.Method))
	io.WriteString(w, fmt.Sprintf("<tr><td>URL</td><td>%s</td></tr>", r.URL))
	io.WriteString(w, fmt.Sprintf("<tr><td>URL.Path</td><td>%s</td></tr>", r.URL.Path))
	io.WriteString(w, fmt.Sprintf("<tr><td>Proto</td><td>%s</td></tr>", r.Proto))
	io.WriteString(w, fmt.Sprintf("<tr><td>Host</td><td>%s</td></tr>", r.Host))
	io.WriteString(w, fmt.Sprintf("<tr><td>RemoteAddr</td><td>%s</td></tr>", r.RemoteAddr))
	io.WriteString(w, fmt.Sprintf("<tr><td>RequestURI</td><td>%s</td></tr>", r.RequestURI))
	io.WriteString(w, fmt.Sprintf("<tr><td>Header</td><td>%s</td></tr>", r.Header))
	io.WriteString(w, fmt.Sprintf("<tr><td>Body</td><td>%s</td></tr>", body))
	io.WriteString(w, fmt.Sprintf("<tr><td>Request Index</td><td>%d</td></tr>", index))
	io.WriteString(w, "</table></body></html>")
	index = index + 1
}

func main() {
	http.HandleFunc("/", ServeHTTP)

	var err error
	if app_addr != "" { // Run as a local web server
		err = http.ListenAndServe(app_addr, nil)
	} else { // Run as FCGI via standard I/O
		err = fcgi.Serve(nil, nil)
	}
	if err != nil {
		log.Fatal(err)
	}
}
