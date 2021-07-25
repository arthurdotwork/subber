package cmd

import (
	"context"
	"errors"
	"fmt"

	"github.com/arthureichelberger/subber/model"
	"github.com/arthureichelberger/subber/pkg/pubsub"
	"github.com/arthureichelberger/subber/service"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var maxMessages uint
var interactively bool
var readSubName string

// readSubCmd represents the readSub command
var readSubCmd = &cobra.Command{
	Use:   "readSub",
	Short: "readSub allows to read messages in a subscription.",
	Run: func(cmd *cobra.Command, args []string) {
		var subName string
		var err error

		if readSubName == "" {
			subName, err = service.NewPrompt("Please enter a subscriptionName", func(value string) error {
				if len(value) == 0 {
					return errors.New("sub name cannot be null")
				}

				return nil
			})

			if err != nil {
				pterm.Error.Println(err.Error())
				return
			}
		} else {
			subName = readSubName
		}

		ctx := context.Background()
		client, err := pubsub.NewPubsubClient(ctx, fmt.Sprintf("%v", viper.Get("PUBSUB_PROJECT_ID")), fmt.Sprintf("%v", viper.Get("EMULATOR_HOST")))
		if err != nil {
			pterm.Error.Println(err.Error())
			return
		}

		pubSubService := service.NewPubSubService(client)
		if !interactively {
			read(ctx, pubSubService, subName)
		} else {
			readInteractively(ctx, pubSubService, subName)
		}

	},
}

func init() {
	rootCmd.AddCommand(readSubCmd)

	readSubCmd.Flags().UintVar(&maxMessages, "maxMessages", 10, "Number of messages before stopping reception")
	readSubCmd.Flags().BoolVar(&interactively, "interactively", false, "Whether or not you want to ack messages interactively")
	readSubCmd.Flags().StringVar(&readSubName, "sub", "", "The name of the sub from which you want to read messages")
}

func read(ctx context.Context, pubSubService service.PubSubServiceInterface, subName string) {
	channel := make(chan model.Message)

	go func() {
		if err := pubSubService.ReadSub(ctx, subName, channel, maxMessages); err != nil {
			pterm.Error.Printfln("Cannot read from subscription %s. (%s)", subName, err.Error())
			return
		}
	}()

	for {
		msg := <-channel
		pterm.Success.Printfln("Received message : %s. (%d/%d)", string(msg.Message), msg.Id, maxMessages)

		if msg.Id == maxMessages {
			close(channel)
			return
		}
	}
}

func readInteractively(ctx context.Context, pubSubService service.PubSubServiceInterface, subName string) {
	channel := make(chan model.Message)
	ackChan := make(chan bool)

	defer close(channel)
	defer close(ackChan)

	go func() {
		if err := pubSubService.ReadSubInteractive(ctx, subName, channel, ackChan, maxMessages); err != nil {
			pterm.Error.Printfln("Cannot read from subscription %s. (%s)", subName, err.Error())
			return
		}
	}()

	for {
		msg := <-channel
		pterm.Success.Printfln("Received message : %s. (%d/%d)", string(msg.Message), msg.Id, maxMessages)
		shouldAck, err := service.Confirm("Would you like to ack this message")
		if err != nil {
			pterm.Error.Println(err.Error())
			return
		}
		ackChan <- shouldAck
		acked := <-ackChan
		if acked {
			pterm.Info.Println("Message have been acked successfully.")
		} else {
			pterm.Info.Println("Message have been nacked successfully.")
		}

		if msg.Id == maxMessages {
			close(channel)
			return
		}
	}
}
