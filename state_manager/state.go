package state_manager

import (
	"github.com/tidwall/gjson"
)

func Get(json, path string) gjson.Result {
	return gjson.Get(json, path)
}
