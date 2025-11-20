package cmd

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"

	grpcAdapter "github.com/mukappalambda/my-trader/internal/adapters/grpc"
	restAdapter "github.com/mukappalambda/my-trader/internal/adapters/rest"
)

var ginMode = flag.String("gin-mode", "release", "gin mode (debug|release)")

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "",
	RunE: func(cmd *cobra.Command, args []string) error {
		flag.Parse()
		dsn := os.Getenv("DATABASE_URL")
		srv, err := grpcAdapter.NewGrpcServer(dsn)
		if err != nil {
			return fmt.Errorf("failed to create grpc server: %s", err)
		}
		schemaRegistryServer, err := restAdapter.NewSchemaRegistryServer(*ginMode)
		if err != nil {
			return fmt.Errorf("failed to create schema registry server: %s", err)
		}
		registryPort, err := cmd.Flags().GetInt("schema-registry-server-port")
		if err != nil {
			return fmt.Errorf("failed to get schema registry port: %s", err)
		}

		go func() {
			log.Printf("schema registry server is listening at port: %d\n", registryPort)
			if err := schemaRegistryServer.Run(fmt.Sprintf("localhost:%d", registryPort)); err != nil {
				fmt.Fprintf(os.Stderr, "failed to run schema registry server: %s\n", err)
			}
		}()
		port, err := cmd.Flags().GetInt("grpc-server-port")
		if err != nil {
			return fmt.Errorf("failed to get grpc-server-port: %s", err)
		}
		if err := srv.Run(port); err != nil {
			return err
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(runCmd)

	runCmd.Flags().Int("grpc-server-port", 50051, "gRPC server port")
	runCmd.Flags().Int("schema-registry-server-port", 8081, "schema registry server port")
}
