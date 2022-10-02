package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ikhsanfrcn/vix-btpn/config"
	"github.com/ikhsanfrcn/vix-btpn/controller"
	"github.com/ikhsanfrcn/vix-btpn/middleware"
	"github.com/ikhsanfrcn/vix-btpn/repository"
	"github.com/ikhsanfrcn/vix-btpn/service"
	"gorm.io/gorm"
)

var (
	db              *gorm.DB                   = config.SetupDatabaseConnection()
	userRepository  repository.UserRepository  = repository.NewUserRepository(db)
	photoRepository repository.PhotoRepository = repository.NewPhotoRepository(db)
	jwtService      service.JWTService         = service.NewJWTService()
	userService     service.UserService        = service.NewUserService(userRepository)
	photoService    service.PhotoService       = service.NewPhotoService(photoRepository)
	authService     service.AuthService        = service.NewAuthService(userRepository)
	authController  controller.AuthController  = controller.NewAuthController(authService, jwtService)
	userController  controller.UserController  = controller.NewUserController(userService, jwtService)
	photoController controller.PhotoController = controller.NewPhotoController(photoService, jwtService)
)

func main() {
	defer config.CloseDatabaseConnection(db)
	r := gin.Default()

	authRoutes := r.Group("api/users")
	{
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/register", authController.Register)
	}

	userRoutes := r.Group("api/users", middleware.AuthorizeJWT(jwtService))
	{
		userRoutes.GET("/profile", userController.Profile)
		userRoutes.PUT("/userId", userController.Update)
	}

	photoRoutes := r.Group("api/", middleware.AuthorizeJWT(jwtService))
	{
		photoRoutes.GET("/photos", photoController.All)
		photoRoutes.POST("/photos", photoController.Insert)
		photoRoutes.PUT("/photoId", photoController.Update)
		photoRoutes.DELETE("/:photoId", photoController.Delete)
	}

	r.Run()
}
