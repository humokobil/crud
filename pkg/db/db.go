package db

import (
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	postgresDbCon *gorm.DB
)

func initPostgresDatabase(connString string) (*gorm.DB, error) {
	// Openning file
	return  gorm.Open(postgres.Open(connString), &gorm.Config{Logger: logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
		  SlowThreshold:              time.Second,   // Slow SQL threshold
		  LogLevel:                   logger.Info, // Log level
		  IgnoreRecordNotFoundError: true,           // Ignore ErrRecordNotFound error for logger
		  Colorful:                  false,          // Disable color
		},)})
}

// GetPostgresDb is singleton func for create one global connection
func GetPostgresDb(connString string) (*gorm.DB, error) {
	var err error = nil
	if postgresDbCon == nil {
		postgresDbCon, err = initPostgresDatabase(connString)
	}
	return postgresDbCon, err
}
