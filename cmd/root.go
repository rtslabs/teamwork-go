package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/rtslabs/teamwork-go/configuration"
)

var cfgFile string

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "teamworkgo",
	Short: "Go CLI to interact with the teamwork API",
	Long:  ``,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {

	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(func() { configuration.InitConfig(cfgFile) }, ) // doesn't work?
	configuration.InitConfig("")

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/"+configuration.FILENAME+".yaml/json)")
	RootCmd.PersistentFlags().StringP("companyName", "", "", "Name of Company")
	RootCmd.PersistentFlags().StringP("apiKey", "", "", "API Key to access Teamwork")
	RootCmd.PersistentFlags().StringP("userId", "", "User ID", "ID of the user in Teamwork")

	viper.BindPFlag("global.companyName", RootCmd.PersistentFlags().Lookup("companyName"))
	viper.BindPFlag("global.apiKey", RootCmd.PersistentFlags().Lookup("apiKey"))
	viper.BindPFlag("global.userId", RootCmd.PersistentFlags().Lookup("userId"))

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// RootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}

// initConfig reads in config file and ENV variables if set.
