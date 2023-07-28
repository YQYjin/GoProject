package main

import (
	"strconv"
)

// 页面数据,Code为状态码
type PageData struct {
	Code int64       `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func QueryPageInfo(topicIDStr string) *PageData {
	//将传入的字符串id转换为int64
	topicID, err := strconv.ParseInt(topicIDStr, 10, 64)
	if err != nil {
		return &PageData{
			Code: -1,
			Msg:  err.Error(),
			Data: nil,
		}
	}
	//调用service层的查询方法,查询页面
	pageInfo, err := ServiceQueryPageInfo(topicID)
	if err != nil {
		return &PageData{
			Code: -1,
			Msg:  err.Error(),
			Data: nil,
		}
	}
	return &PageData{
		Code: 0,
		Msg:  "success",
		Data: pageInfo,
	}
}

type postFlow struct {
}

// 整合帖子数据
func packPostInfo(topicId int64, content string, create_time int64) error {
	post := &Post{
		TopicId: topicId,
		Content: content,
		Date:    create_time,
	}
	//添加帖子ID(服务层方法)并进行下一步操作
	paperPostInfo(post)
	return nil
}
