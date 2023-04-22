package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gen"
	"gorm.io/gorm"

	"os"
	"whisprite/model"
)

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "../query",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
	})

	db_path := os.Getenv("DB_PATH")
	if db_path == "" {
		db_path = "data.db"
	}

	gormdb, _ := gorm.Open(sqlite.Open(db_path))
	g.UseDB(gormdb) // reuse your gorm db

	// Generate basic type-safe DAO API for struct `model.User` following conventions
	g.ApplyBasic(model.Alias{}, model.Quote{}, model.Counter{}, model.CounterContribution{}, model.User{})

	// Generate the code
	g.Execute()
}
