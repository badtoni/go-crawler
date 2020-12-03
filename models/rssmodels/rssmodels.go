package rssmodels

// Enclosure : model for
type Enclosure struct {
	URL    string `xml:"url,attr"`
	Length int64  `xml:"length,attr"`
	Type   string `xml:"type,attr"`
}

// Item : model for the articles
type Item struct {
	Title     string    `xml:"title"`
	Link      string    `xml:"link"`
	Desc      string    `xml:"description"`
	GUID      string    `xml:"guid"`
	Enclosure Enclosure `xml:"enclosure"`
	PubDate   string    `xml:"pubDate"`
}

// Channel : model for the rss channels
type Channel struct {
	Title string `xml:"title"`
	Link  string `xml:"link"`
	Desc  string `xml:"description"`
	Items []Item `xml:"item"`
}

// SimpleChannel : model for the saved channel in the db
type SimpleChannel struct {
	ID  string
	URL string
}

// Rss : model for the rss feeds from a channel
type Rss struct {
	Channel Channel `xml:"channel"`
}

// type SentimentAnalysisResult struct {
// 	Sentiment float32 `json:"sentiment"`
// }

// type NerEntity struct {
// 	EntityName string `json:"entity_name"`
// 	EntityType string `json:"entity_type"`
// }

// type NerAnalysisResult struct {
// 	Entities []NerEntity `json:"entities"`
// }

// type Image struct {
// 	Url   string `xml:"url,attr"`
// 	Title string `xml:"title,attr"`
// 	Link  string `xml:"link,attr"`
// }
