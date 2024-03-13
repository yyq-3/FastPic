package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yyq-3/FastPic/pic"
)

func main() {
	defer handlerErr()

	r := gin.Default()
	picGroup := r.Group("/pic")
	{
		picGroup.POST("/cutTo1600x900", pic.CutTo1600x900)
		picGroup.POST("/cutTo800x600", pic.CutTo800x600)
		picGroup.POST("/customerCut", pic.CustomerCut)
	}
	r.Run(":8081")
}

// global err handler
func handlerErr() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// 记录日志或进行其他必要的错误处理
				// log.Printf("Panic: %v", err)

				// 返回错误信息给客户端
				c.JSON(http.StatusInternalServerError, gin.H{
					"message": "Internal Server Error",
				})

				// 防止继续执行后续的处理函数
				c.Abort()
			}
		}()

		// 调用该请求的剩余处理程序
		c.Next()
	}
}