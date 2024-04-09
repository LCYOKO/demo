package mysql

import (
	"fmt"
	"github.com/pkg/errors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"sync"
)

type Datastore struct {
	db *gorm.DB
}

func (ds *Datastore) Close() error {
	db, err := ds.db.DB()
	if err != nil {
		return errors.Wrap(err, "get gorm db instance failed")
	}
	return db.Close()
}

var (
	store *Datastore
	once  sync.Once
)

// GetMySQLFactoryOr create mysql factory with the given config.
func GetMySQLFactoryOr() (store *Datastore, err error) {
	var dbIns *gorm.DB
	once.Do(func() {
		options := &gorm.Config{
			//Host:                  opts.Host,
			//Username:              opts.Username,
			//Password:              opts.Password,
			//Database:              opts.Database,
			//MaxIdleConnections:    opts.MaxIdleConnections,
			//MaxOpenConnections:    opts.MaxOpenConnections,
			//MaxConnectionLifeTime: opts.MaxConnectionLifeTime,
			//LogLevel:              opts.LogLevel,
			//Logger:                logger.New(opts.LogLevel),
		}
		dbIns, err = gorm.Open(mysql.Open(""), options)
		store = &Datastore{dbIns}
	})

	if store == nil || err != nil {
		return nil, fmt.Errorf("failed to get mysql store fatory, Datastore: %+v, error: %w", store, err)
	}

	return store, nil
}
