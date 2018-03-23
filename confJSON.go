/*Package confJSON - is a simple tool to read/store configuration in the JSON format file.
Values stored as empty interface. So, to use values yo have to convert them to required format.

Note that saved int value will be read as float64!
*/
package confJSON

import (
	"encoding/json"
	"log"
	"os"
)

// Load reads JSON file in to cfg map.
func Load(file string, cfg *map[string]interface{}) {
	f, err := os.Open(file)
	if err != nil {
		log.Fatal("Configuration file can't be read:", err)
	}
	defer f.Close()
	json.NewDecoder(f).Decode(&cfg)
}

// Save stores configuration map in to JSON-formatted file.
func Save(file string, cfg map[string]interface{}) {
	f, err := os.Create(file)
	if err != nil {
		log.Fatal("Can't access to configuration file:", err)
	}
	defer f.Close()
	buf, _ := json.Marshal(cfg)
	f.Write(buf)
}
