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
		w.WriteHeader(func() int {
			ctx := r.Context()

			err := r.ParseForm()
			if err != nil {
				return 400
			}

			form, err := FormRequire(r.Form, "model", "prompt")
			if err != nil {
				return 400
			}

			println("Prompt:", form["prompt"])

			result := make(chan string, 1)

			if err := b.Instance.Generate(ctx, form["model"], form["prompt"], result); err != nil {
				println(err.Error())
				return 400
			}

			println("OK")

			select {
			case result := <-result:
				println(result)
				if _, err := w.Write([]byte(result)); err != nil {
					return 500
				}
			}

			return 200
		}())
	})

	return http.ListenAndServe(addr, mux)
}
