package middlewares

import (
	"fmt"
	"reflect"
	response "urlshortener/utils/response"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func PayloadValidatorMiddleware(payload interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		payloadType := reflect.TypeOf(payload)
		if payloadType.Kind() == reflect.Ptr {
			payloadType = payloadType.Elem()
		}

		payloadInstance := reflect.New(payloadType).Interface()

		if err := c.ShouldBindJSON(payloadInstance); err != nil {
			response.Error(c, "Invalid request payload")
			c.Abort()
			return
		}

		validate := validator.New()

		if err := validate.Struct(payloadInstance); err != nil {
			fmt.Printf("Orror binding JSON: %v\n", err)
			response.Error(c, "Validation failed")
			c.Abort()
			return
		}

		c.Set("payload", payloadInstance)
		c.Next()
	}
}
