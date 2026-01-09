package api

import (
	"github.com/gin-gonic/gin"
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
		c.JSON(200, gin.H{
			"status": "up",
		})
	})

	api := s.router.Group("/api/v1")
	{
		api.GET("/plugins", s.listPlugins)
		api.POST("/scan", s.startScan)
	}
}

func (s *Server) listPlugins(c *gin.Context) {
	// Placeholder for listing plugins
	c.JSON(200, gin.H{"plugins": []string{"twitter", "youtube"}})
}

func (s *Server) startScan(c *gin.Context) {
	// Placeholder for starting a scan
	c.JSON(200, gin.H{"message": "scan started"})
}

func (s *Server) Run(addr string) error {
	return s.router.Run(addr)
}
