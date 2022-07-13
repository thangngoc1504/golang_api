package main

import (
	"PR_gin_g/controller"
	"PR_gin_g/middlewares"
	"PR_gin_g/repository"
	"PR_gin_g/service"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
	"github.com/swaggo/swag/example/basic/docs"
)

var (
	videoRepository repository.VideoRepository = repository.NewVidepRepository()
	videoService    service.VideoService       = service.New(videoRepository)
	loginService    service.LoginService       = service.NewLoginService()
	jwtService      service.JWTService         = service.NewJWTService()

	loginController controller.LoginController = controller.NewLoginController(loginService, jwtService)
	videoController controller.VideoController = controller.New(videoService)
)

func main() {

	// Swagger
	docs.SwaggerInfo.Title = "Gin API"
	docs.SwaggerInfo.Description = "Gin API Documentation"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = "/api"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	defer videoRepository.CloseDB()

	r := gin.New()

	r.Static("/css", "./templates/css")

	r.LoadHTMLGlob("templates/*.html")

	r.Use(gin.Recovery(), gin.Logger())

	r.POST("/login", func(ctx *gin.Context) {
		token := loginController.Login(ctx)
		if token != "" {
			ctx.JSON(http.StatusOK, gin.H{"token": token})
		} else {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		}

	})

	apiRoutes := r.Group("/api", middlewares.AuthorizeJWT())
	{
		apiRoutes.GET("/videos", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, videoController.FindAll())
		})

		apiRoutes.GET("/videos/:id", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, videoController.Detail(ctx))
		})

		apiRoutes.POST("/videos", func(ctx *gin.Context) {
			err := videoController.Save(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "Video created"})
			}
		})

		apiRoutes.PUT("/videos/:id", func(ctx *gin.Context) {
			err := videoController.Update(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "Video updated"})
			}
		})

		apiRoutes.DELETE("/videos/:id", func(ctx *gin.Context) {
			err := videoController.Delete(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "Video deleted"})
			}
		})
	}

	viewRoutes := r.Group("/view")
	{
		viewRoutes.GET("/videos", videoController.ShowAll)
	}

	// Swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r.Run(":" + port)

}
