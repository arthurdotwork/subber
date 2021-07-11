package pubsub

import (
	"context"
	"os"

	"cloud.google.com/go/pubsub"
	"github.com/pterm/pterm"
)

func NewPubsubClient(ctx context.Context, projectId string, host string) (*pubsub.Client, error) {
	os.Setenv("PUBSUB_EMULATOR_HOST", host)
	client, err := pubsub.NewClient(ctx, projectId)
	if err != nil {
		pterm.Error.Println("Cannot start pubsub client :", err)
		return nil, err
	}

	return client, nil
}
