package cmd

import (
	"context"
	"dysn/character/app"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "Init dysn/characters microservice",
	Short: "Dysn characters microservice ",
	Long: `Microservice contains fuctions for working with dysn characters`,
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		ctx, cancel := context.WithCancel(ctx)
		defer cancel()

		go func() {
			sgn := make(chan os.Signal, 1)
			signal.Notify(sgn, syscall.SIGINT, syscall.SIGTERM)

			select {
			case <-ctx.Done():
			case <-sgn:
			}
			cancel()
		}()

		app.Run(ctx)
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}