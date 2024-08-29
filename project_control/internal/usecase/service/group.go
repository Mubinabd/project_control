package usecase

import (
	"context"

	"github.com/Mubinabd/project_control/internal/repository"
	pb "github.com/Mubinabd/project_control/pkg/genproto"
)

type GroupService struct {
	stg repository.StorageI
	pb.UnimplementedGroupServiceServer
}

func NewGroupService(stg repository.StorageI) *GroupService {
	return &GroupService{stg: stg}
}
func (s *GroupService) CreateGroup(c context.Context, req *pb.CreateGroupReq) (*pb.Void, error) {
	return s.stg.Group().CreateGroup(req)
}
func (s *GroupService) GetGroup(c context.Context, req *pb.ById) (*pb.GroupGet, error) {
	return s.stg.Group().GetGroup(req)
}
func (s *GroupService) ListGroups(c context.Context, req *pb.GroupListReq) (*pb.GroupListRes, error) {
	return s.stg.Group().ListGroups(req)
}
func (s *GroupService) UpdateGroup(c context.Context, req *pb.UpdateGr) (*pb.Void, error) {
	return s.stg.Group().UpdateGroup(req)
}
func (s *GroupService) DeleteGroup(c context.Context, req *pb.DeleteGr) (*pb.Void, error) {
	return s.stg.Group().DeleteGroup(req)
}
