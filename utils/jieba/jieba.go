package main

import (
	"fmt"
	"strings"

	"github.com/yanyiwu/gojieba"
)

// 中文分词
func main() {
	var s string
	var words []string
	useHmm := true
	x := gojieba.NewJieba()
	//defer x.Free()

	s = "我来到北京清华大学"
	words = x.CutAll(s)
	fmt.Println(s)
	fmt.Println("全模式:", strings.Join(words, "/"))

	words = x.Cut(s, useHmm)
	fmt.Println(s)
	fmt.Println("精确模式:", strings.Join(words, "/"))
	s = "比特币"
	words = x.Cut(s, useHmm)
	fmt.Println(s)
	fmt.Println("精确模式:", strings.Join(words, "/"))

	x.AddWord("比特币")
	s = "比特币"
	words = x.Cut(s, useHmm)
	fmt.Println(s)
	fmt.Println("添加词典后,精确模式:", strings.Join(words, "/"))

	s = "他来到了网易杭研大厦"
	words = x.Cut(s, useHmm)
	fmt.Println(s)
	fmt.Println("新词识别:", strings.Join(words, "/"))

	s = "小明硕士毕业于中国科学院计算所，后在日本京都大学深造"
	words = x.CutForSearch(s, useHmm)
	fmt.Println(s)
	fmt.Println("搜索引擎模式:", strings.Join(words, "/"))

	s = "长春市长春药店"
	words = x.Tag(s)
	fmt.Println(s)
	fmt.Println("词性标注:", strings.Join(words, ","))

	s = "区块链"
	words = x.Tag(s)
	fmt.Println(s)
	fmt.Println("词性标注:", strings.Join(words, ","))

	s = "长江大桥"
	words = x.CutForSearch(s, !useHmm)
	fmt.Println(s)
	fmt.Println("搜索引擎模式:", strings.Join(words, "/"))

	wordinfos := x.Tokenize(s, gojieba.SearchMode, !useHmm)
	fmt.Println(s)
	fmt.Println("Tokenize:(搜索引擎模式)", wordinfos)

	wordinfos = x.Tokenize(s, gojieba.DefaultMode, !useHmm)
	fmt.Println(s)
	fmt.Println("Tokenize:(默认模式)", wordinfos)

	keywords := x.ExtractWithWeight(s, 5)
	fmt.Println("Extract:", keywords)
}
