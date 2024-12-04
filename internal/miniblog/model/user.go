package model

type User struct {
	Id   int64
	Name string
}

type UserInfo struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

func (u *UserInfo) Info() *UserInfo {
	return &UserInfo{Id: u.Id, Name: u.Name}
}
