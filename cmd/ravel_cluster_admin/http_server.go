package main

import (
	"encoding/binary"
	"encoding/json"
	"math"
	"strconv"

	"github.com/gin-gonic/gin"
)

// float64ToByte converts a float64 to a []byte
func float64ToByte(f float64) []byte {
	var buf [8]byte
	binary.BigEndian.PutUint64(buf[:], math.Float64bits(f))
	return buf[:]
}

// byteToFloat64 converts a []byte to float64
func byteToFloat64(bytes []byte) float64 {
	bits := binary.BigEndian.Uint64(bytes)
	float := math.Float64frombits(bits)
	return float
}

// keyType represents the data type of the value of a key
type keyType string

// ClusterAdminHTTPServer is the entity that represents the HTTP server on the Cluster Admin
type ClusterAdminHTTPServer struct {
	Router     *gin.Engine
	KeyTypeMap map[string]keyType
}

// NewClusterAdminHTTPServer constructs and returns a ClusterAdminHTTPServer object
func NewClusterAdminHTTPServer() *ClusterAdminHTTPServer {
	var server ClusterAdminHTTPServer
	server.Router = gin.Default()
	server.setupPaths()
	server.KeyTypeMap = make(map[string]keyType)
	return &server
}

// setupPaths sets up HTTP endpoints for ClusterAdminHTTPServer.Router
func (s *ClusterAdminHTTPServer) setupPaths() {
	// Data for a "put" request
	type putRequest struct {
		Key string      `json:"key"`
		Val interface{} `json:"val"`
	}

	// Data for a "get" request
	type getRequest struct {
		Key string `json:"key"`
	}

	s.Router.GET("/", func(c *gin.Context) {
		c.JSON(200, "HTTP Server for Ravel Cluster Admin")
	})

	// /get reads the key from the request, locates the cluster with that key
	// reads the data from that cluster, decodes it into the appropriate type and returns it
	s.Router.POST("/get", func(c *gin.Context) {
		var req getRequest
		if err := c.Bind(&req); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		// get cluster id -> read key and value from that cluster
		clusterID := consistentHash.LocateKey([]byte(req.Key))
		val, err := clusterAdminGRPCServer.ReadKey([]byte(req.Key), clusterID.String())
		if err != nil {
			c.JSON(404, gin.H{"error": err.Error()})
			return
		}

		// check for data type of value and send response accordingly
		switch clusterAdminHTTPServer.KeyTypeMap[req.Key] {
		case "float":
			c.JSON(200, gin.H{"key": req.Key, "val": byteToFloat64(val)})
		case "string":
			c.JSON(200, gin.H{"key": req.Key, "val": string(val)})
		case "json":
			var r interface{}
			err := json.Unmarshal(val, &r)
			if err != nil {
				c.JSON(500, gin.H{"error": err.Error()})
			}

			c.JSON(200, gin.H{"key": req.Key, "val": r})
		case "bool":
			boolValue, err := strconv.ParseBool(string(val))
			if err != nil {
				c.JSON(500, gin.H{"error": err.Error()})
			}

			c.JSON(200, gin.H{"key": req.Key, "val": boolValue})
		}
	})

	// /put reads the key and value from the request, updates the type map, locates the cluster for it
	// and writes it to the leader node of that cluster
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
		case bool: // convert bool to string and write that as a []byte
			clusterAdminHTTPServer.KeyTypeMap[req.Key] = "bool"
			boolToString := strconv.FormatBool(req.Val.(bool))
			err := clusterAdminGRPCServer.WriteKeyValue([]byte(req.Key), []byte(boolToString), clusterID.String())
			if err != nil {
				c.JSON(500, gin.H{"error": err.Error()})
				return
			}

			c.JSON(200, gin.H{"msg": "ok"})
		}
	})
}
