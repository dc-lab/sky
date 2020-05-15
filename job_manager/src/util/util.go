package util

import (
	"encoding/json"
	"io"
	"net/http"
)

func PanicOnError(err error, body http.ResponseWriter, code int) {
	if err != nil {
		m := make(map[string]error)
		m["error"] = err
		_ = EncodeBody(body, m)
		http.Error(body, "", code)
		panic(err)
	}
}

func DecodeBody(body io.Reader, v interface{}) error {
	decoder := json.NewDecoder(body)
	decoder.DisallowUnknownFields()
	return decoder.Decode(v)
}

func EncodeBody(body io.Writer, v interface{}) error {
	encoder := json.NewEncoder(body)
	encoder.SetIndent("", "    ")
	return encoder.Encode(v)
}
