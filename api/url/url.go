package url

import "time"

type UrlPayload struct {
	Url          string `json:"url" binding:"required"`
	DaysToExpire int    `json:"days_to_expire" binding:"required"`
}

type UrlGetParam struct {
	ShortUrl string `uri:"shortUrl" binding:"required"`
}

type UrlDB struct {
	Domain         string    `json:"domain"`
	UrlCode        string    `json:"url_code"`
	OriginalUrl    string    `json:"original_url"`
	ExpirationDate time.Time `json:"expiration_date"`
	CreatedAt      time.Time `json:"created_at"`
}
