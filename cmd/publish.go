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

// publishCmd represents the publish command
var publishCmd = &cobra.Command{
	Use: "publish",
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
		client, err := pubsub.NewPubsubClient(ctx, fmt.Sprintf("%v", viper.Get("PUBSUB_PROJECT_ID")), fmt.Sprintf("%v", viper.Get("EMULATOR_HOST")))
		if err != nil {
			pterm.Error.Println(err.Error())
			return
		}

		pubSubService := service.NewPubSubService(client)
		if err = pubSubService.Publish(ctx, topicName, "arthur"); err != nil {
			pterm.Error.Println(err.Error())
			return
		}

		pterm.Success.Printfln("Message published.")
	},
}

func init() {
	rootCmd.AddCommand(publishCmd)
}
