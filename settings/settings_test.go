package settings

import (
	"log"
	"testing"
)

func TestConfigs_Init(t *testing.T) {

	InitSettings()

	log.Printf("appconf:%#v", App)
}
