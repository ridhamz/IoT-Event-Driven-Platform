package events

import (
	"encoding/json"
	"go-cqrs-api/infrastructure"
)

func PublishUserCreated(id, name string) error {
	msg := map[string]string{"id": id, "name": name}
	data, _ := json.Marshal(msg)
	return infrastructure.RedisClient().Publish(infrastructure.Ctx(), "user_created", data).Err()
}
