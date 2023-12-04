package main

import "encoding/json"

func main() {

}

func obj2String(obj interface{}) (string, error) {
	bytes, err := json.Marshal(&obj)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}
