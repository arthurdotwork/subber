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

var createSubCmd = &cobra.Command{
	Use:   "createSub",
	Short: "createTopics allows to create a subscription on the emulator.",
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
		client, err := pubsub.NewPubsubClient(ctx, fmt.Sprintf("%v", viper.Get("PUBSUB_PROJECT_ID")), fmt.Sprintf("%v", viper.Get("EMULATOR_HOST")))
		if err != nil {
			pterm.Error.Println(err.Error())
			return
		}

		pubSubService := service.NewPubSubService(client)
		if err = pubSubService.CreateSub(ctx, subName, topicName); err != nil {
			pterm.Error.Printfln("Cannot create sub %s. (%s)", subName, err.Error())
			return
		}

		pterm.Success.Printfln("Subscription %s created.", subName)
	},
}

func init() {
	rootCmd.AddCommand(createSubCmd)
}
