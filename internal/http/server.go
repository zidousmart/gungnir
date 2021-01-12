package http

import (
	"github.com/gin-gonic/gin"
	"github.com/zidousmart/gungnir/internal/mode"
)

type GHttpEngine struct {
	e *gin.Engine
}

func NewHttpServer(env string) *GHttpEngine {
	if env == mode.EnvDev {
		gin.SetMode(gin.DebugMode)
	} else if env == mode.EnvTest {
		gin.SetMode(gin.TestMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	return &GHttpEngine{
		e: gin.New(),
	}
}

func (s *GHttpEngine) Run(addr ...string) error {
	return s.e.Run(addr...)
}
