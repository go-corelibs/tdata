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
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestTempFile(t *testing.T) {
	Convey("basics", t, func() {
		filename := TempFile("", "tdata.*.txt")
		So(filename, ShouldNotEqual, "")
		stat, err := os.Stat(filename)
		So(err, ShouldBeNil)
		So(stat.Name(), ShouldEqual, filepath.Base(filename))
		So(stat.IsDir(), ShouldBeFalse)
		err = os.Remove(filename)
	})
}
