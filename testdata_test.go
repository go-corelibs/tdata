// Copyright (c) 2024  The Go-Enjin Authors
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
	"fmt"
	"path/filepath"
	"runtime"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestTestData(t *testing.T) {

	_, src, _, _ := runtime.Caller(0)
	tdPath := filepath.Join(filepath.Dir(src), "testdata")
	_tdPath := filepath.Join(filepath.Dir(src), "_testdata")

	Convey("Default Path", t, func() {
		td := New()
		So(td, ShouldNotBeNil)
		So(td.Path(), ShouldEqual, tdPath)
		So(td.E("file.txt"), ShouldBeTrue)
		So(td.F("file.txt"), ShouldEqual, "test file\n")
		So(td.L("."), ShouldEqual, []string{
			td.Join("dir"),
			td.Join("file.txt"),
		})
		So(td.LD("."), ShouldEqual, []string{
			td.Join("dir"),
		})
		So(td.LF("."), ShouldEqual, []string{
			td.Join("file.txt"),
		})
		So(td.LA("."), ShouldEqual, []string{
			td.Join("dir"),
			td.Join("file.txt"),
		})
		So(td.LAH("."), ShouldEqual, []string{
			td.Join("dir"),
			td.Join("dir/.gitkeep"),
			td.Join("file.txt"),
		})
		So(td.LH("dir"), ShouldEqual, []string{
			td.Join("dir/.gitkeep"),
		})
	})

	Convey("Custom Path", t, func() {
		td := NewNamed("_testdata")
		So(td, ShouldNotBeNil)
		So(td.Path(), ShouldEqual, _tdPath)
		So(td.E("another.txt"), ShouldBeTrue)
		td = NewNamed("")
		So(td, ShouldNotBeNil)
		So(td.Path(), ShouldEqual, tdPath)
	})

	Convey("Constructor Errors", t, func() {

		Convey("Runtime Caller", func() {
			defer func() {
				r := recover()
				So(r, ShouldNotBeNil)
				So(r, ShouldEqual, ErrRuntimeCaller)
			}()

			td := newTestData(10000, DefaultTestData)
			So(td, ShouldBeNil)
		})

		Convey("Directory Not Found", func() {
			defer func() {
				r := recover()
				So(r, ShouldNotBeNil)
				So(r, ShouldEqual, fmt.Errorf("%w: %v", ErrNotFound, "not-a-thing"))
			}()

			td := NewNamed("not-a-thing")
			So(td, ShouldBeNil)
		})

	})

}
