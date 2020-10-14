package Controller

import (
	"io"
	"log"
	"net/http"
)

type Server struct {
	port string
}

func (s *Server) Init() {
	log.Println("Initializing Server")
	s.port = ":8080"
}

func helloHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello, world!\n")
}

func (s *Server) Start() {
	http.HandleFunc("/hello", helloHandler)
	log.Println("Listing for requests at http://localhost:8080/hello")
	log.Fatal(http.ListenAndServe(s.port, nil))
}
