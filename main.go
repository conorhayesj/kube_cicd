package main

import (
	"net/http"
	"time"
	"fmt"
)

func greeting(w http.ResponseWriter, r *http.Request) {
	dt := time.Now()
	fmt.Fprintf(w, "The current date and time is:\n%s", dt.Format("02-01-2006 15:04:05 Monday"))
}

func main() {
	http.HandleFunc("/", greeting)

	http.ListenAndServe(":8080", nil)
}
