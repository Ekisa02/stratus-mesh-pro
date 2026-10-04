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

func setupTestServer(t *testing.T) (*gin.Engine, *gorm.DB) {
	// setup sqlite in-memory
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	require.NoError(t, err)
	err = db.AutoMigrate(&models.Node{}, &models.Schedule{})
	require.NoError(t, err)

	discovery := services.NewDiscovery()
	h := handlers.NewHandler(db, discovery)
	r := gin.New()
	a := r.Group("/api/v1")
	{
		n := a.Group("/nodes")
		{
			n.POST("/", h.CreateNode)
			n.GET("/", h.ListNodes)
		}
	}
	return r, db
}

func TestCreateAndListNode(t *testing.T) {
	r, _ := setupTestServer(t)
	w := httptest.NewRecorder()
	body := map[string]string{"name": "test-node", "address": "http://127.0.0.1:9000"}
	b, _ := json.Marshal(body)
	req, _ := http.NewRequest("POST", "/api/v1/nodes/", bytes.NewReader(b))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)
	require.Equal(t, http.StatusCreated, w.Code)

	// list
	w2 := httptest.NewRecorder()
	req2, _ := http.NewRequest("GET", "/api/v1/nodes/", nil)
	r.ServeHTTP(w2, req2)
	require.Equal(t, http.StatusOK, w2.Code)
	var nodes []models.Node
	err := json.Unmarshal(w2.Body.Bytes(), &nodes)
	require.NoError(t, err)
	require.Len(t, nodes, 1)
	require.Equal(t, "test-node", nodes[0].Name)
}
