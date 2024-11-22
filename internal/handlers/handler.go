package handlers

import (
	"awesomeProject/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
	router  *gin.Engine
	service *service.Service
}

func NewHandler(router *gin.Engine, s *service.Service) *Handler {
	return &Handler{
		router:  router,
		service: s,
	}
}

func (h *Handler) HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "The service is up and running"})
}

func (h *Handler) InitRoutes() {
	{
		h.router.GET("/health", h.HealthCheck)
		h.router.POST("/users/add", h.SignUp)
		h.router.POST("/users/login", h.SignIn)
	}

	v1 := h.router.Group("/v1")

	articleGroup := v1.Group("/articles")
	{
		articleGroup.GET("", h.GetArticles)
		articleGroup.GET("/:id", h.GetArticleByID)
		articleGroup.POST("/add", h.AddArticle)
		articleGroup.DELETE("/delete/:id", h.DeleteArticle)
		articleGroup.PUT("/update/:id", h.UpdateArticle)
		//articleGroup.PUT("", h.AddArticle)
	}

	userGroup := v1.Group("/users")
	{
		userGroup.GET("", h.GetUsers)
		userGroup.GET("/:id", h.FindUserById)
		userGroup.DELETE("/delete/:id", h.DeleteUser)
		userGroup.PUT("/:id", h.UpdateUser)
		userGroup.POST("/add", h.SignUp)
	}

	profileGroup := v1.Group("/profiles")
	{
		profileGroup.GET("", h.ListProfiles)
		profileGroup.GET("/:id", h.GetProfileByID)
		profileGroup.POST("/add", h.CreateProfile)
		profileGroup.DELETE("/delete/:id", h.DeleteProfile)
		profileGroup.PUT("/update", h.UpdateProfile)
	}

	commentGroup := v1.Group("/comments")
	{
		commentGroup.POST("/add/:id", h.AddComment)
		//commentGroup.GET("", h.GetComments)
		//commentGroup.DELETE("/delete/:id", h.)
	}

	supportGroup := v1.Group("/supports")
	{
		supportGroup.GET("/:id", h.GetMessageById)
		supportGroup.GET("", h.GetAllMessages)
		supportGroup.DELETE("/delete/:id", h.DeleteMessage)
		supportGroup.POST("add", h.SendMessageToSupport)
	}

	readLaterGroup := v1.Group("/readLater")
	{
		readLaterGroup.POST("/add/:id", h.AddToReadLater)
		readLaterGroup.DELETE("/delete/:id", h.DeleteFromFavorites)
		readLaterGroup.GET("", h.GetArticlesFromFavorites)
	}
}
