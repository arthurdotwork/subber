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

var stdinPayload string
var publishTopicNameOption string

// publishCmd represents the publish command
var publishCmd = &cobra.Command{
	Use:   "publish",
	Short: "publish allows to publish message in a topic.",
	Run: func(cmd *cobra.Command, args []string) {
		var topicName string
		var err error
		if publishTopicNameOption == "" {
			topicName, err = service.NewPrompt("Please enter the topicName", func(value string) error {
				if len(value) == 0 {
					return errors.New("topic name cannot be null")
				}

				return nil
			})

			if err != nil {
				pterm.Error.Println(err.Error())
				return
			}
		} else {
			topicName = publishTopicNameOption
		}

		var payload string
		if stdinPayload == "" {
			payload, err = service.NewPrompt("Please enter your payload", func(value string) error {
				if len(value) == 0 {
					return errors.New("payload cannot be null")
				}

				return nil
			})
		} else {
			payload = stdinPayload
		}

		if err != nil {
			pterm.Error.Println(err.Error())
			return
		}

		ctx := context.Background()
		client, err := pubsub.NewPubsubClient(ctx, fmt.Sprintf("%v", viper.Get("PUBSUB_PROJECT_ID")), fmt.Sprintf("%v", viper.Get("EMULATOR_HOST")))
		if err != nil {
			pterm.Error.Println(err.Error())
			return
		}

		pubSubService := service.NewPubSubService(client)
		if err = pubSubService.Publish(ctx, topicName, payload); err != nil {
			pterm.Error.Println(err.Error())
			return
		}

		pterm.Success.Printfln("Message published.")
	},
}

func init() {
	rootCmd.AddCommand(publishCmd)

	publishCmd.Flags().StringVar(&stdinPayload, "payload", "", "Payload that will be published in the message data")
	publishCmd.Flags().StringVar(&publishTopicNameOption, "topic", "", "The topic in which you want your message to be sent")
}
