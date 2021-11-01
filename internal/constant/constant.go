package constant

import (
	"fmt"
	"os"
	"ride_plus/internal/constant/errors"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

const (
	SuperAdmin   = "SUPER-ADMIN"
	CompanyAdmin = "COMPANY-ADMIN"
	CompanyClerk = "COMPANY-CLERK"
)

//StructValidator validates specific struct
func StructValidator(structName interface{}, validate *validator.Validate, trans ut.Translator) []string {
	var errorList []string
	errV := validate.Struct(structName)
	if errV != nil {
		errs := errV.(validator.ValidationErrors)
		for _, e := range errs {
			errorList = append(errorList, e.Translate(trans))
		}
		return errorList
	}
	return nil
}

// wrap field validator with error code
func VerifyInput(structName interface{}, validate *validator.Validate, trans ut.Translator) *errors.ErrorModel {
	errs := StructValidator(structName, validate, trans)
	if errs == nil {
		return nil
	}
	return &errors.ErrorModel{
		ErrorCode:        errors.ErrCodes[errors.ErrInvalidField],
		ErrorDescription: errors.Descriptions[errors.ErrInvalidField],
		ErrorMessage:     errors.ErrOneOrMoreFieldsInvalid.Error(),
		ErrorDetail:      errs,
	}
}

//DbConnectionString connction string finder from the .env file
func DbConnectionString() (string, error) {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")
	dbname := os.Getenv("DB_NAME")
	addr := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v", host, user, password, dbname, port)
	return addr, nil
}
