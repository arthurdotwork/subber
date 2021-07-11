package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var createTopicCmd = &cobra.Command{
	Use: "createTopic",
	Run: func(cmd *cobra.Command, args []string) {
		topicName := args[0]
		fmt.Printf("Creating topic %s.", topicName)
	},
	Args: cobra.ExactArgs(1),
}

func init() {
	rootCmd.AddCommand(createTopicCmd)
}
