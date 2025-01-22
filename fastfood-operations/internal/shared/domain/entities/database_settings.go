package entities

import "fmt"

type DatabaseSettings struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
	SSL      string
}

func NewDatabaseSettings(host, port, user, password, database, ssl string) *DatabaseSettings {
	return &DatabaseSettings{
		host,
		port,
		user,
		password,
		database,
		ssl,
	}
}

// GetDSN Returns the Data Source Name (DSN)
func (itself *DatabaseSettings) GetDSN() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		itself.Host, itself.User, itself.Password, itself.Database, itself.Port, itself.SSL)
}
