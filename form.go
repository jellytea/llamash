// Copyright 2024 JetERA Creative
// This Source Code Form is subject to the terms of the Mozilla Public License, v. 2.0
// that can be found in the LICENSE file and https://mozilla.org/MPL/2.0/.

package main

import (
	"errors"
	"net/url"
)

func FormRequire(r url.Values, keys ...string) (values map[string]string, _ error) {
	values = make(map[string]string, len(keys))
	for _, key := range keys {
		if v := r.Get(key); v == "" {
			return nil, errors.New("missing parameter: " + key)
		} else {
			values[key] = v
		}
	}

	return
}
