package database

import (
	"fmt"

	"github.com/abhirajranjan/dailydsa/internal/model"
	"github.com/abhirajranjan/dailydsa/internal/permissions"
)

var (
	Connection string
)

func init() {
}

func CreateUser(jwt *model.JWT) (int, error) {
	fmt.Println(jwt.FirstName)
	return 0, nil
}

// auth
func GetUserRolesBySessionID(sessionID int) (permissions.Permissions, error) {
	return permissions.MultiPermission(permissions.Admin, permissions.User), nil
}

func GetProfileBySessionID(sessionID int) (*model.UserProfileModel, error) {
	return new(model.UserProfileModel), nil
}

func GetHistoryBySessionID(int) (*model.HistoryResponse, error) {
	return new(model.HistoryResponse), nil
}
