package service

import (
	"context"

	"github.com/arthureichelberger/subber/model"
)

type PubSubServiceInterface interface {
	CreateTopic(ctx context.Context, topicName string) error
	CreateSub(ctx context.Context, subName string, topicName string) error
	ListTopics(ctx context.Context) ([]string, error)
	ListSubs(ctx context.Context) (map[string]string, error)
	ReadSub(ctx context.Context, subName string, channel chan model.Message, maxMessages uint) error
	Publish(ctx context.Context, topicName string, payload string) error
}
