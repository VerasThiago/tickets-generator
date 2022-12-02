package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/verasthiago/tickets-generator/login/pkg/builder"
	"github.com/verasthiago/tickets-generator/login/pkg/validator"
	"github.com/verasthiago/tickets-generator/shared/auth"
	error_handler "github.com/verasthiago/tickets-generator/shared/errors"
	"github.com/verasthiago/tickets-generator/shared/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type LoginUserAPI interface {
	Handler(context *gin.Context)
}

type LoginUserHandler struct {
	builder.Builder
}

func (l *LoginUserHandler) InitFromBuilder(builder builder.Builder) *LoginUserHandler {
	l.Builder = builder
	return l
}
func (l *LoginUserHandler) Handler(context *gin.Context) {
	var err error
	var user *models.User
	var tokenString string
	var request validator.SignInRequest

	if err := context.ShouldBindJSON(&request); err != nil {
		error_handler.HandleBadRequestError(context, err)
		return
	}

	if errors := request.Validate(); len(errors) > 0 {
		error_handler.HandleBadRequestErrors(context, errors)
		return
	}

	if user, err = l.GetRepository().GetUserByCPF(request.CPF); err != nil {
		error_handler.HandleInternalServerError(context, err, l.GetLog())
		return
	}

	if err := user.CheckPassword(request.Password); err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			err = gorm.ErrRecordNotFound
		}
		error_handler.HandleInternalServerError(context, err, l.GetLog())
		return
	}

	if tokenString, err = auth.GenerateJWT(user.Email, user.Username, user.ID, user.CPF, l.GetSharedFlags().JwtKey, user.IsAdmin); err != nil {
		error_handler.HandleInternalServerError(context, err, l.GetLog())
		return
	}

	context.JSON(http.StatusOK, gin.H{"status": "success", "token": tokenString})
}
