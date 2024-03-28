// Copyright 2024 JetERA Creative
// This Source Code Form is subject to the terms of the Mozilla Public License, v. 2.0
// that can be found in the LICENSE file and https://mozilla.org/MPL/2.0/.

package main

import (
	"github.com/jellytea/go-httpform"
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

			form, err := httpform.WrapFromRequest(r)
			if err != nil {
				return 400
			}

			model := form.String("model")
			prompt := form.String("prompt")

			err = form.Parse()
			if err != nil {
				return 500
			}

			resp, err := b.Instance.Generate(ctx, *model, *prompt)
			if err != nil {
				return 500
			}

			_, _ = w.Write([]byte(resp))

			return 200
		}())
	})

	return http.ListenAndServe(addr, mux)
}
