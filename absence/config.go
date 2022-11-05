package absence

import (
	"gopkg.in/yaml.v2"
	"log"
	"os"
	"time"
)

type Config struct {
	ID          string         `yaml:"absence_id"`
	Key         string         `yaml:"absence_key"`
	TypeOfWork  string         `yaml:"type_of_work"`
	StartHour   string         `yaml:"start_hour"`
	EndHour     string         `yaml:"end_hour"`
	WorkingDays []time.Weekday `yaml:"working_days"`
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
