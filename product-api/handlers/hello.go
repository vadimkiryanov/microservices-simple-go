package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type Hello struct {
	l *log.Logger
}

// Конструктор
func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

// Методы
func (h *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	h.l.Println("hello")

	d, err := io.ReadAll(r.Body) // d - данные

	if err != nil {
		http.Error(rw, "It was error :(", http.StatusBadRequest)
		return
	}

	// Отправляем ответ клиенту с полученными данными
	fmt.Fprintf(rw, "Hello: %v\n", string(d))
}
