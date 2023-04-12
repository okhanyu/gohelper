package gohelper_server

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

var engine *gin.Engine

type Router struct {
	Prefix   string
	Uri      string
	Handlers []gin.HandlerFunc
}

type Routers struct {
	Get     []*Router
	Post    []*Router
	Delete  []*Router
	Put     []*Router
	Patch   []*Router
	Options []*Router
}

func GetRoutersInstance() *Routers {
	return &Routers{}
}

func (routers *Routers) BuildGet(getRouters []*Router) *Routers {
	routers.Get = getRouters
	return routers
}

func (routers *Routers) BuildPost(getRouters []*Router) *Routers {
	routers.Post = getRouters
	return routers
}

func (routers *Routers) BuildDelete(getRouters []*Router) *Routers {
	routers.Delete = getRouters
	return routers
}

func (routers *Routers) BuildPut(getRouters []*Router) *Routers {
	routers.Put = getRouters
	return routers
}

func (routers *Routers) BuildPatch(getRouters []*Router) *Routers {
	routers.Patch = getRouters
	return routers
}

func (routers *Routers) BuildOption(getRouters []*Router) *Routers {
	routers.Options = getRouters
	return routers
}

// NewRouter 实例化一个新的路由
func NewRouter(prefix string, Uri string, Handlers []gin.HandlerFunc) *Router {
	return &Router{
		Uri: fmt.Sprintf("%s%s", prefix, Uri), Handlers: Handlers,
	}
}

// StartServer 启动服务
func StartServer(routers *Routers, ipPort string, uses ...gin.HandlerFunc) *gin.Engine {
	//engine = gin.New()
	//engine.Use(gin.Recovery())
	engine = gin.Default()

	if uses != nil && len(uses) != 0 {
		engine.Use(uses...)
	}

	for _, router := range routers.Get {
		engine.GET(router.Uri, router.Handlers...)
	}

	for _, router := range routers.Delete {
		engine.DELETE(router.Uri, router.Handlers...)
	}

	for _, router := range routers.Put {
		engine.PUT(router.Uri, router.Handlers...)
	}

	for _, router := range routers.Post {
		engine.POST(router.Uri, router.Handlers...)
	}

	for _, router := range routers.Patch {
		engine.PATCH(router.Uri, router.Handlers...)
	}

	for _, router := range routers.Options {
		engine.OPTIONS(router.Uri, router.Handlers...)
	}

	if ipPort == "" {
		ipPort = ":8080"
	}
	err := engine.Run(ipPort)

	if err != nil {
		fmt.Printf("Gin engine start error is :%v", err)
		return nil
	}
	return engine
}
