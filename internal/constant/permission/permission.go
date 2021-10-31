package permission

var (
	Create = "create"
	Draft  = "draft"
	Update = "update"
	Read   = "read"
)
var (
	CreateUser       = "create user"
	CreateCompany    = "create company"
	UpdateUser       = "update user"
	CreateSystemRole = "create system role"

	DraftCreateCompany = "draft create company"
	DraftCreateUser    = "draft create user"

	DraftUpdateUser = "draft update user"
)

var PermissionObjects = map[string]string{
	CreateUser:         CreateUser,
	UpdateUser:         UpdateUser,
	CreateSystemRole:   CreateSystemRole,
	CreateCompany:      CreateCompany,
	DraftCreateUser:    DraftCreateUser,
	DraftCreateCompany: DraftCreateCompany,
}

var PermissionActions = map[string]string{
	CreateUser:         Create,
	CreateSystemRole:   Draft,
	DraftCreateUser:    Create,
	DraftCreateCompany: Update,
	CreateCompany:      Create,
}

// if a permission has an associated permission for drafting data i.e
// if a resource can be created without being active
var DraftPermissions = map[string]string{
	CreateUser:    DraftCreateUser,
	CreateCompany: DraftCreateCompany,
}
