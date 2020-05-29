package main

import (
	"encoding/json"
	"log"
	"net"
	"net/http"
	"os"
)

var sockPath = "./test.sock"

func respondOk(w http.ResponseWriter, body interface{}) {
	w.Header().Set("Content-Type", "application/json")

	if body == nil {
		w.WriteHeader(http.StatusNoContent)
	} else {
		w.WriteHeader(http.StatusOK)
		enc := json.NewEncoder(w)
		enc.Encode(body)
	}
}

type status struct {
	Status string `json:"status"`
}

func handlerStatus() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		respondOk(w, &status{
			Status: "ok",
		})
	})
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/status", handlerStatus())

	os.Remove(sockPath)

	l, err := net.Listen("unix", sockPath)
	if err != nil {
		log.Fatal(err)
	}

	log.Fatal(http.Serve(l, mux))
}
