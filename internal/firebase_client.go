package internal

import (
	"context"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/db"
	"google.golang.org/api/option"
)

// GetDatabaseClient TODO: make me a singleton!
func GetDatabaseClient(ctx context.Context) (*db.Client, error) {
	opt := option.WithCredentialsFile("/home/manuel/dev/oss/short-go/firebase_auth.json")
	config := &firebase.Config{
		ProjectID:   "go-short-56756",
		DatabaseURL: "https://go-short-56756.firebaseio.com",
	}

	client, err := firebase.NewApp(ctx, config, opt)

	if err != nil {
		return nil, err
	}

	db, databaseError := client.Database(ctx)

	if databaseError != nil {
		return nil, err
	}

	return db, nil
}
