package providerwire

import (
	"log"
	"os"

	"github.com/sonymuhamad/nexttalent-test/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func provideGormSQLDatabase(cfg config.EnvConfig) *gorm.DB {
	logger2 := logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), logger.Config{
		LogLevel: logger.Info,
	})

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  cfg.DatabaseURL,
		PreferSimpleProtocol: true,
	}), &gorm.Config{
		SkipDefaultTransaction: true,
		FullSaveAssociations:   false,
		Logger:                 logger2,
	})

	if err != nil {
		panic(err)
	}

	_, err = db.DB()
	if err != nil {
		panic(err)
	}

	return db
}
