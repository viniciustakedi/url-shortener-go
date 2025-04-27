package url

import (
	"net/http"
	urldto "urlshortener/api/url/dto"
	response "urlshortener/utils/response"

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
	payload := c.MustGet("payload").(*UrlPayload)

	if err := urldto.CheckPostPayload(payload.Url, payload.DaysToExpire); err != nil {
		response.Error(c, err.Error())
		return
	}

	shortUrl, err := ctx.urlService.ShortenUrl(*payload)
	if err != nil {
		response.Error(c, "Error to shorten URL, contact support.")
		return
	}

	response.Data(c, shortUrl, "Url shortened successfully.", http.StatusOK)
}
