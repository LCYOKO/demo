package config

import (
	"fmt"
	"time"
)

type Database struct {
	Host        string
	Port        int
	Username    string
	Password    string
	Database    string
	MaxIdeConns int
	MaxConns    int
	MaxIdleTime time.Duration
}

func (d *Database) ToMsqlDNS() string {
	return fmt.Sprintf(`%s:%s@tcp(%s)/%s?charset=utf8&parseTime=%t&loc=%s`,
		d.Username,
		d.Password,
		d.Host,
		d.Database,
		true,
		"Local")
}
