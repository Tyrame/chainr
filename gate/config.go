package main

import (
	"log"

	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Configuration struct {
	Port int
}

func LoadConfig(filename string) (Configuration, error) {
	log.Println("Reading configuration from", filename)
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return Configuration{}, err
	}

	var c Configuration
	err = yaml.Unmarshal(bytes, &c)
	if err != nil {
		return Configuration{}, err
	}

	return c, nil
}