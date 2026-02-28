package http

import (
	"chat_operator_service/internal/domain/user"
	"chat_operator_service/internal/usecase/chat"
	"encoding/json"
	"net/http"
)

type Handler struct {
	usecase *chat.UseCase
}

func NewHandler(usecase *chat.UseCase) *Handler {
	return &Handler{
		usecase: usecase,
	}
}

type Request struct {
	UserID  string `json:"user_id"`
	Role    string `json:"role"`
	Message string `json:"message"`
}

type Response struct {
	Answer string `json:"answer"`
}

func (h *Handler) HandleMessage(w http.ResponseWriter, r *http.Request) {
	var req Request

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	u := user.User{
		ID:   req.UserID,
		Role: user.Role(req.Role),
	}

	answer := h.usecase.HandleMessage(u, req.Message)

	resp := Response{
		Answer: answer,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
