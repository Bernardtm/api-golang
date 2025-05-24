package di

import (
	"bernardtm/backend/configs"
	"bernardtm/backend/internal/core/auth"
	"bernardtm/backend/internal/core/auth/token"
	"bernardtm/backend/internal/core/email"
	"bernardtm/backend/internal/core/files"
	"bernardtm/backend/internal/core/menus"
	"bernardtm/backend/internal/core/shareds"
	"bernardtm/backend/internal/core/status"
	"bernardtm/backend/internal/core/storage"
	"bernardtm/backend/internal/core/users"
	"bernardtm/backend/internal/infra/socket"
	"bernardtm/backend/pkg/providers/emails"
	"bernardtm/backend/pkg/providers/storages"
	"database/sql"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
)

type Container struct {
	AuthController        auth.AuthController
	MenusController       menus.MenusController
	UserController        users.UsersController
	StatusController      status.StatusController
	TokenService          token.TokenService
	HealthcheckController shareds.HealthcheckController
	FilesController       files.FilesController
	SocketHandler         socket.SocketController
}

func NewContainer(db *sql.DB, mongoClient *mongo.Client, appConfig *configs.AppConfig) *Container {

	// Providers
	var emailProvider emails.EmailProvider

	if os.Getenv("ENVIRONMENT") == "dev" {
		emailProvider = emails.NewMailpitSMTPEmailProvider(appConfig)
	} else {
		emailProvider = emails.NewMailgunProvider(appConfig)
	}
	storageProvider := storages.NewS3StorageProvider(appConfig)

	// redisClient, err := redis_client.ConnectRedis(appConfig.REDIS_ADDRESS)
	// if err != nil {
	// 	log.Printf("Failed to connect to Redis: %v", err)
	// }
	// queueProvider := queues.NewRedisQueueProvider(redisClient)

	// Repositories
	statusRepo := status.NewStatusRepository(db)
	twoFactorRepo := auth.NewTwoFactorCodesRepository(db)
	userRepo := users.NewUserRepository(db)
	filesRepo := files.NewFilesRepository(db)
	menusRepo := menus.NewMenusRepository(db)

	// Services
	emailService := email.NewEmailService(emailProvider)
	tokenService := token.NewTokenService(appConfig)
	twoFactorService := auth.NewTwoFactorCodesService(twoFactorRepo, statusRepo)
	authService := auth.NewAuthService(userRepo, appConfig, *emailService, twoFactorService, tokenService)
	statusService := status.NewStatusService(statusRepo)
	storageService := storage.NewStorageService(storageProvider)
	filesService := files.NewFilesService(filesRepo, statusRepo, storageService)
	menusService := menus.NewMenusService(menusRepo)
	userService := users.NewUsersService(userRepo, statusRepo, storageService)

	// Controllers
	authController := auth.NewAuthController(authService)
	statusController := status.NewStatusController(statusService)
	healthcheckController := shareds.NewHealthcheckController()
	filesController := files.NewFilesController(filesService)
	menusController := menus.NewMenusController(menusService)
	userController := users.NewUsersController(userService)

	socketHandler := socket.NewSocketController()

	return &Container{
		AuthController:        authController,
		StatusController:      statusController,
		TokenService:          tokenService,
		HealthcheckController: healthcheckController,
		SocketHandler:         socketHandler,
		FilesController:       filesController,
		MenusController:       menusController,
		UserController:        userController,
	}
}
