package main

import (
	"log"
	"os"

	"github.com/mci-its/backend-service/cmd"
	"github.com/mci-its/backend-service/config"
	"github.com/mci-its/backend-service/controller"
	"github.com/mci-its/backend-service/middleware"
	"github.com/mci-its/backend-service/repository"
	"github.com/mci-its/backend-service/routes"
	"github.com/mci-its/backend-service/service"

	"github.com/gin-gonic/gin"
)

func main() {
	db := config.SetUpDatabaseConnection()
	defer config.CloseDatabaseConnection(db)

	if len(os.Args) > 1 {
		cmd.Commands(db)
		return
	}

	var (
		jwtService service.JWTService = service.NewJWTService()

		// Implementation Dependency Injection
		// Repository
		userRepository            repository.UserRepository            = repository.NewUserRepository(db)
		violationRepository       repository.ViolationRepository       = repository.NewViolationRepository(db)
		violationActionRepository repository.ViolationActionRepository = repository.NewViolationActionRepository(db)
		ViolationNoteRepository   repository.ViolationNoteRepository   = repository.NewViolationNoteRepository(db)		
		// Service
		userService            service.UserService            = service.NewUserService(userRepository, jwtService)
		violationService       service.ViolationService       = service.NewViolationService(violationRepository)
		violationActionService service.ViolationActionService = service.NewViolationActionService(violationActionRepository)
		violationNoteService   service.ViolationNoteService   = service.NewViolationNoteService(ViolationNoteRepository)
		// Controller
		userController            controller.UserController            = controller.NewUserController(userService)
		violationController       controller.ViolationController       = controller.NewViolationController(violationService)
		violationActionController controller.ViolationActionController = controller.NewViolationActionController(violationActionService)
		violationNoteController   controller.ViolationNoteController   = controller.NewViolationNoteController(violationNoteService)
	)

	server := gin.Default()
	server.Use(middleware.CORSMiddleware())

	// routes
	routes.User(server, userController, jwtService)
	routes.Violation(server, violationController)
	routes.ViolationAction(server, violationActionController)
	routes.ViolationNote(server, violationNoteController)

	server.Static("/assets", "./assets")
	port := os.Getenv("PORT")
	if port == "" {
		port = "8888"
	}

	var serve string
	if os.Getenv("APP_ENV") == "localhost" {
		serve = "127.0.0.1:" + port
	} else {
		serve = ":" + port
	}

	if err := server.Run(serve); err != nil {
		log.Fatalf("error running server: %v", err)
	}
}
