package common

import (
	"io/ioutil"
	"log"
	"os"
	"strconv"

	"gopkg.in/yaml.v3"
)

// Conf is the application configuration
type Conf struct {
	DB struct {
		Host string `yaml:"host"`
		Port int    `yaml:"port"`
		Name string `yaml:"name"`
		User string `yaml:"user"`
		Pass string `yaml:"pass"`
		Test bool   `yaml:"test"`
	}
}

// NewConf reads and save application config
func NewConf(YAMLFilePath string) *Conf {
	c := new(Conf)
	readFromYAML(c, YAMLFilePath)
	readFromEnvironment(c)
	return c
}

func readFromYAML(c *Conf, YAMLFilePath string) {
	yamlFile, err := ioutil.ReadFile(YAMLFilePath)
	if err != nil {
		log.Printf("error opening %s -> %v ", YAMLFilePath, err)
	}

	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Printf("error reading %s -> %v ", YAMLFilePath, err)
	}
}

func readFromEnvironment(c *Conf) {
	if c.DB.Host == "" {
		c.DB.Host = os.Getenv("DB_HOST")
	}
	if c.DB.Port == 0 {
		c.DB.Port, _ = strconv.Atoi(os.Getenv("DB_Port"))
	}
	if c.DB.Name == "" {
		c.DB.Name = os.Getenv("DB_NAME")
	}
	if c.DB.User == "" {
		c.DB.User = os.Getenv("DB_USER")
	}
	if c.DB.Pass == "" {
		c.DB.Pass = os.Getenv("DB_PASS")
	}
}
