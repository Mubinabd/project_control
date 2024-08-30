package handlers

import (
	s "github.com/Mubinabd/project_control/internal/usecase/service"
	kafka "github.com/Mubinabd/project_control/pkg/kafka/producer"
	"github.com/go-redis/redis/v8"
)

type Handlers struct {
	Group    *s.GroupService
	Private  *s.PrivateService
	Auth     *s.AuthService
	User     *s.UserService
	RDB      *redis.Client
	Producer kafka.KafkaProducer
}

func NewHandler(group *s.GroupService, private *s.PrivateService, auth *s.AuthService, user *s.UserService, rdb *redis.Client, pr *kafka.KafkaProducer) *Handlers {
	return &Handlers{
		Group:    group,
		Private:  private,
		Auth:     auth,
		User:     user,
		RDB:      rdb,
		Producer: *pr,
	}
}
