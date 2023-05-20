package model

import "time"

type JWT struct {
	FirstName string `jwt:"firstname"`
	LastName  string `jwt:"lastname"`
	Picture   string `jwt:"picture"`
	Locales   string `jwt:"locale"`
	Email     string `jwt:"email"`
	Sub       string `jwt:"sub"`
}

type UserProfileModel struct {
	FirstName string
	LastName  string
	Picture   string
	Locales   string
}

// per element format
type QuestionDetails struct {
	Link       string    `json:"link"`
	Difficulty string    `json:"difficulty"`
	Solution   string    `json:"solution"`
	DateTime   time.Time `json:"datetime"`
}

// response format
type HistoryResponse struct {
	Count    int             `json:"count"`
	UserID   string          `json:"userid"`
	Elements QuestionDetails `json:"child"`
}
