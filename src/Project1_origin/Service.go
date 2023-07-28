package main

import (
	"errors"
	"sync"
)

// 页面信息
type PageInfo struct {
	Topic *Topic
	Posts []*Post
}

// 查询页面流,在流程中暂存信息,最后将信息打包成PageInfo后取出
type QueryPageInfoFlow struct {
	topicId  int64
	pageInfo *PageInfo
	topic    *Topic
	posts    []*Post
}

func ServiceQueryPageInfo(topicID int64) (*PageInfo, error) {
	return NewQueryPageInfoFlow(topicID).Do()

}

// 创建一个查询页面流并初始化话题ID
func NewQueryPageInfoFlow(topicID int64) *QueryPageInfoFlow {
	return &QueryPageInfoFlow{
		topicId: topicID,
	}
}

// 主要流程
func (f *QueryPageInfoFlow) Do() (*PageInfo, error) {
	if err := f.checkParam(); err != nil {
		return nil, err
	}
	if err := f.parperInfo(); err != nil {
		return nil, err
	}
	if err := f.packPageInfo(); err != nil {
		return nil, err
	}
	//返回打包后的信息
	return f.pageInfo, nil
}

// 检查参数是否正确
func (f *QueryPageInfoFlow) checkParam() error {
	if f.topicId <= 0 {
		return errors.New("话题ID应大于0")
	}
	return nil
}

// 准备信息,查询话题信息和帖子信息
func (f *QueryPageInfoFlow) parperInfo() error {
	println("话题ID为:", f.topicId)
	var waitgroup sync.WaitGroup
	waitgroup.Add(2)
	//查询话题信息
	go func() {
		defer waitgroup.Done()
		topic := NewTopicDaoInstance().QueryTopicById(f.topicId)
		println("话题信息为:", topic.Content)
		f.topic = topic
	}()
	//查询帖子信息
	go func() {
		defer waitgroup.Done()
		posts := NewPostDaoInstance().QueryPostByTopicId(f.topicId)

		print("帖子信息为:")
		for _, post := range posts {
			print(post.Content, " ")
		}
		println()
		f.posts = posts
	}()
	waitgroup.Wait()
	return nil
}
func (f *QueryPageInfoFlow) packPageInfo() error {
	//向页面信息流中添加PageInfo
	f.pageInfo = &PageInfo{
		Topic: f.topic,
		Posts: f.posts,
	}
	return nil
}
