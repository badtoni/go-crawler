package main

import (
	"context"
	"go-crawler/crud"
	"go-crawler/database"
	"go-crawler/routes"
	"go-crawler/utils"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/lib/pq"

	"gopkg.in/natefinch/lumberjack.v2"
)

func main() {

	// Get enviroment variables
	utils.GetEnvVars()

	// TODO: save logs also into a log file
	// LOG_FILE_LOCATION="/logs/rss-crawler.log"

	logFileLocation := os.Getenv("LOG_FILE_LOCATION")
	serverPort := os.Getenv("SERVER_PORT")

	// Configure Logging
	if logFileLocation != "" {
		log.SetOutput(&lumberjack.Logger{
			Filename:   logFileLocation,
			MaxSize:    500, // megabytes
			MaxBackups: 3,
			MaxAge:     28,   //days
			Compress:   true, // disabled by default
		})
	}

	// Create Server and Route Handlers
	srv := routes.CreateRoutes(serverPort)

	db := database.GetSQLiteDB()

	// db, err := gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/ormdemo?charset=utf8&parseTime=True")
	// if err != nil {
	// 	log.Panic(err)
	// }

	database.MigrateUser(db)
	crud.CreateUser(db)

	crud.GetUser(db)

	// Crawl the rss channels for ever
	go infinityCrawler()

	// Graceful Shutdown
	waitForShutdown(srv)
}

// TODO: understand that method
func waitForShutdown(srv *http.Server) {
	interruptChan := make(chan os.Signal, 1)
	signal.Notify(interruptChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// Block until we receive our signal.
	<-interruptChan

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	srv.Shutdown(ctx)

	log.Println("Shutting down")
	os.Exit(0)
}

// infinityCrawler : Crwal the rss channels for ever
func infinityCrawler() {
	// Crawl-Loop
	for {
		// crawlers.CrawlRSSToDB()
		log.Println("Loop")
		time.Sleep(10 * time.Second)
	}
}
