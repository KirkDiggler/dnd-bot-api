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
				Name:        "Google",
				Description: "Googlr like uuids",
			},
		})
		if err != nil {
			panic(err)

			return
		}

		slog.Info("create returned successfully: ", "result", result)

		listResult, err := client.ListRooms(ctx, &dndbotv1alpha1api.ListRoomsRequest{})
		if err != nil {
			panic(err)

			return
		}

		slog.Info("list returned successfully", "listResult", listResult)

		if len(listResult.Rooms) > 0 {
			getResult, err := client.GetRoom(ctx, &dndbotv1alpha1api.GetRoomRequest{
				Id: listResult.Rooms[0].Id,
			})
			if err != nil {
				panic(err)
			}

			slog.Info("get returned successfully", "getResult", getResult)
		}
	},
}

func init() {
	rootCmd.AddCommand(testCommand)
}
