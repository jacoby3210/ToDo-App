package users_postgres_repository

import (
	"github.com/google/uuid"
	"github.com/jacoby3210/ToDo-App/internal/core/domain"
)

type UserModel struct {
	ID          uuid.UUID
	Version     int64
	FullName    string
	PhoneNumber *string
}

func userDomainsFromModels(userModels []UserModel) []domain.User {
	userDomains := make([]domain.User, len(userModels))
	for i, userModel := range userModels {
		userDomains[i] = domain.NewUser(
			userModel.ID,
			userModel.Version,
			userModel.FullName,
			userModel.PhoneNumber,
		)
	}
	return userDomains
}