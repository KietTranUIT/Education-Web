package response

import (
	"user-service/internal/core/entity/Error_Code"
)

type Response struct {
	Data       interface{}          `json:"data"`
	Status     bool                 `json:"status"`
	Error_code Error_Code.ErrorCode `json:"error_code"`
	Error_msg  string               `json:"error_msg"`
}

type SignUpResponse struct {
	DisplayName string
}
