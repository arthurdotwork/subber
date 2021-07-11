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

// listSubsCmd represents the listSubs command
var listSubsCmd = &cobra.Command{
	Use: "listSubs",
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		client, err := pubsub.NewPubsubClient(ctx, fmt.Sprintf("%v", viper.Get("PUBSUB_PROJECT_ID")), fmt.Sprintf("%v", viper.Get("EMULATOR_HOST")))
		if err != nil {
			pterm.Error.Println(err.Error())
			return
		}

		pubSubService := service.NewPubSubService(client)
		subs, err := pubSubService.ListSubs(ctx)
		if err != nil {
			pterm.Error.Println(err.Error())
			return
		}

		for sub, topic := range subs {
			pterm.Info.Printfln("Subscription : %s. (topic: %s)", sub, topic)
		}

	},
}

func init() {
	rootCmd.AddCommand(listSubsCmd)
}
