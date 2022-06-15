package consistenthash

import (
	"hash/crc32"
	"sort"
	"strconv"
)

type (
	hashFunc func(val []byte) uint32

	consistentHash struct {
		hashFn   hashFunc
		nodes    []uint32
		nodesMap map[uint32]string
		replicas int64 // number of virtual nodes
	}
)

// New create a new consistentHash, replicas means number of virtual nodes,
// if hashFn is nil, crc32 is used by default.
func New(replicas int64, hashFn hashFunc) *consistentHash {
	c := &consistentHash{replicas: replicas, nodesMap: make(map[uint32]string)}
	if hashFn == nil {
		c.hashFn = crc32.ChecksumIEEE
	} else {
		c.hashFn = hashFn
	}
	return c
}

// Add will add nodes to consistentHash
func (c *consistentHash) Add(nodes ...string) {
	for _, node := range nodes {
		// create virtual nodes
		for i := 0; i < int(c.replicas); i++ {
			hash := c.hashFn([]byte(strconv.Itoa(i) + node))
			c.nodes = append(c.nodes, hash)
			c.nodesMap[hash] = node
		}
	}
	// sort the nodes because Get() will look for it clockwise.
	sort.Slice(c.nodes, func(i, j int) bool { return c.nodes[i] < c.nodes[j] })
}

// Get gets the closest item in the hash to the provided key.
func (c *consistentHash) Get(key string) string {
	if c.isEmpty() {
		return ""
	}
	hash := c.hashFn([]byte(key))
	// find it clockwise.
	index := sort.Search(len(c.nodes), func(i int) bool { return c.nodes[i] >= hash })
	if index == len(c.nodes) {
		// means we have cycled back to the first replica.
		index = 0
	}
	return c.nodesMap[c.nodes[index]]
}

func (c *consistentHash) isEmpty() bool {
	return len(c.nodes) == 0
}
