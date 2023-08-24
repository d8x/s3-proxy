package main

import (
	"fmt"
	"net/http"

	"github.com/d8x/sgw/providers"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	version        string
	configFilePath string
)

var rootCmd = &cobra.Command{
	Use:   "sgw",
	Short: "sgw is a gateway to multiple storage providers",
}

var versionCmd = &cobra.Command{
	Use:  "version",
	Long: "print the version number of sgw",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("sgw version %s\n", version)
	},
}

func main() {
	rootCmd.AddCommand(versionCmd)
	rootCmd.Flags().BoolP("debug", "d", false, "enable debug mode")
	if err := viper.BindPFlag("debug", rootCmd.Flags().Lookup("debug")); err != nil {
		panic(err)
	}
	rootCmd.Flags().StringVarP(&configFilePath, "config", "c", ".", "config file path")

	cfg := NewConfig()
	if err := cfg.ReadConfig(configFilePath); err != nil {
		panic(err)
	}

	if cfg.Debug {
		logrus.SetLevel(logrus.DebugLevel)
	}

	router := http.NewServeMux()
	NewAPI(cfg).Create(router)

	logrus.Infof("listening on port %s", cfg.ListenPort)
	logrus.Infof("supported providers: %v", providers.GetSupportedProviders())
	logrus.Fatal(http.ListenAndServe(cfg.ListenPort, router))

}
