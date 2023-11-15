package handler

import (
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/gin-gonic/gin"
	"github.com/hosseinmirzapur/elk-example/db"
	"github.com/rs/zerolog"
)

type Handler struct {
	DB       db.Database
	Logger   zerolog.Logger
	ESClient *elasticsearch.Client
}

func New(database db.Database, esClient *elasticsearch.Client, logger zerolog.Logger) *Handler {
	return &Handler{
		DB:       database,
		ESClient: esClient,
		Logger:   logger,
	}
}

func (h *Handler) Register(group *gin.RouterGroup) {
	// // including params
	// group.GET("/posts/:id", h.GetPost)
	// group.PATCH("/posts/:id", h.UpdatePost)
	// group.DELETE("/posts/:id", h.DeletePost)

	// // no params
	// group.GET("/posts", h.GetPosts)
	group.POST("/posts", h.CreatePost)

	group.GET("/search", h.SearchPosts)
}
