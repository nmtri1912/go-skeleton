package db

import (
	"fmt"
	"log"
	"time"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func NewDB() *gorm.DB {
	log.Println("Connecting to database")
	dns := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		viper.GetString("postgres.host"),
		viper.GetString("postgres.username"),
		viper.GetString("postgres.password"),
		viper.GetString("postgres.database"),
		viper.GetString("postgres.port"),
	)
	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})

	if err != nil {
		log.Fatal("Cannot init postgres connection ", err)
	}

	sqlDb, err := db.DB()
	if err != nil {
		log.Fatal("Can not get database connection")
	}

	// Maximum Idle Connections
	sqlDb.SetMaxIdleConns(12)
	// Maximum Open Connections
	sqlDb.SetMaxOpenConns(24)
	// Idle Connection Timeout
	sqlDb.SetConnMaxIdleTime(600000 * time.Millisecond)
	// Connection Lifetime
	sqlDb.SetConnMaxLifetime(1800000 * time.Millisecond)

	log.Println("Connect to database successfully")
	return db
}
