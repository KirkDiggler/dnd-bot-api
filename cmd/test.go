package cmd

import (
	"context"
	"log/slog"

	v1alpha1 "github.com/KirkDiggler/dnd-bot-api/protos/gen/go/api/v1alpha1"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

var testCommand = &cobra.Command{
	Use:   "test",
	Short: "Run the test command",
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()

		conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
		if err != nil {
			panic(err)
		}

		defer conn.Close()

		client := v1alpha1.NewPlayerAPIClient(conn)

		result, err := client.ListClasses(ctx, &v1alpha1.ListClassesRequest{})
		if err != nil {
			panic(err)

			return
		}

		slog.Info("list classes returned successfully: ", "result", result)

		getResult, err := client.GetClass(ctx, &v1alpha1.GetClassRequest{
			Id: result.Classes[0].Id,
		})
		if err != nil {
			panic(err)

			return
		}

		slog.Info("get class returned successfully: ", "result", getResult)
	},
}

func init() {
	rootCmd.AddCommand(testCommand)
}
