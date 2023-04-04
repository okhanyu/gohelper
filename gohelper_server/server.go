package gohelper_server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

var engine *gin.Engine

type Router struct {
	Prefix   string
	Uri      string
	Handlers []gin.HandlerFunc
}

func GenerateRouter(prefix string, Uri string, Handlers []gin.HandlerFunc) *Router {
	return &Router{
		Uri: fmt.Sprintf("%s%s", prefix, Uri), Handlers: Handlers,
	}
}

// StartServer 启动服务
func StartServer(routers map[string][]*Router, ipPort string, f ...gin.HandlerFunc) *gin.Engine {
	//engine = gin.New()
	//engine.Use(gin.Recovery())
	engine = gin.Default()

	if f != nil && len(f) != 0 {
		engine.Use(f...)
	}

	for method, routerGroup := range routers {
		if method == http.MethodGet {
			for _, info := range routerGroup {
				engine.GET(info.Uri, info.Handlers...)
			}
		}
		if method == http.MethodPost {
			for _, info := range routerGroup {
				engine.POST(info.Uri, info.Handlers...)
			}
		}
		if method == http.MethodPut {
			for _, info := range routerGroup {
				engine.PUT(info.Uri, info.Handlers...)
			}
		}
		if method == http.MethodDelete {
			for _, info := range routerGroup {
				engine.DELETE(info.Uri, info.Handlers...)
			}
		}
		if method == http.MethodPatch {
			for _, info := range routerGroup {
				engine.PATCH(info.Uri, info.Handlers...)
			}
		}
	}

	err := engine.Run(ipPort)

	if err != nil {
		fmt.Println(err)
		return nil
	}

	return engine
}
