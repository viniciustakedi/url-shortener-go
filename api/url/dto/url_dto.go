package urldto

import "fmt"

func CheckPostPayload(url string, daysToExpire int) error {
	if url == "" {
		return fmt.Errorf("url is required")
	}
	if len(url) < 5 {
		return fmt.Errorf("url must be at least 5 characters long")
	}
	if len(url) > 2048 {
		return fmt.Errorf("url must be less than 2048 characters long")
	}
	if daysToExpire < 1 {
		return fmt.Errorf("days to expire must be greater than 0")
	}
	if daysToExpire > 30 {
		return fmt.Errorf("days to expire must be less than 30")
	}
	return nil
}

func CheckGetPayload(shortUrl string) error {
	if shortUrl == "" {
		return fmt.Errorf("short url is required")
	}
	if len(shortUrl) < 5 {
		return fmt.Errorf("short url must be at least 5 characters long")
	}
	if len(shortUrl) > 2048 {
		return fmt.Errorf("short url must be less than 2048 characters long")
	}
	return nil
}
