package cmd

import (
	"fmt"
	"os"

	"github.com/mcupdater/mcu-go/pkg/util"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "mcu-go",
	Short: "MCU-GO is a Go implementation of MCUpdater",
	Long:  `MCU-GO is a Go implementation of MCUpdater`,
	Run: func(cmd *cobra.Command, args []string) {
		// dummy operation
		if jrePath := viper.GetString("jrePath"); jrePath == "" {
			fmt.Printf("Unable to read `jrePath` from config.\n")
		} else if util.FileExists(jrePath) {
			fmt.Printf("Found `jrePath` at %s.\n", jrePath)
		} else {
			fmt.Printf("Unable to find `jrePath` at %s.\n", jrePath)
		}
	},
}

func init() {
	cobra.OnInitialize(util.InitConfig)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
