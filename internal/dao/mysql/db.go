package mysql

import (
	"database/sql"
	"demo/internal/conf"
	"fmt"
	"github.com/pkg/errors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)


type Datastore struct {
	userDb *gorm.DB
}

func (ds *Datastore) Close() error {
	db, err := ds.userDb.DB()
	if err != nil {
		return errors.Wrap(err, "get gorm db instance failed")
	}
	return db.Close()
}

func NewMySQLFactory(config *conf.Config) (store *Datastore, err error) {
	var userDb *gorm.DB
	userDb, err = crateDatabase(config.UserDataBase)
	if err != nil {
		return nil, fmt.Errorf("failed to get mysql store fatory, Datastore: %+v, error: %w", config.UserDataBase.Database, err)
	}
	return &Datastore{userDb}, nil
}

func crateDatabase(conf conf.Database) (dbIns *gorm.DB, err error) {
	dbIns, err = gorm.Open(mysql.New(mysql.Config{
		// DSN data source name
		DSN: conf.ToMsqlDNS(),
		// string 类型字段的默认长度
		DefaultStringSize: 256,
		// 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DisableDatetimePrecision: true,
		// 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameIndex: true,
		// 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		DontSupportRenameColumn: true,
		// 根据当前 MySQL 版本自动配置
		SkipInitializeWithVersion: false,
	}), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	var db *sql.DB
	db, err = dbIns.DB()
	if err != nil {
		return nil, err
	}
	db.SetMaxIdleConns(conf.MaxIdeConns)
	db.SetMaxOpenConns(conf.MaxConns)
	db.SetConnMaxIdleTime(conf.MaxIdleTime)
	return dbIns, nil
}
