/*
Copyright © 2022 Jelmer Draaijer info@jelmert.nl
*/
package cmd

import (
	"github.com/spf13/cobra"
	jww "github.com/spf13/jwalterweatherman"
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

		jww.INFO.Print(viper.GetString("pushgateway.auth.password"))
		jww.INFO.Print(viper.GetString("pushgateway.auth.user"))
		serverIP := viper.GetString("youless.ip")
		jww.INFO.Print(serverIP)

		prometheus.Background(serverIP)
	},
}

var (
	PushgatewayHost     string
	PushgatewayUser     string
	PushgatewayPassword string
)

func init() {
	rootCmd.AddCommand(pushgatewayCmd)

	pushgatewayCmd.Flags().StringVar(&PushgatewayHost, "host", "", "The host url of the Prometheus Pushgateway")
	pushgatewayCmd.Flags().StringVar(&PushgatewayUser, "user", "", "The user of the Prometheus Pushgateway")
	pushgatewayCmd.Flags().StringVar(&PushgatewayPassword, "password", "", "The password of the Prometheus Pushgateway")

	viper.BindPFlag("pushgateway.host", pushgatewayCmd.Flags().Lookup("host"))
	viper.BindPFlag("pushgateway.auth.user", pushgatewayCmd.Flags().Lookup("user"))
	viper.BindPFlag("pushgateway.auth.password", pushgatewayCmd.Flags().Lookup("password"))
}
