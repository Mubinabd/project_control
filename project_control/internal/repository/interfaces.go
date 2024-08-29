package repository

import (
	pb "github.com/Mubinabd/project_control/pkg/genproto"
	auth "github.com/Mubinabd/project_control/pkg/genproto/auth"
)

type StorageI interface {
	Group() GroupI
	Private() PrivateI
	Auth() AuthI
	User() UserI
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

type AuthI interface {
	Register(*auth.RegisterReq) (*auth.Void, error)
    Login(*auth.LoginReq) (*auth.User, error)
    ForgotPassword(*auth.GetByEmail) (*auth.Void, error)
    ResetPassword(*auth.ResetPassReq) (*auth.Void, error)
    SaveRefreshToken(*auth.RefToken) (*auth.Void, error)
    GetAllUsers(*auth.ListUserReq) (*auth.ListUserRes, error)
}

type UserI interface {
	GetProfile(*auth.GetById) (*auth.UserRes, error)
    EditProfile(*auth.UserRes) (*auth.UserRes, error)
    ChangePassword(*auth.ChangePasswordReq) (*auth.Void, error)
    GetSetting(*auth.GetById) (*auth.Setting, error)
    EditSetting(*auth.SettingReq) (*auth.Void, error)
    DeleteUser(*auth.GetById) (*auth.Void, error)
}