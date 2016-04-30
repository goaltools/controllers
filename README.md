# Controllers
Officially supported controllers for the Goal based applications.

[![GoDoc](https://godoc.org/github.com/goaltools/controllers?status.svg)](https://godoc.org/github.com/goaltools/controllers)
[![Build Status](https://travis-ci.org/goaltools/controllers.svg?branch=master)](https://travis-ci.org/goaltools/controllers)
[![Coverage](https://codecov.io/github/goaltools/controllers/coverage.svg?branch=master)](https://codecov.io/github/goaltools/controllers?branch=master)
[![Go Report Card](http://goreportcard.com/badge/goaltools/controllers?t=3)](http:/goreportcard.com/report/goaltools/controllers)

### Usage
Update your `controllers/init.go` file by embedding necessary controllers as follows:
```go
package controllers

import (
	"github.com/goaltools/controllers/something"
)

type Controllers struct {
	*something.Something
}
```

### License
Distributed under the BSD 2-clause "Simplified" License unless otherwise noted.
