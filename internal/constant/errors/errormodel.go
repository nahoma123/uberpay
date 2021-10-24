package errors

type ErrorModel struct {
	ErrorCode        int    `json:"errorCode"`
	ErrorMessage     string `json:"errorMessage"`
	ErrorDescription string `json:"errorDescription"`
}

func NewErrorResponse(err error) ErrorModel {
	return ErrorModel{
		ErrorMessage:     err.Error(),
		ErrorDescription: Descriptions[err],
		ErrorCode:        ErrCodes[err],
	}
}
func ConvertionError() ErrorModel {
	errValue := ErrorModel{
		ErrorCode:        ErrCodes[ErrorUnableToConvert],
		ErrorDescription: Descriptions[ErrorUnableToConvert],
		ErrorMessage:     ErrorUnableToConvert.Error(),
	}
	return errValue
}
