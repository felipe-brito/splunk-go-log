package util

import "github.com/sirupsen/logrus"

type logger struct {
	log *logrus.Logger
}

func Log() logger {

	l := logrus.New()
	l.SetFormatter(&logrus.JSONFormatter{})

	return logger{
		log: l,
	}
}

func (l logger) Info(i ...interface{}) {
	l.log.SetLevel(logrus.InfoLevel)
	l.log.Info(i)
}

func (l logger) Error(i interface{}, err error) {
	l.log.SetLevel(logrus.ErrorLevel)
	l.log.WithField("msg", i).WithError(err)
}

func (l logger) Warn(i interface{}) {
	l.log.SetLevel(logrus.WarnLevel)
	l.log.Warn(i)
}
