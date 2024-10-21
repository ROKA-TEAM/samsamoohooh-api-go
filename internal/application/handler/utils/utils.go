package utils

import (
	"crypto/rand"
	"encoding/base64"
	"github.com/gofiber/fiber/v3"
	"samsamoohooh-go-api/internal/infra/middleware/guard"
	"samsamoohooh-go-api/pkg/token"
)

func GenerateState() string {
	// 32 바이트의 랜덤 데이터를 생성
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		return ""
	}

	// base64 인코딩을 사용하여 랜덤 데이터를 문자열로 변환
	return base64.URLEncoding.EncodeToString(b)
}

func GetToken(c fiber.Ctx) (*token.Token, error) {
	t, ok := c.Locals(guard.TokenKey).(*token.Token)
	if !ok {
		return nil, fiber.NewError(fiber.StatusInternalServerError, "unable to get token")
	}

	return t, nil
}
