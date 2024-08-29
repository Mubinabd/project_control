package kafka

import (
	"context"
	"encoding/json"

	pb "github.com/Mubinabd/project_control/internal/genproto/auth"
	"github.com/Mubinabd/project_control/service"
	"golang.org/x/exp/slog"
)

func UserRegisterHandler(u *service.AuthService) func(message []byte) {
	return func(message []byte) {
		var user pb.RegisterReq
		if err := json.Unmarshal(message, &user); err != nil {
			slog.Error("Cannot unmarshal JSON: %v", err)
			return
		}

		_, err := u.Register(context.Background(), &user)
		if err != nil {
			slog.Error("failed to register user via Kafka: %v", err)
			return
		}
		slog.Info("User registered")
	}
}
