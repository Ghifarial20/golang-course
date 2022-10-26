package router

import (
	"final-project/controllers"
	"final-project/database"
	"final-project/middlewares"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	db := database.StartDB()
	inDB := &controllers.InDB{DB: db}
	r := gin.Default()

	userRouter := r.Group("/users")
	{
		userRouter.POST("/register", inDB.UserRegister)
		userRouter.POST("/login", inDB.UserLogin)
		userRouter.PUT("/:userID", middlewares.Authentication(), middlewares.UserAuthorization(), inDB.UpdateUser)
		userRouter.DELETE("/:userID", middlewares.Authentication(), middlewares.UserAuthorization(), inDB.DeleteUser)
	}

	photoRouter := r.Group("/photos")
	{
		photoRouter.Use(middlewares.Authentication())
		photoRouter.GET("/", middlewares.UserAuthorization(), inDB.GetPhotos)
		photoRouter.POST("/", middlewares.UserAuthorization(), inDB.CreatePhoto)
		photoRouter.PUT("/:photoId", middlewares.UserAuthorization(), inDB.UpdatePhoto)
		photoRouter.DELETE("/:photoId", middlewares.UserAuthorization(), inDB.DeletePhoto)
	}
	commentRouter := r.Group("/comments")
	{
		commentRouter.Use(middlewares.Authentication())
		commentRouter.GET("/", middlewares.UserAuthorization(), inDB.GetComment)
		commentRouter.POST("/", middlewares.UserAuthorization(), inDB.CreateComment)
		commentRouter.PUT("/:commentId", middlewares.UserAuthorization(), inDB.UpdateComment)
		commentRouter.DELETE("/:commentId", middlewares.UserAuthorization(), inDB.DeleteComment)
	}
	sosmedRouter := r.Group("/socialmedias")
	{
		sosmedRouter.Use(middlewares.Authentication())
		sosmedRouter.GET("/", middlewares.UserAuthorization(), inDB.GetSocialMedias)
		sosmedRouter.POST("/", middlewares.UserAuthorization(), inDB.CreateSocialMedia)
		sosmedRouter.PUT("/:socialMediaId", middlewares.UserAuthorization(), inDB.UpdateSocialMedia)
		sosmedRouter.DELETE("/:socialMediaId", middlewares.UserAuthorization(), inDB.DeleteSocialMedia)
	}

	return r
}
