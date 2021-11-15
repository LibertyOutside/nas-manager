package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"nas-manager/settings"
	"testing"
)

func TestGetFile(t *testing.T) {
	f := GetFile()
	var c settings.AppConfig
	err := yaml.Unmarshal(f, &c)
	if err != nil {
		t.Fail()
	}
	fmt.Printf("%#v", c)
}
