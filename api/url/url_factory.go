package url

import "urlshortener/infra/db"

func MakeUrlController() *UrlController {
	urlService := NewUrlService(db.GetMongoDB())
	urlController := NewUrlController(urlService)

	return urlController
}
