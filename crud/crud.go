package crud

import (
	"database/sql"
	"go-crawler/models/dbmodels"
	"log"

	"gorm.io/gorm"
)

// GetRssChannels : Function to get the rss channels from the db
func GetRssChannels(db *sql.DB) *sql.Rows {
	rssChannels, err := db.Query("select id, url from rss_channels")
	if err != nil {
		log.Print("Panic - db Query ")
		panic(err)
	}
	return rssChannels
}

func CreateUser(db *gorm.DB) {
	//You can insert multiple records too
	var users []dbmodels.User = []dbmodels.User{
		dbmodels.User{Name: "Ricky", Address: "Sydney"},
		dbmodels.User{Name: "Adam", Address: "Brisbane"},
		dbmodels.User{Name: "Justin", Address: "California"},
	}

	for _, user := range users {
		db.Create(&user)
	}
}

func GetUser(db *gorm.DB) {

	// Get first record, order by primary key
	db.First(&User)
	// Get last record, order by primary key
	db.Last(&dbmodels.User)
	// Get all records
	db.Find(&dbmodels.User)
	// Get record with primary key (only works for integer primary key)
	db.First(&dbmodels.User, 10)
}
