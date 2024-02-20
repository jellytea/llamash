// Copyright 2024 JetERA Creative
// This Source Code Form is subject to the terms of the Mozilla Public License, v. 2.0
// that can be found in the LICENSE file and https://mozilla.org/MPL/2.0/.

package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	err := _main()
	if err != nil {
		fmt.Println(err)
		return
	}
}

func _main() error {
	llamaHost := os.Getenv("OLLAMA_HOST")
	if llamaHost == "" {
		flag.StringVar(&llamaHost, "i", "http://127.0.0.1:11434", "O-LLaMA server address.")
	}

	port := flag.String("p", "11444", "Bridge port")

	flag.Parse()

	println("OLLAMA_HOST:", llamaHost)

	b := Bridge{Instance: &Instance{llamaHost}}

	err := b.Serve(":" + *port)
	if err != nil {
		return err
	}

	return nil
}
