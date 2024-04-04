// Copyright (c) 2024  The Go-CoreLibs Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package tdata provides simple unit testing utilities
package tdata

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	clPath "github.com/go-corelibs/path"
)

var _ TestData = (*testdata)(nil)

var (
	// DefaultTestData is the default name of the top-level package directory used
	// for storing unit test data files, "testdata" by default
	DefaultTestData = "testdata"
)

// TestData is an interface for interacting with a project's top-level
// unit testing files directory
type TestData interface {
	// Name returns the name of the testdata directory
	Name() (name string)

	TData
}

type testdata struct {
	tdata

	name string
}

// New is a convenience wrapper around NewNamed and the DefaultTestData
// directory name
func New() TestData {
	return newTestData(1, DefaultTestData)
}

// NewNamed constructs a new TestData instance using the directory of the
// runtime.Caller filename to find the correct package source directory, finds
// the `go.mod` file indicating the top-level of the Go package and then looks
// for the specified directory there. NewNamed will panic if the runtime.Caller
// response is not ok or if the test data directory is not found
func NewNamed(custom string) TestData {
	return newTestData(1, custom)
}

func newTestData(depth int, custom string) (td *testdata) {
	var ok bool
	var fn string
	if _, fn, _, ok = runtime.Caller(depth + 1); !ok {
		panic(ErrRuntimeCaller)
	}
	td = &testdata{name: custom}
	if custom == "" {
		custom = DefaultTestData
	}

	dirnames := strings.Split(filepath.Dir(fn), string(os.PathSeparator))
	for i := len(dirnames) - 1; i >= 0; i-- {
		check := "/" + filepath.Join(dirnames[:i+1]...)
		tdpath := filepath.Join(check, custom)
		if stopped := clPath.IsFile(filepath.Join(check, "go.mod")); stopped {
			if present := clPath.IsDir(tdpath); present {
				td.path = tdpath
				return
			}
			break
		}
	}

	panic(fmt.Errorf("%w: %s", ErrNotFound, custom))
}

func (td *testdata) Name() (name string) {
	return td.name
}
