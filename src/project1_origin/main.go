package main

import (
	"github.com/gin-gonic/gin"
	"os"
)

func main() {
	//初始化数据
	if err := mainInit("./data/"); err != nil {
		println(err.Error())
		os.Exit(-1)
	}
	//通过gin框架提供的http服务
	r := gin.Default()
	//注册路由
	r.GET("/topic/:topicId", func(c *gin.Context) {
		topicId := c.Param("topicId")
		data := QueryPageInfo(topicId)
		c.JSON(200, data)
	})
	//启动服务
	err := r.Run()
	if err != nil {
		return
	}
}
func mainInit(filePath string) error {
	if err := Init(filePath); err != nil {
		return err
	}
	return nil
}
