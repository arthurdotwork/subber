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

// listTopicsCmd represents the listTopics command
var listTopicsCmd = &cobra.Command{
	Use: "listTopics",
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		client, err := pubsub.NewPubsubClient(ctx, fmt.Sprintf("%v", viper.Get("PUBSUB_PROJECT_ID")))
		cobra.CheckErr(err)

		topics := client.Topics(ctx)
		for {
			topic, err := topics.Next()
			if err == iterator.Done {
				break
			}

			pterm.Info.Printfln("Topic : %s.", topic.String())
		}
	},
}

func init() {
	rootCmd.AddCommand(listTopicsCmd)
}
