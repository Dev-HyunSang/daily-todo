package users_test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/dev-hyunsang/daily-todo/database"
	"github.com/dev-hyunsang/daily-todo/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func TestJoinUser(t *testing.T) {
	log.Println("[TEST] TestJoinUser")

	app := fiber.New()
	app.Post("/api/user/join", func(c *fiber.Ctx) error {
		req := new(models.RequestJoinUser)

		if err := c.BodyParser(req); err != nil {
			log.Println(err)
			return c.Status(fiber.StatusBadRequest).JSON(models.ErrResponse{
				ErrMetaData: models.ErrMetaData{
					IsSuccess:  false,
					StatusCode: fiber.StatusBadRequest,
					Message:    "잘못된 요청입니다. 확인 후 다시 시도해 주세요.",
					ErrMessage: err.Error(),
				},
				ResponsedAt: time.Now(),
			})
		}

		if len(req.Email) == 0 || len(req.Password) == 0 || len(req.NickName) == 0 {
			log.Println("NOT INPUT VALUE")
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
					Message:    "서버 내부에서 동일한 메일을 찾던 도중 오류가 발생했어요.",
					ErrMessage: err.Error(),
				},
				ResponsedAt: time.Now(),
			})
		}
		if result.Email == req.Email {
			log.Println("ALREADY EXIST EMAIL")
			return c.Status(fiber.StatusBadRequest).JSON(models.ErrResponse{
				ErrMetaData: models.ErrMetaData{
					IsSuccess:  false,
					StatusCode: fiber.StatusBadRequest,
					Message:    "이미 존재하는 이메일입니다. 확인 후 다시 시도해 주세요.",
				},
				ResponsedAt: time.Now(),
			})
		}

		userUUID := uuid.New()
		pwHash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
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
			log.Println(err)
			return c.Status(fiber.StatusInternalServerError).JSON(models.ErrResponse{
				ErrMetaData: models.ErrMetaData{
					IsSuccess:  false,
					StatusCode: fiber.StatusInternalServerError,
					Message:    "사용자의 소중한 정보를 저장하던 도중 오류가 발생했어요. 잠시후 다시 시도해 주세요.",
					ErrMessage: err.Error(),
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
	})

	testReqJson := models.RequestJoinUser{
		Email:    "parkhyunsang0625@gmail.com",
		Password: "password",
		NickName: "HyunSang Park",
	}

	payload, err := json.Marshal(testReqJson)
	if err != nil {
		t.Fatal(err)
	}
	buff := bytes.NewBuffer(payload)

	req := httptest.NewRequest(http.MethodPost, "/api/user/join", buff)
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	if err != nil {
		t.Fatal(err)
	}

	defer resp.Body.Close()

	respBodyData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	log.Println(string(respBodyData))

	// 만약 상태 코드가 오류 상태이면 Fail 함.
	if resp.StatusCode != http.StatusOK {
		log.Println(resp.Request.Response.Body)
		log.Println(resp.Body)
		t.Errorf("Expected status code %d but got %d", http.StatusOK, resp.StatusCode)
	} else {
		t.Log("[OK] DONE!")
	}
}

func TestLoginUser(t *testing.T) {

}

func TestLogoutUser(t *testing.T) {

}
