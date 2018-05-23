package common

import "encoding/json"

func Map2json(m map[string][]string) ([]byte, error) {

	j, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}

	return j, nil
}


func Json2map(j []byte) (map[string][]string, error) {

	var resMap = make(map[string][]string)
	if err := json.Unmarshal(j, &resMap); err != nil {
		return nil, err
	}

	return resMap, nil
}
