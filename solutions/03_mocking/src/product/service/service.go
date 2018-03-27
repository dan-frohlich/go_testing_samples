package service

//NewService make a service
func NewService(cat Catalog, tweetapi TweetAPI) Service {
	return &service{cat, tweetapi}
}

type service struct {
	cat      Catalog
	tweetapi TweetAPI
}

func (s *service) Mashup(query string) []Mashup {
	hits := s.cat.Search(query)

	var result = []Mashup{}
	for _, hit := range hits {
		pid := hit.ProductID
		product := s.cat.Get(pid)
		tweets := s.tweetapi.Query(product.Name)
		result = append(result, Mashup{Product: product, Tweets: tweets})
	}
	return result
}
