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
	"path/filepath"
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	clPath "github.com/go-corelibs/path"
)

func TestTempData(t *testing.T) {
	Convey("Instance Checks", t, func() {
		td, err := NewTempData("", "tdata.*")
		So(err, ShouldBeNil)
		So(td, ShouldNotBeNil)
		path := td.Path()
		So(path, ShouldNotBeEmpty)
		So(clPath.IsDir(path), ShouldBeTrue)
		So(td.Join(), ShouldEqual, path)
		So(td.Join("nope"), ShouldEqual, filepath.Join(path, "nope"))
		So(td.Destroy(), ShouldBeNil)
		So(clPath.IsDir(td.Path()), ShouldBeFalse)
		So(td.Create(), ShouldBeNil)
		So(clPath.IsDir(td.Path()), ShouldBeTrue)
		So(td.Create(), ShouldBeNil) // nop check
		So(td.Destroy(), ShouldBeNil)
	})
}
