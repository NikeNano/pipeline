/*
Copyright 2020 The Tekton Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package diff

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func FailIfWantGot(want, got interface{}, errorFunc func(string, ...interface{}), format string, t *testing.T) {
	t.Helper()
	if d := cmp.Diff(want, got); d != "" {
		errorFunc(format+", want: [%v], got: [%v], diff: [%v]", want, got, d)
	}
}
