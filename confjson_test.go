package confjson

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

var fileName string

func TestMain(m *testing.M) {
	fileName = filepath.Join(os.TempDir(), "temp_fie_xyz")
	err := m.Run()
	os.Remove(fileName)
	os.Exit(err)
}

func ExampleSave() {

	conf := map[string]interface{}{
		"key1": "value1",
		"key2": 42,
		"key3": true,
	}
	Save(fileName, conf)
	// read the file
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	data := make([]byte, 100)
	l, err := file.Read(data)
	if err != nil {
		log.Fatal(err)
	}
	// remove "{" and "}"
	s := string(data[1 : l-1])
	s = strings.Replace(s, ",", "\n", -1)
	fmt.Println(s)
	// Unordered output:
	// "key1":"value1"
	// "key2":42
	// "key3":true
}

func ExampleLoad() {
	readCfg := make(map[string]interface{})
	Load(fileName, &readCfg)

	for k, v := range readCfg {
		fmt.Printf("%s: %#v\n", k, v)
	}
	// Unordered output:
	// key1: "value1"
	// key2: 42
	// key3: true

}
