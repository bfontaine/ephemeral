package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/bfontaine/ephemeral"
)

func done(s *ephemeral.Server, w http.ResponseWriter, r *http.Request) {
	log.Printf("Got a request")
	w.Write([]byte("I got you.\n"))

	s.Stop(fmt.Sprintf("%s %v", r.Method, r.URL))
}

func main() {
	s := ephemeral.New()

	s.HandleFunc("/", done)

	log.Println("Listening to port 8000...")
	msg, err := s.Listen(":8000")
	if err != nil {
		log.Printf("ERROR: %v\n", err)
	}
	log.Printf("Done. Got '%s'\n", msg)
}
