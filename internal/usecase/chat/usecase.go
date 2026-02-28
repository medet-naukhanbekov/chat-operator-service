package chat

import (
	"chat_operator_service/internal/domain/store"
	"chat_operator_service/internal/domain/user"
	storeRepo "chat_operator_service/internal/repository/store"
	"fmt"
	"strings"
)

type UseCase struct {
	storeRepo storeRepo.Repository
}

func NewUseCase(storeRepo storeRepo.Repository) *UseCase {
	return &UseCase{
		storeRepo: storeRepo,
	}
}

func (uc *UseCase) HandleMessage(u user.User, message string) string {
	if !strings.HasPrefix(message, "/") {
		return "Unknown command. Use /help"
	}

	parts := strings.Split(message, " ")
	command := parts[0]

	switch command {
	case "/help":
		return uc.help()

	case "/store_hours":
		return uc.storeHours(parts)
	case "/add_store":
		return uc.addStore(u, parts)
	default:
		return "Command not found. Use /help"
	}
}

func (uc *UseCase) help() string {
	return `
	Available commands:
	/help - show available commands
	/store_hours <address> - show store working hours
	`
}

func (uc *UseCase) storeHours(parts []string) string {
	if len(parts) < 2 {
		return "Usage: /store_hours <address>"
	}
	address := strings.Join(parts[1:], " ")

	store, err := uc.storeRepo.GetByAddress(address)
	if err != nil {
		return "Store not found"
	}

	return fmt.Sprintf(
		"Store at %s works from %s to %s",
		store.Address,
		store.OpenAt,
		store.CloseAt,
	)
}

func (uc *UseCase) addStore(u user.User, parts []string) string {
	if u.Role != user.RoleSenior {
		return "Access denied: only senior operators can add stores"
	}

	if len(parts) < 4 {
		return "Usage: /add_store <address> <open> <close>"
	}

	open := parts[len(parts)-2]
	close := parts[len(parts)-1]
	address := strings.Join(parts[1:len(parts)-2], " ")

	err := uc.storeRepo.Add(store.Store{
		Address: address,
		OpenAt:  open,
		CloseAt: close,
	})

	if err != nil {
		return "Failed to add store"
	}

	return "Store added successfully"
}
