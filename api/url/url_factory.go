package url

func MakeUrlController() *UrlController {
	urlService := NewHealthService()
	urlController := NewUrlController(urlService)

	return urlController
}
