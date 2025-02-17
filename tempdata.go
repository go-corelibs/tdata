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

package tdata

import (
	"os"

	clPath "github.com/go-corelibs/path"
)

var _ TempData = (*tempdata)(nil)

// TempData is an interface for creating and interacting with temporary
// directories for unit testing purposes
type TempData interface {
	// Create will make the temporary directory associated with this TempData
	// instance if it does not exist. Create does nothing if the directory
	// exists. Create is only useful after a call to Destroy because the
	// temporary directory for this instance has already been created during
	// the call to NewTempData
	Create() (err error)
	// Destroy attempts to correct any file permissions and removes the
	// entire temporary directory associated with this TempData instance
	Destroy() (err error)

	TData
}

// NewTempData constructs a new TempData instance using the given `dir` and
// `pattern` arguments in a call to os.MkdirTemp
func NewTempData(dir, pattern string) (td TempData, err error) {
	return newTempData(dir, pattern)
}

func newTempData(dir, pattern string) (td *tempdata, err error) {
	var path string
	if path, err = os.MkdirTemp(dir, pattern); err == nil {
		td = &tempdata{}
		td.path = path
	}
	return
}

type tempdata struct {
	tdata
}

func (td *tempdata) Create() (err error) {
	if clPath.IsDir(td.path) {
		return
	}
	// same permissions as os.MkdirTemp
	err = os.MkdirAll(td.path, 0700)
	return
}

func (td *tempdata) Destroy() (err error) {
	if clPath.IsDir(td.path) {
		if err = clPath.ChmodAll(td.path); err == nil {
			err = os.RemoveAll(td.path)
		}
	}
	return
}
