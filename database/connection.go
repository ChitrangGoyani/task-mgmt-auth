package database

import (
	"github.com/ChitrangGoyani/task-mgmt-auth/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := "host=localhost user=postgres password=postgres dbname=task_mgmt_auth port=5432 sslmode=disable"
	connection, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Could not connect to postgres")
	}

	DB = connection
	connection.AutoMigrate(&models.User{})
}
