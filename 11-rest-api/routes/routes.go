package routes

import (
	"github.com/gin-gonic/gin"
	"res-api.com/apis/middlewares"
)

func RegisterRoutes(server *gin.Engine) {
	// Route, Middleware, Handler
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)

	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticated)
	authenticated.POST("/events", createEvent)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("/events/:id", deleteEvent)
	authenticated.POST("/events/:id/register", registerForEvent)
	authenticated.DELETE("/events/:id/register", cancelRegistration)

	server.POST("/signup", signup)
	server.POST("/login", login)
}
