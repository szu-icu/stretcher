package router

import (
	"github.com/gin-gonic/gin"
	log "github.com/shengkehua/xlog4go"
)

var groupList = []func(*gin.Engine){}

func RouterRegister() *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())

	// setup log
	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
		log.Info("%v: %v \n %v \n %v\n", absolutePath, httpMethod, handlerName, nuHandlers)
	}

	// setup total middleware

	// setup group
	for _, group := range groupList {
		group(r)
	}

	return r
}
