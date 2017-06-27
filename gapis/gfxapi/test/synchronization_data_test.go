// Copyright (C) 2017 Google Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package test

import (
	"testing"

	"github.com/google/gapid/core/assert"
	"github.com/google/gapid/core/log"
	"github.com/google/gapid/gapis/gfxapi/synchronization"
)

func TestSubcommandLessThan(t *testing.T) {
	ctx := log.Testing(t)
	assert.With(ctx).That(synchronization.SubcommandIndex{0}.LessThan(synchronization.SubcommandIndex{1})).Equals(true)
	assert.With(ctx).That(synchronization.SubcommandIndex{1}.LessThan(synchronization.SubcommandIndex{0})).Equals(false)
	assert.With(ctx).That(synchronization.SubcommandIndex{0}.LessThan(synchronization.SubcommandIndex{0})).Equals(false)
	assert.With(ctx).That(synchronization.SubcommandIndex{0, 0}.LessThan(synchronization.SubcommandIndex{0, 1})).Equals(true)
	assert.With(ctx).That(synchronization.SubcommandIndex{1, 0}.LessThan(synchronization.SubcommandIndex{0, 1})).Equals(false)
	assert.With(ctx).That(synchronization.SubcommandIndex{1, 0}.LessThan(synchronization.SubcommandIndex{0, 1})).Equals(false)

	assert.With(ctx).That(synchronization.SubcommandIndex{1, 0}.LessThan(synchronization.SubcommandIndex{1})).Equals(true)
	assert.With(ctx).That(synchronization.SubcommandIndex{1}.LessThan(synchronization.SubcommandIndex{1, 0})).Equals(false)
}

func deceq(s synchronization.SubcommandIndex, s2 synchronization.SubcommandIndex) bool {
	r := s
	r.Decrement()
	return !((r.LessThan(s2)) || s2.LessThan(r))
}

func TestDecrement(t *testing.T) {
	ctx := log.Testing(t)
	assert.With(ctx).That(deceq(synchronization.SubcommandIndex{1}, synchronization.SubcommandIndex{0})).Equals(true)
	assert.With(ctx).That(deceq(synchronization.SubcommandIndex{1, 1}, synchronization.SubcommandIndex{1, 0})).Equals(true)
	assert.With(ctx).That(deceq(synchronization.SubcommandIndex{1, 0}, synchronization.SubcommandIndex{0})).Equals(true)
	assert.With(ctx).That(deceq(synchronization.SubcommandIndex{2, 3, 4}, synchronization.SubcommandIndex{2, 3, 3})).Equals(true)
	assert.With(ctx).That(deceq(synchronization.SubcommandIndex{0}, synchronization.SubcommandIndex{})).Equals(true)
	assert.With(ctx).That(deceq(synchronization.SubcommandIndex{2, 3, 0}, synchronization.SubcommandIndex{2, 2})).Equals(true)
}