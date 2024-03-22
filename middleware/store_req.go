package middleware

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"gitlab.com/hedwig-phan/assignment-3/cmd/db"
	"gitlab.com/hedwig-phan/assignment-3/ent"
	"gitlab.com/hedwig-phan/assignment-3/internal/util"
)

// func MiddlewareAPILogging() fiber.Handler {
// 	return func(c *fiber.Ctx) error {
// 		_ = c.Next()

// 		//Request
// 		requestTime := time.Now()
// 		requestParametersStr := c.Queries()
// 		requestBodyBytes := c.Body()

// 		requestHeadersStr := c.GetReqHeaders()
// 		requestMetadata := map[string]interface{}{
// 			"method": c.Method(),
// 			"path":   c.Path(),
// 		}

// 		// Convert requestParameters from map[string]string to map[string]interface{}
// 		requestParameters := make(map[string]interface{})
// 		for key, value := range requestParametersStr {
// 			requestParameters[key] = value
// 		}

// 		// Convert requestBody to map[string]interface{}
// 		var requestBody map[string]interface{}
// 		json.Unmarshal(requestBodyBytes, &requestBody)

// 		// Convert requestHeaders from map[string]string to map[string]interface{}
// 		requestHeaders := make(map[string]interface{})
// 		for key, value := range requestHeadersStr {
// 			requestHeaders[key] = value
// 		}

// 		// Store into database
// 		_, err := StoreAPIRequestRespond(util.Ctx, requestTime,
// 			requestParameters, requestHeaders, requestBody, requestMetadata, time.Now(), 0, nil, nil)

// 		if err != nil {
// 			return err
// 		}

// 		return nil
// 	}
// }

func MiddlewareAPILogging() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Your existing middleware logic
			requestTime := time.Now()
			requestParametersStr := r.URL.Query()
			requestBodyBytes := []byte{} // Read the request body here

			requestHeadersStr := r.Header
			requestMetadata := map[string]interface{}{
				"method": r.Method,
				"path":   r.URL.Path,
			}

			// Convert requestParameters from map[string][]string to map[string]interface{}
			requestParameters := make(map[string]interface{})
			for key, value := range requestParametersStr {
				requestParameters[key] = value
			}

			// Convert requestBody to map[string]interface{}
			var requestBody map[string]interface{}
			json.Unmarshal(requestBodyBytes, &requestBody)

			// Convert requestHeaders from map[string][]string to map[string]interface{}
			requestHeaders := make(map[string]interface{})
			for key, value := range requestHeadersStr {
				requestHeaders[key] = value
			}

			// Store into database (replace with your actual database logic)
			_, err := StoreAPIRequestRespond(util.Ctx, requestTime,
				requestParameters, requestHeaders, requestBody, requestMetadata, time.Now(), 0, nil, nil)

			if err != nil {
				log.Fatal(err)
			}

			// Call the next handler
			next.ServeHTTP(w, r)
		})
	}
}

func StoreAPIRequestRespond(ctx context.Context,
	requestTime time.Time,
	requestParameter map[string]interface{},
	requestHeaders map[string]interface{},
	requestBody map[string]interface{},
	requestMetadata map[string]interface{},
	responseTime time.Time,
	responseStatus int,
	responseHeaders map[string]interface{},
	responseBody map[string]interface{}) (*ent.Apireq, error) {

	currentTime := time.Now()

	apireq := db.Client.Apireq.Create().
		SetReqTime(requestTime).
		SetReqParam(requestParameter).
		SetReqHeaders(requestHeaders).
		SetReqBody(requestBody).
		SetReqMetadata(requestMetadata).
		SetRespTime(responseTime).
		SetRespStatus(responseStatus).
		SetRespHeaders(responseHeaders).
		SetRespBody(responseBody).
		SetCreatedAt(currentTime).
		SetUpdatedAt(currentTime).
		SaveX(ctx)

	return apireq, nil
}
