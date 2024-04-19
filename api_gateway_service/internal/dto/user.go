package dto

type User struct {
	UserId   int    `json:"userid,omitempty"`
	Login    string `json:"login,omitempty"`
	Password string `json:"password,omitempty"`
}

type UserId struct {
	UserId int `json:"userid"`
}
