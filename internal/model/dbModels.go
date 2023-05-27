package model

import (
	"database/sql/driver"

	"gorm.io/datatypes"
)

type enumDifficulty string

const (
	EasyDifficulty   enumDifficulty = "easy"
	MediumDifficulty enumDifficulty = "medium"
	HardDifficulty   enumDifficulty = "hard"
)

func (diff *enumDifficulty) Scan(value interface{}) error {
	*diff = enumDifficulty(value.([]byte))
	return nil
}

func (diff enumDifficulty) Value() (driver.Value, error) {
	return string(diff), nil
}

type ModelUser struct {
	ID        int    `gorm:"column:id;         primaryKey"`
	FirstName string `gorm:"column:firstname;  not null;"`
	LastName  string `gorm:"column:lastname;   not null;"`
	Name      string `gorm:"column:name;       not null;"`
	Picture   string `gorm:"column:picture;    default:null"`
	Locale    string `gorm:"column:locale;     default:null"`
	Email     string `gorm:"column:email;      not null; unique;"`
}

func (ModelUser) TableName() string {
	return "user"
}

type ModelQuestions struct {
	ID         int            `gorm:"column:id;         primaryKey"`
	Link       string         `gorm:"column:link;       not null; default:null"`
	Difficulty enumDifficulty `gorm:"column:difficulty; not null; default:null"`
	Date       datatypes.Date `gorm:"column:date;       default:null"`
}

func (ModelQuestions) TableName() string {
	return "questions"
}

type ModelQuestionStack struct {
	ID         int            `gorm:"column:id;         primaryKey"`
	Priority   int            `gorm:"column:priority;   default:null"`
	Link       string         `gorm:"column:link;       not null; default:null"`
	Difficulty enumDifficulty `gorm:"column:difficulty; not null; default:null"`
}

func (ModelQuestionStack) TableName() string {
	return "question_stack"
}

type ModelSubmission struct {
	ID         int            `gorm:"column:id;         primaryKey"`
	Submission string         `gorm:"column:submission; unique; not null; default:null"`
	Time       datatypes.Time `gorm:"column:time;       default:null"`
	UserID     int            `gorm:"column:userid;     not null; default:null"`
	QuestionID int            `gorm:"column:questionid; not null; default:null"`
}

func (ModelSubmission) TableName() string {
	return "submission"
}

type ModelUserSession struct {
	ID     int `gorm:"column:id;     primaryKey"`
	UserID int `gorm:"column:userid; not null; default:null"`
}

func (ModelUserSession) TableName() string {
	return "user_session"
}

type ModelUserRoles struct {
	UserID int `gorm:"column:userid; not null"`
	Role   int `gorm:"column:role; not null"`
}

func (ModelUserRoles) TableName() string {
	return "user_role"
}
