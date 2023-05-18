package auth

import (
	"errors"
	"log"
	"time"

	"github.com/dev-hyunsang/daily-todo/config"
	"github.com/dev-hyunsang/daily-todo/models"
	"github.com/go-redis/redis/v7"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

func RedisInit() *redis.Client {
	dsn := config.GetEnv("REDIS_ADDR")
	if len(dsn) == 0 {
		log.Panic("Redis와 관련된 필수적인 환경변수를 찾을 수 없네요.")
	}

	client := redis.NewClient(&redis.Options{
		Addr: dsn,
	})

	return client
}

func CreateJWT(userUUID uuid.UUID) (*models.TokenDetails, error) {
	var err error
	td := new(models.TokenDetails)

	td.AtExpires = time.Now().Add(time.Minute * 15).Unix()
	td.AccessUUID = uuid.New()
	td.UserUUID = userUUID

	td.RtExpires = time.Now().Add(time.Hour * 24 * 7).Unix()
	td.RefreshToken = uuid.New().String()

	key := config.GetEnv("ACCESS_SECRET")
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_uuid"] = td.UserUUID
	atClaims["access_uuid"] = td.AccessUUID
	atClaims["exp"] = td.AtExpires
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	td.AccessToken, err = at.SignedString([]byte(key))
	if err != nil {
		return nil, err
	}

	ref := config.GetEnv("REFRESH_SECRET")
	rtClaims := jwt.MapClaims{}
	rtClaims["refresh_uuid"] = td.RefreshUUID
	rtClaims["user_uuid"] = td.UserUUID
	rtClaims["exp"] = td.RtExpires
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	td.RefreshToken, err = rt.SignedString([]byte(ref))
	if err != nil {
		return nil, err
	}

	return td, nil
}

func InsertRedisAuth(userUUID uuid.UUID, td *models.TokenDetails) error {
	at := time.Unix(td.AtExpires, 0)
	rt := time.Unix(td.RtExpires, 0)

	client := RedisInit()
	err := client.Set(td.AccessUUID.String(), userUUID.String(), at.Sub(time.Now())).Err()
	if err != nil {
		return err
	}
	err = client.Set(td.RefreshUUID.String(), userUUID.String(), rt.Sub(time.Now())).Err()
	if err != nil {
		return err
	}
	return nil
}

func FetchAuth(authD *models.AccessDetails) (string, error) {
	client := RedisInit()

	result, err := client.Get(authD.AccessUUID).Result()
	if err != nil {
		return "", err
	}
	if len(result) == 0 {
		return "", errors.New("입력하신 정보를 통해서 인증 정보를 찾을 수 없어요. 다시 시도해 주세요.")
	}

	return result, nil
}

func ExtractTokenMetaData(r *fiber.Ctx) (*models.AccessDetails, error) {
	token, err := VerifyToken(r)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)

	if ok && token.Valid {
		accessUUID, ok := claims["access_uuid"].(string)
		if !ok {
			return nil, nil
		}

		userUUID, ok := claims["user_uuid"].(string)
		if !ok {
			return nil, nil
		}

		// 구조체 관련 오류 발생으로 인해서 검증이 되지 않았음.
		return &models.AccessDetails{
			AccessUUID: accessUUID,
			UserUUID:   userUUID,
		}, nil
	}
	return nil, err
}

func DeleteAuth(tokenUUID string) (int64, error) {
	client := RedisInit()
	deleted, err := client.Del(tokenUUID).Result()

	return deleted, err
}
