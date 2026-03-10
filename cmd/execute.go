package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ygo-cli",
	Short: "A CLI tool for looking up Yu-Gi-Oh card data",
	Long:  "ygo-cli is a terminal tool for looking up Yu-Gi-Oh card stats, types, attributes and more!",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
