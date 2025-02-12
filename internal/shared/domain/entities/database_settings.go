package entities

import (
	"fmt"
	"net"
)

const (
	MinPoolSize = 3
	MaxPoolSize = 10
)

type DatabaseSettings struct {
	host     string
	port     string
	user     string
	password string
}

func NewDatabaseSettings(host, port, user, password string) *DatabaseSettings {
	return &DatabaseSettings{
		host,
		port,
		user,
		password,
	}
}

func (itself *DatabaseSettings) GetMongoURI() string {
	return fmt.Sprintf(
		"mongodb://%s:%s@%s",
		itself.user,
		itself.password,
		net.JoinHostPort(itself.host, itself.port),
	)
}

func (itself *DatabaseSettings) GetMinPoolSize() *uint64 {
	minPoolSize := uint64(MinPoolSize)
	return &minPoolSize
}

func (itself *DatabaseSettings) GetMaxPoolSize() *uint64 {
	maxPoolSize := uint64(MaxPoolSize)
	return &maxPoolSize
}
