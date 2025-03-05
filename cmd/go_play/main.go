package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"sync"
	"time"

	"github.com/blevesearch/bleve/v2"
)

type ConfigChangeEventDocument struct {
	Id         uint64    `json:"detailId"`
	ChangeTime time.Time `json:"changeTime"`
	Type       int       `json:"type"`
	Content    string    `json:"content"`
}

// 定义事件类型
type Event struct {
	Data string
}

// 定义事件总线（通过 Channel 实现）
type EventBus struct {
	subscribers map[chan Event]bool
	lock        sync.RWMutex
}

// 创建事件总线
func NewEventBus() *EventBus {
	return &EventBus{
		subscribers: make(map[chan Event]bool),
	}
}

// 订阅事件
func (b *EventBus) Subscribe() chan Event {
	b.lock.Lock()
	defer b.lock.Unlock()

	ch := make(chan Event, 10) // 缓冲 Channel，避免阻塞
	b.subscribers[ch] = true
	return ch
}

// 取消订阅
func (b *EventBus) Unsubscribe(ch chan Event) {
	b.lock.Lock()
	defer b.lock.Unlock()

	if _, ok := b.subscribers[ch]; ok {
		close(ch) // 关闭 Channel
		delete(b.subscribers, ch)
	}
}

// 发布事件
func (b *EventBus) Publish(event Event) {
	b.lock.RLock()
	defer b.lock.RUnlock()

	for ch := range b.subscribers {
		ch <- event // 发送事件到所有订阅者
	}
}

func main() {
	// // 创建事件总线
	// eventBus := NewEventBus()

	// // 订阅者 1
	// subscriber1 := eventBus.Subscribe()
	// go func() {
	// 	for event := range subscriber1 {
	// 		fmt.Printf("Subscriber 1 received event: %s\n", event.Data)
	// 	}
	// 	fmt.Println("Subscriber 1 stopped")
	// }()

	// // 订阅者 2
	// subscriber2 := eventBus.Subscribe()
	// go func() {
	// 	for event := range subscriber2 {
	// 		fmt.Printf("Subscriber 2 received event: %s\n", event.Data)
	// 	}
	// 	fmt.Println("Subscriber 2 stopped")
	// }()

	// // 发布事件
	// eventBus.Publish(Event{Data: "Hello, World!"})
	// time.Sleep(100 * time.Millisecond) // 等待订阅者处理

	// // 取消订阅者 1
	// eventBus.Unsubscribe(subscriber1)

	// // 再次发布事件
	// eventBus.Publish(Event{Data: "Goodbye!"})
	// time.Sleep(100 * time.Millisecond) // 等待订阅者处理

	// // 关闭事件总线
	// eventBus.Unsubscribe(subscriber2)
	m, e := regexp.MatchString("wish_photo/list", "/hacking-tree/v1/internal/admin/wish_photo/list?sign=3295eb4ef750ef9a023043aa8f472778")
	if e != nil {
		return
	}
	fmt.Println(m)
}

const (
	indexPath = "config_change_events.bleve"
)

func prepareIndex() {
	// 创建自定义索引映射
	mapping := bleve.NewIndexMapping()

	// 为 ConfigChangeEventDocument 创建文档映射
	docMapping := bleve.NewDocumentMapping()

	// 配置字段映射
	contentField := bleve.NewTextFieldMapping()
	contentField.Analyzer = "standard" // 使用标准分词器
	contentField.Store = true          // 确保存储字段内容
	contentField.Index = true          // 确保可索引

	// 添加字段映射
	docMapping.AddFieldMappingsAt("content", contentField)

	// 设置默认文档映射
	mapping.AddDocumentMapping("_default", docMapping)

	// 删除索引文件（如果存在）以确保重新创建
	os.RemoveAll(indexPath)

	// 创建新索引
	index, err := bleve.New(indexPath, mapping)
	if err != nil {
		log.Fatalf("无法创建索引: %v", err)
	}
	defer index.Close()

	// 示例数据
	events := []ConfigChangeEventDocument{
		{
			Id:         1,
			ChangeTime: time.Now(),
			Type:       1,
			Content:    content1,
		},
		{
			Id:         2,
			ChangeTime: time.Now().Add(-24 * time.Hour),
			Type:       2,
			Content:    content2,
		},
		{
			Id:         3,
			ChangeTime: time.Now().Add(-48 * time.Hour),
			Type:       2,
			Content:    "abc",
		},
		{
			Id:         4,
			ChangeTime: time.Now().AddDate(0, 0, -2),
			Type:       1,
			Content:    "这个",
		},
		{
			Id:         5,
			ChangeTime: time.Now(),
			Type:       1,
			Content:    "已可以",
		},
	}
	// 索引文档
	for _, event := range events {
		if err := index.Index(fmt.Sprintf("%d", event.Id), event); err != nil {
			log.Printf("索引文档 %d 失败: %v", event.Id, err)
		}
	}
}

func execQuery() {
	index, err := bleve.Open(indexPath)
	if err != nil {
		return
	}
	defer index.Close()

	// 打印索引的详细信息
	indexMapping := index.Mapping()
	fmt.Println("索引映射详情:")
	fmt.Printf("索引映射: %+v\n", indexMapping)

	// 执行搜索
	query := bleve.NewMatchPhraseQuery("过期可以")

	searchRequest := bleve.NewSearchRequest(query)
	searchRequest.SortBy([]string{"-_score", "-changeTime"})
	results, err := index.Search(searchRequest)
	if err != nil {
		log.Fatalf("搜索失败: %v", err)
	}

	fmt.Println("搜索结果:")
	for _, hit := range results.Hits {
		fmt.Printf("文档ID: %s, 相关性得分: %f\n", hit.ID, hit.Score)
	}
	fmt.Printf("总命中数: %d\n", results.Total)
	fmt.Printf("查询: %v\n", query)
}
