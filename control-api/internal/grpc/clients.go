package grpc

import (
    "github.com/Mubinabd/project_control/internal/pkg/config"
    pb "github.com/Mubinabd/project_control/internal/pkg/genproto"
    "google.golang.org/grpc"
    "google.golang.org/grpc/credentials/insecure"
)

type Clients struct {
    Group   pb.GroupServiceClient
    Private pb.PrivateServiceClient
}

func NewClients(cfg *config.Config) (*Clients, error) {
    group_conn, err := grpc.Dial(cfg.ControlURL, grpc.WithTransportCredentials(insecure.NewCredentials()))
    if err != nil {
        return nil, err
    }

    groupClient := pb.NewGroupServiceClient(group_conn)
    privateClient := pb.NewPrivateServiceClient(group_conn)

    return &Clients{
        Group:   groupClient,
        Private: privateClient,
    }, nil
}
