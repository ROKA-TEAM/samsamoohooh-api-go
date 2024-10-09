package jwt

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/pkg/errors"
	"samsamoohooh-go-api/internal/core/domain"
	"samsamoohooh-go-api/internal/core/port"
	"samsamoohooh-go-api/internal/infra/config"
	"strconv"
	"strings"
	"time"
)

var _ port.JWTService = (*JWT)(nil)

type JWT struct {
	config *config.Config

	durations map[domain.TokenType]time.Duration
}

func New(config *config.Config) (*JWT, error) {
	j := JWT{config: config}
	access, err := time.ParseDuration(config.Token.Duration.Access)
	if err != nil {
		return nil, err
	}
	refresh, err := time.ParseDuration(config.Token.Duration.Refresh)
	if err != nil {
		return nil, err
	}

	temp, err := time.ParseDuration(config.Token.Duration.Temp)
	if err != nil {
		return nil, err
	}

	j.durations = map[domain.TokenType]time.Duration{
		domain.Access:  access,
		domain.Refresh: refresh,
		domain.Temp:    temp,
	}

	return &j, nil
}

func (j *JWT) CreateAccessToken(user *domain.User) (string, error) {
	now := time.Now()

	payload := domain.TokenPayload{
		Issuer:    j.config.Token.Issuer,
		Subject:   strconv.Itoa(int(user.ID)),
		Audience:  j.config.Token.Audience,
		ExpiresAt: now.Add(j.durations[domain.Access]),
		NotBefore: now,
		IssuedAt:  now,
	}

	claims := customClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    payload.Issuer,
			Subject:   payload.Subject,
			Audience:  jwt.ClaimStrings{payload.Audience},
			ExpiresAt: jwt.NewNumericDate(payload.ExpiresAt),
			NotBefore: jwt.NewNumericDate(now),
			IssuedAt:  jwt.NewNumericDate(payload.IssuedAt),
		},
		Social: string(user.Social),
		Role:   string(user.Role),
		Type:   string(domain.Access),
	}

	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(j.config.Token.Key))
}

func (j *JWT) CreateRefreshToken(user *domain.User) (string, error) {
	now := time.Now()

	payload := domain.TokenPayload{
		Issuer:    j.config.Token.Issuer,
		Subject:   strconv.Itoa(int(user.ID)),
		Audience:  j.config.Token.Audience,
		ExpiresAt: now.Add(j.durations[domain.Refresh]),
		NotBefore: now,
		IssuedAt:  now,
	}

	var claims = struct {
		jwt.RegisteredClaims
		Social string `json:"social"`
		Type   string `json:"type"`
		Role   string `json:"role"`
	}{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    payload.Issuer,
			Subject:   payload.Subject,
			Audience:  jwt.ClaimStrings{payload.Audience},
			ExpiresAt: jwt.NewNumericDate(payload.ExpiresAt),
			NotBefore: jwt.NewNumericDate(now),
			IssuedAt:  jwt.NewNumericDate(payload.IssuedAt),
		},
		Social: string(user.Social),
		Role:   string(user.Role),
		Type:   string(domain.Refresh),
	}

	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(j.config.Token.Key))
}

func (j *JWT) CreateTempToken(userID uint, sub, social string) (string, error) {
	now := time.Now()

	payload := domain.TokenPayload{
		Issuer:    j.config.Token.Issuer,
		Subject:   strconv.Itoa(int(userID)),
		Audience:  j.config.Token.Audience,
		ExpiresAt: now.Add(j.durations[domain.Temp]),
		NotBefore: now,
		IssuedAt:  now,
		Social:    domain.Google,
	}

	var claims = struct {
		jwt.RegisteredClaims
		Social string `json:"social"`
		Type   string `json:"type"`
	}{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    payload.Issuer,
			Audience:  jwt.ClaimStrings{payload.Audience},
			ExpiresAt: jwt.NewNumericDate(payload.ExpiresAt),
			NotBefore: jwt.NewNumericDate(payload.NotBefore),
			IssuedAt:  jwt.NewNumericDate(payload.IssuedAt),
		},
		Social: string(payload.Social),
		Type:   string(domain.Temp),
	}

	fmt.Printf("Created Token %+v\n", claims)

	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(j.config.Token.Key))
}

