package bloom_filters

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func BenchmarkBloomFilter_Add(b *testing.B) {
	bf := NewBloomFilter(10000, 7)

	testWithSameKey := randomString()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bf.Add(testWithSameKey)
	}
}

func BenchmarkBloomFilter_Contains(b *testing.B) {
	bf := NewBloomFilter(10000, 7)

	insertedStrings := make([]string, 10000)
	// Inserting random elements and storing them
	for i := 0; i < 10000; i++ {
		str := randomString()
		insertedStrings[i] = str
		bf.Add(str)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		bf.Contains(insertedStrings[(i*13)%10000])
	}
}

func TestBloomFilter(t *testing.T) {
	numInserts := 5000
	numTests := 10000

	for size := 30000; size <= 50000; size += 10000 {
		for hashCount := 4; hashCount <= 7; hashCount++ {
			startTime := time.Now() // Start time measurement

			fpRate, falsePositives := findFalsePositiveRate(size, hashCount, numInserts, numTests)

			// Report detailed stats
			fmt.Printf("Size: %d, Hash Count: %d, False Positives: %d, Total Checks: %d, False Positive Rate: %.2f%%, Duration: %s\n",
				size, hashCount, falsePositives, numTests, fpRate*100, time.Since(startTime))

			if fpRate > 0.02 {
				t.Errorf("False positive rate exceeds threshold for size %d and hash count %d", size, hashCount)
			}
		}
	}
}

// findFalsePositiveRate finds the false positive rate for a Bloom filter.
// It inserts numInserts random strings into the Bloom filter and then
// tests numTests random strings to see if they are false positives.
// It returns the false positive rate as a float64.
// Example usage:
// fpRate := findFalsePositiveRate(8000, 4, 1000, 10000)
func findFalsePositiveRate(size int, hashCount int, numInserts int, numTests int) (float64, int) {
	bf := NewBloomFilter(size, hashCount)

	insertedStrings := make([]string, numInserts)
	// Inserting random elements and storing them
	for i := 0; i < numInserts; i++ {
		str := randomString()
		insertedStrings[i] = str
		bf.Add(str)
	}

	// Checking for false positives
	falsePositives := 0
	for i := 0; i < numTests; i++ {
		testStr := randomString()

		// Ensure the random string was not one of the inserted strings
		if !contains(insertedStrings, testStr) && bf.Contains(testStr) {
			falsePositives++
		}
	}

	return float64(falsePositives) / float64(numTests), falsePositives
}

// contains checks if a slice contains a string
func contains(slice []string, str string) bool {
	for _, item := range slice {
		if item == str {
			return true
		}
	}
	return false
}

// randomString generates a random string.
func randomString() string {
	rand.Seed(time.Now().UnixNano())
	letterBytes := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, 10)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
