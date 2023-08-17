package service

import (
	"user-service/internal/common"
	"user-service/internal/core/dto"
	"user-service/internal/core/entity/Error_Code"
	"user-service/internal/core/model/request"
	"user-service/internal/core/model/response"
	"user-service/internal/core/port/repository"
)

type UserService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return UserService{
		userRepo: userRepo,
	}
}

func (u UserService) SignUp(req request.SignUpRequest) *response.Response {
	if req.Username == "" {
		return u.CreateFailResponse(Error_Code.InvalidUsernameMsg, Error_Code.InvalidRequest)
	}

	if req.Password == "" {
		return u.CreateFailResponse(Error_Code.InvalidPasswordMsg, Error_Code.InternalError)
	}

	// Insert the user account into database
	user_account := dto.UserAccountDTO{
		Username:    req.Username,
		Password:    req.Password,
		CreatedDate: common.GetTimeNow(),
		UpdatedDate: common.GetTimeNow(),
		Type:        1,
	}
	err_insertUser := u.userRepo.InsertUserAccount(user_account)

	if err_insertUser != nil {
		if err_insertUser == repository.DuplicateUserAccount {
			return u.CreateFailResponse(Error_Code.DuplicateUserErrMsg, Error_Code.DuplicateUser)
		}
		return u.CreateFailResponse(Error_Code.InternalErrMsg, Error_Code.InternalError)
	}

	// Insert the infomation student account into database
	student := dto.StudentDTO{
		Mahv:         common.GetID(),
		Surname:      req.Surname,
		Name:         req.Name,
		Gender:       req.Gender,
		DayOfBirth:   req.DateOfBirth,
		PlaceOfBirth: req.PlaceOfBirth,
	}
	err_insertStudent := u.userRepo.InsertStudent(student, req.Username)
	if err_insertStudent != nil {
		if err_insertStudent == repository.DuplicateStudent {
			return u.CreateFailResponse(Error_Code.DuplicateStudentErrMsg, Error_Code.DuplicateUser)
		}
		return u.CreateFailResponse(Error_Code.InternalErrMsg, Error_Code.InternalError)
	}

	rep := response.SignUpResponse{
		DisplayName: user_account.Username,
	}
	return u.CreateSuccessResponse(rep)
}

// Create a Fail Response
func (u UserService) CreateFailResponse(message string, err Error_Code.ErrorCode) *response.Response {
	return &response.Response{
		Status:     false,
		Error_code: err,
		Error_msg:  message,
	}
}

// Create a Success Response
func (u UserService) CreateSuccessResponse(rep response.SignUpResponse) *response.Response {
	return &response.Response{
		Data:       rep.DisplayName,
		Status:     true,
		Error_code: Error_Code.Success,
		Error_msg:  Error_Code.SuccessErrMsg,
	}
}
