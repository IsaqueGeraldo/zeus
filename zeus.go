package zeus

import (
	"github.com/IsaqueGeraldo/agni"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var conn *gorm.DB

func Bootstrap(database string) (*gorm.DB, error) {
	agni.Logger("[zeus]: Initializing database connection")

	db, err := gorm.Open(sqlite.Open(database), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		msg := "[zeus]: Error opening the database connection: " + err.Error()
		agni.Logger(msg, agni.RedText)
		return nil, err
	}

	if err := db.AutoMigrate(&Environment{}); err != nil {
		msg := "[zeus]: Error performing automatic migrations: " + err.Error()
		agni.Logger(msg, agni.RedText)
		return nil, err
	}

	return db, nil
}
