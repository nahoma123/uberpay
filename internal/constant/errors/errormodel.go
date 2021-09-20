package errors

import (
	"strconv"

	"github.com/go-playground/validator/v10"
)

type ErrorModel struct {
	ErrorCode        string `json:"errorCode"`
	ErrorMessage     string `json:"errorMessage"`
	ErrorDescription string `json:"errorDescription"`
}

// type ValErrorModel struct {
// 	ErrorCode string            `json:"errorCode"`
// 	ValError  map[string]string `json:"validationErrors"`
// }

type ValErr map[string]string

func (e ValErr) Error() string {
	return "I am a validation error"
}

func NewErrorResponse(err error) ErrorModel {
	return ErrorModel{
		ErrorMessage:     err.Error(),
		ErrorDescription: Descriptions[err],
		ErrorCode:        strconv.Itoa(ErrCodes[err]),
	}
}

func NewValErrResponse(err validator.ValidationErrorsTranslations) ValErr {
	valErr := ValErr{}
	valErr["errorCode"] = strconv.Itoa(ErrCodes[ErrUnknown])
	for k, v := range err {
		valErr[k] = v
	}
	return valErr
}

func GetStatusCode(err error) int {
	return StatusCodes[err]
}
