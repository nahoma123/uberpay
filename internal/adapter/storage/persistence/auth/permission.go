package auth

import (
	storage "ride_plus/internal/adapter/storage/persistence"
	utils "ride_plus/internal/constant/model/init"

	"gorm.io/gorm"
)

type permissionPersistence struct {
	conn *gorm.DB
}

func PermissionInit(utils utils.Utils) storage.RolePersistence {
	return &rolePersistence{
		conn: utils.Conn,
	}
}
