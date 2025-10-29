package repositories

import "context"

// IRepository is the base repository interface
type IRepository[T any, ID comparable] interface {
	// Query
	GetByID(ctx context.Context, id ID) (*T, error)
	GetAll(ctx context.Context) ([]T, error)
	GetPaged(ctx context.Context, pageNumber, pageSize int) ([]T, int64, error)
	Find(ctx context.Context, condition interface{}, args ...interface{}) ([]T, error)
	FirstOrDefault(ctx context.Context, condition interface{}, args ...interface{}) (*T, error)
	Count(ctx context.Context) (int64, error)
	
	// Command
	Insert(ctx context.Context, entity *T) error
	InsertMany(ctx context.Context, entities []T) error
	Update(ctx context.Context, entity *T) error
	Delete(ctx context.Context, id ID) error
	SoftDelete(ctx context.Context, id ID, userID int) error
}

// IProjectRepository extends base repository with project-specific methods
type IProjectRepository interface {
	IRepository[interface{}, int]
	GetAllIncludingOcrProjects(ctx context.Context, pageNumber, pageSize int, filters map[string]interface{}) ([]interface{}, int64, error)
	GetByIDIncludingOcrProjects(ctx context.Context, id int) (interface{}, error)
}

// IOcrProjectRepository extends base repository with OCR project-specific methods
type IOcrProjectRepository interface {
	IRepository[interface{}, int]
	GetByProjectID(ctx context.Context, projectID int) ([]interface{}, error)
	GetByType(ctx context.Context, projectType int) ([]interface{}, error)
}

// IUserRepository extends base repository with user-specific methods
type IUserRepository interface {
	IRepository[interface{}, int]
	GetByUsername(ctx context.Context, username string) (interface{}, error)
	GetByEmail(ctx context.Context, email string) (interface{}, error)
	GetWithRoles(ctx context.Context, id int) (interface{}, error)
}

// IRoleRepository extends base repository with role-specific methods
type IRoleRepository interface {
	IRepository[interface{}, int]
	GetByName(ctx context.Context, name string) (interface{}, error)
	GetWithPermissions(ctx context.Context, id int) (interface{}, error)
}

// ITenantRepository extends base repository with tenant-specific methods
type ITenantRepository interface {
	IRepository[interface{}, int]
	GetByTenancyName(ctx context.Context, tenancyName string) (interface{}, error)
}
