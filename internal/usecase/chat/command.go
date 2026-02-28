package chat

import "chat_operator_service/internal/domain/user"

type Command interface {
	Name() string
	Execute(u user.User, args []string) string
}
