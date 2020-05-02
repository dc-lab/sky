package modeldb

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type TagsMap map[string]string

type File struct {
	Id          string  `json:"id"`
	Owner       string  `json:"owner,omitempty"`
	Name        string  `json:"name,omitempty"`
	Tags        TagsMap `json:"tags,omitempty"`
	Hash        string  `json:"-"`
	UploadToken string  `json:"-"`
	ContentType string  `json:"-"`
}

func (m TagsMap) Value() (driver.Value, error) {
	return json.Marshal(m)
}

func (m *TagsMap) Scan(v interface{}) error {
	if v == nil {
		return nil
	}
	switch data := v.(type) {
	case string:
		return json.Unmarshal([]byte(data), &m)
	case []byte:
		return json.Unmarshal(data, &m)
	default:
		return fmt.Errorf("cannot scan type %t into TagsMap", v)
	}
}
