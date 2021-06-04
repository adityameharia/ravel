package main

import (
	"github.com/buraksezer/consistent"
	"github.com/cespare/xxhash"
	"math/rand"
	"time"
)

type cluster string
func (c cluster) String() string {
	return string(c)
}

type hasher struct {}
func (h hasher) Sum64(data []byte) uint64 {
	return xxhash.Sum64(data)
}

type key []byte
type RavelConsistentHash struct {
	config consistent.Config
	PartitionKeyMap map[int][]key // PartitionID -> []key
	PartitionOwners map[int]cluster // PartitionID -> cluster
	HashRing *consistent.Consistent
}

func (rch *RavelConsistentHash) Init(partitionCount int, replicationFactor int, load float64) {
	rand.Seed(time.Now().UTC().UnixNano())

	rch.PartitionKeyMap = make(map[int][]key)
	rch.PartitionOwners = make(map[int]cluster)

	for i := 0; i<partitionCount; i++ {
		rch.PartitionOwners[i] = ""
	}

	rch.config = consistent.Config{
		PartitionCount: partitionCount,
		ReplicationFactor: replicationFactor,
		Load: load,
		Hasher: hasher{},
	}

	rch.HashRing = consistent.New(nil, rch.config)
}

func (rch *RavelConsistentHash) AddCluster(clusterName cluster) {
	rch.HashRing.Add(clusterName)
	for partID, owner := range rch.PartitionOwners {
		currentOwner := rch.HashRing.GetPartitionOwner(partID)
		if currentOwner != owner {
			// relocate this partID to currentOwner
			rch.PartitionOwners[partID] = cluster(currentOwner.String())
		}
	}
}

func (rch *RavelConsistentHash) LocateKey(k key) consistent.Member {
	partID := rch.HashRing.FindPartitionID(k)
	rch.PartitionKeyMap[partID] = append(rch.PartitionKeyMap[partID], k)
	return rch.HashRing.LocateKey(k)
}

