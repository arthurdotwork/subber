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

// readSubCmd represents the readSub command
var readSubCmd = &cobra.Command{
	Use: "readSub",
	Run: func(cmd *cobra.Command, args []string) {
		subName, err := service.NewPrompt("Please enter a subscriptionName", func(value string) error {
			if len(value) == 0 {
				return errors.New("sub name cannot be null")
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
		if err = pubSubService.ReadSub(ctx, subName); err != nil {
			pterm.Error.Printfln("Cannot read from subscription %s. (%s)", subName, err.Error())
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(readSubCmd)
}
