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

var createTopicNameOption string

var createTopicCmd = &cobra.Command{
	Use:   "createTopic",
	Short: "createTopics allows to create a topic on the emulator.",
	Run: func(cmd *cobra.Command, args []string) {
		var topicName string
		var err error
		if createTopicNameOption == "" {
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
			topicName = createTopicNameOption
		}

		ctx := context.Background()
		client, err := pubsub.NewPubsubClient(ctx, fmt.Sprintf("%v", viper.Get("PUBSUB_PROJECT_ID")), fmt.Sprintf("%v", viper.Get("EMULATOR_HOST")))
		if err != nil {
			pterm.Error.Println(err.Error())
			return
		}

		pubSubService := service.NewPubSubService(client)
		if err = pubSubService.CreateTopic(ctx, topicName); err != nil {
			pterm.Error.Printfln("Cannot create topic %s. (%s)", topicName, err.Error())
			return
		}

		pterm.Success.Printfln("Topic %s created.", topicName)
	},
}

func init() {
	rootCmd.AddCommand(createTopicCmd)

	createTopicCmd.Flags().StringVar(&createTopicNameOption, "topic", "", "The name of the topic you want to create")
}
