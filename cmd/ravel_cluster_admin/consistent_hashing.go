package main

import (
	"github.com/buraksezer/consistent"
	"github.com/cespare/xxhash"
	"log"
	"math/rand"
	"sync"
	"time"
)

type clusterID string
func (c clusterID) String() string {
	return string(c)
}

type hash struct {}
func (h hash) Sum64(data []byte) uint64 {
	return xxhash.Sum64(data)
}

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

type RavelConsistentHash struct {
	mutex sync.Mutex
	config consistent.Config
	PartitionKeyMap map[int]keySet    // PartitionID -> []keySet
	PartitionOwners map[int]clusterID // PartitionID -> clusterID
	HashRing *consistent.Consistent
}

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

func (rch *RavelConsistentHash) AddCluster(clusterName clusterID) {
	log.Println("AddCluster: ")
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

func (rch *RavelConsistentHash) LocateKey(key []byte) consistent.Member {
	log.Println("LocateKey: ")
	rch.mutex.Lock()
	defer rch.mutex.Unlock()

	partID := rch.HashRing.FindPartitionID(key)
	rch.PartitionKeyMap[partID].Insert(key)

	return rch.HashRing.LocateKey(key)
}

