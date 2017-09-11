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

func TestGenerateRandomBytes(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("test did panic: %v", r)
		}
	}()

	b32 := GenerateRandomBytes(32)
	if len(b32) != 32 {
		t.Errorf("wrong result size: got %d want 32", len(b32))
	}

	b16 := GenerateRandomBytes(16)
	if len(b16) != 16 {
		t.Errorf("wrong result size: got %d want 16", len(b16))
	}
}

func TestReadRandomBytes(t *testing.T) {
	b := make([]byte, 12)
	n, err := ReadRandomBytes(b)
	if err != nil {
		t.Errorf("failed to read random bytes: %v", err)
	}
	if n != len(b) {
		t.Errorf("wrong return value: got %d want %d", n, len(b))
	}
}
