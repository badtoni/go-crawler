package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// GetDB : Connect to and return the db
func GetDB() *sql.DB {

	databaseHost := os.Getenv("DB_HOST")
	databasePassword := os.Getenv("DB_PASSWORD")
	databasePort := os.Getenv("DB_PORT")
	databaseName := os.Getenv("DB_NAME")
	databaseUser := os.Getenv("DB_USER")

	db := connectToPostgresDB(databaseHost, databasePort, databaseUser, databasePassword, databaseName)

	return db
}

// GetGormDB : Connect to a db and return the gorm
func GetGormDB() *gorm.DB {

	databaseHost := os.Getenv("DB_HOST")
	databasePassword := os.Getenv("DB_PASSWORD")
	databasePort := os.Getenv("DB_PORT")
	databaseName := os.Getenv("DB_NAME")
	databaseUser := os.Getenv("DB_USER")
	timeZone := os.Getenv("TIME_ZONE")

	db := connectToPostgresWithGorm(databaseHost, databasePort, databaseUser, databasePassword, databaseName, timeZone)

	return db
}

// GetSQLiteDB : Connect to a db and return the gorm
func GetSQLiteDB() *gorm.DB {

	db := connectToSQLiteWithGorm()

	return db
}

// InitGormWithDB : Initialize a gorm instance with ab existing databse connection
func InitGormWithDB(sqlDB *sql.DB) *gorm.DB {

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return gormDB
}

// connectToPostgresWithGorm : Connect to a postgres database with gorm, based on the input
func connectToPostgresWithGorm(databaseHost string, databasePort string, databaseUser string, databasePassword string, databaseName string, timeZone string) *gorm.DB {

	//Creating the connection string
	dsn := fmt.Sprintf("user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		databaseHost, databasePort, databaseUser, databasePassword, databaseName, timeZone)

	// Opening a connection to our database
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	// defer db.Close()

	//validate whether or not our connection string was 100% correct.
	// err = db.Ping()
	// if err != nil {
	// 	panic(err)
	// }

	log.Print("Successfully connected to DB")
	return db
}

// connectToSQLiteWithGorm : Function to connect to a sqlite database with gorm, based on the input
func connectToSQLiteWithGorm() *gorm.DB {

	//Opening a connection to our database
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	// defer db.Close()

	//validate whether or not our connection string was 100% correct.
	// err = db.Ping()
	// if err != nil {
	// 	panic(err)
	// }

	log.Print("Successfully connected to DB")
	return db
}

// connectToPostgresDB : Function to connect to a postgres database based on the input
func connectToPostgresDB(databaseHost string, databasePort string, databaseUser string, databasePassword string, databaseName string) *sql.DB {

	//Creating the connection string
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		databaseHost, databasePort, databaseUser, databasePassword, databaseName)

	//Opening a connection to our database
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	//validate whether or not our connection string was 100% correct.
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	log.Print("Successfully connected to DB")
	return db
}
