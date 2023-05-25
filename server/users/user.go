package users

import (
	"log"
	"time"

	"github.com/dev-hyunsang/daily-todo/auth"
	"github.com/dev-hyunsang/daily-todo/database"
	"github.com/dev-hyunsang/daily-todo/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func JoinUserHandler(c *fiber.Ctx) error {
	req := new(models.RequestJoinUser)

	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrResponse{
			ErrMetaData: models.ErrMetaData{
				IsSuccess:  false,
				StatusCode: fiber.StatusBadRequest,
				Message:    err.Error(),
			},
			ResponsedAt: time.Now(),
		})
	}

	// 입력 값 검증.
	if len(req.Email) == 0 || len(req.Password) == 0 || len(req.NickName) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrResponse{
			ErrMetaData: models.ErrMetaData{
				IsSuccess:  false,
				StatusCode: fiber.StatusBadRequest,
				Message:    "필수 입력값이 입력되지 않았습니다. 확인 후 다시 시도해 주세요.",
			},
			ResponsedAt: time.Now(),
		})
	}

	// 데이터가 없는 경우 오류 발생
	// TODO: 이메일 중복 검사 알고리즘 개선 필요.
	result, err := database.SearchUserByEmail(req.Email)
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(models.ErrResponse{
			ErrMetaData: models.ErrMetaData{
				IsSuccess:  false,
				StatusCode: fiber.StatusInternalServerError,
				Message:    err.Error(),
			},
			ResponsedAt: time.Now(),
		})
	} else if result.Email == req.Email {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrResponse{
			ErrMetaData: models.ErrMetaData{
				IsSuccess:  false,
				StatusCode: fiber.StatusBadRequest,
				Message:    "이미 존재하는 이메일입니다. 확인 후 다시 시도해 주세요.",
				ErrMessage: err.Error(),
			},
			ResponsedAt: time.Now(),
		})
	}

	userUUID := uuid.New()
	pwHash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ErrResponse{
			ErrMetaData: models.ErrMetaData{
				IsSuccess:  false,
				StatusCode: fiber.StatusInternalServerError,
				Message:    err.Error(),
			},
			ResponsedAt: time.Now(),
		})
	}

	result, err = database.CreateUser(models.User{
		UserUUID:  userUUID,
		Email:     req.Email,
		Password:  string(pwHash),
		NickName:  req.NickName,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ErrResponse{
			ErrMetaData: models.ErrMetaData{
				IsSuccess:  false,
				StatusCode: fiber.StatusInternalServerError,
				Message:    err.Error(),
			},
			ResponsedAt: time.Now(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(models.ResponseDoneJoinUser{
		MetaData: models.MetaData{
			IsSuccess:  true,
			StatusCode: fiber.StatusOK,
			Message:    "어서와요! 성공적으로 회원가입이 완료되었습니다!",
		},
		Data:        result,
		ResponsedAt: time.Now(),
	})
}

func LoginUserHandler(c *fiber.Ctx) error {
	req := new(models.RequestLoginUser)

	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrResponse{
			ErrMetaData: models.ErrMetaData{
				IsSuccess:  false,
				StatusCode: fiber.StatusBadRequest,
				Message:    err.Error(),
			},
			ResponsedAt: time.Now(),
		})
	}

	user, err := database.SearchUserByEmail(req.Email)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ErrResponse{
			ErrMetaData: models.ErrMetaData{
				IsSuccess:  false,
				StatusCode: fiber.StatusInternalServerError,
				Message:    err.Error(),
			},
		})
	} else if user.Email != req.Email {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrResponse{
			ErrMetaData: models.ErrMetaData{
				IsSuccess:  false,
				StatusCode: fiber.StatusBadRequest,
				Message:    "존재하지 않는 이메일입니다. 확인 후 다시 시도해 주세요.",
			},
			ResponsedAt: time.Now(),
		})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(models.ErrResponse{
			ErrMetaData: models.ErrMetaData{
				IsSuccess:  false,
				StatusCode: fiber.StatusUnauthorized,
				Message:    "비밀번호가 일치하지 않습니다. 확인 후 다시 시도해 주세요.",
			},
			ResponsedAt: time.Now(),
		})
	}

	ts, err := auth.CreateToken(user.UserUUID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ErrResponse{
			ErrMetaData: models.ErrMetaData{
				IsSuccess:  false,
				StatusCode: fiber.StatusInternalServerError,
				Message:    err.Error(),
			},
			ResponsedAt: time.Now(),
		})
	}

	err = auth.InsertRedisAuth(user.UserUUID, ts)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ErrResponse{
			ErrMetaData: models.ErrMetaData{
				IsSuccess:  false,
				StatusCode: fiber.StatusInternalServerError,
				Message:    err.Error(),
			},
			ResponsedAt: time.Now(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(models.ResponseDoneLoginUser{
		MetaData: models.MetaData{
			IsSuccess:  true,
			StatusCode: fiber.StatusOK,
			Message:    "어서와요! 성공적으로 로그인이 완료되었습니다!",
		},
		Data:        *ts,
		ResponsedAt: time.Now(),
	})
}

func LogoutUserHandler(c *fiber.Ctx) error {
	au, err := auth.ExtractTokenMetaData(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(models.ErrResponse{
			ErrMetaData: models.ErrMetaData{
				IsSuccess:  false,
				StatusCode: fiber.StatusUnauthorized,
				Message:    "잘못된 접근이예요. 로그인 이후에 시도해 주세요!",
			},
			ResponsedAt: time.Now(),
		})
	}

	deleted, err := auth.DeleteAuth(au.AccessUUID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ErrResponse{
			ErrMetaData: models.ErrMetaData{
				IsSuccess:  false,
				StatusCode: fiber.StatusInternalServerError,
				Message:    err.Error(),
			},
			ResponsedAt: time.Now(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(models.ResponseDoneLogoutUser{
		MetaData: models.MetaData{
			IsSuccess:  false,
			StatusCode: fiber.StatusOK,
			Message:    "성공적으로 로그아웃 하였습니다.",
		},
		Data:        deleted,
		ResponsedAt: time.Now(),
	})
}
