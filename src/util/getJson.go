package util

import (
	"encoding/json"
	"fmt"
	"os"
)

func GetObjectFromJSON(path string, entity interface{}) interface{} {
	file, _ := os.Open(path)
	defer file.Close()
	decoder := json.NewDecoder(file)

	conf := entity
	err := decoder.Decode(&conf)
	if err != nil {
		fmt.Println("Error:", err)
	}
	return conf
}
