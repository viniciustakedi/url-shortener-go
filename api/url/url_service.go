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

func (ctx *UrlService) ShortenUrl(data UrlPayload) (PostUrlResponse, error) {
	urlsCollection := ctx.mongoDB.Collection("urls")

	ctxBg := context.Background()

	expireDate := time.Now().Add(time.Hour * 24 * time.Duration(data.DaysToExpire))
	urlCode := code.GenerateRandom(6)

	var existingUrl UrlDB
	err := urlsCollection.FindOne(ctxBg, map[string]interface{}{
		"originalUrl": data.Url,
	}).Decode(&existingUrl)
	if err == nil {
		return PostUrlResponse{
			Url:            fmt.Sprintf("%s/%s", domain, existingUrl.UrlCode),
			ExpirationDate: existingUrl.ExpirationDate,
		}, nil
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
		return PostUrlResponse{}, err
	}

	return PostUrlResponse{
		Url:            fmt.Sprintf("%s/%s", domain, urlCode),
		ExpirationDate: expireDate,
	}, nil
}

func (ctx *UrlService) GetOriginalUrl(urlCode string) (GetUrlResponse, error) {
	urlsCollection := ctx.mongoDB.Collection("urls")

	ctxBg := context.Background()

	var existingUrl UrlDB
	err := urlsCollection.FindOne(ctxBg, map[string]interface{}{
		"urlCode": urlCode,
	}).Decode(&existingUrl)
	if err == mongo.ErrNoDocuments {
		return GetUrlResponse{}, fmt.Errorf("node url was found")
	} else if err != nil {
		return GetUrlResponse{}, err
	}

	if existingUrl.ExpirationDate.Before(time.Now()) {
		return GetUrlResponse{}, fmt.Errorf("url expired")
	}

	return GetUrlResponse{
		OriginalUrl:    existingUrl.OriginalUrl,
		ExpirationDate: existingUrl.ExpirationDate,
	}, nil
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
