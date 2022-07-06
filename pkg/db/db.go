package db

import (
	"fmt"
	"github.com/StepanShevelev/l0/pkg/api"
	cfg "github.com/StepanShevelev/l0/pkg/config"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DbInstance struct {
	Db *gorm.DB
}

var Database DbInstance

func ConnectToDb(config *cfg.Config) {

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s timezone=Europe/Moscow",
		config.DB.Host,
		config.DB.Port,
		config.DB.Username,
		config.DB.Password,
		config.DB.Name,
		config.DB.SSLMode,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logrus.Fatal("Failed to connect to the database! \n", err)
	}

	db.AutoMigrate(&Order{})
	db.AutoMigrate(&Delivery{})
	db.AutoMigrate(&Payment{})
	db.AutoMigrate(&Items{})

	Database = DbInstance{
		Db: db,
	}
}

func GetOrderByUid() error {

	var order *Order
	result := Database.Db.Preload("Delivery").Preload("Items").Preload("Payment").Find(&order)
	if result.Error != nil {
		logrus.Info("Can`t find order", result.Error)
		return result.Error
	}

	api.Caching.SetCache(order.OrderUID, order)

	return nil

}
