package init

import (
	storage "ride_plus/internal/adapter/storage/persistence"
	"time"

	"github.com/casbin/casbin/v2"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"

	"gorm.io/gorm"
)

type Utils struct {
	Conn             *gorm.DB
	EmailPersistence storage.EmailPersistence
	GoValidator      *validator.Validate
	Translator       ut.Translator
	Timeout          time.Duration
	Enforcer         *casbin.Enforcer
}
