package project1_hw

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
func PackPostInfo(topicId string, content string, create_time string) error {
	//将传入的字符串id转换为int64
	topicId_int, err := strconv.ParseInt(topicId, 10, 64)
	if err != nil {
		println(err.Error())
		return err
	}
	//将时间字符串转为int64
	create_time_int, err := strconv.ParseInt(create_time, 10, 64)
	if err != nil {
		println(err.Error())
		return err
	}
	post := &Post{
		TopicId: topicId_int,
		Content: content,
		Date:    create_time_int,
	}
	//添加帖子ID(服务层方法)并进行下一步操作
	err = paperPostInfo(post)
	if err != nil {
		println(err.Error())
		return err
	}
	return nil
}
