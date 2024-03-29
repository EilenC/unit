package consistenthash

import (
	"fmt"
	"hash/crc32"
	"sort"
)

// Hash maps bytes to uint32
type Hash func(data []byte) uint32

// Map constains all hashed keys
type Map struct {
	hash     Hash
	replicas int            //虚拟节点数量
	keys     []int          //Sorted 哈希环
	hashMap  map[int]string //键是虚拟节点的哈希值,值是真实节点的名称
}

func New(replicas int, fn Hash) *Map {
	m := &Map{
		hash:     fn,
		replicas: replicas,
		hashMap:  make(map[int]string),
	}

	if m.hash == nil {
		m.hash = crc32.ChecksumIEEE //default hash algorithms
	}

	return m
}

// Add adds some keys to the hash.
func (m *Map) Add(keys ...string) {
	for _, key := range keys {
		for i := 0; i < m.replicas; i++ {
			hash := int(m.hash([]byte(fmt.Sprint(i) + key))) //calc hash
			m.keys = append(m.keys, hash)
			m.hashMap[hash] = key
		}
	}
	sort.Ints(m.keys) //sort keys
}

// Get gets the closest item in the hash to the provided key.
func (m *Map) Get(key string) string {
	if len(m.keys) == 0 {
		return ""
	}

	hash := int(m.hash([]byte(key)))
	idx := sort.Search(len(m.keys), func(i int) bool {
		return m.keys[i] >= hash
	})

	return m.hashMap[m.keys[idx%len(m.keys)]]
}

// Remove hash delete some keys
func (m *Map) Remove(key string) {
	for i := 0; i < m.replicas; i++ {
		hash := int(m.hash([]byte(fmt.Sprint(i) + key))) //计算key的hash
		idx := sort.Search(len(m.keys), func(i int) bool {
			return m.keys[i] >= hash
		})
		m.keys = append(m.keys[:idx], m.keys[idx+1:]...) //删除环上所有虚拟节点(保持顺序)
		delete(m.hashMap, hash)                          //删除节点
	}
}
