package monster

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Monster struct {
	Name  string
	Age   int
	Skill string
}

func (m *Monster) Store(filePath string) bool {
	// secialize
	data, err := json.Marshal(m)
	if err != nil {
		fmt.Println("\tSerialize Error:", err)
		return false
	}

	// save
	err = ioutil.WriteFile(filePath, data, 0666)
	if err != nil {
		fmt.Println("\tWrite File Error:", err)
		return false
	}
	return true
}

func (m *Monster) ReStore(filePath string) bool {
	// "./src/go_code/chapter15/test02/exercise/store.txt"
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("\tRead File Error:", err)
		return false
	}
	// fmt.Println(content)
	err = json.Unmarshal(content, &m)
	if err != nil {
		fmt.Println("\tDeserialize Error:", err)
		return false
	}
	return true
}
