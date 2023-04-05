package util

import (
	"go.uber.org/zap"
	"sync"
)

var (
	log     *zap.Logger
	onceLog sync.Once
)

func GetLogger() *zap.Logger {
	onceLog.Do(func() {
		logger, _ := zap.NewProduction()
		log = logger
	})
	return log
}

func SyncLog() {
	onceLog.Do(func() {
		err := log.Sync()
		if err != nil {
			log.Error("sync log err", zap.Error(err))
		}
	})
}
