package service

import (
	"io/ioutil"
	"log"
	"gopkg.in/yaml.v2"
)

type Conf struct {
	Users []User
	Rank int64
	Scenes []Scene
}

type User struct {
	Name string `yaml:"name"`
	Compare bool `yaml:"compare"`
}

type Scene struct {
	Track string `yaml:"track"`
	Url string `yaml:"url"`
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
