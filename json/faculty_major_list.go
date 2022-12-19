package json

import (
	_ "embed"
	"encoding/json"
)

//go:embed faculty_major_list.json
var files []byte
var NIMToString map[string]string

func InitNimToString() {
	json.Unmarshal(files, &NIMToString)
}
