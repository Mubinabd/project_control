package kafka

import (
	"context"
	"encoding/json"

	"github.com/Mubinabd/project_control/internal/usecase/service"
	pb "github.com/Mubinabd/project_control/pkg/genproto/auth"
	"golang.org/x/exp/slog"
)

func UserEditProfileHandler(u *service.UserService) func(message []byte) {
	return func(message []byte) {
		var user pb.UserRes
		if err := json.Unmarshal(message, &user); err != nil {
			slog.Error("Cannot unmarshal JSON: %v", err)
			return
		}

		_, err := u.EditProfile(context.Background(), &user)
		if err != nil {
			slog.Error("failed to edit user via Kafka: %v", err)
			return
		}
		slog.Info("User updated")
	}
}

func UserEditPasswordHandler(u *service.UserService) func(message []byte) {
	return func(message []byte) {
		var user pb.ChangePasswordReq
		if err := json.Unmarshal(message, &user); err != nil {
			slog.Error("Cannot unmarshal JSON: %v", err)
			return
		}

		_, err := u.ChangePassword(context.Background(), &user)
		if err != nil {
			slog.Error("failed to edit password via Kafka: %v", err)
			return
		}
		slog.Info("Password updated")
	}
}

func UserEditSettingHandler(u *service.UserService) func(message []byte) {
	return func(message []byte) {
		var user pb.SettingReq
		if err := json.Unmarshal(message, &user); err != nil {
			slog.Error("Cannot unmarshal JSON: %v", err)
			return
		}

		_, err := u.EditSetting(context.Background(), &user)
		if err != nil {
			slog.Error("failed to edit user settings via Kafka: %v", err)
			return
		}
		slog.Info("User settings updated")
	}
}
