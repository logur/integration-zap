# Logur integration for Uber's [Zap](https://github.com/uber-go/zap)

[![CircleCI](https://circleci.com/gh/logur/integration-zap.svg?style=svg)](https://circleci.com/gh/logur/integration-zap)
[![Coverage](https://gocover.io/_badge/logur.dev/integration/zap)](https://gocover.io/logur.dev/integration/zap)
[![Go Report Card](https://goreportcard.com/badge/logur.dev/integration/zap?style=flat-square)](https://goreportcard.com/report/logur.dev/integration/zap)
[![GolangCI](https://golangci.com/badges/github.com/logur/integration-zap.svg)](https://golangci.com/r/github.com/logur/integration-zap)
[![Go Version](https://img.shields.io/badge/go%20version-%3E=1.11-61CFDD.svg?style=flat-square)](https://github.com/logur/integration-zap)
[![GoDoc](http://img.shields.io/badge/godoc-reference-5272B4.svg?style=flat-square)](https://godoc.org/logur.dev/integration/zap)


## Installation

```bash
go get logur.dev/integration/zap
```


## Usage

```go
package main

import (
	"logur.dev/logur"
	zapintegration "logur.dev/integration/zap"
)

func main() {
	logger := zapintegration.New(logur.NewNoopLogger())
}
```


## Development

When all coding and testing is done, please run the test suite:

``` bash
$ make check
```


## License

The MIT License (MIT). Please see [License File](LICENSE) for more information.
