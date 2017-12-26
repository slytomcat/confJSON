package сonfJSON

import (
	"encoding/json"
	"log"
	"os"
)

func Load(filePath string, cfg *map[string]interface{}) {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Configurations' file can't be read:", err)
	}
	defer f.Close()
	json.NewDecoder(f).Decode(&cfg)
}

func Save(filePath string, cfg map[string]interface{}) {
	f, err := os.Create(filePath)
	if err != nil {
		log.Fatal("Can't access to configuration file:", err)
	}
	defer f.Close()
	buf, _ := json.Marshal(cfg)
	f.Write(buf)

}
