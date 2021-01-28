package handler

import (
	"github.com/fautherdtd/user-restapi/pkg/service"
	"github.com/gin-gonic/gin"
)

// Handler ...
type Handler struct {
	services *service.Service
}

// NewHandler ...
func NewHandler(services *service.Service) *Handler {
	return &Handler{
		services: services,
	}
}

// InitRoutes ...
func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		signUp := auth.Group("/sign-up")
		{
			signUp.POST("/start", h.signUpStart)
			signUp.POST("/confirm", h.signUpConfirm)
			signUp.POST("/re-confirm", h.signUpReConfirm)
		}

		authWith := auth.Group("/sign-in")
		{
			authWith.POST("/password", h.signInWithPassword)
			authWith.POST("/sms", h.signInWithSms)
		}
	}

	user := router.Group("/user")
	{
		user.GET("/get/:id", h.getUser)
		user.PUT("/update/:id", h.updateUser)
		user.DELETE("/delete/:id", h.deleteUser)
	}

	return router
}
