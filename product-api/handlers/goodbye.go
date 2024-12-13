package handlers

import (
	"log"
	"net/http"
)

// Class
type Goodbye struct {
	l *log.Logger
}

// Constructor
func NewGoodbye(l *log.Logger) *Goodbye {
	return &Goodbye{l}
}

// Methods
func (h *Goodbye) ServeHTTP(rw http.ResponseWriter, r *http.Request) {

	rw.Write([]byte("Goodbye..."))
}
