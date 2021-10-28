package permission

var (
	CreateUser       = "create user"
	DraftUser        = "draft user"
	CreateSystemRole = "create system role"
)

var PermissionObjects = map[string]string{
	CreateUser:       CreateUser,
	CreateSystemRole: CreateSystemRole,
	DraftUser:        DraftUser,
}

var PermissionAction = map[string]string{
	CreateUser:       "create",
	CreateSystemRole: "create",
	DraftUser:        "draft",
}
