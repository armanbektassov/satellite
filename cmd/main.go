package main

import (
	"net/http"
	"os"
	"os/signal"
)

func main() {
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt)

	go func() {
		http.ListenAndServe("0.0.0.0:6060", nil)
	}()
	<-done
}
