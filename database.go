package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func SetupDatabase() (*gorm.DB, error) {
	// Connect to the database using the SQLite driver
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Migrate the database schema
	err = db.AutoMigrate(&Todo{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
