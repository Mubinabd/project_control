package kafka

import (
	"context"
	"encoding/json"

	s "github.com/Mubinabd/project_control/internal/usecase/service"
	pb "github.com/Mubinabd/project_control/pkg/genproto/auth"
	"golang.org/x/exp/slog"
)

func UserRegisterHandler(u *s.AuthService) func(message []byte) {
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
