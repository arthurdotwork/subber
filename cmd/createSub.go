package cmd

import (
	"context"
	"fmt"

	ps "cloud.google.com/go/pubsub"
	"github.com/arthureichelberger/subber/pkg/pubsub"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var createSubCmd = &cobra.Command{
	Use: "createSub",
	Run: func(cmd *cobra.Command, args []string) {
		subName := args[0]
		topicName := args[1]

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
			pterm.Error.Printfln("Could not create subscription %s.", subName)
			return
		}

		pterm.Success.Printfln("Subscription %s created.\n", subName)
	},
	Args: cobra.ExactArgs(2),
}

func init() {
	rootCmd.AddCommand(createSubCmd)
}
