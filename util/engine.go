package util

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"os"
	"sync"
)

var (
	r    *gin.Engine
	once sync.Once
)

func GetInstance() *gin.Engine {
	once.Do(func() {
		r = gin.Default()
	})
	return r
}

func StartServer(addr ...string) {
	instance := GetInstance()
	err := instance.Run(addr...)
	if err != nil {
		log.Error("start server err", zap.Error(err))
		os.Exit(1)
	}
}
