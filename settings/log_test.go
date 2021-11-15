package settings

import (
	log "github.com/sirupsen/logrus"
	"testing"
)

func TestAssistantFormatter_Format(t *testing.T) {
	InitLogger()
	InitSettings()
	log.WithField("db", App.Mysql.Db).Warn()
	log.WithFields(log.Fields{"host": App.Mysql.Host}).Info()
}
