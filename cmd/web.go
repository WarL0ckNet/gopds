package cmd

import (
	"gopds/server"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	port   int
	host   string
	webCmd = &cobra.Command{
		Use:   "web",
		Short: "Start web server",
		Run: func(cmd *cobra.Command, args []string) {
			server.Start(viper.GetString("server.host"), viper.GetInt("server.port"))
		},
	}
)

func init() {
	rootCmd.AddCommand(webCmd)
	webCmd.Flags().StringVarP(&host, "host", "H", "", "host address")
	webCmd.Flags().IntVarP(&port, "port", "p", 8088, "port")
	viper.BindPFlag("server.host", webCmd.Flags().Lookup("host"))
	viper.BindPFlag("server.port", webCmd.Flags().Lookup("port"))
}
