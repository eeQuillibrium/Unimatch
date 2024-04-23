package dto


type Profile struct {
	UserId int    `json:"userid"`
	Name   string `json:"name"`
	Age    int    `json:"age"`
	About  string `json:"about"`
}
