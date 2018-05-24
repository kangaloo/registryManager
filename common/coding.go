package common

import "encoding/json"

func Map2json(m map[string]interface{}) ([]byte, error) {

	j, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}
	return j, nil
}

func Json2map(j []byte) (map[string]interface{}, error) {

	var resMap = make(map[string]interface{})
	if err := json.Unmarshal(j, &resMap); err != nil {
		return nil, err
	}

	return resMap, nil
}
