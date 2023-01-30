package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"

	"github.com/store/config"
	"github.com/store/model"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Database instance
type Dbinstance struct {
	StoreDb    *gorm.DB
	TestdataDb *gorm.DB
}

var DB Dbinstance

// Connect function
func Connect() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", config.Get("DB_HOST"), config.Get("DB_USER"), config.Get("DB_PASSWORD"), config.Get("DB_NAME"), config.GetUint("DB_PORT"))
	storedb, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
		os.Exit(2)
	}
	log.Println("Connected Stores")
	storedb.Logger = logger.Default.LogMode(logger.Info)
	log.Println("running migrations")
	storedb.AutoMigrate(&model.Store{})

	// testdata database
	testdatadb, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
		os.Exit(2)
	}
	log.Println("Connected Testdata")
	testdatadb.Logger = logger.Default.LogMode(logger.Info)
	log.Println("running migrations")
	testdatadb.AutoMigrate(&model.Testdata{})

	// get instance
	DB = Dbinstance{
		StoreDb:    storedb,
		TestdataDb: testdatadb,
	}
}
