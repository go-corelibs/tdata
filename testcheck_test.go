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
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestTestCheck(t *testing.T) {

	Convey("Check", t, func() {

		tc := NewTestCheck[string](true)
		So(tc, ShouldNotBeNil)
		So(tc.Check(), ShouldBeTrue)
		So(tc.List(), ShouldEqual, []string(nil))

		tc = NewTestCheck[string](false)
		So(tc, ShouldNotBeNil)
		So(tc.Check(), ShouldBeFalse)
		So(tc.List(), ShouldEqual, []string(nil))

	})

	Convey("List", t, func() {

		tc := NewTestCheck(false, "one", "two")
		So(tc, ShouldNotBeNil)
		So(tc.Check(), ShouldBeFalse)
		So(tc.List(), ShouldEqual, []string{"one", "two"})

	})

	Convey("Present", t, func() {

		tc := NewTestCheck(false, "one", "two")
		So(tc, ShouldNotBeNil)
		So(tc.Check(), ShouldBeFalse)
		So(tc.Present("one"), ShouldBeFalse)

		tc = NewTestCheck(true, "one", "two")
		So(tc, ShouldNotBeNil)
		So(tc.Check(), ShouldBeTrue)
		So(tc.Present("nope"), ShouldBeFalse)
		So(tc.Present("one"), ShouldBeTrue)

	})

	Convey("NotPresent", t, func() {

		tc := NewTestCheck(false, "one", "two")
		So(tc, ShouldNotBeNil)
		So(tc.Check(), ShouldBeFalse)
		So(tc.NotPresent("nope"), ShouldBeFalse)

		tc = NewTestCheck(true, "one", "two")
		So(tc, ShouldNotBeNil)
		So(tc.Check(), ShouldBeTrue)
		So(tc.NotPresent("nope"), ShouldBeTrue)
		So(tc.NotPresent("one"), ShouldBeFalse)

	})

}
