package main

import (
	"github.com/gin-gonic/gin"
)

type ClusterAdminHTTPServer struct {
	Router *gin.Engine
}

func NewClusterAdminHTTPServer() (*ClusterAdminHTTPServer, error) {
	var server ClusterAdminHTTPServer
	server.Router = gin.Default()
	server.setupPaths()
	return &server, nil
}

func (s *ClusterAdminHTTPServer) setupPaths() {
	s.Router.GET("/", func(c *gin.Context) {
		c.JSON(200, "HTTP Server for Ravel Cluster Admin")
	})
}