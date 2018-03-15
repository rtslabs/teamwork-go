package cmd

import (
	"fmt"
	"os"

	"github.com/rtslabs/teamwork-go/configuration"
	"github.com/spf13/cobra"
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
	cobra.OnInitialize(func() { configuration.InitConfig(cfgFile) }) // doesn't work?
	configuration.InitConfig("")
}

// initConfig reads in config file and ENV variables if set.
