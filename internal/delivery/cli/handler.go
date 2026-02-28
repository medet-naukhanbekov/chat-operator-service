package cli

import (
	"bufio"
	"chat_operator_service/internal/domain/user"
	storeRepo "chat_operator_service/internal/repository/store"
	"chat_operator_service/internal/usecase/chat"
	"fmt"
	"os"
	"strings"
)

type Handler struct {
	currentUser user.User
	usecase     *chat.UseCase
}

func NewHandler() *Handler {
	repository := storeRepo.NewMemoryRepository()

	return &Handler{
		currentUser: user.User{
			ID:   "1",
			Role: user.RoleOperation, //user.RoleSenior
		},
		usecase: chat.NewUseCase(repository),
	}
}

func (h *Handler) Run() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Printf(
		"Chat started. User role: %s\n",
		h.currentUser.Role,
	)
	fmt.Println("Type a command (type 'exit' to quit):")

	for {
		fmt.Print("> ")

		if !scanner.Scan() {
			break
		}

		input := strings.TrimSpace(scanner.Text())

		if input == "exit" {
			fmt.Println("Bye 👋")
			return
		}

		h.handleMessage(input)
	}
}

func (h *Handler) handleMessage(message string) {
	response := h.usecase.HandleMessage(h.currentUser, message)
	fmt.Println(response)
}
