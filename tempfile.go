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
)

// TempFile is a convenience wrapper around os.CreateTemp. It is the caller's
// responsibility to remove any temporary files made with TempFile when done
//
// Note: will return an empty string if os.CreateTemp returns an error
func TempFile(dir, pattern string) (filename string) {
	var err error
	var fh *os.File
	if fh, err = os.CreateTemp(dir, pattern); err == nil {
		filename = fh.Name()
		_ = fh.Close()
	}
	return
}
