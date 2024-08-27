package data

type UserInfo struct {
	Name string
}

func GetInfoByUID(uid int64) (*UserInfo, error) {
	return &UserInfo{
		Name: "lisi",
	}, nil
}