func (j *JWT) VerifyTempToken(tokenString string) (*domain.TempTokenPayload, error) {
	var myTempTokenClaims tempTokenClaims

	_, err := jwt.ParseWithClaims(tokenString, &myTempTokenClaims, func(token *jwt.Token) (any, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.Wrap(domain.ErrUnauthorized, fmt.Sprintf("unexpected signing method: %v", token.Header["alg"]))
		}
		return []byte(j.config.Token.Key), nil
	})
	if err != nil {
		if errors.Is(err, domain.ErrUnauthorized) {
			return nil, err
		}
		return nil, errors.Wrap(domain.ErrInternal, err.Error())
	}

	fmt.Printf("Parsed temp: %+v\n", myTempTokenClaims)

	// check Audience field
	var check bool = false
	for _, aud := range myTempTokenClaims.Audience {
		if strings.Compare(j.config.Token.Audience, aud) == 0 {
			check = true
			break
		}
	}

	if check == false {
		return nil, errors.Wrap(domain.ErrUnauthorized, "this token cannot be processed by this server, please check the Audience field during token claims.")
	}

	// Check ExpiresAt, 토큰이 만료 되었나?
	// 현재 시간이 expireseAt보다 뒤에(before)에 있을 때
	// 과거 <----expirseAt --- time.Now ---> 미래
	if myTempTokenClaims.ExpiresAt.Before(time.Now()) {
		// 토큰이 만료되었음!
		return nil, errors.Wrap(domain.ErrUnauthorized, "token has expired.")
	}

	// only temp token인가? (token의 type payload 검사를 통해 검사하자)
	if !(strings.Compare(myTempTokenClaims.Type, string(domain.Temp)) == 0) {
		return nil, domain.ErrTokenNotTemporary
	}
	return myTempTokenClaims.toDomain(), nil
}

func (j *JWT) VerifyToken(tokenString string) (*domain.TokenPayload, error) {
	myCustomClaims := customClaims{}

	_, err := jwt.ParseWithClaims(tokenString, &myCustomClaims, func(token *jwt.Token) (any, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.Wrap(domain.ErrUnauthorized, fmt.Sprintf("unexpected signing method: %v", token.Header["alg"]))
		}
		return []byte(j.config.Token.Key), nil
	})
	if err != nil {
		if errors.Is(err, domain.ErrUnauthorized) {
			return nil, err
		}
		return nil, errors.Wrap(domain.ErrInternal, err.Error())
	}

	fmt.Printf("Parsed : %+v\n", myCustomClaims)

	// check Audience field
	var check bool = false
	for _, aud := range myCustomClaims.Audience {
		// is same
		if strings.Compare(j.config.Token.Audience, aud) == 0 {
			check = true
			break
		}
	}

	if check == false {
		return nil, errors.Wrap(domain.ErrUnauthorized, "this token cannot be processed by this server, please check the Audience field during token claims.")
	}

	// Check ExpiresAt, 토큰이 만료 되었나?
	// 현재 시간이 expireseAt보다 뒤에(before)에 있을 때
	// 과거 <----expirseAt --- time.Now ---> 미래
	if myCustomClaims.ExpiresAt.Before(time.Now()) {
		// 토큰이 만료되었음!
		return nil, errors.Wrap(domain.ErrUnauthorized, "token has expired.")

	}

	// TODO: not before
	fmt.Println("myCustomClaims.Type: ", myCustomClaims.Type)

	if strings.Compare(myCustomClaims.Type, "temp") == 0 {
		return nil, domain.ErrTokenTemporary
	}

	return myCustomClaims.toDomain(), nil
}
