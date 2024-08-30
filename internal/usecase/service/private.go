package service

import (
	"context"

	"github.com/Mubinabd/project_control/internal/repository"
	pb "github.com/Mubinabd/project_control/pkg/genproto"
)

type PrivateService struct {
	stg repository.StorageI
	pb.UnimplementedPrivateServiceServer
}

func NewPrivateService(stg repository.StorageI) *PrivateService {
	return &PrivateService{stg: stg}
}
func (s *PrivateService) CreatePrivate(c context.Context, req *pb.CreatePrivateReq) (*pb.Void, error) {
	return s.stg.Private().CreatePrivate(req)
}
func (s *PrivateService) GetPrivate(c context.Context, req *pb.ById) (*pb.PrivateGet, error) {
	return s.stg.Private().GetPrivate(req)
}
func (s *PrivateService) ListPrivates(c context.Context, req *pb.PrivateListReq) (*pb.PrivateListRes, error) {
	return s.stg.Private().ListPrivates(req)
}
func (s *PrivateService) UpdatePrivate(c context.Context, req *pb.UpdatePrivat) (*pb.Void, error) {
	return s.stg.Private().UpdatePrivate(req)
}
func (s *PrivateService) DeletePrivate(c context.Context, req *pb.DeletePrivat) (*pb.Void, error) {
	return s.stg.Private().DeletePrivate(req)
}
