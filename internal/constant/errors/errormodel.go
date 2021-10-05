package errors

import (
	"github.com/go-playground/validator/v10"
	"strconv"
)

type ErrorModel struct {
	ErrorCode        int    `json:"errorCode"`
	ErrorMessage     string `json:"errorMessage"`
	ErrorDescription string `json:"errorDescription"`
}

type ValErr map[string]string

func (e ValErr) Error() string {
	return "I am a validation error"
}

func NewErrorResponse(err error) ErrorModel {
	return ErrorModel{
		ErrorMessage:     err.Error(),
		ErrorDescription: Descriptions[err],
		ErrorCode:        ErrCodes[err],
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
func ConvertionError() ErrorModel {
	errValue := ErrorModel{
		ErrorCode:        ErrCodes[ErrorUnableToConvert],
		ErrorDescription: Descriptions[ErrorUnableToConvert],
		ErrorMessage:     ErrorUnableToConvert.Error(),
	}
	return errValue
}
