package config

import (
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	"path/filepath"
)

// logging font color
const (
	colorNormal = "\x1b[0m"  //默认
	colorRed    = "\x1b[31m" //红
	colorGreen  = "\x1b[32m" //绿
	colorYellow = "\x1b[33m" //黄
	colorBlue   = "\x1b[34m" //蓝
	colorPurple = "\x1b[35m" //紫
	colorCyan   = "\x1b[36m" //青
)

type AssistantFormatter struct{}

func (f AssistantFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}
	// time format
	timestamp := entry.Time.Format("2006-01-02 15:04:05")
	var newLog string
	// log format
	logLevel := logLevelStr(entry)
	if entry.HasCaller() {
		fName := filepath.Base(entry.Caller.File)
		newLog = fmt.Sprintf("[%s] [%s] %s%s%s [%s:%d %s]\n",
			timestamp, logLevel, colorCyan, msgWithFields(entry), colorNormal,
			fName, entry.Caller.Line, entry.Caller.Function)
	} else {
		newLog = fmt.Sprintf("[%s] %s[%s]%s %s\n", timestamp, colorGreen, entry.Level, colorNormal, entry.Message)
	}
	b.WriteString(newLog)
	return b.Bytes(), nil
}

func msgWithFields(entry *logrus.Entry) string {
	var s string
	for key := range entry.Data {
		s += key + ":" + entry.Data[key].(string) + " "
	}
	s += entry.Message
	return s
}

func logLevelStr(entry *logrus.Entry) string {
	level := entry.Level.String()
	var color string
	switch entry.Level {
	case logrus.DebugLevel, logrus.TraceLevel:
		color = colorBlue
	case logrus.InfoLevel:
		color = colorGreen
	case logrus.WarnLevel:
		color = colorYellow
	case logrus.ErrorLevel:
		color = colorRed
	case logrus.FatalLevel:
		color = colorPurple
	case logrus.PanicLevel:
		color = colorPurple
	default:
		color = colorNormal
	}
	return color + level + colorNormal
}

func InitLogger() {
	logrus.SetLevel(logrus.TraceLevel)
	logrus.SetReportCaller(true)
	logrus.SetFormatter(&AssistantFormatter{})
	logrus.Infoln("Logger Has Been Init")
}

func init() {
	InitLogger()
}
