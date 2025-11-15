package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "fclean",
	Short: "🧹 Flutter Cleaner — simple and smart Flutter project cleanup",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Use one of the commands below or run 'fclean wizard' for interactive mode.")
			cmd.Help()
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}


func init() {
	rootCmd.AddCommand(scanCmd)
	rootCmd.AddCommand(cleanCmd)
	rootCmd.AddCommand(suggestCmd)
	rootCmd.AddCommand(statsCmd)
}