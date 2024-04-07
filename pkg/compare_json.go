package pkg

import (
	"encoding/json"
	"reflect"
)

func CompareJSONMaps(jsonStr1, jsonStr2 string) (bool, error) {
	var map1, map2 map[string]interface{}

	if err := json.Unmarshal([]byte(jsonStr1), &map1); err != nil {
		return false, err
	}

	if err := json.Unmarshal([]byte(jsonStr2), &map2); err != nil {
		return false, err
	}

	return reflect.DeepEqual(map1, map2), nil
}
