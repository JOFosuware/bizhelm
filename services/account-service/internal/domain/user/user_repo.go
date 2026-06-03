package user

type Repository interface {
	Create(user *User) error
	FetchByToken(tokenHash string) (*User, error)
	FindByEmail(email string) (*User, error)
	AssignRoles(userID string, roleIDs []string) error
	AssignPermissions(userID string, permissionIDs []string) error
	FetchByID(id string) (*User, error)
	FetchRoles(userID string) ([]string, error)
	FetchPermissions(userID string) ([]string, error)
	Update(user *User) error
	UpdatePermissions(userID string, permissionIDs []string) error
	UpdateRoles(userID string, roleIDs []string) error
	Delete(id string) error
	RemoveRoles(userID string, roleIDs []string) error
	RemovePermissions(userID string, permissionIDs []string) error
	CreateRole(roleID string) error
	DeleteRole(roleID string) error
	CreatePermission(permissionID string) error
	DeletePermission(permissionID string) error
	GetRolePermissions(roleID string) ([]string, error)
}