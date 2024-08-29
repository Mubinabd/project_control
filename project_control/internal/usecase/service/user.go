package service

import (
	"context"

	pb "github.com/Mubinabd/project_control/pkg/genproto/auth"
	st "github.com/Mubinabd/project_control/internal/repository/postgresql"
)

type UserService struct {
	storage st.Storage
	pb.UnimplementedUserServiceServer
}

func NewUserService(storage *st.Storage) *UserService {
	return &UserService{
		storage: *storage,
	}
}

func (s *UserService) GetProfile(ctx context.Context, req *pb.GetById) (*pb.UserRes, error) {
	res, err := s.storage.UserS.GetProfile(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *UserService) EditProfile(ctx context.Context, req *pb.UserRes) (*pb.UserRes, error) {
	res, err := s.storage.UserS.EditProfile(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *UserService) ChangePassword(ctx context.Context, req *pb.ChangePasswordReq) (*pb.Void, error) {
	res, err := s.storage.UserS.ChangePassword(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *UserService) GetSetting(ctx context.Context, req *pb.GetById) (*pb.Setting, error) {
	res, err := s.storage.UserS.GetSetting(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *UserService) EditSetting(ctx context.Context, req *pb.SettingReq) (*pb.Void, error) {
	res, err := s.storage.UserS.EditSetting(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *UserService) DeleteUser(ctx context.Context, req *pb.GetById) (*pb.Void, error) {
	res, err := s.storage.UserS.DeleteUser(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}
