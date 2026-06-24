package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Регистрируем обработчик для корневого URL ("/")
	// Мы используем анонимную функцию (closure) прямо здесь.
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// fmt.Fprintf очень удобен, он записывает форматированную строку
		// прямо в http.ResponseWriter, который отправляет ее клиенту.
		_, _ = fmt.Fprintf(w, "Hello, World!")
	})

	// Определяем порт
	port := ":8080"

	// Сообщение в консоль, что сервер стартовал
	log.Printf("Запускаю веб-сервер на http://localhost%s\n", port)

	// Запускаем сервер на указанном порту.
	// Используем nil для http.Handler, что означает использование
	// DefaultServeMux (куда мы и зарегистрировали наш обработчик).
	// http.ListenAndServe блокирует выполнение, пока сервер работает.
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal("Сервер не смог запуститься: ", err)
	}
}
