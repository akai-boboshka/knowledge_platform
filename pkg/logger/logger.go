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

	// Запись логов в файл или консоль
	logFile, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		Log.SetOutput(logFile)
	} else {
		Log.SetOutput(os.Stdout) // Если файл недоступен, логируем в консоль
	}

	// Установка уровня логирования (например, Debug)
	Log.SetLevel(logrus.DebugLevel)
}
