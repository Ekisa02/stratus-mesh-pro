package router

import (
	"github.com/gin-gonic/gin"
	"github.com/example/stratus-mesh/internal/handlers"
)

func NewRouter(h *handlers.Handler) *gin.Engine {
	r := gin.Default()
	a := r.Group("/api/v1")
	{
		n := a.Group("/nodes")
		{
			n.POST("/", h.CreateNode)
			n.GET("/", h.ListNodes)
		}
		 s := a.Group("/schedules")
		{
			s.POST("/", h.CreateSchedule)
			s.GET("/", h.ListSchedules)
		}
	}
	return r
}
