package main

//数据层

import (
	"bufio"
	"encoding/json"
	"os"
	"sync"
)

// 定义结构体
type Topic struct {
	Id      int64  `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Date    int64  `json:"date"`
}
type Post struct {
	Id      int64  `json:"id"`
	TopicId int64  `json:"topic_id"`
	Content string `json:"content"`
	Date    int64  `json:"date"`
}

type PostDao struct {
}

// 使用sync.once实现单例
var (
	postDao  *PostDao
	postOnce sync.Once
)

func NewPostDaoInstance() *PostDao {
	postOnce.Do(func() {
		postDao = &PostDao{}
	})
	return postDao
}

// 使用postDao实现单例查询,提供效率
func (*PostDao) QueryPostByTopicId(topicId int64) []*Post {
	return postIndexMap[topicId]
}

// 同理
type TopicDao struct{}

var (
	topicDao  *TopicDao
	topicOnce sync.Once
)

func NewTopicDaoInstance() *TopicDao {
	topicOnce.Do(func() {
		topicDao = &TopicDao{}
	})
	return topicDao
}
func (*TopicDao) QueryTopicById(topicId int64) *Topic {
	return topicIndexMap[topicId]
}

// 索引,需求是通过话题ID获取话题信息和话题下所有帖子
var (
	topicIndexMap map[int64]*Topic
	postIndexMap  map[int64][]*Post
)

// ------数据层方法
// 初始化数据层
func Init(filePath string) error {
	if err := initTopicIndexMap(filePath); err != nil {
		return err
	}
	if err := initPostIndexMap(filePath); err != nil {
		return err
	}
	return nil
}

// 初始化话题索引
func initTopicIndexMap(filePath string) error {
	open, err := os.Open(filePath + "topic")
	if err != nil {
		return err
	}
	scanner := bufio.NewScanner(open)
	topicTmpMap := make(map[int64]*Topic)
	//读取文件
	for scanner.Scan() {
		text := scanner.Text()
		var topic Topic
		if err := json.Unmarshal([]byte(text), &topic); err != nil {
			return err
		}
		topicTmpMap[topic.Id] = &topic
	}
	topicIndexMap = topicTmpMap
	return nil
}

// 初始化帖子引索
func initPostIndexMap(filePath string) error {
	open, err := os.Open(filePath + "post")
	if err != nil {
		return err
	}
	scanner := bufio.NewScanner(open)
	postTmpMap := make(map[int64][]*Post)
	for scanner.Scan() {
		text := scanner.Text()
		var post Post
		if err := json.Unmarshal([]byte(text), &post); err != nil {
			return err
		}
		//将帖子放入对应话题下的列表中
		postTmpMap[post.TopicId] = append(postTmpMap[post.TopicId], &post)
	}
	postIndexMap = postTmpMap
	return nil
}
