package cmd

import (
	"context"
	"fmt"

	"github.com/arthureichelberger/subber/pkg/pubsub"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var createTopicCmd = &cobra.Command{
	Use: "createTopic",
	Run: func(cmd *cobra.Command, args []string) {
		topicName := args[0]

		ctx := context.Background()
		client, err := pubsub.NewPubsubClient(ctx, fmt.Sprintf("%v", viper.Get("PUBSUB_PROJECT_ID")))
		cobra.CheckErr(err)

		_, err = client.CreateTopic(ctx, topicName)
		if err != nil {
			pterm.Error.Printfln("Could not create topic %s.", topicName)
			return
		}

		pterm.Success.Printfln("Topic %s created.\n", topicName)
	},
	Args: cobra.ExactArgs(1),
}

func init() {
	rootCmd.AddCommand(createTopicCmd)
}
