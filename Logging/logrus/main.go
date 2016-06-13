package main

import (
	"github.com/Sirupsen/logrus"
	"github.com/evalphobia/logrus_sentry"
	"github.com/getsentry/raven-go"
	"github.com/plumbum/mgorus"
	"time"
)

type HS map[string]string

func main() {

	// log.SetFormatter(&log.JSONFormatter{})
	// log.SetOutput(os.Stderr)

	log := logrus.New()
	log.Level = logrus.DebugLevel
	/*
		hook := graylog.NewGraylogHook("127.0.0.1:12201", "myFacility", map[string]interface{}{"startTime": time.Now().String()})
		log.Hooks.Add(hook)
	*/

	ravenClient, err := raven.New("http://78d5df1b220e47958c28fbab30ac92d5:db59624047fd4616939520e684a62dd7@172.17.0.5:9000/2")
	if err == nil {
		ravenClient.CaptureMessage("Запущен Sentry", HS{"tag1": "one", "tag2": "two"})
		hookSentry, err := logrus_sentry.NewWithClientSentryHook(
			ravenClient,
			[]logrus.Level{
				logrus.PanicLevel,
				logrus.FatalLevel,
				logrus.ErrorLevel,
			})
		if err == nil {
			log.Hooks.Add(hookSentry)
			log.Info("Sentry logger OK")
			ravenClient.CaptureMessage("Подключили Sentry к логу", HS{"tag1": "one", "tag2": "two"})
		} else {
			log.Warn("Can't create Sentry hook: ", err)
		}


		ravenClient.CapturePanic(func () {
			panic("Здесь перехватываем панику")
		}, HS{"status": "panic"})

		ravenClient.

	} else {
		log.Warn("Can't connect to Sentry: ", err)
	}

	hookMongo, err := mgorus.NewHooker("localhost:27017", "logrus", "log")
	if err == nil {
		log.Hooks.Add(hookMongo)
		log.Info("MongoDB logger OK")
	} else {
		log.Warn("Can't create Mongo hook", err)
	}

	log.Print("Simple print")
	log.Warn("warn")
	log.Info("some logging message")
	log.Debug("debug")
	time.Sleep(time.Second) // Ждём одну секунду, что бы логи вывалились в graylog
	log.Error("Is great error")
	log.WithField("lang", "ru-RU").Print("Сообщение на русском")

	log.WithFields(logrus.Fields{
		"name": "zhangsan",
		"age":  28,
	}).Error("Hello world!")

	log.WithField("extra", "Is extra message").WithField("date", time.Now().String()).Info("Item")

	ravenClient.Wait()
	time.Sleep(time.Second) // Ждём одну секунду, что бы логи вывалились в graylog

}
