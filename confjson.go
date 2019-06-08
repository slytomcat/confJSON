/*Package confjson - is a simple tool to read/store configuration in the JSON format file.
Values stored as empty interface. So, to use values yo have to convert them to required format.

Note that saved int value will be read as float64!
*/
package confjson

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

// Load reads JSON file in to cfg map.
func Load(file string) (map[string]interface{}, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, fmt.Errorf("Configuration file can't be read: %v", err)
	}
	defer f.Close()
	cfg := make(map[string]interface{})
	err = json.NewDecoder(f).Decode(&cfg)
	if err != nil && err != io.EOF {
		return nil, fmt.Errorf("Configuration file can be parsed: %v", err)
	}
	return cfg, nil
}

// Save stores configuration map in to JSON-formatted file.
func Save(file string, cfg map[string]interface{}) error {
	f, err := os.Create(file)
	if err != nil {
		return fmt.Errorf("Can't access configuration file: %v", err)
	}
	defer f.Close()
	buf, err := json.Marshal(cfg)
	if err != nil {
		return fmt.Errorf("Can't encode configuration data: %v", err)
	}
	f.Write(buf)
	return nil
}
