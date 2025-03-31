package router

import (
	"urlshortener/api/url"
	"urlshortener/middlewares"
	"urlshortener/models"

	"github.com/gin-gonic/gin"
)

func RegisterUrlRoutes(router *gin.RouterGroup) {
	urlController := url.MakeUrlController()

	routes := []struct {
		method      string
		route       string
		handler     gin.HandlerFunc
		middlewares []gin.HandlerFunc
	}{
		{
			method:      "POST",
			route:       "/url/shorten",
			handler:     urlController.ShortenUrl,
			middlewares: []gin.HandlerFunc{middlewares.PayloadValidatorMiddleware(&models.UrlPayload{})},
		},
	}

	for _, route := range routes {
		switch route.method {
		case "GET":
			router.GET(route.route, append(route.middlewares, route.handler)...)
		case "POST":
			router.POST(route.route, append(route.middlewares, route.handler)...)
		}
	}
}
