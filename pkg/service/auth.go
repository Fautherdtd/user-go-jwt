package service

import (
	"crypto/sha1"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/fautherdtd/user-restapi/entities"
	"github.com/fautherdtd/user-restapi/pkg/repository"
	"github.com/fautherdtd/user-restapi/pkg/smsc"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

const (
	tokenTTL = 12 * time.Hour
)

type tokenClaims struct {
	jwt.StandardClaims
	UserID    int    `json:"user_id"`
	UserName  string `json:"user_name"`
	UserPhone string `json:"user_phone"`
}

// AuthService ...
type AuthService struct {
	repo repository.Authorization
}

// NewAuthService ...
func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{
		repo: repo,
	}
}

// CreateUserWithPhone ...
// Регистрация пользователя с sign-up/start
func (s *AuthService) CreateUserWithPhone(user entities.User) (int, error) {
	return s.repo.CreateUserWithPhone(user)
}

// GenerateSmsCode ...
// Генерация смс-кода для подтверждения
func (s *AuthService) GenerateSmsCode(id int, phone string) (bool, error) {
	code, err := smsc.GenerateCodeAndSave(viper.GetString("sms-code.auth"), id)
	if err != nil {
		return false, err
	}

	err = smsc.SendSmsCode(&smsc.SmsClient{
		Phones:  phone,
		Message: strconv.Itoa(code),
	})
	if err != nil {
		logrus.Errorf("error send-sms: %s", err.Error())
		return false, err
	}

	return true, nil
}

// ConfirmUserByCode ...
// Подтверждение пользователя по коду
func (s *AuthService) ConfirmUserByCode(id int) error {
	return s.repo.ConfirmUserByCode(id)
}

// CheckUserVerification ...
func (s *AuthService) CheckUserVerification(id int) (bool, error) {
	return s.repo.CheckUserVerification(id)
}

// GenerateToken ...
// Генерация JWT token
func (s *AuthService) GenerateToken(user entities.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.ID,
		user.Name,
		user.Phone,
	})

	return token.SignedString([]byte(os.Getenv("SIGNING_KEY")))
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(os.Getenv("SALT_KEY"))))
}
