package cmd

import (
	"context"
	"errors"
	"fmt"

	"github.com/arthureichelberger/subber/pkg/pubsub"
	"github.com/arthureichelberger/subber/service"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var createTopicCmd = &cobra.Command{
	Use: "createTopic",
	Run: func(cmd *cobra.Command, args []string) {
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

		_, err = client.CreateTopic(ctx, topicName)
		if err != nil {
			pterm.Error.Printfln("Could not create topic %s. (%s)", topicName, err.Error())
			return
		}

		pterm.Success.Printfln("Topic %s created.", topicName)
	},
}

func init() {
	rootCmd.AddCommand(createTopicCmd)
}
