package gohelper_demo

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/okhanyu/gohelper/gohelper_server"
)

func StartHttpServerDemo() {
	routers := make(map[string][]*gohelper_server.Router)

	prefix := "api/"

	Gets := make([]*gohelper_server.Router, 0)
	Gets = append(Gets, gohelper_server.GenerateRouter(prefix, "chat", []gin.HandlerFunc{GetHandlerFunc}))
	routers["GET"] = Gets

	Posts := make([]*gohelper_server.Router, 0)
	Posts = append(Posts, gohelper_server.GenerateRouter(prefix, "chat", []gin.HandlerFunc{PostHandlerFunc}))
	routers["POST"] = Posts

	//routers["GET"] = []general.Router{
	//	{Uri: "/", Handlers: []gin.HandlerFunc{a}},
	//	{Uri: "/abc", Handlers: []gin.HandlerFunc{b}},
	//}

	gohelper_server.StartServer(routers, ":8080")
}

func GetHandlerFunc(context *gin.Context) {
	fmt.Println("GetHandlerFunc")
	gohelper_server.Success(context, nil)

}

func PostHandlerFunc(context *gin.Context) {
	fmt.Println("PostHandlerFunc")
	context.String(200, "Hello, PostHandlerFunc")
}
