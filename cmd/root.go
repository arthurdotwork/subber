package cmd

import (
	"errors"
	"os"

	"github.com/arthureichelberger/subber/service"
	"github.com/spf13/cobra"

	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "subber",
	Short: "Subber is a CLI tool to interact with a pubsub local emulator",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	home, err := os.UserHomeDir()
	cobra.CheckErr(err)
	checkConfigFile(home + "/.subber.yaml")
	viper.AddConfigPath(home)
	viper.SetConfigType("yaml")
	viper.SetConfigName(".subber")

	viper.AutomaticEnv()
	cobra.CheckErr(viper.ReadInConfig())

	checkConfig()
}

func checkConfigFile(cfgFile string) {
	if _, err := os.Stat(cfgFile); os.IsNotExist(err) {
		_, err = os.Create(cfgFile)
		cobra.CheckErr(err)
	}
}

func checkConfig() {
	if viper.Get("PUBSUB_PROJECT_ID") == nil {
		pubsubProjectId, err := service.NewPrompt("Pubsub Project ID", func(value string) error {
			if len(value) == 0 {
				return errors.New("pubsub Project ID cannot be null")
			}

			return nil
		})

		if err != nil {
			panic(err.Error())
		}

		viper.Set("PUBSUB_PROJECT_ID", pubsubProjectId)
	}

	if viper.Get("EMULATOR_HOST") == nil {
		pubsubProjectId, err := service.NewPrompt("Emulator Host", func(value string) error {
			if len(value) == 0 {
				return errors.New("emulator Host cannot be null")
			}

			return nil
		})

		if err != nil {
			panic(err.Error())
		}

		viper.Set("EMULATOR_HOST", pubsubProjectId)
	}

	_ = viper.WriteConfig()
}
