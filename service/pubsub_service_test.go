package service_test

import (
	"context"
	"testing"

	"cloud.google.com/go/pubsub"
	"cloud.google.com/go/pubsub/pstest"
	"github.com/arthureichelberger/subber/model"
	"github.com/arthureichelberger/subber/service"
	"github.com/stretchr/testify/assert"
	"google.golang.org/api/option"
	"google.golang.org/grpc"
)

func TestPubSubService(t *testing.T) {
	ctx := context.Background()

	srv := pstest.NewServer()
	defer srv.Close()

	conn, err := grpc.Dial(srv.Addr, grpc.WithInsecure())
	assert.NoError(t, err)
	defer conn.Close()

	client, err := pubsub.NewClient(ctx, "project", option.WithGRPCConn(conn))
	assert.NoError(t, err)
	defer client.Close()

	pubSubService := service.NewPubSubService(client)

	t.Run("It should be able to create a Topic.", func(t *testing.T) {
		err := pubSubService.CreateTopic(ctx, "subber")
		assert.NoError(t, err)
	})

	t.Run("It should be able to create a Subscription.", func(t *testing.T) {
		err := pubSubService.CreateSub(ctx, "subber", "subber")
		assert.NoError(t, err)
	})

	t.Run("It should be able to retrieve the list of all Topics.", func(t *testing.T) {
		topics, err := pubSubService.ListTopics(ctx)
		assert.NoError(t, err)
		assert.Equal(t, 1, len(topics))
		assert.Equal(t, "projects/project/topics/subber", topics[0])
	})

	t.Run("It should be able to retrieve the list of all Subscriptions.", func(t *testing.T) {
		subs, err := pubSubService.ListSubs(ctx)
		assert.NoError(t, err)
		assert.Equal(t, 1, len(subs))
		assert.Equal(t, "projects/project/topics/subber", subs["projects/project/subscriptions/subber"])
	})

	t.Run("It should be able to retrieve messages from an existing Subscription.", func(t *testing.T) {
		topic := client.Topic("subber")
		_ = pubSubService.Publish(ctx, topic.ID(), "arthur")
		c := make(chan model.Message)
		maxMessages := uint(1)
		go func() {
			err := pubSubService.ReadSub(ctx, "subber", c, maxMessages)
			assert.NoError(t, err)
		}()

		for {
			msg := <-c
			assert.Equal(t, "arthur", string(msg.Message))
			assert.Equal(t, maxMessages, msg.Id)

			if msg.Id == maxMessages {
				close(c)
				return
			}
		}
	})
}
