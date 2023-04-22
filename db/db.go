package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"os"
)

var Connection *gorm.DB
var Path string

func init() {
	Path = os.Getenv("DB_PATH")
	if Path == "" {
		Path = "data.db"
	}

	Connection, _ = gorm.Open(sqlite.Open(Path))
}
