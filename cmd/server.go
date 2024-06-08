package cmd

import (
	"context"
	"log"
	"log/slog"
	"net"
	"net/http"
	"os"

	"github.com/KirkDiggler/dnd-bot-api/internal/clients/dnd5e"

	"github.com/KirkDiggler/dnd-bot-api/internal/handlers/v1alpha1"
	dndbotv1alpha1api "github.com/KirkDiggler/dnd-bot-api/protos/gen/go/api/v1alpha1"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
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

		dnd5eClient, err := dnd5e.New(&dnd5e.Config{
			HttpClient: http.DefaultClient,
		})
		if err != nil {
			panic(err)
		}

		hander, err := v1alpha1.NewHandler(&v1alpha1.HandlerConfig{
			Client: dnd5eClient,
		})
		if err != nil {
			log.Fatal(ctx, "error in v1alpha1.NewHandler", err)
		}

		dndbotv1alpha1api.RegisterPlayerAPIServer(srv, hander)
		err = srv.Serve(lis)
		if err != nil {
			log.Fatal(ctx, "error in srv.Serve", err)
		}

	},
}

func init() {
	rootCmd.AddCommand(serverCommand)
}
