package main

import (
	"fmt"
	"net/http"

	httpDelivery "chat_operator_service/internal/delivery/http"
	storeRepo "chat_operator_service/internal/repository/store"
	"chat_operator_service/internal/usecase/chat"
)

func main() {
	repository := storeRepo.NewMemoryRepository()
	usecase := chat.NewUseCase(repository)
	handler := httpDelivery.NewHandler(usecase)

	http.HandleFunc("/api/chat/message", handler.HandleMessage)

	fmt.Println("HTTP server started on:8080")
	http.ListenAndServe(":8080", nil)
}
