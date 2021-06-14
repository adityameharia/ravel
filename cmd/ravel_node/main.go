package main

import (
	"log"
	"os"

	"github.com/adityameharia/ravel/db"
	"github.com/urfave/cli/v2"

	"github.com/adityameharia/ravel/RavelClusterAdminPB"
	"github.com/google/uuid"
)

// Config is the struct containing the configuration details of the node
type Config struct {
	ClusterID        string // ClusterID is ID of th cluster the node is a part of
	NodeID           string // NodeID is the nodes unique ID
	StorageDir       string // StorageDir is the Data Directory for Raft
	GRPCAddr         string // GRPCAddr is the Address (with port) at which gRPC server is started
	RaftInternalAddr string // RaftInternalAddr is the Raft internal communication address with port
	AdminGRPCAddr    string // AdminGRPCAddr is the address at which the cluster admin gRPC server is hosted
	IsLeader         bool   // IsLeader is a bool defining whether the node is a leader or not
}

var nodeConfig Config
var adminClient RavelClusterAdminPB.RavelClusterAdminClient
var conf db.RavelDatabase
var dirname string
var yamlFile string

func init() {
	nodeConfig.NodeID = uuid.New().String()
	var err error
	dirname, err = os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	// nodeConfig.GRPCAddr = "0.0.0.0:50000"
	// nodeConfig.RaftInternalAddr = "localhost:60000"
	// nodeConfig.StorageDir = "/tmp/ravel_node"

	// flag.StringVar(&nodeConfig.StorageDir, "storageDir", "", "Storage Dir")
	// flag.StringVar(&nodeConfig.GRPCAddr, "gRPCAddr", "", "GRPC Addr of this node")
	// flag.StringVar(&nodeConfig.RaftInternalAddr, "raftAddr", "", "Raft Internal address for this node")
	// flag.StringVar(&nodeConfig.AdminGRPCAddr, "adminRPCAddr", "", "GRPC address of the cluster admin")
	// flag.BoolVar(&nodeConfig.IsLeader, "leader", false, "Register this node as a new leader or not")
}

func main() {
	app := &cli.App{}
	app.Name = "Ravel Replica"
	app.Usage = "Manage a Ravel replica server"
	app.Commands = []*cli.Command{
		{
			Name:  "start",
			Usage: "Starts a replica server",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:        "storagedir",
					Usage:       "Storage Dir",
					Value:       dirname + "/ravel_replica",
					Aliases:     []string{"s"},
					Destination: &nodeConfig.StorageDir},
				&cli.StringFlag{
					Name:        "grpcaddr",
					Usage:       "GRPC Addr of this replica",
					Value:       "localhost:50000",
					Aliases:     []string{"g"},
					Destination: &nodeConfig.GRPCAddr,
				},
				&cli.StringFlag{
					Name:        "raftaddr",
					Usage:       "Raft Internal address for this replica",
					Value:       "localhost:60000",
					Aliases:     []string{"r"},
					Destination: &nodeConfig.RaftInternalAddr,
				},
				&cli.StringFlag{
					Name:        "adminrpcaddr",
					Usage:       "GRPC address of the cluster admin",
					Value:       "localhost:42000",
					Aliases:     []string{"a"},
					Destination: &nodeConfig.AdminGRPCAddr,
				},
				&cli.StringFlag{
					Name:        "yaml",
					Usage:       "yaml file containing the config",
					Value:       "",
					Aliases:     []string{"y"},
					Destination: &yamlFile,
				},
				&cli.BoolFlag{
					Name:        "leader",
					Usage:       "Register this node as a new leader or not",
					Value:       false,
					Aliases:     []string{"l"},
					Destination: &nodeConfig.IsLeader,
				},
			},
			Action: func(c *cli.Context) error {
				setUpConf()
				startReplica()
				return nil
			},
		},
		{
			Name:  "kill",
			Usage: "Removes and deletes all the data in the cluster",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:        "storagedir",
					Usage:       "Storage Dir",
					Value:       dirname + "/ravel_replica",
					Required:    true,
					Aliases:     []string{"s"},
					Destination: &nodeConfig.StorageDir},
			},
			Action: func(c *cli.Context) error {
				setUpConf()
				killCluster()
				return nil
			},
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func setUpConf() {
	err := conf.Init(nodeConfig.StorageDir + "/config")
	if err != nil {
		log.Println(err)
		log.Fatal("Conf: Unable to Setup Database")
	}
	cID, err := conf.Read([]byte("clusterID"))
	if err == nil {
		nodeConfig.ClusterID = string(cID)
	}
	nID, err := conf.Read([]byte("nodeID"))
	if err == nil {
		nodeConfig.NodeID = string(nID)
	}
}
