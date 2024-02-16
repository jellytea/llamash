// Copyright 2024 JetERA Creative
// This Source Code Form is subject to the terms of the Mozilla Public License, v. 2.0
// that can be found in the LICENSE file and https://mozilla.org/MPL/2.0/.

package main

import (
	"net/http"
)

type Bridge struct {
	Instance *Instance
}

func (b *Bridge) Serve(addr string) error {
	mux := http.NewServeMux()
	mux.HandleFunc("/generate", func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			w.WriteHeader(400)
			return
		}

		model := r.Form.Get("model")
		prompt := r.Form.Get("prompt")

		if model != "" && prompt != "" {
			c := make(chan string, 1)
			if b.Instance.Generate(model, prompt, c) != nil {
				goto FAIL
			}

			select {
			case result := <-c:
				if _, err := w.Write([]byte(result)); err != nil {
					goto FAIL
				}
			}

			return
		}
	FAIL:
		w.WriteHeader(400)
	})

	return http.ListenAndServe(addr, mux)
}
