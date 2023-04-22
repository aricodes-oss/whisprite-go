package main

import (
	"gorm.io/gen"

	"whisprite/db"
	"whisprite/model"
)

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "../query",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
	})

	g.UseDB(db.Connection) // reuse your gorm db

	// Generate basic type-safe DAO API for struct `model.User` following conventions
	g.ApplyBasic(model.CommandAlias{}, model.Quote{}, model.Counter{}, model.CounterContribution{}, model.User{})
	g.GenerateAllTable()

	// Generate the code
	g.Execute()
}
