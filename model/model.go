package model

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

type Goly struct {
	ID       uint64 `json:"id" gorm:"primary"`
	Redirect string `json:"redirect"`
	Goly     string `json:"goly" gorm:"unique;not null"`
	Clicked  uint64 `json:"clicked"`
	Random   bool   `json:"random"`
}

func Setup() {
	dburl := os.Getenv("DB_URL")
	var err error

	db, err = gorm.Open(postgres.Open(dburl), &gorm.Config{})
	if err != nil {
		log.Println("Error connecting to the Database")
		panic(err)
	}

	err = db.AutoMigrate(&Goly{})
	if err != nil {
		log.Println(err)
	}
	log.Println("Successfully connected to the Database")
}
