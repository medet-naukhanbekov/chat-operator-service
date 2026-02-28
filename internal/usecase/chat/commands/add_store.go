package commands

import (
	"chat_operator_service/internal/domain/store"
	"chat_operator_service/internal/domain/user"
	storeRepo "chat_operator_service/internal/repository/store"
	"strings"
)

type AddStoreCommand struct {
	repo storeRepo.Repository
}

func NewAddStoreCommand(repo storeRepo.Repository) *AddStoreCommand {
	return &AddStoreCommand{repo: repo}
}

func (c *AddStoreCommand) Name() string {
	return "/add_store"
}

func (c *AddStoreCommand) Execute(u user.User, args []string) string {
	if u.Role != user.RoleSenior {
		return "Access denied: only senior operation can add stores"
	}

	if len(args) < 3 {
		return "Usage: /add_store <address> <open> <close>"
	}

	open := args[len(args)-2]
	close := args[len(args)-1]
	address := strings.Join(args[:len(args)-2], " ")

	err := c.repo.Add(store.Store{
		Address: address,
		OpenAt:  open,
		CloseAt: close,
	})

	if err != nil {
		return "Failed to add store"
	}

	return "Store added successfully"
}
