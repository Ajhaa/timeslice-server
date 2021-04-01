package task

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterHandlers(group *gin.RouterGroup, s Service) {
	group.POST("/", s.post)
	group.GET("/", s.index)
}

type partialTask struct {
	Name     string `json:"name"`
	Duration uint   `json:"duration"`
}

func (s Service) post(ctx *gin.Context) {
	var task partialTask

	if err := ctx.ShouldBindJSON(&task); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdTask := s.Repo.create(ctx, task)

	ctx.JSON(200, createdTask)
}

func (s Service) index(ctx *gin.Context) {
	tasks := s.Repo.findAll(ctx)
	ctx.JSON(200, tasks)
}
