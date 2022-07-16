package util

import (
	"encoding/json"
	"fmt"
	"os"
)

// 从json文件中获取 Object
func GetObjectFromJSON(path string, entity interface{}) interface{} {
	file, _ := os.Open(path)
	defer file.Close()
	decoder := json.NewDecoder(file)

	err := decoder.Decode(&entity)
	if err != nil {
		fmt.Println("Error:", err)
	}
	return entity
}
