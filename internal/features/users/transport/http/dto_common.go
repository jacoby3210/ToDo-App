package users_transport_http

import (
	"github.com/google/uuid"
	"github.com/jacoby3210/ToDo-App/internal/core/domain"
)

type UserDTOResponse struct {
	ID          uuid.UUID `json:"id"`
	Version     int64     `json:"version"`
	FullName    string    `json:"full_name"`
	PhoneNumber *string   `json:"phone_number,omitempty"`
}

func userDTOFromDomain(user domain.User) UserDTOResponse {
	return UserDTOResponse{
		ID:          user.ID,
		Version:     user.Version,
		FullName:    user.FullName,
		PhoneNumber: user.PhoneNumber,
	}
}

func userDTOFromDomains(users []domain.User) []UserDTOResponse {
	userDTOs := make([]UserDTOResponse, len(users))
	for i, user := range users {
		userDTOs[i] = userDTOFromDomain(user)	
	}
	return userDTOs
}