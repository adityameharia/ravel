package main

import (
	"encoding/binary"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"math"
)

type fileType string

type keyType string

const (
	stringKeyType keyType = "string"
	floatKeyType  keyType = "float"
	jsonKeyType   keyType = "json"
	fileKeyType   keyType = "file"
	imageKeyType  keyType = "image"
)

type ClusterAdminHTTPServer struct {
	Router     *gin.Engine
	KeyTypeMap map[string]keyType
}

func NewClusterAdminHTTPServer() *ClusterAdminHTTPServer {
	var server ClusterAdminHTTPServer
	server.Router = gin.Default()
	server.setupPaths()
	server.KeyTypeMap = make(map[string]keyType)
	return &server
}

func float64ToByte(f float64) []byte {
	var buf [8]byte
	binary.BigEndian.PutUint64(buf[:], math.Float64bits(f))
	return buf[:]
}

func byteTofloat64(bytes []byte) float64 {
	bits := binary.BigEndian.Uint64(bytes)
	float := math.Float64frombits(bits)
	return float
}

func (s *ClusterAdminHTTPServer) setupPaths() {
	type putRequest struct {
		Key string      `json:"key"`
		Val interface{} `json:"val"`
	}

	type getRequest struct {
		Key string `json:"key"`
	}

	s.Router.GET("/", func(c *gin.Context) {
		c.JSON(200, "HTTP Server for Ravel Cluster Admin")
	})

	s.Router.POST("/get", func(c *gin.Context) {
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

		switch clusterAdminHTTPServer.KeyTypeMap[req.Key] {
		case "float":
			c.JSON(200, gin.H{"key": req.Key, "val": byteTofloat64(val)})
		case "string":
			c.JSON(200, gin.H{"key": req.Key, "val": string(val)})
		case "json":
			var r interface{}
			err := json.Unmarshal(val, &r)
			if err != nil {
				c.JSON(500, gin.H{"error": err.Error()})
			}

			c.JSON(200, gin.H{"key": req.Key, "val": r})
		}
	})

	s.Router.POST("/put", func(c *gin.Context) {
		var req putRequest
		if err := c.Bind(&req); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		clusterID := consistentHash.LocateKey([]byte(req.Key))

		switch req.Val.(type) {
		case float64:
			clusterAdminHTTPServer.KeyTypeMap[req.Key] = "float"
			v := float64ToByte(req.Val.(float64))
			err := clusterAdminGRPCServer.WriteKeyValue([]byte(req.Key), v, clusterID.String())
			if err != nil {
				c.JSON(500, gin.H{"error": err.Error()})
				return
			}

			c.JSON(200, gin.H{"msg": "ok"})
		case string:
			clusterAdminHTTPServer.KeyTypeMap[req.Key] = "string"
			err := clusterAdminGRPCServer.WriteKeyValue([]byte(req.Key), []byte(req.Val.(string)), clusterID.String())
			if err != nil {
				c.JSON(500, gin.H{"error": err.Error()})
				return
			}

			c.JSON(200, gin.H{"msg": "ok"})
		case map[string]interface{}: // json object
			clusterAdminHTTPServer.KeyTypeMap[req.Key] = "json"
			jsonBytes, err := json.Marshal(req.Val)
			err = clusterAdminGRPCServer.WriteKeyValue([]byte(req.Key), jsonBytes, clusterID.String())
			if err != nil {
				c.JSON(500, gin.H{"error": err.Error()})
				return
			}

			c.JSON(200, gin.H{"msg": "ok"})
		}
	})
}
