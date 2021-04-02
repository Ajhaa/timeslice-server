package task

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterHandlers(group *gin.RouterGroup, s Service) {
	r := resource{service: s}
	group.POST("/", r.post)
	group.GET("/", r.index)
}

type resource struct {
	service Service
}

func (r resource) post(ctx *gin.Context) {
	var task PartialTask

	if err := ctx.ShouldBindJSON(&task); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdTask := r.service.Create(ctx, task)

	ctx.JSON(200, createdTask)
}

func (r resource) index(ctx *gin.Context) {
	tasks := r.service.Query(ctx)
	ctx.JSON(200, tasks)
}
