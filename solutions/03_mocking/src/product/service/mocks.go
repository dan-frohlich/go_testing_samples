package service 

//CatalogMock mock a Catalog
type CatalogMock struct {
	search func(query string) []SearchHit
	get    func(id string) Product
}

//Search mock search
func (c *CatalogMock) Search(query string) []SearchHit {
	return c.search(query)
}

//Get mock get
func (c *CatalogMock) Get(id string) Product {
	return c.get(id)
}

//TweetAPIMock mock a TweetAPI
type TweetAPIMock struct {
	query func(query string) []Tweet
}

//Query mock query
func (api *TweetAPIMock) Query(query string) []Tweet {
	return api.query(query)
}


