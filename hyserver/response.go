package hyserver

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	CodeError   = -1 // 未知错误
	CodeSuccess = 0
)

var CodeMsgMap = make(map[int]string)
var CustomCodeMsgMap = make(map[int]string)

func init() {
	CodeMsgMap[CodeSuccess] = "成功"
	CodeMsgMap[CodeError] = "未知错误"
}

// Init 初始化自定义错误信息表
func Init(customCodeMsg map[int]string) {
	CustomCodeMsgMap = customCodeMsg
}

// Success 固定格式
func Success(c *gin.Context, data interface{}) {
	c.JSON(
		http.StatusOK,
		gin.H{
			"code": CodeSuccess,
			"msg":  "success",
			"data": data,
		},
	)
}

// SuccessPlain 原样
func SuccessPlain(c *gin.Context, data interface{}) {
	c.JSON(
		http.StatusOK,
		data,
	)
}

// Fail 未知错误
func Fail(c *gin.Context, code int) {
	c.JSON(
		http.StatusOK,
		gin.H{
			"code": code,
			"msg":  CodeMsgMap[CodeError],
		},
	)
}

// FailWithErrCode 错误码返回，使用前请先调用Init方法
func FailWithErrCode(c *gin.Context, code int) {
	c.JSON(
		http.StatusOK,
		gin.H{
			"code": code,
			"msg":  CustomCodeMsgMap[code],
		},
	)
}

// FailWithErrCodeAndMsg 错误码和信息返回
func FailWithErrCodeAndMsg(c *gin.Context, code int, msg string) {
	c.JSON(
		http.StatusOK,
		gin.H{
			"code": code,
			"msg":  msg,
		},
	)
}

// FailWithHttpStatus 状态码的返回
func FailWithHttpStatus(c *gin.Context, httpStatus int) {
	c.JSON(
		httpStatus,
		gin.H{},
	)
}

// FailWithHttpStatusAndCodeAndMsg 状态码、错误码、错误信息的返回
func FailWithHttpStatusAndCodeAndMsg(c *gin.Context, httpStatus int, code int, msg string) {
	c.JSON(
		httpStatus,
		gin.H{
			"code": code,
			"msg":  msg,
		},
	)
}

// FailWithError Error
func FailWithError(c *gin.Context, err error) {
	c.JSON(
		http.StatusOK,
		gin.H{
			"code": CodeError,
			"msg":  err.Error(),
		},
	)
	// 转换gorm错误
	//if errors.Is(err, gorm.ErrRecordNotFound) {
	//	FailWithError(c, NewErrorWithCode(RETCOD_NOT_EXISTS))
	//	return
	//}
	//if cErr, ok := err.(BizError); ok {
	//	Fail(c, cErr.Code(), cErr.Error())
	//	return
	//}
}

// JsonResult 返回结构
//type JsonResult struct {
//	Code int         `json:"code"`
//	Msg  string      `json:"msg,omitempty"`
//	Data interface{} `json:"data"`
//}
