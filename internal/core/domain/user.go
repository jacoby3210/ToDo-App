package domain

import (
	"fmt"
	"regexp"

	"github.com/google/uuid"

	core_errors "github.com/jacoby3210/ToDo-App/internal/core/errors"
)

type User struct {
	ID          uuid.UUID
	Version     int64
	FullName    string
	PhoneNumber *string
}

// GET /users/{id}

func NewUser(
	id uuid.UUID,
	version int64,
	fullName string,
	phoneNumber *string,
) User {
	return User{
		ID:          id,
		Version:     version,
		FullName:    fullName,
		PhoneNumber: phoneNumber,
	}
}

func NewUserUninitialized(
	fullName string,
	phoneNumber *string,
) User {
	return NewUser(
		UninitializedID,
		UninitializedVersion,
		fullName,
		phoneNumber,
	)
}

func (u *User) Validate() error {
	fullNameLength := len([]rune(u.FullName))
	if fullNameLength < 3 || fullNameLength > 100 {
		return fmt.Errorf(
			"invalid `FullName` len: %d: %w", 
			fullNameLength, 
			core_errors.ErrInvalidArgument,
		)
	}
	if u.PhoneNumber != nil {
		phoneNumberLen := len([]rune(*u.PhoneNumber))
		if phoneNumberLen < 10 || phoneNumberLen > 15 {
			return fmt.Errorf(
				"invalid `PhoneNumber` len: %d: %w", 
				phoneNumberLen, 
				core_errors.ErrInvalidArgument,
			)
		}

		re := regexp.MustCompile(`^\+?\d+$`)
		if !re.MatchString(*u.PhoneNumber) {
			return fmt.Errorf(
				"invalid `PhoneNumber` format: %s: %w", 
				*u.PhoneNumber, 
				core_errors.ErrInvalidArgument,
			)
		}
	}
	
	return nil

}