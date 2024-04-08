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

// TestCheck is a generic interface for working with lists of comparable
// values during unit testing or other conditional branching procedures
//
// Example:
//
// if the original code does something like this:
//
//	slice := []string{"one", "two"}
//	check := len(slice) > 0
//	lookup := make(map[string]struct{})
//	for _, item := range slice {
//	  lookup[item] = struct{}{}
//	}
//
// and then further on, the above is used in cases like:
//
//	if _, present := lookup[item]; check && !present {
//	  ... do stuff ...
//	}
//
// then migrating to TestCheck would look like the following:
//
//	slice := []string{"one", "two"}
//	tc := NewTestCheck(len(slice) > 0, slice...)
//	... further on ...
//	if tc.NotPresent(item) {
//	  ... do stuff ...
//	}
type TestCheck[V comparable] interface {
	// Check returns conditional given during TestCheck construction
	Check() (ok bool)
	// List returns the argument list given during TestCheck construction
	List() (argv []V)
	// Present returns true if the conditional is true and the value
	// given is also present in the constructor arguments list
	Present(value V) (ok bool)
	// NotPresent returns true if the conditional is true and the value
	// given is not present in the constructor arguments list
	NotPresent(value V) (ok bool)
}

type testcheck[V comparable] struct {
	check  bool
	argv   []V
	lookup map[V]struct{}
}

// NewTestCheck constructs a new TestCheck instance of any comparable type
func NewTestCheck[V comparable](condition bool, arguments ...V) TestCheck[V] {
	tc := &testcheck[V]{
		check:  condition,
		argv:   arguments,
		lookup: make(map[V]struct{}),
	}
	for _, arg := range arguments {
		tc.lookup[arg] = struct{}{}
	}
	return tc
}

func (tc *testcheck[V]) Check() (ok bool) {
	return tc.check
}

func (tc *testcheck[V]) List() (argv []V) {
	return tc.argv[:]
}

func (tc *testcheck[V]) Present(value V) (ok bool) {
	_, present := tc.lookup[value]
	return tc.check && present
}

func (tc *testcheck[V]) NotPresent(value V) (ok bool) {
	_, present := tc.lookup[value]
	return tc.check && !present
}
