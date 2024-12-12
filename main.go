package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	// Устанавливаем обработчик для корневого пути "/"
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		// Выводим в консоль сообщение "hello"
		fmt.Println("hello")

		// Читаем тело запроса
		d, err := io.ReadAll(r.Body) // d - данные

		// Если произошла ошибка при чтении тела запроса
		if err != nil {
			// Возвращаем ошибку "Bad Request" клиенту
			http.Error(rw, "It was error :(", http.StatusBadRequest)
			return
		}

		// Отправляем ответ клиенту с полученными данными
		fmt.Fprintf(rw, "Hello: %v\n", string(d))
	})

	// Устанавливаем обработчик для пути "/goodbye"
	http.HandleFunc("/goodbye", func(http.ResponseWriter, *http.Request) {
		// Выводим в консоль сообщение "Goodbye"
		fmt.Println("Goodbye")
	})

	// Запускаем сервер на порту 8080
	http.ListenAndServe(":8080", nil)
}
