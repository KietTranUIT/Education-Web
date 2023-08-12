package repository

import (
	"fmt"
	"user-service/internal/core/dto"
	"user-service/internal/core/port/repository"

	_ "github.com/go-sql-driver/mysql"
)

type UserRepository struct {
	db repository.Database
}

const (
	duplicateEntryMsg = "Duplicate entry"
	numRowInserted    = 1
)

var (
	InsertUserAccountStatement = "INSERT INTO USER_ACCOUNT VALUES('%s','%s','%s','%s','%d')"
	InsertStudentStatement     = "INSERT INTO STUDENT VALUES ('%s','%s','%s','%s','%s','%s', '%s')"
)

func NewUserRepository(db repository.Database) UserRepository {
	return UserRepository{
		db: db,
	}
}

func (u UserRepository) InsertUserAccount(user dto.UserAccountDTO) error {
	result, err := u.db.GetDB().Exec(fmt.Sprintf(InsertUserAccountStatement, user.Username, user.Password, user.CreatedDate, user.UpdatedDate, user.Type))

	if err != nil {
		fmt.Println(err.Error())

		if err.Error() == duplicateEntryMsg {
			return repository.DuplicateUserAccount
		}
	}
	rowAffected, err := result.RowsAffected()

	if err != nil {
		return err
	}

	if rowAffected != numRowInserted {
		return err
	}
	return nil
}

func (u UserRepository) InsertStudent(student dto.StudentDTO, username string) error {
	result, err := u.db.GetDB().Exec(fmt.Sprintf(InsertStudentStatement,
		student.Mahv,
		student.Surname,
		student.Name,
		student.Gender,
		student.DayOfBirth,
		student.PlaceOfBirth,
		username))

	if err != nil {
		fmt.Println(err.Error())
		if err.Error() == duplicateEntryMsg {
			return repository.DuplicateStudent
		}
	}

	rowAffected, err := result.RowsAffected()

	if err != nil {
		return err
	}

	if rowAffected != numRowInserted {
		return repository.InsertStudentError
	}
	return nil
}
