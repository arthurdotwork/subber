package cmd

import (
	"context"
	"errors"
	"fmt"

	ps "cloud.google.com/go/pubsub"
	"github.com/arthureichelberger/subber/pkg/pubsub"
	"github.com/arthureichelberger/subber/service"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var createSubCmd = &cobra.Command{
	Use: "createSub",
	Run: func(cmd *cobra.Command, args []string) {
		subName, err := service.NewPrompt("Please enter the subscriptionName", func(value string) error {
			if len(value) == 0 {
				return errors.New("sub name cannot be null")
			}

			return nil
		})

		if err != nil {
			pterm.Error.Println(err.Error())
			return
		}

		topicName, err := service.NewPrompt("Please enter the topicName", func(value string) error {
			if len(value) == 0 {
				return errors.New("topic name cannot be null")
			}

			return nil
		})

		if err != nil {
			pterm.Error.Println(err.Error())
			return
		}

		ctx := context.Background()
		client, err := pubsub.NewPubsubClient(ctx, fmt.Sprintf("%v", viper.Get("PUBSUB_PROJECT_ID")))
		cobra.CheckErr(err)

		topic := client.Topic(topicName)
		if ok, err := topic.Exists(ctx); !ok || err != nil {
			pterm.Error.Printfln("Could not get topic %s.", topicName)
			return
		}

		_, err = client.CreateSubscription(ctx, subName, ps.SubscriptionConfig{
			Topic: topic,
		})
		if err != nil {
			pterm.Error.Printfln("Could not create subscription %s. (%s)", subName, err.Error())
			return
		}

		pterm.Success.Printfln("Subscription %s created.", subName)
	},
}

func init() {
	rootCmd.AddCommand(createSubCmd)
}
