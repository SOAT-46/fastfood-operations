package configuration

import (
	"github.com/SOAT-46/fastfood-operations/internal/shared/domain/entities"
	logger "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func postgresDialector(dsn string) gorm.Dialector {
	return postgres.Open(dsn)
}

func GormDB(settings *entities.DatabaseSettings) *gorm.DB {
	dialector := postgresDialector(settings.GetDSN())
	db, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		logger.Errorf("Could not open connection: %s", err.Error())
		panic(err)
	}

	logger.Infof("Using %s database", db.Dialector.Name())
	return db
}
