package config

import (
	"log"
	"testing"
)

func TestConfigs_Init(t *testing.T) {
	c := Configs{
		Path: "../conf/conf-dev1.yaml",
	}

	if err := c.Init(); err != nil {
		log.Printf("config read error!")
		t.Fail()
	}
	log.Printf("appconf:%#v", c.AppConfig)
}
