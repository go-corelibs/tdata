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
	"path/filepath"

	clPath "github.com/go-corelibs/path"
)

var _ TData = (*tdata)(nil)

// TData is the filesystem interface common to both TestData and TempData
// implementations
type TData interface {
	// Path returns the absolute path to this instance's directory
	Path() (path string)
	// Join is a convenience wrapper around Path and filepath.Join
	Join(names ...string) (joined string)
	// E reports whether the given file exists (as a file or a directory)
	E(filename string) (exists bool)
	// F reads the given file and returns the contents
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

type tdata struct {
	path string
}

func (td *tdata) Path() (abs string) {
	return td.path
}

func (td *tdata) Join(names ...string) (joined string) {
	return filepath.Join(append([]string{td.path}, names...)...)
}

func (td *tdata) E(filename string) (exists bool) {
	exists = clPath.Exists(filepath.Join(td.path, filename))
	return
}

func (td *tdata) F(filename string) (contents string) {
	if data, err := os.ReadFile(filepath.Join(td.path, filename)); err == nil {
		contents = string(data)
	}
	return
}

func (td *tdata) L(dirname string) (found []string) {
	found, _ = clPath.List(filepath.Join(td.path, dirname), false)
	return
}

func (td *tdata) LD(dirname string) (found []string) {
	found, _ = clPath.ListDirs(filepath.Join(td.path, dirname), false)
	return
}

func (td *tdata) LF(dirname string) (found []string) {
	found, _ = clPath.ListFiles(filepath.Join(td.path, dirname), false)
	return
}

func (td *tdata) LA(dirname string) (found []string) {
	found = append(found, td.LAD(dirname)...)
	found = append(found, td.LAF(dirname)...)
	return
}

func (td *tdata) LAD(dirname string) (found []string) {
	found, _ = clPath.ListAllDirs(filepath.Join(td.path, dirname), false)
	return
}

func (td *tdata) LAF(dirname string) (found []string) {
	found, _ = clPath.ListAllFiles(filepath.Join(td.path, dirname), false)
	return
}

func (td *tdata) LH(dirname string) (found []string) {
	found, _ = clPath.List(filepath.Join(td.path, dirname), true)
	return
}

func (td *tdata) LAH(dirname string) (found []string) {
	found = append(found, td.LADH(dirname)...)
	found = append(found, td.LAFH(dirname)...)
	return
}

func (td *tdata) LADH(dirname string) (found []string) {
	found, _ = clPath.ListAllDirs(filepath.Join(td.path, dirname), true)
	return
}

func (td *tdata) LAFH(dirname string) (found []string) {
	found, _ = clPath.ListAllFiles(filepath.Join(td.path, dirname), true)
	return
}
