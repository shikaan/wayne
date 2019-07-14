package internal

import (
	"context"
	"log"
	"net/http"
	"strings"
)

// Redirect test
func Redirect(writer http.ResponseWriter, request *http.Request) {
	readableHash := request.URL.Path
	readableHash = strings.TrimPrefix(readableHash, "/")

	ctx := context.Background()
	client, databaseError := GetDatabaseClient(ctx)

	if databaseError != nil {
		log.Panic(databaseError)
		return
	}

	query := client.NewRef("urls").OrderByChild("readableHash").EqualTo(readableHash)

	queryResults, queryError := query.GetOrdered(ctx)

	if queryError != nil {
		log.Print("Query Error!")
		log.Panic(queryError)
		return
	}

	for _, result := range queryResults {
		var url URLRecord
		if err := result.Unmarshal(&url); err != nil {
			log.Fatalln("Error unmarshaling result:", err)
			return
		}

		http.Redirect(writer, request, url.TargetURL, 301)
		return
	}
}
