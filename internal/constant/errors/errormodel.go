package errors

type ErrorModel struct {
	ErrorCode        string `json:"errorCode"`
	ErrorMessage     string `json:"errorMessage"`
	ErrorDescription string `json:"errorDescription"`
}
