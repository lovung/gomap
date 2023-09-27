# GoMap

[![Go Report Card](https://goreportcard.com/badge/github.com/lovung/gomap)](https://goreportcard.com/report/github.com/lovung/gomap)
[![GoDoc](https://godoc.org/github.com/lovung/gomap?status.svg)](https://pkg.go.dev/github.com/lovung/gomap)
[![License](https://img.shields.io/badge/License-MIT-blue.svg)](https://github.com/lovung/gomap/blob/main/LICENSE)


In Golang, we have the built-in `map` which is using hashmap to implement O(N). 

However, the performance of hashing algorithm is not go at all.

By creating this package, I am aiming to implement many types of Map which support 
more purposes but easily to change between the algorithm in the blackbox.

## Goals

- [x] Generics - Type Safe.
- [x] Support thread-safe.
- [x] Easy to switch map types.
- [x] Support Bloom filter if key's type is `integer`.
- [ ] Stat functions: hit-rate, size, time of operations.


## API Documentation
For detailed documentation, please refer to the [GoDoc](https://pkg.go.dev/github.com/loving/gomap) page.

## Contributing
Contributions are welcome! Please feel free to open issues or pull requests for bug fixes, improvements, or new features.

## License
This project is licensed under the MIT License - see the LICENSE file for details.
