package router

import (
	"contest/internal/server/handlers"
	"contest/internal/server/middleware"
	"contest/internal/service"
	"contest/internal/storage"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"time"
)

func Run(storage *storage.Storage, services *service.Manager, jwtSecret string, port int) error {
	runTestHandler := handlers.NewRunTestHandler(services.TestRunner, services.Logger)
	testHandler := handlers.NewTestHandler(storage.TestRepository, services.Logger)
	taskHandler := handlers.NewTaskHandler(storage.TaskRepository, services.Logger)
	launchHandler := handlers.NewLaunchHandler(storage.LaunchRepository, services.Logger)

	gin.SetMode(gin.ReleaseMode)
	g := gin.New()

	g.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"*"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "*"
		},
		MaxAge: 12 * time.Hour,
	}))

	g.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	g.Use(middleware.AuthMiddleware(services.Logger, jwtSecret))
	api := g.Group("/api")
	{
		api.POST("/run", runTestHandler.RunTest)
		api.PUT("/test", testHandler.AddTest)
		api.DELETE("/test/:id", testHandler.DeleteTest)
		api.PATCH("/test/:id", testHandler.UpdateTest)
		api.GET("/test/:id", testHandler.GetTest)

		api.GET("/tests/:task_id", testHandler.GetTestsByTaskID)
		api.GET("/tests", testHandler.GetTests)

		api.PUT("/task", taskHandler.AddTask)
		api.DELETE("/task/:id", taskHandler.DeleteTask)
		api.PATCH("/task/:id", taskHandler.UpdateTask)
		api.GET("/task/:id", taskHandler.GetTask)

		api.GET("/tasks", taskHandler.GetAllTasks)

		api.GET("/launches/success/:user_id", launchHandler.GetSuccessLaunchesByUser)
		api.GET("/launches/:user_id/:contest_id", launchHandler.GetLaunchesByUserAndContest)
		api.GET("/launches/:user_id", launchHandler.GetLaunchesByUser)
	}

	if err := g.Run(fmt.Sprintf(":%d", port)); err != nil {
		return err
	}
	return nil
}
