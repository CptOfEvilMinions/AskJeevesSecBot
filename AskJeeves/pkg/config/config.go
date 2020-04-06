package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Slack struct {
		Token string `yaml:"token"`
	} `yaml:"Slack"`
	Kafka struct {
		Hostname      string `yaml:"hostname"`
		Port          int    `yaml:"port"`
		GroupId       int    `yaml:"groupId"`
		Offset        string `yaml:"offset"`
		ConsumerTopic string `yaml:"consumer_topic"`
		PollInterval  int    `yaml:"poll_interval"`
	} `yaml:"kafka"`
	MySQL struct {
		Hostname string `yaml:"hostname"`
		Port     int    `yaml:"port"`
		Protocol string `yaml:"protocol"`
		Database string `yaml:"database"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		Expire   int    `yaml:"expire"`
		Interval int    `yaml:"interval"`
	} `yaml:"mysql"`
	GeoIP struct {
		FilePath   string `yaml:"file_path"`
		URL        string `yaml:"url"`
		LicenseKey string `yaml:"license_key"`
	} `yaml:"GeoIP"`
	ButlingButler struct {
		URL      string `yaml:"url"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		Interval int    `yaml:"interval"`
	} `yaml:"ButlingButler"`
	TheHive struct {
		URL    string `yaml:"url"`
		APIkey string `yaml:"api_key"`
	} `yaml:"theHive"`
}

// NewConfig returns a new decoded Config struct
func NewConfig(configPath string) (*Config, error) {
	// Create config structure
	config := &Config{}

	// Open config file
	file, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Init new YAML decode
	d := yaml.NewDecoder(file)

	// Start YAML decoding from file
	if err := d.Decode(&config); err != nil {
		return nil, err
	}

	return config, nil
}

// ValidateConfigPath just makes sure, that the path provided is a file,
// that can be read
func ValidateConfigPath(path string) error {
	s, err := os.Stat(path)
	if err != nil {
		return err
	}
	if s.IsDir() {
		return fmt.Errorf("'%s' is a directory, not a normal file", path)
	}
	return nil
}
