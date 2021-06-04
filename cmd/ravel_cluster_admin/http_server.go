package main

import (
	"encoding/binary"
	"github.com/gin-gonic/gin"
	"math"
)

type ClusterAdminHTTPServer struct {
	Router *gin.Engine
}

func NewClusterAdminHTTPServer() *ClusterAdminHTTPServer {
	var server ClusterAdminHTTPServer
	server.Router = gin.Default()
	server.setupPaths()
	return &server
}

func float64ToByte(f float64) []byte {
	var buf [8]byte
	binary.BigEndian.PutUint64(buf[:], math.Float64bits(f))
	return buf[:]
}

func (s *ClusterAdminHTTPServer) setupPaths() {
	type putRequest struct {
		Key string `json:"key"`
		Val interface{} `json:"val"`
	}

	type getRequest struct {
		Key string `json:"key"`
	}

	s.Router.GET("/", func(c *gin.Context) {
		c.JSON(200, "HTTP Server for Ravel Cluster Admin")
	})

	s.Router.POST("/get", func(c *gin.Context){
		var req getRequest
		if err := c.Bind(&req); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		clusterID := consistentHash.LocateKey([]byte(req.Key))
		val, err := clusterAdminGRPCServer.ReadKey([]byte(req.Key), clusterID.String())
		if err != nil {
			c.JSON(404, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, gin.H{"key": req.Key, "msg": string(val)})
	})

	s.Router.POST("/put", func(c *gin.Context){
		var req putRequest
		if err := c.Bind(&req); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		clusterID := consistentHash.LocateKey([]byte(req.Key))

		switch req.Val.(type) {
		case float64:
			v := float64ToByte(req.Val.(float64))
			err := clusterAdminGRPCServer.WriteKeyValue([]byte(req.Key), v ,clusterID.String())
			if err != nil {
				c.JSON(500, gin.H{"error": err.Error()})
				return
			}
		case string:
			err := clusterAdminGRPCServer.WriteKeyValue([]byte(req.Key), []byte(req.Val.(string)) ,clusterID.String())
			if err != nil {
				c.JSON(500, gin.H{"error": err.Error()})
				return
			}
		case map[string]interface{}: // js object
			c.JSON(200, gin.H{"msg": "yet to implement"})
			return
		}

		c.JSON(200, gin.H{"msg": "ok"})
	})
}