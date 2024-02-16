// Copyright 2024 JetERA Creative
// This Source Code Form is subject to the terms of the Mozilla Public License, v. 2.0
// that can be found in the LICENSE file and https://mozilla.org/MPL/2.0/.

package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"time"
)

type GenerationReq struct {
	Model  string `json:"model"`
	Prompt string `json:"prompt"`
	Stream bool   `json:"stream"`
}

type Response struct {
	Model              string    `json:"model"`
	CreatedAt          time.Time `json:"created_at"`
	Response           string    `json:"response"`
	Done               bool      `json:"done"`
	TotalDuration      int64     `json:"total_duration"`
	LoadDuration       int64     `json:"load_duration"`
	PromptEvalCount    int       `json:"prompt_eval_count"`
	PromptEvalDuration int       `json:"prompt_eval_duration"`
	EvalCount          int       `json:"eval_count"`
	EvalDuration       int       `json:"eval_duration"`
}

type Instance struct {
	URL string
}

func (i *Instance) Generate(model string, prompt string, C chan<- string) error {
	s, err := json.Marshal(GenerationReq{
		Model:  model,
		Prompt: prompt,
		Stream: false,
	})
	if err != nil {
		return err
	}

	resp, err := http.Post(i.URL+"/api/generate", "application/json", bytes.NewBuffer(s))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	b := bytes.NewBuffer(nil)
	_, err = io.Copy(b, resp.Body)
	if err != nil {
		return err
	}

	var r Response

	err = json.Unmarshal(b.Bytes(), &r)
	if err != nil {
		return err
	}

	C <- r.Response

	return nil
}
