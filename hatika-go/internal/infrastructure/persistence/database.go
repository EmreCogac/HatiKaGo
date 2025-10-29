package persistence

import (
	"fmt"
	"log"
	"time"

	"hatika-go/internal/domain/entities"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// DatabaseConfig holds database configuration
type DatabaseConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
	SSLMode  string
}

// NewDatabase creates a new database connection
func NewDatabase(config *DatabaseConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		config.Host,
		config.Port,
		config.User,
		config.Password,
		config.DBName,
		config.SSLMode,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		NowFunc: func() time.Time {
			return time.Now().UTC()
		},
	})

	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get database instance: %w", err)
	}

	// Connection pool settings
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	log.Println("Database connection established successfully")

	return db, nil
}

// database migrations
func AutoMigrate(db *gorm.DB) error {
	log.Println("Running database migrations...")

	err := db.AutoMigrate(
		&entities.User{},
		&entities.Role{},
		&entities.Permission{},
		&entities.Tenant{},
		&entities.Project{},
		&entities.OcrProject{},
	)

	if err != nil {
		return fmt.Errorf("failed to run migrations: %w", err)
	}

	log.Println("Database migrations completed successfully")
	return nil
}

func SeedData(db *gorm.DB) error {
	log.Println("Seeding initial data...")

	permissions := []entities.Permission{
		{Name: entities.PagesUsers, DisplayName: "Users", Description: "Access to users page"},
		{Name: entities.PagesRoles, DisplayName: "Roles", Description: "Access to roles page"},
		{Name: entities.PagesProjects, DisplayName: "Projects", Description: "Access to projects page"},
		{Name: entities.PagesOcrProjects, DisplayName: "OCR Projects", Description: "Access to OCR projects page"},
		{Name: entities.PagesTenants, DisplayName: "Tenants", Description: "Access to tenants page"},
		{Name: entities.UsersCreate, DisplayName: "Create User", Description: "Can create users"},
		{Name: entities.UsersEdit, DisplayName: "Edit User", Description: "Can edit users"},
		{Name: entities.UsersDelete, DisplayName: "Delete User", Description: "Can delete users"},
		{Name: entities.RolesCreate, DisplayName: "Create Role", Description: "Can create roles"},
		{Name: entities.RolesEdit, DisplayName: "Edit Role", Description: "Can edit roles"},
		{Name: entities.RolesDelete, DisplayName: "Delete Role", Description: "Can delete roles"},
		{Name: entities.ProjectsCreate, DisplayName: "Create Project", Description: "Can create projects"},
		{Name: entities.ProjectsEdit, DisplayName: "Edit Project", Description: "Can edit projects"},
		{Name: entities.ProjectsDelete, DisplayName: "Delete Project", Description: "Can delete projects"},
	}

	for _, perm := range permissions {
		var existingPerm entities.Permission
		result := db.Where("name = ?", perm.Name).First(&existingPerm)
		if result.Error == gorm.ErrRecordNotFound {
			if err := db.Create(&perm).Error; err != nil {
				log.Printf("Warning: Failed to create permission %s: %v", perm.Name, err)
			}
		}
	}

	adminRole := entities.Role{
		Name:        entities.AdminRoleName,
		DisplayName: "Administrator",
		Description: "System administrator with full access",
		IsStatic:    true,
		IsDefault:   false,
	}

	var existingAdminRole entities.Role
	result := db.Where("name = ?", adminRole.Name).First(&existingAdminRole)
	if result.Error == gorm.ErrRecordNotFound {
		if err := db.Create(&adminRole).Error; err != nil {
			return fmt.Errorf("failed to create admin role: %w", err)
		}

		var allPermissions []entities.Permission
		db.Find(&allPermissions)
		if err := db.Model(&adminRole).Association("Permissions").Append(allPermissions); err != nil {
			log.Printf("Warning: Failed to assign permissions to admin role: %v", err)
		}
	}

	userRole := entities.Role{
		Name:        entities.UserRoleName,
		DisplayName: "User",
		Description: "Standard user with limited access",
		IsStatic:    true,
		IsDefault:   true,
	}

	var existingUserRole entities.Role
	result = db.Where("name = ?", userRole.Name).First(&existingUserRole)
	if result.Error == gorm.ErrRecordNotFound {
		if err := db.Create(&userRole).Error; err != nil {
			return fmt.Errorf("failed to create user role: %w", err)
		}
	}

	log.Println("Initial data seeded successfully")
	return nil
}
