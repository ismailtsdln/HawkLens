package api

import (
	"context"
	"io"

	"github.com/gin-gonic/gin"
	"github.com/ismailtsdln/HawkLens/internal/db"
	"github.com/ismailtsdln/HawkLens/internal/engine"
	"github.com/ismailtsdln/HawkLens/pkg/plugins"
)

type Server struct {
	router *gin.Engine
}

func NewServer() *Server {
	router := gin.Default()

	s := &Server{
		router: router,
	}

	s.setupRoutes()
	return s
}

func (s *Server) setupRoutes() {
	s.router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "up"})
	})

	api := s.router.Group("/api/v1")
	{
		api.GET("/plugins", s.listPlugins)
		api.GET("/results", s.getResults)
		api.GET("/scan-stream", s.streamScan)
	}
}

func (s *Server) listPlugins(c *gin.Context) {
	c.JSON(200, gin.H{"plugins": plugins.ListPlugins()})
}

func (s *Server) getResults(c *gin.Context) {
	platform := c.Query("platform")
	pg, err := db.NewPostgresDB("postgres://user:pass@localhost:5432/hawklens?sslmode=disable")
	if err != nil {
		c.JSON(500, gin.H{"error": "Database connection failed"})
		return
	}
	defer pg.Close()

	results, err := pg.ListResults(platform)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to fetch results"})
		return
	}
	c.JSON(200, results)
}

func (s *Server) streamScan(c *gin.Context) {
	query := c.Query("query")
	if query == "" {
		c.JSON(400, gin.H{"error": "query is required"})
		return
	}

	c.Writer.Header().Set("Content-Type", "text/event-stream")
	c.Writer.Header().Set("Cache-Control", "no-cache")
	c.Writer.Header().Set("Connection", "keep-alive")
	c.Writer.Header().Set("Transfer-Encoding", "chunked")

	pluginNames := plugins.ListPlugins()
	dispatcher := engine.NewDispatcher(5)

	// Create a context that is cancelled when the client disconnects
	ctx, cancel := context.WithCancel(c.Request.Context())
	defer cancel()

	dispatcher.Run(ctx)

	for _, name := range pluginNames {
		dispatcher.Submit(name, query)
	}

	c.Stream(func(w io.Writer) bool {
		for wrapper := range dispatcher.Results() {
			if wrapper.Error == nil {
				for _, res := range wrapper.Results {
					c.SSEvent("message", res)
				}
			}
		}
		return false // Close stream after dispatcher is done
	})

	dispatcher.Wait()
}

func (s *Server) Run(addr string) error {
	return s.router.Run(addr)
}
