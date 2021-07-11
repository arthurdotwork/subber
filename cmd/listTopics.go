package cmd

import (
	"context"
	"fmt"

	"github.com/arthureichelberger/subber/pkg/pubsub"
	"github.com/arthureichelberger/subber/service"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// listTopicsCmd represents the listTopics command
var listTopicsCmd = &cobra.Command{
	Use: "listTopics",
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		client, err := pubsub.NewPubsubClient(ctx, fmt.Sprintf("%v", viper.Get("PUBSUB_PROJECT_ID")), fmt.Sprintf("%v", viper.Get("EMULATOR_HOST")))
		if err != nil {
			pterm.Error.Println(err.Error())
		}

		pubSubService := service.NewPubSubService(client)
		topics, err := pubSubService.ListTopics(ctx)

		if err != nil {
			pterm.Error.Println(err.Error())
			return
		}

		for _, topic := range topics {
			pterm.Info.Printfln("Topic : %s.", topic)
		}
	},
}

func init() {
	rootCmd.AddCommand(listTopicsCmd)
}
