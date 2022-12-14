package absence

import (
	"gopkg.in/yaml.v2"
	"log"
	"os"
)

type Config struct {
	ID  string `yaml:"absence_id"`
	Key string `yaml:"absence_key"`
}

func NewFromPath(path string) *Config {
	c := new(Config)
	yamlFile, err := os.ReadFile(path)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return c
}
