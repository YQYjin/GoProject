package main

// 页面
type PageInfo struct {
	Topic *Topic
	Posts []*Post
}

type QueryPageInfoFlow struct {
	topicId  int64
	pageInfo *PageInfo

	topic *Topic
	posts []*Post
}

func (f *QueryPageInfoFlow) Do(*PageInfo, error) {

}
