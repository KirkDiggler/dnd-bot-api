package cmd

import (
	"context"
	"log/slog"

	dndbotv1alpha1api "github.com/KirkDiggler/dnd-bot-api/protos/gen-external/go/api/admin"
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

		client := dndbotv1alpha1api.NewAminAPIClient(conn)

		result, err := client.CreateRoom(ctx, &dndbotv1alpha1api.CreateRoomRequest{
			Room: &dndbotv1alpha1api.Room{
				Id:          "1",
				Name:        "bob",
				Description: "bob likes turtles",
			},
		})

		if err != nil {
			panic(err)

			return
		}

		slog.Info("create returned successfully: ", result)

		getResult, err := client.GetRoom(ctx, &dndbotv1alpha1api.GetRoomRequest{
			Id: "1",
		})
		if err != nil {
			panic(err)

			return
		}

		slog.Info("get returned successfully", getResult)

	},
}

func init() {
	rootCmd.AddCommand(testCommand)
}
