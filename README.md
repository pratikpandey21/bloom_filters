# Bloom Filter Implementation in Go ⚡️

Bloom filters are a probabilistic data structure that are used to test whether an element is a member of a set. This project includes a simple implementation of Bloom Filters in Go.

## Features 📋
- Quick to add elements & check membership in the Bloom Filter.
- Memory-efficient representation of a set (at the cost of a small false positive rate).
- The false positive rate decreases as we increase the filter's size or the number of hash functions.

## Setup 🚀
First, make sure you have Go installed on your machine. If you do not have it installed, you can download it from the official website [here](https://golang.org/dl/).

To experiment with the Bloom Filter, clone the project.

```shell
git clone  https://github.com/pratikpandey21/bloom_filters.git
```

Understand how the bloom filters accuracy changes with, check out the tests and run them to see the results -

```shell
go test ./...
```

## Usage 💻

You can use the bloom filter in your own code.

```shell
import "github.com/pratikpandey21/bloom_filters"

bf := NewBloomFilter(size, hashCount)
// Add an element
bf.Add("element")
// Check membership
exists := bf.Contains("element")
fmt.Println(exists) 
```

## Contributing 🤝
Contributions, issues, and feature requests are welcome! See [CONTRIBUTING.md](CONTRIBUTING.md) for more information.

## License 🏷️
This project is licensed under the terms of the MIT license. See [LICENSE](LICENSE) for more details.
