package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "i18n",
	Short: "A demo CLI for i18n",
	Long:  `A longer description that spans multiple lines and likely contains examples and usage of using your application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello from the i18n CLI!")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
