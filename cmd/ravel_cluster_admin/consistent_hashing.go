package main

import (
	"encoding/json"
	"github.com/adityameharia/ravel/db"
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
type hash struct{}

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

func (k keySet) AllStrings() []string {
	var all []string
	for key := range k.m {
		all = append(all, key)
	}

	return all
}

// RavelConsistentHash is the main entity that implements the logic for sharding and data relocation
type RavelConsistentHash struct {
	mutex           sync.Mutex
	config          consistent.Config
	PartitionKeyMap map[uint64]keySet    // PartitionID -> []keySet
	PartitionOwners map[uint64]clusterID // PartitionID -> clusterID
	HashRing        *consistent.Consistent
}

// Init initialises a RavelConsistentHash object
func (rch *RavelConsistentHash) Init(partitionCount int, replicationFactor int, load float64) {
	rch.mutex.Lock()
	defer rch.mutex.Unlock()
	rand.Seed(time.Now().UTC().UnixNano())

	diskPartitionKeyMap, err := ReadPartitionKeyMapFromDisk(RavelClusterAdminBackupPath)
	if err != nil {
		log.Fatal("Error in rch.Init:", err.Error())
	}

	diskPartitionOwnersMap, err := ReadPartitionOwnersFromDisk(RavelClusterAdminBackupPath)
	if err != nil {
		log.Fatal("Error in rch.Init:", err.Error())
	}

	rch.PartitionKeyMap = diskPartitionKeyMap
	rch.PartitionOwners = diskPartitionOwnersMap

	if len(rch.PartitionOwners) == 0 && len(rch.PartitionKeyMap) == 0 {
		for i := 0; i < partitionCount; i++ {
			rch.PartitionOwners[uint64(i)] = ""
			rch.PartitionKeyMap[uint64(i)] = newKeySet()
		}
	}

	rch.config = consistent.Config{
		PartitionCount:    partitionCount,
		ReplicationFactor: replicationFactor,
		Load:              load,
		Hasher:            hash{},
	}

	rch.HashRing = consistent.New(nil, rch.config)
}

// Reset resets the RavelConsistentHash object to its initial state
func (rch *RavelConsistentHash) Reset(partitionCount int, replicationFactor int, load float64) {
	rch.mutex.Lock()
	defer rch.mutex.Unlock()
	rand.Seed(time.Now().UTC().UnixNano())

	rch.PartitionKeyMap = make(map[uint64]keySet)
	rch.PartitionOwners = make(map[uint64]clusterID)

	for i := 0; i < partitionCount; i++ {
		rch.PartitionOwners[uint64(i)] = ""
		rch.PartitionKeyMap[uint64(i)] = newKeySet()
	}

	rch.config = consistent.Config{
		PartitionCount:    partitionCount,
		ReplicationFactor: replicationFactor,
		Load:              load,
		Hasher:            hash{},
	}

	rch.HashRing = consistent.New(nil, rch.config)
}

// BackupToDisk writes the RavelConsistentHash.PartitionKeyMap and RavelConsistentHash.PartitionOwners maps to disk using BadgerDB
func (rch *RavelConsistentHash) BackupToDisk(badgerPath string) error {
	log.Println("Running Backup")
	var backupDB db.RavelDatabase

	err := backupDB.Init(badgerPath + "/partition_owners")
	if err != nil {
		return err
	}
	for partID, cluster := range rch.PartitionOwners {
		err = backupDB.Write(uint64ToBytes(partID), []byte(cluster.String()))
	}

	backupDB.Close()

	err = backupDB.Init(badgerPath + "/partition_keymap")
	if err != nil {
		return err
	}

	type kSetJSON struct {
		Keys []string `json:"keys"`
	}
	for partID, kSet := range rch.PartitionKeyMap {
		kSetJSONBytes, err := json.Marshal(kSetJSON{
			Keys: kSet.AllStrings(),
		})

		if err != nil {
			log.Println("Error in RavelConsistentHash.BackupOnDisk:", err.Error())
		}

		err = backupDB.Write(uint64ToBytes(partID), kSetJSONBytes)
	}

	backupDB.Close()
	return nil
}

// AddCluster adds a new cluster to the ring, as a result some partitions are relocated to this new cluster,
// the keys in the relocated partition are looked up in the RavelConsistentHash.PartitionKeyMap and are moved
// to the new cluster
func (rch *RavelConsistentHash) AddCluster(clusterName clusterID) {
	log.Println("Len Partition")
	log.Println("Adding Cluster:", clusterName)
	rch.mutex.Lock()
	defer rch.mutex.Unlock()

	rch.HashRing.Add(clusterName)
	rch.relocatePartitions() // dont do this when this is the very first one

	err := rch.BackupToDisk(RavelClusterAdminBackupPath)
	if err != nil {
		log.Println("Error in Backing Up to Disk:", err.Error())
	}
}

// DeleteCluster deletes a cluster from the owners map
func (rch *RavelConsistentHash) DeleteCluster(clusterName clusterID) {
	log.Println("Removing Cluster:", clusterName)
	rch.mutex.Lock()
	defer rch.mutex.Unlock()

	rch.HashRing.Remove(clusterName.String())
	rch.relocatePartitions()

	err := rch.BackupToDisk(RavelClusterAdminBackupPath)
	if err != nil {
		log.Println("Error in Backing Up to Disk:", err.Error())
	}
}

// LocateKey returns the cluster for a given key
func (rch *RavelConsistentHash) LocateKey(key []byte) consistent.Member {
	rch.mutex.Lock()
	defer rch.mutex.Unlock()

	partID := rch.HashRing.FindPartitionID(key)
	rch.PartitionKeyMap[uint64(partID)].Insert(key)

	err := rch.BackupToDisk(RavelClusterAdminBackupPath)
	if err != nil {
		log.Println("Error in Backing Up to Disk:", err.Error())
	}

	return rch.HashRing.LocateKey(key)
}

// relocatePartitions checks for owner changes and then relocates the keys in that partition to the new owner
func (rch *RavelConsistentHash) relocatePartitions() {
	log.Println("Relocating Partitions")
	for partID, owner := range rch.PartitionOwners {
		newOwner := rch.HashRing.GetPartitionOwner(int(partID))
		if newOwner != owner {
			// relocate this partID to newOwner
			keys := rch.PartitionKeyMap[partID].All()

			for i := 0; i < len(keys); i++ {
				log.Printf("Relocating key: %v from cluster: %v to cluster: %v\n", string(keys[i]), owner.String(), newOwner.String())
				val, err := clusterAdminGRPCServer.ReadKeyAndDelete(keys[i], owner.String())
				if err != nil {
					log.Println(err)
				}

				err = clusterAdminGRPCServer.WriteKeyValue(keys[i], val, newOwner.String())
				if err != nil {
					log.Println("Yo:", err)
				}
			}

			rch.PartitionOwners[partID] = clusterID(newOwner.String())
		}
	}
}