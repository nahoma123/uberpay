package constant

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"os"
	"strconv"
	"template/internal/constant/errors"
)

type SuccessData struct {
	Code int
	Data interface{}
}

//ResponseJson returns json data
func ResponseJson(c *gin.Context, responseData interface{}, statusCode int) {
	c.JSON(statusCode, responseData)
}

//StructValidator validates  interfaces
func StructValidator(structName interface{}, validate *validator.Validate) *errors.ErrorModel {
	err := validate.Struct(structName)
	if err != nil {
		return &errors.ErrorModel{
			ErrorCode:        strconv.Itoa(errors.StatusCodes[errors.ErrorUnableToBindJsonToStruct]),
			ErrorDescription: errors.Descriptions[errors.ErrorUnableToBindJsonToStruct],
			ErrorMessage:     errors.ErrorUnableToBindJsonToStruct.Error(),
		}
	}
	return nil
}

//ValidateVariable validates variables
func ValidateVariable(parm interface{}, validate *validator.Validate) error {
	errs := validate.Var(parm, "required")
	if errs != nil {
		return errs
	}
	return nil
}

//DbConnectionString returns  connection string from the env file using os package
func DbConnectionString() (string, error) {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")
	dbname := os.Getenv("DB_NAME")
	addr := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v", host, user, password, dbname, port)
	return addr, nil

}
