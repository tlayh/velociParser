package service

import (
	"io/ioutil"
	"log"
	"gopkg.in/yaml.v2"
)

type Conf struct {
	User struct {
		Name string `yaml:"name"`
	}
	Scene struct {
		Track string `yaml:"track"`
		Url string `yaml:"url"`
	}
}

func ReadConfig() Conf {
	var config Conf

	yamlFile, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}

	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return config
}
