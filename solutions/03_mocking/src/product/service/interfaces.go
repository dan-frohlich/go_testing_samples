package service

//Service api service
type Service interface {
	Mashup(query string) []Mashup
}

//Catalog product catalog datasource
type Catalog interface {
	Search(query string) []SearchHit
	Get(id string) Product
}

//TweetAPI twitter interface
type TweetAPI interface {
	Query(query string) []Tweet
}

//SearchHit for Catalog search interface
type SearchHit struct {
	Query, ProductID string
	Confidence       float64
}

//Product represents a product in the catalog
type Product struct {
	ID, Name string
}

//Tweet an entity
type Tweet struct {
	Timestamp int64
	Text      string
}

//Mashup an entity
type Mashup struct {
	Product Product
	Tweets  []Tweet
}
