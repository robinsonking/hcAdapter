/*
 * Copyright 2018 The openwallet Authors
 * This file is part of the openwallet library.
 *
 * The openwallet library is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * The openwallet library is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 * GNU Lesser General Public License for more details.
 */

package hdkeystore

import (
	"testing"
	"path/filepath"
)


func TestStoreHDKey(t *testing.T) {
	path := filepath.Join(".", "keys")
	_, rootId, err := StoreHDKey(path, "sogosdfo", "", StandardScryptN, StandardScryptP)
	if err != nil {
		t.Errorf("StoreHDKey failed unexpected error: %v", err)
	} else {
		t.Logf("StoreHDKey root id = %s", rootId)
	}
}

func TestGetKey(t *testing.T) {
	path := filepath.Join(".", "keys")
	ks := &HDKeystore{path, StandardScryptN, StandardScryptP}

	key, err := ks.GetKey("WAeAP5ggYYZ1euSJqURNEoGBRP6ucfPq2g",
		"sogosdfo-WAeAP5ggYYZ1euSJqURNEoGBRP6ucfPq2g.key",
		"")

	if err != nil {
		t.Errorf("GetKey failed unexpected error: %v\n", err)
	} else {
		t.Logf("GetKey root id = %s", key.KeyID)
	}
}