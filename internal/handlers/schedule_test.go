package handlers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/example/stratus-mesh/internal/handlers"
	"github.com/example/stratus-mesh/internal/models"
	"github.com/example/stratus-mesh/internal/services"
)

func setupScheduleServer(t *testing.T) (*gin.Engine, *gorm.DB) {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	require.NoError(t, err)
	err = db.AutoMigrate(&models.Node{}, &models.Schedule{})
	require.NoError(t, err)

	// seed a node
	node := models.Node{Name: "seed", Address: "http://127.0.0.1:9001"}
	require.NoError(t, db.Create(&node).Error)

	discovery := services.NewDiscovery()
	discovery.Register(node)

	h := handlers.NewHandler(db, discovery)
	r := gin.New()
	a := r.Group("/api/v1")
	{
		s := a.Group("/schedules")
		{
			s.POST("/", h.CreateSchedule)
			s.GET("/", h.ListSchedules)
		}
	}
	return r, db
}

func TestCreateSchedule(t *testing.T) {
	r, db := setupScheduleServer(t)
	var node models.Node
	require.NoError(t, db.First(&node).Error)

	w := httptest.NewRecorder()
	body := map[string]interface{}{"node_id": node.ID, "task": "do-work", "cron": "@every 5m"}
	b, _ := json.Marshal(body)
	req, _ := http.NewRequest("POST", "/api/v1/schedules/", bytes.NewReader(b))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)
	require.Equal(t, http.StatusCreated, w.Code)

	w2 := httptest.NewRecorder()
	req2, _ := http.NewRequest("GET", "/api/v1/schedules/", nil)
	r.ServeHTTP(w2, req2)
	require.Equal(t, http.StatusOK, w2.Code)
	var scheds []models.Schedule
	require.NoError(t, json.Unmarshal(w2.Body.Bytes(), &scheds))
	require.Len(t, scheds, 1)
	require.Equal(t, "do-work", scheds[0].Task)
}
