package handler

import (
	"net/http"

	"github.com/fautherdtd/user-restapi/entities"
	"github.com/fautherdtd/user-restapi/pkg/smsc"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func (h *Handler) signUpStart(c *gin.Context) {
	var user entities.User
	if err := c.BindJSON(&user); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	// Создание пользователя
	id, err := h.services.Authorization.CreateUserWithPhone(user)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	user.ID = id

	// Генерация JWT-token
	token, err := h.services.Authorization.GenerateToken(user)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	// Генерация sms-кода
	status, err := h.services.GenerateSmsCode(user.ID, user.Phone)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token":   token,
		"confirm": status,
	})
}

func (h *Handler) signUpConfirm(c *gin.Context) {
	var input entities.Confirm

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	// Проверка на уже подтвержденного пользователя
	verification, err := h.services.Authorization.CheckUserVerification(input.UserID)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	if verification {
		c.JSON(http.StatusOK, map[string]interface{}{
			"status":  false,
			"message": "Вы уже верифицированы.",
		})
		return
	}

	// Подтверждение смс-кода
	status, err := smsc.ConfirmCodeFromStorage(viper.GetString("sms-code.auth"), input.UserID, input.Code)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if !status {
		c.JSON(http.StatusOK, map[string]interface{}{
			"status":  status,
			"message": "Неверный код.",
		})
		return
	}

	// Верификация пользователя
	if err = h.services.ConfirmUserByCode(input.UserID); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"status":  status,
		"message": "Код подтвержден. Успешная верификация пользователя.",
	})
}

func (h *Handler) signUpReConfirm(c *gin.Context) {
	var user entities.Confirm
	if err := c.BindJSON(&user); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	// Проверка на уже подтвержденного пользователя
	verification, err := h.services.Authorization.CheckUserVerification(user.UserID)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	if verification {
		c.JSON(http.StatusOK, map[string]interface{}{
			"status":  false,
			"message": "Вы уже верифицированы.",
		})
		return
	}

	// Генерация sms-кода
	status, err := h.services.GenerateSmsCode(user.UserID, user.Phone)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"confirm": status,
	})
}

func (h *Handler) signInWithPassword(c *gin.Context) {
	// ...
}

func (h *Handler) signInWithSms(c *gin.Context) {
	// ...
}
