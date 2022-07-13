package api

import (
	"PR_gin_g/controller"
	"PR_gin_g/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

type VideoApi struct {
	loginController controller.LoginController
	videoController controller.VideoController
}

func NewVideoApi(loginController controller.LoginController, videoController controller.VideoController) *VideoApi {
	return &VideoApi{
		loginController: loginController,
		videoController: videoController,
	}
}

func (api *VideoApi) Authenticate(ctx *gin.Context) {
	token := api.loginController.Login(ctx)
	if token != "" {
		ctx.JSON(http.StatusOK, gin.H{"token": token})
	} else {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
	}
}

func (api *VideoApi) GetVideos(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"videos": api.videoController.FindAll()})
}

func (api *VideoApi) CreateVideo(ctx *gin.Context) {
	err := api.videoController.Save(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, &dto.Response{Message: "video created"})
	}
}

func (api *VideoApi) UpdateVideo(ctx *gin.Context) {
	err := api.videoController.Update(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, &dto.Response{Message: "video updated"})
	}
}

func (api *VideoApi) DeleteVideo(ctx *gin.Context) {
	err := api.videoController.Delete(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, &dto.Response{Message: "video deleted"})
	}
}
