package chat

import (
	"chat_operator_service/internal/domain/user"
	storeRepo "chat_operator_service/internal/repository/store"
	"chat_operator_service/internal/usecase/chat/commands"
	"strings"
)

type UseCase struct {
	storeRepo storeRepo.Repository
	commands  map[string]Command
}

func NewUseCase(storeRepo storeRepo.Repository) *UseCase {
	uc := &UseCase{
		storeRepo: storeRepo,
		commands:  make(map[string]Command),
	}

	//регистрируем команды
	uc.RegisterCommand(commands.NewHelpCommand())
	uc.RegisterCommand(commands.NewStoreHoursCommand(storeRepo))
	uc.RegisterCommand(commands.NewAddStoreCommand(storeRepo))

	return uc
}

func (uc *UseCase) RegisterCommand(cmd Command) {
	uc.commands[cmd.Name()] = cmd
}

func (uc *UseCase) HandleMessage(u user.User, message string) string {
	message = strings.TrimSpace(message)

	if !strings.HasPrefix(message, "/") {
		return "Unknown command. Use /help"
	}

	parts := strings.Fields(message)
	commandName := parts[0]
	args := parts[1:]

	cmd, ok := uc.commands[commandName]
	if !ok {
		return "Command not found. Use /help"
	}

	return cmd.Execute(u, args)
}
