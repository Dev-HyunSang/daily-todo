package todos

import (
	"time"

	"github.com/dev-hyunsang/daily-todo/auth"
	"github.com/dev-hyunsang/daily-todo/database"
	"github.com/dev-hyunsang/daily-todo/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func CreateToDoHandler(c *fiber.Ctx) error {
	req := new(models.RequestCreateToDo)
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrResponse{
			ErrMetaData: models.ErrMetaData{
				IsSuccess: false,
				Message:   "잘못된 요청입니다. 다시 시도해 주세요.",
			},
			ResponsedAt: time.Now(),
		})
	}

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

	data, err := database.SearchUserByUUID(au.UserUUID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ErrResponse{
			ErrMetaData: models.ErrMetaData{
				IsSuccess:  false,
				StatusCode: fiber.StatusInternalServerError,
				Message:    "사용자를 찾을 수  없습니다. 잠시후 다시 시도해 주세요.",
			},
			ResponsedAt: time.Now(),
		})
	}

	resultToDo, err := database.CreateToDo(models.ToDo{
		ToDoUUID:  uuid.New(),
		UserUUID:  data.UserUUID,
		IsDone:    false,
		Context:   req.Context,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ErrResponse{
			ErrMetaData: models.ErrMetaData{
				IsSuccess:  false,
				StatusCode: fiber.StatusInternalServerError,
				Message:    "사용자 분께서 입력해 주신 소중한 정보를 저장하던 도중 오류가 발생했어요. 잠시후 다시 시도해 주세요.",
				ErrMessage: err.Error(),
			},
			ResponsedAt: time.Now(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(models.ResponseDoneCreateToDo{
		MetaData: models.MetaData{
			IsSuccess:  true,
			StatusCode: fiber.StatusOK,
			Message:    "사용자 분의 소중한 할일이 성공적으로 등록되었어요!",
		},
		Data:        resultToDo,
		ResponsedAt: time.Now(),
	})
}

func AllListToDoHandler(c *fiber.Ctx) error {
	au, err := auth.ExtractTokenMetaData(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(models.ErrResponse{
			ErrMetaData: models.ErrMetaData{
				IsSuccess:  false,
				StatusCode: fiber.StatusUnauthorized,
				Message:    "잘못된 접근이예요. 로그인 이후에 시도해 주세요!",
				ErrMessage: err.Error(),
			},
			ResponsedAt: time.Now(),
		})
	}

	result, err := database.AllListToDo(au.UserUUID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ErrResponse{})
	}

	return c.Status(fiber.StatusOK).JSON(models.ResponseDoneAllListToDo{
		MetaData: models.MetaData{
			IsSuccess:  true,
			StatusCode: fiber.StatusOK,
			Message:    "성공적으로 기록된 할 일들을 불러왔어요!",
		},
		Data:        result,
		ResponsedAt: time.Now(),
	})
}

func EditToDoHandler(c *fiber.Ctx) error {
	return nil
}

func DeleteToDoHandler(c *fiber.Ctx) error {
	req := new(models.RequestDeleteToDo)
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusMethodNotAllowed).JSON(models.ErrResponse{})
	}

	au, err := auth.ExtractTokenMetaData(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(models.ErrResponse{
			ErrMetaData: models.ErrMetaData{
				IsSuccess:  false,
				StatusCode: fiber.StatusUnauthorized,
				Message:    "잘못된 접근이예요. 로그인 이후에 시도해 주세요!",
				ErrMessage: err.Error(),
			},
			ResponsedAt: time.Now(),
		})
	}

	result, err := database.DeleteToDo(au.UserUUID, req.ToDoUUID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ErrResponse{
			ErrMetaData: models.ErrMetaData{
				IsSuccess:  false,
				StatusCode: fiber.StatusInternalServerError,
				Message:    "삭제하던 도중 오류가 발생 했어요. 잠시후 다시 시도해 주세요.",
				ErrMessage: err.Error(),
			},
			ResponsedAt: time.Now(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(models.ResponseDoneDeleteToDo{
		MetaData: models.MetaData{
			IsSuccess:  true,
			StatusCode: fiber.StatusOK,
			Message:    "성공적으로 사용자의 할 일을 삭제했습니다.",
		},
		Data:        result,
		ResponsedAt: time.Now(),
	})
}
