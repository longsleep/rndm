/*
 * Copyright 2017 Kopano
 *
 * Use of this source code is governed by a MIT license
 * that can be found in the LICENSE.txt file.
 *
 */

package rndm

import (
	"testing"
)

func TestGenerateRandomString(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("test did panic: %v", r)
		}
	}()

	s32 := GenerateRandomString(32)
	if len(s32) != 44 {
		// Base64 encoded without padding.
		t.Errorf("wrong result size: got %d want 44", len(s32))
	}
}
