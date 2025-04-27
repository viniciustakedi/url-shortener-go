package cron

import (
	urlcron "urlshortener/cron/url"
)

func Init() {
	urlcron.DeleteExpiredUrls()
	// Add more cron jobs here as needed
}
