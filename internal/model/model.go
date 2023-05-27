package model

type JWT struct {
	FirstName string `jwt:"firstname"`
	LastName  string `jwt:"lastname"`
	Picture   string `jwt:"picture"`
	Name      string `jwt:"name"`
	Locales   string `jwt:"locale"`
	Email     string `jwt:"email"`
}
