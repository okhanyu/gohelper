package gohelper_server

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

//var GlobalEngine *gin.Engine

type Router struct {
	Prefix   string
	Uri      string
	Handlers []gin.HandlerFunc
}

type Server struct {
	Get     []*Router
	Post    []*Router
	Delete  []*Router
	Put     []*Router
	Patch   []*Router
	Options []*Router
	Engine  *gin.Engine
}

// GetServerInstance 构造Server实例
func GetServerInstance() *Server {
	return &Server{Engine: gin.Default()}
}

func (server *Server) BuildGet(getRouters []*Router) *Server {
	server.Get = getRouters
	return server
}

func (server *Server) BuildPost(getRouters []*Router) *Server {
	server.Post = getRouters
	return server
}

func (server *Server) BuildDelete(getRouters []*Router) *Server {
	server.Delete = getRouters
	return server
}

func (server *Server) BuildPut(getRouters []*Router) *Server {
	server.Put = getRouters
	return server
}

func (server *Server) BuildPatch(getRouters []*Router) *Server {
	server.Patch = getRouters
	return server
}

func (server *Server) BuildOption(getRouters []*Router) *Server {
	server.Options = getRouters
	return server
}

func (server *Server) BuildUsesFunc(uses ...gin.HandlerFunc) *Server {
	if uses != nil && len(uses) != 0 {
		server.Engine.Use(uses...)
	}
	return server
}

// NewRouter 实例化一个新的路由
func NewRouter(prefix string, Uri string, Handlers []gin.HandlerFunc) *Router {
	return &Router{
		Uri: fmt.Sprintf("%s%s", prefix, Uri), Handlers: Handlers,
	}
}

// Build 启动服务
func (server *Server) Build(ipPort string) *gin.Engine {
	//engine = gin.New()
	//engine.Use(gin.Recovery())
	//engine = gin.Default()

	for _, router := range server.Get {
		server.Engine.GET(router.Uri, router.Handlers...)
	}

	for _, router := range server.Delete {
		server.Engine.DELETE(router.Uri, router.Handlers...)
	}

	for _, router := range server.Put {
		server.Engine.PUT(router.Uri, router.Handlers...)
	}

	for _, router := range server.Post {
		server.Engine.POST(router.Uri, router.Handlers...)
	}

	for _, router := range server.Patch {
		server.Engine.PATCH(router.Uri, router.Handlers...)
	}

	for _, router := range server.Options {
		server.Engine.OPTIONS(router.Uri, router.Handlers...)
	}

	if ipPort == "" {
		ipPort = ":8080"
	}

	err := server.Engine.Run(ipPort)

	if err != nil {
		fmt.Printf("Gin engine start error is :%v", err)
		return nil
	}
	return server.Engine
}
