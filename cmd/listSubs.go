package cmd

import (
	"context"
	"fmt"

	"github.com/arthureichelberger/subber/pkg/pubsub"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"google.golang.org/api/iterator"
)

// listSubsCmd represents the listSubs command
var listSubsCmd = &cobra.Command{
	Use: "listSubs",
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		client, err := pubsub.NewPubsubClient(ctx, fmt.Sprintf("%v", viper.Get("PUBSUB_PROJECT_ID")))
		cobra.CheckErr(err)

		subscriptions := client.Subscriptions(ctx)
		for {
			sub, err := subscriptions.Next()
			if err == iterator.Done {
				break
			}

			subConfig, err := sub.Config(ctx)
			if err != nil {
				pterm.Error.Println(err.Error())
				return
			}
			pterm.Info.Printfln("Subscription : %s. (topic: %s)", sub.String(), subConfig.Topic.String())
		}
	},
}

func init() {
	rootCmd.AddCommand(listSubsCmd)
}
