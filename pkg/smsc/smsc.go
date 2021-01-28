package smsc

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/fautherdtd/user-restapi/pkg/storage"
	"github.com/sirupsen/logrus"
)

const (
	// URL's
	sendURL = "https://smsc.ru/sys/send.php?login=%s&psw=%s&phones=%s&mes=%s&fmt=3"

	// Type codes (prefix)
	codeVerify = "code:sms"
)

// SmsClient ...
type SmsClient struct {
	Phones  string
	Message string
}

// SmsResponse ...
type SmsResponse struct {
	ID int `json:"id"`
}

// SendSmsCode ...
// Отправка кода в смс по номеру телефона
func SendSmsCode(sc *SmsClient) error {
	var sr SmsResponse
	resp, err := http.Get(fmt.Sprintf(sendURL,
		os.Getenv("SMSC_USER"), os.Getenv("SMSC_KEY"), sc.Phones, sc.Message))
	if err != nil {
		logrus.Errorf("error get sms-send: %s", err.Error())
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logrus.Errorf("error ioutil.ReadAll: %s", err.Error())
		return err
	}

	err = json.Unmarshal(body, &sr)
	if err != nil {
		logrus.Errorf("error json.Unmarshal: %s", err.Error())
		return err
	}

	fmt.Print(sr.ID)
	return nil
}

// GenerateCodeAndSave ...
// Генерация и сохранение кода в хранилище
func GenerateCodeAndSave(key string, userID int) (int, error) {
	redis, err := storage.InitRedis()
	if err != nil {
		logrus.Fatalf("error initializing storage: %s", err.Error())
	}

	code := rand.Intn(3000-1000) + 1000
	keySet := fmt.Sprintf("%s-%s-%s", codeVerify, key, strconv.Itoa(userID))
	redis.Set(keySet, code, 900*time.Second).Err()

	return code, nil
}

// ConfirmCodeFromStorage ...
// Подтверждение кода из хранилища по ключу
func ConfirmCodeFromStorage(key string, userID int, code int) (bool, error) {
	redis, err := storage.InitRedis()
	if err != nil {
		logrus.Fatalf("error initializing storage: %s", err.Error())
	}

	key = fmt.Sprintf("%s-%s-%s", codeVerify, key, strconv.Itoa(userID))
	val, err := redis.Get(key).Result()
	if err != nil {
		fmt.Println(err)
	}

	return val == fmt.Sprint((strconv.Itoa(code))), nil

}
