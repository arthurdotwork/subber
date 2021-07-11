package pubsub

import (
	"context"
	"log"

	"cloud.google.com/go/pubsub"
	"google.golang.org/api/option"
)

func NewPubsubClient(ctx context.Context, projectId string) (*pubsub.Client, error) {
	client, err := pubsub.NewClient(ctx, "local-emulator-project", option.WithCredentialsFile("{}"))
	if err != nil {
		log.Fatal("Cannot start pubsub client :", err)
		return nil, err
	}

	return client, nil
}
