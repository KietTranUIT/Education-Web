package config

import "fmt"

type ConfigServer struct {
	Port uint
	Host string
}

func (conf ConfigServer) GetURL() string {
	return fmt.Sprintf("%s:%d", conf.Host, conf.Port)
}
