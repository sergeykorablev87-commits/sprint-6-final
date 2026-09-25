package main

import (
	"log"
	"os"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
	logger := log.New(os.Stdout, "MorseServer: ", log.LstdFlags)

	srv := server.NewServer(logger)

	logger.Printf("Сервер запущен на порту 8080")

	if err := srv.ListenAndServe(); err != nil {
		logger.Fatalf("Ошибка запуска сервера: %v", err)
	}
}
