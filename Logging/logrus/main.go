package main

import (
	"github.com/Sirupsen/logrus"
	"github.com/gemnasium/logrus-hooks/graylog"
	"github.com/plumbum/mgorus"
	"time"
)

func main() {

	// log.SetFormatter(&log.JSONFormatter{})
	// log.SetOutput(os.Stderr)

	log := logrus.New()
	log.Level = logrus.DebugLevel
	hook := graylog.NewGraylogHook("127.0.0.1:12201", "myFacility", map[string]interface{}{"startTime": time.Now().String()})
	log.Hooks.Add(hook)

	hooker, err := mgorus.NewHooker("localhost:27017", "logrus", "log")
	if err == nil {
		log.Hooks.Add(hooker)
		log.Info("MongoDB log ok")
	}

	log.Print("Simple print")
	log.Warn("warn")
	log.Info("some logging message")
	log.Debug("debug")
	log.Error("Is great error")
	log.Print("Сообщение на русском")

	log.WithFields(logrus.Fields{
		"name": "zhangsan",
		"age":  28,
	}).Error("Hello world!")

	log.WithField("extra", "Is extra message").WithField("date", time.Now().String()).Info("Item")

	time.Sleep(time.Second) // Ждём одну секунду, что бы логи вывалились в graylog

}
