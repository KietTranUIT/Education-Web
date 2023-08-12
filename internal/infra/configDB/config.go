package configDB

import "fmt"

type ConfigDatabase struct {
	Driver   string
	Username string
	Password string
	Host     string
	Port     uint
	Dbname   string
}

func (conf ConfigDatabase) URL() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", conf.Username, conf.Password, conf.Host, conf.Port, conf.Dbname)
}
