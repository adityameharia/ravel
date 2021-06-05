package main

import (
	"github.com/buraksezer/consistent"
	"github.com/cespare/xxhash"
	"log"
	"math/rand"
	"sync"
	"time"
)

// clusterID is a string that is used to communicate the ID of a cluster. It implements the consistent.Member interface
type clusterID string
func (c clusterID) String() string {
	return string(c)
}

// hash implements the consistent.Hasher interface
type hash struct {}
func (h hash) Sum64(data []byte) uint64 {
	return xxhash.Sum64(data)
}

// ketSet implements a simple Set to store unique values of the keys
type keySet struct {
	m map[string]struct{}
}

func newKeySet() keySet {
	var k keySet
	k.m = make(map[string]struct{})
	return k
}

func (k keySet) Insert(key []byte) {
	k.m[string(key)] = struct{}{}
}

func (k keySet) Delete(key []byte) {
	delete(k.m, string(key))
}

func (k keySet) All() [][]byte {
	var all [][]byte
	for key := range k.m {
		all = append(all, []byte(key))
	}

	return all
}

// RavelConsistentHash is the main entity that implements the logic for sharding and data relocation
type RavelConsistentHash struct {
	mutex sync.Mutex
	config consistent.Config
	PartitionKeyMap map[int]keySet    // PartitionID -> []keySet
	PartitionOwners map[int]clusterID // PartitionID -> clusterID
	HashRing *consistent.Consistent
}

// Init initialises a RavelConsistentHash object
func (rch *RavelConsistentHash) Init(partitionCount int, replicationFactor int, load float64) {
	rch.mutex.Lock()
	defer rch.mutex.Unlock()
	rand.Seed(time.Now().UTC().UnixNano())

	rch.PartitionKeyMap = make(map[int]keySet)
	rch.PartitionOwners = make(map[int]clusterID)

	for i := 0; i<partitionCount; i++ {
		rch.PartitionOwners[i] = ""
		rch.PartitionKeyMap[i] = newKeySet()
	}

	rch.config = consistent.Config{
		PartitionCount: partitionCount,
		ReplicationFactor: replicationFactor,
		Load: load,
		Hasher: hash{},
	}

	rch.HashRing = consistent.New(nil, rch.config)
}

// AddCluster adds a new cluster to the ring, as a result some partitions are relocated to this new cluster,
// the keys in the relocated partition are looked up in the RavelConsistentHash.PartitionKeyMap and are moved
// to the new cluster
func (rch *RavelConsistentHash) AddCluster(clusterName clusterID) {
	rch.mutex.Lock()
	defer rch.mutex.Unlock()

	rch.HashRing.Add(clusterName)
	for partID, owner := range rch.PartitionOwners {
		newOwner := rch.HashRing.GetPartitionOwner(partID)
		if newOwner != owner {
			// relocate this partID to currentOwner
			keys := rch.PartitionKeyMap[partID].All()

			for i:=0; i<len(keys); i++ {
				val, err := clusterAdminGRPCServer.ReadKey(keys[i], owner.String())
				if err != nil {
					log.Println(err)
				}

				err = clusterAdminGRPCServer.DeleteKey(keys[i], owner.String())
				if err != nil {
					log.Println(err)
				}

				err = clusterAdminGRPCServer.WriteKeyValue(keys[i], val, newOwner.String())
				if err != nil {
					log.Println(err)
				}
			}

			rch.PartitionOwners[partID] = clusterID(newOwner.String())
		}
	}
}

// LocateKey returns the cluster for a given key
func (rch *RavelConsistentHash) LocateKey(key []byte) consistent.Member {
	rch.mutex.Lock()
	defer rch.mutex.Unlock()

	partID := rch.HashRing.FindPartitionID(key)
	rch.PartitionKeyMap[partID].Insert(key)

	return rch.HashRing.LocateKey(key)
}

