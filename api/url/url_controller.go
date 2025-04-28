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

	response.Data(c, shortUrl, "Url shortened successfully.", http.StatusCreated)
}

func (ctx *UrlController) GetOriginalUrl(c *gin.Context) {
	shortUrl := c.Param("shortUrl")
	if err := urldto.CheckGetPayload(shortUrl); err != nil {
		response.Error(c, err.Error())
		return
	}

	urlReponse, err := ctx.urlService.GetOriginalUrl(shortUrl)
	if err != nil {
		response.Error(c, "Error to get original URL, contact support - "+err.Error())
		return
	}

	response.Data(c, urlReponse, "Original URL found successfully!", http.StatusOK)
}

func (ctx *UrlController) DeleteExpiredUrlsWithCron() (string, error) {
	msg, err := ctx.urlService.DeleteExpiredUrls()
	if err != nil {
		return "", err
	}

	return msg, nil
}
