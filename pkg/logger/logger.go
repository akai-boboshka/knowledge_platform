package logger

import (
	"github.com/sirupsen/logrus"
	"os"
)

func Init() {
	Log := logrus.New()

	// Настройка формата вывода
	Log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	logFile, err := os.OpenFile("./pkg/logger/app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		Log.SetOutput(logFile)
	} else {
		Log.SetOutput(os.Stdout)
	}

	// Установка уровня логирования (например, Debug)
	Log.SetLevel(logrus.InfoLevel)
}
