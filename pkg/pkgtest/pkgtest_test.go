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

package pkgtest

import (
	"path/filepath"
	"runtime"
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/go-corelibs/tdata"
)

func TestPkgTest(t *testing.T) {

	_, src, _, _ := runtime.Caller(0)
	topdir := filepath.Dir(filepath.Dir(filepath.Dir(src)))
	tdPath := filepath.Join(topdir, "testdata")

	Convey("Default Path", t, func() {
		td := tdata.New()
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

}
