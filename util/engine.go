package util

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"os"
	"sync"
)

var (
	engine *gin.Engine
	once   sync.Once
)

func GetInstance() *gin.Engine {
	once.Do(func() {
		engine = gin.Default()
	})
	return engine
}

func StartServer(addr ...string) {
	instance := GetInstance()
	err := instance.Run(addr...)
	if err != nil {
		log.Error("start server err", zap.Error(err))
		os.Exit(1)
	}
}
