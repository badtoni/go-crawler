package crawlers

import (
	"database/sql"
	"encoding/xml"
	"go-crawler/crud"
	"go-crawler/database"
	"go-crawler/models/rssmodels"
	"log"
	"net/http"
	"time"
)

// CrawlRSSAndSaveToDB : Function to crawl the rss channels, defined in the db, and afterwards save them in the db
func CrawlRSSAndSaveToDB() {

	db := database.GetDB()

	rssChannels := crud.GetRssChannels(db)

	// Loop through all Channels
	for rssChannels.Next() {

		var rssChannel rssmodels.SimpleChannel

		if err := rssChannels.Scan(&rssChannel.ID, &rssChannel.URL); err != nil {
			// do something with error
		} else {
			log.Print("Querying: " + rssChannel.URL)
			log.Printf("Consuming feeds - Current Time: %v\n", time.Now())
			consumeArticlesAndSave(db, rssChannel.URL, rssChannel.ID)
		}
	}
}

// CrawlRSSAndReturn : Function to crawl the rss channels, defined in the db, and afterwards return them back
func CrawlRSSAndReturn() interface{} {
	var rssfeeds []rssmodels.Channel

	db := database.GetDB()

	rssChannels := crud.GetRssChannels(db)

	for rssChannels.Next() {

		var rssChannel rssmodels.SimpleChannel

		if err := rssChannels.Scan(&rssChannel.ID, &rssChannel.URL); err != nil {
			// do something with error
		} else {
			log.Print("Querying: " + rssChannel.URL)
			log.Printf("Consuming feeds - Current Time: %v\n", time.Now())

			// TODO change functions to return the rss feeds
			rssfeeds = append(rssfeeds, consumeArticlesAndReturn(db, rssChannel.URL, rssChannel.ID))
		}
	}
	return rssfeeds
}

// consumeArticlesAndReturn : gets and saves all the articles for the given rss feed channel
func consumeArticlesAndReturn(db *sql.DB, rssChannelURL string, rssChannelID string) rssmodels.Channel {
	// Get rss response from rss Channel
	resp, err := http.Get(rssChannelURL)
	if err != nil {
		log.Printf("Error GET consumeArticles: %v\n", err.Error())
		return rssmodels.Channel{}
	}
	defer resp.Body.Close()

	rss := rssmodels.Rss{}

	// Create a new xml parser
	decoder := xml.NewDecoder(resp.Body)
	// Docodes the response body and saves it into the tss variable
	err = decoder.Decode(&rss)
	if err != nil {
		log.Printf("Error Decode: %v\n", err)
		return rssmodels.Channel{}
	}

	log.Printf("Channel title: %v\n", rss.Channel.Title)
	log.Printf("Channel link: %v\n", rss.Channel.Link)

	//Loop through rss feed link articles:
	for i, item := range rss.Channel.Items {
		log.Printf("%v. item title: %v\n", i, item.Title)
		log.Printf("%v. item date: %v\n", i, item.PubDate)
		// log.Printf("%v. item guid: %v\n", i, item.Guid)

		// var olditem rssmodels.Item
		// olditem = getArticle(item.Guid, db)

		// // If olditem/article is empty, save it in the db
		// if (rssmodels.Item{}) == olditem {
		// 	/* write rss article to db */
		// 	postArticle(item, db, rssChannelID)

		// } else {
		// 	// log.Printf("%v. olditem guid: %v\n", i, olditem.Guid)
		// 	log.Printf("Article already exists in the DB\n")
		// }
	}
	return rss.Channel
}

// consumeArticlesAndSave : gets and saves all the articles for the given channel
func consumeArticlesAndSave(db *sql.DB, rssChannelURL string, rssChannelID string) {

	// Get rss response from rss Channel
	resp, err := http.Get(rssChannelURL)
	if err != nil {
		log.Printf("Error GET consumeArticles: %v\n", err.Error())
		return
	}
	defer resp.Body.Close()

	rss := rssmodels.Rss{}

	decoder := xml.NewDecoder(resp.Body)
	err = decoder.Decode(&rss)
	if err != nil {
		log.Printf("Error Decode: %v\n", err)
		return
	}

	log.Printf("Channel title: %v\n", rss.Channel.Title)
	log.Printf("Channel link: %v\n", rss.Channel.Link)

	//Loop through rss feed link articles:
	for i, item := range rss.Channel.Items {
		log.Printf("%v. item title: %v\n", i, item.Title)
		log.Printf("%v. item date: %v\n", i, item.PubDate)
		// log.Printf("%v. item guid: %v\n", i, item.Guid)

		var olditem rssmodels.Item
		olditem = getArticle(item.GUID, db)

		// If olditem/article is empty, save it in the db
		if (rssmodels.Item{}) == olditem {
			/* write rss article to db */
			postArticle(item, db, rssChannelID)

		} else {
			// log.Printf("%v. olditem guid: %v\n", i, olditem.Guid)
			log.Printf("Article already exists in the DB\n")
		}
	}
}

// getArticle : function querying for a single article from the db, based on the given guid
func getArticle(guid string, db *sql.DB) rssmodels.Item {
	sqlStatement := `SELECT id, title, link, description, guid FROM rss_feed WHERE guid=$1;`
	var item rssmodels.Item
	var id int
	row := db.QueryRow(sqlStatement, guid)
	switch err := row.Scan(&id, &item.Title, &item.Link, &item.Desc, &item.GUID); err {
	case sql.ErrNoRows:
		// log.Println("No rows were returned!")
	case nil:
		// log.Println(id, item.Title, item.Link, item.Desc, item.Guid, item.PubDate)
	default:
		panic(err)
	}
	return item
}

// checkIfNull : Function checks if string is null
func checkIfNull(s string) sql.NullString {
	if len(s) == 0 {
		return sql.NullString{}
	}
	return sql.NullString{
		String: s,
		Valid:  true,
	}
}

// postArticle : Function inserting a single article into the db, based on the given item
func postArticle(item rssmodels.Item, db *sql.DB, rssChannelID string) int {
	//Inserting records with database/sql

	//fmt.Print(item.Image.Link + " " + item.Image.Title + " " + item.Image.Url + "\n")

	sqlStatement := `INSERT INTO rss_feed (title, link, description, guid, pub_date, rssChannelID, crawled_date )
					VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
					RETURNING id`
	//Retrieve the ID of new records
	id := 0
	var err error
	// println(item.Title)
	crawledDate := "NOW()"

	err = db.QueryRow(sqlStatement, item.Title, item.Link, item.Desc, item.GUID, checkIfNull(item.PubDate), rssChannelID, crawledDate).Scan(&id)
	if err != nil {
		panic(err)
	}
	log.Println("New article ID is:", id)
	return id
}
