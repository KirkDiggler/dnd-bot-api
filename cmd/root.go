package cmd

import (
	"context"
	"github.com/spf13/cobra"
	"log"
)

var rootCmd = &cobra.Command{
	Use: "dnd-bot-api",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(context.Background(), "error in rootCommand.Execute", err)
	}
}
