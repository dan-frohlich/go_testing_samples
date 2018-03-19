package service

//CatalogMock mock a Catalog
type CatalogMock struct {
}

//Search mock search
func (c *CatalogMock) Search(query string) []SearchHit {
}

//Get mock get
func (c *CatalogMock) Get(id string) Product {
}

//TweetAPIMock mock a TweetAPI
type TweetAPIMock struct {
}

//Query mock query
func (api *TweetAPIMock) Query(query string) []Tweet {
}
