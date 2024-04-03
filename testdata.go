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

var (
	// DefaultTestData is the default name of the top-level package directory used
	// for storing unit test data files, "testdata" by default
	DefaultTestData = "testdata"
)

// TestData is an interface for interacting with a project's top-level
// unit testing files directory
type TestData interface {
	// Path returns the absolute path to this instance's testdata directory
	Path() (path string)
	// Join is a convenience wrapper around Path and filepath.Join
	Join(names ...string) (joined string)
	// E reports whether the given file exists (as a file or a directory) within
	// the testdata filesystem
	E(filename string) (exists bool)
	// F reads the given file from the testdata filesystem and returns the contents
	F(filename string) (contents string)
	// L lists files and directories within the dirname given
	L(dirname string) (found []string)
	// LD lists directories within the dirname given
	LD(dirname string) (found []string)
	// LF lists files within the dirname given
	LF(dirname string) (found []string)
	// LA lists files and directories, recursively
	LA(dirname string) (found []string)
	// LAF lists files, recursively
	LAF(dirname string) (found []string)
	// LAD lists directories, recursively
	LAD(dirname string) (found []string)
	// LH is the same as L except including hidden files
	LH(dirname string) (found []string)
	// LAH is the same as LA except including hidden files
	LAH(dirname string) (found []string)
	// LAFH is the same as LAF except including hidden files
	LAFH(dirname string) (found []string)
	// LADH is the same as LAD except including hidden files
	LADH(dirname string) (found []string)
}

type testdata struct {
	name string
	path string
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

func (td *testdata) Path() (path string) {
	return td.path
}

func (td *testdata) Join(names ...string) (joined string) {
	return filepath.Join(append([]string{td.path}, names...)...)
}

func (td *testdata) E(filename string) (exists bool) {
	exists = clPath.Exists(filepath.Join(td.path, filename))
	return
}

func (td *testdata) F(filename string) (contents string) {
	if data, err := os.ReadFile(filepath.Join(td.path, filename)); err == nil {
		contents = string(data)
	}
	return
}

func (td *testdata) L(dirname string) (found []string) {
	found, _ = clPath.List(filepath.Join(td.path, dirname), false)
	return
}

func (td *testdata) LD(dirname string) (found []string) {
	found, _ = clPath.ListDirs(filepath.Join(td.path, dirname), false)
	return
}

func (td *testdata) LF(dirname string) (found []string) {
	found, _ = clPath.ListFiles(filepath.Join(td.path, dirname), false)
	return
}

func (td *testdata) LA(dirname string) (found []string) {
	found = append(found, td.LAD(dirname)...)
	found = append(found, td.LAF(dirname)...)
	return
}

func (td *testdata) LAD(dirname string) (found []string) {
	found, _ = clPath.ListAllDirs(filepath.Join(td.path, dirname), false)
	return
}

func (td *testdata) LAF(dirname string) (found []string) {
	found, _ = clPath.ListAllFiles(filepath.Join(td.path, dirname), false)
	return
}

func (td *testdata) LH(dirname string) (found []string) {
	found, _ = clPath.List(filepath.Join(td.path, dirname), true)
	return
}

func (td *testdata) LAH(dirname string) (found []string) {
	found = append(found, td.LADH(dirname)...)
	found = append(found, td.LAFH(dirname)...)
	return
}

func (td *testdata) LADH(dirname string) (found []string) {
	found, _ = clPath.ListAllDirs(filepath.Join(td.path, dirname), true)
	return
}

func (td *testdata) LAFH(dirname string) (found []string) {
	found, _ = clPath.ListAllFiles(filepath.Join(td.path, dirname), true)
	return
}
