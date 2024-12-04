package store

import (
	"demo/pkg/config"
	"demo/pkg/db"
	"fmt"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"sync"
)

type Datastore struct {
	UserDb *gorm.DB
}

func (ds *Datastore) Close() error {
	db, err := ds.UserDb.DB()
	if err != nil {
		return errors.Wrap(err, "get gorm db instance failed")
	}
	return db.Close()
}

var Store *Datastore
var once sync.Once

func Init(config *config.Config) (err error) {
	err = nil
	once.Do(func() {
		var userDb *gorm.DB
		userDb, err = db.CrateDatabase(config.UserDataBase)
		if err != nil {
			err = fmt.Errorf("failed to get mysql store fatory, Datastore: %+v, error: %w", config.UserDataBase.Database, err)
			return
		}
		Store = &Datastore{
			UserDb: userDb,
		}
	})
	return err
}
