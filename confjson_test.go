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
	err := Save(fileName, conf)
	if err != nil {
		log.Fatal(err)
	}
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

	readCfg, err := Load(fileName)

	if err != nil {
		log.Fatal(err)
	}

	for k, v := range readCfg {
		fmt.Printf("%s: %#v\n", k, v)
	}
	// Unordered output:
	// key1: "value1"
	// key2: 42
	// key3: true
}

func TestNoFileRead(t *testing.T) {

	_, err := Load("/fie_xyz_bad")

	if err == nil {
		t.Error("Succesful read from non-existing file")
	}
}

func TestNoFileSave(t *testing.T) {
	readCfg := make(map[string]interface{})

	err := Save("/fie_xyz_bad", readCfg)

	if err == nil {
		t.Error("Succesful save to inaccessible file")
	}

}

func TestLoadParseError(t *testing.T) {

	file, err := os.Create(fileName)
	if err != nil {
		t.Fatal(err)
	}

	_, err = file.WriteString("{\"key1\":\"value1\",\"key2\":42,\"key3\"=true}")
	if err != nil {
		file.Close()
		t.Fatal(err)
	}
	file.Sync()
	file.Close()

	_, err = Load(fileName)
	if err == nil {
		t.Error("Succesful read from incorrect file")
	}
}
