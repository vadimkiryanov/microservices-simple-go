package main

import (
	"context"
	"log"
	"main-mode/handlers"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)

	ph := handlers.NewProducts(l)

	// Мультиплексер
	sm := http.NewServeMux()
	sm.Handle("/", ph)

	// Конфик сервера
	s := http.Server{
		Addr:         ":9090",
		Handler:      sm,
		IdleTimeout:  time.Second * 120,
		ReadTimeout:  time.Second * 1,
		WriteTimeout: time.Second * 1,
	}

	go func() {
		// Запускаем сервер на порту 8080
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal, 1) // Создание каналана
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	l.Println("Получено завершение, корректное завершение работы", sig)

	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc)
}
