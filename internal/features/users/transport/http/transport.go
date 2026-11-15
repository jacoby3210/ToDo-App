package users_transport_http

import (
	"context"
	"net/http"

	"github.com/jacoby3210/ToDo-App/internal/core/domain"
	core_http_server "github.com/jacoby3210/ToDo-App/internal/core/transport/http/server"
)

type UsersHTTPHandler struct {
	usersService UsersService
}

type UsersService interface {
	CreateUser(
		ctx context.Context,
		user domain.User,
	)(domain.User, error)
}

func NewUsersHTTPHandler(
	usersService UsersService,
) *UsersHTTPHandler {
	return &UsersHTTPHandler{
		usersService: usersService,
	}
}

func (h *UsersHTTPHandler) Routes() []core_http_server.Route {
	return []core_http_server.Route{
		{
			Handler: h.CreateUser,
			Method: http.MethodPost,
			Path:"/users",
		},
	}
}

