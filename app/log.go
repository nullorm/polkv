package app

import (
	"log"

	"go.uber.org/zap"
)

var Log = newLogger()

func newLogger() *zap.Logger {
	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatalln("could not inititalize zap", err)
	}
	return logger
}
