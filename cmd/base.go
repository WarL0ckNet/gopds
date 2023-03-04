package cmd

import (
	"gopds/db"
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cnt     db.DB_connect
	migrCmd = &cobra.Command{
		Use:   "migrate",
		Short: "Init or update database",
		Run: func(cmd *cobra.Command, args []string) {
			base, err := db.ConnectDatabase(&db.DB_connect{
				Host: viper.GetString("db.host"),
				Port: viper.GetInt("db.port"),
				Name: viper.GetString("db.name"),
				User: viper.GetString("db.user"),
				Pass: viper.GetString("db.pass"),
			})
			if err != nil {
				log.Fatalf("Connect database error: %v", err)
			} else {
				db.InitDatabase(base)
			}
		},
	}
	clearCmd = &cobra.Command{
		Use:   "clear",
		Short: "Clear database",
		Run: func(cmd *cobra.Command, args []string) {
			base, err := db.ConnectDatabase(&db.DB_connect{
				Host: viper.GetString("db.host"),
				Port: viper.GetInt("db.port"),
				Name: viper.GetString("db.name"),
				User: viper.GetString("db.user"),
				Pass: viper.GetString("db.pass"),
			})
			if err != nil {
				log.Fatalf("Connect database error: %v", err)
			} else {
				db.ClearDatabase(base)
			}
		},
	}
)

func init() {
	rootCmd.AddCommand(migrCmd)
	rootCmd.AddCommand(clearCmd)
	migrCmd.Flags().StringVarP(&cnt.Host, "db_host", "H", "localhost", "database host")
	migrCmd.Flags().StringVarP(&cnt.Name, "db_name", "d", "gopds", "database name")
	migrCmd.Flags().StringVarP(&cnt.User, "db_user", "u", "gopds", "database username")
	migrCmd.Flags().StringVarP(&cnt.Pass, "db_pass", "P", "", "database password")
	migrCmd.Flags().IntVarP(&cnt.Port, "db_port", "p", 5432, "database port")
	clearCmd.Flags().AddFlagSet(migrCmd.Flags())
	viper.BindPFlag("db.host", migrCmd.Flags().Lookup("db_host"))
	viper.BindPFlag("db.name", migrCmd.Flags().Lookup("db_name"))
	viper.BindPFlag("db.user", migrCmd.Flags().Lookup("db_user"))
	viper.BindPFlag("db.pass", migrCmd.Flags().Lookup("db_pass"))
	viper.BindPFlag("db.port", migrCmd.Flags().Lookup("db_port"))
}
