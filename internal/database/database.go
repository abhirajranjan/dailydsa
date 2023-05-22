package database

import (
	"fmt"
	"log"

	"github.com/abhirajranjan/dailydsa/internal/model"
	"github.com/abhirajranjan/dailydsa/internal/permissions"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DatabaseBridge struct {
	db *gorm.DB
}

func CreateDatabaseBridge(cfg DatabaseConfig) *DatabaseBridge {
	db, err := gorm.Open(postgres.Open(cfg.Host), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database: ", err)
	}

	// var now time.Time
	// db.Raw("SELECT NOW()").Scan(&now)

	// fmt.Println(now)
	return &DatabaseBridge{db: db}
}

func (*DatabaseBridge) CreateUser(jwt *model.JWT) (int, error) {
	fmt.Println(jwt.FirstName)
	return 0, nil
}

// auth
func (*DatabaseBridge) GetUserRolesBySessionID(sessionID int) (permissions.Permissions, error) {
	return permissions.MultiPermission(permissions.Admin, permissions.User), nil
}

func (*DatabaseBridge) GetProfileBySessionID(sessionID int) (*model.UserProfileModel, error) {
	return new(model.UserProfileModel), nil
}

func (*DatabaseBridge) GetHistoryBySessionID(int) (*model.HistoryResponse, error) {
	return new(model.HistoryResponse), nil
}
