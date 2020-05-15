package main

import (
	"encoding/json"
	"io"
	"net/http"
)

func panicOnError(err error, body http.ResponseWriter, code int) {
	if err != nil {
		m := make(map[string]error)
		m["error"] = err
		_ = encodeBody(body, m)
		http.Error(body, "", code)
		panic(err)
	}
}

func decodeBody(body io.Reader, v interface{}) error {
	decoder := json.NewDecoder(body)
	decoder.DisallowUnknownFields()
	return decoder.Decode(v)
}

func encodeBody(body io.Writer, v interface{}) error {
	encoder := json.NewEncoder(body)
	encoder.SetIndent("", "    ")
	return encoder.Encode(v)
}
