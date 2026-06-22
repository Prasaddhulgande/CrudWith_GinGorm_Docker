package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)
   var DB gorm.DB
func connectDB() {
	dsn:= "host=localhost port=5432 user=postgres password=postgres sslmode=disable dbname=gorm"
db, err:= gorm.Open(postgres.OPen(dsn), &gorm.Config{})

if err!=nil {
	panic(err)
}
db.AutoMigrate(&models.User{})
DB=db

}