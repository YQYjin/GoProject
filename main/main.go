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
	//设置静态文件路径,防止跨域错误
	r.Static("/assets", "../assets")
	//配置路由，即访问网址，:topicId为参数
	//该网址接收到get请求时的操作
	r.GET("topic/:topicId", func(c *gin.Context) {
		topicId := c.Param("topicId")
		data := QueryPageInfo(topicId)
		c.JSON(200, data)
	})
	//定义post方法,接收发送的帖子
	r.POST("post", func(c *gin.Context) {
		topicId := c.PostForm("topicId")
		content := c.PostForm("content")
		create_time := c.PostForm("create_time")
		//使用协程进行添加,防止阻塞
		go func() {
			err := packPostInfo(topicId, content, create_time)
			if err != nil {
				return
			}
		}()
		println("接收到的帖子信息为:", topicId, " ", content, " ", create_time)
		c.JSON(200, "success")
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
