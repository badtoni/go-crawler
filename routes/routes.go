package routes

import (
	"encoding/json"
	"fmt"
	"go-crawler/crawlers"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// Example route handler
func handler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	name := query.Get("name")
	if name == "" {
		name = "Guest"
	}
	log.Printf("Received request for %s\n", name)
	w.Write([]byte(fmt.Sprintf("Hello, %s\n", name)))
}

// crawlAndSave : Crawl rss feeds and save them in db
func crawlAndSave(w http.ResponseWriter, r *http.Request) {

	// TODO: maybe leave this code and check for URL parameters
	// vars := mux.Vars(r)
	// id, _ := strconv.Atoi(vars["id"])
	query := r.URL.Query()
	name := query.Get("name")
	if name == "" {
		name = "Guest"
	}
	log.Printf("Received a Crawl-and-save-to-the-DB request \n")

	// TODO make a security check
	// crawlers.CrawlRSSAndSaveToDB()
	// TODO: make a response which informs if the crawl was succesfull or not
}

// crawlAndReturn : Crawl rss feeds and send them back
func crawlAndReturn(w http.ResponseWriter, r *http.Request) {

	// TODO: maybe leave this code and check for URL parameters
	// vars := mux.Vars(r)
	// id, _ := strconv.Atoi(vars["id"])
	query := r.URL.Query()
	name := query.Get("name")
	if name == "" {
		name = "Guest"
	}
	log.Printf("Received a Crawl-and-send-back request \n")

	rssfeeds := crawlers.CrawlRSSAndReturn()

	// TODO make a security check
	// // TODO: make a response which informs if the crawl was succesfull or not, amd returns the rss feeds
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(rssfeeds)
}

// CreateRoutes : Create the routes for the calls of this API
func CreateRoutes(serverPort string) *http.Server {
	// Create Server and Route Handlers
	router := mux.NewRouter()

	router.HandleFunc("/", handler)
	router.HandleFunc("/crawlSave/", crawlAndSave).Methods("GET")
	router.HandleFunc("/crawlReturn/", crawlAndReturn).Methods("GET")

	srv := &http.Server{
		Handler:      router,
		Addr:         ":" + serverPort,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	// Start Server
	go func() {
		log.Println("Starting Crawler API Server")
		if err := srv.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	return srv
}
