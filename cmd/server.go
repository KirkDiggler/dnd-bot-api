package cmd

import (
	"context"
	"github.com/KirkDiggler/dnd-bot-api/internal/handlers/v1alpha1"
	"github.com/KirkDiggler/dnd-bot-api/internal/repositories/rooms"
	dndbotv1alpha1api "github.com/KirkDiggler/dnd-bot-api/protos/gen-external/go/api/admin"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"log"
	"log/slog"
	"net"
	"os"
)

var serverCommand = &cobra.Command{
	Use:   "server",
	Short: "Run the gRPC server",
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

		lis, err := net.Listen("tcp", ":8080")
		if err != nil {
			log.Fatal(ctx, "error in net.Listen", err)
		}

		srv := grpc.NewServer()
		logger.Info("server listening")

		redisClient := redis.NewUniversalClient(&redis.UniversalOptions{
			Addrs: []string{"localhost:6379"},
		})

		roomRepo, err := rooms.NewRedis(&rooms.RedisConfig{
			Client: redisClient,
		})
		if err != nil {
			log.Fatal(ctx, "error in rooms.NewRedis", err)
		}

		hander, err := v1alpha1.NewHandler(&v1alpha1.HandlerConfig{
			RoomRepo: roomRepo,
		})
		if err != nil {
			log.Fatal(ctx, "error in v1alpha1.NewHandler", err)
		}

		dndbotv1alpha1api.RegisterAminAPIServer(srv, hander)
		err = srv.Serve(lis)
		if err != nil {
			log.Fatal(ctx, "error in srv.Serve", err)
		}

	},
}

func init() {
	rootCmd.AddCommand(serverCommand)
}
