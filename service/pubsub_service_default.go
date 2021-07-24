package service

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"cloud.google.com/go/pubsub"
	"github.com/pterm/pterm"
	"google.golang.org/api/iterator"
)

type PubSubService struct {
	Client *pubsub.Client
}

func NewPubSubService(client *pubsub.Client) PubSubServiceInterface {
	return PubSubService{
		Client: client,
	}
}

func (ps PubSubService) CreateTopic(ctx context.Context, topicName string) error {
	if _, err := ps.Client.CreateTopic(ctx, topicName); err != nil {
		return err
	}

	return nil
}

func (ps PubSubService) CreateSub(ctx context.Context, subName string, topicName string) error {
	topic := ps.Client.Topic(topicName)
	if ok, err := topic.Exists(ctx); !ok || err != nil {
		return errors.New("topic not found")
	}

	if _, err := ps.Client.CreateSubscription(ctx, subName, pubsub.SubscriptionConfig{Topic: topic}); err != nil {
		return err
	}

	return nil
}

func (ps PubSubService) ListTopics(ctx context.Context) ([]string, error) {
	topics := ps.Client.Topics(ctx)
	var topicsName []string
	for {
		topic, err := topics.Next()
		if err == iterator.Done {
			break
		}

		if err != nil {
			return nil, err
		}

		topicsName = append(topicsName, topic.String())
	}

	return topicsName, nil
}

func (ps PubSubService) ListSubs(ctx context.Context) (map[string]string, error) {
	subscriptions := ps.Client.Subscriptions(ctx)
	subs := make(map[string]string)
	for {
		sub, err := subscriptions.Next()
		if err == iterator.Done {
			break
		}

		subConfig, err := sub.Config(ctx)
		if err != nil {
			return nil, err
		}

		subs[sub.String()] = subConfig.Topic.String()
	}

	return subs, nil
}

func (ps PubSubService) ReadSub(ctx context.Context, subName string) error {
	sub := ps.Client.Subscription(subName)
	if ok, err := sub.Exists(ctx); !ok || err != nil {
		return errors.New("subscription does not exist")
	}

	// Consume 10 messages.
	var mu sync.Mutex
	received := 0
	cctx, cancel := context.WithCancel(ctx)
	err := sub.Receive(cctx, func(ctx context.Context, msg *pubsub.Message) {
		mu.Lock()
		defer mu.Unlock()
		pterm.Success.Printfln("Got message: %q", string(msg.Data))
		msg.Ack()
		received++
		if received == 10 {
			cancel()
		}
	})

	if err != nil {
		return fmt.Errorf("receive: %v", err)
	}

	return nil
}

func (ps PubSubService) Publish(ctx context.Context, topicName string, payload string) error {
	topic := ps.Client.Topic(topicName)
	ok, err := topic.Exists(ctx)

	if !ok || err != nil {
		return fmt.Errorf("topic %s does not exist", topicName)
	}
	result := topic.Publish(ctx, &pubsub.Message{
		Data: []byte(payload),
	})

	_, err = result.Get(ctx)
	if err != nil {
		return fmt.Errorf("could not publish message in topic. (%s)", err.Error())
	}

	return nil
}
