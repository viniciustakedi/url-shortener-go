package url

type UrlService struct{}

func NewHealthService() *UrlService {
	return &UrlService{}
}

func (ctx *UrlService) ShortenUrl() (string, error) {
	return "http://takedi.dev/T9K1LD", nil
}
