package store

import (
	"demo/internal/conf"
	"demo/internal/store/mysql"
)

var (
	UserRepo UserRepository
)

func Init(conf *conf.Config) error {
	if err := mysql.Init(conf); err != nil {
		return err
	}
	UserRepo = &userRepository{
		mysql.Store.UserDb,
	}
	return nil
}
