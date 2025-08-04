package queries

import (
	"encoding/json"
	"fmt"
	"go-cqrs-api/infrastructure"
)

func GetUserFromReadModel(id string) (map[string]string, error) {
	val, err := infrastructure.RedisClient().Get(infrastructure.Ctx(), fmt.Sprintf("user:%s", id)).Result()
	if err != nil {
		return nil, err
	}
	var user map[string]string
	json.Unmarshal([]byte(val), &user)
	return user, nil
}
