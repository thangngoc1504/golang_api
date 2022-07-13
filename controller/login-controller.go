package controller

import (
	"PR_gin_g/dto"
	"PR_gin_g/service"

	"github.com/gin-gonic/gin"
)

type LoginController interface {
	Login(ctx *gin.Context) string
}

type loginController struct {
	loginService service.LoginService
	jwtService   service.JWTService
}

func NewLoginController(loginService service.LoginService, jwtService service.JWTService) LoginController {
	return &loginController{
		loginService: loginService,
		jwtService:   jwtService,
	}
}

func (c *loginController) Login(ctx *gin.Context) string {
	var credentials dto.Credentials
	err := ctx.ShouldBindJSON(&credentials)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return ""
	}
	isAuthenticated := c.loginService.Login(credentials.Username, credentials.Password)
	if isAuthenticated {
		return c.jwtService.GenerateToken(credentials.Username, true)
	}
	return ""
}
