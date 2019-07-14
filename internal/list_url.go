package internal

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
)

// URLRecord lol
type URLRecord struct {
	ReadableHash string `json:"readableHash,omitempty"`
	TargetURL    string `json:"targetURL,omitempty"`
}

// URLEntity lol
type URLEntity struct {
	ID           string `json:"id"`
	ReadableHash string `json:"readableHash,omitempty"`
	TargetURL    string `json:"targetURL,omitempty"`
}

// ListURLs asdf
func ListURLs(writer http.ResponseWriter, request *http.Request) {
	ctx := context.Background()
	client, databaseError := GetDatabaseClient(ctx)

	if databaseError != nil {
		log.Panic(databaseError)
		return
	}

	query := client.NewRef("urls").OrderByKey()

	var queryResult map[string]URLRecord

	if queryError := query.Get(ctx, &queryResult); queryError != nil {
		log.Print("Query Error!")
		log.Panic(queryError)
		return
	}

	var result []URLEntity

	for key, value := range queryResult {
		entity := URLEntity{
			ID:           key,
			ReadableHash: value.ReadableHash,
			TargetURL:    value.TargetURL,
		}

		result = append(result, entity)
	}

	item, err := json.Marshal(result)

	if err != nil {
		log.Panic(err)
	}

	writer.WriteHeader(200)
	writer.Write([]byte(item))
}
