package app

import (
	"log"
	"net"

	"github.com/Mubinabd/project_control/internal/pkg/config"
	pb "github.com/Mubinabd/project_control/internal/pkg/genproto"
	ps "github.com/Mubinabd/project_control/internal/pkg/postgres"
	"github.com/Mubinabd/project_control/internal/repository/postgresql"
	service "github.com/Mubinabd/project_control/internal/usecase/service"

	"google.golang.org/grpc"
)

func Run(cfg config.Config) {

	// connect to postgres
	db, err := ps.New(&cfg)
	if err != nil {
		log.Fatal(err)
	}
	stg := postgresql.NewStorage(db.DB)

	group_service := service.NewGroupService(stg)
	private_service := service.NewPrivateService(stg)

	lis, err := net.Listen("tcp", ":20020")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	server := grpc.NewServer()

	pb.RegisterGroupServiceServer(server, group_service)
	pb.RegisterPrivateServiceServer(server, private_service)
	log.Println("Starting gRPC server on port", cfg.GRPCPort)
	if err := server.Serve(lis); err != nil {
		log.Fatal("gRPC server failed to start: ", err)
	}

}
