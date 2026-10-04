package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/example/stratus-mesh/internal/models"
	"github.com/example/stratus-mesh/internal/services"
)

type Handler struct {
	db *gorm.DB
	discovery *services.Discovery
}

func NewHandler(db *gorm.DB, discovery *services.Discovery) *Handler {
	return &Handler{db: db, discovery: discovery}
}

// CreateNode registers a node and returns it
func (h *Handler) CreateNode(c *gin.Context) {
	var payload models.Node
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.db.Create(&payload).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// add to discovery for simulation
	h.discovery.Register(payload)
	c.JSON(http.StatusCreated, payload)
}

// ListNodes returns a list of nodes
func (h *Handler) ListNodes(c *gin.Context) {
	var nodes []models.Node
	if err := h.db.Find(&nodes).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, nodes)
}

// CreateSchedule creates a scheduled task for a node
func (h *Handler) CreateSchedule(c *gin.Context) {
	var payload models.Schedule
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.db.Create(&payload).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, payload)
}

// ListSchedules returns schedules
func (h *Handler) ListSchedules(c *gin.Context) {
	var schedules []models.Schedule
	if err := h.db.Preload("Node").Find(&schedules).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, schedules)
}
