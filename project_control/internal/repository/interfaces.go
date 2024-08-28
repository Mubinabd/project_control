package repository

import (
	pb "github.com/Mubinabd/project_control/internal/pkg/genproto"
)

type StorageI interface {
	Group() GroupI
	Private() PrivateI
}

type GroupI interface {
	CreateGroup(req *pb.CreateGroupReq) (*pb.Void, error)
	GetGroup(req *pb.ById) (*pb.GroupGet, error)
	UpdateGroup(req *pb.UpdateGr) (*pb.Void, error)
	DeleteGroup(req *pb.DeleteGr) (*pb.Void, error)
	ListGroups(req *pb.GroupListReq) (*pb.GroupListRes, error)
}

type PrivateI interface {
	CreatePrivate(req *pb.CreatePrivateReq) (*pb.Void, error)
	UpdatePrivate(req *pb.UpdatePrivat) (*pb.Void, error)
	DeletePrivate(req *pb.DeletePrivat) (*pb.Void, error)
	GetPrivate(req *pb.ById) (*pb.PrivateGet, error)
	ListPrivates(req *pb.PrivateListReq) (*pb.PrivateListRes, error)
}
