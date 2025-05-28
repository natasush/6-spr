package main

import (
	"log"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
	// логгер
	logger := log.New(log.Writer(), "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)

	// сервер
	server := server.NewServer(logger)

	// логируем сообщение о старте сервера
	server.Logger.Println("Starting server on", server.HTTPServer.Addr)

	// запускаем сервер
	if err := server.HTTPServer.ListenAndServe(); err != nil {
		server.Logger.Fatal("ListenAndServe:", err)
	}
}
