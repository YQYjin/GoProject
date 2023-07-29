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
