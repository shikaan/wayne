package internal

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
)

// StoreURL test
func StoreURL(writer http.ResponseWriter, request *http.Request) {
	ctx := context.Background()
	client, databaseError := GetDatabaseClient(ctx)

	if databaseError != nil {
		log.Panic(databaseError)
		return
	}

	allURLsQuery := client.NewRef("urls")

	var url URLRecord

	decodeErr := json.NewDecoder(request.Body).Decode(&url)

	if decodeErr != nil {
		log.Panic(databaseError)
		return
	}

	newURLQuery, newURLQueryError := allURLsQuery.Push(ctx, &url)

	if newURLQueryError != nil {
		log.Panic(newURLQueryError)

		return
	}

	fetchError := newURLQuery.Get(ctx, &url)

	if fetchError != nil {
		log.Panic(fetchError)
		return
	}

	writer.WriteHeader(201)
	json.NewEncoder(writer).Encode(&url)
}
