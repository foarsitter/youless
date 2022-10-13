/*
Package cmd
Copyright © 2022 Jelmer Draaijer info@jelmert.nl
*/
package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"youless-example/prometheus"
)

// pushgatewayCmd represents the pushgateway command
var pushgatewayCmd = &cobra.Command{
	Use:   "pushgateway",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		pushgatewayUrl := viper.GetString("pushgateway.url")
		job := viper.GetString("pushgateway.job")
		groupingName := viper.GetString("pushgateway.group.name")
		groupingValue := viper.GetString("pushgateway.group.value")
		pushgatewayUser := viper.GetString("pushgateway.auth.user")
		pushgatewayPassword := viper.GetString("pushgateway.auth.password")

		serverIP := viper.GetString("youless.ip")

		prometheus.UpdateValues(serverIP, pushgatewayUrl, pushgatewayUser, pushgatewayPassword, job, groupingName, groupingValue)

		if !DoNotRunInBackground {
			prometheus.Background(serverIP, pushgatewayUrl, pushgatewayUser, pushgatewayPassword, job, groupingName, groupingValue)
		}

	},
}

var (
	PushgatewayHost      string
	PushgatewayUser      string
	PushgatewayPassword  string
	DoNotRunInBackground bool
)

func init() {
	rootCmd.AddCommand(pushgatewayCmd)

	pushgatewayCmd.Flags().StringVar(&PushgatewayHost, "host", "", "The host url of the Prometheus Pushgateway")
	pushgatewayCmd.Flags().StringVar(&PushgatewayUser, "user", "", "The user of the Prometheus Pushgateway")
	pushgatewayCmd.Flags().StringVar(&PushgatewayPassword, "password", "", "The password of the Prometheus Pushgateway")
	pushgatewayCmd.Flags().BoolVar(&DoNotRunInBackground, "single", false, "The password of the Prometheus Pushgateway")

	_ = viper.BindPFlag("pushgateway.host", pushgatewayCmd.Flags().Lookup("host"))
	_ = viper.BindPFlag("pushgateway.auth.user", pushgatewayCmd.Flags().Lookup("user"))
	_ = viper.BindPFlag("pushgateway.auth.password", pushgatewayCmd.Flags().Lookup("password"))
}
