package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
)

func main() {
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt)

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "OK")
	})

	go func() {
		if err := http.ListenAndServe("0.0.0.0:8080", nil); err != nil {
			panic(err)
		}
	}()
	<-done
}
