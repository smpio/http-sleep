package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"
)

var usage = "Usage:\nGET /10s\nGET /1m\nGET /3h\n"

func main() {
	port := 80

	flag.IntVar(&port, "port", port, "Port to listen on")
	flag.Parse()

	http.HandleFunc("/", serve)

	httpServer := &http.Server{
		Addr: fmt.Sprint(":", port),
	}

	fmt.Print("Listening on port ", port, "\n", usage)
	httpServer.ListenAndServe()
}

func serve(w http.ResponseWriter, r *http.Request) {
	duration, err := time.ParseDuration(r.RequestURI[1:])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Invalid duration specifier in the URL.\n", usage)
	} else {
		time.Sleep(duration)
		fmt.Fprint(w, "OK")
	}
}
