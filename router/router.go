package router

import (
	"time"

	"github.com/Gaurav-malwe/login-service/internal/controller"
	"github.com/Gaurav-malwe/login-service/internal/repository"
	"github.com/Gaurav-malwe/login-service/mongodb"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

const (
	basePath = "/login-service"
)

type LoginServer struct {
	Repository repository.Repository
	Router     *gin.Engine
}

func NewLoginServiceAPIServer(mongodbProvider *mongodb.MongoDB) *LoginServer {
	server := new(LoginServer)
	server.Repository = repository.New(mongodbProvider)
	server.Router = gin.New()

	// CORS middleware configuration
	config := cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	})
	server.Router.Use(config)
	return server
}

func middlewareSetPath(path string) func(ginCtx *gin.Context) {
	return func(c *gin.Context) {
		c.Set("path", path)
		c.Next()
	}
}

func (s *LoginServer) RegisterHandlers(c controller.Controller, middlewares ...gin.HandlerFunc) {
	group := s.Router.Group(basePath) // middleware.TransactionInMiddleware(),
	// middleware.SetPath(basePath),
	// middleware.RequestLogger(),

	group.Use(middlewares...)

	group.POST("/signup", c.Register)
	group.POST("/login", c.Login)

}
