package modeldb

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type TagsMap map[string]string

type File struct {
	Id         string  `json:"id" example:"6d83a3d2-16a6-486a-91a2-5d44ba74e326"`
	Owner      string  `json:"owner,omitempty" example:"b14bf169-3df5-4d61-ba94-1a09103cbdb2"`
	Name       string  `json:"name,omitempty" example:"file.txt"`
	Tags       TagsMap `json:"tags,omitempty"`
	TaskId     string  `json:"task_id,omitempty" example:""`
	Executable bool    `json:"executable"`

	Hash        string `json:"-"`
	UploadToken string `json:"-"`
	ContentType string `json:"-"`

	UploadUrls []string `json:"-"`
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
