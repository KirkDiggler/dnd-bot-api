package cmd

import (
	"context"
	dndbotv1alpha1api "github.com/KirkDiggler/dnd-bot-api/protos/gen-external/go/api/admin"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"log/slog"
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

		client := dndbotv1alpha1api.NewAminAPIClient(conn)

		result, err := client.ListRooms(ctx, &dndbotv1alpha1api.ListRoomsRequest{
			Name: "bob",
		})
		if err != nil {
			panic(err)

			return
		}

		slog.Info("result returned successfully: ", result)
	},
}

func init() {
	rootCmd.AddCommand(testCommand)
}
