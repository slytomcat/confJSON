package сonfJSON

import (
	"encoding/json"
	"log"
	"os"
)
/*
 * ConfJSON - is a simple tool to read/store configuration in the JSON format file.
 *
 * Functions:
 *
 *   Load(file string, cfg *map[string]interface{})
 *
 * Load reads JSON file in to cfg map. All JSON keys become keys of map and all values
 * are stored as empty interface. So to use values yo have to convert them to requered format.
 *
 *   Save(file string, cfg map[string]interface{})
 *
 * Save stores configuration map in to JSON-formated file.
 * */

func Load(file string, cfg *map[string]interface{}) {
	f, err := os.Open(file)
	if err != nil {
		log.Fatal("Configuration file can't be read:", err)
	}
	defer f.Close()
	json.NewDecoder(f).Decode(&cfg)
}

func Save(file string, cfg map[string]interface{}) {
	f, err := os.Create(file)
	if err != nil {
		log.Fatal("Can't access to configuration file:", err)
	}
	defer f.Close()
	buf, _ := json.Marshal(cfg)
	f.Write(buf)
}
