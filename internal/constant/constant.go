package constant

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"os"
	"strings"
)

const (
	SuperAdmin   = "SUPER-ADMIN"
	CompanyAdmin = "COMPANY-ADMIN"
	CompanyClerk = "COMPANY-CLERK"
)
//ResponseJson creates new json object
func ResponseJson(c *gin.Context, responseData interface{}, statusCode int) {
	c.JSON(statusCode, responseData)
	return
}
//StructValidator validates specific struct
func StructValidator(structName interface{}, validate *validator.Validate, trans ut.Translator) error {
	errV := validate.Struct(structName)
	if errV != nil {
		errs := errV.(validator.ValidationErrors)
		valErr := errs.Translate(trans)
		for key, _ := range valErr {
			value := strings.TrimSpace(valErr[key])
			value += " " + os.Getenv("ErrSecretKey")
			return errors.New(value)
		}
	}
	return nil
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
