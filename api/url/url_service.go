package url

import (
	"context"
	"fmt"
	"time"
	"urlshortener/utils/code"

	"go.mongodb.org/mongo-driver/mongo"
)

const (
	domain = "https://takedi.com"
)

type UrlService struct {
	mongoDB *mongo.Database
}

func NewUrlService(mongoDB *mongo.Database) *UrlService {
	return &UrlService{
		mongoDB: mongoDB,
	}
}

func (ctx *UrlService) ShortenUrl(data UrlPayload) (string, error) {
	urlsCollection := ctx.mongoDB.Collection("urls")

	ctxBg := context.Background()

	expireDate := time.Now().Add(time.Hour * 24 * time.Duration(data.DaysToExpire))
	urlCode := code.GenerateRandom(6)

	var existingUrl UrlDB
	err := urlsCollection.FindOne(ctxBg, map[string]interface{}{
		"OriginalUrl": data.Url,
	}).Decode(&existingUrl)
	if err == nil {
		return domain + "/" + existingUrl.UrlCode, nil
	}

	// Prepare the data to insert
	dataToInsert := UrlDB{
		Domain:         domain,
		UrlCode:        urlCode,
		OriginalUrl:    data.Url,
		ExpirationDate: expireDate,
		CreatedAt:      time.Now(),
	}

	_, err = urlsCollection.InsertOne(ctxBg, dataToInsert)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s/%s", domain, urlCode), nil
}

func (ctx *UrlService) GetOriginalUrl(urlCode string) (string, error) {
	urlsCollection := ctx.mongoDB.Collection("urls")

	ctxBg := context.Background()

	var existingUrl UrlDB
	err := urlsCollection.FindOne(ctxBg, map[string]interface{}{
		"urlCode": urlCode,
	}).Decode(&existingUrl)
	if err == mongo.ErrNoDocuments {
		return "", fmt.Errorf("node url was found")
	} else if err != nil {
		return "", err
	}

	if existingUrl.ExpirationDate.Before(time.Now()) {
		return "", fmt.Errorf("url expired")
	}

	return existingUrl.OriginalUrl, nil
}

func (ctx *UrlService) DeleteExpiredUrls() (string, error) {
	urlsCollection := ctx.mongoDB.Collection("urls")

	ctxBg := context.Background()

	result, err := urlsCollection.DeleteMany(ctxBg, map[string]interface{}{
		"expirationDate": map[string]interface{}{
			"$lt": time.Now(),
		},
	})
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%d urls were deleted", result.DeletedCount), nil
}
