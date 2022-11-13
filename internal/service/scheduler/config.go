package scheduler

import (
	"gopkg.in/yaml.v2"
	"log"
	"os"
	"time"
)

type Config struct {
	CronExecutionTime string         `yaml:"cron_execution_time"`
	TypeOfWork        string         `yaml:"type_of_work"`
	StartHour         string         `yaml:"start_hour"`
	EndHour           string         `yaml:"end_hour"`
	WorkingDays       []time.Weekday `yaml:"working_days"`

	TelegramApiToken    string `yaml:"telegram_api_token"`
	TelegramChannelID   string `yaml:"telegram_channel_id"`
	TelegramChannelName string `yaml:"telegram_channel_name"`
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
