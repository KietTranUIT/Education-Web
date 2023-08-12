package repository

import (
	"errors"
	"user-service/internal/core/dto"
)

var (
	DuplicateUserAccount   = errors.New("duplicate user account")
	DuplicateStudent       = errors.New("duplicate student")
	InsertUserAccountError = errors.New("insert user account error")
	InsertStudentError     = errors.New("insert student error")
)

type UserRepository interface {
	InsertUserAccount(user dto.UserAccountDTO) error
	InsertStudent(student dto.StudentDTO, username string) error
}
