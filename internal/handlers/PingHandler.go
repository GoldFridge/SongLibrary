package handlers

import (
	"fmt"
	"log"
	"net/http"
)

func PingHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, err := fmt.Fprint(w, "ok")
	if err != nil {
		log.Fatal(err)
	}
}
