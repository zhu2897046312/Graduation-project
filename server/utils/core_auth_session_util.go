package utils

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"golang.org/x/net/context"
)

// CoreLoginUserInfoModel 定义用户信息结构体
type CoreLoginUserInfoModel struct {
	ID         int64  `json:"id"`
	Nickname   string `json:"nickname"`
	Account    string `json:"account"`
	DeptID     int64 `json:"dept_id"`
	Avatar     string `json:"avatar"`
	Permission string `json:"permission"`
}

// RedisPrefixKey Redis 键前缀
const RedisPrefixKey = "yex:auth:"

// SessionID 生成会话ID
func SessionID(id int64) string {
	now := time.Now().String()
	hashID := md5.Sum([]byte(fmt.Sprintf("%d", id)))
	hashNow := md5.Sum([]byte(now))
	randomUUID := strings.ReplaceAll(uuid.New().String(), "-", "")[:5]

	return fmt.Sprintf("%s.%s.%s",
		hex.EncodeToString(hashID[:]),
		hex.EncodeToString(hashNow[:]),
		randomUUID)
}

// PutSession 将用户信息存入Redis
func PutSession(rdb *redis.Client, id int64, obj *CoreLoginUserInfoModel) (string, error) {
	key := SessionID(id)
	jsonData, err := json.Marshal(obj)
	if err != nil {
		return "", err
	}

	err = rdb.Set(context.Background(),
		fmt.Sprintf("%s:%s", RedisPrefixKey, key),
		jsonData,
		8*time.Hour).Err()
	if err != nil {
		return "", err
	}

	return key, nil
}

// GetSession 从Redis获取用户信息
func GetSession(rdb *redis.Client, key string, defaultVal *CoreLoginUserInfoModel) (*CoreLoginUserInfoModel, error) {
	cache, err := rdb.Get(context.Background(), fmt.Sprintf("%s:%s", RedisPrefixKey, key)).Result()
	if err == redis.Nil {
		return defaultVal, nil
	} else if err != nil {
		return nil, err
	}

	var user CoreLoginUserInfoModel
	err = json.Unmarshal([]byte(cache), &user)
	if err != nil {
		return defaultVal, err
	}

	return &user, nil
}

// GetIDFromSession 从会话中获取用户ID
func GetIDFromSession(rdb *redis.Client, key string, defaultVal int64) (int64, error) {
	user, err := GetSession(rdb, key, nil)
	if err != nil {
		return defaultVal, err
	}
	if user == nil {
		return defaultVal, nil
	}
	return user.ID, nil
}
