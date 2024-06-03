package store

import (
	conf2 "demo/internal/web/conf"
	mysql2 "demo/internal/web/store/mysql"
)

var (
	UserRepo UserRepository
)

func Init(conf *conf2.Config) error {
	if err := mysql2.Init(conf); err != nil {
		return err
	}
	UserRepo = &userRepository{
		mysql2.Store.UserDb,
	}
	return nil
}
