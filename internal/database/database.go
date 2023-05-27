package database

import (
	"fmt"
	"log"
	"strings"

	"github.com/abhirajranjan/dailydsa/internal/model"
	"github.com/abhirajranjan/dailydsa/internal/permissions"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	ErrUserDoesntExist = fmt.Errorf("user doesnt exists")
)

type DatabaseBridge struct {
	db *gorm.DB
}

func CreateDatabaseBridge(cfg DatabaseConfig) *DatabaseBridge {
	db, err := gorm.Open(mysql.Open(cfg.Host), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		TranslateError:                           true,
	})
	if err != nil {
		log.Fatal("failed to connect database: ", err)
	}

	bridge := &DatabaseBridge{
		db: db,
	}
	return bridge
}

func (bridge *DatabaseBridge) CreateUser(jwt *model.JWT) (int, error) {
	user := model.ModelUser{
		FirstName: jwt.FirstName,
		LastName:  jwt.LastName,
		Name:      jwt.Name,
		Email:     jwt.Email,
	}

	if trimmed := strings.Trim(jwt.Locales, " "); trimmed != "" {
		user.Locale = trimmed
	}
	if trimmed := strings.Trim(jwt.Picture, " "); trimmed != "" {
		user.Picture = trimmed
	}

	var userexists model.ModelUser
	err := bridge.db.Transaction(func(tx *gorm.DB) error {
		if m := tx.Model(model.ModelUser{}).Where(user).FirstOrCreate(&userexists); m.Error != nil {
			return m.Error
		}

		if m := tx.Model(model.ModelUserRoles{}).
			Where(model.ModelUserRoles{UserID: userexists.ID, Role: 0}).
			FirstOrCreate(&model.ModelUserRoles{}); m.Error != nil {
			return m.Error
		}
		user.ID = userexists.ID
		return nil
	})
	if err != nil {
		return 0, err
	}
	return user.ID, nil
}

// auth
func (bridge *DatabaseBridge) GetUserRolesBySessionID(sessionID int) (permissions.Permissions, error) {
	roles := []permissions.Permissions{}

	err := bridge.db.Transaction(func(tx *gorm.DB) error {
		session_user := model.ModelUserSession{}
		if m := tx.Model(&model.ModelUserSession{}).
			Select("userid").
			Where("id = ?", sessionID).
			First(&session_user); m.Error != nil {
			return m.Error
		}

		m := tx.Model(&model.ModelUserRoles{}).
			Select("role").
			Where("userid = ?", session_user.UserID)

		if m.Error != nil {
			return m.Error
		}

		row, err := m.Rows()
		if err != nil {
			return err
		}
		var i int
		for row.Next() {
			if err := row.Scan(&i); err != nil {
				return err
			}
			roles = append(roles, permissions.ToPermission(i))
		}

		return nil
	})
	if err != nil {
		return nil, err
	}
	return permissions.MultiPermission(roles...), nil
}

func (bridge *DatabaseBridge) GetProfileBySessionID(sessionID int) (*model.ModelUser, error) {
	user := model.ModelUser{}
	err := bridge.db.Transaction(func(tx *gorm.DB) error {

		session_user := model.ModelUserSession{}
		if m := tx.Model(&model.ModelUserSession{}).
			Select("userid").
			Where("id = ?", sessionID).
			First(&session_user); m.Error != nil {
			return m.Error
		}

		user = model.ModelUser{}

		if m := tx.Model(&model.ModelUser{}).
			Select("id", "firstname", "lastname", "name", "locale", "picture", "email").
			Where("id = ?", session_user.UserID).
			First(&user); m.Error != nil {
			return m.Error
		}
		return nil
	})

	if err != nil {
		return nil, ErrUserDoesntExist
	}

	return &user, nil
}

func (bridge *DatabaseBridge) GetHistoryBySessionID(sessionID int) ([]model.ModelSubmission, error) {
	dest := []model.ModelSubmission{}
	err := bridge.db.Transaction(func(tx *gorm.DB) error {
		session_user := model.ModelUserSession{}
		if m := tx.Model(&model.ModelUserSession{}).
			Select("userid").
			Where("id = ?", sessionID).
			First(&session_user); m.Error != nil {
			return m.Error
		}

		m := tx.Model(&model.ModelSubmission{}).Where("userid = ?", session_user.UserID)
		if m.Error != nil {
			return m.Error
		}

		rows, err := m.Rows()
		if err != nil {
			return err
		}

		for rows.Next() {
			sub := model.ModelSubmission{}
			if err := rows.Scan(&sub.ID, &sub.UserID, &sub.QuestionID, &sub.Submission, &sub.Time); err != nil {
				return err
			}
			dest = append(dest, sub)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return dest, nil
}

func (bridge *DatabaseBridge) CreateSession(userid int) (int, error) {
	sessionID := 0
	err := bridge.db.Transaction(func(tx *gorm.DB) error {
		session := &model.ModelUserSession{UserID: userid}
		if m := tx.Model(&model.ModelUserSession{}).Create(&session); m.Error != nil {
			return m.Error
		}
		sessionID = session.ID
		return nil
	})

	if err != nil {
		return 0, err
	}
	return sessionID, nil
}

func (bridge *DatabaseBridge) UserSubmission(sessionID int, sub *model.ModelSubmission) (int, error) {
	submissionID := 0
	err := bridge.db.Transaction(func(tx *gorm.DB) error {
		session_user := model.ModelUserSession{}
		if m := tx.Model(&model.ModelUserSession{}).Where("id = ?", sessionID).First(&session_user); m.Error != nil {
			return m.Error
		}
		sub.UserID = session_user.UserID
		if m := tx.Model(&model.ModelSubmission{}).Create(sub); m.Error != nil {
			return m.Error
		}
		submissionID = sub.ID
		return nil
	})
	if err != nil {
		return 0, err
	}
	return submissionID, nil
}
