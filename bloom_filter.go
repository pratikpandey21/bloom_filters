package bloom_filters

import (
	"github.com/spaolacci/murmur3"
	"math/rand"
)

type BloomFilter struct {
	bitArray  []bool
	size      int
	hashCount int
}

func NewBloomFilter(size, hashCount int) *BloomFilter {
	return &BloomFilter{bitArray: make([]bool, size), size: size, hashCount: hashCount}
}

func (bf *BloomFilter) Add(str string) {
	for i := 0; i < bf.hashCount; i++ {
		hash := hash(str, i) % bf.size
		bf.bitArray[hash] = true
	}
}

func (bf *BloomFilter) Contains(str string) bool {
	for i := 0; i < bf.hashCount; i++ {
		hash := hash(str, i) % bf.size
		if !bf.bitArray[hash] {
			return false
		}
	}
	return true
}

func hash(str string, seed int) int {
	mmr := murmur3.New32WithSeed(uint32(rand.Uint32() * uint32(seed)))
	_, err := mmr.Write([]byte(str))
	if err != nil {
		return 0
	}
	return int(mmr.Sum32())
}
