package url

import (
	"net/http"
	response "urlshortener/utils"

	"github.com/gin-gonic/gin"
)

type UrlController struct {
	urlService *UrlService
}

func NewUrlController(
	urlService *UrlService,
) *UrlController {
	return &UrlController{
		urlService: urlService,
	}
}

func (ctx *UrlController) ShortenUrl(c *gin.Context) {
	shortUrl, err := ctx.urlService.ShortenUrl()
	if err != nil {
		response.Error(c, "Error to shorten URL, contact support.")
		return
	}

	response.Data(c, shortUrl, "Url shortened successfully.", http.StatusOK)
}
