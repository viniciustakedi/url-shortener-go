package url

import "time"

type UrlPayload struct {
	Url          string `json:"url" binding:"required"`
	DaysToExpire int    `json:"daysToExpire" binding:"required"`
}

type UrlGetParam struct {
	ShortUrl string `uri:"shortUrl" binding:"required"`
}

type UrlDB struct {
	Domain         string    `json:"domain" bson:"domain"`
	UrlCode        string    `json:"urlCode" bson:"urlCode"`
	OriginalUrl    string    `json:"originalUrl" bson:"originalUrl"`
	ExpirationDate time.Time `json:"expirationDate" bson:"expirationDate"`
	CreatedAt      time.Time `json:"createdAt" bson:"createdAt"`
}
