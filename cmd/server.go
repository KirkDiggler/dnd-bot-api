package cmd

import (
	"context"
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
		
		err = srv.Serve(lis)
		if err != nil {
			log.Fatal(ctx, "error in srv.Serve", err)
		}

	},
}

func init() {
	rootCmd.AddCommand(serverCommand)
}
