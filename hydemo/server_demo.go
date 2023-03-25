package hydemo

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/okhanyu/gohelper/hyserver"
)

func StartHttpServerDemo() {
	routers := make(map[string][]*hyserver.Router)

	routers = make(map[string][]*hyserver.Router)
	prefix := "api/"

	Gets := make([]*hyserver.Router, 0)
	Gets = append(Gets, hyserver.GenerateRouter(prefix, "chat", []gin.HandlerFunc{GetHandlerFunc}))
	routers["GET"] = Gets

	Posts := make([]*hyserver.Router, 0)
	Posts = append(Posts, hyserver.GenerateRouter(prefix, "chat", []gin.HandlerFunc{PostHandlerFunc}))
	routers["POST"] = Posts

	//routers["GET"] = []general.Router{
	//	{Uri: "/", Handlers: []gin.HandlerFunc{a}},
	//	{Uri: "/abc", Handlers: []gin.HandlerFunc{b}},
	//}

	hyserver.StartServer(routers, ":8080")
}

func GetHandlerFunc(context *gin.Context) {
	fmt.Println("GetHandlerFunc")
	hyserver.Success(context, nil)

}

func PostHandlerFunc(context *gin.Context) {
	fmt.Println("PostHandlerFunc")
	context.String(200, "Hello, PostHandlerFunc")
}
