package bleve

import (
	"fmt"
	"log"
	"testing"

	"github.com/blevesearch/bleve/v2"
	"github.com/blevesearch/bleve/v2/analysis"
	"github.com/blevesearch/bleve/v2/analysis/analyzer/custom"
	"github.com/blevesearch/bleve/v2/analysis/tokenizer/unicode"
	"github.com/blevesearch/bleve/v2/registry"
	"github.com/go-ego/gse"
)

func TestBleveCase1(t *testing.T) {
	// 创建一个内存索引
	mapping := bleve.NewIndexMapping()

	// 添加自定义分词器
	err := mapping.AddCustomTokenizer("unicode", map[string]interface{}{
		"type": unicode.Name,
	})
	if err != nil {
		log.Fatal(err)
	}

	err = mapping.AddCustomAnalyzer("chinese", map[string]interface{}{
		"type":      custom.Name,
		"tokenizer": "unicode", // 使用 segment 分词器
	})
	if err != nil {
		log.Fatal(err)
	}

	mapping.DefaultAnalyzer = "chinese"

	index, err := bleve.NewMemOnly(mapping)
	if err != nil {
		log.Fatal(err)
	}

	// 索引文档
	doc1 := map[string]interface{}{
		"content": "我是一个好学生",
	}
	err = index.Index("1", doc1)
	if err != nil {
		log.Fatal(err)
	}

	doc2 := map[string]interface{}{
		"content": "我才是一个坏学生",
	}
	err = index.Index("2", doc2)
	if err != nil {
		log.Fatal(err)
	}

	// 查询 "好学生"
	query := bleve.NewMatchQuery("好学生")
	searchRequest := bleve.NewSearchRequest(query)
	searchResult, err := index.Search(searchRequest)
	if err != nil {
		log.Fatal(err)
	}

	// 输出查询结果
	fmt.Printf("查询 '好学生' 的结果: %v\n", searchResult)
}

// 注册自定义分词器
func init() {
	registry.RegisterTokenizer("gse", func(config map[string]interface{}, cache *registry.Cache) (analysis.Tokenizer, error) {
		return NewGseTokenizer(), nil
	})
}

// GseTokenizer 是一个自定义的分词器，基于 gse
type GseTokenizer struct {
	segmenter *gse.Segmenter
}

func NewGseTokenizer() *GseTokenizer {
	segmenter := &gse.Segmenter{}
	segmenter.LoadDict() // 加载默认字典
	return &GseTokenizer{
		segmenter: segmenter,
	}
}

func (t *GseTokenizer) Tokenize(input []byte) analysis.TokenStream {
	// 使用 gse 分词
	segments := t.segmenter.Cut(string(input), true)

	// 将分词结果转换为 analysis.Token
	tokens := make(analysis.TokenStream, 0, len(segments))
	for _, seg := range segments {
		tokens = append(tokens, &analysis.Token{
			Term: []byte(seg),
		})
	}
	return tokens
}

func TestBleveCase2(t *testing.T) {
	// 创建自定义分词器
	mapping := bleve.NewIndexMapping()

	// // 注册自定义分词器
	// err := mapping.AddCustomTokenizer("gse", map[string]interface{}{
	// 	"type": "gse",
	// })
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// // 注册自定义分析器
	// err = mapping.AddCustomAnalyzer("chinese", map[string]interface{}{
	// 	"type":      custom.Name,
	// 	"tokenizer": "gse", // 使用自定义的 gse 分词器
	// })
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// mapping.DefaultAnalyzer = "chinese"

	// 创建索引
	index, err := bleve.NewMemOnly(mapping)
	if err != nil {
		log.Fatal(err)
	}

	// 索引文档
	doc1 := map[string]interface{}{
		"content": "我是一个好学生",
	}
	err = index.Index("1", doc1)
	if err != nil {
		log.Fatal(err)
	}

	doc2 := map[string]interface{}{
		"content": "我是一个坏学生",
	}
	err = index.Index("2", doc2)
	if err != nil {
		log.Fatal(err)
	}

	// 查询 "好学生"
	query := bleve.NewMatchPhraseQuery("好学生")
	searchRequest := bleve.NewSearchRequest(query)
	searchResult, err := index.Search(searchRequest)
	if err != nil {
		log.Fatal(err)
	}

	// 输出查询结果
	fmt.Printf("查询 '好学生' 的结果: %v\n", searchResult)
}
