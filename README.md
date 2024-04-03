[![godoc](https://img.shields.io/badge/godoc-reference-blue.svg)](https://pkg.go.dev/github.com/go-corelibs/tdata)
[![codecov](https://codecov.io/gh/go-corelibs/tdata/graph/badge.svg?token=KziW6VPHy1)](https://codecov.io/gh/go-corelibs/tdata)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-corelibs/tdata)](https://goreportcard.com/report/github.com/go-corelibs/tdata)

# tdata - go testing data utilities

A collection of utilities for working with go package unit testing data files.

There are two main interfaces provided by this package. The first is TestData
which makes working with `testdata` directories trivial. The second is TempData
which makes working with temporary directories more convenient than using the
standard `os` functions.


# Installation

``` shell
> go get github.com/go-corelibs/tdata@latest
```

# Examples

## TestData

``` go
var (
    // construct a new instance for accessing the `testdata` top-level path
    // this can be within functions or as a global shared instance
    td = tdata.New()
)

func TestSomeThing(t *testing.T) {
    contents := td.F("filename.txt")
    if SomeThing() == contents {
        t.Log("SomeThing is correct")
    } else {
        t.Error("SomeThing is not correct")
    }
}
```

## TempData

``` go
func TestSomeThing(t *testing.T) {
    tmpd, err := tdata.NewTempData("", "prefix-*.d")
    if err != nil {
        t.Fatalf("error creating TempData: %v", err)
    }
    defer tmpd.Destroy() // destroy the temporary directory when done
    
    // do stuff with tmpd.Path() or tmpd.Join()
}
```

# Go-CoreLibs

[Go-CoreLibs] is a repository of shared code between the [Go-Curses] and
[Go-Enjin] projects.

# License

```
Copyright 2024 The Go-CoreLibs Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use file except in compliance with the License.
You may obtain a copy of the license at

 http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
```

[Go-CoreLibs]: https://github.com/go-corelibs
[Go-Curses]: https://github.com/go-curses
[Go-Enjin]: https://github.com/go-enjin
