package commands

import "chat_operator_service/internal/domain/user"

type HelpCommand struct{}

func NewHelpCommand() *HelpCommand {
	return &HelpCommand{}
}

func (c *HelpCommand) Name() string {
	return "/help"
}

func (c *HelpCommand) Execute(u user.User, args []string) string {
	return `
	Available commands:
	/help - show available commands 
	/store_hours <address> - show store working hours
	/add_store <address> <open> <close> - add new store
	`
}
