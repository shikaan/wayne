package internal

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"os"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/db"
	"google.golang.org/api/option"
)

// FirebaseAuth lol
type FirebaseAuth struct {
	ProjectID string `json:"project_id"`
}

// GetDatabaseClient TODO: make me a singleton!
func GetDatabaseClient(ctx context.Context) (*db.Client, error) {
	firebaseAuthData, ioError := ioutil.ReadFile(os.Getenv("FIREBASE_AUTH_JSON"))

	if ioError != nil {
		return nil, ioError
	}

	var firebaseAuth FirebaseAuth

	jsonError := json.Unmarshal(firebaseAuthData, &firebaseAuth)

	if jsonError != nil {
		return nil, jsonError
	}

	opt := option.WithCredentialsFile(os.Getenv("FIREBASE_AUTH_JSON"))

	config := &firebase.Config{
		ProjectID:   firebaseAuth.ProjectID,
		DatabaseURL: "https://" + firebaseAuth.ProjectID + ".firebaseio.com",
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
