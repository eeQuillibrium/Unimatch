package models

type User struct {
	Id       int    `db:"id"`
	Login    string `db:"login"`
	PassHash string `db:"passhash"`
}
