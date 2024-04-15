package dto

type User struct {
	UserId   int    `json:"userid"`
	Login    string `json:"login"`
	Password string `json:"password"`
}

type UserId struct {
	UserId int `json:"userid"`
}
