package commands

import (
	"chat_operator_service/internal/domain/user"
	storeRepo "chat_operator_service/internal/repository/store"
	"fmt"
	"strings"
)

type StoreHoursCommand struct {
	repo storeRepo.Repository
}

func NewStoreHoursCommand(repo storeRepo.Repository) *StoreHoursCommand {
	return &StoreHoursCommand{repo: repo}
}

func (c *StoreHoursCommand) Name() string {
	return "/store_hours"
}

func (c *StoreHoursCommand) Execute(u user.User, args []string) string {
	if len(args) < 1 {
		return "Usage: /store_hours <address>"
	}

	address := strings.Join(args, " ")

	store, err := c.repo.GetByAddress(address)
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
